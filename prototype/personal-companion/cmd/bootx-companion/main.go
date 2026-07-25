package main

import (
	"context"
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
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/ethicalreview"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/evidence"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/lawclarity"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/model"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/openaiadvisory"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/tui"
)

const maxJSONInput = 1 << 20

func main() {
	inputPath := flag.String("input", "", "strict JSON request path, or - for stdin; omit for TUI")
	lawInputPath := flag.String("law-input", "", "strict Law Clarity JSON request path, or - for stdin")
	reviewInputPath := flag.String("review-input", "", "strict public ethical-review JSON request path, or - for stdin; calls OpenAI after explicit consent")
	openAIModel := flag.String("openai-model", "", "OpenAI model override for -review-input; defaults to BOOTX_OPENAI_MODEL or the documented current default")
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
	modeCount := 0
	if *inputPath != "" {
		modeCount++
	}
	if *lawInputPath != "" {
		modeCount++
	}
	if *reviewInputPath != "" {
		modeCount++
	}
	if *documentPath != "" {
		modeCount++
	}
	if modeCount > 1 {
		fatal(fmt.Errorf("-input, -law-input, -review-input, and -document are mutually exclusive"))
	}
	if *lawInputPath != "" {
		if *workspaceRoot != "" || *documentPublic || *goal != "" || *question != "" || *priorities != "" {
			fatal(fmt.Errorf("document-mode flags cannot be combined with -law-input"))
		}
		request, err := readLawRequest(*lawInputPath)
		if err != nil {
			fatal(err)
		}
		report, err := lawclarity.Evaluate(request)
		if err != nil {
			fatal(err)
		}
		encodeJSON(report, *compact)
		return
	}
	if *reviewInputPath != "" {
		if *workspaceRoot != "" || *documentPublic || *goal != "" || *question != "" || *priorities != "" {
			fatal(fmt.Errorf("document-mode flags cannot be combined with -review-input"))
		}
		request, err := readReviewRequest(*reviewInputPath)
		if err != nil {
			fatal(err)
		}
		preflight, err := ethicalreview.Evaluate(request)
		if err != nil {
			fatal(err)
		}
		client, err := openaiadvisory.NewFromEnvironment(*openAIModel)
		if err != nil {
			fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 65*time.Second)
		defer cancel()
		advisory, receipt, err := client.Review(ctx, request, preflight)
		if err != nil {
			fatal(err)
		}
		envelope := ethicalreview.Envelope{
			RequestID:              request.RequestID,
			CapabilityID:           ethicalreview.CapabilityID,
			GeneratedAt:            time.Now().UTC(),
			RuntimeNotice:          "Decision support only: declared-evidence mathematics plus bounded OpenAI review. This output is not fact verification, approval, legal advice, guilt, sentence, or authority to act.",
			DeterministicPreflight: preflight,
			OpenAIAdvisory:         advisory,
			RemoteReceipt:          receipt,
			BlockedActions: []string{
				"automatic publication, broadcast, messaging, or account action",
				"guilt, punishment, detention, legal sentence, or denial-of-rights decision",
				"scoring a person's worth or protected identity",
				"claiming that a source, statement, rewrite, or recommendation is verified true",
				"external execution without a separate informed human decision",
			},
			UserDecision: nil,
		}
		encodeJSON(envelope, *compact)
		return
	}
	decisionEngine, err := engine.New()
	if err != nil {
		fatal(err)
	}
	if *documentPath != "" {
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
		encodeJSON(packet, *compact)
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
	encodeJSON(packet, *compact)
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

func encodeJSON(value any, compact bool) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(true)
	if !compact {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(value); err != nil {
		fatal(err)
	}
}

func readRequest(path string) (model.Request, error) {
	return readStrictJSON[model.Request](path, "request")
}

func readLawRequest(path string) (lawclarity.Request, error) {
	return readStrictJSON[lawclarity.Request](path, "law-clarity request")
}

func readReviewRequest(path string) (ethicalreview.Request, error) {
	return readStrictJSON[ethicalreview.Request](path, "ethical-review request")
}

func readStrictJSON[T any](path, label string) (T, error) {
	var result T
	var source io.Reader
	var file *os.File
	if path == "-" {
		source = os.Stdin
	} else {
		opened, err := os.Open(path)
		if err != nil {
			return result, err
		}
		file = opened
		defer file.Close()
		source = file
	}
	decoder := json.NewDecoder(io.LimitReader(source, maxJSONInput))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&result); err != nil {
		return result, fmt.Errorf("decode strict %s JSON: %w", label, err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return result, fmt.Errorf("%s must contain exactly one JSON object", label)
	}
	return result, nil
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "bootx-companion:", err)
	os.Exit(1)
}
