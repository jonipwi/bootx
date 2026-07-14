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
