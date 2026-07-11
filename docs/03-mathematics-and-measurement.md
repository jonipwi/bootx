# Mathematics and Measurement Framework

## 1. Measurement principles

Mathematics clarifies assumptions; it does not manufacture objective truth from subjective ratings. Every metric needs a construct definition, unit, data source, population, time window, missing-data rule, uncertainty estimate, and decision purpose.

Use a dashboard before a single score. A composite index can summarize, but must never hide a failing safety-critical dimension.

## 2. Decision under uncertainty

Let \(a\) be an action, \(s\) a possible state of the world, \(P(s\mid I)\) the probability of that state given information \(I\), and \(L(a,s)\) the loss caused by action \(a\) in state \(s\).

$$
\operatorname{ExpectedLoss}(a)=\sum_s P(s\mid I)L(a,s)
$$

Choose the action with the lowest expected loss only after applying rights and safety constraints. Some actions remain forbidden even if a crude utility calculation appears favorable.

For the repository's prevention example:

$$
EL_{reactive}=pL
$$

$$
EL_{preventive}=C+pR
$$

where \(p\) is event probability, \(L\) unprepared loss, \(C\) preparation cost, and \(R\) residual loss after preparation. Prevention has lower expected loss when:

$$
C<p(L-R)
$$

This equation is correct only if inputs cover all material impacts, including unequal burdens, warning fatigue, opportunity cost, privacy, and long-term trust.

## 3. Precaution and tail risk

Expected loss can conceal rare catastrophic outcomes. For high-consequence decisions, also report:

- worst credible case;
- probability above an unacceptable-loss threshold;
- reversibility and recovery time;
- distribution of harm across groups;
- conditional tail expectation, where justified;
- confidence ranges and model uncertainty.

Precaution does not mean “always intervene.” It means the evidence burden and safeguards increase with plausible irreversible harm.

## 4. Forecast quality

### Calibration

Of events assigned 70% probability, approximately 70% should occur over a sufficiently large, relevant sample. Calibration plots should be reported by context and subgroup.

### Brier score

For binary outcomes with predicted probability \(p_i\) and outcome \(y_i\in\{0,1\}\):

$$
BS=\frac{1}{N}\sum_{i=1}^{N}(p_i-y_i)^2
$$

Lower is better. Compare against simple baselines and report uncertainty.

### Classification measures

For protective alerts, report at minimum:

$$
Precision=\frac{TP}{TP+FP}
\qquad
Recall=\frac{TP}{TP+FN}
$$

Also report specificity, negative predictive value, prevalence, alert rate per user, time to action, and cost-weighted errors. Accuracy alone is misleading for rare scams or emergencies.

## 5. Human outcome measures

The companion is intended to strengthen people, so system metrics are insufficient.

| Construct | Example operational measure | Guardrail |
|---|---|---|
| Comprehension | correct explanation of reason and options after an alert | reading level and accessibility |
| Calibration | gap between confidence and correctness | no confidence inflation |
| Agency | informed-choice score; override and exit success | no dark patterns |
| Independent skill | unassisted performance after training period | detect dependency |
| Safety benefit | verified loss or exposure avoided versus comparator | monitor displacement of harm |
| Privacy | data minimization, retention compliance, deletion verification | zero unauthorized secondary use |
| Equity | error and benefit differences across relevant groups | investigate material disparity |
| Trustworthiness | appropriate reliance under known limits | avoid maximizing reported trust |

The goal is **appropriate trust**, not maximum trust. Users should rely on the companion when evidence supports reliance and question it when uncertainty or limits are material.

## 6. AI DNA measurement model

Each AI DNA dimension in the next document receives a 0–4 maturity rating:

| Rating | Meaning |
|---:|---|
| 0 | absent or contradicted |
| 1 | stated intent only |
| 2 | implemented control with limited internal evidence |
| 3 | tested with representative users or adversarial cases |
| 4 | independently evaluated and continuously monitored |

Every rating must include evidence, evaluator, date, context, uncertainty, known failures, and expiry. If evidence is missing, the rating is “not established,” not an optimistic midpoint.

### Mandatory gates

Regardless of average score, release is blocked when any applicable gate fails:

- critical security control;
- meaningful consent and exit;
- high-impact human review;
- incident response and rollback;
- protected-group harm review;
- memory deletion and access control;
- truthful identity and limitation disclosure.

### Profile before composite

Report the dimension vector:

$$
D=(T,R,L,C,A,E,S,H,G)
$$

for truth, reasoning, learning, communication, adaptability, ethics, safety, humility, and common good. A radar chart may aid discussion but does not prove validity.

If governance requires a summary, use a transparent weighted mean only after gates pass:

$$
M=\frac{\sum_i w_i d_i}{4\sum_i w_i}\times 100
$$

Publish all \(w_i\), dimension values, confidence ranges, and sensitivity analysis. Never compare unlike contexts as if they form a universal league table.

## 7. Reframing ACGI as a research dashboard

The original AI Companion Gravitation Index is preserved as an inspiring metaphor, but its ratio should not be used operationally. Replace it with two independently visible profiles.

### Protection capacity

$$
P=(H_p,T_r,S_f,A_u,G_v)
$$

covering human protection, warranted trust, safety performance, autonomy, and governance capacity.

### Capability pressure

$$
Q=(C_a,D_p,X_r,U_n)
$$

covering AI capability, dependency/replacement pressure, extraction/manipulation risk, and uncertainty.

Normalize indicators to \([0,1]\) only after defining anchors from data. Display trends, confidence intervals, and subgroup distributions.

For exploratory research—not release authorization—a log balance may be calculated:

$$
B=\sum_i \alpha_i\ln(P_i+\epsilon)-\sum_j\beta_j\ln(Q_j+\epsilon)
$$

where weights sum to one on each side and \(\epsilon\) is a disclosed small constant. This avoids division by zero and makes assumptions explicit, but the result remains model-dependent. Thresholds must be empirically calibrated; the repository's original 0.7, 1.0, and 1.5 bands are not validated.

## 8. Trust is not a popularity score

Model warranted reliance for task \(k\) and context \(c\) as a profile:

$$
W_{k,c}=f(Reliability, Competence, Integrity, Transparency, Recourse, Context)
$$

Do not infer it solely from satisfaction. A charming system may receive high ratings while being unsafe. Behavioral measures should test whether users accept correct advice, reject incorrect advice, seek help appropriately, and understand uncertainty.

## 9. Experimental design

Before a study, record:

1. research question and causal hypothesis;
2. population, inclusion, exclusion, and recruitment;
3. comparator and assignment method;
4. primary and secondary outcomes;
5. harm and stop criteria;
6. sample-size rationale;
7. subgroup analyses;
8. data handling and consent;
9. analysis plan and missing-data rules;
10. deviations, adverse events, and publication commitment.

Use confidence intervals and effect sizes, not significance labels alone. Replicate across contexts. Separate exploratory analysis from confirmatory evaluation.

## 10. Worked anti-scam example

Suppose 1,000 test messages include 100 verified scams. The system flags 120 messages: 90 scams and 30 legitimate messages.

```text
TP = 90, FP = 30, FN = 10, TN = 870
Precision = 90 / 120 = 75%
Recall = 90 / 100 = 90%
Specificity = 870 / 900 ≈ 96.7%
```

These numbers are incomplete. Researchers must also measure whether users understand the alert, whether scammers adapt, which groups experience false positives, whether users safely verify messages, and whether repeated alerts cause habituation.

## 11. Measurement anti-patterns

- assigning precise percentages from narrative impressions;
- changing the metric after seeing the result;
- optimizing engagement or self-reported trust as the primary goal;
- averaging away a catastrophic subgroup failure;
- claiming causation from before/after observation alone;
- hiding missing data or model changes;
- choosing thresholds without cost, rights, or calibration analysis;
- treating an AI evaluator as independent evidence.
