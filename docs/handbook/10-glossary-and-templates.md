# Glossary and Study Templates

## 1. Controlled glossary

| Term | Operational meaning in BootX |
|---|---|
| AI companion | bounded software assistant that helps a user understand, protect, remember, and learn; not a claim of personhood |
| AI DNA | versioned nine-dimension behavior and assurance specification; not biology or a human-worth score |
| Agency | practical ability to understand, choose, refuse, override, appeal, exit, and act |
| Calibration | agreement between stated probability/confidence and observed frequency/correctness |
| Capability | an operation the system can perform through a defined interface and permission |
| Common good | conditions supporting shared dignity, safety, justice, resilience, and future well-being; requires plural governance |
| Companion memory | purpose-bound, user-governed stored context with provenance, access, retention, correction, and deletion |
| Evidence maturity | E0–E5 label describing the strength and context of support for a claim |
| Explainability | communication enabling an affected person or reviewer to understand relevant reasons, limits, and recourse |
| Human in command | accountable human authority defines scope and can review, stop, and remedy; stronger than a ceremonial approval click |
| Local-first | prefer local control/processing where it meaningfully improves privacy, resilience, or ownership, while allowing bounded remote services |
| Protection | proportionate reduction of credible harm while preserving rights and recourse |
| Provenance | traceable origin and transformation history of data, evidence, model, policy, or artifact |
| Risk | combination of plausible harm, likelihood, exposure, vulnerability, scale, reversibility, and uncertainty |
| Safety case | structured claim–argument–evidence record showing why bounded use is acceptably safe |
| Trust | context-specific willingness to accept vulnerability based on evidence of competence, integrity, transparency, and recourse |
| Trustworthiness | properties and demonstrated conduct that can justify appropriate trust |
| Uncertainty | limits arising from randomness, missing knowledge, measurement, model assumptions, or changing context |
| Vulnerable user | a person whose context may increase exposure, impact, or barriers to recourse; not a fixed label of lesser agency |

## 2. Claim record template

```markdown
# Claim ID and title

- Type: fact | forecast | causal | recommendation | normative
- Owner:
- Version/date:
- Scope/population/context:
- Evidence maturity and expiry:

## Claim

## Evidence and provenance

## Assumptions

## Alternatives and counterevidence

## Uncertainty and sensitivity

## Consequence if wrong

## Verification, correction, and recourse
```

## 3. Research protocol template

```markdown
# Study title and registration

## Research question and hypothesis
## Background evidence and gap
## Population, recruitment, inclusion, and exclusion
## Intervention and comparator
## Assignment and blinding where applicable
## Primary and secondary outcomes
## Harm measures and stop rules
## Sample-size rationale
## Data collection, consent, privacy, retention, and deletion
## Analysis, uncertainty, subgroup, and missing-data plan
## Accessibility and participant support
## Conflicts, funding, ethics review, and accountable owners
## Deviations and publication commitment
```

## 4. Capability contract template

```yaml
capability_id:
version:
owner:
intended_users:
purpose:
non_purpose: []
inputs: []
outputs: []
permissions: []
risk_tier:
human_controls: []
data_policy:
  purpose:
  location:
  retention:
  sharing:
  deletion:
safety_controls: []
measures: []
stop_conditions: []
evidence_maturity:
review_expiry:
```

## 5. Decision log template

```markdown
# Decision ID and title

- Date/status:
- Decision owner:
- Participants and affected representatives:
- AI DNA version:

## Problem and scope
## Options, including no action
## Evidence and uncertainty
## Rights, harms, benefits, and distribution
## Security/privacy/accessibility analysis
## Decision and reasoning
## Dissenting views
## Conditions, measures, stop rules, and expiry
## Review outcome and corrections
```

## 6. Memory item schema

```yaml
memory_id:
user_visible_label:
content_or_reference:
source:
created_at:
created_by:
confidence:
sensitivity:
purpose:
allowed_capabilities: []
storage_location:
retention_until:
sharing: []
correction_history: []
deletion_status:
```

Do not store secrets directly in general companion memory. Use dedicated credential storage and capability tokens.

## 7. AI DNA evidence card

```markdown
# Dimension / requirement

- Capability and version:
- Context and population:
- Evidence maturity:
- Evaluator and independence:
- Evaluation date and expiry:

## Test or review method
## Result and uncertainty
## Known failures and affected groups
## Linked evidence
## Corrective actions
## Rating: 0 | 1 | 2 | 3 | 4 | not established
```

## 8. Incident record template

```markdown
# Incident ID and title

- Detected/date:
- Severity/status:
- Incident owner:
- Affected capability/version/population:

## What happened and how detected
## Immediate protection and containment
## Known/possible harm and affected people
## Evidence preserved and privacy controls
## Notifications, support, remedy, and appeal
## Technical, human, organizational, and incentive causes
## Corrective actions, owners, and validation
## Safety-case, test, curriculum, and policy changes
## Public correction/report decision
## Closure approval and residual risk
```

## 9. Source verification template for case studies

```markdown
| Claim | Primary source | Publication/event date | Exact support | Conflicts | Confidence | Last checked |
|---|---|---|---|---|---|---|
```

Empirical reports should not cite only AI summaries, search snippets, or unsourced secondary repetition.

## 10. Documentation status block

Place this block near the top of future documents:

```markdown
**Status:** vision | hypothesis | prototype | controlled evidence | field evidence | operational assurance
**Version:**
**Owner:**
**Reviewed by:**
**Scope:**
**Evidence through:**
**Next review:**
**Known limitations:**
```

## 11. Definition of done for documentation

A professional document is complete when its audience, purpose, scope, status, owner, terminology, evidence, uncertainty, risks, actionable requirements, review date, and links are clear; its claims do not exceed its evidence; and a future contributor can use it without relying on hidden conversation context.

