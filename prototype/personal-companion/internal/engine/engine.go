package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/jonipwi/bootx/prototype/personal-companion/internal/model"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/policy"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/warning"
)

type Engine struct {
	policy *policy.Engine
	now    func() time.Time
}

func New() (*Engine, error) {
	p, err := policy.Load()
	if err != nil {
		return nil, err
	}
	return &Engine{policy: p, now: time.Now}, nil
}

func (e *Engine) Process(request model.Request) (model.Packet, error) {
	if err := e.validate(request); err != nil {
		return model.Packet{}, err
	}

	base := e.basePacket(request)
	if request.DataClass == model.DataProhibited {
		return e.blockedPacket(base, "The selected input is classified as prohibited and was not analyzed."), nil
	}

	analysis := e.policy.Analyze(request)
	if analysis.Prohibited {
		return e.blockedPacket(base, "The requested purpose matches a prohibited capability in the deterministic policy."), nil
	}

	appendAnalysisFindings(&base, request, analysis)
	if request.Warning != nil {
		return e.processWarning(base, request, analysis)
	}
	return e.processGeneral(base, request, analysis), nil
}

func (e *Engine) validate(request model.Request) error {
	if request.CapabilityID != model.CapabilityID {
		return fmt.Errorf("capability_id must be %q", model.CapabilityID)
	}
	if strings.TrimSpace(request.RequestID) == "" || strings.TrimSpace(request.UserID) == "" {
		return fmt.Errorf("request_id and user_id are required")
	}
	if strings.TrimSpace(request.Goal) == "" || strings.TrimSpace(request.Question) == "" {
		return fmt.Errorf("goal and question are required")
	}
	if request.CreatedAt.IsZero() {
		return fmt.Errorf("created_at is required")
	}
	if strings.TrimSpace(request.ContentSource.Type) == "" {
		return fmt.Errorf("content_source.type is required")
	}
	if err := validateContentIntegrity(request); err != nil {
		return err
	}
	if len([]byte(request.SelectedContent)) > e.policy.Rules.MaxSelectedContentBytes {
		return fmt.Errorf("selected_content exceeds %d bytes", e.policy.Rules.MaxSelectedContentBytes)
	}
	if !validDataClass(request.DataClass) {
		return fmt.Errorf("unsupported data_class %q", request.DataClass)
	}
	if !validDomain(request.DeclaredDomain) {
		return fmt.Errorf("unsupported declared_domain %q", request.DeclaredDomain)
	}
	if request.DataClass == model.DataSensitive && !request.Synthetic {
		return fmt.Errorf("real sensitive input is not authorized before security review; use only an explicitly synthetic test case")
	}
	if request.Warning != nil && !request.Synthetic {
		return fmt.Errorf("warning assessment is synthetic-only in this unvalidated MVP")
	}
	if request.MemoryPermission != "none" && request.MemoryPermission != "session" {
		return fmt.Errorf("memory_permission must be none or session in MVP version 1")
	}
	if request.RemotePermission != "deny" {
		return fmt.Errorf("remote_permission must be deny: this build has no remote-processing capability")
	}
	if !oneOf(request.OutputPreference, "concise", "standard", "detailed", "checklist", "comparison", "warning_card") {
		return fmt.Errorf("unsupported output_preference %q", request.OutputPreference)
	}
	return nil
}

func (e *Engine) basePacket(request model.Request) model.Packet {
	return model.Packet{
		RequestID:      request.RequestID,
		CapabilityID:   model.CapabilityID,
		GeneratedAt:    e.now().UTC(),
		RuntimeNotice:  "Deterministic DEV-1 prototype with contained read-only local-document ingestion; no AI model, network lookup, external action, or safety certification.",
		GoalUnderstood: strings.TrimSpace(request.Goal),
		Observations:   []model.Finding{},
		Assumptions: []string{
			"The goal, domain, user identifier, and priorities were declared locally; this prototype has no user-authentication system.",
		},
		Unknowns: []string{},
		Options:  []model.Option{},
		BlockedActions: []string{
			"open_active_link",
			"send_message",
			"place_call",
			"submit_credentials",
			"transfer_money",
			"change_account",
			"control_device_or_robot",
			"broadcast_family_or_public_alert",
		},
		Limitations: []string{
			"No external source, identity, account, location, forecast, or event was independently verified by this offline build.",
			"A matching content hash proves which bytes were processed, not who authored them or whether their claims are true.",
			"Recommendations are advisory; Joni remains the decision-maker.",
		},
		EvidenceReceipt: evidenceReceipt(request.ContentSource, len([]byte(request.SelectedContent))),
		DataReceipt: model.DataReceipt{
			MemoryUsed:       false,
			RemoteProcessing: false,
			Synthetic:        request.Synthetic,
			RawRetention:     "process_memory_only; discarded when the process exits",
			LocationUse:      "none unless manually declared in a warning request; never persisted",
		},
	}
}

func (e *Engine) blockedPacket(packet model.Packet, reason string) model.Packet {
	packet.DecisionClass = model.D5Prohibited
	packet.ResponseMode = model.ModeBlock
	packet.Recommendation = model.Recommendation{Status: "blocked", Basis: reason}
	packet.NextSafeStep = "Restate a lawful, non-coercive, protective goal without prohibited data or external action."
	packet.Limitations = append(packet.Limitations, reason)
	packet.Assurance = assuranceChecks(true, false)
	return packet
}

func (e *Engine) processGeneral(packet model.Packet, request model.Request, analysis policy.Analysis) model.Packet {
	class, mode := classifyGeneral(request, analysis)
	packet.DecisionClass = class
	packet.ResponseMode = mode
	packet.Unknowns = append(packet.Unknowns, "The offline MVP has no retrieved evidence beyond the deliberately selected input.")
	isLocalDocument := request.ContentSource.Type == "local_workspace_file" && request.ContentSource.IntegrityVerified
	localSummary := localDocumentSummary{}
	if isLocalDocument {
		localSummary = analyzeLocalDocument(request.SelectedContent)
		appendLocalDocumentFindings(&packet, localSummary)
	}

	switch {
	case analysis.HasCredential || analysis.HasPayment || (analysis.HasURL && analysis.HasUrgency):
		packet.Options = suspiciousMessageOptions()
		packet.Recommendation = model.Recommendation{
			Status:   "advisory",
			OptionID: "verify-independent",
			Basis:    "Independent verification avoids acting through an unverified channel while credential, payment, link, or urgency indicators remain unresolved.",
		}
		packet.NextSafeStep = "Do not use the selected link or provide secrets; obtain the organization's official channel independently."
		packet.ResponseMode = model.ModeVerify
		packet.DecisionClass = model.D2Consequential
	case class == model.D3Qualified:
		packet.Options = qualifiedDomainOptions(request.DeclaredDomain)
		packet.Recommendation = model.Recommendation{
			Status:   "advisory",
			OptionID: "prepare-qualified-review",
			Basis:    "The domain requires qualified judgment that this prototype does not possess.",
		}
		packet.NextSafeStep = "Write down the facts, questions, time constraints, and records needed for an appropriate qualified professional."
		packet.ResponseMode = model.ModeAbstain
	case class == model.D4Emergency:
		packet.Options = emergencyOptions()
		packet.Recommendation = model.Recommendation{
			Status: "limited",
			Basis:  "The input may concern immediate safety, but no event or official instruction was verified.",
		}
		packet.NextSafeStep = "If danger is immediate, move away from it and use an independently known local emergency channel; otherwise verify the responsible authority."
		packet.ResponseMode = model.ModeUrgentGuidance
	case isLocalDocument:
		packet.Options = localDocumentOptions(localSummary)
		packet.Recommendation = model.Recommendation{
			Status:   "advisory",
			OptionID: "review-evidence-gap",
			Basis:    "A contained deterministic scan can identify structural review candidates, but a human must interpret context and verify sources.",
		}
		if localSummary.FirstGap != "" {
			packet.NextSafeStep = "Review this candidate in context and verify its most material supporting source: " + localSummary.FirstGap
		} else {
			packet.NextSafeStep = "Choose one material claim and record the exact source, evidence maturity, uncertainty, and correction condition."
		}
	default:
		packet.Options = generalOptions(request.UserPriorities)
		packet.Recommendation = model.Recommendation{
			Status:   "advisory",
			OptionID: "gather-evidence",
			Basis:    "A reversible evidence-gathering step preserves agency and reduces avoidable error before commitment.",
		}
		packet.NextSafeStep = "Identify the one missing fact most likely to change the decision, then verify it from an appropriate source."
	}
	packet.Assurance = assuranceChecks(false, len(packet.Unknowns) > 0)
	return packet
}

func (e *Engine) processWarning(packet model.Packet, request model.Request, analysis policy.Analysis) (model.Packet, error) {
	result, err := warning.Evaluate(*request.Warning)
	if err != nil {
		return model.Packet{}, err
	}
	packet.Warning = result.Card
	packet.DecisionClass = result.DecisionClass
	packet.ResponseMode = result.ResponseMode
	packet.Unknowns = append(packet.Unknowns, result.Card.Unknowns...)
	packet.Limitations = append(packet.Limitations, result.Card.Limitations...)
	packet.NextSafeStep = result.Card.NextSafeStep
	packet.Options = warningOptions(result.Card.Level)
	packet.Recommendation = model.Recommendation{
		Status:   "advisory",
		OptionID: warningOptionID(result.Card.Level),
		Basis:    result.Card.Observed[0],
	}
	conditional := result.Card.Level == warning.WX || result.Card.Level == warning.W0 || len(result.Card.Unknowns) > 0
	packet.Assurance = assuranceChecks(false, conditional)
	if result.Card.Level == warning.WX {
		for i := range packet.Assurance {
			if packet.Assurance[i].Dimension == "Truth" || packet.Assurance[i].Dimension == "Safety" {
				packet.Assurance[i].Status = "CONDITIONAL"
				packet.Assurance[i].Basis = "Safe verification output only; warning reliability is degraded or conflicted."
			}
		}
	}
	_ = analysis
	return packet, nil
}

func appendAnalysisFindings(packet *model.Packet, request model.Request, analysis policy.Analysis) {
	if request.ContentSource.IntegrityVerified {
		packet.Observations = append(packet.Observations, model.Finding{
			Claim:  "Selected content bytes match the declared SHA-256 and byte length.",
			Source: fallbackReference(request.ContentSource.Reference, request.ContentSource.Type),
			Status: "integrity_verified_origin_not_authenticated",
		})
	} else if request.ContentSource.OriginVerified {
		packet.Observations = append(packet.Observations, model.Finding{
			Claim:  "The input declares the content origin verified, but this offline prototype did not authenticate that origin.",
			Source: request.ContentSource.Type,
			Status: "user_claim_not_independently_verified",
		})
	}
	if len(analysis.MatchedSignals) == 0 {
		packet.Observations = append(packet.Observations, model.Finding{
			Claim:  "No configured scam, credential, payment, urgency, or prompt-injection text indicator was detected.",
			Source: "deterministic_local_rules",
			Status: "observed",
		})
		return
	}
	for _, signal := range analysis.MatchedSignals {
		packet.Observations = append(packet.Observations, model.Finding{
			Claim:  "Selected input contains " + signal + ".",
			Source: request.ContentSource.Type,
			Status: "observed_indicator_not_verdict",
		})
	}
}

func validateContentIntegrity(request model.Request) error {
	source := request.ContentSource
	if !source.IntegrityVerified {
		if source.SHA256 != "" || source.ByteLength != 0 || source.Reference != "" || source.ModifiedAt != "" {
			return fmt.Errorf("content-source integrity metadata requires integrity_verified=true")
		}
		return nil
	}
	if strings.TrimSpace(source.Reference) == "" {
		return fmt.Errorf("content_source.reference is required for integrity-verified content")
	}
	digest, err := hex.DecodeString(source.SHA256)
	if err != nil || len(digest) != sha256.Size {
		return fmt.Errorf("content_source.sha256 must be a 64-character hexadecimal SHA-256 digest")
	}
	contentBytes := []byte(request.SelectedContent)
	if source.ByteLength != len(contentBytes) {
		return fmt.Errorf("content_source.byte_length does not match selected_content bytes")
	}
	actual := sha256.Sum256(contentBytes)
	if !strings.EqualFold(source.SHA256, hex.EncodeToString(actual[:])) {
		return fmt.Errorf("content_source.sha256 does not match selected_content bytes")
	}
	if source.ModifiedAt != "" {
		if _, err := time.Parse(time.RFC3339Nano, source.ModifiedAt); err != nil {
			return fmt.Errorf("content_source.modified_at must be RFC 3339: %w", err)
		}
	}
	return nil
}

func evidenceReceipt(source model.ContentSource, selectedBytes int) model.EvidenceReceipt {
	originStatus := "not_authenticated"
	if source.OriginVerified {
		originStatus = "user_claimed_verified_not_authenticated"
	}
	return model.EvidenceReceipt{
		SourceType:        source.Type,
		Reference:         source.Reference,
		IntegrityVerified: source.IntegrityVerified,
		SHA256:            strings.ToLower(source.SHA256),
		ByteLength:        selectedBytes,
		ModifiedAt:        source.ModifiedAt,
		OriginStatus:      originStatus,
	}
}

func fallbackReference(reference, sourceType string) string {
	if strings.TrimSpace(reference) != "" {
		return reference
	}
	return sourceType
}

func classifyGeneral(request model.Request, analysis policy.Analysis) (model.DecisionClass, model.ResponseMode) {
	if request.DeclaredDomain == model.DomainHealth || request.DeclaredDomain == model.DomainLegal {
		return model.D3Qualified, model.ModeAbstain
	}
	if request.DeclaredDomain == model.DomainEmergency {
		return model.D4Emergency, model.ModeUrgentGuidance
	}
	if request.DeclaredDomain == model.DomainFinance || request.DeclaredDomain == model.DomainRelationship || request.DeclaredDomain == model.DomainFaith || request.DataClass == model.DataSensitive {
		return model.D2Consequential, model.ModeCompare
	}
	if analysis.HasCredential || analysis.HasPayment || (analysis.HasURL && analysis.HasUrgency) {
		return model.D2Consequential, model.ModeVerify
	}
	if request.DeclaredDomain == model.DomainStudy {
		return model.D0Informational, model.ModeInform
	}
	return model.D1Reversible, model.ModeCompare
}

func generalOptions(priorities []string) []model.Option {
	priorityNote := "declared priorities"
	if len(priorities) == 0 {
		priorityNote = "a short list of user-selected priorities"
	}
	return []model.Option{
		{OptionID: "defer", Summary: "Defer the decision until a defined review time.", Benefits: []string{"avoids premature commitment"}, Risks: []string{"delay may have an opportunity cost"}, Reversibility: "easy", ExternalEffect: "none"},
		{OptionID: "gather-evidence", Summary: "Verify the most decision-relevant missing fact.", Benefits: []string{"reduces avoidable uncertainty"}, Risks: []string{"takes time and source judgment"}, Reversibility: "easy", ExternalEffect: "none"},
		{OptionID: "reversible-step", Summary: "Take the smallest safe reversible step.", Benefits: []string{"creates learning while limiting downside"}, Risks: []string{"may not resolve the full decision"}, Reversibility: "easy", ExternalEffect: "none"},
		{OptionID: "compare-priorities", Summary: "Compare permissible options against " + priorityNote + ".", Benefits: []string{"keeps the user's values visible"}, Risks: []string{"values do not replace missing facts"}, Reversibility: "easy", ExternalEffect: "none"},
	}
}

func suspiciousMessageOptions() []model.Option {
	return []model.Option{
		{OptionID: "do-not-use-channel", Summary: "Do not open the selected link or reply through the unverified channel.", Benefits: []string{"avoids immediate credential or payment exposure"}, Risks: []string{"a legitimate request may be delayed"}, Reversibility: "easy", ExternalEffect: "none"},
		{OptionID: "verify-independent", Summary: "Use an independently obtained official application, website, number, or in-person channel.", Benefits: []string{"separates verification from the suspicious message"}, Risks: []string{"requires careful source selection"}, Reversibility: "easy", ExternalEffect: "none"},
		{OptionID: "trusted-review", Summary: "Ask a preselected trusted person to review the indicators without sharing secrets.", Benefits: []string{"adds a second perspective"}, Risks: []string{"may expose unnecessary personal context"}, Reversibility: "easy", ExternalEffect: "draft_only"},
	}
}

func qualifiedDomainOptions(domain model.Domain) []model.Option {
	return []model.Option{
		{OptionID: "prepare-qualified-review", Summary: "Prepare facts and questions for an appropriate qualified " + string(domain) + " professional.", Benefits: []string{"supports informed professional review"}, Risks: []string{"availability and cost may vary"}, Reversibility: "easy", ExternalEffect: "none"},
		{OptionID: "official-education", Summary: "Read current general information from a responsible official source.", Benefits: []string{"improves vocabulary and question quality"}, Risks: []string{"general information may not fit the case"}, Reversibility: "easy", ExternalEffect: "none"},
		{OptionID: "safe-hold", Summary: "Avoid irreversible action until qualified review unless delay itself creates immediate danger.", Benefits: []string{"limits unsupported commitment"}, Risks: []string{"delay can matter in urgent cases"}, Reversibility: "easy", ExternalEffect: "none"},
	}
}

func emergencyOptions() []model.Option {
	return []model.Option{
		{OptionID: "move-from-danger", Summary: "Move away from direct immediate danger when it is safe to do so.", Benefits: []string{"reduces exposure"}, Risks: []string{"movement may be unsafe in some hazards"}, Reversibility: "limited", ExternalEffect: "user_controlled"},
		{OptionID: "official-contact", Summary: "Use an independently known responsible emergency or local authority channel.", Benefits: []string{"connects to accountable responders"}, Risks: []string{"service may be unavailable"}, Reversibility: "easy", ExternalEffect: "user_controlled"},
		{OptionID: "verify-if-safe", Summary: "If there is no immediate danger, verify event, area, time, and instruction before disruptive action.", Benefits: []string{"reduces false escalation"}, Risks: []string{"must not delay urgent safety action"}, Reversibility: "easy", ExternalEffect: "none"},
	}
}

func warningOptions(level string) []model.Option {
	base := []model.Option{
		{OptionID: "official-instruction", Summary: "In this exercise, review the instruction attributed to the authority; real use requires independent authentication outside BootX.", Benefits: []string{"preserves the source wording for review"}, Risks: []string{"issuer, area, and freshness are not authenticated by this prototype"}, Reversibility: "varies", ExternalEffect: "user_controlled"},
		{OptionID: "verify-source", Summary: "Verify issuer, status, area, time, and update/cancellation through an independently obtained official channel.", Benefits: []string{"reduces spoofing and stale-alert risk"}, Risks: []string{"verification takes time"}, Reversibility: "easy", ExternalEffect: "none"},
	}
	switch level {
	case warning.W4:
		return append([]model.Option{{OptionID: "act-now", Summary: "Take immediate personal safety action without waiting for AI certainty.", Benefits: []string{"limits the cost of dangerous delay"}, Risks: []string{"the safest movement depends on the hazard"}, Reversibility: "limited", ExternalEffect: "user_controlled"}}, base...)
	case warning.W3:
		return append([]model.Option{{OptionID: "protect-now", Summary: "Activate the personal protection plan within the available lead time.", Benefits: []string{"reduces exposure before impact"}, Risks: []string{"must remain consistent with official instruction"}, Reversibility: "limited", ExternalEffect: "user_controlled"}}, base...)
	case warning.W2:
		return append([]model.Option{{OptionID: "prepare-now", Summary: "Take low-burden reversible preparedness steps.", Benefits: []string{"improves readiness with limited downside"}, Risks: []string{"may consume time or supplies if event changes"}, Reversibility: "easy", ExternalEffect: "user_controlled"}}, base...)
	case warning.W1, warning.W0:
		return append([]model.Option{{OptionID: "monitor", Summary: "Monitor a named official source at a defined review time.", Benefits: []string{"maintains awareness without false escalation"}, Risks: []string{"updates can be missed"}, Reversibility: "easy", ExternalEffect: "none"}}, base...)
	default:
		return append([]model.Option{{OptionID: "verify-before-action", Summary: "Verify source integrity and personal relevance before consequential action.", Benefits: []string{"contains conflicted or unverified evidence"}, Risks: []string{"reversible precautions may still be needed"}, Reversibility: "easy", ExternalEffect: "none"}}, base...)
	}
}

func warningOptionID(level string) string {
	switch level {
	case warning.W4:
		return "act-now"
	case warning.W3:
		return "protect-now"
	case warning.W2:
		return "prepare-now"
	case warning.W1, warning.W0:
		return "monitor"
	default:
		return "verify-before-action"
	}
}

func assuranceChecks(blocked, conditional bool) []model.AssuranceCheck {
	status := "PASS"
	truthBasis := "Claims are limited to typed user input and deterministic observations."
	if conditional {
		status = "CONDITIONAL"
		truthBasis = "Output is usable only with displayed unknowns and independent verification."
	}
	checks := []model.AssuranceCheck{
		{Dimension: "Truth", Status: status, Basis: truthBasis},
		{Dimension: "Reasoning", Status: status, Basis: "Decision class, alternatives, reversibility, and missing evidence are explicit."},
		{Dimension: "Learning", Status: "CONDITIONAL", Basis: "Corrections can be reviewed in-session; persistent learning is intentionally absent."},
		{Dimension: "Communication", Status: "PASS", Basis: "The packet separates observations, unknowns, options, recommendation status, and limits."},
		{Dimension: "Adaptability", Status: "CONDITIONAL", Basis: "Text-only local operation is supported; multilingual and accessibility validation remain incomplete."},
		{Dimension: "Ethics", Status: "PASS", Basis: "No hidden memory, remote processing, coercive action, or external execution exists."},
		{Dimension: "Safety", Status: "PASS", Basis: "Consequential tools are technically absent and untrusted content cannot grant authority."},
		{Dimension: "Humility", Status: "PASS", Basis: "The prototype discloses limits and preserves Joni's decision authority."},
		{Dimension: "Common good", Status: "CONDITIONAL", Basis: "Personal protection is in scope; family and community dissemination remain blocked."},
	}
	if blocked {
		checks[5].Basis = "The deterministic policy blocked the prohibited request."
		checks[6].Basis = "The deterministic policy blocked the prohibited request before content analysis or external effect."
	}
	return checks
}

func validDataClass(value model.DataClass) bool {
	return value == model.DataPublic || value == model.DataPersonal || value == model.DataSensitive || value == model.DataProhibited
}

func validDomain(value model.Domain) bool {
	switch value {
	case model.DomainGeneral, model.DomainStudy, model.DomainDigitalSafety, model.DomainHousehold, model.DomainHealth, model.DomainFinance, model.DomainLegal, model.DomainRelationship, model.DomainFaith, model.DomainEmergency, model.DomainUnknown:
		return true
	default:
		return false
	}
}

func oneOf(value string, allowed ...string) bool {
	for _, candidate := range allowed {
		if value == candidate {
			return true
		}
	}
	return false
}
