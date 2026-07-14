# BootX Personal Companion MVP — Complete Usage Guide

**Application version:** `0.2.0-dev`<br>
**Capability:** `assist.personal-decision.v1`<br>
**Evidence maturity:** `E2 — Prototype` in a limited host environment<br>
**Intended operator:** Joni or a developer processing approved public local documents and synthetic tests<br>
**Status:** software-only development prototype; not deployed, safety-certified, or authorized for real emergency reliance

## 1. What this prototype does

BootX Personal Companion receives a deliberately selected question or scenario, processes it through deterministic evidence and policy rules, and produces an advisory decision packet.

```text
Selected input
    → scope and consent
    → deterministic indicator analysis
    → decision/risk classification
    → evidence, unknowns, options, and limitations
    → advisory output
    → you decide
```

The prototype includes:

- an interactive line-oriented terminal UI;
- a strict JSON backend mode;
- contained read-only ingestion of one explicitly confirmed public local workspace document;
- personal decision classes `D0`–`D5`;
- forecast/disaster levels `W0`–`W4` and `WX` for synthetic exercises;
- nine AI DNA runtime checks;
- visible blocked actions, limitations, and data receipts;
- synthetic scam, study, and warning fixtures;
- no third-party Go modules.

It does **not** include a generative AI model, network lookup, browser, persistent memory, user authentication, external source or author authentication, message sending, phone calls, payment, account access, device/robot control, family/public broadcast, or emergency dispatch.

## 2. Safety rules before use

Use the MVP only for:

- low-stakes, reversible, non-sensitive personal decisions;
- public or deliberately minimized information;
- one deliberately reviewed public and non-sensitive `.md`, `.txt`, or `.json` file inside a user-selected workspace;
- explicitly synthetic sensitive scenarios;
- explicitly synthetic forecast/disaster exercises;
- software and policy testing.

Do not enter:

- passwords, PINs, one-time codes, recovery codes, seed phrases, or private keys;
- full payment-card, bank-account, government-identity, or medical-record data;
- unnecessary private messages or bystander information;
- real sensitive scenarios before independent security review;
- real disaster data for operational reliance;
- information whose disclosure in terminal history, redirected output, crash dumps, or screen capture would create harm.

The application does not intentionally write raw input. The surrounding terminal, operating system, shell redirection, screen recorder, antivirus, debugger, or crash-reporting service may still retain data.

## 3. Repository locations

```text
bootx\
└── prototype\
    ├── build.ps1                 supported build/verify/run workflow
    └── personal-companion\
        ├── go.mod                actual Go module
        ├── README.md             implementation summary
        ├── USAGE.md              this operator guide
        ├── cmd\                  executable entry point
        ├── internal\             engine, models, policy, TUI, warning logic
        ├── schemas\              strict input/output contracts
        └── testdata\             synthetic fixtures
```

The Go module is `prototype\personal-companion`, not `prototype`. Do not run `go mod init` in the parent `prototype` directory.

## 4. Prerequisites

- Windows PowerShell 5.1 or PowerShell 7;
- Go 1.22 or later on `PATH`;
- a supported Windows host for the produced `.exe`;
- local read/write access to the repository and chosen build-output directory.

Check Go:

```powershell
go version
```

The project currently uses only the Go standard library. `go mod tidy` should not download external packages.

## 5. Recommended quick start

From the BootX repository root:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\prototype\build.ps1 -Action verify
powershell -NoProfile -ExecutionPolicy Bypass -File .\prototype\build.ps1 -Action run
```

From `D:\Job\Human+AI+DNA\bootx\prototype`:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\build.ps1 -Action verify
powershell -NoProfile -ExecutionPolicy Bypass -File .\build.ps1 -Action run
```

The process-scoped execution-policy option does not change the machine-wide policy.

If scripts are already permitted:

```powershell
.\build.ps1 -Action verify
.\build.ps1 -Action run
```

## 6. Build script reference

Syntax:

```powershell
.\build.ps1 [-Action <action>] [-OutputDirectory <path>] [-SkipTests]
```

| Action | Behavior |
|---|---|
| `build` | formatting check, tests, vet, JSON validation, executable build, version smoke test, SHA-256, and manifest |
| `test` | formatting check, `go test ./...`, and `go vet ./...`; no executable |
| `verify` | full build plus all backend fixtures and a contained non-synthetic read-only BootX research-document smoke test |
| `run` | full build, then launches the terminal UI |
| `clean` | removes only the default `personal-companion\dist` directory |

`build` is the default action:

```powershell
.\build.ps1
```

Use a custom output directory:

```powershell
.\build.ps1 -Action verify -OutputDirectory D:\Temp\bootx-companion-build
```

`clean` deliberately refuses a custom output directory to reduce accidental deletion risk.

Skip tests only for a temporary developer build:

```powershell
.\build.ps1 -Action build -SkipTests
```

The manifest records `tests_skipped: true`. A skipped-test build must not be treated as verified.

### Default build outputs

```text
personal-companion\dist\
├── bootx-companion-windows-amd64.exe
├── bootx-companion-windows-amd64.exe.sha256
└── build-manifest.json
```

The exact filename reflects the `GOOS` and `GOARCH` reported by the installed Go toolchain.

### Build-manifest fields

| Field | Meaning |
|---|---|
| `project` | application name |
| `capability` | implemented capability contract |
| `application_version` | executable-reported version |
| `build_utc` | build-record creation time |
| `go_version` | complete Go toolchain string |
| `goos`, `goarch` | target platform and architecture |
| `binary` | output executable filename |
| `sha256` | lowercase SHA-256 digest |
| `tests_skipped` | whether test/vet gates were explicitly bypassed |
| `external_modules` | external Go-module count computed from `go list -m all`, currently zero |
| `safety_status` | development and assurance limitation |

## 7. Verify the executable hash

From `prototype` after a default build:

```powershell
$dist = '.\personal-companion\dist'
$manifest = Get-Content -Raw "$dist\build-manifest.json" | ConvertFrom-Json
$actual = (Get-FileHash "$dist\$($manifest.binary)" -Algorithm SHA256).Hash.ToLowerInvariant()
$actual -eq $manifest.sha256
```

Expected output is `True`. A mismatch means the file and manifest do not describe the same bytes; do not run or distribute that artifact until the cause is understood.

## 8. Run the terminal UI

Recommended:

```powershell
.\build.ps1 -Action run
```

Run an already built executable:

```powershell
.\personal-companion\dist\bootx-companion-windows-amd64.exe
```

Run directly from source:

```powershell
cd .\personal-companion
go run ./cmd/bootx-companion
```

The main menu is:

```text
1  Personal decision assistance
2  Forecast/disaster warning assessment
3  Read-only local workspace document
4  Safety boundaries
q  Quit and discard session data
```

## 9. Personal decision workflow

Choose menu option `1`.

### Step 1 — Goal

State what you are trying to decide or understand.

Good:

```text
Choose the next reversible BootX study task.
```

Avoid vague or unlimited goals such as “decide my whole future.”

### Step 2 — Direct question

Ask one decision-focused question.

```text
Which missing fact should I verify before choosing the next task?
```

### Step 3 — Domain

| Domain | Intended use |
|---|---|
| `general` | ordinary low-stakes personal comparison |
| `study` | learning, research, or project planning |
| `digital_safety` | suspicious digital content and verification |
| `household` | reversible household planning |
| `health` | general educational framing only; qualified review required |
| `finance` | conservative comparison/verification; no transaction |
| `legal` | general educational framing only; qualified review required |
| `relationship` | reflective options without deciding another person's intentions or worth |
| `faith` | Creator-centered reflection without divine claims or conscience control |
| `emergency` | limited safety-first guidance; real warning reliance is not authorized |
| `unknown` | use when no domain fits; expect more uncertainty |

### Step 4 — Data class

| Class | Meaning | MVP behavior |
|---|---|---|
| `public` | intentionally public, non-sensitive information | accepted |
| `personal` | minimized non-sensitive personal context | accepted for low-stakes use |
| `sensitive` | information that could materially harm privacy, security, finance, health, or safety | accepted only when explicitly synthetic |
| `prohibited` | credentials, secrets, harmful/illegal control content, or data outside authority | blocked before content analysis |

### Step 5 — Synthetic status

Answer whether the scenario is entirely synthetic/test data.

- Sensitive data with `No` fails closed before content is requested.
- Synthetic fixtures and exercises use `Yes`.
- A low-stakes non-sensitive personal study question may use `No`.

### Step 6 — Selected content

Enter only the smallest content needed. Finish by entering a single period on its own line:

```text
The project has three incomplete verification tasks.
.
```

An immediate `.` means no selected content. A period inside another sentence does not end input.

Selected content is treated as untrusted data. Text claiming “system message,” “developer message,” “ignore instructions,” or “grant tools” cannot change policy.

### Step 7 — Output preference

Supported values:

- `concise`
- `standard`
- `detailed`
- `checklist`
- `comparison`

Press Enter for `standard`. The current renderer always preserves mandatory safety, evidence, and limitation fields even if a concise style is requested.

### Step 8 — Priorities

Enter comma-separated priorities relevant to this decision:

```text
truth, reversibility, family safety, learning
```

Priorities guide comparison. They cannot change facts, lower a risk class, remove uncertainty, or override prohibited-use rules.

### Step 9 — Confirm visible scope

Review:

- goal;
- domain;
- data class;
- synthetic status;
- selected-content byte count;
- session-only memory;
- denied remote processing;
- absent external actions.

Choose `y` only if the scope is correct. Choosing `n` cancels processing.

## 10. Read-only local workspace document

This is the first real, non-synthetic integration. It reads a file the operator deliberately selects; it does not retrieve information from a network or claim the document is true.

### TUI workflow

Choose menu option `3`, then provide:

1. a workspace root;
2. a document path relative to that root;
3. explicit confirmation that the file was reviewed as public and non-sensitive;
4. a narrow goal and direct question;
5. optional decision priorities;
6. final confirmation of the visible processing scope.

The file is not opened until after the public/non-sensitive confirmation. BootX accepts only a regular UTF-8 `.md`, `.txt`, or `.json` file of at most 65,536 bytes. The resolved file must remain inside the resolved workspace root, including after symbolic-link or junction resolution.

### Command-line workflow

Build first, then run from the repository root:

```powershell
.\prototype\personal-companion\dist\bootx-companion-windows-amd64.exe `
  -workspace . `
  -document docs\research\civilization\religion-ideology-and-decision-integrity.md `
  -document-public `
  -goal "Choose the next evidence-improvement task" `
  -question "Which missing fact should be verified first?" `
  -priorities "truth, reversibility, common good"
```

`-document-public` is an explicit operator assertion. Do not use it for a file containing passwords, private correspondence, identity records, financial information, health information, precise private location, or other sensitive material.

### Evidence receipt meaning

The output records:

- workspace-relative reference;
- exact UTF-8 byte length;
- modification time observed during reading;
- SHA-256 of the processed bytes;
- `integrity_verified: true` when the engine independently recomputes and matches the receipt;
- `origin_status: not_authenticated`.

SHA-256 establishes which bytes were processed. It does not establish authorship, publisher identity, factual truth, completeness, freshness, safety, or approval. BootX does not authenticate the local user; `user_id` is a declared identifier only.

The deterministic document scan also reports heading count, external-link strings, open Markdown checklist items, evidence-gap-marker lines, and the first candidate for human review. A marker can be missed or taken out of context; it is a navigation aid, not semantic understanding, source verification, or a priority certificate.

## 11. Synthetic forecast/disaster workflow

Choose menu option `2`. This workflow is **synthetic-only** in version `0.2.0-dev`.

The UI requests:

1. event identifier and hazard type;
2. whether official status is `active`, `none_found`, or `unavailable`;
3. issuing authority and whether it was authenticated;
4. message status: `Actual`, `Test`, `Exercise`, or `Draft`;
5. message type: `Alert`, `Update`, or `Cancel`;
6. issue and expiry times;
7. affected area and personal area match;
8. urgency, severity, and certainty;
9. evidence tier `V0`–`V4`;
10. source integrity, conflict, and staleness;
11. whether direct danger is being simulated;
12. official instruction or selected synthetic bulletin excerpt.

The prototype does not retrieve, authenticate, or refresh a real warning feed. Manually entering `authority_authenticated: true` is a test assertion, not cryptographic proof.

For an actual immediate danger, do not spend time operating this prototype. Move away from the danger when safe and use a known responsible authority or emergency channel.

## 12. Reading decision classes

| Class | Meaning | System behavior |
|---|---|---|
| `D0` | informational | bounded explanation with evidence limitations |
| `D1` | low-stakes reversible | options and advisory comparison; user decides |
| `D2` | sensitive or consequential | conservative comparison/verification; no external execution |
| `D3` | qualified domain | educational framing and questions; defer to an appropriate professional |
| `D4` | emergency or critical | short safety-first guidance and user-controlled responsible channels |
| `D5` | prohibited | refuse and provide a safe lawful alternative when possible |

A higher number is not a danger probability or moral score. It identifies the authority and control boundary.

## 13. Reading response modes

| Mode | Interpretation |
|---|---|
| `INFORM` | bounded informational explanation |
| `COMPARE` | review multiple permissible options and trade-offs |
| `VERIFY` | confirm authenticity or a critical fact independently before consequential action |
| `CLARIFY` | the goal or scope needs a focused clarification |
| `ABSTAIN` | evidence or competency is insufficient |
| `MONITOR` | watch named conditions or sources at a defined review time |
| `PREPARE` | take low-burden reversible preparation steps |
| `PROTECT` | follow applicable authenticated instructions and protective plans |
| `WAIT` | delay is currently less harmful; an evidence deadline and action trigger are required |
| `URGENT_GUIDANCE` | immediate safety-first information; not emergency authority |
| `BLOCK` | request or data is prohibited |
| `DEGRADED` | a dependency or integrity condition prevents normal output |

## 14. Reading warning levels

| Level | Meaning | Personal posture |
|---|---|---|
| `W0 — NO ACTIVE VERIFIED SIGNAL` | no relevant active verified signal found in supplied synthetic data | ordinary preparedness; absence is not proof of safety |
| `W1 — MONITOR` | an attributable preliminary signal warrants attention | monitor and review plans |
| `W2 — PREPARE` | plausible relevant impact supports reversible preparation | prepare now and monitor |
| `W3 — PROTECT` | authenticated relevant warning or strong evidence supports protection | follow applicable instruction and plan |
| `W4 — URGENT ACTION` | simulated immediate/direct danger has high delay cost | act now in the exercise; do not wait for AI certainty |
| `WX — VERIFY / CONFLICT / DEGRADED` | integrity, area, freshness, authority, or evidence is unresolved | verify before consequential action; use justified reversible precautions |

BootX warning level, official status, evidence tier, and AI DNA runtime checks are separate fields. None is an overall probability that an event will occur.

## 15. Reading the decision packet

The TUI displays:

- `Request` — local request identifier;
- `Class` and `Mode` — authority/risk class and response behavior;
- `Goal` — the interpreted user goal;
- `Notice` — prototype and assurance limitation;
- `Warning Card` — present only for a warning exercise;
- `Observations` — deterministic indicators, not verdicts;
- `Unknowns` — facts that remain unresolved;
- `Options` — permissible choices, benefits, and risks;
- `Advisory Result` — recommendation status and basis;
- `Next safe step` — a specific human-controlled step;
- `AI DNA runtime checks` — dimension-level conformance observations, not certification;
- `Blocked external actions` — capabilities technically absent from the prototype;
- `Limitations` — what was not verified, decided, or performed;
- `Data receipt` — memory, remote processing, synthetic status, retention, and location use.

For contained local documents, the packet also shows an evidence receipt containing exact selected-byte integrity and an explicit origin-authentication status.

After viewing the packet, the UI can print structured JSON. Selecting this option does not save a file by itself, but terminal capture or redirection can persist it.

## 16. AI DNA runtime checks

The output reports nine dimensions:

- Truth
- Reasoning
- Learning
- Communication
- Adaptability
- Ethics
- Safety
- Humility
- Common good

Statuses are:

- `PASS` — the local runtime condition passed its implemented rule;
- `CONDITIONAL` — usable only with stated limitations or verification;
- `FAIL` — the affected generated recommendation must be blocked;
- `NOT ESTABLISHED` — evidence is insufficient.

These checks are produced by the prototype itself and are not independent assurance. They must not be interpreted as a forecast probability, accuracy score, moral rating, or safety certification.

## 17. Strict JSON backend mode

Change to the real module:

```powershell
cd D:\Job\Human+AI+DNA\bootx\prototype\personal-companion
```

Run a fixture:

```powershell
go run ./cmd/bootx-companion -input .\testdata\safe-study.json
go run ./cmd/bootx-companion -input .\testdata\suspicious-message.json
go run ./cmd/bootx-companion -input .\testdata\warning-prepare.json
```

Use a built executable from `prototype`:

```powershell
.\personal-companion\dist\bootx-companion-windows-amd64.exe `
  -input .\personal-companion\testdata\suspicious-message.json
```

Compact JSON:

```powershell
go run ./cmd/bootx-companion -input .\testdata\safe-study.json -compact
```

Read one JSON object from standard input:

```powershell
Get-Content -Raw .\testdata\safe-study.json |
  go run ./cmd/bootx-companion -input -
```

Show version:

```powershell
go run ./cmd/bootx-companion -version
```

CLI flags:

| Flag | Meaning |
|---|---|
| `-input <path>` | read one strict request object from a file |
| `-input -` | read one strict request object from standard input |
| `-compact` | emit compact JSON in backend mode |
| `-version` | print application version and exit |
| `-workspace <path>` | workspace boundary for read-only document mode |
| `-document <relative-path>` | contained `.md`, `.txt`, or `.json` file to read |
| `-document-public` | required confirmation that the document is public and non-sensitive |
| `-goal <text>` | decision goal for document mode |
| `-question <text>` | direct question for document mode |
| `-priorities <csv>` | optional comma-separated priorities for document mode |

Unknown fields, unsupported enum values, non-denied remote permission, multiple JSON objects, oversized selected content, non-synthetic sensitive data, and non-synthetic warning input fail closed.

## 18. JSON input contract

Primary schema: [`schemas/personal-decision-input.schema.json`](schemas/personal-decision-input.schema.json)

| Field | Requirement |
|---|---|
| `request_id` | non-empty local identifier |
| `capability_id` | exactly `assist.personal-decision.v1` |
| `user_id` | locally declared identifier; this prototype does not authenticate it |
| `created_at` | RFC 3339 date-time |
| `goal` | non-empty decision purpose |
| `question` | non-empty direct question |
| `selected_content` | maximum 65,536 bytes at runtime |
| `content_source` | typed source; optional contained-file reference, byte length, modification time, and SHA-256 integrity fields; origin claims are not authenticated |
| `data_class` | `public`, `personal`, `sensitive`, or `prohibited` |
| `declared_domain` | one supported domain |
| `memory_permission` | `none` or `session` |
| `remote_permission` | exactly `deny` |
| `output_preference` | supported presentation preference |
| `synthetic` | required boolean; must be `true` for sensitive or warning scenarios |
| `user_priorities` | optional array of distinct short values |
| `warning` | optional typed warning extension; requires `synthetic: true` |

Minimal low-stakes request:

```json
{
  "request_id": "local-example-1",
  "capability_id": "assist.personal-decision.v1",
  "user_id": "owner-local",
  "created_at": "2026-07-12T17:00:00+07:00",
  "goal": "Choose the next reversible study task",
  "question": "Which missing fact should I verify first?",
  "selected_content": "Two study tasks have incomplete evidence.",
  "content_source": {
    "type": "user_selected_content",
    "origin_verified": false
  },
  "data_class": "personal",
  "declared_domain": "study",
  "memory_permission": "session",
  "remote_permission": "deny",
  "output_preference": "standard",
  "synthetic": false,
  "user_priorities": ["truth", "reversibility", "learning"]
}
```

The warning extension is fully specified in the schema and demonstrated by [`testdata/warning-prepare.json`](testdata/warning-prepare.json).

## 19. JSON output contract

Primary schema: [`schemas/personal-decision-output.schema.json`](schemas/personal-decision-output.schema.json)

Important fields:

| Field | Meaning |
|---|---|
| `request_id` | correlates output to input |
| `generated_at` | UTC packet-generation time |
| `runtime_notice` | prototype limitations |
| `decision_class` | `D0`–`D5` authority/risk class |
| `response_mode` | output posture |
| `observations` | sourced deterministic findings |
| `assumptions` | declared contextual assumptions |
| `unknowns` | unresolved information |
| `options` | permissible option records |
| `recommendation` | advisory status, option, and basis |
| `next_safe_step` | human-controlled next action |
| `blocked_actions` | capabilities not available |
| `limitations` | explicit non-claims |
| `warning` | optional `W0`–`W4`/`WX` card |
| `ai_dna_runtime_checks` | nine dimension records |
| `evidence_receipt` | source type, relative reference, integrity result, hash, bytes, observed modification time, and explicit origin-authentication status |
| `data_receipt` | memory, remote, synthetic, retention, and location record |
| `user_decision` | `null` in this MVP; BootX does not choose for the user |

## 20. Synthetic fixtures and expected outputs

| Fixture | Purpose | Expected result |
|---|---|---|
| `safe-study.json` | low-stakes non-sensitive study decision | `D0 / INFORM` |
| `suspicious-message.json` | synthetic credential/link/urgency indicators | `D2 / VERIFY` |
| `warning-prepare.json` | synthetic possible relevant hazard | `D2 / PREPARE / W2` |
| `warning-urgent.json` | synthetic authenticated immediate warning | `D4 / URGENT_GUIDANCE / W4` |
| `typhoon-bavi-exercise.json` | full synthetic Typhoon Bavi preparedness case | `D2 / PREPARE / W2` |
| `space-governance-readiness-exercise.json` | fictional settlement request requiring qualified governance review | `D3 / ABSTAIN` |

Run every fixture and all automated checks:

```powershell
.\build.ps1 -Action verify
```

Passing fixtures prove only that implemented rules behave as expected for those cases. They do not establish real-world accuracy or safety.

The dedicated synthetic Typhoon Bavi end-to-end case is documented at [`../TEST_CASE.md`](../TEST_CASE.md). Run it from `prototype` with:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass `
  -File .\test-case-typhoon-bavi.ps1
```

It records the complete synthetic input, observable deterministic processing evidence, assertions, and complete JSON output in `prototype\typhoon-bavi.log`.

The space-governance boundary case is documented at [`../TEST_CASE_SPACE_GOVERNANCE.md`](../TEST_CASE_SPACE_GOVERNANCE.md). It verifies that the MVP abstains from certifying a fictional lunar settlement and that the research chapter preserves mandatory evidence boundaries.

## 21. Troubleshooting

### `go.mod file not found`

You are probably in `prototype` instead of `prototype\personal-companion`.

Use the parent build script:

```powershell
.\build.ps1 -Action verify
```

Or enter the real module before direct Go commands:

```powershell
cd .\personal-companion
go test ./...
```

Do not run `go mod init` in `prototype`.

### `warning: "all" matched no packages`

This normally means an empty parent module was initialized by mistake. Remove only the accidental `prototype\go.mod`, then use `prototype\personal-companion\go.mod`.

```powershell
Remove-Item D:\Job\Human+AI+DNA\bootx\prototype\go.mod
```

Before deleting, confirm that the file contains the accidental empty module and that `personal-companion\go.mod` still exists.

### PowerShell says script execution is disabled

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\build.ps1 -Action verify
```

This bypass applies only to that new PowerShell process.

### Go cache reports `Access is denied`

Use `build.ps1`. It places `GOCACHE` and `GOTMPDIR` under the current user's temporary directory for the build process.

### `real sensitive input is not authorized`

The security gate is working. Remove real sensitive data. Use a minimized non-sensitive scenario or an explicitly synthetic test fixture.

### `warning assessment is synthetic-only`

The warning workflow is not validated for real use. Set `synthetic: true` only for an actual test exercise; do not relabel real warning reliance as synthetic.

### `unknown field` in JSON

The backend uses strict decoding. Correct the field name or update the reviewed contract and code together. Do not silently discard unknown fields.

### `request must contain exactly one JSON object`

The input contains trailing data or multiple objects. Submit one object per process invocation.

### The output does not sound like a generative AI

That is expected. Version `0.2.0-dev` is a deterministic baseline with contained read-only local-document ingestion. A bounded model is intentionally disconnected until the baseline is frozen, measured, and reviewed.

### Build succeeds but there is no `dist` directory where expected

Check whether `-OutputDirectory` was supplied. The build summary prints the absolute binary and manifest paths.

## 22. Safe cleanup

Remove only the default generated build output:

```powershell
.\build.ps1 -Action clean
```

The script refuses to clean a custom output path.

The application creates no installed service, registry entry, scheduled task, persistent memory database, or background process. To stop the TUI, choose `q` or press `Ctrl+C`. To remove the source, use normal repository/version-control procedures only after preserving any work you intend to keep.

## 23. Operator checklist

Before processing:

- [ ] I am using the correct `personal-companion` module or parent build script.
- [ ] I have not entered credentials or unnecessary sensitive information.
- [ ] Any sensitive or warning scenario is genuinely synthetic.
- [ ] My goal and direct question are narrow and understandable.
- [ ] I reviewed the visible data scope before confirmation.

Before acting on output:

- [ ] I distinguished observations from unknowns and recommendations.
- [ ] I understand the `D` class, response mode, and any `W` level.
- [ ] I did not interpret AI DNA runtime checks as certification or probability.
- [ ] I verified consequential facts through an appropriate independent source.
- [ ] I understand that BootX did not execute an external action.
- [ ] The final decision remains mine.

## 24. Governing documentation

- [Personal Decision-Assistance Pipeline](../../docs/handbook/14-personal-decision-pipeline.md)
- [Development Guideline](../../DEVELOPMENT_GUIDELINE.md)
- [AI DNA Operational Specification](../../docs/handbook/05-ai-dna-specification.md)
- [Companion System Architecture](../../docs/handbook/06-companion-system-architecture.md)
- [Safety Case and Risk Register](../../docs/handbook/09-safety-case-and-risk-register.md)
- [Project Progress and Readiness](../../PROGRESS.md)

Where this usage guide conflicts with a higher-authority project document, follow the higher-authority requirement and record the inconsistency for correction.
