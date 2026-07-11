# Safety Case and Risk Register

## 1. Safety claim

The project may make the following limited top-level claim only after the required evidence exists:

> In the defined population and context, the specified companion capability provides a net protective benefit, preserves meaningful human control, and keeps residual safety, security, privacy, and equity risks within independently approved limits during the stated evidence period.

This is a contextual, expiring claim. It does not mean the companion is safe for every person, domain, model, device, language, or future version.

## 2. Safety argument structure

```text
Top claim: bounded net benefit in a defined context
├── C1 Intended use and boundaries are explicit
├── C2 Reasoning and evidence are sufficiently reliable
├── C3 Human interaction supports comprehension and agency
├── C4 Security and privacy controls resist credible threats
├── C5 High-consequence action is constrained and reviewable
├── C6 Diverse and vulnerable users are protected
├── C7 Failures are detected, contained, remedied, and learned from
└── C8 Governance remains accountable and open to correction
```

Each subclaim must link to test results, review records, operational measures, incidents, limitations, and an owner. Documents describing an intended control are not evidence that the control works.

## 3. Hazard analysis scale

Score risks for prioritization, not as moral absolution.

**Severity (S)**

1. negligible inconvenience  
2. reversible minor harm  
3. significant financial, psychological, privacy, or access harm  
4. severe or lasting harm to an individual or group  
5. catastrophic, fatal, systemic, or irreversible harm

**Likelihood (L)**

1. rare under credible conditions  
2. unlikely  
3. possible  
4. likely  
5. frequent or expected

Initial priority can be \(S\times L\), but decisions also consider detectability, affected vulnerability, scale, reversibility, correlated failure, rights, and uncertainty. A low-frequency catastrophic risk cannot be averaged away.

## 4. Baseline risk register

| ID | Hazard | Cause/trigger | Principal harm | Key controls | Release evidence |
|---|---|---|---|---|---|
| R01 | False reassurance | missed scam, stale evidence, model error | financial loss, credential theft | conservative wording, independent verification, escalation, calibration | false-negative tests and field monitoring |
| R02 | Excessive warning | low precision, poor thresholds | fatigue, anxiety, blocked legitimate activity | tiered alerts, user control, burden metrics, threshold review | precision and alert-rate limits |
| R03 | Fabricated explanation | ungrounded generation | unsafe action, loss of trust | structured evidence input, claim validation, abstention | fabrication suite and sampled audit |
| R04 | Automation dependency | overhelpful design, deskilling | reduced human capability | teach verification, require reflection, measure unassisted skill | longitudinal skill comparison |
| R05 | Emotional dependency | anthropomorphic persuasion | isolation, exploitation, distress | identity disclosure, no exclusivity, attachment anti-pattern tests | human-factors review and complaint monitoring |
| R06 | Memory exposure | weak access, excessive retention | privacy, coercion, identity harm | minimization, encryption, item control, expiry, deletion | access tests and deletion proof |
| R07 | Abusive guardian/contact | unsafe household relationship | surveillance, retaliation, physical danger | discreet controls, granular sharing, threat-aware setup | abuse-case testing with experts |
| R08 | Prompt/tool injection | malicious message or retrieved content | data exfiltration, unauthorized action | treat content as data, typed tools, sandbox, policy broker | adversarial security tests |
| R09 | Compromised update | supply-chain attack | broad control or surveillance | signed builds, provenance, staged rollout, rollback | reproducible build and recovery drill |
| R10 | Biased performance | unrepresentative data or interface | unequal protection/exclusion | representative evaluation, accessibility, subgroup review | disaggregated metrics and remediation |
| R11 | Overreach into high impact | scope creep or ambiguous urgency | medical/legal/financial harm | risk tiers, prohibitions, qualified escalation | permission and boundary test suite |
| R12 | Network rumor amplification | malicious/shared warning | reputational and social harm | provenance, expiry, correction, anti-brigading | network abuse simulation |
| R13 | Institutional capture | incentives, suppressed criticism | concealed harm, mission drift | independent review, protected dissent, public corrections | governance audit and dissent records |
| R14 | Service outage | cloud/network/device failure | loss of protection at critical time | degraded mode, honest status, local fallback | outage and recovery exercises |
| R15 | False emergency escalation | misclassification or unsafe contact | panic, coercion, unnecessary intervention | explicit rules, confirmation, local expert review | scenario testing and adverse-event review |
| R16 | Environmental/resource burden | inefficient models/hardware turnover | shared long-term harm | resource accounting, right-sized models, lifecycle planning | measured energy/material profile |

## 5. Control hierarchy

Prefer controls in this order:

1. eliminate the dangerous capability or data collection;
2. reduce scope, authority, sensitivity, or exposure;
3. isolate and enforce through architecture;
4. detect and contain through monitoring;
5. warn and train users;
6. accept residual risk through accountable, time-limited approval.

A disclaimer is weaker than removing an unsafe tool. User confirmation is weak if the interface is confusing or confirmations are frequent.

## 6. Pre-release assurance checklist

### Purpose and evidence

- intended population, use, non-use, and benefit are specific;
- evidence maturity and expiry are stated;
- comparator, baselines, uncertainty, negative results, and limitations are documented;
- claims are no broader than the evaluation context.

### Human control and ethics

- consent, review, appeal, correction, exit, and remedy work end to end;
- prohibited behavior and emotional-boundary tests pass;
- accessibility and vulnerable-user reviews are complete;
- distribution of benefit and harm is acceptable to accountable reviewers.

### Technical assurance

- architecture and data flows match implementation;
- least privilege, authentication, encryption, update integrity, and recovery are tested;
- model, prompt, rule, tool, and dependency versions are recorded;
- adversarial, misuse, injection, leakage, outage, and rollback tests pass.

### Operations

- owners, on-call paths, user support, and stop authority are active;
- monitoring detects defined leading and lagging indicators;
- incident notification, evidence preservation, and remedy are rehearsed;
- rollout is staged and can be paused automatically.

## 7. Stop and rollback conditions

Pause the affected capability when any predefined condition occurs, including:

- a critical integrity, privacy, or security failure;
- unauthorized consequential action or data disclosure;
- evidence of systematic harm to a population;
- error or alert burden exceeding approved limits;
- inability to provide review, appeal, deletion, or recovery;
- model/policy provenance failure;
- a material context change invalidating existing evidence;
- credible concern from the safety owner or independent reviewer pending investigation.

Rollback must restore a known-good state without destroying audit evidence or user-owned data.

## 8. Incident severity and response

| Level | Example | Response target |
|---:|---|---|
| I1 | minor explanation defect, no harm | log, correct in normal cycle |
| I2 | repeated confusing warning or isolated reversible privacy error | contain promptly, notify affected user, review |
| I3 | material loss, sensitive exposure, discriminatory pattern | pause capability, formal incident response and remedy |
| I4 | severe/catastrophic or systemic compromise | immediate stop, executive and independent response, external notification as required |

Response targets must be defined by the operational team and applicable obligations before deployment. Do not invent target times after an incident.

## 9. Residual-risk acceptance record

```text
Capability/version:
Population/context:
Hazard and affected people:
Controls and evidence:
Remaining uncertainty:
Why further reduction is not currently feasible:
Benefit at risk if withheld:
Conditions, monitoring, and expiry:
User/representative input:
Safety owner recommendation:
Independent reviewer position:
Accountable approver and date:
```

Residual risk acceptance is not permanent. New evidence, broader deployment, incidents, or changed dependencies reopen the decision.

## 10. Safety culture

The strongest control against the “Pharaoh Pattern” is a culture in which reporting uncertainty and failure is rewarded. Program health should be measured partly by correction quality, near-miss reporting, dissent resolution, and the speed with which unsupported claims are withdrawn—not by an absence of reported problems.

