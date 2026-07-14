package model

import "time"

const (
	CapabilityID = "assist.personal-decision.v1"
	Version      = "0.2.0-dev"
)

type DataClass string

const (
	DataPublic     DataClass = "public"
	DataPersonal   DataClass = "personal"
	DataSensitive  DataClass = "sensitive"
	DataProhibited DataClass = "prohibited"
)

type Domain string

const (
	DomainGeneral       Domain = "general"
	DomainStudy         Domain = "study"
	DomainDigitalSafety Domain = "digital_safety"
	DomainHousehold     Domain = "household"
	DomainHealth        Domain = "health"
	DomainFinance       Domain = "finance"
	DomainLegal         Domain = "legal"
	DomainRelationship  Domain = "relationship"
	DomainFaith         Domain = "faith"
	DomainEmergency     Domain = "emergency"
	DomainUnknown       Domain = "unknown"
)

type DecisionClass string

const (
	D0Informational DecisionClass = "D0"
	D1Reversible    DecisionClass = "D1"
	D2Consequential DecisionClass = "D2"
	D3Qualified     DecisionClass = "D3"
	D4Emergency     DecisionClass = "D4"
	D5Prohibited    DecisionClass = "D5"
)

type ResponseMode string

const (
	ModeInform         ResponseMode = "INFORM"
	ModeCompare        ResponseMode = "COMPARE"
	ModeVerify         ResponseMode = "VERIFY"
	ModeClarify        ResponseMode = "CLARIFY"
	ModeAbstain        ResponseMode = "ABSTAIN"
	ModeMonitor        ResponseMode = "MONITOR"
	ModePrepare        ResponseMode = "PREPARE"
	ModeProtect        ResponseMode = "PROTECT"
	ModeWait           ResponseMode = "WAIT"
	ModeUrgentGuidance ResponseMode = "URGENT_GUIDANCE"
	ModeBlock          ResponseMode = "BLOCK"
	ModeDegraded       ResponseMode = "DEGRADED"
)

type ContentSource struct {
	Type              string `json:"type"`
	OriginVerified    bool   `json:"origin_verified"`
	Reference         string `json:"reference,omitempty"`
	IntegrityVerified bool   `json:"integrity_verified,omitempty"`
	SHA256            string `json:"sha256,omitempty"`
	ByteLength        int    `json:"byte_length,omitempty"`
	ModifiedAt        string `json:"modified_at,omitempty"`
}

type WarningInput struct {
	EventID                string `json:"event_id"`
	HazardType             string `json:"hazard_type"`
	OfficialStatus         string `json:"official_status"`
	Issuer                 string `json:"issuer"`
	OfficialURL            string `json:"official_url,omitempty"`
	AuthorityAuthenticated bool   `json:"authority_authenticated"`
	MessageStatus          string `json:"message_status"`
	MessageType            string `json:"message_type"`
	IssuedAt               string `json:"issued_at,omitempty"`
	UpdatedAt              string `json:"updated_at,omitempty"`
	ExpiresAt              string `json:"expires_at,omitempty"`
	AffectedArea           string `json:"affected_area,omitempty"`
	AreaMatch              string `json:"area_match"`
	Urgency                string `json:"urgency"`
	Severity               string `json:"severity"`
	Certainty              string `json:"certainty"`
	Instruction            string `json:"instruction,omitempty"`
	EvidenceTier           string `json:"evidence_tier"`
	IntegrityStatus        string `json:"integrity_status"`
	SourceConflict         bool   `json:"source_conflict"`
	Stale                  bool   `json:"stale"`
	DirectDanger           bool   `json:"direct_danger"`
}

type Request struct {
	RequestID        string        `json:"request_id"`
	CapabilityID     string        `json:"capability_id"`
	UserID           string        `json:"user_id"`
	CreatedAt        time.Time     `json:"created_at"`
	Goal             string        `json:"goal"`
	Question         string        `json:"question"`
	SelectedContent  string        `json:"selected_content"`
	ContentSource    ContentSource `json:"content_source"`
	DataClass        DataClass     `json:"data_class"`
	DeclaredDomain   Domain        `json:"declared_domain"`
	MemoryPermission string        `json:"memory_permission"`
	RemotePermission string        `json:"remote_permission"`
	OutputPreference string        `json:"output_preference"`
	UserPriorities   []string      `json:"user_priorities,omitempty"`
	Synthetic        bool          `json:"synthetic"`
	Warning          *WarningInput `json:"warning,omitempty"`
}

type Finding struct {
	Claim  string `json:"claim"`
	Source string `json:"source"`
	Status string `json:"status"`
}

type Option struct {
	OptionID       string   `json:"option_id"`
	Summary        string   `json:"summary"`
	Benefits       []string `json:"benefits,omitempty"`
	Risks          []string `json:"risks,omitempty"`
	Reversibility  string   `json:"reversibility"`
	ExternalEffect string   `json:"external_effect"`
}

type Recommendation struct {
	Status   string `json:"status"`
	OptionID string `json:"option_id,omitempty"`
	Basis    string `json:"basis"`
}

type AssuranceCheck struct {
	Dimension string `json:"dimension"`
	Status    string `json:"status"`
	Basis     string `json:"basis"`
}

type DataReceipt struct {
	MemoryUsed       bool   `json:"memory_used"`
	RemoteProcessing bool   `json:"remote_processing"`
	Synthetic        bool   `json:"synthetic"`
	RawRetention     string `json:"raw_content_retention"`
	LocationUse      string `json:"location_use"`
}

type EvidenceReceipt struct {
	SourceType        string `json:"source_type"`
	Reference         string `json:"reference,omitempty"`
	IntegrityVerified bool   `json:"integrity_verified"`
	SHA256            string `json:"sha256,omitempty"`
	ByteLength        int    `json:"byte_length"`
	ModifiedAt        string `json:"modified_at,omitempty"`
	OriginStatus      string `json:"origin_status"`
}

type WarningCard struct {
	EventID             string   `json:"event_id"`
	HazardType          string   `json:"hazard_type"`
	OfficialStatus      string   `json:"official_status"`
	Issuer              string   `json:"issuer"`
	OfficialURL         string   `json:"official_url,omitempty"`
	IssuedAt            string   `json:"issued_at,omitempty"`
	UpdatedAt           string   `json:"updated_at,omitempty"`
	ExpiresAt           string   `json:"expires_at,omitempty"`
	AffectedArea        string   `json:"affected_area,omitempty"`
	AreaMatch           string   `json:"area_match"`
	Urgency             string   `json:"urgency"`
	Severity            string   `json:"severity"`
	Certainty           string   `json:"certainty"`
	Level               string   `json:"level"`
	LevelLabel          string   `json:"level_label"`
	EvidenceTier        string   `json:"evidence_tier"`
	DecisionPosture     string   `json:"decision_posture"`
	OfficialInstruction string   `json:"official_instruction,omitempty"`
	Observed            []string `json:"observed,omitempty"`
	Forecast            []string `json:"forecast,omitempty"`
	Unknowns            []string `json:"unknowns,omitempty"`
	NextSafeStep        string   `json:"next_safe_step"`
	NextReview          string   `json:"next_review"`
	Limitations         []string `json:"limitations,omitempty"`
}

type Packet struct {
	RequestID       string           `json:"request_id"`
	CapabilityID    string           `json:"capability_id"`
	GeneratedAt     time.Time        `json:"generated_at"`
	RuntimeNotice   string           `json:"runtime_notice"`
	DecisionClass   DecisionClass    `json:"decision_class"`
	ResponseMode    ResponseMode     `json:"response_mode"`
	GoalUnderstood  string           `json:"goal_understood"`
	Observations    []Finding        `json:"observations"`
	Assumptions     []string         `json:"assumptions"`
	Unknowns        []string         `json:"unknowns"`
	Options         []Option         `json:"options"`
	Recommendation  Recommendation   `json:"recommendation"`
	NextSafeStep    string           `json:"next_safe_step"`
	BlockedActions  []string         `json:"blocked_actions"`
	Limitations     []string         `json:"limitations"`
	Warning         *WarningCard     `json:"warning,omitempty"`
	Assurance       []AssuranceCheck `json:"ai_dna_runtime_checks"`
	EvidenceReceipt EvidenceReceipt  `json:"evidence_receipt"`
	DataReceipt     DataReceipt      `json:"data_receipt"`
	UserDecision    *string          `json:"user_decision"`
}
