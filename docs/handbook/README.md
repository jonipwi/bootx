# BootX Human + AI Companion Research Handbook

**Document status:** Research baseline 1.2<br>
**Last implementation review:** 2026-07-24<br>
**Audience:** learners, researchers, engineers, educators, reviewers, and future maintainers<br>
**Purpose:** turn the ideas in the BootX repository into a testable, teachable, and implementable Human + AI companion program

## Project constitution

The handbook operates under these repository-level documents:

- [BootX Charter](../../CHARTER.md) — enduring mission and protected commitments;
- [Governance](../../GOVERNANCE.md) — human authority, review, amendments, dissent, and release decisions;
- [Succession](../../SUCCESSION.md) — continuity beyond any founder, maintainer, vendor, or AI model;
- [Common Good Ethical Use License](../../LICENSE) — adopted custom royalty-free permissions and ethical-use conditions; enforceability and compatibility remain subject to qualified legal review;
- [Licensing policy](../../LICENSING.md) — classification, AI DNA conditions, and distribution guidance;
- [Disclaimer](../../DISCLAIMER.md) — limitations that do not substitute for safety controls;
- [Development guideline](../../DEVELOPMENT_GUIDELINE.md) — personal-first companion and robotics boundaries, forecast/disaster warning status, do/don't contract, staged implementation, and quantitative evidence gates;
- [Progress and inheritance plan](../../PROGRESS.md) — evidence-based readiness, incomplete gates, and ordered future work.

When an ordinary handbook provision conflicts with the Charter, the conflict must be resolved through the governance process rather than silently interpreted away.

## Mission

BootX explores a future in which AI strengthens human judgment, safety, dignity, learning, and long-term resilience. This handbook converts that aspiration into a disciplined research and development framework.

The governing mission is:

> Develop Human + AI companions that help people understand reality, make safer decisions, retain meaningful control, learn from error, and cooperate for the common good.

"Benefit humanity" is not treated as a slogan or a single optimization target. It is decomposed into rights, safety constraints, measurable outcomes, accountable processes, and continuous correction.

## How to use this handbook

Read the documents in order for a complete study. Engineers can begin with the system specification and roadmap after reviewing the ethical constitution. Researchers should begin with the evidence audit, logic, and measurement framework.

The current implementation evidence is the [host-based Go Personal Companion `0.4.0-dev`](../../prototype/personal-companion/README.md). Its personal, warning, local-document, and Law Clarity paths remain deterministic. A separate public/non-sensitive ethical-review path uses OpenAI only after explicit consent, with a fixed no-tools structured-output contract and visible remote receipt. Its integrity and declared-source records do not authenticate people, publishers, authorities, sources, or claims; its law and review indices are not legal findings or probabilities. It does not close the handbook's research, security, identity, human-factors, legal, fairness, or deployment gates.

| Part | Document | Primary question |
|---:|---|---|
| 0 | [Foundation and scope](00-foundation-and-scope.md) | What is BootX trying to achieve, and what is it not claiming? |
| 1 | [Repository evidence audit](01-repository-evidence-audit.md) | What is implemented, proposed, illustrative, or unverified? |
| 2 | [Logic and epistemology](02-logic-and-epistemology.md) | How should the companion reason and correct itself? |
| 3 | [Mathematics and measurement](03-mathematics-and-measurement.md) | How are risk, trust, uncertainty, and benefit measured? |
| 4 | [Ethics, rights, and governance](04-ethics-rights-and-governance.md) | What must the companion protect, and who remains accountable? |
| 5 | [AI DNA specification](05-ai-dna-specification.md) | What operational traits define trustworthy behavior? |
| 6 | [Companion system architecture](06-companion-system-architecture.md) | How can the principles become a secure system? |
| 7 | [Research and implementation roadmap](07-research-and-implementation-roadmap.md) | What should be built, tested, and reviewed first? |
| 8 | [Curriculum and assessment](08-curriculum-and-assessment.md) | How can future contributors learn and demonstrate competence? |
| 9 | [Safety case and risk register](09-safety-case-and-risk-register.md) | What can go wrong, and what evidence is required before release? |
| 10 | [Glossary and study templates](10-glossary-and-templates.md) | Which shared terms and reusable records should the project use? |
| 11 | [Standards and authoritative reading map](11-standards-and-reading-map.md) | Which external frameworks should future research compare and extend? |
| 12 | [Anti-abuse and evidence integrity](12-anti-abuse-and-evidence-integrity.md) | How should BootX resist deepfakes, poisoned evidence, correlated sources, and abuse by authorities? |
| 13 | [Documentation verification and corrections register](13-documentation-verification-register.md) | Which claims were checked, corrected, withdrawn, or remain unverified? |
| 14 | [Personal decision-assistance pipeline](14-personal-decision-pipeline.md) | How should personal input become evidence-bounded options and a human-controlled decision output? |
| 15 | [Law Clarity Logic](15-law-clarity-logic.md) | How can a public rule be screened for clarity, fairness, consistent enforcement, accountability, loopholes, ambiguity, and unchecked power without manufacturing legal authority? |
| 16 | [Ethical Publication and Decision-Rationale Review](16-ethical-publication-review.md) | How can declared-evidence mathematics and bounded OpenAI critique help a human pause before publication without producing censorship, truth authority, or automated action? |

## Foundational commitments

1. **Human dignity:** A person has value independent of productivity, intelligence, wealth, health, nationality, belief, or data value.
2. **Human agency:** The system supports decisions; it does not covertly control them.
3. **Truth with humility:** Claims are tied to evidence, uncertainty is visible, and correction is expected.
4. **Protection without domination:** Safety controls must be proportional, contestable, and accountable.
5. **Privacy and ownership:** Personal memory and context remain under meaningful user control.
6. **Justice:** Benefits, errors, burdens, and access are evaluated across different populations, especially vulnerable people.
7. **Common good:** The design considers families, communities, institutions, future generations, and the environment—not only individual convenience.
8. **Continuous learning:** Decisions and failures produce auditable lessons rather than hidden mistakes.

Only a **conforming AI agent** may participate in BootX implementations. Vendor neutrality does not remove evaluation, least-privilege, provenance, release-gate, monitoring, and human-accountability requirements.

## Research posture

This handbook separates four kinds of statement:

- **Verified artifact fact:** directly inspectable in this repository.
- **Supported design proposition:** justified by explicit reasoning and available evidence.
- **Research hypothesis:** plausible and testable, but not yet established.
- **Value commitment:** a normative choice that must be governed openly rather than disguised as a scientific fact.

No composite score, metaphor, benchmark, or AI-generated explanation is sufficient by itself to authorize a high-impact decision. Human accountability, domain expertise, independent evidence, and appeal mechanisms remain necessary.

## Definition of success

The program succeeds only when evidence shows that the companion:

- reduces preventable harm without creating larger hidden harms;
- improves calibrated understanding rather than confidence alone;
- preserves or increases human agency;
- resists scams, manipulation, coercion, and dependency;
- protects sensitive data throughout its life cycle;
- works acceptably for diverse and vulnerable users;
- exposes limitations and supports correction;
- can be safely disabled, repaired, audited, and exited.

The initial target should be a narrow, reversible protective companion—not an autonomous governor, medical authority, therapist, police system, or controller of critical infrastructure.
