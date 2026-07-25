# BootX Personal Companion MVP

**Status:** `DEV-1` software-only development prototype; not a deployed AI companion, emergency service, or safety-certified system<br>
**Version:** `0.4.0-dev`<br>
**Runtime:** Go 1.22 or later on a mature supported host operating system<br>
**Capabilities:** `assist.personal-decision.v1`, `assist.law-clarity.v1`, `assist.ethical-review.v1`

This module implements the first BootX `Input → Process → Output` decision-assistance pipeline as a dependency-free Go backend and terminal UI.

This `0.4.0-dev` module is the final planned publicly source-available development baseline. It remains unvalidated and may receive bounded maintenance, but new risk-management and decision-assistance development is intended for the controlled private successor described in [RELEASE_BOUNDARY.md](../../RELEASE_BOUNDARY.md). Existing public-baseline license rights remain intact.

For complete operator instructions, field definitions, examples, build verification, and troubleshooting, read [USAGE.md](USAGE.md).

## Safety boundary

- deterministic personal, warning, Law Clarity, and declared-evidence rules remain authoritative;
- OpenAI is connected only for the separate `assist.ethical-review.v1` public-draft advisory mode after explicit remote-processing and human-authority confirmation;
- one explicitly confirmed public, non-sensitive `.md`, `.txt`, or `.json` document can be read from inside a user-selected workspace with a SHA-256 integrity receipt;
- public, non-sensitive legal text may be screened from reviewer-supplied scores and rationales; the result is not legal advice, validity, constitutionality, guilt, liability, or enforcement authority;
- network access is limited to a fixed OpenAI Responses API endpoint in ethical-review mode; no browser, message, call, payment, account, device, robot, siren, or broadcast capability exists;
- personal-decision and Law Clarity remote processing remains denied;
- sensitive and forecast/disaster scenarios must be explicitly synthetic in this unvalidated MVP;
- selected content is untrusted data and cannot change policy;
- the application does not write raw content; local-mode content remains process-local, while ethical-review draft content is explicitly sent to OpenAI and disclosed in its remote receipt;
- health and legal questions defer to qualified professionals;
- disaster output preserves official-status distinctions and is not an official alert;
- AI DNA runtime checks are not certification or forecast probability;
- file integrity does not authenticate authorship or establish that the document's claims are true;
- `user_id`, `origin_verified`, and warning-authority fields are declarations, not authenticated identities;
- the human user remains the decision-maker.
- no OpenAI output can lower a deterministic warning, approve publication, determine guilt or sentence, or trigger an external action.

Shell redirection, terminal history, screen capture, crash dumps, or the surrounding operating system can still persist output. Do not enter passwords, PINs, recovery codes, private keys, or unnecessary personal data.

## Run the terminal UI

From this directory:

```powershell
go run ./cmd/bootx-companion
```

The UI supports:

1. personal decision assistance;
2. forecast/disaster warning assessment using manually selected information;
3. contained read-only processing of one confirmed public workspace document;
4. Law Clarity Logic deterministic screening;
5. a visible safety-boundary screen.

The TUI is deliberately line-oriented so it remains inspectable, keyboard-accessible, and dependency-free. It is not a background agent.

## Backend JSON mode

```powershell
go run ./cmd/bootx-companion -input ./testdata/suspicious-message.json
go run ./cmd/bootx-companion -input ./testdata/warning-prepare.json
go run ./cmd/bootx-companion -law-input ./testdata/law-clarity-gray-zone.json
go run ./cmd/bootx-companion -review-input ./testdata/ethical-review-synthetic-publication.json
```

Use `-input -` to read exactly one strict JSON request from standard input. Unknown JSON fields, unsupported permissions, malformed warning values, and multiple JSON objects fail closed.

Schemas:

- [`schemas/personal-decision-input.schema.json`](schemas/personal-decision-input.schema.json)
- [`schemas/personal-decision-output.schema.json`](schemas/personal-decision-output.schema.json)
- [`schemas/law-clarity-input.schema.json`](schemas/law-clarity-input.schema.json)
- [`schemas/law-clarity-output.schema.json`](schemas/law-clarity-output.schema.json)
- [`schemas/ethical-review-input.schema.json`](schemas/ethical-review-input.schema.json)
- [`schemas/ethical-review-output.schema.json`](schemas/ethical-review-output.schema.json)

The complete formulas, gates, workflow, safeguards, worked example, limitations, and inheritance requirements are recorded in the canonical [Law Clarity Logic reference](../../docs/handbook/15-law-clarity-logic.md).

## OpenAI ethical-review mode

This opt-in CLI capability reviews a public, non-sensitive social post, speech, statement, proposal, or decision rationale before a separate human decision:

```powershell
-not [string]::IsNullOrWhiteSpace($env:OPENAI_API_KEY)
go run ./cmd/bootx-companion `
  -review-input ./testdata/ethical-review-synthetic-publication.json
```

The key is read only from `OPENAI_API_KEY`; never put it in a command argument, JSON fixture, source file, or repository. The request must explicitly confirm public/non-sensitive data, OpenAI remote processing, and continuing human authority.

BootX first calculates transparent declared-evidence metrics and a review-priority index. OpenAI then returns a strict structured critique with no tools, conversation state, application persistence, or external actions. The API request uses `store:false`. The raw public draft is sent to OpenAI, and the output contains a visible receipt. Neither the local index nor the model output is a truth finding, harm probability, moral certificate, legal opinion, publication approval, guilt finding, or sentence recommendation.

Read the complete [Ethical Publication and Decision-Rationale Review](../../docs/handbook/16-ethical-publication-review.md) contract before use.

## Read-only real workspace mode

From the repository root, process one deliberately reviewed public document without copying it into a synthetic fixture:

```powershell
.\prototype\personal-companion\dist\bootx-companion-windows-amd64.exe `
  -workspace . `
  -document docs\research\civilization\religion-ideology-and-decision-integrity.md `
  -document-public `
  -goal "Choose the next evidence-improvement task" `
  -question "Which missing fact should be verified first?" `
  -priorities "truth, reversibility, common good"
```

The document path must be relative and remain inside the workspace after link resolution. Only regular UTF-8 `.md`, `.txt`, and `.json` files up to 65,536 bytes are accepted. The result records the relative reference, byte length, modification time, and SHA-256. It also states `origin_status: not_authenticated` because a hash identifies bytes, not an author or truthful claim.

A deterministic scan reports heading count, external-link strings, open Markdown checklist items, lines containing configured evidence-gap markers, and the first review candidate. These are structural keyword observations, not semantic understanding or proof that a candidate is truly the highest priority.

## Test and build

From the repository's `prototype` directory, use the supported build script:

```powershell
.\build.ps1                 # test, vet, JSON validation, and build
.\build.ps1 -Action verify  # also run fixtures, real-document, and Law Clarity smoke tests
.\build.ps1 -Action test
.\build.ps1 -Action run     # build and start the TUI
.\build.ps1 -Action clean
```

The default output is `personal-companion\dist\` and includes the executable, SHA-256 file, and `build-manifest.json`. The script uses `$PSScriptRoot`, so it finds the correct nested Go module even when launched from another directory.

If Windows reports that script execution is disabled, use a process-scoped bypass that does not change the machine policy:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\build.ps1 -Action verify
```

Direct Go commands remain available from `prototype\personal-companion`:

```powershell
go test ./...
go build ./cmd/bootx-companion
```

Synthetic fixtures are under [`testdata/`](testdata/). They are exercises, not real events or operational evidence.

The complete synthetic Typhoon Bavi end-to-end regression, runner, and captured evidence are documented in [`../TEST_CASE.md`](../TEST_CASE.md).

The synthetic space-governance abstention test, documentation assertions, runner, and captured evidence are documented in [`../TEST_CASE_SPACE_GOVERNANCE.md`](../TEST_CASE_SPACE_GOVERNANCE.md).

The paired synthetic Law Clarity capital-corruption test checks an abusive scapegoating proposal against a non-capital, evidence-based revision. Its complete specification, runner, and captured evidence are documented in [`../TEST_CASE_LAW_CLARITY_CAPITAL_CORRUPTION.md`](../TEST_CASE_LAW_CLARITY_CAPITAL_CORRUPTION.md).

The synthetic OpenAI ethical-review test verifies the live fixed-endpoint connection and no-storage/no-tools/no-actions/human-authority boundaries without retaining the API key or full generated prose. See [`../TEST_CASE_OPENAI_ETHICAL_REVIEW.md`](../TEST_CASE_OPENAI_ETHICAL_REVIEW.md).

## Package structure

```text
cmd/bootx-companion/   CLI and strict JSON entry point
internal/engine/       input validation and decision-packet composition
internal/ethicalreview/ deterministic declared-evidence formulas, gates, and advisory types
internal/evidence/     contained read-only local-file loading and integrity receipts
internal/lawclarity/   deterministic legal-clarity formulas, gates, review prompts, and tests
internal/model/        typed request, warning, option, assurance, and output models
internal/openaiadvisory/ bounded Responses API client, prompt, schema, receipts, and tests
internal/policy/       embedded deterministic policy and indicator rules
internal/tui/          interactive terminal workflow and rendering
internal/warning/      deterministic W0-W4/WX warning evaluation
schemas/               machine-readable input and output contracts
testdata/              synthetic personal, scam, disaster, governance, and legal-clarity cases
```

## Current limitations

- rule indicators are an auditable baseline, not a validated fraud detector;
- no authoritative source retrieval or cryptographic alert-authority registry exists;
- timestamps and URLs are preserved but not independently fetched or authenticated;
- local document authorship, publisher identity, and factual claims are not authenticated;
- Law Clarity ratings, thresholds, weights, bands, phrase scan, and risk index are unvalidated research aids, not legal findings or probabilities;
- OpenAI review can be incomplete, biased, mistaken, over-cautious, or falsely reassuring; no representative fairness or benefit study exists;
- no user-login or operating-system identity authentication exists;
- no persistent encrypted memory exists;
- no multilingual or completed accessibility study exists;
- no calibrated benefit/harm percentages or representative field results exist;
- no family, community, operating-system, or robotics integration is authorized.

See the [Personal Decision-Assistance Pipeline](../../docs/handbook/14-personal-decision-pipeline.md) and [Development Guideline](../../DEVELOPMENT_GUIDELINE.md) for requirements and gates.
