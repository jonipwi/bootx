# BootX Public-Source Release Boundary

**Status:** Active conservative release-control decision; not a license amendment or safety certification  
**Decision date:** 2026-07-25  
**Final public-source development baseline:** Personal Companion `0.4.0-dev` and the accompanying repository documentation  
**Private successor namespace:** `0.5.0-private-dev` or a later clearly marked private-development version  
**Release commit/tag:** Pending designation when this boundary is committed and reviewed

## 1. Decision

The present BootX repository is the final planned **publicly source-available development baseline**. New development of risk-management and decision-assistance capabilities should continue in a controlled private-source environment.

This decision is intended to reduce premature exposure, copying, deployment, or misuse of unvalidated high-consequence decision logic while BootX lacks constituted governance, complete provenance, independent security and legal review, representative evaluation, and operational assurance.

“Final public baseline” does **not** mean that BootX is stable, validated, production-ready, safety-certified, legally approved, or authorized for public or High-Impact deployment. The `-dev` suffix and all existing warnings remain in force.

## 2. Precise terminology

BootX `0.4.0-dev` is **ethical source-available**, not OSI-approved open source or standard free software. Its public availability is governed by the [BootX Common Good Ethical Use License 1.0](LICENSE).

The planned successor is described as **private-source** or **closed-source development** because access to its source, security materials, model behavior, and high-risk decision logic will be restricted. Confidentiality is a risk-control measure, not evidence that the successor is safe, ethical, correct, or secure.

## 3. Rights already granted remain intact

This boundary does not revoke, narrow, or retroactively replace rights already granted for public BootX material under License 1.0. A recipient may continue to exercise those rights subject to the license.

The official project must:

- keep the exact License 1.0 text available with the public baseline;
- preserve public source and notices for the designated baseline;
- avoid claiming that earlier recipients accepted a future license;
- distinguish official continuity from independent compliant forks;
- preserve correction, incident, and material limitation records.

Publicly released material cannot be made secret merely by changing the project’s future development policy.

## 4. Public repository freeze

After the release commit and tag are designated, the public repository should accept only bounded maintenance that protects the value and safety of the inherited baseline:

- factual, mathematical, logical, accessibility, and documentation corrections;
- security advisories and safely publishable remediation;
- dependency, build, reproducibility, provenance, and license corrections;
- tests that verify existing public behavior without adding a new high-risk capability;
- removal or containment of exposed secrets, unlawful material, or credible hazards;
- clear deprecation, migration, archival, and continuity information.

The public repository should not receive:

- new high-consequence decision-making or decision-ranking capabilities;
- private successor source, prompts, model weights, credentials, exploit details, or internal threat intelligence;
- autonomous publication, enforcement, surveillance, financial, emergency, legal, medical, physical, or robotics authority;
- material that would allow a private safety control to be bypassed;
- marketing claims that the public baseline is complete, safe, reliable, or validated.

A safety-critical public fix may be backported when disclosure is lawful and does not create a greater immediate risk. Security-sensitive details may use coordinated disclosure, but confidentiality must not conceal harm, affected-user obligations, or accountability.

## 5. Private successor controls

Private development must begin in a separate private repository or an equivalently access-controlled environment. It must not rely on an unprotected branch in a public remote.

Minimum controls are:

1. named human owner and approved purpose for every capability;
2. least-privilege membership, multi-factor authentication, separated administrative roles, and prompt access revocation;
3. protected branches, required review, signed or otherwise attributable changes, and immutable audit records;
4. secret scanning, dependency and provenance inventory, backup, restore, incident response, and key-rotation exercises;
5. separated development, evaluation, and production data with no real sensitive data until the applicable gates pass;
6. deterministic policy and stop controls outside the generative model;
7. no model self-approval, no direct model-to-actuator or model-to-enforcement authority, and no irreversible action without accountable human control;
8. independent security, privacy, rights, domain, accessibility, and human-factors review proportionate to risk;
9. versioned test evidence, subgroup analysis, rollback, expiry, appeal, correction, and remedy;
10. a sanitized public accountability record of material limitations, incidents, and evidence when disclosure can be made safely and lawfully.

Private access does not relax License 1.0 obligations that apply to inherited public material, the Charter, AI DNA mandatory gates, human-rights protections, the Development Guideline, or the prohibition on autonomous High-Impact decisions.

## 6. Licensing and ownership gate for a closed-source successor

Private modification for internal research and closed-source distribution are different legal events. Before any private successor is distributed, licensed to another party, hosted for external users, commercialized, or transferred to a successor, qualified counsel and the accountable rights holders must determine:

- who owns every inherited and new contribution;
- whether contributors granted separate relicensing rights;
- whether License 1.0 share-alike, source, notice, deployment, or other duties apply;
- whether third-party code, data, models, documentation, and patents permit the intended use;
- whether a contributor agreement, dual-license structure, clean-room boundary, or replacement implementation is required;
- which privacy, consumer, product-safety, AI, sector, export, and jurisdictional obligations apply.

Code received only under License 1.0 must not be represented or redistributed as unrestricted proprietary code. If the project lacks the necessary rights, the affected component must remain under its existing terms, be separately licensed by the rights holder, or be replaced lawfully.

This document records project policy; it is not legal advice and does not itself create a proprietary license.

## 7. Publication from private development

The private project may publish bounded public-interest outputs such as:

- corrected scientific or mathematical findings;
- high-level safety lessons and known limitations;
- vulnerability advisories after coordinated remediation;
- evaluation methods, non-sensitive aggregate results, and accountability reports;
- interoperability specifications that do not expose protected controls.

Publishing successor source or reopening public feature development requires a documented D4 governance decision, applicable rights-holder approval, legal and security review, an explicit license, and a new public release record. Silence, an informal upload, or an AI recommendation is not approval.

## 8. Human decision authority

The private successor remains decision support. It may organize evidence, expose uncertainty, compare options, identify missing safeguards, and recommend review. It must not:

- determine guilt, sentence, eligibility, human worth, political loyalty, or religious truth;
- replace a qualified judge, clinician, emergency authority, safety engineer, or other accountable professional;
- hide policy choices behind “the AI decided”;
- turn a score, confidence value, or model response into automatic authority;
- manipulate a user into obedience, dependency, or surrender of conscience;
- act externally without a separately authorized, visible, reversible, and audited human-controlled capability.

## 9. Inheritance and continuity

The continuity register must identify separate custodians and recovery procedures for:

- the immutable public baseline and its license;
- the private successor source and encrypted backups;
- credentials, signing keys, deployment access, and audit evidence;
- contribution and relicensing records;
- incident, disclosure, appeal, and affected-user obligations;
- a safe shutdown or transfer if no qualified successor exists.

No successor inherits permission merely because they possess a repository, key, company role, or founder relationship. Authority must transfer through the governance and rights-holder process.

## 10. Required designation record

Before calling the public baseline final in a release announcement:

- [ ] commit this boundary and all intended public-baseline corrections;
- [ ] record the exact commit hash;
- [ ] create a signed or otherwise verifiable final public-baseline tag;
- [ ] archive a verified source copy with License 1.0 and checksums;
- [ ] complete a secret and private-information scan;
- [ ] record known limitations, unsupported artifacts, and open security/legal gates;
- [ ] verify public links, tests, build instructions, and artifact hashes;
- [ ] configure branch/repository protections against accidental private-source publication;
- [ ] obtain qualified legal review before representing any successor as proprietary or distributing it closed source.

Until these items are complete, `0.4.0-dev` is the declared final **planned** public-source baseline, but its immutable release identifier remains pending.
