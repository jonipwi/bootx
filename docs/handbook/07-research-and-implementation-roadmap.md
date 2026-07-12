# Research and Implementation Roadmap

Implementation status is tracked in the repository-level [PROGRESS.md](../../PROGRESS.md). This roadmap defines the intended sequence; a roadmap item is not complete unless the progress record links to adequate evidence.

## 1. Strategy

Begin with one narrow, high-value, reversible use case: explanation and verification support for suspicious digital messages. This directly reflects the repository's family-protection and cybersecurity vision and permits meaningful testing without granting autonomous control.

Each phase is evidence-gated. A calendar date or feature count does not authorize progression.

## 2. Phase 0 — Research and repository integrity

**Objective:** make the project reproducible and governable.

Deliverables:

- restored BootX source, build scripts, dependencies, ownership records, and third-party license inventory;
- reproducible-build instructions and artifact manifest;
- adopted project Charter, governance, succession, contributor, licensing, ethical-constitution, and decision-log processes;
- stakeholder map including older adults, accessibility representatives, security experts, and caregivers;
- threat model, data map, research protocol templates, and incident process;
- sourced review of anti-scam interventions, fake-evidence and deepfake defenses, human factors, and relevant legal/standards obligations for target jurisdictions.

Exit gate:

- independent reviewer can reproduce artifacts;
- scope, accountable owners, prohibited uses, and stop authority are approved;
- no collection of participant data before ethics/privacy review.

## 3. Phase 1 — Non-AI interaction prototype

**Objective:** test explanation and choice design before adding model variability.

Build:

- static examples of suspicious and legitimate messages;
- clear observed-indicator display;
- low/uncertain/high concern language without false certainty;
- safe verification workflow using independently obtained contact information;
- accessible interface variants and trusted-contact setup.

Research:

- formative interviews and cognitive walkthroughs;
- comprehension, decision time, false-alarm reaction, accessibility, and emotional response;
- whether users learn transferable verification skills.

Exit gate:

- users understand the difference between indicator and verdict;
- critical actions are accessible;
- no dark pattern or coercive language remains;
- predefined minimum comprehension and maximum harmful-action rates are met.

## 4. Phase 2 — Local protective prototype

**Objective:** add deterministic local analysis with no autonomous external action.

**Current evidence:** the [Go personal-companion MVP](../../prototype/personal-companion/README.md) implements a limited deterministic baseline for this phase. It does not close the phase exit gate; representative evaluation, alert-burden study, parser/security review, human deletion/disable study, and rollback evidence remain incomplete.

Build:

- safe parsing of sender, link, urgency, credential, and payment indicators;
- local policy engine and typed analysis record;
- session-only processing by default;
- test corpus with provenance and clearly defined labels;
- audit, deletion, and degraded-mode controls.

Research:

- precision, recall, calibration, per-user alert burden, and subgroup error;
- adversarial links, malformed messages, Unicode deception, and prompt injection;
- security review of parsing and storage.

Exit gate:

- safety and security tests pass approved thresholds;
- false-positive harms and alert fatigue are acceptable in the study context;
- users can delete records and disable the feature successfully;
- rollback is demonstrated.

## 5. Phase 3 — Bounded AI explanation

**Objective:** determine whether a model improves comprehension without increasing unsafe reliance.

Build:

- evidence-grounded explanation generated only from structured indicators;
- uncertainty and limitation template;
- capability broker with no send, payment, credential, or account-change tools;
- versioned model/prompt/policy evaluation;
- local-first operation with a minimized remote option.

Study design:

- compare rule-only explanation with AI-assisted explanation;
- preregister primary outcomes: correct safe action and calibrated comprehension;
- secondary outcomes: time, trust calibration, anxiety, independent skill, accessibility;
- record adverse effects and model disagreement.

Exit gate:

- meaningful improvement over comparator with no material degradation in safety, agency, privacy, or equity;
- claim fabrication and unsupported verdicts below approved limits;
- independent red-team and privacy review complete.

## 6. Phase 4 — Consented personal continuity

**Objective:** test whether limited memory improves protection while preserving autonomy.

Possible memory:

- preferred language and explanation level;
- user-selected official organizations and trusted contacts;
- prior confirmed scam patterns;
- accessibility preferences.

Do not initially store broad conversation history, intimate profiling, credentials, or inferred psychological traits.

Exit gate:

- users can accurately inspect, correct, export, and delete memory;
- purpose and retention comprehension meets threshold;
- memory poisoning, cross-user leakage, and abusive-guardian tests pass;
- benefit over non-memory baseline justifies residual privacy risk.

## 7. Phase 5 — Supervised field pilot

**Objective:** evaluate real workflows in a small, representative population.

Controls:

- voluntary recruitment and clear withdrawal;
- trained support and incident channel;
- conservative capability limits;
- staged rollout and predefined automatic pause criteria;
- independent monitoring and scheduled safety review;
- no experimental withholding of essential protection.

Exit gate:

- benefits persist outside the laboratory;
- incident rates and subgroup disparities remain within approved bounds;
- users show stable or improved independent skill;
- support, rollback, and remedy processes work in practice.

## 8. Phase 6 — Trusted-circle warning network

**Objective:** share high-quality safety signals without centralizing personal content.

Build only after governance maturity:

- signed, minimal indicators with provenance and expiry;
- reputation based on evidence quality, not social popularity;
- resistance to brigading, rumor, and malicious reporting;
- local choice over subscription and action;
- appeal and correction propagation.

Exit gate:

- abuse resistance and governance are independently evaluated;
- personal messages are not uploaded by default;
- false reports can be corrected across the network;
- concentration of authority and conflicts are controlled.

## 9. Phase 7 — BootX integration research

**Objective:** explore a secure platform for selected validated functions.

Prerequisites:

- reproducible BootX source and builds;
- defined hardware target and threat model;
- process/memory isolation, signed updates, recovery, key protection, and test infrastructure;
- a stable companion service contract;
- no real personal data until security review approves it.

Start with a read-only demonstration and synthetic data. Networked tools and personal memory come later, if assurance supports them.

## 10. Workstreams

| Workstream | Primary outputs |
|---|---|
| Logic and evidence | claim schema, provenance, calibration, correction workflow |
| Ethics and governance | constitution, impact assessments, dissent and appeal |
| Human factors and education | accessible UI, curriculum, comprehension and dependency measures |
| Security and privacy | threat model, broker, memory controls, incident response |
| AI/model evaluation | grounded explanation, robustness, fairness, regression suites |
| Platform engineering | reproducible BootX, isolation, signed update, recovery |
| Community research | representative participation, field evaluation, distributional analysis |

## 11. Program metrics

Maintain a balanced scorecard:

- verified protective benefit;
- decision comprehension and calibration;
- false positives, false negatives, and alert burden;
- independent skill and dependency indicators;
- privacy and deletion conformance;
- security findings and recovery performance;
- accessibility completion;
- subgroup benefit and harm;
- correction and incident resolution time;
- portability and exit success;
- evidence maturity per AI DNA dimension.

Do not use engagement, conversation length, emotional attachment, or raw user trust as primary success measures.

## 12. Research publication standard

Every public study should include protocol, population, context, comparator, data exclusions, measures, uncertainty, negative results, incidents, conflicts, funding, limitations, and materials sufficient for replication where privacy permits. Illustrative examples must be labeled at the point of use.

## 13. Long-term horizon

Only after protective companionship is demonstrably safe should the program study broader education, disaster preparedness, family continuity, infrastructure resilience, or space-habitat applications. Each domain needs its own experts, data, risk analysis, law, and validation. AI DNA is a shared assurance language, not evidence that one successful use transfers automatically to another.
