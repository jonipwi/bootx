package warning

import (
	"fmt"
	"strings"

	"github.com/jonipwi/bootx/prototype/personal-companion/internal/model"
)

const (
	W0 = "W0"
	W1 = "W1"
	W2 = "W2"
	W3 = "W3"
	W4 = "W4"
	WX = "WX"
)

type Result struct {
	Card          *model.WarningCard
	DecisionClass model.DecisionClass
	ResponseMode  model.ResponseMode
}

func Evaluate(in model.WarningInput) (Result, error) {
	if strings.TrimSpace(in.EventID) == "" || strings.TrimSpace(in.HazardType) == "" {
		return Result{}, fmt.Errorf("warning event_id and hazard_type are required")
	}
	if !oneOf(in.OfficialStatus, "active", "none_found", "unavailable") {
		return Result{}, fmt.Errorf("official_status must be active, none_found, or unavailable")
	}
	if !oneOf(in.AreaMatch, "inside", "near", "outside", "unknown") {
		return Result{}, fmt.Errorf("area_match must be inside, near, outside, or unknown")
	}
	if !oneOf(in.IntegrityStatus, "pass", "fail", "unknown") {
		return Result{}, fmt.Errorf("integrity_status must be pass, fail, or unknown")
	}
	if !oneOf(in.MessageStatus, "Actual", "Exercise", "System", "Test", "Draft") {
		return Result{}, fmt.Errorf("message_status must be a supported CAP status")
	}
	if !oneOf(in.MessageType, "Alert", "Update", "Cancel", "Ack", "Error") {
		return Result{}, fmt.Errorf("message_type must be a supported CAP message type")
	}
	if !oneOf(in.Urgency, "Immediate", "Expected", "Future", "Past", "Unknown") {
		return Result{}, fmt.Errorf("urgency must be Immediate, Expected, Future, Past, or Unknown")
	}
	if !oneOf(in.Severity, "Extreme", "Severe", "Moderate", "Minor", "Unknown") {
		return Result{}, fmt.Errorf("severity must be Extreme, Severe, Moderate, Minor, or Unknown")
	}
	if !oneOf(in.Certainty, "Observed", "Likely", "Possible", "Unlikely", "Unknown") {
		return Result{}, fmt.Errorf("certainty must be Observed, Likely, Possible, Unlikely, or Unknown")
	}
	if !validTier(in.EvidenceTier) {
		return Result{}, fmt.Errorf("evidence_tier must be V0, V1, V2, V3, or V4")
	}

	level, reason := assignLevel(in)
	label, mode, posture, class, step, review := describe(level)
	card := &model.WarningCard{
		EventID:             in.EventID,
		HazardType:          in.HazardType,
		OfficialStatus:      in.OfficialStatus,
		Issuer:              fallback(in.Issuer, "not supplied"),
		OfficialURL:         in.OfficialURL,
		IssuedAt:            in.IssuedAt,
		UpdatedAt:           in.UpdatedAt,
		ExpiresAt:           in.ExpiresAt,
		AffectedArea:        in.AffectedArea,
		AreaMatch:           in.AreaMatch,
		Urgency:             fallback(in.Urgency, "Unknown"),
		Severity:            fallback(in.Severity, "Unknown"),
		Certainty:           fallback(in.Certainty, "Unknown"),
		Level:               level,
		LevelLabel:          label,
		EvidenceTier:        strings.ToUpper(in.EvidenceTier),
		DecisionPosture:     posture,
		OfficialInstruction: in.Instruction,
		Observed:            []string{reason},
		NextSafeStep:        step,
		NextReview:          review,
		Limitations: []string{
			"BootX is personal decision support, not an alerting or emergency authority.",
			"The level is deterministic guidance and does not replace the issuer's original category or instruction.",
		},
	}

	if in.OfficialStatus == "none_found" {
		card.Unknowns = append(card.Unknowns, "No official alert found is not proof that no hazard exists.")
	}
	if !in.AuthorityAuthenticated && in.OfficialStatus == "active" {
		card.Unknowns = append(card.Unknowns, "The claimed issuing authority was not authenticated.")
	}
	if in.AreaMatch == "unknown" {
		card.Unknowns = append(card.Unknowns, "Personal location relevance is unknown.")
	}
	if in.SourceConflict {
		card.Unknowns = append(card.Unknowns, "Material source conflict requires independent verification.")
	}
	if in.Stale {
		card.Unknowns = append(card.Unknowns, "The supplied alert or forecast may be stale.")
	}
	if in.DirectDanger {
		card.Limitations = append(card.Limitations, "Immediate guidance is based on direct/user-reported danger, not an independently verified official alert.")
	}
	if in.OfficialStatus == "active" && in.AuthorityAuthenticated {
		card.Observed = append(card.Observed, "An active alert was supplied with an authenticated issuing authority.")
	}
	if in.Certainty != "" && !strings.EqualFold(in.Certainty, "Observed") {
		card.Forecast = append(card.Forecast, "The source characterizes event certainty as "+in.Certainty+".")
	}
	return Result{Card: card, DecisionClass: class, ResponseMode: mode}, nil
}

func assignLevel(in model.WarningInput) (string, string) {
	if in.DirectDanger {
		return W4, "Direct or user-reported immediate danger was declared; delay may create serious harm."
	}
	if in.AreaMatch == "outside" {
		return W0, "The supplied event area does not match the configured personal area."
	}
	status := strings.ToLower(in.MessageStatus)
	msgType := strings.ToLower(in.MessageType)
	if status == "test" || status == "exercise" || status == "draft" {
		return W0, "The supplied message is marked as " + status + ", not an actual active warning."
	}
	if msgType == "cancel" || strings.EqualFold(in.Urgency, "Past") {
		return W0, "The supplied message is canceled or no longer requires responsive action."
	}
	if in.IntegrityStatus == "fail" || in.SourceConflict || in.Stale {
		return WX, "Source integrity, conflict, or freshness prevents a reliable personal warning level."
	}
	if in.OfficialStatus == "unavailable" || in.IntegrityStatus == "unknown" {
		return WX, "Official status or source integrity is unavailable or unknown."
	}
	if in.OfficialStatus == "active" {
		if !in.AuthorityAuthenticated || in.AreaMatch == "unknown" {
			return WX, "The claimed active alert cannot be safely personalized until authority and area relevance are verified."
		}
		instruction := strings.ToLower(in.Instruction)
		if strings.EqualFold(in.Urgency, "Immediate") || containsAny(instruction, []string{"evacuate now", "take shelter now", "move to higher ground", "leave immediately"}) {
			return W4, "An authenticated relevant alert requires immediate responsive action."
		}
		if strings.EqualFold(in.Severity, "Extreme") || strings.EqualFold(in.Severity, "Severe") || strings.EqualFold(in.Urgency, "Expected") {
			return W3, "An authenticated relevant alert supports protective action within the available lead time."
		}
		return W2, "An authenticated relevant alert supports low-burden preparation while conditions are monitored."
	}

	switch strings.ToUpper(in.EvidenceTier) {
	case "V0":
		return WX, "The signal has unknown or unauthenticated provenance."
	case "V1":
		return W1, "One attributable source supports monitoring and independent verification."
	case "V2", "V3", "V4":
		if in.AreaMatch == "inside" || in.AreaMatch == "near" {
			return W2, "Corroborated evidence and possible personal relevance support reversible preparation."
		}
		return W1, "Corroborated evidence supports monitoring, but personal area relevance is limited."
	default:
		return WX, "Evidence state is not established."
	}
}

func describe(level string) (label string, mode model.ResponseMode, posture string, class model.DecisionClass, step, review string) {
	switch level {
	case W0:
		return "NO ACTIVE VERIFIED SIGNAL", model.ModeMonitor, "MONITOR", model.D1Reversible,
			"Continue ordinary preparedness and use the named official source for material decisions.", "Review when new authoritative information is issued."
	case W1:
		return "MONITOR", model.ModeMonitor, "MONITOR", model.D1Reversible,
			"Monitor the identified official source and verify the household communication plan.", "Set a user-chosen review time appropriate to the forecast lead time."
	case W2:
		return "PREPARE", model.ModePrepare, "PREPARE NOW", model.D2Consequential,
			"Take low-burden reversible preparation steps and preserve access to official updates.", "Review at the next official bulletin or earlier if conditions change."
	case W3:
		return "PROTECT", model.ModeProtect, "PROTECT NOW", model.D4Emergency,
			"Follow the authenticated issuing authority's instruction and activate the personal safety plan.", "Continuously monitor official updates until the alert is canceled or expires."
	case W4:
		return "URGENT ACTION", model.ModeUrgentGuidance, "ACT NOW", model.D4Emergency,
			"Move away from immediate danger and follow official or on-scene responder instructions now.", "Do not delay immediate safety action for another AI update."
	default:
		return "VERIFY / CONFLICT / DEGRADED", model.ModeVerify, "VERIFY BEFORE CONSEQUENT ACTION", model.D2Consequential,
			"Use an independently obtained official channel; take only justified reversible precautions while verifying.", "Review as soon as source integrity, freshness, and area relevance can be checked."
	}
}

func validTier(value string) bool {
	return oneOf(strings.ToUpper(value), "V0", "V1", "V2", "V3", "V4")
}

func oneOf(value string, allowed ...string) bool {
	for _, candidate := range allowed {
		if strings.EqualFold(value, candidate) {
			return true
		}
	}
	return false
}

func containsAny(text string, terms []string) bool {
	for _, term := range terms {
		if strings.Contains(text, term) {
			return true
		}
	}
	return false
}

func fallback(value, fallbackValue string) string {
	if strings.TrimSpace(value) == "" {
		return fallbackValue
	}
	return value
}
