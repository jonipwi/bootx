package evidence

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadLocalDocumentReturnsIntegrityReceipt(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "study.md")
	content := []byte("# Study\n\nVerify the source before deciding.\n")
	if err := os.WriteFile(path, content, 0o600); err != nil {
		t.Fatal(err)
	}

	document, err := LoadLocalDocument(root, "study.md")
	if err != nil {
		t.Fatal(err)
	}
	want := sha256.Sum256(content)
	if document.Reference != "study.md" || document.Content != string(content) {
		t.Fatalf("unexpected document receipt: %+v", document)
	}
	if document.SHA256 != hex.EncodeToString(want[:]) || document.ByteLength != len(content) {
		t.Fatal("document hash or byte length does not match selected bytes")
	}
}

func TestLoadLocalDocumentRejectsTraversal(t *testing.T) {
	parent := t.TempDir()
	root := filepath.Join(parent, "workspace")
	if err := os.Mkdir(root, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(parent, "outside.md"), []byte("outside"), 0o600); err != nil {
		t.Fatal(err)
	}

	_, err := LoadLocalDocument(root, filepath.Join("..", "outside.md"))
	if err == nil || !strings.Contains(err.Error(), "inside the workspace") {
		t.Fatalf("expected containment rejection, got %v", err)
	}
}

func TestLoadLocalDocumentRejectsLinkEscapeWhenSupported(t *testing.T) {
	parent := t.TempDir()
	root := filepath.Join(parent, "workspace")
	if err := os.Mkdir(root, 0o700); err != nil {
		t.Fatal(err)
	}
	outside := filepath.Join(parent, "outside.md")
	if err := os.WriteFile(outside, []byte("outside"), 0o600); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(root, "linked.md")
	if err := os.Symlink(outside, link); err != nil {
		t.Skipf("file symlinks are unavailable in this test environment: %v", err)
	}
	if _, err := LoadLocalDocument(root, "linked.md"); err == nil || !strings.Contains(err.Error(), "inside the workspace") {
		t.Fatalf("expected resolved-link containment rejection, got %v", err)
	}
}

func TestLoadLocalDocumentRejectsAbsolutePath(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "study.md")
	if err := os.WriteFile(path, []byte("study"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadLocalDocument(root, path); err == nil || !strings.Contains(err.Error(), "relative") {
		t.Fatalf("expected absolute-path rejection, got %v", err)
	}
}

func TestLoadLocalDocumentRejectsUnsupportedOrOversizedFiles(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "program.exe"), []byte("text"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadLocalDocument(root, "program.exe"); err == nil || !strings.Contains(err.Error(), "not allowed") {
		t.Fatalf("expected extension rejection, got %v", err)
	}

	large := strings.Repeat("a", MaxDocumentBytes+1)
	if err := os.WriteFile(filepath.Join(root, "large.txt"), []byte(large), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadLocalDocument(root, "large.txt"); err == nil || !strings.Contains(err.Error(), "exceeds") {
		t.Fatalf("expected size rejection, got %v", err)
	}
}

func TestLoadLocalDocumentRejectsInvalidText(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "bad.txt"), []byte{0xff, 0x00}, 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadLocalDocument(root, "bad.txt"); err == nil || !strings.Contains(err.Error(), "UTF-8") {
		t.Fatalf("expected text rejection, got %v", err)
	}
}
