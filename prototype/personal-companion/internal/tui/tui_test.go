package tui

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jonipwi/bootx/prototype/personal-companion/internal/engine"
)

func TestScriptedPersonalWorkflow(t *testing.T) {
	e, err := engine.New()
	if err != nil {
		t.Fatal(err)
	}
	input := strings.Join([]string{
		"1",
		"Choose a study plan",
		"What fact should I verify first?",
		"2",
		"2",
		"n",
		".",
		"",
		"truth, learning",
		"y",
		"n",
		"q",
		"",
	}, "\n")
	var output bytes.Buffer
	if err := Run(strings.NewReader(input), &output, e); err != nil {
		t.Fatal(err)
	}
	text := output.String()
	for _, required := range []string{"BOOTX PERSONAL DECISION PACKET", "Class: D0", "AI DNA RUNTIME CHECKS", "no external actions", "Session closed"} {
		if !strings.Contains(text, required) {
			t.Fatalf("output missing %q", required)
		}
	}
}

func TestScriptedLocalDocumentWorkflow(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "work.md"), []byte("# Work\n\nOne public task needs evidence.\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	e, err := engine.New()
	if err != nil {
		t.Fatal(err)
	}
	input := strings.Join([]string{
		"3",
		root,
		"work.md",
		"y",
		"Choose the next evidence task",
		"Which missing fact should I verify first?",
		"truth, reversibility",
		"y",
		"n",
		"q",
		"",
	}, "\n")
	var output bytes.Buffer
	if err := Run(strings.NewReader(input), &output, e); err != nil {
		t.Fatal(err)
	}
	text := output.String()
	for _, required := range []string{"Read-only integrity receipt", "work.md", "Integrity verified: true", "Origin: not_authenticated", "Class: D0", "Session closed"} {
		if !strings.Contains(text, required) {
			t.Fatalf("output missing %q\n%s", required, text)
		}
	}
}

func TestScriptedLawClarityWorkflow(t *testing.T) {
	e, err := engine.New()
	if err != nil {
		t.Fatal(err)
	}
	input := strings.Join([]string{
		"4",
		"y",
		"Public-order clause example",
		"fictional educational jurisdiction",
		"2",
		"synthetic://law-clarity-example",
		"Identify ambiguity and rights-review questions",
		"Authorities may take appropriate action against persons whose conduct may disturb public interest.",
		".",
		"45", "The operative terms are not defined.",
		"25", "Actors, conduct, evidence, and boundaries are incomplete.",
		"50", "Notice, appeal, and remedy are not stated.",
		"35", "Different authorities could reach different results.",
		"30", "Written reasons and independent review are absent.",
		"20", "Open wording permits broad exceptions.",
		"80", "Appropriate and public interest are materially vague here.",
		"85", "Material terms have no definitions.",
		"40", "No direct contradiction is visible in this excerpt.",
		"90", "The authority receives broad unexplained discretion.",
		"75", "No exception boundary or expiry is stated.",
		"90", "One authority appears able to interpret and enforce.",
		"20", "No independent oversight is stated.",
		"y",
		"n",
		"q",
		"",
	}, "\n")
	var output bytes.Buffer
	if err := Run(strings.NewReader(input), &output, e); err != nil {
		t.Fatal(err)
	}
	text := output.String()
	for _, required := range []string{
		"BOOTX LAW CLARITY LOGIC SCREENING REPORT",
		"LAW QUALITY: 35.75/100",
		"GRAY-ZONE RISK: 74.25/100",
		"STRICT GATE: FAIL | RIGHTS GATE: FAIL",
		"FUNDAMENTAL_REVISION_REQUIRED",
		"No legal decision was made",
		"Session closed",
	} {
		if !strings.Contains(text, required) {
			t.Fatalf("output missing %q\n%s", required, text)
		}
	}
}
