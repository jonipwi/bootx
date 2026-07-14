package engine

import (
	"strings"
	"testing"
)

func TestAnalyzeLocalDocumentFindsReviewStructure(t *testing.T) {
	content := `# Study

## Evidence

- [ ] Verify the primary source.
- [x] Preserve the correction.

Evidence is not established for the main causal claim.
See https://example.invalid/source.
`
	summary := analyzeLocalDocument(content)
	if summary.Headings != 2 || summary.ExternalLinks != 1 || summary.OpenChecklist != 1 {
		t.Fatalf("unexpected summary: %+v", summary)
	}
	if summary.EvidenceGapLines < 1 || !strings.Contains(summary.FirstGap, "not established") {
		t.Fatalf("evidence gap not detected: %+v", summary)
	}
}

func TestAnalyzeLocalDocumentDoesNotInventGap(t *testing.T) {
	summary := analyzeLocalDocument("# Complete note\n\nThe cited source was checked on the recorded date.\n")
	if summary.EvidenceGapLines != 0 || summary.FirstGap != "" {
		t.Fatalf("unexpected gap: %+v", summary)
	}
}
