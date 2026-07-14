package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/jonipwi/bootx/prototype/personal-companion/internal/engine"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/evidence"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/model"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/tui"
)

const maxJSONInput = 1 << 20

func main() {
	inputPath := flag.String("input", "", "strict JSON request path, or - for stdin; omit for TUI")
	compact := flag.Bool("compact", false, "emit compact JSON in backend mode")
	showVersion := flag.Bool("version", false, "print version and exit")
	workspaceRoot := flag.String("workspace", "", "workspace root for contained read-only document mode")
	documentPath := flag.String("document", "", "relative .md, .txt, or .json path inside -workspace")
	documentPublic := flag.Bool("document-public", false, "confirm the selected document was reviewed as public and non-sensitive")
	goal := flag.String("goal", "", "decision goal for read-only document mode")
	question := flag.String("question", "", "direct question for read-only document mode")
	priorities := flag.String("priorities", "", "optional comma-separated priorities for read-only document mode")
	flag.Parse()

	if *showVersion {
		fmt.Println(model.Version)
		return
	}
	decisionEngine, err := engine.New()
	if err != nil {
		fatal(err)
	}
	if *documentPath != "" {
		if *inputPath != "" {
			fatal(fmt.Errorf("-input and -document are mutually exclusive"))
		}
		request, err := buildDocumentRequest(documentRequestOptions{
			WorkspaceRoot:   *workspaceRoot,
			DocumentPath:    *documentPath,
			PublicConfirmed: *documentPublic,
			Goal:            *goal,
			Question:        *question,
			Priorities:      *priorities,
		})
		if err != nil {
			fatal(err)
		}
		packet, err := decisionEngine.Process(request)
		if err != nil {
			fatal(err)
		}
		encodePacket(packet, *compact)
		return
	}
	if *workspaceRoot != "" || *documentPublic || *goal != "" || *question != "" || *priorities != "" {
		fatal(fmt.Errorf("document-mode flags require -document"))
	}
	if *inputPath == "" {
		if err := tui.Run(os.Stdin, os.Stdout, decisionEngine); err != nil {
			fatal(err)
		}
		return
	}

	request, err := readRequest(*inputPath)
	if err != nil {
		fatal(err)
	}
	packet, err := decisionEngine.Process(request)
	if err != nil {
		fatal(err)
	}
	encodePacket(packet, *compact)
}

type documentRequestOptions struct {
	WorkspaceRoot   string
	DocumentPath    string
	PublicConfirmed bool
	Goal            string
	Question        string
	Priorities      string
}

func buildDocumentRequest(options documentRequestOptions) (model.Request, error) {
	if !options.PublicConfirmed {
		return model.Request{}, fmt.Errorf("read-only document mode requires -document-public after reviewing the file as public and non-sensitive")
	}
	if strings.TrimSpace(options.Goal) == "" || strings.TrimSpace(options.Question) == "" {
		return model.Request{}, fmt.Errorf("read-only document mode requires -goal and -question")
	}
	document, err := evidence.LoadLocalDocument(options.WorkspaceRoot, options.DocumentPath)
	if err != nil {
		return model.Request{}, err
	}
	return model.Request{
		RequestID:       newDocumentRequestID(),
		CapabilityID:    model.CapabilityID,
		UserID:          "declared-local-operator",
		CreatedAt:       time.Now(),
		Goal:            strings.TrimSpace(options.Goal),
		Question:        strings.TrimSpace(options.Question),
		SelectedContent: document.Content,
		ContentSource: model.ContentSource{
			Type:              "local_workspace_file",
			OriginVerified:    false,
			Reference:         document.Reference,
			IntegrityVerified: true,
			SHA256:            document.SHA256,
			ByteLength:        document.ByteLength,
			ModifiedAt:        document.ModifiedAt.Format(time.RFC3339Nano),
		},
		DataClass:        model.DataPublic,
		DeclaredDomain:   model.DomainStudy,
		MemoryPermission: "none",
		RemotePermission: "deny",
		OutputPreference: "standard",
		UserPriorities:   splitPriorities(options.Priorities),
		Synthetic:        false,
	}, nil
}

func splitPriorities(input string) []string {
	var result []string
	seen := map[string]bool{}
	for _, value := range strings.Split(input, ",") {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	return result
}

func newDocumentRequestID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err == nil {
		return "local-doc-" + hex.EncodeToString(b)
	}
	return fmt.Sprintf("local-doc-%d", time.Now().UnixNano())
}

func encodePacket(packet model.Packet, compact bool) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(true)
	if !compact {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(packet); err != nil {
		fatal(err)
	}
}

func readRequest(path string) (model.Request, error) {
	var source io.Reader
	var file *os.File
	if path == "-" {
		source = os.Stdin
	} else {
		opened, err := os.Open(path)
		if err != nil {
			return model.Request{}, err
		}
		file = opened
		defer file.Close()
		source = file
	}
	decoder := json.NewDecoder(io.LimitReader(source, maxJSONInput))
	decoder.DisallowUnknownFields()
	var request model.Request
	if err := decoder.Decode(&request); err != nil {
		return model.Request{}, fmt.Errorf("decode strict request JSON: %w", err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return model.Request{}, fmt.Errorf("request must contain exactly one JSON object")
	}
	return request, nil
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "bootx-companion:", err)
	os.Exit(1)
}
