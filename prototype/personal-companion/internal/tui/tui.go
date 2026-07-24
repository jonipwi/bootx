package tui

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/jonipwi/bootx/prototype/personal-companion/internal/engine"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/evidence"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/lawclarity"
	"github.com/jonipwi/bootx/prototype/personal-companion/internal/model"
)

type UI struct {
	reader *bufio.Reader
	out    io.Writer
	engine *engine.Engine
}

func Run(in io.Reader, out io.Writer, decisionEngine *engine.Engine) error {
	ui := &UI{reader: bufio.NewReader(in), out: out, engine: decisionEngine}
	return ui.run()
}

func (ui *UI) run() error {
	ui.clear()
	ui.header()
	for {
		fmt.Fprintln(ui.out, "\nMain menu")
		fmt.Fprintln(ui.out, "  1  Personal decision assistance")
		fmt.Fprintln(ui.out, "  2  Forecast/disaster warning assessment")
		fmt.Fprintln(ui.out, "  3  Read-only local workspace document")
		fmt.Fprintln(ui.out, "  4  Law Clarity Logic screening")
		fmt.Fprintln(ui.out, "  5  Safety boundaries")
		fmt.Fprintln(ui.out, "  q  Quit and discard session data")
		choice, err := ui.prompt("Choose", false)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		switch strings.ToLower(choice) {
		case "1":
			if err := ui.personalWorkflow(); err != nil {
				return err
			}
		case "2":
			if err := ui.warningWorkflow(); err != nil {
				return err
			}
		case "3":
			if err := ui.localDocumentWorkflow(); err != nil {
				return err
			}
		case "4":
			if err := ui.lawClarityWorkflow(); err != nil {
				return err
			}
		case "5":
			ui.boundaries()
		case "q", "quit", "exit":
			fmt.Fprintln(ui.out, "Session closed. No raw input was written by BootX Companion.")
			return nil
		default:
			fmt.Fprintln(ui.out, "Please choose 1, 2, 3, 4, 5, or q.")
		}
	}
}

func (ui *UI) lawClarityWorkflow() error {
	fmt.Fprintln(ui.out, "\n=== Law Clarity Logic screening ===")
	fmt.Fprintln(ui.out, "This feature calculates a transparent screening report from your ratings and rationales.")
	fmt.Fprintln(ui.out, "It does not decide legal validity, constitutionality, guilt, liability, or enforcement authority.")
	publicConfirmed, err := ui.confirm("Is the clause public, non-sensitive, and appropriate for educational screening")
	if err != nil {
		return err
	}
	if !publicConfirmed {
		fmt.Fprintln(ui.out, "Canceled before clause entry. Sensitive, privileged, sealed, or uncertain legal material is not authorized.")
		return nil
	}
	title, err := ui.prompt("Law, rule, policy, or procedure title", true)
	if err != nil {
		return err
	}
	jurisdiction, err := ui.prompt("Jurisdiction or organizational context", true)
	if err != nil {
		return err
	}
	instrumentType, err := ui.choose("Instrument type", []string{"law", "regulation", "company_policy", "court_procedure", "contract", "other"})
	if err != nil {
		return err
	}
	sourceReference, err := ui.prompt("Public source reference or citation", true)
	if err != nil {
		return err
	}
	purpose, err := ui.prompt("Screening purpose", true)
	if err != nil {
		return err
	}
	clause, err := ui.multiline("Public clause text (finish with a single period on its own line)")
	if err != nil {
		return err
	}

	fmt.Fprintln(ui.out, "\nQuality ratings: 0 means the condition fails; 100 means strong support.")
	clarity, err := ui.lawRating("Clarity (C)")
	if err != nil {
		return err
	}
	specificity, err := ui.lawRating("Specificity and boundaries (S)")
	if err != nil {
		return err
	}
	fairness, err := ui.lawRating("Fairness and rights protection (F)")
	if err != nil {
		return err
	}
	consistency, err := ui.lawRating("Consistent enforceability (I)")
	if err != nil {
		return err
	}
	accountability, err := ui.lawRating("Accountability and auditability (A)")
	if err != nil {
		return err
	}
	lowLoophole, err := ui.lawRating("Low loophole risk (L)")
	if err != nil {
		return err
	}

	fmt.Fprintln(ui.out, "\nGray-zone ratings: 0 means low observed risk; 100 means high observed risk.")
	vagueRisk, err := ui.lawRating("Vague-language risk (V)")
	if err != nil {
		return err
	}
	definitionRisk, err := ui.lawRating("Undefined/circular-definition risk (D)")
	if err != nil {
		return err
	}
	contradictionRisk, err := ui.lawRating("Contradictory-clause risk (X)")
	if err != nil {
		return err
	}
	discretionRisk, err := ui.lawRating("Enforcement-discretion risk (E)")
	if err != nil {
		return err
	}
	exceptionRisk, err := ui.lawRating("Unclear-exception-boundary risk (U)")
	if err != nil {
		return err
	}

	fmt.Fprintln(ui.out, "\nPower context: power concentration is risk; oversight strength is protection.")
	powerRisk, err := ui.lawRating("Power-concentration risk (P)")
	if err != nil {
		return err
	}
	oversightStrength, err := ui.lawRating("Oversight strength (O)")
	if err != nil {
		return err
	}

	fmt.Fprintln(ui.out, "\nVisible processing scope")
	fmt.Fprintf(ui.out, "  Instrument: %s | Jurisdiction/context: %s | Clause: %d bytes\n", instrumentType, jurisdiction, len([]byte(clause)))
	fmt.Fprintln(ui.out, "  Input: public/non-sensitive | Memory: process only | Remote: denied | Legal authority: none")
	confirmed, err := ui.confirm("Calculate the screening report from these reviewer-supplied ratings")
	if err != nil {
		return err
	}
	if !confirmed {
		fmt.Fprintln(ui.out, "Canceled. No Law Clarity report was produced.")
		return nil
	}

	report, err := lawclarity.Evaluate(lawclarity.Request{
		RequestID:                   newRequestID(),
		CapabilityID:                lawclarity.CapabilityID,
		UserID:                      "declared-local-reviewer",
		CreatedAt:                   time.Now(),
		Title:                       title,
		Jurisdiction:                jurisdiction,
		InstrumentType:              instrumentType,
		SourceReference:             sourceReference,
		Purpose:                     purpose,
		ClauseText:                  clause,
		PublicNonSensitiveConfirmed: true,
		Quality: lawclarity.QualityRatings{
			Clarity:         clarity,
			Specificity:     specificity,
			Fairness:        fairness,
			Consistency:     consistency,
			Accountability:  accountability,
			LowLoopholeRisk: lowLoophole,
		},
		GrayZone: lawclarity.GrayZoneRatings{
			VagueLanguageRisk:     vagueRisk,
			DefinitionRisk:        definitionRisk,
			ContradictionRisk:     contradictionRisk,
			EnforcementDiscretion: discretionRisk,
			ExceptionBoundaryRisk: exceptionRisk,
		},
		Power: lawclarity.PowerContext{
			PowerConcentrationRisk: powerRisk,
			OversightStrength:      oversightStrength,
		},
	})
	if err != nil {
		fmt.Fprintf(ui.out, "Law Clarity request rejected safely: %v\n", err)
		return nil
	}
	return ui.showLawClarityReport(report)
}

func (ui *UI) lawRating(label string) (lawclarity.Rating, error) {
	var score int
	for {
		value, err := ui.prompt(label+" score [0-100]", true)
		if err != nil {
			return lawclarity.Rating{}, err
		}
		score, err = strconv.Atoi(value)
		if err == nil && score >= 0 && score <= 100 {
			break
		}
		fmt.Fprintln(ui.out, "Enter a whole number from 0 through 100.")
	}
	rationale, err := ui.prompt(label+" rationale/evidence note", true)
	if err != nil {
		return lawclarity.Rating{}, err
	}
	return lawclarity.Rating{Score: score, Rationale: rationale}, nil
}

func (ui *UI) showLawClarityReport(report lawclarity.Report) error {
	fmt.Fprintln(ui.out, "\n============================================================")
	fmt.Fprintln(ui.out, "BOOTX LAW CLARITY LOGIC SCREENING REPORT")
	fmt.Fprintln(ui.out, "============================================================")
	fmt.Fprintf(ui.out, "Request: %s\nInstrument: %s | Context: %s\n", report.RequestID, report.InstrumentType, report.Jurisdiction)
	fmt.Fprintf(ui.out, "Notice: %s\n", report.RuntimeNotice)
	fmt.Fprintf(ui.out, "\nLAW QUALITY: %.2f/100\n", report.LawQualityScore)
	fmt.Fprintf(ui.out, "QUALITY BAND: %s\n", report.QualityBand)
	fmt.Fprintf(ui.out, "GRAY-ZONE RISK: %.2f/100\n", report.GrayZoneRiskScore)
	fmt.Fprintf(ui.out, "MANIPULATION INDEX: %.2f/100 (experimental; not a probability)\n", report.ManipulationRiskIndex)
	fmt.Fprintf(ui.out, "STRICT GATE: %s | RIGHTS GATE: %s\n", report.StrictGoodLawGate.Status, report.HumanRightsGate.Status)
	fmt.Fprintf(ui.out, "DISPOSITION: %s\n", report.Disposition)

	if len(report.VisiblePhraseHits) > 0 {
		fmt.Fprintln(ui.out, "\nVISIBLE AMBIGUITY PHRASE HITS (context review required)")
		for _, hit := range report.VisiblePhraseHits {
			fmt.Fprintf(ui.out, "  • %q: %d\n", hit.Phrase, hit.Count)
		}
	}
	printStrings(ui.out, "FINDINGS", report.Findings)
	printStrings(ui.out, "REWRITE REQUIREMENTS", report.RewriteRequirements)
	fmt.Fprintln(ui.out, "\nNON-BINDING REWRITE TEMPLATE")
	fmt.Fprintln(ui.out, "  "+report.RewriteTemplate)
	printStrings(ui.out, "BLOCKED CONCLUSIONS", report.BlockedConclusions)
	printStrings(ui.out, "LIMITATIONS", report.Limitations)
	fmt.Fprintln(ui.out, "\nAI DNA RUNTIME CHECKS (report process, not legal certification)")
	for _, check := range report.Assurance {
		fmt.Fprintf(ui.out, "  %-12s %-15s %s\n", check.Dimension, check.Status, check.Basis)
	}
	fmt.Fprintln(ui.out, "\nINPUT RECEIPT")
	fmt.Fprintf(ui.out, "  Reference: %s | Bytes: %d | Source: %s\n", report.InputReceipt.SourceReference, report.InputReceipt.ClauseBytes, report.InputReceipt.SourceStatus)
	fmt.Fprintln(ui.out, "  Remote processing: false | Persistent memory: false")

	showJSON, err := ui.confirm("Show the structured JSON report")
	if err != nil {
		return err
	}
	if showJSON {
		encoded, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(ui.out, string(encoded))
	}
	fmt.Fprintln(ui.out, "\nNo legal decision was made. Qualified human review remains required.")
	return nil
}

func (ui *UI) localDocumentWorkflow() error {
	fmt.Fprintln(ui.out, "\n=== Read-only local workspace document ===")
	fmt.Fprintln(ui.out, "This mode reads one user-selected public UTF-8 .md, .txt, or .json file. It does not authenticate the author or truth of its claims.")
	workspaceRoot, err := ui.prompt("Workspace root", true)
	if err != nil {
		return err
	}
	relativePath, err := ui.prompt("Document path relative to workspace", true)
	if err != nil {
		return err
	}
	publicConfirmed, err := ui.confirm("Have you reviewed this file as public and non-sensitive")
	if err != nil {
		return err
	}
	if !publicConfirmed {
		fmt.Fprintln(ui.out, "Canceled before reading the file. Sensitive or uncertain local documents are not authorized in this mode.")
		return nil
	}
	document, err := evidence.LoadLocalDocument(workspaceRoot, relativePath)
	if err != nil {
		fmt.Fprintf(ui.out, "Document rejected safely: %v\n", err)
		return nil
	}

	fmt.Fprintln(ui.out, "\nRead-only integrity receipt")
	fmt.Fprintf(ui.out, "  Reference: %s\n  Bytes: %d\n  SHA-256: %s\n  Modified UTC: %s\n", document.Reference, document.ByteLength, document.SHA256, document.ModifiedAt.Format(time.RFC3339Nano))
	fmt.Fprintln(ui.out, "  Integrity: selected bytes hashed | Origin/claims: not authenticated")
	goal, err := ui.prompt("Goal", true)
	if err != nil {
		return err
	}
	question, err := ui.prompt("Direct question", true)
	if err != nil {
		return err
	}
	prioritiesText, err := ui.prompt("Priorities, comma-separated (optional)", false)
	if err != nil {
		return err
	}
	fmt.Fprintln(ui.out, "\nVisible processing scope")
	fmt.Fprintf(ui.out, "  Public document: %s | %d bytes\n", document.Reference, document.ByteLength)
	fmt.Fprintln(ui.out, "  Read-only: yes | Memory: none | Remote processing: denied | External action: absent")
	confirmed, err := ui.confirm("Process this document scope")
	if err != nil {
		return err
	}
	if !confirmed {
		fmt.Fprintln(ui.out, "Canceled. The selected document was not processed further.")
		return nil
	}

	request := model.Request{
		RequestID:       newRequestID(),
		CapabilityID:    model.CapabilityID,
		UserID:          "declared-local-operator",
		CreatedAt:       time.Now(),
		Goal:            goal,
		Question:        question,
		SelectedContent: document.Content,
		ContentSource: model.ContentSource{
			Type:              "local_workspace_file",
			Reference:         document.Reference,
			IntegrityVerified: true,
			SHA256:            document.SHA256,
			ByteLength:        document.ByteLength,
			ModifiedAt:        document.ModifiedAt.Format(time.RFC3339Nano),
		},
		DataClass:        model.DataPublic,
		DeclaredDomain:   model.DomainStudy,
		MemoryPermission: "none",
		RemotePermission: "deny",
		OutputPreference: "standard",
		UserPriorities:   splitComma(prioritiesText),
		Synthetic:        false,
	}
	packet, err := ui.engine.Process(request)
	if err != nil {
		fmt.Fprintf(ui.out, "Document request rejected safely: %v\n", err)
		return nil
	}
	return ui.showPacket(packet)
}

func (ui *UI) personalWorkflow() error {
	fmt.Fprintln(ui.out, "\n=== New personal decision ===")
	goal, err := ui.prompt("Goal", true)
	if err != nil {
		return err
	}
	question, err := ui.prompt("Direct question", true)
	if err != nil {
		return err
	}
	domain, err := ui.chooseDomain()
	if err != nil {
		return err
	}
	dataClass, err := ui.chooseDataClass()
	if err != nil {
		return err
	}
	synthetic, err := ui.confirm("Is this entirely synthetic/test data")
	if err != nil {
		return err
	}
	if dataClass == model.DataSensitive && !synthetic {
		fmt.Fprintln(ui.out, "Real sensitive input is not authorized before independent security review. No content will be requested.")
		return nil
	}
	content, err := ui.multiline("Selected content (finish with a single period on its own line; use an immediate period for none)")
	if err != nil {
		return err
	}
	preference, err := ui.promptDefault("Output preference [concise/standard/detailed/checklist/comparison]", "standard")
	if err != nil {
		return err
	}
	prioritiesText, err := ui.prompt("Priorities, comma-separated (optional)", false)
	if err != nil {
		return err
	}
	priorities := splitComma(prioritiesText)

	fmt.Fprintln(ui.out, "\nVisible processing scope")
	fmt.Fprintf(ui.out, "  Goal: %s\n  Domain: %s\n  Data class: %s\n  Synthetic: %t\n  Selected content: %d bytes\n", goal, domain, dataClass, synthetic, len([]byte(content)))
	fmt.Fprintln(ui.out, "  Memory: session only | Remote processing: denied | External action: absent")
	confirmed, err := ui.confirm("Process this scope")
	if err != nil {
		return err
	}
	if !confirmed {
		fmt.Fprintln(ui.out, "Canceled. The selected content was not processed further.")
		return nil
	}

	request := model.Request{
		RequestID:        newRequestID(),
		CapabilityID:     model.CapabilityID,
		UserID:           "owner-local",
		CreatedAt:        time.Now(),
		Goal:             goal,
		Question:         question,
		SelectedContent:  content,
		ContentSource:    model.ContentSource{Type: "user_selected_content", OriginVerified: false},
		DataClass:        dataClass,
		DeclaredDomain:   domain,
		MemoryPermission: "session",
		RemotePermission: "deny",
		OutputPreference: preference,
		UserPriorities:   priorities,
		Synthetic:        synthetic,
	}
	packet, err := ui.engine.Process(request)
	if err != nil {
		fmt.Fprintf(ui.out, "Request rejected safely: %v\n", err)
		return nil
	}
	return ui.showPacket(packet)
}

func (ui *UI) warningWorkflow() error {
	fmt.Fprintln(ui.out, "\n=== Forecast/disaster assessment ===")
	fmt.Fprintln(ui.out, "This unvalidated workflow accepts synthetic/test events only and performs no network lookup.")
	continueSynthetic, err := ui.confirm("Continue with a synthetic warning exercise")
	if err != nil {
		return err
	}
	if !continueSynthetic {
		fmt.Fprintln(ui.out, "Canceled. For a real immediate danger, move away from danger when safe and use a known responsible authority or emergency channel.")
		return nil
	}
	eventID, err := ui.prompt("Event identifier", true)
	if err != nil {
		return err
	}
	hazard, err := ui.prompt("Hazard type", true)
	if err != nil {
		return err
	}
	officialStatus, err := ui.choose("Official status", []string{"active", "none_found", "unavailable"})
	if err != nil {
		return err
	}
	issuer, err := ui.prompt("Issuing authority (or 'not supplied')", true)
	if err != nil {
		return err
	}
	authenticated, err := ui.confirm("For this synthetic exercise, do you declare that the authority was authenticated through a known official source")
	if err != nil {
		return err
	}
	messageStatus, err := ui.choose("Message status", []string{"Actual", "Test", "Exercise", "Draft"})
	if err != nil {
		return err
	}
	messageType, err := ui.choose("Message type", []string{"Alert", "Update", "Cancel"})
	if err != nil {
		return err
	}
	issuedAt, err := ui.prompt("Issued time with timezone (optional)", false)
	if err != nil {
		return err
	}
	expiresAt, err := ui.prompt("Expiry time with timezone (optional)", false)
	if err != nil {
		return err
	}
	affectedArea, err := ui.prompt("Affected area as written by source (optional)", false)
	if err != nil {
		return err
	}
	areaMatch, err := ui.choose("Personal area match", []string{"inside", "near", "outside", "unknown"})
	if err != nil {
		return err
	}
	urgency, err := ui.choose("Urgency", []string{"Immediate", "Expected", "Future", "Past", "Unknown"})
	if err != nil {
		return err
	}
	severity, err := ui.choose("Severity", []string{"Extreme", "Severe", "Moderate", "Minor", "Unknown"})
	if err != nil {
		return err
	}
	certainty, err := ui.choose("Certainty", []string{"Observed", "Likely", "Possible", "Unlikely", "Unknown"})
	if err != nil {
		return err
	}
	evidenceTier, err := ui.choose("Evidence verification tier", []string{"V0", "V1", "V2", "V3", "V4"})
	if err != nil {
		return err
	}
	integrity, err := ui.choose("Source integrity", []string{"pass", "fail", "unknown"})
	if err != nil {
		return err
	}
	conflict, err := ui.confirm("Is there a material source conflict")
	if err != nil {
		return err
	}
	stale, err := ui.confirm("Is the information stale or expired")
	if err != nil {
		return err
	}
	directDanger, err := ui.confirm("Are you directly observing immediate danger")
	if err != nil {
		return err
	}
	instruction, err := ui.multiline("Official instruction or selected bulletin excerpt (finish with a single period)")
	if err != nil {
		return err
	}

	fmt.Fprintln(ui.out, "\nVisible processing scope")
	fmt.Fprintf(ui.out, "  Event: %s | Hazard: %s | Official status: %s\n", eventID, hazard, officialStatus)
	fmt.Fprintf(ui.out, "  Issuer authentication declared in exercise: %t | Area match: %s | Evidence: %s | Integrity: %s\n", authenticated, areaMatch, evidenceTier, integrity)
	fmt.Fprintln(ui.out, "  Location: manually declared match only; not persisted | Network lookup: absent | Broadcasting: absent")
	confirmed, err := ui.confirm("Process this warning scope")
	if err != nil {
		return err
	}
	if !confirmed {
		fmt.Fprintln(ui.out, "Canceled. The selected warning data was not processed further.")
		return nil
	}

	warningInput := &model.WarningInput{
		EventID:                eventID,
		HazardType:             hazard,
		OfficialStatus:         officialStatus,
		Issuer:                 issuer,
		AuthorityAuthenticated: authenticated,
		MessageStatus:          messageStatus,
		MessageType:            messageType,
		IssuedAt:               issuedAt,
		ExpiresAt:              expiresAt,
		AffectedArea:           affectedArea,
		AreaMatch:              areaMatch,
		Urgency:                urgency,
		Severity:               severity,
		Certainty:              certainty,
		Instruction:            instruction,
		EvidenceTier:           evidenceTier,
		IntegrityStatus:        integrity,
		SourceConflict:         conflict,
		Stale:                  stale,
		DirectDanger:           directDanger,
	}
	request := model.Request{
		RequestID:        newRequestID(),
		CapabilityID:     model.CapabilityID,
		UserID:           "owner-local",
		CreatedAt:        time.Now(),
		Goal:             "Determine a safe personal posture for the selected " + hazard + " information",
		Question:         "Should I act, protect, prepare, monitor, verify, wait, or seek the responsible authority?",
		SelectedContent:  instruction,
		ContentSource:    model.ContentSource{Type: "user_selected_warning", OriginVerified: authenticated},
		DataClass:        model.DataPersonal,
		DeclaredDomain:   model.DomainEmergency,
		MemoryPermission: "session",
		RemotePermission: "deny",
		OutputPreference: "warning_card",
		UserPriorities:   []string{"life safety", "minimal false alarm", "official verification", "family readiness"},
		Synthetic:        true,
		Warning:          warningInput,
	}
	packet, err := ui.engine.Process(request)
	if err != nil {
		fmt.Fprintf(ui.out, "Warning request rejected safely: %v\n", err)
		return nil
	}
	return ui.showPacket(packet)
}

func (ui *UI) showPacket(packet model.Packet) error {
	fmt.Fprintln(ui.out, "\n============================================================")
	fmt.Fprintln(ui.out, "BOOTX PERSONAL DECISION PACKET")
	fmt.Fprintln(ui.out, "============================================================")
	fmt.Fprintf(ui.out, "Request: %s\nClass: %s | Mode: %s\n", packet.RequestID, packet.DecisionClass, packet.ResponseMode)
	fmt.Fprintf(ui.out, "Goal: %s\nNotice: %s\n", packet.GoalUnderstood, packet.RuntimeNotice)

	if packet.Warning != nil {
		w := packet.Warning
		fmt.Fprintln(ui.out, "\n--- WARNING CARD ---")
		fmt.Fprintf(ui.out, "Event: %s | Hazard: %s\n", w.EventID, w.HazardType)
		fmt.Fprintf(ui.out, "Official status: %s | Issuer: %s\n", w.OfficialStatus, w.Issuer)
		fmt.Fprintf(ui.out, "Issued: %s | Expires: %s\n", fallback(w.IssuedAt, "not supplied"), fallback(w.ExpiresAt, "not supplied"))
		fmt.Fprintf(ui.out, "Area match: %s | Urgency: %s | Severity: %s | Certainty: %s\n", w.AreaMatch, w.Urgency, w.Severity, w.Certainty)
		fmt.Fprintf(ui.out, "BOOTX LEVEL: %s — %s\n", w.Level, w.LevelLabel)
		fmt.Fprintf(ui.out, "Evidence: %s | Decision posture: %s\n", w.EvidenceTier, w.DecisionPosture)
		if w.OfficialInstruction != "" {
			fmt.Fprintf(ui.out, "Official/user-supplied instruction: %s\n", w.OfficialInstruction)
		}
		fmt.Fprintf(ui.out, "Next review: %s\n", w.NextReview)
	}

	printFindings(ui.out, "OBSERVATIONS (indicators, not verdicts)", packet.Observations)
	printStrings(ui.out, "UNKNOWNS", packet.Unknowns)
	fmt.Fprintln(ui.out, "\nOPTIONS")
	for _, option := range packet.Options {
		fmt.Fprintf(ui.out, "  [%s] %s\n", option.OptionID, option.Summary)
		if len(option.Benefits) > 0 {
			fmt.Fprintf(ui.out, "      Benefit: %s\n", strings.Join(option.Benefits, "; "))
		}
		if len(option.Risks) > 0 {
			fmt.Fprintf(ui.out, "      Risk: %s\n", strings.Join(option.Risks, "; "))
		}
	}
	fmt.Fprintln(ui.out, "\nADVISORY RESULT")
	fmt.Fprintf(ui.out, "  Status: %s", packet.Recommendation.Status)
	if packet.Recommendation.OptionID != "" {
		fmt.Fprintf(ui.out, " | Option: %s", packet.Recommendation.OptionID)
	}
	fmt.Fprintf(ui.out, "\n  Basis: %s\n", packet.Recommendation.Basis)
	fmt.Fprintf(ui.out, "  Next safe step: %s\n", packet.NextSafeStep)

	fmt.Fprintln(ui.out, "\nAI DNA RUNTIME CHECKS (not certification or forecast probability)")
	for _, check := range packet.Assurance {
		fmt.Fprintf(ui.out, "  %-12s %-15s %s\n", check.Dimension, check.Status, check.Basis)
	}
	printStrings(ui.out, "BLOCKED EXTERNAL ACTIONS", packet.BlockedActions)
	printStrings(ui.out, "LIMITATIONS", packet.Limitations)
	fmt.Fprintln(ui.out, "\nEVIDENCE RECEIPT")
	fmt.Fprintf(ui.out, "  Source: %s | Reference: %s\n", packet.EvidenceReceipt.SourceType, fallback(packet.EvidenceReceipt.Reference, "not supplied"))
	fmt.Fprintf(ui.out, "  Integrity verified: %t | Bytes: %d | Origin: %s\n", packet.EvidenceReceipt.IntegrityVerified, packet.EvidenceReceipt.ByteLength, packet.EvidenceReceipt.OriginStatus)
	if packet.EvidenceReceipt.SHA256 != "" {
		fmt.Fprintf(ui.out, "  SHA-256: %s\n", packet.EvidenceReceipt.SHA256)
	}
	fmt.Fprintln(ui.out, "\nDATA RECEIPT")
	fmt.Fprintf(ui.out, "  Memory used: %t | Remote processing: %t | Synthetic input: %t\n", packet.DataReceipt.MemoryUsed, packet.DataReceipt.RemoteProcessing, packet.DataReceipt.Synthetic)
	fmt.Fprintf(ui.out, "  Raw retention: %s\n", packet.DataReceipt.RawRetention)
	fmt.Fprintf(ui.out, "  Location: %s\n", packet.DataReceipt.LocationUse)

	showJSON, err := ui.confirm("Show the structured JSON packet")
	if err != nil {
		return err
	}
	if showJSON {
		encoded, err := json.MarshalIndent(packet, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(ui.out, string(encoded))
	}
	fmt.Fprintln(ui.out, "\nDecision remains yours. This session is not saved by the program.")
	return nil
}

func (ui *UI) chooseDomain() (model.Domain, error) {
	value, err := ui.choose("Domain", []string{"general", "study", "digital_safety", "household", "health", "finance", "legal", "relationship", "faith", "emergency", "unknown"})
	return model.Domain(value), err
}

func (ui *UI) chooseDataClass() (model.DataClass, error) {
	value, err := ui.choose("Data class", []string{"public", "personal", "sensitive", "prohibited"})
	return model.DataClass(value), err
}

func (ui *UI) choose(label string, values []string) (string, error) {
	for {
		fmt.Fprintf(ui.out, "%s:\n", label)
		for i, value := range values {
			fmt.Fprintf(ui.out, "  %d  %s\n", i+1, value)
		}
		answer, err := ui.prompt("Choose", true)
		if err != nil {
			return "", err
		}
		for i, value := range values {
			if answer == fmt.Sprint(i+1) || strings.EqualFold(answer, value) {
				return value, nil
			}
		}
		fmt.Fprintln(ui.out, "Choose one of the numbered values.")
	}
}

func (ui *UI) prompt(label string, required bool) (string, error) {
	for {
		fmt.Fprintf(ui.out, "%s: ", label)
		line, err := ui.reader.ReadString('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			return "", err
		}
		value := strings.TrimSpace(line)
		if required && value == "" {
			if errors.Is(err, io.EOF) {
				return "", io.EOF
			}
			fmt.Fprintln(ui.out, "A value is required.")
			continue
		}
		if errors.Is(err, io.EOF) && value == "" {
			return "", io.EOF
		}
		return value, nil
	}
}

func (ui *UI) promptDefault(label, defaultValue string) (string, error) {
	value, err := ui.prompt(label, false)
	if err != nil {
		return "", err
	}
	if value == "" {
		return defaultValue, nil
	}
	return value, nil
}

func (ui *UI) multiline(label string) (string, error) {
	fmt.Fprintln(ui.out, label+":")
	var lines []string
	for {
		line, err := ui.reader.ReadString('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			return "", err
		}
		trimmed := strings.TrimRight(line, "\r\n")
		if trimmed == "." {
			return strings.Join(lines, "\n"), nil
		}
		if trimmed != "" || len(lines) > 0 {
			lines = append(lines, trimmed)
		}
		if errors.Is(err, io.EOF) {
			return strings.Join(lines, "\n"), nil
		}
	}
}

func (ui *UI) confirm(label string) (bool, error) {
	for {
		answer, err := ui.prompt(label+" [y/N]", false)
		if err != nil {
			return false, err
		}
		switch strings.ToLower(answer) {
		case "y", "yes":
			return true, nil
		case "", "n", "no":
			return false, nil
		default:
			fmt.Fprintln(ui.out, "Please answer y or n.")
		}
	}
}

func (ui *UI) header() {
	fmt.Fprintf(ui.out, "BOOTX PERSONAL COMPANION MVP %s\n", model.Version)
	fmt.Fprintln(ui.out, "Input → deterministic evidence/policy process → advisory output → Joni decides")
	fmt.Fprintln(ui.out, "Offline | session-only | no external actions | no robotics | not emergency authority")
}

func (ui *UI) boundaries() {
	fmt.Fprintln(ui.out, "\nSafety boundaries")
	fmt.Fprintln(ui.out, "  • No AI model is connected in this baseline.")
	fmt.Fprintln(ui.out, "  • No network, browser, message, payment, account, call, device, or motor capability exists.")
	fmt.Fprintln(ui.out, "  • Selected content is treated as untrusted data and cannot change policy.")
	fmt.Fprintln(ui.out, "  • Health/legal output is educational and defers to qualified professionals.")
	fmt.Fprintln(ui.out, "  • Disaster output preserves official distinctions and never becomes an official alert.")
	fmt.Fprintln(ui.out, "  • Law Clarity scores are reviewer-supplied research screening, not legal advice, validity, or authority.")
	fmt.Fprintln(ui.out, "  • Warning assessment and sensitive scenarios are synthetic-only until security and human-study gates pass.")
	fmt.Fprintln(ui.out, "  • The process retains no raw content after exit unless the operator redirects output externally.")
}

func (ui *UI) clear() {
	fmt.Fprint(ui.out, "\x1b[2J\x1b[H")
}

func printFindings(out io.Writer, title string, values []model.Finding) {
	if len(values) == 0 {
		return
	}
	fmt.Fprintln(out, "\n"+title)
	for _, value := range values {
		fmt.Fprintf(out, "  • %s [%s; %s]\n", value.Claim, value.Source, value.Status)
	}
}

func printStrings(out io.Writer, title string, values []string) {
	if len(values) == 0 {
		return
	}
	fmt.Fprintln(out, "\n"+title)
	for _, value := range values {
		fmt.Fprintln(out, "  • "+value)
	}
}

func splitComma(value string) []string {
	var result []string
	for _, part := range strings.Split(value, ",") {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func newRequestID() string {
	b := make([]byte, 6)
	if _, err := rand.Read(b); err == nil {
		return "local-" + hex.EncodeToString(b)
	}
	return fmt.Sprintf("local-%d", time.Now().UnixNano())
}

func fallback(value, fallbackValue string) string {
	if strings.TrimSpace(value) == "" {
		return fallbackValue
	}
	return value
}
