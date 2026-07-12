# BootX Prototype Test Case — Synthetic Typhoon Bavi Preparedness

**Test ID:** `TC-BAVI-001`<br>
**Test type:** deterministic end-to-end regression and evidence-capture test<br>
**Status:** synthetic exercise only; not a real forecast, warning, or emergency recommendation<br>
**Capability:** `assist.personal-decision.v1`<br>
**Runner:** [`test-case-typhoon-bavi.ps1`](test-case-typhoon-bavi.ps1)<br>
**Input fixture:** [`personal-companion/testdata/typhoon-bavi-exercise.json`](personal-companion/testdata/typhoon-bavi-exercise.json)<br>
**Evidence log:** [`typhoon-bavi.log`](typhoon-bavi.log)

For a short non-technical interpretation, see [`TYPHOON_BAVI_SUMMARY.md`](TYPHOON_BAVI_SUMMARY.md).

## 1. Purpose

Verify that the BootX Personal Companion MVP converts a synthetic Typhoon Bavi scenario into a conservative personal preparedness decision packet without claiming official authority, executing an external action, using remote processing, or escalating a preliminary scenario into an urgent warning.

The expected result is:

```text
Decision class:   D2
Response mode:    PREPARE
Warning level:    W2 — PREPARE
Decision posture: PREPARE NOW
Human authority:  Joni decides
```

## 2. Non-purpose

This test does not:

- represent an actual Typhoon Bavi forecast or event update;
- authenticate a meteorological or emergency authority;
- predict landfall, track, wind, rainfall, flood, landslide, damage, or casualties;
- authorize evacuation, travel, public warning, emergency contact, or family broadcast;
- measure real-world warning accuracy, benefit, harm, or AI DNA validity;
- connect a generative AI model or network service.

The name “Typhoon Bavi” is a synthetic scenario label for exercising BootX logic.

## 3. Preconditions

- Go 1.22 or later is available on `PATH`.
- PowerShell can launch the runner with a process-scoped execution-policy bypass if required.
- The real Go module exists at `prototype/personal-companion/go.mod`.
- The repository contains no required external Go modules.
- The test input contains only synthetic data.
- No real emergency reliance is permitted.

## 4. Input scenario

The fixture declares:

- two independent mock forecast lineages;
- possible typhoon-related wind and rainfall within 48 hours;
- possible cascading flood and landslide impacts;
- no official local warning or evacuation instruction in the exercise;
- a configured personal area near the modeled impact area;
- evidence tier `V2`;
- future urgency, severe potential impact, and possible certainty;
- low-burden reversible preparation as the appropriate candidate posture;
- remote processing denied and synthetic status true.

The test intentionally uses `official_status: none_found`. BootX must distinguish that from an authenticated official all-clear and must not say the user is safe.

## 5. Observable processing stages

The runner records process evidence without inventing hidden model reasoning:

1. **Build integrity:** formatting, tests, vet, JSON validation, binary build, version, and SHA-256.
2. **Input validation:** capability, synthetic status, permissions, domain, data class, warning fields, and content size.
3. **Policy boundary:** remote processing denied; external actions absent; warning scenario allowed only because it is synthetic.
4. **Indicator analysis:** deterministic observations from selected content.
5. **Decision classification:** decision class and response mode.
6. **Warning evaluation:** official status, area match, urgency, severity, certainty, evidence tier, warning level, and posture.
7. **Option construction:** permitted options and reversibility.
8. **AI DNA runtime checks:** all nine dimension statuses and bases.
9. **Data receipt:** no remote processing, no persistent memory, synthetic status, and location-use statement.
10. **Human authority:** `user_decision` remains `null`; no external action occurs.

The process section is an auditable summary of typed inputs, deterministic program outputs, test gates, and safety controls. It is not private chain-of-thought.

## 6. Required assertions

| ID | Assertion |
|---|---|
| BAVI-A01 | Build manifest reports `tests_skipped: false`. |
| BAVI-A02 | Build manifest reports `external_modules: 0`. |
| BAVI-A03 | Application capability is `assist.personal-decision.v1`. |
| BAVI-A04 | Decision class is `D2`. |
| BAVI-A05 | Response mode is `PREPARE`. |
| BAVI-A06 | Warning level is `W2`. |
| BAVI-A07 | Warning label is `PREPARE`. |
| BAVI-A08 | Decision posture is `PREPARE NOW`. |
| BAVI-A09 | Evidence tier is `V2`. |
| BAVI-A10 | Official status remains `none_found`. |
| BAVI-A11 | Recommendation selects `prepare-now`. |
| BAVI-A12 | Nine AI DNA runtime-check records are present. |
| BAVI-A13 | Remote processing is false. |
| BAVI-A14 | Persistent memory use is false. |
| BAVI-A15 | Data receipt records `synthetic: true`. |
| BAVI-A16 | Family/public broadcast is in `blocked_actions`. |
| BAVI-A17 | Device/robot control is in `blocked_actions`. |
| BAVI-A18 | User decision remains `null`. |
| BAVI-A19 | Output includes the limitation that BootX is not an alerting or emergency authority. |
| BAVI-A20 | Backend exits successfully and the evidence log finishes with `RESULT: PASS`. |

Any failed assertion makes the test fail. No averaging is permitted.

## 7. Run the test

From the repository root:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass `
  -File .\prototype\test-case-typhoon-bavi.ps1
```

From the `prototype` directory:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass `
  -File .\test-case-typhoon-bavi.ps1
```

The script overwrites only its declared log file, creates a unique temporary build directory, and removes that temporary directory after the test.

Use a different log path when needed:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass `
  -File .\test-case-typhoon-bavi.ps1 `
  -LogPath D:\Temp\typhoon-bavi-test.log
```

## 8. Log structure

`typhoon-bavi.log` contains:

```text
TEST METADATA
BUILD AND VERIFICATION
INPUT — COMPLETE JSON REQUEST
PROCESS — OBSERVABLE DETERMINISTIC EVIDENCE
ASSERTIONS
OUTPUT — COMPLETE JSON DECISION PACKET
RESULT
CLEANUP
```

The input and output are synthetic but complete. Do not adapt the runner to log real credentials, personal records, precise private location, or operational emergency data.

## 9. Expected protective interpretation

The expected `W2 — PREPARE` result means:

- take low-cost reversible preparation steps;
- maintain access to independently obtained official updates;
- review communication, supplies, documents, shelter, and routes;
- do not interpret forecast boundaries as precise impact boundaries;
- do not evacuate, travel into risk, or broadcast a public warning based on this synthetic exercise;
- escalate only when authenticated, relevant, current evidence satisfies deterministic rules.

“Prepare” is not “panic,” “evacuate,” “safe,” or “certain.”

## 10. Pass limitations

A passing result establishes only that the current source and test fixture produce the specified deterministic packet in this host environment. It does not establish:

- correct behavior for every cyclone or hazard;
- calibrated false-positive or false-negative rates;
- usability under stress, disability, language difference, outage, or device loss;
- official source authentication;
- representative human comprehension;
- real personal, family, or community benefit;
- readiness for bounded AI, operating-system integration, robotics, or public deployment.

## 11. Related documents

- [Complete MVP Usage Guide](personal-companion/USAGE.md)
- [Development Guideline](../DEVELOPMENT_GUIDELINE.md)
- [Personal Decision-Assistance Pipeline](../docs/handbook/14-personal-decision-pipeline.md)
- [Anti-Abuse and Evidence Integrity](../docs/handbook/12-anti-abuse-and-evidence-integrity.md)
- [Safety Case and Risk Register](../docs/handbook/09-safety-case-and-risk-register.md)
