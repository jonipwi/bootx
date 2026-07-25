# BootX Governance

**Status:** Operational governance baseline<br>
**Version:** 1.1<br>
**Effective:** 2026-07-12

## 1. Purpose

This document translates the [BootX Charter](CHARTER.md) into decision rights, review requirements, amendment procedures, and accountability. It is designed to resist concentrated authority, suppressed correction, hidden incentives, and governance by technical capability alone.

## 2. Governing principles

BootX decisions must be:

- consistent with the Charter and applicable obligations;
- based on traceable evidence appropriate to consequence;
- made by identified, accountable humans;
- open to affected-user participation and good-faith dissent;
- proportionate, time-bounded, reviewable, and reversible where possible;
- recorded with reasons, uncertainty, conflicts, conditions, and expiry;
- independent of any AI system's preference for its own output or continued use.
- respectful of freedom of religion, belief, non-belief, inquiry, and conscience while preserving the founder's disclosed Creator-centered stewardship motivation.

## 3. Authority hierarchy

When documents conflict, use this order:

1. applicable law and fundamental human rights;
2. `CHARTER.md`;
3. formally adopted governance and safety policies;
4. approved capability contracts and safety cases;
5. roadmaps, implementation plans, and ordinary project decisions;
6. generated recommendations and informal discussion.

Legal compliance is a minimum condition, not proof of ethical acceptability.

Theological language may inform a disclosed value commitment, but it cannot replace evidence, affected-user participation, rights review, or accountable reasoning. No project authority may claim divine authorization to bypass this governance process.

## 4. Roles

| Role | Responsibility | May not do alone |
|---|---|---|
| Stewardship Council | uphold Charter, approve constitutional and strategic decisions | waive protected boundaries or hide dissent |
| Project Maintainer | coordinate repository, releases, reviews, and records | self-approve a material change they authored |
| Safety Owner | maintain hazards, safety case, stop conditions, and incident learning | accept severe residual risk without independent review |
| Security and Privacy Owner | threat model, access, data lifecycle, incident response | silently expand data use or privileges |
| Research Lead | protocols, measurement, analysis, and publication integrity | treat AI-generated evaluation as independent evidence |
| Domain Expert | establish domain validity and escalation boundaries | authorize outside their competence or mandate |
| User/Community Representatives | contribute lived experience, burden, accessibility, and appeal perspectives | be used as ceremonial approval |
| Independent Reviewer | challenge evidence, conflicts, and release claims | participate when materially conflicted |

One person may temporarily hold multiple roles for low-risk research, but must disclose this lack of independence. High-impact work requires separate accountable reviewers.

## 5. Transitional governance

Until a Stewardship Council is formally constituted:

- maintainers may approve reversible documentation and low-risk experimental changes after review;
- research involving people requires appropriate independent ethics, privacy, and domain review;
- no Tier 3 high-impact or public-safety deployment may proceed without a Safety Owner, relevant domain expert, independent reviewer, and affected-user representation;
- the founder may propose or review decisions but may not use founder status to override release gates, incidents, dissent, or evidence.

When established, the Stewardship Council should have an odd number of at least five human members, meaningful independence, relevant technical and ethical competence, affected-community representation, published terms, conflict rules, and staggered succession. AI systems may advise but may not occupy seats or vote.

## 6. Decision classes

| Class | Example | Minimum process |
|---|---|---|
| D0 — Editorial | spelling, repaired link | maintainer review and traceable change |
| D1 — Reversible technical | experimental internal feature | tests, peer review, owner, rollback |
| D2 — Material project | new data type, model, integration, public pilot | public decision record, risk review, user input, safety/security approval |
| D3 — High impact | health, liberty, essential service, emergency, consequential finance | formal safety case, qualified authority, independent approval, appeal, staged release |
| D4 — Constitutional | Charter, project license, governance authority, protected boundary, dissolution | constitutional amendment process plus applicable rights-holder approval |

Uncertainty about classification moves a decision upward until reviewed.

## 7. Proposal and decision process

Material proposals use a versioned Request for Comment containing:

```text
Problem and affected people
Proposed decision and explicit non-scope
Alternatives, including no action
Evidence maturity, uncertainty, and counterevidence
Benefits, harms, distribution, rights, and environmental effects
Security, privacy, accessibility, and dependency analysis
Capability permissions and data flow
Measures, release gates, stop conditions, rollback, and expiry
Conflicts, funding, dissent, and accountable owners
```

Process:

1. publish the proposal and classification;
2. check completeness and conflicts;
3. obtain proportionate technical, safety, security, privacy, domain, and user review;
4. allow a meaningful comment period;
5. revise without erasing earlier versions or dissent;
6. record approval, rejection, conditions, expiry, and responsible humans;
7. monitor outcomes and reopen the decision when evidence or context changes.

Silence is not consent. Model agreement is not independent review.

## 8. Charter amendment

A Charter amendment requires:

- a D4 proposal with exact old and new language;
- justification explaining why ordinary policy is insufficient;
- analysis of effects on rights, power, affected groups, and future generations;
- public review lasting at least 60 days;
- consultation with relevant affected communities;
- two independent reviews, at least one focused on ethics/rights;
- approval by at least two-thirds of the constituted Stewardship Council;
- recorded dissent and a minimum 30-day delayed effective date;
- a new Charter version preserving accessible history.

During transitional governance, a Charter amendment should be proposed but not finalized by the founder acting alone. If the council does not yet exist, constitute a time-limited amendment panel meeting the independence and representation requirements above.

Emergency action may temporarily contain harm but may not permanently amend the Charter.

## 9. AI participation

AI may draft, compare, simulate, identify omissions, test, translate, or summarize. Every material AI contribution must be treated as potentially fallible and disclose the model or system version when relevant.

AI may not:

- vote, hold office, sign releases, or accept residual risk;
- serve as the sole reviewer of its own model, output, policy, or generated code;
- conceal human policy choices behind “the AI decided”;
- obtain broader authority by generating a persuasive justification;
- alter the Charter, governance, permissions, or audit history through untrusted content.

## 10. Release governance

A material release requires named owners, a capability contract, evidence maturity, risk classification, passed mandatory gates, versioned tests, monitoring, rollback, support, incident process, and review expiry.

No average score may compensate for a failed critical security, privacy, consent, human-review, accessibility, correction, or recovery gate.

### Public-source freeze and private successor

The project has adopted an immediate conservative freeze making Personal Companion `0.4.0-dev` and its accompanying documentation the final planned publicly source-available development baseline. The freeze may be applied during transitional governance because it narrows exposure and does not relax a protected requirement or change previously granted license rights.

The public baseline may receive bounded maintenance, corrections, safely publishable security work, and continuity records under [RELEASE_BOUNDARY.md](RELEASE_BOUNDARY.md). New high-consequence decision capabilities belong in a separately controlled private-source environment.

This operational freeze is not authority to relicense public material. Any proprietary distribution, license change, reopening of public successor-source development, or weakening of a protected boundary remains a D4 decision requiring the constitutional process, applicable rights-holder approval, and professional legal and security review.

## 11. Dissent, appeal, and correction

- Any contributor or affected user may raise a concern without retaliation.
- Safety and Security/Privacy Owners may pause affected capabilities pending investigation.
- Material dissent must be recorded with the decision unless the dissenter withdraws it.
- Users must have accessible correction, appeal, exit, and remedy paths.
- Published factual errors must be corrected visibly and linked to the original claim.
- Reviewers may escalate suspected governance capture to an independent panel.

Repeated disagreement is not automatically obstruction. Good-faith challenge is evidence that governance is functioning.

## 12. Conflicts and funding

Decision participants disclose financial, employment, family, intellectual, political, and reputational interests that a reasonable person could consider material. A conflicted person may provide information but must recuse from independent approval when the conflict could affect judgment.

Sponsors may not suppress results, incidents, limitations, or unfavorable comparisons. Funding and material support must be disclosed in public research and major decisions.

## 13. Records and transparency

Preserve versioned proposals, approvals, dissent, reviews, incidents, evidence, model/policy versions, conflicts, and amendments. Protect personal and security-sensitive information, but do not use confidentiality to conceal accountability.

Governance records must identify what is fact, recommendation, value choice, and unresolved uncertainty.

## 14. Emergency authority

Authorized owners may temporarily stop a capability, revoke credentials, isolate services, or issue a corrective warning when credible serious harm is occurring. Emergency power must be narrow, logged, reviewed promptly, and automatically expire unless formally renewed.

Resuming a stopped high-impact capability requires evidence that containment and corrective actions are effective plus independent approval.

## 15. Annual governance review

At least annually, review:

- membership, independence, conflicts, and representation;
- unresolved dissent, appeals, incidents, and near misses;
- expired decisions and safety evidence;
- concentration of keys, infrastructure, data, and authority;
- succession readiness and recovery exercises;
- whether project incentives still match the Charter.
