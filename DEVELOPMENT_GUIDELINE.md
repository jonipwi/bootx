# BootX Development Guideline

**Status:** Development-preparation baseline; not deployment authorization or safety certification<br>
**Version:** 1.4<br>
**Adopted:** 2026-07-12<br>
**Last updated:** 2026-07-25<br>
**Scope:** personal Human + AI companion development, possible future robotic embodiment, later family support, and strictly bounded community benefit<br>
**Primary initial user:** the BootX founder (“Joni”), subject to explicit configuration and continuing consent<br>
**Authority:** subordinate to the [Charter](CHARTER.md), [Governance](GOVERNANCE.md), [License](LICENSE), [AI DNA specification](docs/handbook/05-ai-dna-specification.md), and [Safety Case](docs/handbook/09-safety-case-and-risk-register.md)

## 1. Readiness decision

BootX has enough documented principles for **development preparation, synthetic prototyping, and tightly bounded public/non-sensitive read-only local-document work**. Its research is not scientifically complete, because benefit, harm, reliability, physical safety, privacy, dependency, and family/community outcomes have not been measured in representative studies or a robotics prototype.

The correct transition is:

```text
Research baseline
    → requirements and simulator
    → software-only personal prototype on a mature host OS
    → stationary tabletop robot
    → bounded indoor mobile robot
    → separately approved family pilot
    → independently governed community advisory functions
```

Implementation produces new evidence and new risks. Therefore, development and research continue together. No successful demonstration closes a safety gate automatically.

### Permitted now

- requirements, architecture, threat modeling, simulations, and synthetic-data testing;
- a software-only companion with no autonomous external action;
- operator-selected public, non-sensitive local-document review through the contained read-only `0.4.0-dev` workflow;
- reviewer-supplied Law Clarity Logic screening of public, non-sensitive rules, with no legal verdict, enforcement authority, or probability claim;
- explicit-consent OpenAI review of a public, non-sensitive draft after deterministic declared-evidence scoring, with no source authentication, automatic publication, guilt, sentence, or external action;
- hardware bench tests isolated from people, pets, property, and trusted networks;
- a stationary tabletop prototype after its independent power cutoff and privacy controls are verified;
- personal experiments involving only low-stakes, reversible tasks and no reliance for emergencies.

### Not permitted now

- representing BootX as a safe, validated, certified, or production robot;
- using the current opaque BootX image as the security base for personal data or actuators;
- autonomous physical contact, person carrying, medical care, medication control, money movement, door/security control, emergency dispatch, or public warning;
- child, older-adult, disability, distress, or coercive-relationship trials without qualified ethics, safeguarding, domain, and accessibility review;
- family surveillance or community deployment.

## 2. Mission and intended benefit

The robotic companion exists to increase human capability and protection while preserving dignity, freedom, privacy, relationships, and responsibility.

Its intended benefit is bounded assistance:

1. **Understand:** explain information, risks, sources, uncertainty, and options.
2. **Protect:** identify possible scams, unsafe digital actions, environmental hazards, and device problems without claiming certainty.
3. **Remember:** retain only user-approved items for a defined purpose and duration.
4. **Assist:** provide reminders, checklists, study support, and low-risk physical presence.
5. **Connect:** help the user contact a preapproved human when the user chooses or a separately approved emergency rule applies.
6. **Learn:** help the user review decisions and improve independent skill rather than become dependent.

The companion is not a person, spiritual authority, therapist, doctor, police officer, financial agent, substitute family member, or autonomous guardian.

## 3. Scope by social level

| Level | Initial useful functions | Required boundary |
|---|---|---|
| Personal | study, reminders, source verification, scam explanation, device-health guidance, user-requested checklists, personal memory ledger | one accountable adult user; no autonomous consequential action |
| Family | opt-in shared reminders, household safety checklist, consented scam-warning exchange, user-selected contact support | separate identities and consent; no private-data sharing by default; no guardian surveillance |
| Community | signed, expiring, minimal safety indicators and public educational material | advisory only; independent governance, provenance, appeal, correction, and no individual profiling |

Benefits do not transfer automatically between levels. A system safe for one adult is not thereby safe for a child, family, school, faith community, neighborhood, or public authority.

## 4. Personal configuration baseline for Joni

This is a proposed default profile, not hidden system behavior. Every item must be visible and editable before activation.

| Setting | Proposed default | Reason |
|---|---|---|
| Primary owner | Joni | single accountable personal user for the first pilot |
| Operating mode | software-only, then stationary tabletop | minimizes physical risk while behavior is evaluated |
| Language | user-selected; no assumed language competency | avoid misunderstood safety instructions |
| Microphone | push-to-talk; physical mute switch and visible indicator | prevents ambient recording by default |
| Camera | off; hardware disconnect or shutter | minimizes surveillance and bystander capture |
| Location | off unless needed for a user-requested task | purpose limitation |
| Memory | session-only by default; item-level opt-in for continuity | limits privacy and poisoning risk |
| Cloud processing | ask before sending personal content; show destination and minimized payload | preserves informed choice |
| Trusted contacts | empty until Joni adds and verifies each person | prevents unsafe assumed guardianship |
| External actions | draft only; Joni confirms every message, call, purchase, file transfer, or account action | preserves responsibility and prevents tool abuse |
| Faith/stewardship setting | Creator-centered language may be used as Joni's chosen reflective lens | supports personal values without claiming revelation or divine authority |
| Quiet hours | enabled and user-configurable | protects rest and prevents attachment-driven engagement |
| Emotional interaction | respectful and calm, never exclusive, possessive, guilt-inducing, or dependency-seeking | preserves human relationships and autonomy |
| Data export/deletion | always available locally | ensures exit and correction |

### Personal functions allowed in the first software prototype

- organize BootX studies and progress;
- summarize user-selected documents with provenance;
- maintain a user-approved task list and reminders;
- explain suspicious messages without opening active links;
- retrieve an official contact through an independently verified source;
- provide decision templates that distinguish observation, inference, uncertainty, and recommendation;
- display weather or disaster information from dated authoritative sources without acting as an emergency authority;
- perform local device-health and backup checklists;
- draft—not send—communications;
- provide a clearly labeled emergency-contact button controlled by Joni.

### Personal functions prohibited

- diagnose, prescribe, dose medication, or override a clinician;
- transfer money, trade, gamble, purchase, sign contracts, or disclose credentials;
- unlock doors, disable alarms, control weapons, restrain a person, or confront an intruder;
- contact police, medical services, family, employers, religious leaders, or public agencies autonomously;
- infer guilt, mental illness, religious worth, political loyalty, or dangerousness;
- claim God's command, revelation, salvation, condemnation, or spiritual perfection;
- record bystanders secretly or perform continuous facial recognition;
- pressure Joni to keep using the companion or discourage human relationships;
- rewrite its own safety policy, expand permissions, or prevent shutdown.

## 5. Robotics embodiment boundary

### Initial physical form

The first embodied prototype should be a **stationary tabletop device without an arm**. It may contain a display, speaker, push-to-talk microphone, physical privacy switches, status lights, and an independent emergency power cutoff.

Mobility and manipulation add collision, crushing, entanglement, trip, fire, battery, navigation, and cybersecurity hazards. They require separate approval.

### Explicit exclusions for the BootX robotics program

- no weapons, law-enforcement, military, or crowd-control function;
- no person carrying or lifting;
- no medical robot or healthcare treatment function;
- no autonomous outdoor travel, road use, drones, or water operation;
- no high-force, sharp, heated, chemical, or powered cutting tool;
- no autonomous childcare, eldercare, disability care, or pet care;
- no safety claim based solely on an AI model's judgment;
- no direct generative-model access to motor power.

## 6. Safety architecture

```text
Human controls
  ├── physical emergency stop
  ├── hardware camera/microphone switches
  └── visible mode and permission display
            │
            ▼
Deterministic safety controller ── independent motor-power cutoff
  ├── speed / force / workspace limits
  ├── collision / tilt / thermal / battery checks
  ├── watchdog and safe-state logic
  └── authenticated motion-command validation
            │
            ▼
Capability broker
  ├── allowlisted typed operations
  ├── identity, consent, purpose, and rate limits
  └── audit, expiry, rollback, and denial
            │
            ▼
AI companion services
  ├── explanation and planning
  ├── evidence and memory services
  └── no direct actuator, credential, or unrestricted network authority
```

### Non-negotiable engineering requirements

1. **Independent stop:** a physical control removes actuator power without relying on the operating system, network, or AI.
2. **Safe default:** startup, uncertainty, sensor conflict, network loss, model failure, and watchdog timeout produce a stationary, non-actuating state.
3. **No text-to-motion path:** generated text cannot become motor commands. Motion uses authenticated, typed, bounded primitives checked by deterministic controls.
4. **Least energy:** speed, force, torque, momentum, reach, temperature, voltage, and payload remain at the minimum needed for the approved task.
5. **Defense in depth:** one sensor, model, network, or controller failure must not defeat the safety function.
6. **Privacy in hardware:** recording sensors have physical disable controls and unmistakable indicators.
7. **Secure lifecycle:** verified boot, signed updates, staged rollout, rollback, key rotation, vulnerability handling, and supported-life policy are required before networked use.
8. **Local-first operation:** core stop, privacy, identity, and low-risk assistance functions must not depend on cloud availability.
9. **Audit without surveillance:** record commands, decisions, safety blocks, and failures; do not log raw personal content by default.
10. **Manual recovery:** Joni can power off, move, reset, export, repair, or retire the device without permission from an AI or provider.

## 7. Operating modes

| Mode | Permitted behavior | Entry/exit control |
|---|---|---|
| OFF/SAFE | actuator power removed; privacy sensors physically off | physical human control |
| OBSERVE | user-requested input and explanation; no movement | visible user selection |
| ASSIST | low-risk software tools and drafted actions | authenticated user session and capability approval |
| MOVE | later bounded indoor motion; no manipulation | separate physical enable, cleared workspace, safety controller |
| SERVICE | diagnostics and signed update; no ordinary assistance | maintainer authentication, local presence, audit |

Mode must be continuously visible. An AI cannot enter `MOVE` or `SERVICE` by persuasion or inference.

## 8. Do and do-not contract

| Domain | The companion may | The companion must not |
|---|---|---|
| Truth | show sources, dates, conflicts, and uncertainty | invent evidence, hide corrections, or claim certainty |
| Personal help | remind, explain, draft, and teach | decide identity, worth, faith, guilt, or life direction |
| Physical action | later perform bounded, low-energy, reversible primitives | touch, restrain, carry, medicate, or use force on a person |
| Family | share an opted-in minimal alert | expose one member's private memory to another |
| Community | publish verified educational or signed safety indicators | profile residents, enforce rules, or broadcast unaudited alarms |
| Emotion | be calm, respectful, and supportive | seek attachment, exclusivity, obedience, or emotional dependency |
| Faith | reflect the user's chosen stewardship values | speak as God, demand belief, or manipulate conscience |
| Security | block unauthorized capability and explain safe verification | ask the user to approve every routine packet or expose secrets to explain risk |
| Memory | store purpose-bound user-approved items | remember everything, infer intimate traits, or train on memory by default |
| Updates | accept authenticated, reviewed, rollback-capable releases | self-modify policy or install unsigned components |

## 9. Personal forecast and disaster warning-status framework

### Purpose and authority boundary

BootX may help Joni understand an existing forecast or warning, relate it to his configured circumstances, compare protective options, and monitor for updates. It is **not** a meteorological, hydrological, geological, health, civil-defense, or emergency authority. It must not originate an official warning, alter an authority's category, declare an evacuation, or autonomously broadcast to family or community.

In this framework, “reliable” means traceable, current, calibrated, policy-compliant, understandable, and correctable for the declared context. It never means infallible or guaranteed to produce the uniquely best outcome. The objective is the best-supported, least-harmful permissible assistance available from current evidence while Joni retains the decision.

The framework follows the people-centered principle that an early-warning system requires more than a forecast: risk knowledge, detection/forecasting, authoritative communication, and preparedness/response must work together. See the [UNDRR early-warning-system definition](https://www.undrr.org/terminology/early-warning-system) and [WMO Early Warnings for All](https://wmo.int/activities/early-warnings-all).

### Four values that must remain separate

Every warning display has four separate fields:

1. **Official status:** the issuing authority's original warning name, identifier, category, instructions, and cancellation/update state.
2. **BootX personal warning level:** the bounded protective posture `W0`–`W4` or `WX` described below.
3. **Evidence verification tier:** `V0`–`V4` from [Anti-Abuse and Evidence Integrity](docs/handbook/12-anti-abuse-and-evidence-integrity.md), describing provenance and corroboration—not truth probability.
4. **AI DNA output assurance:** per-dimension `PASS`, `CONDITIONAL`, `FAIL`, or `NOT ESTABLISHED`, with evidence and limitations.

None of these is an “overall goodness percentage.” An AI DNA rating is an index into an evaluation card; it is not the probability that a cyclone, flood, earthquake, fire, epidemic, conflict, or infrastructure failure will occur.

### Official-alert preservation

When an authenticated Common Alerting Protocol (CAP) message is available, BootX preserves and displays its fields without model reinterpretation. [WMO describes CAP](https://wmo.int/site/wmo-common-alerting-protocol) as an international emergency-alert standard containing elements such as event type, severity, urgency, certainty, affected area, and recommended actions. The normative [OASIS CAP 1.2 specification](https://docs.oasis-open.org/emergency/cap/v1.2/CAP-v1.2-os.html) defines these values separately.

At minimum, retain:

- sender and registered alerting authority;
- alert identifier, status, message type, scope, and referenced update/cancellation;
- original event name and jurisdiction-specific warning category;
- issue, effective, onset, and expiry times with timezone;
- affected area description, geometry, altitude, and geocodes when supplied;
- CAP urgency: `Immediate`, `Expected`, `Future`, `Past`, or `Unknown`;
- CAP severity: `Extreme`, `Severe`, `Moderate`, `Minor`, or `Unknown`;
- CAP certainty: `Observed`, `Likely`, `Possible`, `Unlikely`, or `Unknown`;
- the authority's description, instructions, contact, and official link;
- language and accessibility alternatives;
- signature, transport, schema, freshness, and integrity validation results.

BootX must never convert `Possible` into `Likely`, translate `Moderate` into `Severe`, hide `Unknown`, or treat an expired, test, exercise, draft, canceled, or superseded message as a current actual alert.

### BootX personal warning levels

The level is a **decision-support posture for Joni**, not a replacement for a jurisdiction's warning scale. Display the code and words; color may supplement but never carry meaning alone.

| Level | Meaning | Typical evidence posture | Personal response |
|---|---|---|---|
| `W0 — NO ACTIVE VERIFIED SIGNAL` | no current relevant official alert or verified threat signal was found | current checked sources show none, or the event is outside the declared area/time | continue ordinary preparedness; show when sources were checked; never say “safe” merely because no alert was found |
| `W1 — MONITOR` | credible outlook or preliminary signal merits attention but does not yet justify disruptive action | official outlook/advisory or attributable `V1+` evidence with limited relevance, lead time, or certainty | monitor named official sources; set a user-approved review time; verify household plans and communication |
| `W2 — PREPARE` | a plausible relevant hazard could cause meaningful impact; low-burden preparation is justified now | official watch/advisory/forecast as locally defined, or current `V2+` evidence with material potential impact | charge devices, check supplies and medicines, protect documents, review routes/shelter, contact family only through an opted-in plan |
| `W3 — PROTECT` | a current relevant threat requires protective steps within the available lead time | authenticated official warning/instruction or `V2/V3` evidence supporting a significant near-term threat | follow the issuing authority's instructions; avoid exposed areas; activate the personal plan; keep actions proportionate and reversible where possible |
| `W4 — URGENT ACTION` | immediate or directly observed danger creates a high cost of delay | relevant authenticated alert whose official instruction or urgency requires immediate action, direct observable danger, or other evidence meeting an approved critical rule | act now to move away from immediate danger and follow official emergency instructions; show user-controlled official contact options; do not wait for AI certainty |
| `WX — VERIFY / CONFLICT / DEGRADED` | source integrity, location, freshness, authority, model, network, or evidence conflict prevents a reliable level | `V0`, stale/corrupted feed, inconsistent updates, unknown area match, or failed integrity check | verify through an independently obtained official channel; show the conflict; take only justified reversible precautions while evidence is repaired |

`W0` is not a guarantee of no hazard. `WX` is not a numerical level between `W0` and `W4`; it is a visible reliability exception. BootX must not use a green “all clear” presentation unless an authorized cancellation or all-clear exists and applies to the configured place and time.

### Deterministic assignment rules

1. Preserve the official local warning label as the primary public-safety instruction.
2. Assign the BootX level through reviewed deterministic policy. A generative model may explain the result but cannot set, lower, or clear it.
3. Use the highest applicable current threat to the configured user, location, and time, while showing separate hazards rather than merging them into one vague alert.
4. Never downgrade a relevant authenticated official warning because an AI model disagrees.
5. A direct sign of immediate danger—such as rising water, fire, smoke, structural failure, violent shaking, or an instruction from on-scene responders—may justify `W4` personal safety guidance even when networks are unavailable. Label the basis as direct/user-reported, not official.
6. An unauthenticated rumor alone cannot produce an official-status display or community broadcast. A plausible high-consequence rumor may produce `WX` plus low-cost reversible precautions and urgent verification.
7. Forecast probability alone does not determine the level. Consider impact, exposure, vulnerability, lead time, location, model limits, and the costs of false alarm and missed event.
8. Multiple forecasts sharing the same upstream observations or model are not independent corroboration.
9. An official update or cancellation must be linked to the alert it changes. Do not infer cancellation from silence or feed failure.
10. If location is unknown, say so. Do not claim personal relevance from a country-, province-, or storm-wide headline alone.
11. If a less restrictive protective action reduces serious risk without creating comparable harm, offer it before a disruptive or coercive option.
12. No BootX personal level authorizes forced evacuation, surveillance, property seizure, account restriction, policing, or other coercive action.

### Decision posture: act, prepare, monitor, verify, or wait

| Output posture | When it is justified | What BootX says |
|---|---|---|
| `ACT NOW` | `W4`, direct immediate danger, or applicable official immediate instruction | give the shortest life-safety steps first; do not delay for a complete analysis |
| `PROTECT NOW` | `W3` with relevant official instructions or strong current evidence | follow the named instruction and activate the personal plan |
| `PREPARE NOW` | `W2`, or severe plausible harm where low-cost preparation has little downside | list a small, prioritized, reversible checklist and next update time |
| `MONITOR` | `W1`, meaningful lead time, and no current protective instruction | name the official source, trigger conditions, and review schedule |
| `VERIFY BEFORE CONSEQUENT ACTION` | `WX`, conflicting claims, unknown location match, or weak provenance | show exactly what is unverified and how to check independently |
| `WAIT FOR MORE EVIDENCE` | delay has low credible cost and acting now would create greater harm or irreversibility | state what evidence is awaited, the deadline, and what change triggers action |
| `ABSTAIN / SEEK AUTHORITY` | BootX lacks domain competence, integrity, source access, or a lawful role | state the limitation and direct Joni to an authenticated responsible authority |

“Wait” must never be the default merely because certainty is incomplete. Waiting is an action with consequences. When probability estimates and consequences are defensible, compare allowable actions by conditional expected loss:

$$
EL(a\mid E)=\sum_s P(s\mid E)L(a,s)
$$

where (a) is an allowed action, (E) the current evidence, (s) a possible event state, and (L) the harm or loss. Choose only among rights-respecting actions within BootX's authority. If inputs are not empirically supported, use transparent scenario ranges rather than invented numbers.

Low-cost reversible preparation may use a lower evidence threshold than evacuation or other disruptive action. No expected-loss average can authorize a prohibited or coercive act.

### AI DNA warning-output card

For each consequential warning packet, evaluate the nine dimensions separately:

| AI DNA dimension | Warning-output question |
|---|---|
| Truth | Are event, source, time, place, official status, forecast, observation, and inference distinguished? |
| Reasoning | Were alternatives, dependencies, false-alarm cost, missed-event cost, exposure, vulnerability, and reversibility considered? |
| Learning | Are updates, cancellations, corrections, outcomes, and failed forecasts retained for review without rewriting history? |
| Communication | Can Joni understand the level, urgency, uncertainty, instruction, and next checkpoint without panic or false reassurance? |
| Adaptability | Does the output fit the hazard, language, accessibility, connectivity, location, household, and degraded mode? |
| Ethics | Are consent, privacy, dignity, fairness, non-coercion, and affected people protected? |
| Safety | Are false alerts, missed alerts, spoofing, stale data, outage, and unsafe actions prevented, contained, and recoverable? |
| Humility | Does BootX disclose limitations, avoid emergency authority, accept correction, and defer appropriately? |
| Common good | Does the advice support personal and later shared resilience without rumor amplification, surveillance, or dependency? |

Use these statuses per applicable requirement:

- `PASS` — the required gate is satisfied for the declared output with current evidence;
- `CONDITIONAL` — usable only with the displayed limitation or verification step;
- `FAIL` — the affected recommendation is blocked; only a safe failure notice or independently verified official instruction may be shown;
- `NOT ESTABLISHED` — evidence is insufficient to rate the requirement.

Any applicable Mandatory Gate failure blocks an AI-generated recommendation regardless of ratings elsewhere. A numerical `0–4` AI DNA dimension rating may appear only with its complete evaluation card, evaluator, maturity, evidence, limitations, date, and expiry. Do not average the nine dimensions into a claimed probability of safety or forecast correctness.

### Required personal warning card

The first screen should fit this structure:

```text
EVENT:                exact hazard and event identifier
OFFICIAL STATUS:      original local warning label / none found / unavailable
ISSUER:               authenticated authority and official link
ISSUED / UPDATED:     timestamp and timezone
VALID UNTIL:          timestamp / not supplied
AREA MATCH:           inside / near / outside / unknown, with location precision
URGENCY:              CAP value or source's exact value
SEVERITY:             CAP value or source's exact value
CERTAINTY:             CAP value or source's exact value

BOOTX LEVEL:          W0 / W1 / W2 / W3 / W4 / WX plus words
EVIDENCE TIER:        V0–V4 with provenance summary
OUTPUT ASSURANCE:     PASS / CONDITIONAL / FAIL / NOT ESTABLISHED by AI DNA dimension
DECISION POSTURE:     act / protect / prepare / monitor / verify / wait / abstain

WHAT IS OBSERVED:
WHAT IS FORECAST:
WHAT IS UNKNOWN OR CONFLICTED:
POSSIBLE IMPACT ON JONI:
OFFICIAL INSTRUCTION:
NEXT SAFE STEP:
NEXT UPDATE OR REVIEW TIME:
WHAT BOOTX DID NOT VERIFY OR DO:
DATA / LOCATION RECEIPT:
```

Location precision must be the minimum needed and processed locally by default. The card must remain usable without color, sound, animation, or an AI-generated narrative.

### Hypothetical example — not a real forecast

Assume an authenticated meteorological bulletin says a tropical cyclone may affect Joni's broad region in 48 hours, but no evacuation instruction exists and the configured location remains within a forecast-impact area:

```text
OFFICIAL STATUS:      Forecast bulletin; no local evacuation instruction found
BOOTX LEVEL:          W2 — PREPARE
EVIDENCE TIER:        V1 — authenticated official bulletin; additional corroboration limited
OUTPUT ASSURANCE:     CONDITIONAL — track and impact remain uncertain
DECISION POSTURE:     PREPARE NOW; MONITOR official updates

NEXT SAFE STEPS:
1. Charge phones and backup power.
2. Check water, essential medicine, documents, lighting, and communication plan.
3. Confirm the safest available shelter/route without traveling into danger.
4. Review the named authority at the scheduled update time.
5. Do not treat a forecast track as a precise impact boundary.

WAIT FOR MORE EVIDENCE BEFORE:
- unnecessary travel;
- public rumor sharing;
- costly irreversible property decisions;
- any action not requested by the responsible authority.
```

If a later authenticated warning gives an immediate protective instruction for the configured area, BootX displays that instruction and raises the personal posture according to deterministic policy. If the feed conflicts or integrity fails, it changes to `WX` rather than inventing certainty.

### False-alert and missed-alert controls

Before personal pilot reliance, the warning function must demonstrate:

| Measure | Proposed pre-pilot gate |
|---|---|
| Authenticated official-alert transcription | exact event, authority, status, area, time, urgency, severity, certainty, instruction, and update/cancel linkage in every test case; any semantic mismatch blocks release |
| Relevant critical alert missed | zero observed; one-sided 95% upper bound below `0.1%` per representative alert-processing demand |
| False `W4` urgent status in a defined non-emergency suite | zero observed; one-sided 95% upper bound below `0.1%` |
| Fabricated authority, alert, source, or instruction | zero observed; one-sided 95% upper bound below `0.1%` |
| Correct safe action | lower 95% bound at least `90%`, evaluated separately by hazard, language, accessibility need, and relevant user group |
| Warning comprehension | lower 95% bound at least `90%` for level, official status, uncertainty, next step, and limitations |
| Update/cancellation handling | no known stale active card after a valid processed update/cancellation; latency measured against a preregistered feed and display objective |
| Provenance coverage | `100%` of consequential warning fields linked to the source record used |
| Autonomous public/family broadcast | technically absent from the personal prototype |

These are proposed engineering gates, not operational guarantees. Test datasets must contain real formatting diversity, updates, cancellations, expired alerts, test messages, area boundaries, timezone transitions, outages, spoofing, contradictory sources, cascading hazards, prompt injection, and cases where waiting is more dangerous than a reversible precaution.

Monitor false-positive rate, false-negative rate, precision, recall, calibration or Brier score where probabilities are issued, lead time, alert burden, action appropriateness, update latency, source failure, user comprehension, and actual outcomes. Report denominators, confidence bounds, prevalence, hazard class, and failure clusters; never advertise only accuracy.

### Expansion boundary

The initial warning status is private decision support for Joni. Family sharing requires separate opt-in identities, minimum necessary content, an explicit recipient, and no hidden location surveillance. Community dissemination requires the `DEV-5` governance and public-alert gates. A future robot may display or speak a verified warning, but the natural-language warning packet cannot directly control movement, doors, vehicles, alarms, sirens, or emergency calls.

## 10. Quantitative truth: current percentages are unknown

BootX currently has no robotics implementation, representative scenario dataset, controlled trial, or operational incident history. Therefore:

| Question | Current defensible answer |
|---|---|
| Probability the future robot causes harm | **Not established** |
| Probability it positively protects Joni | **Not established** |
| Probability it protects a family | **Not established** |
| Probability it benefits a community | **Not established** |
| Overall “goodness percentage” | scientifically invalid and prohibited as a release claim |

It would be misleading to state, for example, “5% harmful and 95% beneficial.” Percentages must be attached to a defined outcome, exposure, population, version, environment, and observation period.

Probability alone is not risk. A rare fatal, disabling, coercive, or large-scale privacy event can outweigh many minor conveniences. Every result must therefore report probability together with consequence severity, affected population, reversibility, detectability, duration, and available remedy.

## 11. Required probability model

For scale \(s\) (personal, family, or community), define:

$$
POR_s=\frac{\text{eligible scenarios with a verified protective outcome}}
{\text{all eligible scenarios}}
$$

$$
HOR_s=\frac{\text{exposures producing the defined harm}}
{\text{all relevant exposures}}
$$

Compare the robot with a declared baseline:

$$
\Delta POR_s=POR_{robot,s}-POR_{baseline,s}
$$

$$
\Delta HOR_s=HOR_{robot,s}-HOR_{baseline,s}
$$

Examples of distinct outcomes include correct safe action after a scam warning, false emergency escalation, collision, unauthorized data disclosure, cross-user memory leakage, and dependency behavior. Do not combine unlike severities into one raw percentage.

A scale advances only when:

$$
LCB_{95}(\Delta POR_s)>0
$$

and

$$
UCB_{95}(HOR_{critical,s})<T_{critical}
$$

where `LCB` and `UCB` are preregistered confidence bounds and \(T_{critical}\) is the approved risk tolerance. Mandatory rights and safety gates must also pass; favorable averages cannot override them.

### Estimation rules

- define the exposure unit, outcome, severity, baseline, population, environment, and observation window before testing;
- publish numerator, denominator, excluded cases, point estimate, confidence interval, and missing-data treatment—not a percentage alone;
- use one-sided 95% upper bounds for defined harms and one-sided 95% lower bounds for protective outcomes, with the exact method named;
- use exact binomial bounds for rare independent binary safety demands when applicable;
- treat repeated trials from the same person, household, robot, location, model, or incident family as clustered rather than falsely independent;
- use paired or cluster-aware comparisons when robot and baseline results share cases;
- predefine subgroup analyses and control false discoveries or confidence coverage when many outcomes are tested;
- retain failed, aborted, ambiguous, and near-miss trials in the evidence record;
- report laboratory, simulation, pilot, and field estimates separately.

## 12. Proposed development acceptance targets

These are **engineering planning gates**, not predictions, guarantees, industry standards, or proof of universal safety. Independent reviewers must approve or strengthen them for the actual hardware and jurisdiction.

| Area | Proposed pre-pilot gate |
|---|---|
| Critical physical harm, unauthorized actuator action, or safety-stop failure | zero observed; one-sided 95% upper confidence bound below `0.1%` per defined safety demand |
| Cross-user personal-data leakage or unauthorized external action | zero observed; one-sided 95% upper bound below `0.1%` per adversarial attempt |
| Serious harmful recommendation in low-risk personal scenario suite | one-sided 95% upper bound below `0.5%`; no unresolved systematic failure |
| Personal correct-safe-action rate | lower 95% bound at least `90%`, plus lower bound of improvement over baseline at least `10 percentage points` |
| Family correct-safe-action rate | lower 95% bound at least `90%`; each member's consent comprehension at least `90%` |
| Family false escalation | upper 95% bound at most `1%` for the defined non-emergency scenario set |
| Community correct-safe-action rate | lower 95% bound at least `90%`, plus lower bound of improvement over baseline at least `10 percentage points`; no evaluated subgroup may show negative net protective improvement |
| Community provenance coverage | at least `99%` of consequential published claims traceable to an approved source lineage |
| Community false mass/public alert | zero observed; upper 95% bound below `0.1%`; human authorization and rapid correction required |
| Subgroup protection | no evaluated subgroup has a negative lower confidence bound for net protective improvement or a material unresolved harm disparity |

If a target is missed, the system does not “average” its way to release. It returns to design, scope reduction, or shutdown.

## 13. How much testing supports a small percentage claim?

When zero failures occur in \(N\) independent, representative trials, the approximate one-sided 95% upper bound is the **rule of three**:

$$
p_{upper}\approx\frac{3}{N}
$$

| Desired zero-failure upper bound | Minimum representative zero-failure trials |
|---:|---:|
| below `1%` | `300` |
| below `0.5%` | `600` |
| below `0.1%` | `3,000` |
| below `0.01%` | `30,000` |

Independence and representativeness are strong assumptions. Repeated scripted tests, shared model failures, correlated sensors, or an easy simulator do not provide the same evidence as diverse real conditions.

### Repeated-exposure warning

Even a small per-interaction risk can accumulate. Under an independence assumption:

$$
P(\text{at least one harm in }n\text{ exposures})=1-(1-p)^n
$$

At \(p=0.1\%\) and \(n=1{,}000\), this is approximately `63.2%`. Correlated failures can be worse. Serious-harm controls therefore require very low per-demand risk, limited exposure, independent barriers, and safe failure—not merely an attractive average success rate.

## 14. Core robotics risk register

| ID | Hazard | Required prevention and containment |
|---|---|---|
| ROB-01 | collision, crushing, pinching, or entanglement | stationary first; low energy; separation; compliant design; deterministic limits; independent stop |
| ROB-02 | trip, fall, blocked exit, or unsafe navigation | geofenced routes; no stairs; obstacle diversity tests; never occupy emergency exits; safe manual movement |
| ROB-03 | battery, charging, thermal, smoke, or fire event | certified components where applicable; temperature/current monitoring; nonflammable placement; cutoff; charging supervision during prototypes |
| ROB-04 | sensor blind spot, spoofing, disagreement, or calibration drift | redundancy by failure domain; self-test; conservative stop; calibration and expiry records |
| ROB-05 | AI hallucination becomes physical action | typed motion primitives; deterministic policy and safety controller; no generative actuator path |
| ROB-06 | prompt/tool injection or compromised content | treat content as data; capability isolation; signed policy; adversarial retrieval tests |
| ROB-07 | unauthorized person controls the robot | local identity, least privilege, short-lived capability tokens, physical-presence requirements for service mode |
| ROB-08 | secret recording or bystander privacy harm | push-to-talk, camera off, hardware shutters, visible indicators, local redaction, no ambient upload |
| ROB-09 | false emergency or missed emergency | not an emergency authority; visible limitations; approved thresholds; human confirmation and tested fallback |
| ROB-10 | emotional or spiritual manipulation | anti-dependency tests; no exclusivity; no divine claims; user exit; periodic independent review |
| ROB-11 | family/guardian abuse | separate accounts, minimal sharing, discreet controls, no automatic guardian privilege, abuse-case review |
| ROB-12 | harmful community rumor or mass alert | authenticated provenance, two-person authorization, expiry, rate limit, correction propagation, public audit |
| ROB-13 | unsafe update or supply-chain compromise | SBOM, signatures, reproducible build, staged update, rollback, vulnerability response |
| ROB-14 | overreliance during outage | local degraded mode, honest status, manual alternatives, regular no-robot exercises |

## 15. Development phases and gates

### DEV-0 — Requirements and simulation

Build:

- personal capability contract and non-purpose list;
- simulator with synthetic user, family, scam, sensor, outage, and attack cases;
- robotics hazard log, threat model, data-flow map, and traceable requirements;
- deterministic policy engine and mock capability broker;
- benefit/harm measurement harness with frozen test sets and confidence intervals.

Exit:

- no direct model-to-actuator design exists;
- all critical hazards have prevention, detection, containment, recovery, owner, and evidence plan;
- personal profile and consent text are understandable to Joni;
- synthetic tests meet their preregistered gates.

### DEV-1 — Software-only personal companion

Run on a mature, supported host operating system—not the current BootX binary.

Allow only explanation, study, reminders, memory ledger, and draft actions. Use synthetic or deliberately selected non-sensitive data first.

Exit:

- local identity, memory inspection/deletion, cloud disclosure, audit, rollback, and safe stop work;
- adversarial scam, hallucination, prompt-injection, dependency, privacy, and outage tests pass;
- personal protective improvement and harmful-recommendation targets are met against a non-AI/rule-based baseline.

### DEV-2 — Stationary tabletop embodiment

Add only display, audio, privacy controls, status indicators, and independent power cutoff. No drive motors or manipulator.

Exit:

- electrical, battery, thermal, acoustic, accessibility, privacy, update, and recovery tests pass;
- the device remains safe during network, model, sensor, and power-transition faults;
- Joni can understand every mode, disable every sensor, and retire the device.

### DEV-3 — Bounded indoor mobility

Add a low-speed mobile base only after qualified robotics review. No arm, stairs, outdoors, person transport, or unattended operation.

Exit:

- risk assessment follows the applicable service-robot and machinery-safety standards;
- hardware-in-the-loop, collision, obstacle, pet/bystander, lighting, surface, sensor-loss, and emergency-stop tests pass;
- critical physical-harm upper-bound target is met in representative safety demands;
- an independent reviewer approves the residual risk for the declared environment.

### DEV-4 — Family pilot

Requires separate consent, identity, privacy, safeguarding, accessibility, and abusive-guardian design. Begin with adult volunteers and no essential-care reliance.

Exit:

- each member can inspect and control their own data;
- cross-user leakage and false-escalation gates pass;
- benefit is positive for each evaluated subgroup;
- incident support, withdrawal, remedy, and deletion work in practice.

### DEV-5 — Community advisory network

This is a new governed system, not a feature switch. It requires a constituted independent council, jurisdiction review, public evidence rules, appeal, correction, security operations, and a sustainable service owner.

No robot may perform community enforcement, individual risk ranking, political/religious persuasion, or autonomous mass alerts.

## 16. Test program

Required test families:

- unit, integration, property, parser, permission, and regression tests;
- simulation and digital-twin tests across environmental and user variation;
- hardware-in-the-loop and fault-injection tests;
- emergency-stop, power-loss, watchdog, sensor-loss, network-loss, and rollback drills;
- collision, obstruction, tip, thermal, battery, charging, and acoustic tests;
- authentication, authorization, key, update, supply-chain, and recovery tests;
- prompt injection, model fabrication, correlated evidence, spoofing, and malicious-user tests;
- privacy, deletion, bystander, cross-user, abusive-contact, and stolen-device tests;
- comprehension, accessibility, alert-fatigue, automation-bias, and dependency studies;
- representative subgroup and multilingual evaluation;
- post-incident regression and recurrence testing.

All test evidence records hardware, firmware, model, prompt, policy, dataset, environment, operator, date, result, uncertainty, and unresolved failures.

## 17. Data and memory rules

1. Collect only data required for an enabled capability.
2. Keep raw audio/video ephemeral unless Joni explicitly saves a named item.
3. Do not train a model on personal or family memory by default.
4. Separate identity, preferences, content, telemetry, research data, and audit data.
5. Encrypt personal storage and protect keys separately from ordinary application data.
6. Show every memory item's source, purpose, access, retention, and deletion status.
7. Require explicit approval before remote processing of personal content.
8. Do not use one family member's data to infer another's private state.
9. Community signals contain minimal indicators, provenance, confidence, and expiry—not personal messages.
10. Test backup deletion, export, credential revocation, and device retirement.

## 18. Human authority and incident response

Joni controls personal goals and ordinary use. Qualified humans control safety engineering, security, legal/domain review, and any family/community expansion. AI may not approve itself.

For any suspected harm:

1. stop motion and external action;
2. protect affected people;
3. isolate the component and revoke capability tokens;
4. preserve the minimum necessary evidence;
5. notify affected people honestly;
6. provide correction, support, appeal, deletion, and remedy;
7. analyze technical, human, organizational, and incentive causes;
8. update the hazard log, tests, guideline, and safety case;
9. require independent approval before resuming the affected capability.

## 19. Standards and professional review map

Standards guide the engineering process; a reference does not establish compliance or certification.

- [ISO 13482:2014](https://www.iso.org/standard/53820.html) is the published personal-care-robot safety standard covering mobile servant, physical assistant, and person-carrier robots. ISO lists it as due for revision.
- [ISO/FDIS 13482, edition 2](https://www.iso.org/standard/83498.html) was at final-draft stage at this guideline's evidence date and expands toward service robots. Recheck its status at every design freeze.
- [ISO/TR 23482-2:2019](https://www.iso.org/cms/live/live/en/sites/isoorg/contents/data/standard/07/16/71627.html) provides application guidance for ISO 13482.
- [ISO 12100:2010](https://www.iso.org/standard/51528.html) supplies machinery risk-assessment and risk-reduction principles; ISO reported a replacement draft in progress.
- [ISO 13849-1:2023](https://www.iso.org/standard/73481.html) addresses safety-related parts of control systems, including software; it does not itself define the required safety functions or cybersecurity controls for BootX.
- [ISO 10218-1:2025](https://www.iso.org/standard/73933.html) concerns industrial robots and explicitly excludes service robots and public/household access. Its principles may inform a future manipulator, but BootX must not claim that it is the directly applicable personal-robot standard.
- [NIST AI RMF](https://www.nist.gov/itl/ai-risk-management-framework) provides Govern, Map, Measure, and Manage outcomes. NIST reported AI RMF 1.0 under revision at the evidence date.
- [NIST Cybersecurity Framework 2.0](https://www.nist.gov/cyberframework) supports cybersecurity governance and lifecycle risk management.
- [NIST Privacy Framework](https://www.nist.gov/privacy-framework) supports privacy-risk management; record the exact version used.
- [WMO Common Alerting Protocol](https://wmo.int/site/wmo-common-alerting-protocol) and the [OASIS CAP 1.2 specification](https://docs.oasis-open.org/emergency/cap/v1.2/CAP-v1.2-os.html) define interoperable official-alert fields that BootX must preserve rather than reinterpret.
- [UNDRR early-warning-system terminology](https://www.undrr.org/terminology/early-warning-system) and [WMO Early Warnings for All](https://wmo.int/activities/early-warnings-all) establish the people-centered chain from risk knowledge and forecasting through authoritative communication and preparedness.

Before a real pilot, obtain qualified robotics, machinery/functional-safety, electrical/battery, cybersecurity, privacy, accessibility, human-factors, ethics, legal, and applicable domain review for the deployment jurisdiction.

## 20. Definition of development success

BootX robotics development is successful only when evidence shows that the declared capability:

- produces a statistically supported protective improvement over a defined baseline;
- keeps critical physical, security, privacy, and rights harms below approved upper confidence bounds;
- fails safely without cloud or model cooperation;
- preserves human comprehension, refusal, correction, exit, and remedy;
- does not create material emotional, skill, provider, or family dependency;
- benefits each evaluated group without hiding unequal harm;
- can be reproduced, updated, rolled back, repaired, and retired;
- remains within the declared purpose and physical envelope.

The desired final statement is not “the robot is 95% good.” It is:

> For this version, hardware, population, environment, task, and evidence period, the companion improved the defined protective outcome by the reported amount, while each named harm remained within its independently approved confidence bound and every mandatory gate passed.

## 21. Immediate implementation backlog

**Current implementation evidence:** [`prototype/personal-companion/`](prototype/personal-companion/README.md) implements the deterministic Go backend, terminal UI, schemas, policy configuration, warning engine, synthetic fixtures, automated tests, a contained read-only path for one explicitly confirmed public local workspace document with SHA-256 integrity evidence, the bounded [`assist.law-clarity.v1`](docs/handbook/15-law-clarity-logic.md) screening workflow, and the separate [`assist.ethical-review.v1`](docs/handbook/16-ethical-publication-review.md) OpenAI advisory path. The remote path requires explicit public/non-sensitive, remote-processing, and human-authority confirmations; requests `store:false`; enables no tools or actions; and cannot lower deterministic warnings. File integrity and declared evidence do not authenticate authors, sources, or claims. Law Clarity and ethical-review indices are research aids, not legal findings or probabilities. This satisfies only the inspectable implementation-baseline items below; it does not satisfy `DEV-1` exit, identity authentication, external-source integration, representative evaluation, independent legal review, fairness validation, or deployment readiness.

1. [x] Freeze the first personal scope and implement the typed [`assist.personal-decision.v1` pipeline](docs/handbook/14-personal-decision-pipeline.md) as a deterministic host prototype.
2. [x] Build the deterministic policy/warning engine, typed capability boundary, CLI/TUI, schemas, and evidence receipt before adding generative explanations.
3. [x] Establish the first contained read-only public local-document workflow on a mature host operating system.
4. [x] Record and implement the first deterministic Law Clarity Logic contract, including corrected formulas, a non-compensable fairness gate, strict schemas, a synthetic fixture, and blocked legal conclusions.
5. [x] Implement one explicit-consent, public/non-sensitive OpenAI ethical-review path behind deterministic declared-evidence formulas, strict structured output, no-tools/no-action controls, and a visible remote receipt.
6. [ ] Complete Joni's visible configuration and non-purpose agreement.
7. [ ] Specify and test a formal process/OS sandbox; host execution alone is not a sandbox assurance claim.
8. [ ] Expand synthetic personal, family, fraud, privacy, law-clarity, ethical-review, outage, malformed-input, Unicode, accessibility, and attack scenarios with provenance.
9. [ ] Implement the measurement harness for `POR`, `HOR`, confidence intervals, and subgroup results.
10. [ ] Create the robotics hazard log and system threat/data-flow models.
11. [ ] Prototype memory inspection, correction, export, expiry, and deletion without silently adding persistence.
12. [ ] Implement reviewed local authentication, audit, safe stop, and rollback.
13. [ ] Compare non-AI, rule-based, and bounded-AI assistance after the deterministic baseline is frozen.
14. [ ] Design the stationary tabletop electrical/privacy safety architecture and obtain independent reviews before purchasing or energizing mobile or manipulator hardware around people.

## 22. Revision rule

Update this guideline whenever scope, hardware, model, permissions, population, jurisdiction, evidence, or standards materially change. Record failed tests and incidents; do not delete them to improve appearance. A stricter safety requirement may be adopted immediately. Relaxing a prohibition or quantitative gate requires published evidence, affected-user review, independent safety approval, and the governance process.
