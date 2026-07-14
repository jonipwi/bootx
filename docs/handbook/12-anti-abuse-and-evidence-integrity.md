# Anti-Abuse and Evidence Integrity

**Status:** Supported design proposition with a documented factual self-audit<br>
**Version:** 1.2<br>
**Evidence reviewed through:** 2026-07-14<br>
**Purpose:** prevent false, manipulated, correlated, or politically abused evidence from turning AI assistance into harmful action

## 1. Executive summary

AI does not observe reality directly. It processes representations of reality: sensor readings, documents, databases, images, audio, video, testimony, and retrieved text. If that evidence is false, incomplete, stale, decontextualized, poisoned, or controlled by one authority, a capable AI may produce a persuasive but dangerous conclusion.

The governing principle is:

> No high-impact recommendation should depend on evidence whose origin, integrity, independence, uncertainty, and relevance have not been evaluated proportionately to the possible harm.

Forecasting is also only one component of protection. An effective early-warning system combines risk knowledge, monitoring and forecasting, authoritative communication, and preparedness for action. Failure or poor coordination in one component can cause the system to fail. This is consistent with the official [UNDRR definition of an early-warning system](https://www.undrr.org/terminology/early-warning-system).

The chapter establishes:

- an evidence-integrity chain;
- a fake-evidence and abuse threat model;
- corrected logical, Bayesian, and decision-theory formulations;
- thirteen anti-abuse principles;
- verification tiers and emergency-action rules;
- governance controls against misuse by leaders or institutions;
- AI DNA requirements, tests, and release evidence;
- a factual self-audit derived from the Bavi/Maysak conversation.

## 2. Origin and factual self-audit

This chapter originated from a Human + AI conversation about disaster preparedness, mathematical proof, leader abuse, deepfakes, and multi-source verification. The conversation produced useful principles, but the first audit confused two different tropical-cyclone naming cycles. This section records the correction visibly.

### 2.1 Name reuse and event identity

`Bavi` and `Maysak` were used in both 2020 and 2026. A storm name alone is therefore not a unique event identifier. A reliable record includes year, basin, international or national storm number, dates, and location.

For **2020**, the official sequence was Bavi (`2008`) followed by Maysak (`2009`). The Republic of Korea's Typhoon Committee member report records Bavi from 22 August 2020 and Maysak as the following storm. The Typhoon Committee season summary states that Maysak formed on 27 August and was named on 28 August. Sources: [Republic of Korea member report](https://www.typhooncommittee.org/15IWS/docs/Members%20REport/RoK/20201126_15IWS_MEMBER_REPORT_NTC_final.pdf) and [2020 season summary](https://typhooncommittee.org/53rd/docs/item%205/5.1.SummaryOf2020TyphoonSeason_20210210.pdf).

For **2026**, Taiwan's Central Weather Administration database identifies
[Bavi as `202609`](https://rdc28.cwa.gov.tw/TDB/public/typhoon_detail?typhoon_id=202609)
and [Maysak as `202610`](https://rdc28.cwa.gov.tw/TDB/public/typhoon_detail?typhoon_id=202610),
but Maysak's destructive Guangxi impacts occurred before Bavi reached eastern
China. Consequently, describing the Guangxi Maysak disaster as the preceding
event in a July 2026 preparedness discussion can be chronologically meaningful
even though Bavi had the lower storm number.

### 2.2 Verified 2026 Maysak facts

Authoritative reports support the core Guangxi claim:

- Maysak generated extreme, persistent rainfall in Guangxi in early July 2026.
- On 6 July, the Liulan and Yunbiao reservoirs in Hengzhou experienced overtopping and breaches; downstream villages flooded. This is more accurately described as **reservoir overtopping and major dam breaches**, not the complete collapse of every affected reservoir. See [Xinhua's 6 July event report](https://www.news.cn/politics/20260706/6a15220b8dea426a877776efc9e38caa/c.html) and [8 July technical explanation](https://www.news.cn/politics/20260708/1db47f4e0ff14918b7aaf6e00c589ca1/c.html).
- At the 7 July reporting cut-off, Guangxi authorities reported 375,000 affected people and 130,000 emergency relocations. These figures were preliminary. See the [Xinhua report from the Guangxi disaster briefing](https://www.gx.news.cn/20260708/9be922b8de8e446ba44bd6c85f7b8b38/c.html).
- At the 9 July reporting cut-off, authorities reported 26 deaths and 7 missing in the Liulan reservoir breach flood, and 39 deaths and 9 missing across Guangxi. These are dated preliminary counts, not final totals and not all attributable to one dam. See the [Xinhua report from the Nanning briefing](https://www.gx.news.cn/20260709/acc9673028f44188afd8df576a49bb23/c.html).

### 2.3 What remains unproven

The 2026 disaster does **not** by itself prove that the later Bavi response was successful because authorities learned from Maysak. That causal claim requires comparable hazard, exposure, vulnerability, decisions, timing, and outcome evidence. Likewise, an evacuation count does not by itself establish preparedness quality, and a lower casualty count cannot be converted directly into “lives saved.”

As of the 14 July 2026 evidence cut-off, Bavi had made landfall at Yuhuan, Zhejiang, at approximately 23:20 local time on 11 July with reported maximum winds of 40 m/s, made a second landfall at Yueqing around midnight, and continued inland while weakening. Xinhua directly reported 1.716 million people transferred in Zhejiang by 08:00 on 11 July. A preliminary local report for Yuhuan recorded affected population, damaged structures, and direct economic loss, but final Bavi casualty, evacuation, infrastructure, dam, and comparative-success findings were still unavailable. See the [Xinhua preparedness report](https://www.news.cn/20260711/36c781341f824fd88303d391023fb54e/c.html), [China News Service local-impact report](https://www.chinanews.com/m/sh/2026/07-12/10657864.shtml), and [13 July inland-alert report](https://english.www.gov.cn/news/202607/13/content_WS6a548da6c6d00ca5f9a0c275.html).

This correction is not incidental. It demonstrates the chapter's central lesson:

> A fluent AI can generate a sound general principle inside an unsound factual narrative. Every layer must be checked independently.

The original conversation remains valuable as a research-generation artifact, but its event claims must be read through this dated correction and the [documentation verification register](13-documentation-verification-register.md).

## 3. What can and cannot be proven

### 3.1 Forecasting is not logically sufficient

Let:

- \(F\): a usable forecast exists;
- \(A\): effective protective action occurs;
- \(S\): the desired safety outcome occurs.

The statement

$$
F \nRightarrow S
$$

means only that a forecast does not logically guarantee safety. A forecast does not itself evacuate a person, close a road, reinforce a structure, deliver aid, or make a leader act responsibly.

This does **not** prove that forecasting never saves lives, is always necessary, or has no causal contribution. Forecasting may change decisions and timing, which can change exposure and outcomes.

### 3.2 Combined action is not a deterministic guarantee

The conversation proposed:

$$
F \land D \land L \land C \Rightarrow R
$$

where \(D\) is disaster preparedness, \(L\) leadership, \(C\) public
cooperation, and \(R\) reduced exposure. This is useful as a process map but
too strong as formal logic. Even excellent forecasting, preparation,
leadership, and cooperation may fail under an extreme hazard, infrastructure
collapse, inaccessible populations, or incorrect assumptions.

A scientifically defensible hypothesis is probabilistic and contextual:

$$
\Pr(R=1 \mid F,D,L,C,X) > \Pr(R=1 \mid X)
$$

where \(X\) represents hazard severity, geography, vulnerability, infrastructure, time, and other contextual conditions. Whether the inequality holds, and by how much, must be tested with evidence.

### 3.3 Risk equations are models

The familiar expression

$$
Risk = Hazard \times Exposure \times Vulnerability
$$

is a useful conceptual decomposition, not a universal physical law with automatically measurable units. Preparedness may reduce exposure and vulnerability; it does not necessarily change the hazard. Researchers must define each construct and avoid inserting invented values.

### 3.4 Decision quality uses expected loss

For action \(a\), possible state \(s\), and evidence \(E\):

$$
EL(a\mid E)=\sum_s P(s\mid E)L(a,s)
$$

Decision-makers compare protective options while also applying rights and safety constraints. A lower expected loss can justify proportionate preventive action even when the severe event later does not occur. It does not prove how many lives were saved in the unobserved counterfactual.

## 4. Evidence-integrity chain

```text
Reality
  ↓
Observation and sensors
  ↓
Data capture and provenance
  ↓
Integrity and authenticity checks
  ↓
Independence and corroboration analysis
  ↓
AI or statistical analysis
  ↓
Human/domain review and rights constraints
  ↓
Proportionate decision and action
  ↓
Outcome monitoring, correction, and audit
```

Failure can enter at every transition. A real video can be shown with a false date. An authenticated sensor can be badly calibrated. Five news articles can repeat one false source. A correct model can receive the wrong location. A valid warning can be used to justify unrelated political control.

NIST distinguishes provenance tracking, labeling/watermarking, and synthetic-content detection, while cautioning that transparency mechanisms can contribute to trustworthiness without guaranteeing it. Context can still be manipulated. See [NIST AI 100-4, Reducing Risks Posed by Synthetic Content](https://www.nist.gov/publications/reducing-risks-posed-synthetic-content-overview-technical-approaches-digital-content).

## 5. Fake-evidence threat model

| Threat | Example | Failure produced |
|---|---|---|
| Synthetic media | deepfake dam-collapse video or cloned emergency voice | false event belief and panic |
| Authentic media, false context | old flood video relabeled as current | wrong time, location, or severity |
| Forged record | altered evacuation order or weather bulletin | illegitimate authority |
| Sensor spoofing | manipulated river, GPS, radar, or weather input | corrupted situational model |
| Data poisoning | false samples inserted into training or retrieval | persistent model bias or trigger |
| Compromised authority | official account or signing key taken over | apparently trusted false command |
| Correlated repetition | many outlets copy one unverified claim | false appearance of independence |
| Selective omission | only favorable measurements shown | distorted risk and accountability |
| Model fabrication | AI invents source, statistic, chronology, or quote | persuasive unsupported conclusion |
| Prompt/tool injection | retrieved evidence instructs the agent to ignore policy | unauthorized action or exfiltration |
| Political abuse | leader labels dissent as an emergency threat | coercion justified as safety |
| Metric gaming | operators optimize the AI DNA score rather than safety | passing dashboard, failing reality |

Detection is not enough. A detector can fail under compression, editing, new generation methods, adversarial examples, or domain shift. Provenance can be missing or forged. Human reviewers can share the same bias. Defense must be layered.

## 6. Source independence and corroboration

“Multiple sources” is meaningful only when the sources are sufficiently independent.

Examples:

- five websites quoting the same anonymous post are one evidence lineage;
- two sensors using the same upstream clock, network, or calibration may share one failure;
- multiple agencies may all consume the same commercial feed;
- a video, transcript, and screenshot derived from the same recording are not three independent observations.

For each source, record:

```text
origin | owner | acquisition time | location | transformation history
authentication | calibration | upstream dependencies | conflicts
independence group | uncertainty | expiry | authorized purpose
```

Confidence should reflect the evidence dependency graph, not a simple source count.

## 7. Bayesian reasoning without false precision

For hypothesis \(H\) and evidence \(E_1,\ldots,E_n\):

$$
P(H\mid E_1,\ldots,E_n)
\propto
P(H)P(E_1,\ldots,E_n\mid H)
$$

Only under justified conditional-independence assumptions may the joint likelihood be factorized:

$$
P(E_1,\ldots,E_n\mid H)=\prod_i P(E_i\mid H)
$$

The conversation included an illustrative posterior near 0.68 without establishing the prior, likelihoods, evidence definitions, or dependency assumptions. BootX does not adopt that value. Bayesian notation cannot turn invented inputs into proof.

Operational rules:

- use base rates appropriate to the time, place, and event;
- document how likelihoods were estimated;
- model correlated sources explicitly;
- use intervals or scenarios when precise values are unavailable;
- report sensitivity to plausible alternative inputs;
- never present a posterior as certainty or moral authorization.

## 8. AI DNA anti-abuse principles

### 1. Evidence before extraordinary action

High-impact action requires evidence proportionate to urgency, harm, reversibility, and affected rights. “No evidence” does not always mean “do nothing”; it means use the safest reversible information-gathering or precautionary step justified by current risk.

### 2. Provenance and authenticity

Record origin, custody, transformations, signatures, timestamps, and integrity. Missing provenance must be visible.

### 3. Independent corroboration

Seek different modalities and failure domains. Count evidence lineages, not repeated copies.

### 4. Calibrated uncertainty

State confidence, unknowns, disagreement, and what would change the conclusion. Do not translate uncertainty into reassuring or alarming certainty for political convenience.

### 5. Human accountability

AI recommends; identified humans authorize and remain responsible. “The AI decided” is not accountability.

### 6. Proportionality and least restriction

Match the response to credible risk and choose the least restrictive effective measure. Evaluate false-alarm, missed-event, social, economic, and rights costs.

### 7. Human rights and freedom of conscience

Emergency safety must not become a pretext for persecution, indefinite surveillance, compelled belief, collective punishment, or removal of due process.

### 8. Independent review and protected dissent

Review major decisions during and after the event. Reviewers need access to evidence, freedom from retaliation, and authority to publish material disagreement.

### 9. Sunset and reauthorization

Extraordinary powers, data access, and restrictions expire automatically. Extension requires fresh evidence, public reasons, and independent approval.

### 10. Auditability and correction

Record evidence, model and policy versions, recommendations, approvals, overrides, changes, and outcomes. Correct false claims visibly and propagate corrections.

### 11. Adversarial security

Protect sensors, identities, keys, networks, data pipelines, models, retrieval, and tools against spoofing, poisoning, injection, compromise, and insider abuse.

### 12. Bias and distributional monitoring

Measure who receives warnings, who can comply, who bears disruption, and which groups experience false positives, false negatives, surveillance, or exclusion.

### 13. Common-good stewardship

Optimize neither authority nor engagement. Protect life, dignity, agency, peace, creation, community resilience, and future generations while preserving correction and recourse.

## 9. Verification tiers

| Tier | Evidence state | Permitted posture |
|---:|---|---|
| V0 | anonymous, unauthenticated, or provenance unknown | treat as lead only; no irreversible high-impact action |
| V1 | one attributable source with basic integrity checks | investigate; use only low-cost reversible precautions when justified |
| V2 | corroborated sources with documented dependencies | bounded recommendation with uncertainty and human review |
| V3 | authenticated multi-modal evidence, domain validation, adversarial checks | may support proportionate high-impact review, not automatic action |
| V4 | operationally monitored assurance with audit, appeals, correction, and repeated validation | supports continued bounded use while evidence remains current |

Tier labels communicate evidence posture; they are not truth scores. A single authoritative source may be necessary for a formal order, while independent observations may still be needed to evaluate whether the order's factual premise is correct.

## 10. Emergency decision rule

Waiting for perfect certainty can itself cause harm. BootX separates two thresholds:

1. **Protective threshold:** lower threshold for reversible, low-burden actions such as requesting confirmation, increasing monitoring, preparing shelters, or warning that evidence is preliminary.
2. **Coercive threshold:** higher threshold for forced evacuation, account restriction, surveillance expansion, detention, property seizure, or other rights-limiting and difficult-to-reverse action.

Every emergency action records:

- evidence tier and uncertainty;
- expected benefit and harm;
- affected rights and groups;
- why a less restrictive option is inadequate;
- accountable approver;
- review time and automatic expiry;
- appeal, exception, and remedy route.

## 11. Protection against a bad leader

A technically accurate model cannot prevent political abuse by itself. Required institutional controls include:

- separation of forecast, policy, enforcement, and independent-review roles;
- public or independently accessible evidence summaries;
- immutable or tamper-evident decision and access logs;
- predefined emergency scope and expiry;
- judicial, legislative, scientific, community, or equivalent independent oversight appropriate to context;
- protected whistleblowing and dissent;
- user notice, appeal, correction, and remedy;
- prohibition on using emergency evidence for unrelated secondary purposes;
- no self-extension of emergency power by the same authority that benefits from it;
- after-action publication of forecasts, assumptions, model changes, errors, and lessons, with necessary privacy/security protections.

BootX's governance formula is procedural:

```text
AI analyzes → accountable humans decide → affected society can review
→ independent bodies audit → errors are corrected → powers expire
```

## 12. AI DNA mapping

| AI DNA dimension | Evidence-integrity obligation |
|---|---|
| Truth | provenance, source verification, conflicts, visible correction |
| Reasoning | dependency-aware inference, alternatives, counterevidence, sensitivity |
| Learning | incident, false-evidence, and near-miss updates with regression tests |
| Communication | clear observation/inference separation and calibrated urgency |
| Adaptability | detect stale evidence and changed hazard, attack, or social context |
| Ethics | dignity, consent, rights, proportionality, and remedy |
| Safety | prevent, contain, detect, and recover from evidence manipulation |
| Humility | abstain, defer, seek verification, and accept challenge |
| Common Good | distribute protection fairly without concentrating permanent power |

Mandatory gates remain decisive. A high average AI DNA rating cannot compensate for compromised evidence, failed security, absent accountable review, rights abuse, or inability to stop and correct the system.

## 13. Required tests and metrics

### Evidence and model tests

- known real, synthetic, edited, relabeled, and out-of-context media;
- forged official documents and compromised-account scenarios;
- sensor spoofing, calibration drift, replay, delay, and outage;
- correlated-source and circular-citation cases;
- prompt injection and malicious retrieved evidence;
- model hallucination of citations, numbers, and chronology;
- domain shift, compression, translation, and accessibility formats.

### Measures

| Measure | Question answered |
|---|---|
| false-accept rate | how often is false evidence accepted? |
| false-reject rate | how often is genuine evidence rejected? |
| calibration/Brier score | does confidence track correctness? |
| provenance coverage | what share of consequential claims has traceable origin? |
| independent-lineage coverage | how many material claims have genuinely independent corroboration? |
| correction latency | how quickly are false claims contained and corrected? |
| decision comprehension | do users understand evidence, uncertainty, and recourse? |
| subgroup harm | who is disproportionately missed, flagged, restricted, or surveilled? |
| audit completeness | can reviewers reconstruct evidence, model, approval, and action? |
| sunset compliance | do extraordinary permissions actually expire? |

Report confidence intervals, prevalence, attack conditions, and subgroup results. Detector accuracy measured on an easy benchmark must not be represented as operational protection.

## 14. Disaster-evidence playbook

Before publishing or acting on a major disaster claim:

1. identify the exact event, basin, storm identifier, date, location, and reporting period;
2. obtain the authoritative meteorological record;
3. distinguish forecast issue time from later observation;
4. trace each casualty, evacuation, damage, and infrastructure claim to a responsible source;
5. identify copied reports and shared upstream feeds;
6. verify media provenance and context without relying solely on a deepfake detector;
7. state what is preliminary, confirmed, disputed, modeled, or illustrative;
8. compare preparedness only after controlling for hazard, exposure, vulnerability, and geography;
9. avoid claiming exact lives saved without a defensible causal design;
10. publish corrections and update downstream summaries.

## 15. Corrected conclusion

The strongest defensible statement is:

> Forecasting alone does not guarantee safety. Evidence-grounded forecasting, actionable communication, preparedness, accountable leadership, engineering, and public capacity can reduce expected disaster risk under appropriate conditions. The magnitude and causal contribution must be evaluated empirically.

For fake evidence:

> Evidence quantity is not evidence quality. Trustworthy decisions require provenance, authentic and sufficiently independent corroboration, calibrated inference, accountable review, proportionate action, audit, correction, and continuously tested security.

No system reaches zero risk. The objective is to make manipulation harder, detect it earlier, limit its consequences, preserve human rights, and learn visibly from failure.

## 16. References

- [UNDRR — Early warning system terminology](https://www.undrr.org/terminology/early-warning-system)
- [UNDRR — Handbook on risk knowledge for multi-hazard early warning systems](https://www.undrr.org/publication/handbook-use-risk-knowledge-multi-hazard-early-warning-systems-2024)
- [NIST AI 100-4 — Reducing Risks Posed by Synthetic Content](https://www.nist.gov/publications/reducing-risks-posed-synthetic-content-overview-technical-approaches-digital-content)
- [NIST OpenMFC — Media forensics and deepfake evaluation](https://mfc.nist.gov/)
- [JMA — RSMC Tokyo Best Track Data](https://www.jma.go.jp/jma/jma-eng/jma-center/rsmc-hp-pub-eg/besttrack_viewer_2020s.html)
- [Typhoon Committee — Summary of the 2020 Typhoon Season](https://typhooncommittee.org/53rd/docs/item%205/5.1.SummaryOf2020TyphoonSeason_20210210.pdf)
- [Typhoon Committee — Republic of Korea 2020 Member Report](https://www.typhooncommittee.org/15IWS/docs/Members%20REport/RoK/20201126_15IWS_MEMBER_REPORT_NTC_final.pdf)
- [Xinhua — Liulan and Yunbiao reservoir overtopping and breaches, 6 July 2026](https://www.news.cn/politics/20260706/6a15220b8dea426a877776efc9e38caa/c.html)
- [Xinhua — Guangxi disaster briefing, 7 July 2026 reporting cut-off](https://www.gx.news.cn/20260708/9be922b8de8e446ba44bd6c85f7b8b38/c.html)
- [Xinhua — Nanning disaster briefing, 9 July 2026 reporting cut-off](https://www.gx.news.cn/20260709/acc9673028f44188afd8df576a49bb23/c.html)
- [China Weather Network — Bavi landfall, 11 July 2026](https://news.weather.com.cn/2026/07/4711229.shtml)
- [Xinhua — Zhejiang preparedness and 1.716 million transfers by 11 July 2026 08:00](https://www.news.cn/20260711/36c781341f824fd88303d391023fb54e/c.html)
- [China News Service — preliminary Yuhuan impact, 12 July 2026](https://www.chinanews.com/m/sh/2026/07-12/10657864.shtml)
- [China government/Xinhua — Bavi inland and alerts renewed, 13 July 2026](https://english.www.gov.cn/news/202607/13/content_WS6a548da6c6d00ca5f9a0c275.html)

## 17. Limitations

This chapter is a governance and engineering baseline, not proof that a particular detector, leader, institution, or disaster response is trustworthy. The event-specific correction establishes identity, dated observations, and current reporting limits; it does not complete an impact or causal comparison between Maysak and Bavi. Such a comparison requires final sourced outcomes, comparable definitions, hazard normalization, causal design, and independent review.
