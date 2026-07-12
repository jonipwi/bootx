package warning

import (
	"testing"

	"github.com/jonipwi/bootx/prototype/personal-companion/internal/model"
)

func TestEvaluateLevels(t *testing.T) {
	tests := []struct {
		name  string
		input model.WarningInput
		want  string
	}{
		{name: "direct danger overrides degraded feed", input: baseInput(func(in *model.WarningInput) {
			in.DirectDanger = true
			in.IntegrityStatus = "fail"
		}), want: W4},
		{name: "authenticated immediate alert", input: baseInput(func(in *model.WarningInput) {
			in.OfficialStatus = "active"
			in.AuthorityAuthenticated = true
			in.Urgency = "Immediate"
			in.Severity = "Severe"
			in.Certainty = "Likely"
		}), want: W4},
		{name: "authenticated severe expected alert", input: baseInput(func(in *model.WarningInput) {
			in.OfficialStatus = "active"
			in.AuthorityAuthenticated = true
			in.Urgency = "Expected"
			in.Severity = "Severe"
		}), want: W3},
		{name: "corroborated forecast prepare", input: baseInput(func(in *model.WarningInput) {
			in.EvidenceTier = "V2"
			in.AreaMatch = "near"
		}), want: W2},
		{name: "one source monitor", input: baseInput(func(in *model.WarningInput) {
			in.EvidenceTier = "V1"
		}), want: W1},
		{name: "conflict degraded", input: baseInput(func(in *model.WarningInput) {
			in.SourceConflict = true
		}), want: WX},
		{name: "test is not actual alert", input: baseInput(func(in *model.WarningInput) {
			in.OfficialStatus = "active"
			in.AuthorityAuthenticated = true
			in.MessageStatus = "Test"
			in.Urgency = "Immediate"
		}), want: W0},
		{name: "outside configured area", input: baseInput(func(in *model.WarningInput) {
			in.OfficialStatus = "active"
			in.AuthorityAuthenticated = true
			in.AreaMatch = "outside"
			in.Urgency = "Immediate"
		}), want: W0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Evaluate(test.input)
			if err != nil {
				t.Fatalf("Evaluate() error = %v", err)
			}
			if result.Card.Level != test.want {
				t.Fatalf("level = %s, want %s", result.Card.Level, test.want)
			}
		})
	}
}

func TestEvaluateRejectsUnknownCAPValue(t *testing.T) {
	in := baseInput(nil)
	in.Urgency = "VerySoon"
	if _, err := Evaluate(in); err == nil {
		t.Fatal("expected unsupported urgency to fail closed")
	}
}

func baseInput(mutate func(*model.WarningInput)) model.WarningInput {
	in := model.WarningInput{
		EventID:         "event-1",
		HazardType:      "flood",
		OfficialStatus:  "none_found",
		Issuer:          "example authority",
		MessageStatus:   "Actual",
		MessageType:     "Alert",
		AreaMatch:       "inside",
		Urgency:         "Future",
		Severity:        "Moderate",
		Certainty:       "Possible",
		EvidenceTier:    "V0",
		IntegrityStatus: "pass",
	}
	if mutate != nil {
		mutate(&in)
	}
	return in
}
