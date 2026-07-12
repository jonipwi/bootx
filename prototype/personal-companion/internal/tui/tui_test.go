package tui

import (
	"bytes"
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
