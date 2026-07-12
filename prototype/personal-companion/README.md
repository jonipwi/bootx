# BootX Personal Companion MVP

**Status:** `DEV-1` software-only development prototype; not a deployed AI companion, emergency service, or safety-certified system<br>
**Version:** `0.1.0-dev`<br>
**Runtime:** Go 1.22 or later on a mature supported host operating system<br>
**Capability:** `assist.personal-decision.v1`

This module implements the first BootX `Input → Process → Output` decision-assistance pipeline as a dependency-free Go backend and terminal UI.

For complete operator instructions, field definitions, examples, build verification, and troubleshooting, read [USAGE.md](USAGE.md).

## Safety boundary

- deterministic rules only; no AI model is connected;
- no network, browser, message, call, payment, account, device, robot, siren, or broadcast capability;
- remote processing must be `deny`;
- sensitive and forecast/disaster scenarios must be explicitly synthetic in this unvalidated MVP;
- selected content is untrusted data and cannot change policy;
- raw content remains in process memory only and is not written by the program;
- health and legal questions defer to qualified professionals;
- disaster output preserves official-status distinctions and is not an official alert;
- AI DNA runtime checks are not certification or forecast probability;
- the human user remains the decision-maker.

Shell redirection, terminal history, screen capture, crash dumps, or the surrounding operating system can still persist output. Do not enter passwords, PINs, recovery codes, private keys, or unnecessary personal data.

## Run the terminal UI

From this directory:

```powershell
go run ./cmd/bootx-companion
```

The UI supports:

1. personal decision assistance;
2. forecast/disaster warning assessment using manually selected information;
3. a visible safety-boundary screen.

The TUI is deliberately line-oriented so it remains inspectable, keyboard-accessible, and dependency-free. It is not a background agent.

## Backend JSON mode

```powershell
go run ./cmd/bootx-companion -input ./testdata/suspicious-message.json
go run ./cmd/bootx-companion -input ./testdata/warning-prepare.json
```

Use `-input -` to read exactly one strict JSON request from standard input. Unknown JSON fields, unsupported permissions, malformed warning values, and multiple JSON objects fail closed.

Schemas:

- [`schemas/personal-decision-input.schema.json`](schemas/personal-decision-input.schema.json)
- [`schemas/personal-decision-output.schema.json`](schemas/personal-decision-output.schema.json)

## Test and build

From the repository's `prototype` directory, use the supported build script:

```powershell
.\build.ps1                 # test, vet, JSON validation, and build
.\build.ps1 -Action verify  # also run all backend fixture smoke tests
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

## Package structure

```text
cmd/bootx-companion/   CLI and strict JSON entry point
internal/engine/       input validation and decision-packet composition
internal/model/        typed request, warning, option, assurance, and output models
internal/policy/       embedded deterministic policy and indicator rules
internal/tui/          interactive terminal workflow and rendering
internal/warning/      deterministic W0-W4/WX warning evaluation
schemas/               machine-readable input and output contracts
testdata/              synthetic personal, scam, and disaster cases
```

## Current limitations

- rule indicators are an auditable baseline, not a validated fraud detector;
- no authoritative source retrieval or cryptographic alert-authority registry exists;
- timestamps and URLs are preserved but not independently fetched or authenticated;
- no persistent encrypted memory exists;
- no multilingual or completed accessibility study exists;
- no calibrated benefit/harm percentages or representative field results exist;
- no family, community, operating-system, or robotics integration is authorized.

See the [Personal Decision-Assistance Pipeline](../../docs/handbook/14-personal-decision-pipeline.md) and [Development Guideline](../../DEVELOPMENT_GUIDELINE.md) for requirements and gates.
