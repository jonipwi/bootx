package policy

import (
	"embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/jonipwi/bootx/prototype/personal-companion/internal/model"
)

//go:embed config/*.json
var configFS embed.FS

type Rules struct {
	Version                 string   `json:"version"`
	MaxSelectedContentBytes int      `json:"max_selected_content_bytes"`
	ProhibitedIntentPhrases []string `json:"prohibited_intent_phrases"`
	CredentialIndicators    []string `json:"credential_indicators"`
	PaymentIndicators       []string `json:"payment_indicators"`
	UrgencyIndicators       []string `json:"urgency_indicators"`
	InjectionIndicators     []string `json:"injection_indicators"`
}

type DecisionClasses struct {
	Version string `json:"version"`
	Classes []struct {
		Code      string `json:"code"`
		Name      string `json:"name"`
		MaxEffect string `json:"max_effect"`
	} `json:"classes"`
}

type Analysis struct {
	HasURL         bool
	HasCredential  bool
	HasPayment     bool
	HasUrgency     bool
	HasInjection   bool
	Prohibited     bool
	MatchedSignals []string
}

type Engine struct {
	Rules   Rules
	Classes DecisionClasses
}

var urlPattern = regexp.MustCompile(`(?i)\b(?:https?://|www\.)\S+`)

func Load() (*Engine, error) {
	var rules Rules
	if err := decodeEmbedded("config/policy-rules.json", &rules); err != nil {
		return nil, err
	}
	var classes DecisionClasses
	if err := decodeEmbedded("config/decision-classes.json", &classes); err != nil {
		return nil, err
	}
	if rules.MaxSelectedContentBytes <= 0 || len(classes.Classes) != 6 {
		return nil, fmt.Errorf("invalid embedded policy configuration")
	}
	return &Engine{Rules: rules, Classes: classes}, nil
}

func decodeEmbedded(name string, out any) error {
	b, err := configFS.ReadFile(name)
	if err != nil {
		return fmt.Errorf("read embedded %s: %w", name, err)
	}
	if err := json.Unmarshal(b, out); err != nil {
		return fmt.Errorf("decode embedded %s: %w", name, err)
	}
	return nil
}

func (e *Engine) Analyze(request model.Request) Analysis {
	combined := strings.ToLower(strings.Join([]string{request.Goal, request.Question, request.SelectedContent}, "\n"))
	a := Analysis{HasURL: urlPattern.MatchString(combined)}
	a.HasCredential = containsAny(combined, e.Rules.CredentialIndicators)
	a.HasPayment = containsAny(combined, e.Rules.PaymentIndicators)
	a.HasUrgency = containsAny(combined, e.Rules.UrgencyIndicators)
	a.HasInjection = containsAny(combined, e.Rules.InjectionIndicators)
	a.Prohibited = containsAny(strings.ToLower(request.Goal+"\n"+request.Question), e.Rules.ProhibitedIntentPhrases)

	if a.HasURL {
		a.MatchedSignals = append(a.MatchedSignals, "link or URL text")
	}
	if a.HasCredential {
		a.MatchedSignals = append(a.MatchedSignals, "credential-related language")
	}
	if a.HasPayment {
		a.MatchedSignals = append(a.MatchedSignals, "payment or financial language")
	}
	if a.HasUrgency {
		a.MatchedSignals = append(a.MatchedSignals, "urgency or pressure language")
	}
	if a.HasInjection {
		a.MatchedSignals = append(a.MatchedSignals, "instruction-like content treated as untrusted data")
	}
	return a
}

func containsAny(text string, terms []string) bool {
	for _, term := range terms {
		if strings.Contains(text, strings.ToLower(term)) {
			return true
		}
	}
	return false
}
