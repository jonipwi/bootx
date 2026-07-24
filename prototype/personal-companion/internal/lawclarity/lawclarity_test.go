package lawclarity

import (
	"strings"
	"testing"
	"time"
)

func TestEvaluateGrayZoneExample(t *testing.T) {
	request := exampleRequest()
	report, err := Evaluate(request)
	if err != nil {
		t.Fatal(err)
	}
	if report.LawQualityScore != 35.75 {
		t.Fatalf("quality score = %v, want 35.75", report.LawQualityScore)
	}
	if report.GrayZoneRiskScore != 74.25 {
		t.Fatalf("gray-zone score = %v, want 74.25", report.GrayZoneRiskScore)
	}
	if report.ManipulationRiskIndex != 48.11 {
		t.Fatalf("manipulation index = %v, want 48.11", report.ManipulationRiskIndex)
	}
	if report.HumanRightsGate.Status != "FAIL" || report.StrictGoodLawGate.Status != "FAIL" {
		t.Fatalf("unexpected gates: rights=%s strict=%s", report.HumanRightsGate.Status, report.StrictGoodLawGate.Status)
	}
	if !report.HighManipulationTrigger || report.Disposition != "FUNDAMENTAL_REVISION_REQUIRED" {
		t.Fatalf("unexpected disposition/trigger: %s %t", report.Disposition, report.HighManipulationTrigger)
	}
	if len(report.VisiblePhraseHits) != 2 {
		t.Fatalf("phrase hits = %#v", report.VisiblePhraseHits)
	}
	if report.UserDecision != nil || report.InputReceipt.SourceStatus != "user_supplied_reference_not_authenticated" {
		t.Fatal("human authority or source-authentication boundary missing")
	}
}

func TestEvaluateStrongRatingsStillRequiresQualifiedReview(t *testing.T) {
	request := exampleRequest()
	request.ClauseText = "An officer shall issue a written notice when documented evidence establishes an immediate risk. The affected person may appeal within seven days."
	request.Quality = quality(90, 90, 90, 90, 90, 90)
	request.GrayZone = gray(10, 10, 10, 10, 10)
	request.Power = PowerContext{
		PowerConcentrationRisk: rated(10, "Decision and review functions are separated."),
		OversightStrength:      rated(90, "Independent review is available."),
	}
	report, err := Evaluate(request)
	if err != nil {
		t.Fatal(err)
	}
	if report.StrictGoodLawGate.Status != "PASS" || report.HumanRightsGate.Status != "PASS" {
		t.Fatalf("unexpected gates: %+v %+v", report.StrictGoodLawGate, report.HumanRightsGate)
	}
	if report.Disposition != "QUALIFIED_REVIEW_REQUIRED" {
		t.Fatalf("disposition = %s", report.Disposition)
	}
	if len(report.VisiblePhraseHits) != 0 {
		t.Fatalf("unexpected phrase hits: %#v", report.VisiblePhraseHits)
	}
	if report.VisiblePhraseHits == nil {
		t.Fatal("visible_phrase_hits must encode as an empty array, not null")
	}
	if report.Assurance[1].Status != "PASS" || report.Assurance[5].Status != "PASS" {
		t.Fatal("runtime assurance must evaluate the screening process, not certify or grade the reviewed law")
	}
}

func TestEvaluateRejectsUnsafeOrIncompleteInput(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Request)
		want   string
	}{
		{name: "confirmation", mutate: func(r *Request) { r.PublicNonSensitiveConfirmed = false }, want: "public_non_sensitive_confirmed"},
		{name: "capability", mutate: func(r *Request) { r.CapabilityID = "wrong" }, want: "capability_id"},
		{name: "score", mutate: func(r *Request) { r.Quality.Clarity.Score = 101 }, want: "between 0 and 100"},
		{name: "rationale", mutate: func(r *Request) { r.Quality.Clarity.Rationale = "" }, want: "rationale is required"},
		{name: "clause", mutate: func(r *Request) { r.ClauseText = "" }, want: "clause_text is required"},
		{name: "type", mutate: func(r *Request) { r.InstrumentType = "criminal_verdict" }, want: "unsupported instrument_type"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := exampleRequest()
			test.mutate(&request)
			_, err := Evaluate(request)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want %q", err, test.want)
			}
		})
	}
}

func exampleRequest() Request {
	return Request{
		RequestID:                   "law-test-001",
		CapabilityID:                CapabilityID,
		UserID:                      "declared-local-reviewer",
		CreatedAt:                   time.Date(2026, 7, 24, 12, 0, 0, 0, time.UTC),
		Title:                       "Public-order clause example",
		Jurisdiction:                "educational fictional example",
		InstrumentType:              "regulation",
		SourceReference:             "synthetic://law-clarity-example",
		Purpose:                     "Identify ambiguity and rights-review questions.",
		ClauseText:                  "Authorities may take appropriate action against persons whose conduct may disturb public interest.",
		PublicNonSensitiveConfirmed: true,
		Quality:                     quality(45, 25, 50, 35, 30, 20),
		GrayZone:                    gray(80, 85, 40, 90, 75),
		Power: PowerContext{
			PowerConcentrationRisk: rated(90, "One authority defines and enforces the clause."),
			OversightStrength:      rated(20, "No independent review is stated."),
		},
	}
}

func quality(c, s, f, i, a, l int) QualityRatings {
	return QualityRatings{
		Clarity:         rated(c, "Reviewer rationale for clarity."),
		Specificity:     rated(s, "Reviewer rationale for specificity."),
		Fairness:        rated(f, "Reviewer rationale for fairness."),
		Consistency:     rated(i, "Reviewer rationale for consistency."),
		Accountability:  rated(a, "Reviewer rationale for accountability."),
		LowLoopholeRisk: rated(l, "Reviewer rationale for loophole exposure."),
	}
}

func gray(v, d, x, e, u int) GrayZoneRatings {
	return GrayZoneRatings{
		VagueLanguageRisk:     rated(v, "Reviewer rationale for vague language."),
		DefinitionRisk:        rated(d, "Reviewer rationale for definitions."),
		ContradictionRisk:     rated(x, "Reviewer rationale for contradictions."),
		EnforcementDiscretion: rated(e, "Reviewer rationale for discretion."),
		ExceptionBoundaryRisk: rated(u, "Reviewer rationale for exceptions."),
	}
}

func rated(score int, rationale string) Rating {
	return Rating{Score: score, Rationale: rationale}
}
