package evidence

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

const MaxDocumentBytes = 65_536

var allowedExtensions = map[string]bool{
	".json": true,
	".md":   true,
	".txt":  true,
}

// LocalDocument is a read-only receipt for bytes deliberately selected from a
// contained workspace. Integrity proves which bytes were processed; it does
// not authenticate the author, publisher, or truth of the content.
type LocalDocument struct {
	Reference  string
	Content    string
	SHA256     string
	ByteLength int
	ModifiedAt time.Time
}

func LoadLocalDocument(workspaceRoot, relativePath string) (LocalDocument, error) {
	if strings.TrimSpace(workspaceRoot) == "" {
		return LocalDocument{}, fmt.Errorf("workspace root is required")
	}
	if strings.TrimSpace(relativePath) == "" {
		return LocalDocument{}, fmt.Errorf("document path is required")
	}
	if filepath.IsAbs(relativePath) {
		return LocalDocument{}, fmt.Errorf("document path must be relative to the workspace root")
	}

	rootAbs, err := filepath.Abs(workspaceRoot)
	if err != nil {
		return LocalDocument{}, fmt.Errorf("resolve workspace root: %w", err)
	}
	rootReal, err := filepath.EvalSymlinks(rootAbs)
	if err != nil {
		return LocalDocument{}, fmt.Errorf("resolve workspace root links: %w", err)
	}
	rootInfo, err := os.Stat(rootReal)
	if err != nil {
		return LocalDocument{}, fmt.Errorf("inspect workspace root: %w", err)
	}
	if !rootInfo.IsDir() {
		return LocalDocument{}, fmt.Errorf("workspace root is not a directory")
	}

	candidate := filepath.Join(rootReal, filepath.Clean(relativePath))
	candidateReal, err := filepath.EvalSymlinks(candidate)
	if err != nil {
		return LocalDocument{}, fmt.Errorf("resolve selected document: %w", err)
	}
	reference, err := containedReference(rootReal, candidateReal)
	if err != nil {
		return LocalDocument{}, err
	}
	if !allowedExtensions[strings.ToLower(filepath.Ext(candidateReal))] {
		return LocalDocument{}, fmt.Errorf("document type is not allowed; use .md, .txt, or .json")
	}

	file, err := os.Open(candidateReal)
	if err != nil {
		return LocalDocument{}, fmt.Errorf("open selected document read-only: %w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return LocalDocument{}, fmt.Errorf("inspect selected document: %w", err)
	}
	if !info.Mode().IsRegular() {
		return LocalDocument{}, fmt.Errorf("selected document must be a regular file")
	}
	if info.Size() > MaxDocumentBytes {
		return LocalDocument{}, fmt.Errorf("selected document exceeds %d bytes", MaxDocumentBytes)
	}

	content, err := io.ReadAll(io.LimitReader(file, MaxDocumentBytes+1))
	if err != nil {
		return LocalDocument{}, fmt.Errorf("read selected document: %w", err)
	}
	if len(content) > MaxDocumentBytes {
		return LocalDocument{}, fmt.Errorf("selected document exceeds %d bytes", MaxDocumentBytes)
	}
	afterRead, err := file.Stat()
	if err != nil {
		return LocalDocument{}, fmt.Errorf("reinspect selected document: %w", err)
	}
	if afterRead.Size() != int64(len(content)) || afterRead.Size() != info.Size() || !afterRead.ModTime().Equal(info.ModTime()) {
		return LocalDocument{}, fmt.Errorf("selected document changed while it was being read; retry after writes stop")
	}
	if !utf8.Valid(content) || bytes.IndexByte(content, 0) >= 0 {
		return LocalDocument{}, fmt.Errorf("selected document must be UTF-8 text without NUL bytes")
	}

	digest := sha256.Sum256(content)
	return LocalDocument{
		Reference:  filepath.ToSlash(reference),
		Content:    string(content),
		SHA256:     hex.EncodeToString(digest[:]),
		ByteLength: len(content),
		ModifiedAt: afterRead.ModTime().UTC(),
	}, nil
}

func containedReference(root, selected string) (string, error) {
	relative, err := filepath.Rel(root, selected)
	if err != nil {
		return "", fmt.Errorf("compare workspace and document paths: %w", err)
	}
	if relative == "." || relative == "" || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("selected document must remain inside the workspace root")
	}
	return relative, nil
}
