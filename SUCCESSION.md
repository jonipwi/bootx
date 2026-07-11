# BootX Succession and Continuity Plan

**Status:** Stewardship continuity baseline<br>
**Version:** 1.0<br>
**Effective:** 2026-07-12

## 1. Purpose

BootX must be able to continue safely when a founder, maintainer, sponsor, provider, or AI model becomes unavailable. Succession protects the mission, users, knowledge, security, and capacity for correction. It does not preserve any individual's permanent control.

## 2. Succession principle

> Future stewards inherit responsibility and method—not unquestionable conclusions, personal status, or ownership of the people the project serves.

Successors are bound by the [Charter](CHARTER.md), [Governance](GOVERNANCE.md), applicable obligations, documented commitments to users, and current safety evidence.

## 3. Trigger events

This plan applies to:

- planned retirement, resignation, or extended leave;
- death or incapacity;
- loss of availability, credentials, funding, or hosting;
- removal for misconduct, conflict, persistent negligence, or Charter violation;
- organizational closure, acquisition, or mission change;
- compromise of keys or infrastructure;
- dependency on a discontinued AI model, vendor, or platform;
- community fork or responsible project dissolution.

No trigger grants an automatic right to the founder's relatives, employer, sponsor, largest contributor, or technology provider.

## 4. Continuity roles

| Role | Continuity responsibility |
|---|---|
| Continuity Coordinator | maintain the plan, asset register, exercises, and handover checklist |
| Interim Maintainer | preserve safe operations and records without making avoidable strategic changes |
| Security Custodians | recover or rotate keys and credentials using separated control |
| Stewardship Council | select and supervise permanent successors through the governance process |
| Independent Observer | verify that transfer, conflicts, incidents, and user commitments are handled transparently |

Names and private contact details belong in a protected operational register, not in this public document.

## 5. Required continuity register

Maintain and review a protected register covering:

- repositories, mirrors, branches, release artifacts, and reproducible-build instructions;
- domains, websites, package registries, communication channels, and service accounts;
- signing, encryption, backup, recovery, and hardware keys;
- cloud, model, data, telemetry, research, and security providers;
- licenses, copyrights, contributor records, trademarks, contracts, and obligations;
- data inventories, consent commitments, retention, deletion, and breach responsibilities;
- open incidents, vulnerabilities, safety cases, appeals, research protocols, and stop conditions;
- funding, expenses, sponsorship, insurance, and material conflicts;
- current maintainers, reviewers, user representatives, and emergency contacts;
- documents or knowledge held by only one person.

The register must never store plaintext secrets. It records where and how authorized custodians recover them.

## 6. Anti-single-point-of-failure requirements

- No critical production credential should depend on one person.
- Use separated custody and threshold approval for release signing, domain transfer, sensitive backups, and high-impact operations once the project has enough qualified custodians.
- Maintain tested, encrypted backups in more than one controlled location.
- Document build, release, recovery, incident, and data-deletion procedures.
- Preserve open formats and provider-independent export wherever practical.
- Maintain a supported last-known-good mode that does not require a specific remote AI model.
- Rotate credentials and review access promptly when a steward departs or is removed.

## 7. Successor criteria

A successor or stewardship body must demonstrate:

- commitment to the Charter and willingness to be corrected;
- competence appropriate to the role;
- responsible security, privacy, and evidence practice;
- understanding of affected users and limits of authority;
- disclosed conflicts and sufficient independence;
- ability to sustain operations, documentation, and community review;
- no history of concealing material harm or retaliating against dissent;
- acceptance of term limits, review, removal, and handover obligations.

Technical brilliance, funding, family relationship, model access, or founder loyalty alone is insufficient.

## 8. Planned transition

For a planned handover:

1. publish the intended transition and scope without exposing sensitive details;
2. update the continuity register and resolve undocumented dependencies;
3. nominate candidates through an open, conflict-disclosed process;
4. evaluate candidates against role criteria with affected-community input;
5. run parallel operation and recovery exercises;
6. transfer access gradually using least privilege;
7. rotate keys and remove unnecessary former access;
8. record outstanding risks, dissent, obligations, and decision expiry;
9. announce the accountable successor and appeal/contact channels;
10. complete an independent post-transition review.

The departing steward should be available for a defined advisory period where possible but should not retain hidden control.

## 9. Emergency transition

When a steward is suddenly unavailable:

1. protect people and stabilize services;
2. verify the trigger through more than one reliable channel where possible;
3. activate designated interim roles and freeze unnecessary high-risk changes;
4. preserve logs, records, code, research data, and evidence;
5. revoke or rotate affected credentials and inspect for compromise;
6. communicate known operational impact and support paths honestly;
7. restore only validated essential functions;
8. begin the permanent selection process under `GOVERNANCE.md`;
9. review whether concentration of knowledge or authority contributed to the event.

An interim steward may take necessary containment and continuity actions but may not exploit the emergency to amend the Charter, expand surveillance, transfer project assets for personal benefit, or erase dissent and incident evidence.

## 10. Founder unavailability

The founder's historical statement remains part of the record, but founder unavailability does not suspend governance or convert personal interpretation into permanent doctrine.

Future stewards should ask:

- Does this decision protect dignity and agency?
- What evidence supports it and what might disprove it?
- Who benefits, who bears risk, and who has not been heard?
- Can people review, appeal, exit, repair, and improve it?
- Are we preserving the mission or merely preserving our control?

## 11. Model and vendor succession

Every AI-dependent capability must document:

- provider-independent inputs, outputs, and capability contracts;
- minimum conformance and evaluation requirements;
- model/version provenance and known limitations;
- data export, deletion, and provider-exit procedures;
- fallback behavior when the model is unavailable or no replacement qualifies;
- regression, safety, privacy, and human-factors tests required for replacement.

A replacement model does not inherit approval automatically. It must pass the applicable release gates.

## 12. Forks and project integrity

The BootX Common Good Ethical Use License may permit compliant forks. A fork may innovate or disagree, but may not falsely claim governance authority, safety approval, certification, or official continuity it has not received. Trademark and identity rules should distinguish compatible experimentation from official BootX releases.

The official project should not use brand control to suppress truthful criticism or legitimate independent research.

## 13. Responsible dissolution

If safe continuation is impossible or no qualified successor exists:

1. stop new high-risk activity and preserve essential user support;
2. notify affected people and provide migration/export time;
3. revoke credentials and safely shut down services;
4. fulfill data deletion, research, contractual, and legal obligations;
5. archive source, documentation, decisions, and reproducibility materials where lawful and safe;
6. disclose unresolved hazards and unsupported artifacts;
7. transfer only to a qualified steward under the governance process, or close the project;
8. publish a final accountability and continuity report.

Continuing unsafely is not preferable to responsible closure.

## 14. Annual continuity exercise

At least annually, simulate loss of a maintainer, signing key, hosting provider, and AI model. Verify that authorized people can:

- locate authoritative records;
- reproduce and verify releases;
- restore a safe service or shut it down;
- rotate access without one unavailable person;
- contact users and reviewers;
- export/delete governed data;
- preserve incident evidence;
- identify and correct new single points of failure.

Record results, owners, deadlines, and retest evidence.
