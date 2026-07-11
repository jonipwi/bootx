# Human + AI Companion System Architecture

## 1. Architectural objective

Build a local-first, hybrid companion that can explain risk, protect user-approved workflows, manage consent-based memory, and obtain bounded external intelligence without making a remote model or a local agent an unaccountable authority.

The architecture assumes compromise and error are possible. Trust is produced through isolation, least privilege, provenance, verification, monitoring, recovery, and accountable human control.

## 2. Reference architecture

```text
Human / Authorized Guardian / Qualified Reviewer
                     │
       Consent, confirmation, appeal, stop
                     │
┌────────────────────▼─────────────────────┐
│ Experience and Explanation Layer         │
│ accessible UI • evidence • uncertainty   │
│ options • memory ledger • audit viewer   │
└────────────────────┬─────────────────────┘
                     │
┌────────────────────▼─────────────────────┐
│ Companion Orchestrator                   │
│ intent • context • plan • response       │
│ no direct ambient authority              │
└──────────┬──────────────┬────────────────┘
           │              │
┌──────────▼─────────┐ ┌──▼────────────────┐
│ Policy & Safety    │ │ Evidence Service   │
│ risk tier • rules  │ │ provenance • time  │
│ consent • gates    │ │ retrieval • claims │
└──────────┬─────────┘ └──┬────────────────┘
           │              │
┌──────────▼──────────────▼────────────────┐
│ Capability Broker                         │
│ typed tools • least privilege • budgets   │
│ confirmation • sandbox • rate limits      │
└───────┬───────────────┬──────────────────┘
        │               │
┌───────▼────────┐ ┌────▼──────────────────┐
│ Local services │ │ Remote bounded services│
│ rules/models   │ │ reasoning/threat feeds │
│ encrypted data │ │ minimization/redaction │
└───────┬────────┘ └────┬──────────────────┘
        │               │
┌───────▼───────────────▼──────────────────┐
│ Secure platform foundation               │
│ boot integrity • isolation • keys         │
│ signed updates • logging • recovery       │
└───────────────────────────────────────────┘
```

Models propose. Policy constrains. The capability broker mediates. Humans authorize consequential action. The platform records and can recover.

## 3. Core components

### Experience and explanation layer

Provides calm, accessible explanations with observed facts, inference, confidence, options, and recourse. It must support screen readers, keyboard navigation, multiple reading levels, and low-bandwidth/degraded modes. It must not use anthropomorphic or emotional pressure to increase compliance.

### Companion orchestrator

Coordinates tasks but holds no unrestricted credentials. It receives only the context required for the present purpose, produces a typed plan, and requests capabilities through the broker. Free-form generated text cannot itself grant permission.

### Policy and safety engine

Applies deterministic controls around probabilistic models:

- capability allowlists and risk tiers;
- consent and confirmation state;
- user and guardian boundaries;
- data classification and destination policy;
- emergency and qualified-review rules;
- rate, money, time, and scope limits;
- hard prohibitions and safe fallback.

Policy changes require versioning, review, tests, and rollback.

### Evidence service

Maintains source provenance, timestamps, integrity, claim relationships, conflict, and uncertainty. It prevents a generated answer from laundering unknown-origin data into apparent fact.

### Capability broker

The broker exposes narrow, typed operations such as `analyze_selected_message`, `lookup_official_contact`, or `draft_warning`. Each operation declares permissions, sensitive inputs, external effects, confirmation requirements, and audit fields. Tokens are short-lived and scoped to one purpose.

### Memory service

Memory is divided by purpose:

| Memory class | Example | Default |
|---|---|---|
| Session | current task context | delete after session or short expiry |
| Preference | reading level, language | local, user-editable |
| Personal continuity | user-approved goals or history | explicit opt-in, encrypted, item-level deletion |
| Safety event | confirmed scam indicator or incident | minimal record, defined retention |
| Audit | permission and external action record | tamper-evident, access-controlled |

Each item records source, creator, date, confidence, sensitivity, purpose, access, expiry, and correction history. Retrieval must be purpose-filtered. “Remember everything” is prohibited.

### Remote service gateway

Before remote processing, minimize, redact, or pseudonymize inputs; show the user when consequential personal data leaves the device; enforce purpose-bound contracts; authenticate responses; and handle network failure without unsafe guessing.

### Secure platform foundation

A future BootX platform should provide verified boot, signed components, memory/process isolation, hardware-backed keys where available, encrypted storage, secure time, signed updates, recovery partitions, tamper-evident logging, and a tested factory-reset/export path.

The current repository does not contain auditable source for these controls. They are requirements, not descriptions of the supplied binary.

## 4. Trust boundaries and threat actors

Assume threats from:

- external scammers and malware;
- malicious content causing prompt or tool injection;
- compromised remote models, feeds, or updates;
- abusive household members or guardians;
- insiders and administrators;
- commercial pressure to expand data collection;
- user mistakes under urgency or distress;
- model hallucination, bias, or correlated failure;
- loss or theft of the device;
- the companion's own excessive authority.

Content is data, not instruction. A message being analyzed cannot grant tools, disclose memory, or change policy. Remote text and model outputs must be treated as untrusted inputs.

## 5. Permission model

| Action class | Example | Required control |
|---|---|---|
| Read local selected data | analyze one selected message | contextual consent and visible scope |
| Read sensitive memory | retrieve health or financial context | explicit purpose, strong authentication, audit |
| Contact external service | verify an official domain | destination allowlist, data minimization |
| Draft external effect | prepare a reply | user reviews before send |
| Execute consequential effect | send money, change account, call emergency service | outside initial scope or qualified multi-step authorization |
| Change policy/update | install model or rule set | signature, compatibility tests, staged rollout, rollback |

Permissions expire. Denial must not be converted into repeated pressure.

## 6. Protective workflow example

For a suspected bank message:

1. User selects the message; the UI displays exactly what will be analyzed.
2. Local rules extract observable indicators without opening active links.
3. Evidence service compares the sender and link with independently obtained official data.
4. A model explains indicators and uncertainty; it does not declare guilt.
5. Policy engine blocks any direct transfer or credential submission through the analysis flow.
6. User receives choices: dismiss, inspect details, call an official number, or contact a preapproved trusted person.
7. Any external lookup is minimized and logged.
8. User can report a false alert and view/delete the event record.

## 7. Resilience modes

| Condition | Required behavior |
|---|---|
| Remote AI unavailable | local rules and clear degraded-mode label; no fabricated answer |
| Evidence stale | display date and require fresh verification for consequential use |
| Integrity check fails | disable affected capability, preserve safe read-only access, alert operator |
| Memory unavailable | do not reconstruct personal history from guesses |
| Update failure | return to last known-good signed version |
| Suspected compromise | revoke tokens, isolate service, preserve evidence, guide recovery |

## 8. Observability without surveillance

Collect the minimum telemetry needed to evaluate reliability and safety. Prefer local aggregation and privacy-preserving counts. Never log raw sensitive content by default. Operational dashboards should include errors, alert burden, overrides, safety blocks, integrity failures, rollback, deletion success, subgroup performance where ethical and valid, and unresolved incidents.

## 9. Portability and exit

The user must be able to:

- see stored memory and its purpose;
- correct, export, and delete it in understandable formats;
- revoke remote integrations and credentials;
- run essential protective functions in a bounded offline mode where feasible;
- replace the companion provider without losing control of personal data;
- disable companionship features while retaining access to their own files.

Exit is part of safety architecture, not a commercial option.

## 10. BootX platform path

The practical engineering order is:

1. restore and reproduce the existing educational OS build;
2. establish memory safety, testing, image provenance, and recovery;
3. build companion services on a mature host platform first to validate human outcomes;
4. define a narrow IPC/API contract between platform and companion;
5. port only validated minimal services to BootX;
6. do not place personal users or real credentials on BootX until security assurance is credible.

This order preserves the educational OS vision while avoiding use of an experimental kernel as the security base for sensitive AI companionship.

