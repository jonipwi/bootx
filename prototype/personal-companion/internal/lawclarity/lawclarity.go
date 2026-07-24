package lawclarity

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	CapabilityID       = "assist.law-clarity.v1"
	ScorePassThreshold = 60
	MaxClauseBytes     = 32_768
)

type Rating struct {
	Score     int    `json:"score"`
	Rationale string `json:"rationale"`
}

type QualityRatings struct {
	Clarity         Rating `json:"clarity"`
	Specificity     Rating `json:"specificity"`
	Fairness        Rating `json:"fairness"`
	Consistency     Rating `json:"consistency"`
	Accountability  Rating `json:"accountability"`
	LowLoopholeRisk Rating `json:"low_loophole_risk"`
}

type GrayZoneRatings struct {
	VagueLanguageRisk     Rating `json:"vague_language_risk"`
	DefinitionRisk        Rating `json:"definition_risk"`
	ContradictionRisk     Rating `json:"contradiction_risk"`
	EnforcementDiscretion Rating `json:"enforcement_discretion_risk"`
	ExceptionBoundaryRisk Rating `json:"exception_boundary_risk"`
}

type PowerContext struct {
	PowerConcentrationRisk Rating `json:"power_concentration_risk"`
	OversightStrength      Rating `json:"oversight_strength"`
}

type Request struct {
	RequestID                   string          `json:"request_id"`
	CapabilityID                string          `json:"capability_id"`
	UserID                      string          `json:"user_id"`
	CreatedAt                   time.Time       `json:"created_at"`
	Title                       string          `json:"title"`
	Jurisdiction                string          `json:"jurisdiction"`
	InstrumentType              string          `json:"instrument_type"`
	SourceReference             string          `json:"source_reference"`
	Purpose                     string          `json:"purpose"`
	ClauseText                  string          `json:"clause_text"`
	PublicNonSensitiveConfirmed bool            `json:"public_non_sensitive_confirmed"`
	Quality                     QualityRatings  `json:"quality"`
	GrayZone                    GrayZoneRatings `json:"gray_zone"`
	Power                       PowerContext    `json:"power"`
}

type ScoredCriterion struct {
	ID                   string  `json:"id"`
	Label                string  `json:"label"`
	Score                int     `json:"score"`
	Weight               float64 `json:"weight"`
	WeightedContribution float64 `json:"weighted_contribution"`
	Direction            string  `json:"direction"`
	Rationale            string  `json:"rationale"`
	ThresholdStatus      string  `json:"threshold_status"`
}

type GateResult struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Threshold string `json:"threshold"`
	Basis     string `json:"basis"`
}

type PhraseHit struct {
	Phrase string `json:"phrase"`
	Count  int    `json:"count"`
}

type ReviewSection struct {
	Stage     string   `json:"stage"`
	Questions []string `json:"questions"`
}

type AssuranceCheck struct {
	Dimension string `json:"dimension"`
	Status    string `json:"status"`
	Basis     string `json:"basis"`
}

type InputReceipt struct {
	SourceReference             string `json:"source_reference"`
	ClauseBytes                 int    `json:"clause_bytes"`
	PublicNonSensitiveConfirmed bool   `json:"public_non_sensitive_confirmed"`
	RemoteProcessing            bool   `json:"remote_processing"`
	PersistentMemory            bool   `json:"persistent_memory"`
	RawRetention                string `json:"raw_content_retention"`
	SourceStatus                string `json:"source_status"`
}

type Report struct {
	RequestID               string            `json:"request_id"`
	CapabilityID            string            `json:"capability_id"`
	GeneratedAt             time.Time         `json:"generated_at"`
	RuntimeNotice           string            `json:"runtime_notice"`
	Title                   string            `json:"title"`
	Jurisdiction            string            `json:"jurisdiction"`
	InstrumentType          string            `json:"instrument_type"`
	Formulae                []string          `json:"formulae"`
	QualityDimensions       []ScoredCriterion `json:"quality_dimensions"`
	LawQualityScore         float64           `json:"law_quality_score"`
	QualityBand             string            `json:"quality_band"`
	StrictGoodLawGate       GateResult        `json:"strict_good_law_gate"`
	HumanRightsGate         GateResult        `json:"human_rights_gate"`
	GrayZoneDimensions      []ScoredCriterion `json:"gray_zone_dimensions"`
	GrayZoneRiskScore       float64           `json:"gray_zone_risk_score"`
	ManipulationRiskIndex   float64           `json:"manipulation_risk_index"`
	HighManipulationTrigger bool              `json:"high_manipulation_trigger"`
	VisiblePhraseHits       []PhraseHit       `json:"visible_phrase_hits"`
	Findings                []string          `json:"findings"`
	ReviewWorkflow          []ReviewSection   `json:"review_workflow"`
	RewriteRequirements     []string          `json:"rewrite_requirements"`
	RewriteTemplate         string            `json:"rewrite_template"`
	Disposition             string            `json:"disposition"`
	BlockedConclusions      []string          `json:"blocked_conclusions"`
	Limitations             []string          `json:"limitations"`
	Assurance               []AssuranceCheck  `json:"ai_dna_runtime_checks"`
	InputReceipt            InputReceipt      `json:"input_receipt"`
	UserDecision            *string           `json:"user_decision"`
}

type criterionInput struct {
	id        string
	label     string
	weight    float64
	direction string
	rating    Rating
}

var reviewPhrases = []string{
	"reasonable",
	"appropriate",
	"improper",
	"disturbing",
	"public interest",
	"as necessary",
	"other relevant circumstances",
	"according to authority",
	"may take action",
}

func Evaluate(request Request) (Report, error) {
	if err := validate(request); err != nil {
		return Report{}, err
	}

	qualityInputs := []criterionInput{
		{id: "C", label: "Clear language", weight: 0.20, direction: "higher_is_better", rating: request.Quality.Clarity},
		{id: "S", label: "Specific definitions and boundaries", weight: 0.20, direction: "higher_is_better", rating: request.Quality.Specificity},
		{id: "F", label: "Fair and rights-protecting", weight: 0.20, direction: "higher_is_better", rating: request.Quality.Fairness},
		{id: "I", label: "Consistently enforceable", weight: 0.15, direction: "higher_is_better", rating: request.Quality.Consistency},
		{id: "A", label: "Accountable and auditable", weight: 0.15, direction: "higher_is_better", rating: request.Quality.Accountability},
		{id: "L", label: "Loophole risk acceptably low", weight: 0.10, direction: "higher_is_better", rating: request.Quality.LowLoopholeRisk},
	}
	grayInputs := []criterionInput{
		{id: "V", label: "Vague-word risk", weight: 0.30, direction: "higher_is_more_risk", rating: request.GrayZone.VagueLanguageRisk},
		{id: "D", label: "Undefined or circular-definition risk", weight: 0.25, direction: "higher_is_more_risk", rating: request.GrayZone.DefinitionRisk},
		{id: "X", label: "Contradictory-clause risk", weight: 0.20, direction: "higher_is_more_risk", rating: request.GrayZone.ContradictionRisk},
		{id: "E", label: "Enforcement-discretion risk", weight: 0.15, direction: "higher_is_more_risk", rating: request.GrayZone.EnforcementDiscretion},
		{id: "U", label: "Unclear-exception-boundary risk", weight: 0.10, direction: "higher_is_more_risk", rating: request.GrayZone.ExceptionBoundaryRisk},
	}

	qualityDimensions, qualityScore := scoreCriteria(qualityInputs)
	grayDimensions, grayScore := scoreCriteria(grayInputs)
	strictGate := strictGoodLawGate(qualityInputs)
	rightsGate := rightsGate(request.Quality.Fairness)
	phraseHits := scanPhrases(request.ClauseText)
	highManipulation := grayScore >= 60 &&
		request.Power.PowerConcentrationRisk.Score >= 60 &&
		request.Quality.Accountability.Score < 60
	manipulationIndex := round2(
		(grayScore / 100) *
			(float64(request.GrayZone.EnforcementDiscretion.Score) / 100) *
			(float64(request.Power.PowerConcentrationRisk.Score) / 100) *
			(1 - float64(request.Power.OversightStrength.Score)/100) *
			100,
	)

	report := Report{
		RequestID:      request.RequestID,
		CapabilityID:   CapabilityID,
		GeneratedAt:    time.Now().UTC(),
		RuntimeNotice:  "Deterministic legal-clarity screening only; reviewer-supplied ratings are not authenticated facts, legal advice, or a validity judgment.",
		Title:          strings.TrimSpace(request.Title),
		Jurisdiction:   strings.TrimSpace(request.Jurisdiction),
		InstrumentType: request.InstrumentType,
		Formulae: []string{
			"G = C AND S AND F AND I AND A AND L; operational screening pass requires every quality dimension >= 60",
			"Q = 0.20C + 0.20S + 0.20F + 0.15I + 0.15A + 0.10L",
			"Z = 0.30V + 0.25D + 0.20X + 0.15E + 0.10U",
			"M = (Z/100) * (E/100) * (P/100) * (1 - O/100) * 100",
		},
		QualityDimensions:       qualityDimensions,
		LawQualityScore:         qualityScore,
		QualityBand:             qualityBand(qualityScore),
		StrictGoodLawGate:       strictGate,
		HumanRightsGate:         rightsGate,
		GrayZoneDimensions:      grayDimensions,
		GrayZoneRiskScore:       grayScore,
		ManipulationRiskIndex:   manipulationIndex,
		HighManipulationTrigger: highManipulation,
		VisiblePhraseHits:       phraseHits,
		Findings:                findings(request, qualityScore, grayScore, strictGate, rightsGate, highManipulation, phraseHits),
		ReviewWorkflow:          standardReviewWorkflow(),
		RewriteRequirements:     rewriteRequirements(request, phraseHits),
		RewriteTemplate:         "An authorized [role] may [action] only when [defined evidence standard] establishes [specific condition], within [scope and duration]. The written decision must state [evidence and legal basis], provide [notice and opportunity to respond], identify [appeal deadline and independent reviewer], and specify [remedy for error or abuse].",
		Disposition:             disposition(qualityScore, grayScore, strictGate.Status == "PASS", rightsGate.Status == "PASS", highManipulation),
		BlockedConclusions: []string{
			"legal validity or constitutionality",
			"binding legal interpretation",
			"guilt, liability, eligibility, or entitlement",
			"authorization to enforce, punish, detain, discriminate, or deny rights",
			"replacement of a lawyer, court, regulator, legislature, or affected-community review",
		},
		Limitations: []string{
			"All numerical ratings and rationales are supplied by the reviewer and are not independently verified by BootX.",
			"The score bands and 60-point gates are research thresholds, not validated legal standards or probabilities.",
			"The visible phrase scan is literal and context-blind; a hit is a review prompt, while no hit is not proof of clarity.",
			"A clause excerpt may omit definitions, safeguards, remedies, or conflicts located elsewhere in the full instrument.",
			"Different jurisdictions, languages, rights frameworks, and legal traditions require qualified local analysis.",
			"The manipulation-risk index is an experimental multiplicative index, not a corruption forecast.",
		},
		Assurance: lawClarityAssurance(),
		InputReceipt: InputReceipt{
			SourceReference:             strings.TrimSpace(request.SourceReference),
			ClauseBytes:                 len([]byte(request.ClauseText)),
			PublicNonSensitiveConfirmed: request.PublicNonSensitiveConfirmed,
			RemoteProcessing:            false,
			PersistentMemory:            false,
			RawRetention:                "process_memory_only; discarded when the process exits",
			SourceStatus:                "user_supplied_reference_not_authenticated",
		},
		UserDecision: nil,
	}
	return report, nil
}

func validate(request Request) error {
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
		{name: "title", value: request.Title, max: 512},
		{name: "jurisdiction", value: request.Jurisdiction, max: 256},
		{name: "source_reference", value: request.SourceReference, max: 2048},
		{name: "purpose", value: request.Purpose, max: 2048},
	}
	for _, field := range required {
		value := strings.TrimSpace(field.value)
		if value == "" {
			return fmt.Errorf("%s is required", field.name)
		}
		if len([]byte(value)) > field.max {
			return fmt.Errorf("%s exceeds %d bytes", field.name, field.max)
		}
	}
	if request.CreatedAt.IsZero() {
		return fmt.Errorf("created_at is required")
	}
	if !request.PublicNonSensitiveConfirmed {
		return fmt.Errorf("public_non_sensitive_confirmed must be true; sensitive or uncertain legal material is not authorized")
	}
	switch request.InstrumentType {
	case "law", "regulation", "company_policy", "court_procedure", "contract", "other":
	default:
		return fmt.Errorf("unsupported instrument_type %q", request.InstrumentType)
	}
	if strings.TrimSpace(request.ClauseText) == "" {
		return fmt.Errorf("clause_text is required")
	}
	if len([]byte(request.ClauseText)) > MaxClauseBytes {
		return fmt.Errorf("clause_text exceeds %d bytes", MaxClauseBytes)
	}
	if !utf8.ValidString(request.ClauseText) || strings.ContainsRune(request.ClauseText, '\x00') {
		return fmt.Errorf("clause_text must be valid UTF-8 text without NUL bytes")
	}
	ratings := []struct {
		name   string
		rating Rating
	}{
		{name: "quality.clarity", rating: request.Quality.Clarity},
		{name: "quality.specificity", rating: request.Quality.Specificity},
		{name: "quality.fairness", rating: request.Quality.Fairness},
		{name: "quality.consistency", rating: request.Quality.Consistency},
		{name: "quality.accountability", rating: request.Quality.Accountability},
		{name: "quality.low_loophole_risk", rating: request.Quality.LowLoopholeRisk},
		{name: "gray_zone.vague_language_risk", rating: request.GrayZone.VagueLanguageRisk},
		{name: "gray_zone.definition_risk", rating: request.GrayZone.DefinitionRisk},
		{name: "gray_zone.contradiction_risk", rating: request.GrayZone.ContradictionRisk},
		{name: "gray_zone.enforcement_discretion_risk", rating: request.GrayZone.EnforcementDiscretion},
		{name: "gray_zone.exception_boundary_risk", rating: request.GrayZone.ExceptionBoundaryRisk},
		{name: "power.power_concentration_risk", rating: request.Power.PowerConcentrationRisk},
		{name: "power.oversight_strength", rating: request.Power.OversightStrength},
	}
	for _, item := range ratings {
		if item.rating.Score < 0 || item.rating.Score > 100 {
			return fmt.Errorf("%s score must be between 0 and 100", item.name)
		}
		rationale := strings.TrimSpace(item.rating.Rationale)
		if rationale == "" {
			return fmt.Errorf("%s rationale is required", item.name)
		}
		if len([]byte(rationale)) > 2048 {
			return fmt.Errorf("%s rationale exceeds 2048 bytes", item.name)
		}
	}
	return nil
}

func scoreCriteria(inputs []criterionInput) ([]ScoredCriterion, float64) {
	results := make([]ScoredCriterion, 0, len(inputs))
	total := 0.0
	for _, input := range inputs {
		contribution := float64(input.rating.Score) * input.weight
		total += contribution
		status := "PASS"
		if input.direction == "higher_is_more_risk" {
			if input.rating.Score >= ScorePassThreshold {
				status = "REVIEW"
			}
		} else if input.rating.Score < ScorePassThreshold {
			status = "FAIL"
		}
		results = append(results, ScoredCriterion{
			ID:                   input.id,
			Label:                input.label,
			Score:                input.rating.Score,
			Weight:               input.weight,
			WeightedContribution: round2(contribution),
			Direction:            input.direction,
			Rationale:            strings.TrimSpace(input.rating.Rationale),
			ThresholdStatus:      status,
		})
	}
	return results, round2(total)
}

func strictGoodLawGate(inputs []criterionInput) GateResult {
	var failed []string
	for _, input := range inputs {
		if input.rating.Score < ScorePassThreshold {
			failed = append(failed, input.id)
		}
	}
	if len(failed) == 0 {
		return GateResult{
			ID:        "strict_good_law_gate",
			Status:    "PASS",
			Threshold: "every C,S,F,I,A,L score >= 60",
			Basis:     "Every quality dimension met the research screening threshold; qualified legal review is still required.",
		}
	}
	return GateResult{
		ID:        "strict_good_law_gate",
		Status:    "FAIL",
		Threshold: "every C,S,F,I,A,L score >= 60",
		Basis:     "Non-compensable dimensions below threshold: " + strings.Join(failed, ", ") + ".",
	}
}

func rightsGate(fairness Rating) GateResult {
	if fairness.Score >= ScorePassThreshold {
		return GateResult{
			ID:        "human_rights_fairness_gate",
			Status:    "PASS",
			Threshold: "F >= 60",
			Basis:     "The reviewer-supplied fairness score met the research threshold; this is not a human-rights or legal-validity finding.",
		}
	}
	return GateResult{
		ID:        "human_rights_fairness_gate",
		Status:    "FAIL",
		Threshold: "F >= 60",
		Basis:     fmt.Sprintf("Fairness score %d is below 60; averaging cannot override this gate.", fairness.Score),
	}
}

func scanPhrases(clause string) []PhraseHit {
	lower := strings.ToLower(clause)
	hits := make([]PhraseHit, 0)
	for _, phrase := range reviewPhrases {
		if count := strings.Count(lower, phrase); count > 0 {
			hits = append(hits, PhraseHit{Phrase: phrase, Count: count})
		}
	}
	return hits
}

func qualityBand(score float64) string {
	switch {
	case score >= 85:
		return "85-100: strong and dependable candidate"
	case score >= 70:
		return "70 to <85: acceptable candidate; minor revision indicated"
	case score >= 50:
		return "50 to <70: significant gray-zone risk"
	case score >= 30:
		return "30 to <50: high manipulation risk"
	default:
		return "0 to <30: unfit for reliable enforcement"
	}
}

func findings(request Request, quality, gray float64, strict, rights GateResult, manipulation bool, hits []PhraseHit) []string {
	result := []string{
		fmt.Sprintf("Reviewer-supplied weighted law-quality score: %.2f/100.", quality),
		fmt.Sprintf("Reviewer-supplied weighted gray-zone risk score: %.2f/100.", gray),
	}
	if len(hits) > 0 {
		result = append(result, fmt.Sprintf("Literal scan found %d configured ambiguity phrase type(s); context must be reviewed by a human.", len(hits)))
	} else {
		result = append(result, "Literal scan found no configured ambiguity phrases; this does not establish clarity.")
	}
	if strict.Status == "FAIL" {
		result = append(result, strict.Basis)
	}
	if rights.Status == "FAIL" {
		result = append(result, rights.Basis)
	}
	if manipulation {
		result = append(result, "The research manipulation trigger fired: gray-zone risk and concentrated power are high while accountability is below threshold.")
	}
	if request.Power.OversightStrength.Score < ScorePassThreshold {
		result = append(result, "Reviewer-supplied oversight strength is below the research threshold.")
	}
	return result
}

func rewriteRequirements(request Request, hits []PhraseHit) []string {
	var requirements []string
	add := func(value string) {
		requirements = append(requirements, value)
	}
	if request.Quality.Clarity.Score < ScorePassThreshold || len(hits) > 0 {
		add("Define or replace material vague terms and state obligations in direct, ordinary language.")
	}
	if request.Quality.Specificity.Score < ScorePassThreshold || request.GrayZone.DefinitionRisk.Score >= ScorePassThreshold {
		add("State who is covered, what conduct triggers the rule, where and when it applies, and the exact definitions used.")
	}
	if request.Quality.Fairness.Score < ScorePassThreshold {
		add("Add due process, notice, an opportunity to respond, proportionality, non-discrimination safeguards, and effective remedy.")
	}
	if request.Quality.Consistency.Score < ScorePassThreshold || request.GrayZone.ContradictionRisk.Score >= ScorePassThreshold {
		add("Resolve conflicting clauses and state how comparable cases must be treated and reviewed for selective enforcement.")
	}
	if request.Quality.Accountability.Score < ScorePassThreshold || request.Power.OversightStrength.Score < ScorePassThreshold {
		add("Require written reasons, an auditable record, conflict disclosure, independent review, an appeal deadline, and consequences for abuse.")
	}
	if request.Quality.LowLoopholeRisk.Score < ScorePassThreshold || request.GrayZone.ExceptionBoundaryRisk.Score >= ScorePassThreshold {
		add("Bound every exception by actor, evidence, purpose, scope, duration, review, expiry, and remedy.")
	}
	if request.GrayZone.EnforcementDiscretion.Score >= ScorePassThreshold {
		add("Replace open-ended enforcement discretion with defined evidence thresholds, proportional actions, and escalation limits.")
	}
	if request.Power.PowerConcentrationRisk.Score >= ScorePassThreshold {
		add("Separate decision, review, and remedy authority so one actor cannot define, enforce, and finally judge the same rule.")
	}
	add("Have qualified counsel and affected communities review the complete instrument, related definitions, procedures, and jurisdiction-specific law.")
	sort.Strings(requirements)
	return requirements
}

func standardReviewWorkflow() []ReviewSection {
	return []ReviewSection{
		{Stage: "Language test", Questions: []string{
			"Can an ordinary affected person understand the rule?",
			"Are technical and material discretionary terms defined?",
			"Are obligations and prohibitions expressed directly?",
		}},
		{Stage: "Boundary test", Questions: []string{
			"Who and what conduct are covered, where, and for what period?",
			"What exceptions exist and who may invoke them?",
		}},
		{Stage: "Evidence test", Questions: []string{
			"What evidence and standard of proof are required?",
			"Who bears the burden, and may the affected person challenge the evidence?",
		}},
		{Stage: "Enforcement test", Questions: []string{
			"Could comparable decision-makers reach materially different results?",
			"Are sanctions proportionate and is selective enforcement detectable?",
		}},
		{Stage: "Rights test", Questions: []string{
			"Are notice, response, due process, non-discrimination, appeal, and remedy protected?",
			"Does the rule impose punishment or deprivation before an accountable finding?",
		}},
		{Stage: "Accountability test", Questions: []string{
			"Must the authority document evidence, legal basis, reasons, scope, and duration?",
			"Is review independent, timely, accessible, and capable of correcting abuse?",
		}},
	}
}

func disposition(quality, gray float64, strictPass, rightsPass, manipulation bool) string {
	switch {
	case !rightsPass:
		return "FUNDAMENTAL_REVISION_REQUIRED"
	case manipulation || quality < 50 || gray >= 70:
		return "MAJOR_REVISION_REQUIRED"
	case !strictPass || quality < 85 || gray >= 30:
		return "REVISION_REQUIRED"
	default:
		return "QUALIFIED_REVIEW_REQUIRED"
	}
}

func lawClarityAssurance() []AssuranceCheck {
	return []AssuranceCheck{
		{Dimension: "Truth", Status: "CONDITIONAL", Basis: "Reviewer ratings and source identity are declared, not independently authenticated."},
		{Dimension: "Reasoning", Status: "PASS", Basis: "Weights, thresholds, gates, and formula contributions are visible and non-compensable failures remain visible."},
		{Dimension: "Learning", Status: "CONDITIONAL", Basis: "The report supplies review questions but stores no outcomes or corrections."},
		{Dimension: "Communication", Status: "PASS", Basis: "Quality, ambiguity, rights gate, manipulation trigger, limitations, and disposition are separated."},
		{Dimension: "Adaptability", Status: "CONDITIONAL", Basis: "The English-language rubric is not validated across jurisdictions, languages, disabilities, or legal traditions."},
		{Dimension: "Ethics", Status: "PASS", Basis: "Fairness is a mandatory gate and cannot be rescued by a high average score."},
		{Dimension: "Safety", Status: "PASS", Basis: "The feature cannot issue legal verdicts, enforce rules, punish people, or execute external actions."},
		{Dimension: "Humility", Status: "PASS", Basis: "The report identifies itself as screening and requires qualified legal and affected-community review."},
		{Dimension: "Common good", Status: "CONDITIONAL", Basis: "The rubric foregrounds rights and accountability, but social benefit has not been empirically validated."},
	}
}

func round2(value float64) float64 {
	return math.Round(value*100) / 100
}
