package engine

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/jonipwi/bootx/prototype/personal-companion/internal/model"
)

type localDocumentSummary struct {
	Headings         int
	ExternalLinks    int
	OpenChecklist    int
	EvidenceGapLines int
	FirstGap         string
}

var evidenceGapMarkers = []string{
	"not established",
	"not verified",
	"not validated",
	"unvalidated",
	"evidence needed",
	"evidence still needed",
	"requires review",
	"require review",
	"still require",
	"missing evidence",
	"missing fact",
	"unknown",
	"todo",
	"tbd",
}

func analyzeLocalDocument(content string) localDocumentSummary {
	var summary localDocumentSummary
	summary.ExternalLinks = strings.Count(strings.ToLower(content), "https://") + strings.Count(strings.ToLower(content), "http://")
	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		lower := strings.ToLower(trimmed)
		if strings.HasPrefix(trimmed, "# ") || strings.HasPrefix(trimmed, "## ") || strings.HasPrefix(trimmed, "### ") || strings.HasPrefix(trimmed, "#### ") || strings.HasPrefix(trimmed, "##### ") || strings.HasPrefix(trimmed, "###### ") {
			summary.Headings++
		}
		if strings.HasPrefix(lower, "- [ ]") || strings.HasPrefix(lower, "* [ ]") {
			summary.OpenChecklist++
		}
		if trimmed != "" && containsAnyMarker(lower, evidenceGapMarkers) {
			summary.EvidenceGapLines++
			if summary.FirstGap == "" {
				summary.FirstGap = truncateRunes(strings.TrimSpace(strings.TrimLeft(trimmed, "#-* >")), 240)
			}
		}
	}
	return summary
}

func appendLocalDocumentFindings(packet *model.Packet, summary localDocumentSummary) {
	packet.Observations = append(packet.Observations, model.Finding{
		Claim: fmt.Sprintf(
			"Deterministic document scan found %d heading(s), %d external-link string(s), %d open checklist item(s), and %d line(s) with evidence-gap markers.",
			summary.Headings,
			summary.ExternalLinks,
			summary.OpenChecklist,
			summary.EvidenceGapLines,
		),
		Source: "deterministic_local_document_scan",
		Status: "structural_observation_not_semantic_verdict",
	})
	if summary.FirstGap != "" {
		packet.Observations = append(packet.Observations, model.Finding{
			Claim:  "First evidence-gap candidate: " + summary.FirstGap,
			Source: "deterministic_local_document_scan",
			Status: "candidate_requires_human_review",
		})
	}
}

func localDocumentOptions(summary localDocumentSummary) []model.Option {
	gapSummary := "Review the first detected evidence-gap candidate against its surrounding context."
	if summary.FirstGap == "" {
		gapSummary = "Manually identify one material claim whose evidence maturity is not clear."
	}
	return []model.Option{
		{OptionID: "review-evidence-gap", Summary: gapSummary, Benefits: []string{"focuses review on a visible, reversible evidence task"}, Risks: []string{"keyword markers may miss or misclassify context"}, Reversibility: "easy", ExternalEffect: "none"},
		{OptionID: "verify-cited-source", Summary: "Choose one material cited source and independently verify that it supports the nearby claim.", Benefits: []string{"tests claim-to-source alignment"}, Risks: []string{"one source cannot validate the whole document"}, Reversibility: "easy", ExternalEffect: "none"},
		{OptionID: "review-open-checklist", Summary: "Review an open checklist item and confirm its owner, evidence, and completion rule.", Benefits: []string{"turns an inherited task into an auditable next step"}, Risks: []string{"the document may contain no open checkbox or may be stale"}, Reversibility: "easy", ExternalEffect: "none"},
		{OptionID: "defer", Summary: "Defer changes until a qualified reviewer or better source is available.", Benefits: []string{"avoids unsupported edits"}, Risks: []string{"important correction may be delayed"}, Reversibility: "easy", ExternalEffect: "none"},
	}
}

func containsAnyMarker(text string, markers []string) bool {
	for _, marker := range markers {
		if strings.Contains(text, marker) {
			return true
		}
	}
	return false
}

func truncateRunes(value string, maximum int) string {
	if utf8.RuneCountInString(value) <= maximum {
		return value
	}
	runes := []rune(value)
	return string(runes[:maximum]) + "..."
}
