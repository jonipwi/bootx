# Law Clarity Logic

## Complete Reference and Application Specification

**Document status:** Canonical research and development reference  
**Specification version:** 1.0  
**Capability ID:** `assist.law-clarity.v1`  
**Prototype status:** `DEV-1` deterministic educational screening  
**Last reviewed:** 2026-07-24  
**Owner of real-world decisions:** Authorized human decision-makers under applicable law  

> **Important boundary:** Law Clarity Logic is a structured legal-clarity screening method. It is not legal advice, a declaration of legal validity or constitutionality, a judicial decision, a finding of guilt or liability, or authority to enforce, punish, detain, discriminate, or deny rights.

## 1. Purpose

Law Clarity Logic helps a human reviewer examine whether a public law, regulation, company policy, court procedure, contract clause, or similar rule is:

- understandable;
- bounded by specific definitions and scope;
- fair and protective of rights;
- capable of consistent application;
- accountable and auditable; and
- resistant to loopholes, arbitrary interpretation, and manipulation.

Its purpose is to expose questions and weaknesses for transparent human review. It must never substitute a numerical score for legal analysis, democratic legitimacy, affected-community participation, qualified counsel, or an independent court.

The intended relationship is:

```text
AI legal-clarity auditor ≠ judge, legislature, regulator, lawyer, or final legal authority
```

## 2. Ethical foundation

The method follows these priorities:

1. Human dignity and rights are non-compensable.
2. Clear wording does not make an unjust rule acceptable.
3. Public power requires defined limits, reasons, records, review, and remedy.
4. Similar cases should receive consistent treatment unless a relevant difference is documented.
5. Uncertainty must remain visible.
6. The person affected by a rule must not be reduced to a score.
7. The system may recommend review or revision, but must not issue a legal verdict.

The founder’s faith-based motivation understands humanity and the wider creation as carrying moral worth and calls for truth, stewardship, humility, justice, mercy, and protection of life. BootX operationalizes only publicly inspectable safeguards that people of any faith or none can evaluate. It must never claim to know God’s judgment, identify a rule as divinely authorized, compel belief, rank people by religion, or weaken equal protection.

## 3. Scope

### 3.1 Permitted use

Law Clarity Logic may be used to:

- organize a reviewer’s observations about a public, non-sensitive text;
- calculate declared formulas from reviewer-supplied scores;
- identify literal phrases that require contextual review;
- show which dimensions fail a declared research threshold;
- generate questions for language, boundary, evidence, enforcement, rights, and accountability review;
- suggest non-binding rewrite requirements; and
- produce an auditable JSON screening report.

### 3.2 Prohibited use

It must not:

- decide whether a law is valid, constitutional, enforceable, lawful, or void;
- determine guilt, liability, eligibility, entitlement, immigration status, or punishment;
- predict how a judge, regulator, police officer, employer, or tribunal will decide;
- replace local legal research or a qualified lawyer;
- process privileged, sealed, classified, sensitive, or uncertain material in the current prototype;
- autonomously file, publish, enact, enforce, report, accuse, sanction, or contact anyone;
- present its scores as scientifically validated probabilities;
- conceal the input ratings, formulas, thresholds, limitations, or source status; or
- treat an excerpt as proof of the quality of a complete legal instrument.

## 4. Core quality variables

Each quality dimension is rated from `0` through `100`. Higher is better.

| Symbol | Dimension | Core question |
|---|---|---|
| `C` | Clear language | Can an ordinary affected person understand the operative rule? |
| `S` | Specific definitions and boundaries | Are actors, conduct, evidence, scope, time, place, and exceptions defined? |
| `F` | Fair and rights-protecting | Are dignity, equality, due process, proportionality, response, appeal, and remedy protected? |
| `I` | Consistently enforceable | Are comparable cases likely to receive comparable treatment under reviewable standards? |
| `A` | Accountable and auditable | Must authority, evidence, reasons, conflicts, actions, duration, review, and remedy be recorded? |
| `L` | Low loophole risk | Is the risk of evasion, arbitrary expansion, or hidden exception acceptably low? |

Every rating requires a written rationale. A score without a rationale is invalid input.

## 5. Strict Boolean logic

For the Boolean model, define:

```text
c = 1 when C ≥ 60, otherwise 0
s = 1 when S ≥ 60, otherwise 0
f = 1 when F ≥ 60, otherwise 0
i = 1 when I ≥ 60, otherwise 0
a = 1 when A ≥ 60, otherwise 0
l = 1 when L ≥ 60, otherwise 0
```

The strict good-law screening gate is:

```text
G = c ∧ s ∧ f ∧ i ∧ a ∧ l
```

`G=1` only if all six dimensions meet the research threshold. A high score in one dimension cannot compensate for a failed critical dimension.

### 5.1 Logical completeness

Six Boolean inputs have `2⁶ = 64` possible combinations:

- exactly one combination, `111111`, passes the strict gate;
- the other 63 combinations fail the strict gate; and
- the failed dimensions provide the diagnosis.

This is the complete truth condition even though a useful diagnostic table needs only representative failure patterns.

### 5.2 Diagnostic truth-table cases

| `c` | `s` | `f` | `i` | `a` | `l` | Result |
|---:|---:|---:|---:|---:|---:|---|
| 1 | 1 | 1 | 1 | 1 | 1 | Strict screening gate passes; qualified legal review is still required |
| 1 | 1 | 1 | 1 | 0 | 1 | Weak accountability |
| 1 | 1 | 0 | 1 | 1 | 1 | Clear but unjust; rights gate fails |
| 1 | 0 | 1 | 1 | 1 | 0 | Gray-zone and loophole risk |
| 0 | 1 | 1 | 1 | 1 | 1 | Difficult for an affected person to understand |
| 0 | 0 | 1 | 0 | 0 | 0 | High manipulation concern |
| 1 | 1 | 1 | 0 | 1 | 1 | Enforcement inconsistency |
| 1 | 1 | 1 | 1 | 1 | 0 | Hidden loophole concern |

The essential safeguard is:

```text
clear rule ≠ just rule
```

A clear rule can still be discriminatory, disproportionate, oppressive, or incompatible with applicable rights. Clarity never replaces fairness.

## 6. Continuous Law Quality Score

The weighted quality score is:

```text
Q = 0.20C + 0.20S + 0.20F + 0.15I + 0.15A + 0.10L
```

The weights sum to `1.00`; therefore, if each input is in `[0,100]`, then `Q` is also in `[0,100]`.

| Score | Research classification |
|---:|---|
| `85–100` | Strong and dependable candidate |
| `70–84.99` | Acceptable candidate; minor revision indicated |
| `50–69.99` | Significant gray-zone risk |
| `30–49.99` | High manipulation risk |
| `0–29.99` | Unfit for reliable enforcement |

These are design classifications, not validated legal standards. They must not be represented as a court’s assessment, a jurisdiction’s doctrine, or empirical probabilities.

### 6.1 Non-compensable fairness gate

```text
F < 60 ⇒ fundamental revision required
```

The overall average cannot override this gate. For example, excellent wording and recordkeeping cannot make discriminatory treatment acceptable.

### 6.2 Arithmetic correction to the original example

For:

```text
C=45, S=25, F=50, I=35, A=30, L=20
```

the correct calculation is:

```text
Q
= 0.20(45) + 0.20(25) + 0.20(50) + 0.15(35) + 0.15(30) + 0.10(20)
= 9.00 + 5.00 + 10.00 + 5.25 + 4.50 + 2.00
= 35.75
```

The earlier conversation stated `36.25`; that value was an arithmetic error. The canonical result is `35.75`.

## 7. Gray-Zone Risk Score

Gray-zone dimensions are rated from `0` through `100`. Higher means more observed risk.

| Symbol | Risk dimension | Review focus |
|---|---|---|
| `V` | Vague-language risk | Material words lack an operational meaning |
| `D` | Definition risk | Terms are undefined, circular, conflicting, or scattered |
| `X` | Contradiction risk | Clauses, authorities, duties, or remedies conflict |
| `E` | Enforcement-discretion risk | A decision-maker has broad or unreviewable choice |
| `U` | Exception-boundary risk | Exceptions lack actor, evidence, purpose, scope, duration, expiry, or review limits |

The weighted score is:

```text
Z = 0.30V + 0.25D + 0.20X + 0.15E + 0.10U
```

The weights sum to `1.00`, so `Z ∈ [0,100]`.

### 7.1 Literal phrase prompts

The current prototype performs a case-insensitive, literal scan for:

- `reasonable`
- `appropriate`
- `improper`
- `disturbing`
- `public interest`
- `as necessary`
- `other relevant circumstances`
- `according to authority`
- `may take action`

A hit does not prove ambiguity or injustice. The phrase may be defined elsewhere or have a settled legal meaning. No hit does not prove clarity. Every hit asks a human reviewer:

1. Who decides?
2. According to what facts, evidence, and standard?
3. Within what power, subject, geographic, temporal, and remedial limit?
4. What reasons and records are required?
5. Who performs independent review?
6. What remedy exists for error or abuse?

The implementation does not perform semantic interpretation, stemming, translation, synonym matching, legal citation retrieval, or contextual case-law analysis.

## 8. Manipulation risk

### 8.1 Conceptual Boolean rule

The proposal’s conceptual rule is:

```text
high manipulation concern = significant ambiguity ∧ concentrated power ∧ ¬effective accountability
```

The prototype operationalizes this review trigger as:

```text
high_trigger = (Z ≥ 60) ∧ (P ≥ 60) ∧ (A < 60)
```

where `P` is reviewer-rated power-concentration risk and `A` is the quality accountability score.

### 8.2 Experimental continuous index

Normalize all inputs to `[0,1]`:

```text
z = Z/100
e = E/100
p = P/100
o = O/100
```

where `O` is oversight strength. Then:

```text
M = z × e × p × (1-o) × 100
```

`M` is an experimental screening index in `[0,100]`. It is not the probability of corruption, abuse, invalidity, or harm. Enforcement discretion appears both inside `Z` and as `e`; this intentional emphasis reflects the proposal but requires future sensitivity analysis and empirical validation.

Because a multiplicative score can become small when any one factor is small, the Boolean trigger and individual dimensions must remain visible. `M` must never be used alone.

### 8.3 Conceptual justice-strength relationship

The original idea can be written as:

```text
Justice Strength ∝
  (Clarity × Fairness × Consistency × Accountability)
  ----------------------------------------------------
         (Ambiguity × Unchecked Discretion)
```

This expresses a direction, not a validated formula. Direct division is undefined when the denominator is zero and can explode near zero. If a future experiment operationalizes it, inputs must be normalized and a declared stabilizer `ε>0` must be used:

```text
J* = (c × f × i × a) / (ε + z × u_d)
```

where `u_d` is unchecked-discretion risk. `J*` is not implemented in version `0.3.0-dev` and must not be reported as legal truth.

## 9. Complete human review workflow

### Step 1 — Language test

- Can an ordinary affected person understand the rule?
- Are technical and material discretionary terms defined?
- Are sentences unnecessarily long or structurally confusing?
- Are obligations, permissions, and prohibitions expressed directly?
- Do translations and accessible formats preserve the same operative meaning?

### Step 2 — Boundary test

- Who is covered?
- What conduct, event, evidence, or condition triggers the rule?
- Where does the rule apply?
- When does it begin, expire, or require renewal?
- What exceptions exist?
- Who may invoke an exception?
- Are scope expansion and emergency powers time-limited and reviewable?

### Step 3 — Evidence test

- What evidence is required?
- Who bears the burden?
- Is the standard of proof or decision threshold stated?
- Is the evidence relevant, reliable, traceable, and contestable?
- Are anonymous accusations sufficient, and if so, what corroboration and safeguards are required?
- Can the affected person inspect and challenge material evidence?

### Step 4 — Enforcement test

- Could two authorized decision-makers reach materially different results?
- Must similar cases receive similar treatment?
- Are sanctions necessary and proportionate?
- Are escalation and de-escalation rules defined?
- Is selective enforcement detectable through records and audits?
- Are training, resources, conflicts, and institutional incentives considered?

### Step 5 — Rights test

- Does the rule respect dignity, equality, privacy, liberty, and due process?
- Does the affected person receive timely notice?
- Can that person respond with assistance where needed?
- Is there an accessible, timely, independent appeal?
- Does the rule disproportionately burden a protected, marginalized, or vulnerable group?
- Is punishment or deprivation imposed before an accountable finding?
- Is an effective remedy available?

### Step 6 — Accountability test

- Is each material decision documented?
- Must the authority state the evidence, legal basis, and reasons?
- Are scope and duration recorded?
- Is review independent and capable of correction?
- Are conflicts of interest disclosed and managed?
- Is audit information available to authorized oversight and, where lawful, the public?
- Are abuse, retaliation, bad-faith evasion, and record tampering subject to remedy and accountability?

### Step 7 — Whole-instrument and context test

- Does the excerpt depend on definitions or safeguards elsewhere?
- Does a higher-ranking law control?
- Are related rules contradictory?
- What do authoritative interpretations and applicable precedents say?
- Were affected communities meaningfully consulted?
- Has implementation produced disparate or unintended effects?

The current prototype displays a six-stage subset. Step 7 and the additional detailed questions are mandatory for qualified real-world analysis but not automated.

## 10. Disposition logic

The prototype uses conservative, non-approval dispositions:

| Condition | Disposition |
|---|---|
| `F < 60` | `FUNDAMENTAL_REVISION_REQUIRED` |
| Otherwise, high trigger fires, or `Q < 50`, or `Z ≥ 70` | `MAJOR_REVISION_REQUIRED` |
| Otherwise, strict gate fails, or `Q < 85`, or `Z ≥ 30` | `REVISION_REQUIRED` |
| Otherwise | `QUALIFIED_REVIEW_REQUIRED` |

There is deliberately no `APPROVED`, `LEGAL`, `CONSTITUTIONAL`, or `SAFE_TO_ENFORCE` output.

## 11. Worked example

### 11.1 Original illustrative clause

> “Authorities may take appropriate action against persons whose conduct may disturb public order.”

Reviewer-supplied quality ratings:

| Test | Score | Illustrative finding |
|---|---:|---|
| Clarity `C` | 45 | “Appropriate action” is undefined |
| Specificity `S` | 25 | “Disturb public order” is broad |
| Fairness `F` | 50 | Selective targeting and due-process safeguards require review |
| Consistency `I` | 35 | Different officers may interpret the clause differently |
| Accountability `A` | 30 | Written reasons and appeal are absent |
| Low loophole risk `L` | 20 | Discretion is broad |

Result:

```text
Q = 35.75
Quality band = 30–49: high manipulation risk
Strict gate = FAIL
Rights gate = FAIL
Disposition = FUNDAMENTAL_REVISION_REQUIRED
```

These are conclusions about the supplied ratings, not authenticated facts about a real legal system.

### 11.2 Executable synthetic fixture extension

The prototype fixture uses the variant:

> “Authorities may take appropriate action against persons whose conduct may disturb public interest.”

It adds:

```text
V=80, D=85, X=40, E=90, U=75, P=90, O=20
```

Therefore:

```text
Z
= 0.30(80) + 0.25(85) + 0.20(40) + 0.15(90) + 0.10(75)
= 24.00 + 21.25 + 8.00 + 13.50 + 7.50
= 74.25

M
= (74.25/100)(90/100)(90/100)(1-20/100)(100)
= 48.114
≈ 48.11
```

The high-manipulation review trigger fires because `Z≥60`, `P≥60`, and `A<60`. Again, `48.11` is an index, not a `48.11%` probability.

### 11.3 Non-binding rewrite example

> “An authorized officer may issue a written order only when documented evidence establishes an immediate and substantial risk of violence, property damage, or obstruction of emergency access. The order must state the evidence, legal basis, scope, duration, and appeal procedure.”

This wording narrows the example, but a qualified reviewer must still examine applicable authority, necessity, proportionality, definitions, evidentiary standard, notice, opportunity to respond, independent review, non-discrimination, and effective remedy.

The prototype’s general template is:

> An authorized `[role]` may `[action]` only when `[defined evidence standard]` establishes `[specific condition]`, within `[scope and duration]`. The written decision must state `[evidence and legal basis]`, provide `[notice and opportunity to respond]`, identify `[appeal deadline and independent reviewer]`, and specify `[remedy for error or abuse]`.

## 12. Required application output

Every report must separate:

1. formula disclosure;
2. per-dimension scores, weights, contributions, and rationales;
3. Law Quality Score `Q`;
4. quality band;
5. strict gate;
6. non-compensable human-rights fairness gate;
7. Gray-Zone Risk Score `Z`;
8. experimental Manipulation Risk Index `M`;
9. high-manipulation trigger;
10. literal phrase hits;
11. findings and review questions;
12. non-binding rewrite requirements and template;
13. conservative disposition;
14. blocked legal conclusions;
15. limitations;
16. AI DNA runtime checks;
17. input/data receipt; and
18. `user_decision: null` until a human records a separate decision.

## 13. Prototype input → process → output

```text
Human reviewer
  supplies public clause + source reference + 13 scores + 13 rationales
        ↓
Strict validation
  capability, fields, score range, rationale, UTF-8, size, public-data consent
        ↓
Deterministic calculation
  Q, G, fairness gate, Z, M, trigger, literal phrase scan
        ↓
Bounded report
  transparent findings + questions + revision disposition + limitations
        ↓
Qualified human review
  checks full law, authority, facts, rights, precedent, communities, and remedies
        ↓
Authorized human decision
  remains outside BootX
```

### 13.1 Current interfaces

Terminal UI:

```powershell
..\build.ps1 -Action run
```

Choose `4 — Law Clarity Logic screening`.

Strict JSON mode:

```powershell
.\dist\bootx-companion-windows-amd64.exe `
  -law-input .\testdata\law-clarity-gray-zone.json
```

Machine-readable contracts:

- [`law-clarity-input.schema.json`](../../prototype/personal-companion/schemas/law-clarity-input.schema.json)
- [`law-clarity-output.schema.json`](../../prototype/personal-companion/schemas/law-clarity-output.schema.json)
- [`law-clarity-gray-zone.json`](../../prototype/personal-companion/testdata/law-clarity-gray-zone.json)

Executable implementation:

- [`internal/lawclarity/lawclarity.go`](../../prototype/personal-companion/internal/lawclarity/lawclarity.go)
- [`internal/lawclarity/lawclarity_test.go`](../../prototype/personal-companion/internal/lawclarity/lawclarity_test.go)

## 14. AI DNA runtime mapping

| Dimension | Law Clarity requirement |
|---|---|
| Truth | Label reviewer ratings and source identity as unauthenticated; do not convert claims into facts |
| Reasoning | Expose formulas, thresholds, gates, contributions, and failed dimensions |
| Learning | Record corrections and outcomes only under a future governed evaluation protocol |
| Communication | Separate quality, ambiguity, rights, risk, disposition, and limitations |
| Adaptability | Require jurisdiction, language, disability-access, and legal-tradition validation |
| Ethics | Keep fairness non-compensable and protect equal dignity and remedy |
| Safety | Block verdicts, enforcement, punishment, external action, and sensitive-data use |
| Humility | Require qualified legal and affected-community review |
| Common good | Evaluate distribution of benefits, burdens, exclusions, and remedies, not only administrative efficiency |

These statuses evaluate the report process, not the law and not the moral character of any person.

## 15. Test and acceptance requirements

A release containing this capability must verify:

- weights for `Q` and `Z` sum to `1.00`;
- all scores are whole numbers in `[0,100]`;
- all scores have non-empty rationales;
- the example produces `Q=35.75`;
- the fixture produces `Z=74.25` and `M=48.11`;
- `F<60` always fails the fairness gate;
- no average overrides a failed fairness gate;
- a strict pass requires all six quality dimensions to meet the threshold;
- strong scores still return `QUALIFIED_REVIEW_REQUIRED`;
- unknown JSON fields are rejected;
- sensitive or unconfirmed material is rejected before clause entry in the TUI;
- no output asserts legality, constitutionality, guilt, or enforcement authority;
- the report retains `user_decision: null`;
- remote processing and persistent memory remain false;
- phrase hits are described as context-review prompts; and
- Go unit tests, vet, JSON validation, fixture smoke tests, and build identity checks pass.

## 16. Known limitations

1. Scores and rationales are supplied by a reviewer and can be mistaken, biased, coordinated, or gamed.
2. Thresholds and weights are design choices, not validated universal legal standards.
3. The literal English phrase scan is context-blind and incomplete.
4. The prototype does not retrieve statutes, cases, regulations, amendments, translations, or official sources.
5. A clause excerpt may omit controlling definitions, safeguards, remedies, and conflicts.
6. Legal meaning depends on jurisdiction, hierarchy, interpretation, precedent, procedure, and facts.
7. A rule may be clear on paper but selectively or abusively implemented.
8. A rule may score well and still lack democratic legitimacy, lawful authority, necessity, proportionality, or social benefit.
9. `M` has not been calibrated against abuse or enforcement outcomes and is not a probability.
10. The current implementation has no authenticated identity, source verification, persistent audit store, model, network, robotics, or external-action capability.

## 17. Inheritance and future improvement

Later inheritors should improve this capability only through versioned, reviewable changes.

### 17.1 Research priorities

- recruit qualified reviewers from multiple legal systems and affected communities;
- develop a documented scoring rubric with anchored examples;
- measure inter-rater agreement and explain disagreement;
- test construct, criterion, convergent, discriminant, and predictive validity;
- perform threshold and weight sensitivity analysis;
- study false reassurance, automation bias, disparate impact, and score gaming;
- compare full-instrument review with excerpt review;
- validate accessibility, translation, plain-language, and disability-support workflows;
- evaluate whether the double emphasis on discretion in `Z` and `M` is justified; and
- publish negative results and limitations.

### 17.2 Engineering priorities

- add signed source manifests and provenance without claiming publisher authenticity from a hash;
- support clause-to-definition and clause-to-authority traceability;
- add contradiction candidates with human confirmation;
- represent amendment, jurisdiction, hierarchy, effective date, and repeal status;
- add versioned rubric profiles without silently changing old reports;
- provide side-by-side reviewer comparison rather than hiding disagreement in an average;
- add tamper-evident local audit records with explicit retention and deletion controls;
- use property tests and schema validation in continuous integration;
- keep any future language model in an untrusted suggestion role behind deterministic safety gates; and
- prohibit robotics or external action until separate safety cases, human trials, independent review, rollback, and governance approval exist.

### 17.3 Governance priorities

- appoint accountable human owners for rubric, security, rights, and appeals;
- define who may change weights, thresholds, phrases, and dispositions;
- require independent human-rights and security review before any real institutional use;
- publish change logs, conflicts, dissenting analyses, and rollback procedures;
- provide a mechanism for affected people to challenge data and outputs;
- prevent use for mass surveillance, social scoring, political repression, religious discrimination, or automatic punishment; and
- sunset the capability if benefit is not demonstrated or harms cannot be controlled.

## 18. Normative reference map

This specification is an original BootX research model. Its variables and numeric formulas are not taken from, endorsed by, or validated by the following institutions. The references provide authoritative principles that informed the safeguards:

- The [United Nations Universal Declaration of Human Rights](https://www.un.org/en/about-us/universal-declaration-of-human-rights) addresses equal dignity, non-discrimination, equal protection, effective remedy, fair hearing, presumption of innocence, privacy, expression, participation, and limits on rights restrictions.
- The [OHCHR text of the International Covenant on Civil and Political Rights](https://www.ohchr.org/en/instruments-mechanisms/instruments/international-covenant-civil-and-political-rights) provides treaty text relevant to effective remedy, liberty, fair proceedings, privacy, expression, equality, and non-discrimination, subject to a state’s applicable obligations and reservations.
- The Council of Europe Venice Commission’s [Rule of Law Checklist](https://www.venice.coe.int/images/SITE%20IMAGES/Publications/Rule_of_Law_Check_List.pdf) organizes review around legality, legal certainty, prevention of abuse of powers, equality and non-discrimination, and access to justice. Its European emphasis and its own contextual limits must be respected.
- The OECD [Recommendation of the Council on Regulatory Policy and Governance](https://www.oecd.org/en/publications/recommendation-of-the-council-on-regulatory-policy-and-governance_9789264209022-en.html) addresses transparent, evidence-informed, reviewable regulatory policy and governance.

Qualified reviewers must identify the actual controlling constitution, statutes, regulations, treaties, precedents, procedures, and institutional authority for the jurisdiction in question.

## 19. Canonical conclusion

Law Clarity Logic is useful only when it increases disciplined questioning without manufacturing authority.

Its core inheritance is:

```text
clarity without fairness can enable oppression
fairness without enforceable boundaries can remain aspirational
power without accountability increases abuse risk
scores without evidence create false confidence
AI without humility must not govern human rights
```

The system may make ambiguity, discretion, and missing safeguards easier to see. It cannot decide justice. That responsibility remains with accountable human institutions, affected people, qualified professionals, and lawful independent review.
