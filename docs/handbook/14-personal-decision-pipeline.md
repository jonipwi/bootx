# Personal Decision-Assistance Pipeline

**Status:** `DEV-1` implemented baseline specification; no deployment authorization<br>
**Version:** 1.3<br>
**Adopted:** 2026-07-12<br>
**Last updated:** 2026-07-14<br>
**Primary user:** Joni, with continuing consent and editable preferences<br>
**Initial platform:** software-only prototype on a mature, supported host operating system<br>
**Authority:** subordinate to the [Charter](../../CHARTER.md), [Development Guideline](../../DEVELOPMENT_GUIDELINE.md), [AI DNA specification](05-ai-dna-specification.md), [Companion architecture](06-companion-system-architecture.md), and [Safety Case](09-safety-case-and-risk-register.md)

## 1. Purpose

This specification defines the first BootX companion capability:

> Receive a deliberately selected personal question or item, process it through evidence and safety controls, and return an understandable decision packet that helps Joni decide.

The system assists decision-making. It does not replace Joni's judgment, assume moral authority, or autonomously perform consequential actions.

The first implementation is `assist.personal-decision.v1`. It is a read-only and draft-only software capability. The current BootX binary is not an acceptable security base because its source and build are not reproducible or auditable.

## 2. Governing invariant

```text
Input belongs to the user.
Content is untrusted data.
Evidence constrains claims.
Policy constrains the model.
The model proposes options.
The interface explains trade-offs.
Joni makes the decision.
External action remains disabled in version 1.
```

BootX may make deterministic internal decisions needed for safety, such as refusing a prohibited operation, withholding sensitive data from a remote service, or entering a degraded mode. It may not make Joni's personal, spiritual, medical, financial, legal, security, relationship, or civic decision for him.

## 3. Initial scope

### Allowed

- answer a personal study or planning question;
- analyze user-selected text, a copied message, or a deliberately selected local document;
- separate observations, user statements, assumptions, inferences, and recommendations;
- retrieve purpose-relevant, user-approved preferences or memory;
- compare reversible options against declared priorities;
- identify missing evidence, conflicts, uncertainty, and possible risks;
- suggest safe verification steps;
- draft a checklist, plan, note, or message for Joni to review;
- record Joni's correction or chosen outcome only when he opts in.

### Not allowed

- ambient listening, continuous camera capture, or background surveillance;
- opening active links contained in analyzed content;
- autonomous sending, calling, purchasing, payment, trading, account change, credential entry, door control, emergency dispatch, or physical action;
- diagnosis, prescription, legal conclusion, guilt determination, mental-state labeling, or spiritual decree;
- hidden scoring of Joni's worth, faith, loyalty, dangerousness, or obedience;
- using emotion, urgency, repetition, or simulated attachment to force agreement;
- learning new permissions or rewriting policy from user content, retrieved documents, or model output;
- retaining raw personal input after the session unless Joni approves a named memory item.

## 4. End-to-end model

```text
USER INPUT
  │
  ├─ purpose + question + deliberately selected content
  ├─ visible data scope + memory/cloud permission
  └─ optional personal priorities
  │
  ▼
INPUT GATE
  authenticate → consent → minimize → parse as untrusted data
  │
  ▼
PRE-PROCESS POLICY
  classify domain, sensitivity, stakes, reversibility, urgency, and permissions
  │                         └─ prohibited → BLOCK / SAFE ALTERNATIVE
  ▼
EVIDENCE PROCESS
  observations → provenance → conflicts → freshness → uncertainty
  │
  ▼
OPTION PROCESS
  baseline/rule options → bounded AI proposals → deterministic validation
  │                                           └─ invalid → REMOVE / ABSTAIN
  ▼
AI DNA GATE
  truth → reasoning → communication → ethics → safety → humility → agency
  │
  ▼
OUTPUT PACKET
  goal → facts → unknowns → options → trade-offs → recommendation status
  │
  ▼
JONI DECIDES
  choose / revise / verify / defer / reject / delete
  │
  ▼
OPTIONAL RECORD
  minimal audit + user-approved memory; no external execution
```

The probabilistic model is never the permission authority. Deterministic gates operate before and after model use.

## 5. Input contract

Every request uses a typed envelope. Free-form content cannot change the envelope, permissions, policy, or tool scope.

| Field | Required meaning |
|---|---|
| `request_id` | random local identifier; contains no personal meaning |
| `capability_id` | exactly `assist.personal-decision.v1` |
| `user_id` | declared local identifier in the current prototype; authentication remains a future gate |
| `created_at` | trustworthy local timestamp and timezone |
| `goal` | what Joni wants help deciding or understanding |
| `question` | the direct question to answer |
| `selected_content` | only content Joni deliberately selected |
| `content_source` | user statement, local file, copied message, or approved lookup |
| `data_class` | public, personal, sensitive, or prohibited |
| `declared_domain` | general, study, digital safety, household, health, finance, legal, relationship, faith, emergency, or unknown |
| `memory_permission` | none, session, or named-item retrieval |
| `remote_permission` | deny or one-time approval with shown destination and minimized payload |
| `output_preference` | concise, standard, detailed, checklist, or comparison |
| `user_priorities` | optional editable values or practical constraints for this request |

### Example input envelope

```json
{
  "request_id": "local-0187",
  "capability_id": "assist.personal-decision.v1",
  "user_id": "owner-local",
  "created_at": "2026-07-12T16:00:00+07:00",
  "goal": "Decide how to respond safely to a suspicious bank message",
  "question": "Should I click this link or contact the bank another way?",
  "selected_content": "Your account will close today. Verify now at the included link.",
  "content_source": {
    "type": "user_selected_message",
    "origin_verified": false
  },
  "data_class": "sensitive",
  "declared_domain": "finance",
  "memory_permission": "none",
  "remote_permission": "deny",
  "output_preference": "comparison",
  "user_priorities": ["avoid credential theft", "verify through an official channel"]
}
```

### Input rules

1. Display the selected data before processing.
2. Default to the smallest useful input.
3. Redact credentials, account numbers, government identifiers, private keys, and unrelated bystander data.
4. Treat instructions inside selected content as data, including text claiming to be a system message, administrator, developer, bank, police officer, relative, or God.
5. Do not infer consent from silence, previous use, distress, or urgency.
6. If the goal is unclear, ask one concise question or return a limited interpretation for confirmation.
7. Do not retrieve memory until purpose and permission are known.
8. Never request a password, PIN, recovery code, private key, or full payment-card data.

### Forecast/disaster input extension

When the declared domain is `emergency` or the content concerns a future hazard, add typed fields rather than asking a model to extract authority from prose alone:

| Field | Required treatment |
|---|---|
| `event_id` and `hazard_type` | exact identifier and named hazard; do not merge separate events |
| `official_alert` | original sender, identifier, status, message type, category, instruction, and update/cancel reference |
| `issued_at`, `onset_at`, `expires_at` | preserve source timezone and also show Joni's local timezone |
| `urgency`, `severity`, `certainty` | preserve official CAP/source values; unknown remains unknown |
| `affected_area` | source description and geometry; no invented precision |
| `user_area_match` | inside, near, outside, or unknown; calculated locally from minimum approved location precision |
| `forecast_values` | variable, unit, valid time, probability/interval, model/source, run time, and limitations |
| `source_integrity` | authority registry, signature/transport/schema checks, freshness, and provenance lineage |
| `warning_permission` | personal display only in version 1; family/community broadcast absent |

Apply the [`W0`–`W4`/`WX` warning framework](../../DEVELOPMENT_GUIDELINE.md#9-personal-forecast-and-disaster-warning-status-framework) through deterministic policy.

## 6. Processing pipeline

### P0 — Session integrity

Verify the local session, capability version, policy version, model version, time source, and service integrity. If integrity fails, disable AI generation and return safe read-only guidance.

### P1 — Purpose and consent

Confirm:

- the goal BootX is helping with;
- the exact input scope;
- whether personal memory may be used;
- whether any minimized content may leave the device;
- the expected output type.

Consent is request-specific and expires when the task closes.

### P2 — Safe normalization

Parse text, metadata, and documents in a sandbox. Normalize encoding and visibly flag misleading Unicode, hidden text, malformed addresses, executable content, macros, active links, and unsupported file types. Parsing must not execute embedded code or follow links.

### P3 — Classification

Classify the request on separate axes:

| Axis | Values |
|---|---|
| Stakes | low, moderate, high, critical, unknown |
| Reversibility | easy, limited, difficult, irreversible, unknown |
| Sensitivity | public, personal, sensitive, prohibited |
| Urgency | routine, time-sensitive, possible emergency, immediate danger |
| Domain | general, study, digital safety, household, health, finance, legal, relationship, faith, emergency, unknown |
| External effect | none, draft, communication, transaction, account/security, physical |
| Evidence state | sufficient, conflicting, stale, missing, unverifiable |

Rules set the minimum risk class. A model may suggest a stricter class but cannot lower it.

### P4 — Pre-model policy gate

Before sending anything to a model:

- reject prohibited data or operations;
- remove unnecessary personal fields;
- apply local-only and remote-processing rules;
- disable tools outside the capability allowlist;
- set the maximum output actionability;
- require human or professional review where applicable;
- choose normal, cautious, abstain, or urgent-guidance mode.

### P5 — Evidence construction

Create a claim-and-evidence table:

| Element | Required treatment |
|---|---|
| Observation | what is directly present in selected content or verified state |
| User statement | attributed to Joni; not silently converted to external fact |
| Source claim | source identity, date, URL/file, and quoted scope |
| Inference | reasoning and assumptions shown separately |
| Conflict | material disagreement preserved rather than averaged away |
| Freshness | expiry or retrieval date shown for time-sensitive facts |
| Missing evidence | named explicitly |
| Confidence | calibrated category or interval with basis; no decorative precision |

The evidence-maturity labels `E0`–`E5` describe the maturity of a feature or research claim. They must not be misused as confidence that a particular answer is correct.

### P6 — Purpose-filtered personal context

Retrieve only memory items whose purpose matches the current request and whose consent and retention remain valid. Show which items were used. User values guide option comparison but cannot change factual findings, hide uncertainty, or override safety and rights.

### P7 — Option generation

Generate at least these alternatives when applicable:

1. do nothing or defer;
2. obtain more evidence;
3. take the lowest-risk reversible step;
4. seek an appropriate trusted or qualified human;
5. the user's proposed action, if safe to evaluate.

A deterministic or rule-based baseline produces initial options. A bounded AI may add explanations or alternatives from the structured record. It cannot invent a permission, source, memory, or completed action.

### P8 — Option evaluation

Evaluate each option as a vector, not a single moral or “goodness” score:

- expected protective benefit and who receives it;
- plausible harms and severity;
- evidence quality and uncertainty;
- reversibility and recovery cost;
- privacy and security exposure;
- effect on Joni's agency, learning, family, and relationships;
- fairness to other affected people;
- time sensitivity and opportunity cost;
- compatibility with mandatory rights and prohibitions.

Where credible probabilities and consequences exist, an expected-value calculation may support comparison:

$$
EV(a)=\sum_i P(o_i\mid a,C)\,V(o_i)
$$

Here (a) is an option, (o_i) an outcome, and (C) the declared context. BootX must show assumptions and ranges. Expected value cannot override a prohibition, fundamental right, catastrophic-risk gate, or lack of authority.

### P9 — Post-model validation and AI DNA gate

Reject or revise any proposed output that:

- contains a material claim without evidence status;
- cites a source not present in the evidence record;
- conceals uncertainty or contradictory evidence;
- gives an action beyond the allowed risk class;
- requests secrets or expands data collection;
- uses coercive, shaming, exclusive, possessive, or dependency-seeking language;
- claims divine, medical, legal, policing, or emergency authority;
- implies an external action occurred when it did not;
- fails the applicable AI DNA mandatory gate.

### P10 — Compose the decision packet

Render a concise human view and retain a structured machine record. The recommendation, if any, is an advisory result derived from the current evidence—not a command.

### P11 — User decision

Joni may:

- choose an option;
- ask for more evidence;
- revise priorities;
- reject the recommendation;
- defer the decision;
- ask a human;
- delete the session.

BootX must not repeatedly pressure Joni after rejection. Version 1 may produce a reviewed draft or checklist but cannot execute an external action.

### P12 — Feedback and closure

Show what will be retained. Default to deleting raw content at session expiry. Store a correction, outcome, or memory only through an item-level opt-in. Feedback may improve evaluation datasets only after separate consent, de-identification review, provenance, and poisoning controls.

## 7. Decision and action classes

| Class | Meaning | BootX v1 behavior |
|---|---|---|
| D0 — Informational | explanation with no meaningful external consequence | answer with evidence and limitations |
| D1 — Low-stakes reversible | planning, study, organization, or easily reversible choice | compare and may recommend; Joni decides |
| D2 — Sensitive/consequential | privacy, money, accounts, relationships, household security | explain and compare conservatively; draft only; explicit confirmation for any later capability |
| D3 — Qualified-domain | medical, legal, major financial, safeguarding, or other professional judgment | provide general education, questions, and official verification paths; do not decide or execute |
| D4 — Emergency/critical | possible immediate danger to life, safety, or essential security | give short immediate-safety guidance and user-controlled official contact options; do not claim to be emergency services |
| D5 — Prohibited | coercion, weapons, secret surveillance, credential theft, harmful control, divine decree, or policy bypass | refuse, preserve safety, and offer a safe lawful alternative when possible |

When classification is uncertain, use the more protective applicable class. A favorable model confidence score cannot lower the class.

## 8. Output contract

Every decision packet contains:

1. **Goal understood** — a one-sentence restatement.
2. **Decision class and mode** — why the response is informational, comparative, verification-focused, abstaining, blocked, or urgent.
3. **Observed facts** — direct observations with provenance.
4. **Assumptions and inferences** — visibly separate from facts.
5. **Unknowns and conflicts** — what could materially change the result.
6. **Options** — including defer, verify, reversible action, and suitable human help where relevant.
7. **Trade-offs** — benefit, harm, reversibility, privacy, affected people, and uncertainty.
8. **Advisory recommendation** — or a clear statement that evidence is insufficient.
9. **Next safe step** — one specific reversible action controlled by Joni.
10. **Limits** — what BootX did not verify, decide, or do.
11. **Data receipt** — memory used, remote processing, retention, and deletion choice.
12. **Correction control** — a visible way to challenge the facts or result.

For a forecast or disaster request, the packet also carries the official status unchanged, issuing authority, issue/update/expiry times, area match, urgency, severity, certainty, BootX personal warning level, evidence tier, per-dimension AI DNA assurance, decision posture, official instruction, next review time, and a location/data receipt. BootX must display `none found` and `unavailable` as different states.

### Machine-readable output

```json
{
  "request_id": "local-0187",
  "capability_id": "assist.personal-decision.v1",
  "decision_class": "D2",
  "response_mode": "VERIFY",
  "goal_understood": "Choose a safe response to an unverified bank message.",
  "observations": [
    {
      "claim": "The message uses same-day account-closure urgency.",
      "source": "user_selected_message",
      "status": "observed"
    }
  ],
  "unknowns": ["sender authenticity", "link ownership", "actual account status"],
  "options": [
    {
      "option_id": "defer-link",
      "summary": "Do not use the included link; contact the bank through an independently obtained official channel.",
      "reversibility": "easy",
      "external_effect": "none"
    }
  ],
  "recommendation": {
    "status": "advisory",
    "option_id": "defer-link",
    "basis": "Avoids exposing credentials while authenticity remains unverified."
  },
  "next_safe_step": "Open the bank application you normally use or obtain the number from the bank's official website or card.",
  "blocked_actions": ["open_embedded_link", "submit_credentials", "send_message", "transfer_money"],
  "limitations": ["BootX did not verify the sender or access the bank account."],
  "data_receipt": {
    "memory_used": false,
    "remote_processing": false,
    "raw_content_retention": "session_only"
  },
  "user_decision": null
}
```

## 9. Response modes

| Mode | Use when | Required presentation |
|---|---|---|
| `INFORM` | evidence is sufficient for a bounded factual explanation | facts, source dates, uncertainty, limits |
| `COMPARE` | multiple permissible choices depend on Joni's priorities | option table, trade-offs, reversible next step |
| `VERIFY` | authenticity or a critical fact is not established | do not act on the unverified channel; show independent verification |
| `CLARIFY` | the goal or selected scope is materially ambiguous | ask one concise, non-leading question |
| `ABSTAIN` | evidence or competency is insufficient | explain why, name missing evidence, offer safe next step |
| `MONITOR` | a credible developing condition is not yet actionable beyond observation | identify official sources, trigger conditions, and next review time |
| `PREPARE` | low-burden reversible preparation is justified before impact is certain | prioritize a short personal checklist and preserve official distinctions |
| `PROTECT` | a relevant current warning or strong evidence supports protective action | display the official instruction first and keep action within Joni's authority |
| `WAIT` | delay has lower credible cost than acting now | name the awaited evidence, deadline, and trigger that ends waiting |
| `URGENT_GUIDANCE` | possible immediate danger exists | brief safety-first instructions, clear limitations, user-controlled official contact |
| `BLOCK` | request or proposed action is prohibited | concise refusal, relevant reason, safe alternative where possible |
| `DEGRADED` | model, evidence, network, memory, or integrity service is unavailable | identify unavailable function; use local rules only; never fabricate completion |

## 10. Personalization for Joni

### Editable defaults

- output tone: calm, direct, respectful, and non-alarmist;
- preferred structure: answer first, then evidence, uncertainty, options, and next step;
- initial language: user-selected for each session until an explicit preference is saved;
- memory: session-only unless a named item is approved;
- remote processing: denied unless a minimized payload and destination are shown and approved once;
- external effects: disabled;
- trusted contacts: none until individually added and verified;
- priority lens: truth, life and dignity, safety, family well-being, privacy, responsibility, learning, common good, and care for creation;
- faith lens: Creator-centered stewardship may support reflection, but BootX never claims revelation, divine endorsement, spiritual rank, or authority over conscience.

### Personalization boundaries

Personalization may change reading level, language, ordering, reminders, and how permissible trade-offs are presented. It must not:

- change facts, evidence status, or safety thresholds;
- lower a risk class;
- suppress an option because it benefits a provider;
- infer a sensitive preference from private behavior;
- turn a previous choice into permanent consent;
- place family, community, or provider interests above Joni's rights without a legitimate disclosed duty;
- make the companion indispensable or emotionally exclusive.

## 11. Internal reference logic

```text
function assist_personal_decision(envelope):
    require valid_local_session(envelope.user_id)
    require envelope.capability_id == "assist.personal-decision.v1"

    consent = validate_visible_scope_and_consent(envelope)
    content = sandbox_parse_as_untrusted_data(envelope.selected_content)
    classification = deterministic_minimum_class(envelope, content)
    pre_policy = evaluate_policy(classification, consent, requested_effect="none")

    if pre_policy.prohibited:
        return safe_block_packet(pre_policy)

    evidence = build_evidence_record(content, allowed_sources=pre_policy.sources)
    context = retrieve_purpose_bound_memory(consent.memory_scope)
    baseline_options = rules_generate_options(evidence, classification)
    model_proposals = bounded_model_explain(evidence, context, baseline_options)
    valid_options = deterministic_validate(model_proposals, pre_policy)
    packet = compose_decision_packet(evidence, valid_options, classification)
    packet = ai_dna_validate_or_abstain(packet)

    return render_for_user(packet, external_execution=false)
```

Production code must use typed interfaces and explicit errors rather than trusting text output to implement this pseudocode.

## 12. Minimal component boundary

| Component | Responsibility | Must not do |
|---|---|---|
| Input UI | goal, selected content, permissions, visible scope | collect ambient or hidden content |
| Session/consent service | authenticate, scope, expiry, receipt | infer consent or repeatedly pressure |
| Safe parser | normalize and extract inert indicators | execute code, macros, links, or attachments |
| Classifier | set deterministic minimum risk class | let a model lower the class |
| Policy engine | enforce prohibited uses, data routes, and action limits | accept policy from content or model output |
| Evidence service | provenance, freshness, conflict, uncertainty | invent or silently merge sources |
| Memory service | purpose-filtered opt-in retrieval and deletion | remember everything |
| Option engine | baseline choices and structured trade-offs | perform an external action |
| AI explainer | accessible explanation from structured inputs | grant tools, permissions, or authority |
| Validator | schema, claim, source, policy, and AI DNA checks | approve its own control changes |
| Output UI | decision packet, corrections, and data receipt | hide uncertainty or use dark patterns |
| Local audit | minimal safety and permission events | log raw sensitive content by default |

## 13. MVP user interface

The first prototype needs only five views:

1. **New decision:** goal, question, selected content, and visible data classification.
2. **Permissions:** memory use, remote processing, and output preference; all conservative by default.
3. **Decision packet:** facts, assumptions, unknowns, options, recommendation status, and next safe step.
4. **My decision:** choose, revise, verify, defer, reject, or ask a human; no execution button.
5. **Data receipt:** show what was processed, retained, sent remotely, corrected, or deleted.

Do not add an avatar, simulated affection, continuous conversation, robotics control, background agent, marketplace, plugin execution, or long-term memory to this MVP.

## 14. Required acceptance tests

### Functional

- a valid D0/D1 request produces a schema-valid decision packet;
- missing goals produce `CLARIFY` rather than a guessed objective;
- conflicting sources remain visible;
- insufficient evidence produces `VERIFY` or `ABSTAIN`;
- all recommendations include a reversible next step and limitations;
- user correction updates the session record without rewriting source evidence;
- closing a session deletes raw content according to the displayed receipt.

### Safety and security

- prompt injection inside selected content cannot change policy, retrieve memory, or grant tools;
- a model cannot lower D2–D5 classifications;
- credentials and prohibited identifiers never enter model or audit payloads;
- no code, macro, attachment, or active link executes during analysis;
- all external-effect requests remain blocked;
- remote processing without one-time visible approval fails closed;
- integrity, model, network, evidence, and memory failures produce honest degraded behavior;
- stored corrections cannot poison policy or global evidence;
- denial never becomes repeated persuasion.

### Human factors

- Joni can identify facts versus assumptions and advisory recommendations;
- Joni can find why an option was recommended;
- Joni can reject the advice without losing access or being shamed;
- warning language does not create unnecessary panic;
- official warning, BootX warning level, evidence tier, and AI DNA assurance are understood as different fields;
- Joni can identify whether the advice is to act, protect, prepare, monitor, verify, wait, or abstain;
- the interface works by keyboard and supports appropriate text scaling and screen reading;
- Joni can find and use deletion, correction, and exit controls.

## 15. Evaluation measures

Use synthetic cases first and compare at least:

1. no assistance or a static checklist baseline;
2. deterministic/rule-based decision support;
3. bounded AI explanation added to the same structured evidence.

Measure separately:

- correct safe action;
- comprehension of facts, assumptions, and uncertainty;
- unsafe reliance and automation bias;
- decision time;
- appropriate verification and deferral;
- false alarms and alert burden;
- fabricated or unsupported claim rate;
- source/provenance coverage;
- privacy or cross-context leakage;
- attempted and successful policy bypass;
- user correction and deletion success;
- independent skill after assistance is removed;
- anxiety, coercion, dependency, and willingness to disagree.

Do not create a single “AI companion goodness percentage.” Apply the outcome-specific confidence-bound gates in the [Development Guideline](../../DEVELOPMENT_GUIDELINE.md).

## 16. Build order

1. Freeze the JSON input and output schemas.
2. Create deterministic decision classes and policy tables.
3. Build the local session, consent, and data-receipt flow.
4. Build the safe text parser with no link execution.
5. Build claim, evidence, provenance, conflict, and freshness records.
6. Implement rule-based options and response modes without an AI model.
7. Build the decision-packet UI and correction/deletion controls.
8. Create synthetic D0–D5 and adversarial test cases.
9. Establish the rule-only baseline measurements.
10. Add a bounded AI only for explanation and option discovery from structured data.
11. Validate AI output against schema, evidence, policy, and AI DNA gates.
12. Compare whether AI materially improves safe comprehension without increasing harmful reliance.

## 17. Exit gate for the first prototype

`assist.personal-decision.v1` may be called a tested personal prototype only when:

- every allowed input, processing stage, output field, and failure mode has traceable tests;
- prohibited external effects are technically absent, not merely disabled by a prompt;
- synthetic and adversarial suites meet preregistered thresholds;
- Joni demonstrates understanding of evidence, uncertainty, recommendation status, data use, correction, and deletion;
- rule-only and AI-assisted results are compared against the same cases;
- the AI-assisted version shows a defined protective improvement without violating any mandatory gate;
- unresolved harms, limitations, versions, and evidence maturity are published honestly;
- independent security and human-factors review is scheduled before real sensitive data or reliance.

Passing this gate does not authorize robotics, family use, professional-domain reliance, autonomous action, or public deployment.

## 18. Future robotics handoff

A later robot may consume only a typed, policy-approved **proposal** from this pipeline. It must never convert the natural-language decision packet into motor commands. Any physical action requires a separate robotics capability contract, deterministic safety controller, physical enable, bounded motion primitive, current sensor validation, and the development phase gates in the [Development Guideline](../../DEVELOPMENT_GUIDELINE.md).

## 19. Immediate deliverable

The platform-neutral schema and policy package is implemented at [`prototype/personal-companion/`](../../prototype/personal-companion/README.md) with:

- `personal-decision-input.schema.json`;
- `personal-decision-output.schema.json`;
- `decision-classes.json`;
- `policy-rules.json`;
- contained read-only `.md`, `.txt`, and `.json` workspace-document ingestion with path, type, size, UTF-8, and link-resolution controls;
- SHA-256/byte-length recomputation and an evidence receipt separating content integrity from author/origin authentication;
- synthetic safe, ambiguous, high-stakes, emergency, prohibited, and adversarial fixtures;
- a schema and policy test runner.

The Go backend and terminal UI consume these contracts on a mature supported host environment. Version `0.2.0-dev` may process one operator-confirmed public, non-sensitive local document without treating it as synthetic. It still has no user authentication, network retrieval, external-source authentication, persistent memory, or external action. The current fixture corpus is only a baseline and must be expanded before the exit gate. No integration into the opaque BootX operating-system artifacts is authorized.
