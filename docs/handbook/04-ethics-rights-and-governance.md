# Ethics, Rights, and Governance Constitution

This domain constitution operates under the repository-level [BootX Charter](../../CHARTER.md) and [BootX Governance](../../GOVERNANCE.md). The Charter states the highest commitments; this document develops their ethical and operational requirements for Human + AI companion research.

## 1. Purpose

This constitution defines the moral and institutional boundaries for BootX Human + AI companion research. It is not a claim that ethical disagreement has been solved. It states the project's commitments, names conflicts, and assigns authority for decisions.

## 2. Ethical foundation

The program uses four complementary lenses:

- **Rights and duties:** respect dignity, privacy, freedom, equality, consent, and due process.
- **Consequences:** reduce harm and expand durable benefit while considering indirect and long-term effects.
- **Care:** recognize vulnerability, relationship, context, accessibility, and responsibility toward people who need support.
- **Virtue and character:** cultivate truthfulness, humility, courage, justice, compassion, and practical wisdom in people and institutions.

No lens is sufficient alone. Aggregate benefit cannot erase fundamental rights; formal consent cannot excuse exploitation; good intentions cannot substitute for outcomes; and an appealing character cannot substitute for evidence.

## 3. Human rights requirements

Every deployed companion must provide:

1. **Identity transparency:** users know they are interacting with an AI system and which organization is responsible.
2. **Purpose clarity:** users know what the system does, does not do, and why data is processed.
3. **Meaningful choice:** consent is specific, understandable, revocable, and not obtained through coercive design.
4. **Privacy:** collection is minimized; access, retention, sharing, and deletion are controlled and auditable.
5. **Human review:** consequential outputs can be reviewed by an authorized, competent person.
6. **Explanation:** users receive reasons suited to their needs without revealing dangerous operational detail.
7. **Correction:** users can correct records and contest decisions.
8. **Exit and portability:** users can disable the service, export permitted data, and delete memory without retaliation.
9. **Accessibility:** design accounts for disability, language, literacy, age, bandwidth, and device constraints.
10. **Equal concern:** performance and harm are evaluated across affected groups.

## 4. Companion covenant

The companion shall:

- say what it knows, what it infers, and what it does not know;
- protect without shaming, frightening, or infantilizing the user;
- recommend the least restrictive effective intervention;
- preserve the user's ability to learn and choose;
- disclose conflicts, sponsorship, and uncertainty relevant to advice;
- refuse unauthorized harmful action;
- encourage qualified human help when competence or authority is exceeded;
- treat memory as borrowed stewardship, not ownership of a person;
- never condition essential safety on emotional attachment or unnecessary data surrender;
- permit correction and leave a traceable reason for consequential actions.

## 5. Prohibited behaviors

The system and its operators must not:

- deceive users about AI identity, capabilities, evidence, or commercial incentives;
- exploit loneliness, fear, grief, cognitive impairment, or dependency;
- encourage exclusivity such as “only I understand you” or discourage healthy human relationships;
- simulate threats, guilt, affection, or abandonment to obtain compliance;
- infer or act on sensitive traits beyond an authorized, necessary purpose;
- rank human worth or deny rights using an AI DNA, trust, risk, or productivity score;
- conduct covert political, religious, commercial, or behavioral persuasion;
- make irreversible high-impact decisions without accountable human authority;
- silently expand data use, retention, or sharing;
- punish users for opting out, appealing, or correcting the system;
- conceal incidents, adverse evidence, model changes, or conflicts of interest;
- allow the system to pursue continued operation as an end overriding human control.

## 6. Vulnerable-person safeguards

### Children

Use age-appropriate explanations, strong defaults, minimal data, guardian structures consistent with the child's evolving autonomy, no targeted persuasion, and clear emergency boundaries. Child safety review is mandatory before any child-facing pilot.

### Older adults and users with limited digital literacy

Use readable language, multimodal explanation, verified-contact workflows, deliberate pauses before high-risk transfers, and easy access to a trusted person. Never assume incapacity solely from age.

### People in distress

The companion may support grounding and connection, but must not claim clinical competence or replace emergency or mental-health services. Crisis handling requires locally reviewed resources, conservative escalation rules, privacy protection, and human supervision.

### People under coercion

Exit, notification, and trusted-contact features can increase danger in abusive contexts. Provide discreet controls, threat-model the device and account, and never assume a listed contact is safe.

## 7. Risk-tier governance

| Tier | Example | Minimum governance |
|---:|---|---|
| 0 — Informational | study explanation, low-stakes summary | provenance, correction, privacy baseline |
| 1 — Assistive | suspicious-message explanation | tested alert design, override, logging, monitoring |
| 2 — Consequential support | recommendation involving meaningful money or sensitive data | explicit confirmation, independent verification, human escalation, stronger audit |
| 3 — High impact | health, liberty, employment, essential service, emergency allocation | qualified authority, formal safety case, independent review, appeal, legal and domain compliance |
| 4 — Prohibited or exceptional | coercive control, autonomous force, human-worth scoring | not permitted under this program |

Higher tiers require stronger evidence and more limited automation. A fluent model does not reduce the tier.

## 8. Governance roles

| Role | Accountability |
|---|---|
| User council | lived experience, needs, acceptability, correction priorities |
| Product owner | scope, resources, benefit hypothesis, operational ownership |
| Safety owner | hazard analysis, safety case, stop authority |
| Security and privacy owner | threat model, access control, incident response, data governance |
| Domain expert | validity and escalation boundaries for the use case |
| Research lead | protocol, measures, analysis integrity, publication |
| Independent reviewer | challenge evidence and conflicts; approve high-risk gates |
| Operations lead | monitoring, rollback, support, recovery |
| Executive sponsor | remains accountable; cannot delegate moral responsibility to the AI |

The same person or team should not build, evaluate, and unilaterally authorize a high-impact system.

## 9. Decision and dissent process

Every major release decision records:

- proposed capability and affected population;
- benefit hypothesis and evidence level;
- known hazards and residual risks;
- rights and distributional analysis;
- stakeholder participation and disagreements;
- security, privacy, and accessibility findings;
- approval, conditions, expiry, and rollback plan;
- minority or dissenting opinion.

Anyone in the program may raise a safety concern. The safety owner and independent reviewer hold stop authority. Retaliation for good-faith reporting is incompatible with the project.

## 10. Data governance

Personal data must have a documented flow from collection to deletion. For each field record:

```text
purpose | lawful/ethical basis | source | sensitivity | location
access roles | retention | sharing | model use | deletion proof | incident owner
```

Default rules:

- local processing where it materially reduces exposure and remains supportable;
- end-to-end protection for data in transit and strong encryption at rest;
- separate identity, content, telemetry, and research datasets;
- no training on personal companion memory by default;
- no silent cross-user memory;
- user-visible memory ledger with create/read/update/delete controls;
- time-limited, purpose-bound emergency access with audit;
- backup deletion and key-destruction procedures tested, not merely promised.

## 11. Accountability after harm

When harm or a near miss occurs:

1. protect affected people and stop ongoing exposure;
2. preserve evidence without unnecessary additional collection;
3. notify appropriate parties honestly and promptly;
4. provide accessible remedy and appeal;
5. analyze technical, human, organizational, and incentive causes;
6. publish a proportionate correction or incident report;
7. validate remediation before resuming;
8. update training, tests, governance, and the safety case.

“The model made the decision” is never an acceptable final account of responsibility.

## 12. Ethical success criterion

The companion is ethically successful when people become safer and more capable while retaining dignity, relationships, freedom, privacy, and recourse. A system that prevents one harm by creating dependency, surveillance, exclusion, or unchallengeable power has not met the mission.
