package ethicalreview

import (
	"strings"
	"testing"
	"time"
)

func TestEvaluateComputesTransparentRiskIndex(t *testing.T) {
	request := validRequest()
	request.ContentType = "public_statement"
	request.Claims = []EvidenceClaim{
		{ID: "c1", Text: "Supported high-consequence claim", SourceReference: "primary://one", SourceStatus: "primary_confirmed", Consequence: "high"},
		{ID: "c2", Text: "Unsupported irreversible claim", SourceStatus: "unverified", Consequence: "irreversible"},
		{ID: "c3", Text: "Disputed moderate claim", SourceReference: "public://dispute", SourceStatus: "disputed", Consequence: "moderate"},
	}
	result, err := Evaluate(request)
	if err != nil {
		t.Fatal(err)
	}
	if result.EvidenceCoverage != 41.67 || result.HighConsequenceCoverage != 50 ||
		result.UncertaintyGap != 33.33 || result.ContestedRate != 33.33 ||
		result.ConsequenceExposure != 75 || result.ReviewRiskIndex != 50.42 {
		t.Fatalf("unexpected formula result: %+v", result)
	}
	if result.WarningLevel != "W4_STOP" || len(result.HardStops) != 1 {
		t.Fatalf("expected hard-stop warning, got %+v", result)
	}
}

func TestEvaluateNeverApprovesPublication(t *testing.T) {
	result, err := Evaluate(validRequest())
	if err != nil {
		t.Fatal(err)
	}
	if result.WarningLevel != "W1_REVIEW" || result.DecisionPosture != "CONTINUE_HUMAN_REVIEW" {
		t.Fatalf("unexpected posture: %+v", result)
	}
	if strings.Contains(strings.ToLower(result.DecisionPosture), "approve") {
		t.Fatal("preflight must never approve publication")
	}
}

func TestEvaluateLegalDraftRequiresQualifiedReview(t *testing.T) {
	request := validRequest()
	request.ContentType = "legal_reasoning_draft"
	result, err := Evaluate(request)
	if err != nil {
		t.Fatal(err)
	}
	if result.WarningLevel != "W4_STOP" || len(result.HardStops) == 0 {
		t.Fatalf("legal safeguard missing: %+v", result)
	}
}

func TestHighConsequenceClaimNeverReturnsLowReviewLevel(t *testing.T) {
	request := validRequest()
	request.Claims[0].Consequence = "high"
	result, err := Evaluate(request)
	if err != nil {
		t.Fatal(err)
	}
	if result.WarningLevel != "W2_VERIFY" {
		t.Fatalf("high-consequence declared evidence must still be verified, got %+v", result)
	}
}

func TestValidateRequiresExplicitRemoteConsent(t *testing.T) {
	request := validRequest()
	request.RemoteProcessingConsent = false
	_, err := Evaluate(request)
	if err == nil || !strings.Contains(err.Error(), "remote_processing_consent") {
		t.Fatalf("expected remote consent failure, got %v", err)
	}
}

func validRequest() Request {
	return Request{
		RequestID:                   "review-test",
		CapabilityID:                CapabilityID,
		UserID:                      "local-user",
		CreatedAt:                   time.Date(2026, 7, 25, 12, 0, 0, 0, time.UTC),
		ContentType:                 "social_post",
		Purpose:                     "Review before publication.",
		Audience:                    "Public community",
		Context:                     "Synthetic educational example.",
		DraftText:                   "A careful statement supported by a declared primary source.",
		Claims:                      []EvidenceClaim{{ID: "c1", Text: "A declared claim.", SourceReference: "primary://example", SourceStatus: "primary_confirmed", Consequence: "low"}},
		PublicNonSensitiveConfirmed: true,
		RemoteProcessingConsent:     true,
		HumanAuthorityConfirmed:     true,
	}
}
