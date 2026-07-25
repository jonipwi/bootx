package ethicalreview

import (
	"fmt"
	"math"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	CapabilityID  = "assist.ethical-review.v1"
	MaxDraftBytes = 16_384
	MaxClaims     = 25
)

type EvidenceClaim struct {
	ID              string `json:"id"`
	Text            string `json:"text"`
	SourceReference string `json:"source_reference"`
	SourceStatus    string `json:"source_status"`
	Consequence     string `json:"consequence"`
}

type Request struct {
	RequestID                   string          `json:"request_id"`
	CapabilityID                string          `json:"capability_id"`
	UserID                      string          `json:"user_id"`
	CreatedAt                   time.Time       `json:"created_at"`
	ContentType                 string          `json:"content_type"`
	Purpose                     string          `json:"purpose"`
	Audience                    string          `json:"audience"`
	Context                     string          `json:"context"`
	DraftText                   string          `json:"draft_text"`
	Claims                      []EvidenceClaim `json:"claims"`
	PublicNonSensitiveConfirmed bool            `json:"public_non_sensitive_confirmed"`
	RemoteProcessingConsent     bool            `json:"remote_processing_consent"`
	HumanAuthorityConfirmed     bool            `json:"human_authority_confirmed"`
}

type EvidenceMetric struct {
	Name           string  `json:"name"`
	Value          float64 `json:"value"`
	Unit           string  `json:"unit"`
	Interpretation string  `json:"interpretation"`
}

type Preflight struct {
	Formulae                  []string         `json:"formulae"`
	ClaimCount                int              `json:"claim_count"`
	HighConsequenceClaimCount int              `json:"high_consequence_claim_count"`
	Metrics                   []EvidenceMetric `json:"metrics"`
	EvidenceCoverage          float64          `json:"evidence_coverage"`
	HighConsequenceCoverage   float64          `json:"high_consequence_coverage"`
	UncertaintyGap            float64          `json:"uncertainty_gap"`
	ContestedRate             float64          `json:"contested_rate"`
	ConsequenceExposure       float64          `json:"consequence_exposure"`
	ReviewRiskIndex           float64          `json:"review_risk_index"`
	WarningLevel              string           `json:"warning_level"`
	DecisionPosture           string           `json:"decision_posture"`
	HardStops                 []string         `json:"hard_stops"`
	Limitations               []string         `json:"limitations"`
}

type StatementReview struct {
	Excerpt        string `json:"excerpt"`
	Classification string `json:"classification"`
	SupportStatus  string `json:"support_status"`
	Reason         string `json:"reason"`
}

type Finding struct {
	Category  string `json:"category"`
	Severity  string `json:"severity"`
	TextBasis string `json:"text_basis"`
	Concern   string `json:"concern"`
	Repair    string `json:"repair"`
}

type Rewrite struct {
	Provided      bool     `json:"provided"`
	Draft         string   `json:"draft"`
	ChangeSummary []string `json:"change_summary"`
}

type Advice struct {
	Summary               string            `json:"summary"`
	StatementReviews      []StatementReview `json:"statement_reviews"`
	Findings              []Finding         `json:"findings"`
	MissingPerspectives   []string          `json:"missing_perspectives"`
	Counterarguments      []string          `json:"counterarguments"`
	QuestionsBeforeAction []string          `json:"questions_before_action"`
	SuggestedPosture      string            `json:"suggested_posture"`
	Rewrite               Rewrite           `json:"rewrite"`
	Limitations           []string          `json:"limitations"`
}

type ProviderResult struct {
	Status  string  `json:"status"`
	Advice  *Advice `json:"advice,omitempty"`
	Refusal string  `json:"refusal,omitempty"`
}

type RemoteReceipt struct {
	Provider                string   `json:"provider"`
	API                     string   `json:"api"`
	ModelRequested          string   `json:"model_requested"`
	ModelReturned           string   `json:"model_returned,omitempty"`
	ResponseID              string   `json:"response_id,omitempty"`
	ProviderRequestID       string   `json:"provider_request_id,omitempty"`
	StoreRequested          bool     `json:"store_requested"`
	ApplicationPersistence  bool     `json:"application_persistence"`
	ConversationStateUsed   bool     `json:"conversation_state_used"`
	ToolsEnabled            bool     `json:"tools_enabled"`
	ExternalActionsEnabled  bool     `json:"external_actions_enabled"`
	RawDraftSent            bool     `json:"raw_draft_sent"`
	SentFields              []string `json:"sent_fields"`
	ProviderRetentionNotice string   `json:"provider_retention_notice"`
}

type Envelope struct {
	RequestID              string         `json:"request_id"`
	CapabilityID           string         `json:"capability_id"`
	GeneratedAt            time.Time      `json:"generated_at"`
	RuntimeNotice          string         `json:"runtime_notice"`
	DeterministicPreflight Preflight      `json:"deterministic_preflight"`
	OpenAIAdvisory         ProviderResult `json:"openai_advisory"`
	RemoteReceipt          RemoteReceipt  `json:"remote_receipt"`
	BlockedActions         []string       `json:"blocked_actions"`
	UserDecision           *string        `json:"user_decision"`
}

func Evaluate(request Request) (Preflight, error) {
	if err := Validate(request); err != nil {
		return Preflight{}, err
	}

	var evidenceTotal, highEvidenceTotal, consequenceTotal float64
	var highCount, uncertainCount, contestedCount int
	var hardStops []string
	for _, claim := range request.Claims {
		evidence := evidenceWeight(claim.SourceStatus)
		consequence := consequenceWeight(claim.Consequence)
		evidenceTotal += evidence
		consequenceTotal += consequence
		if claim.Consequence == "high" || claim.Consequence == "irreversible" {
			highCount++
			highEvidenceTotal += evidence
			if claim.Consequence == "irreversible" {
				hardStops = append(hardStops, fmt.Sprintf("claim %q has a declared irreversible consequence and requires qualified review regardless of declared source status", claim.ID))
			} else if claim.SourceStatus == "none" || claim.SourceStatus == "unverified" || claim.SourceStatus == "disputed" {
				hardStops = append(hardStops, fmt.Sprintf("claim %q is high-consequence but lacks strong declared evidence", claim.ID))
			}
		}
		if claim.SourceStatus == "none" || claim.SourceStatus == "unverified" {
			uncertainCount++
		}
		if claim.SourceStatus == "disputed" {
			contestedCount++
		}
	}

	claimCount := len(request.Claims)
	evidenceCoverage := 100.0
	uncertaintyGap := 0.0
	contestedRate := 0.0
	consequenceExposure := 0.0
	if claimCount > 0 {
		evidenceCoverage = 100 * evidenceTotal / float64(claimCount)
		uncertaintyGap = 100 * float64(uncertainCount) / float64(claimCount)
		contestedRate = 100 * float64(contestedCount) / float64(claimCount)
		consequenceExposure = 100 * consequenceTotal / float64(claimCount)
	}
	highCoverage := 100.0
	if highCount > 0 {
		highCoverage = 100 * highEvidenceTotal / float64(highCount)
	}
	risk := 0.35*(100-evidenceCoverage) +
		0.25*(100-highCoverage) +
		0.20*uncertaintyGap +
		0.10*contestedRate +
		0.10*consequenceExposure

	if request.ContentType == "legal_reasoning_draft" {
		hardStops = append(hardStops, "legal reasoning requires qualified independent review; BootX and OpenAI cannot determine guilt, liability, or sentence")
	}

	warning, posture := warningFor(risk, len(hardStops) > 0, highCount > 0)
	evidenceCoverage = round2(evidenceCoverage)
	highCoverage = round2(highCoverage)
	uncertaintyGap = round2(uncertaintyGap)
	contestedRate = round2(contestedRate)
	consequenceExposure = round2(consequenceExposure)
	risk = round2(risk)

	return Preflight{
		Formulae: []string{
			"E = 100 * sum(declared source weights) / N; primary=1.00, secondary=0.75, disputed=0.25, unverified/none=0",
			"H = 100 * sum(declared source weights for high or irreversible claims) / Nh; H=100 when Nh=0",
			"U = 100 * count(unverified or unsourced claims) / N",
			"C = 100 * count(disputed claims) / N",
			"I = 100 * mean(consequence weights); low=0.25, moderate=0.50, high=0.75, irreversible=1.00",
			"R = 0.35(100-E) + 0.25(100-H) + 0.20U + 0.10C + 0.10I",
		},
		ClaimCount:                claimCount,
		HighConsequenceClaimCount: highCount,
		Metrics: []EvidenceMetric{
			{Name: "Evidence coverage", Value: evidenceCoverage, Unit: "index_points", Interpretation: "Declared source support only; not truth verification."},
			{Name: "High-consequence coverage", Value: highCoverage, Unit: "index_points", Interpretation: "Declared support for high or irreversible claims."},
			{Name: "Uncertainty gap", Value: uncertaintyGap, Unit: "index_points", Interpretation: "Share of claims declared unverified or without a source."},
			{Name: "Contested rate", Value: contestedRate, Unit: "index_points", Interpretation: "Share of claims declared disputed."},
			{Name: "Consequence exposure", Value: consequenceExposure, Unit: "index_points", Interpretation: "Declared consequence-weight profile."},
		},
		EvidenceCoverage:        evidenceCoverage,
		HighConsequenceCoverage: highCoverage,
		UncertaintyGap:          uncertaintyGap,
		ContestedRate:           contestedRate,
		ConsequenceExposure:     consequenceExposure,
		ReviewRiskIndex:         risk,
		WarningLevel:            warning,
		DecisionPosture:         posture,
		HardStops:               hardStops,
		Limitations: []string{
			"All source statuses and consequence levels are user declarations; BootX does not fetch or authenticate sources in this capability.",
			"R is a transparent review-priority index, not a probability of harm, truth, justice, or future outcome.",
			"OpenAI findings are advisory model output and can be incomplete, biased, or mistaken.",
			"No score or rewrite authorizes publication, punishment, accusation, enforcement, or another external action.",
		},
	}, nil
}

func Validate(request Request) error {
	if request.CapabilityID != CapabilityID {
		return fmt.Errorf("capability_id must be %q", CapabilityID)
	}
	required := []struct {
		name  string
		value string
		max   int
	}{
		{name: "request_id", value: request.RequestID, max: 128},
		{name: "user_id", value: request.UserID, max: 128},
		{name: "purpose", value: request.Purpose, max: 1024},
		{name: "audience", value: request.Audience, max: 512},
		{name: "context", value: request.Context, max: 4096},
	}
	for _, field := range required {
		if strings.TrimSpace(field.value) == "" {
			return fmt.Errorf("%s is required", field.name)
		}
		if len([]byte(field.value)) > field.max {
			return fmt.Errorf("%s exceeds %d bytes", field.name, field.max)
		}
	}
	if request.CreatedAt.IsZero() {
		return fmt.Errorf("created_at is required")
	}
	switch request.ContentType {
	case "social_post", "speech", "public_statement", "proposal", "decision_rationale", "legal_reasoning_draft", "other":
	default:
		return fmt.Errorf("unsupported content_type %q", request.ContentType)
	}
	if strings.TrimSpace(request.DraftText) == "" {
		return fmt.Errorf("draft_text is required")
	}
	if !utf8.ValidString(request.DraftText) {
		return fmt.Errorf("draft_text must be valid UTF-8")
	}
	if len([]byte(request.DraftText)) > MaxDraftBytes {
		return fmt.Errorf("draft_text exceeds %d bytes", MaxDraftBytes)
	}
	if len(request.Claims) > MaxClaims {
		return fmt.Errorf("claims exceeds %d items", MaxClaims)
	}
	seen := map[string]bool{}
	for index, claim := range request.Claims {
		if strings.TrimSpace(claim.ID) == "" || len([]byte(claim.ID)) > 64 {
			return fmt.Errorf("claims[%d].id is required and must not exceed 64 bytes", index)
		}
		if seen[claim.ID] {
			return fmt.Errorf("duplicate claim id %q", claim.ID)
		}
		seen[claim.ID] = true
		if strings.TrimSpace(claim.Text) == "" || len([]byte(claim.Text)) > 1024 {
			return fmt.Errorf("claims[%d].text is required and must not exceed 1024 bytes", index)
		}
		switch claim.SourceStatus {
		case "primary_confirmed", "secondary_confirmed":
			if strings.TrimSpace(claim.SourceReference) == "" {
				return fmt.Errorf("claims[%d].source_reference is required for declared confirmed evidence", index)
			}
		case "disputed", "unverified", "none":
		default:
			return fmt.Errorf("unsupported claims[%d].source_status %q", index, claim.SourceStatus)
		}
		if len([]byte(claim.SourceReference)) > 2048 {
			return fmt.Errorf("claims[%d].source_reference exceeds 2048 bytes", index)
		}
		switch claim.Consequence {
		case "low", "moderate", "high", "irreversible":
		default:
			return fmt.Errorf("unsupported claims[%d].consequence %q", index, claim.Consequence)
		}
	}
	if !request.PublicNonSensitiveConfirmed {
		return fmt.Errorf("public_non_sensitive_confirmed must be true; private, sensitive, or uncertain material is not authorized for remote review")
	}
	if !request.RemoteProcessingConsent {
		return fmt.Errorf("remote_processing_consent must be true for OpenAI review")
	}
	if !request.HumanAuthorityConfirmed {
		return fmt.Errorf("human_authority_confirmed must be true; BootX does not make or execute the decision")
	}
	return nil
}

func ValidateAdvice(advice Advice) error {
	switch advice.SuggestedPosture {
	case "continue_human_review", "revise_before_human_review", "delay_and_verify", "do_not_publish_as_written", "seek_qualified_review":
	default:
		return fmt.Errorf("unsupported suggested_posture %q", advice.SuggestedPosture)
	}
	for index, statement := range advice.StatementReviews {
		switch statement.Classification {
		case "fact_claim", "inference", "opinion", "value_judgment", "question", "unclear":
		default:
			return fmt.Errorf("statement_reviews[%d] has unsupported classification %q", index, statement.Classification)
		}
		switch statement.SupportStatus {
		case "declared_supported", "declared_partial", "declared_unsupported", "declared_disputed", "not_assessable":
		default:
			return fmt.Errorf("statement_reviews[%d] has unsupported support_status %q", index, statement.SupportStatus)
		}
	}
	for index, finding := range advice.Findings {
		switch finding.Category {
		case "evidence", "logic", "fairness", "compassion", "uncertainty", "foreseeable_harm", "privacy", "due_process":
		default:
			return fmt.Errorf("findings[%d] has unsupported category %q", index, finding.Category)
		}
		switch finding.Severity {
		case "low", "medium", "high", "critical":
		default:
			return fmt.Errorf("findings[%d] has unsupported severity %q", index, finding.Severity)
		}
	}
	return nil
}

func evidenceWeight(status string) float64 {
	switch status {
	case "primary_confirmed":
		return 1
	case "secondary_confirmed":
		return 0.75
	case "disputed":
		return 0.25
	default:
		return 0
	}
}

func consequenceWeight(consequence string) float64 {
	switch consequence {
	case "low":
		return 0.25
	case "moderate":
		return 0.50
	case "high":
		return 0.75
	case "irreversible":
		return 1
	default:
		return 0
	}
}

func warningFor(risk float64, hardStop, highConsequence bool) (string, string) {
	if hardStop {
		return "W4_STOP", "STOP_AND_SEEK_QUALIFIED_REVIEW"
	}
	if risk >= 60 {
		return "W3_REVISE", "REVISE_BEFORE_HUMAN_REVIEW"
	}
	if risk >= 35 || highConsequence {
		return "W2_VERIFY", "VERIFY_BEFORE_HUMAN_REVIEW"
	}
	return "W1_REVIEW", "CONTINUE_HUMAN_REVIEW"
}

func round2(value float64) float64 {
	return math.Round(value*100) / 100
}
