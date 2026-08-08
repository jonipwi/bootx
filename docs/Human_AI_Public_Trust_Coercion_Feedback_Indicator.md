# Human + AI Governance Monitor — Public Trust & Coercion Feedback Indicator
## Supplement to the National Common-Good Governance Index (NCGI)

**Version:** 0.1  
**Purpose:** Early-warning monitoring of declining institutional trust, grievance escalation, and coercive state response.

---

# 1. Why This Indicator Exists

A country can appear stable while public trust is quietly deteriorating.

The critical transition is often not simply:

> people are unhappy

but:

> people no longer believe that normal institutions can fairly resolve grievances.

When that happens, a local incident can escalate more quickly because:

```text
Incident
  ↓
Official explanation
  ↓
Public distrust
  ↓
Collective gathering / protest
  ↓
Security response
  ↓
Greater distrust
  ↓
Next incident escalates faster
```

This is the **Public Trust / Coercion Feedback Loop**.

The objective is not to predict revolution or label a government as illegitimate.

It is to detect when the relationship between citizens and institutions is becoming less correctable and more dependent on coercion.

---

# 2. Indicator Name

## PTCF — Public Trust & Coercion Feedback Index

The PTCF score ranges from:

```text
0 = healthy institutional trust / low coercive escalation risk
100 = severe trust breakdown / high coercive escalation risk
```

A high score is an **early-warning signal**, not proof of regime collapse.

---

# 3. Core Variables

Each variable is scored from `0` to `10`.

| Code | Variable | Meaning |
|---|---|---|
| G | Grievance intensity | How serious and widespread the underlying grievance is |
| D | Distrust of official explanation | Degree to which public rejects institutional accounts |
| M | Mobilization | Size, repetition, and organization of collective action |
| C | Coercive response | Degree of force, detention, intimidation, censorship, or security deployment |
| I | Institutional accessibility failure | Difficulty obtaining fair review, appeal, investigation, court access, or remedy |
| T | Transparency deficit | Lack of timely, verifiable, independently reviewable information |
| R | Repetition / contagion | Whether similar incidents are recurring or spreading |
| V | Vulnerable-person impact | Degree to which children, elderly, civilians, detainees, minorities, or other vulnerable groups are affected |

---

# 4. Mathematical Formula

```text
PTCF = 10 × (
  0.15G +
  0.15D +
  0.15M +
  0.15C +
  0.10I +
  0.10T +
  0.10R +
  0.10V
)
```

Weights sum to 1.00, so:

```text
0 ≤ PTCF ≤ 100
```

---

# 5. Interpretation Bands

| PTCF | Interpretation |
|---:|---|
| 0–19 | Low risk / normal institutional grievance handling |
| 20–39 | Watch level |
| 40–59 | Elevated trust-friction |
| 60–74 | High escalation risk |
| 75–89 | Severe institutional trust breakdown |
| 90–100 | Extreme coercion / legitimacy crisis warning |

These are research bands, not scientifically validated thresholds.

---

# 6. Evidence Confidence

Every PTCF score must include a separate confidence score.

```text
CONF = SourceQuality × SourceDiversity × EvidenceCoverage
```

Each factor ranges from `0` to `1`.

Example:

```text
SourceQuality = 0.8
SourceDiversity = 0.7
EvidenceCoverage = 0.6

CONF = 0.8 × 0.7 × 0.6
     = 0.336
     = 33.6%
```

A dramatic incident with only social-media clips may therefore produce:

```text
PTCF = 70
Confidence = 25%
```

This means the pattern looks concerning, but the factual basis is incomplete.

---

# 7. Four Evidence Classes

### FACT
Supported by credible documentation, official records, reputable journalism, court records, or multiple independent sources.

### TESTIMONY
Reported firsthand experience or family/public witness account.

### ALLEGATION
A claim circulating publicly but not independently established.

### INTERPRETATION
An analytical or political conclusion drawn from known events.

This distinction is mandatory.

---

# 8. Event-Level Scoring Sheet

| Variable | Score 0–10 | Evidence | Confidence |
|---|---:|---|---|
| G — grievance intensity |  |  |  |
| D — official-explanation distrust |  |  |  |
| M — mobilization |  |  |  |
| C — coercive response |  |  |  |
| I — institutional accessibility failure |  |  |  |
| T — transparency deficit |  |  |  |
| R — repetition/contagion |  |  |  |
| V — vulnerable-person impact |  |  |  |

Then calculate:

```text
PTCF_event
```

---

# 9. Country-Level PTCF

One dramatic event should not define an entire country.

For event `i`:

```text
E_i = PTCF_i × Confidence_i × SeverityWeight_i
```

Then for a monitoring period:

```text
National PTCF =
Σ(E_i) / Σ(Confidence_i × SeverityWeight_i)
```

This reduces the influence of low-confidence rumors.

---

# 10. Time Decay

```text
AdjustedEventWeight =
OriginalWeight × e^(-λt)
```

Where:

- `t` = time since incident
- `λ` = decay constant

For a one-year half-life:

```text
λ = ln(2) / 365
```

Repeated incidents prevent the risk score from naturally decaying.

---

# 11. Momentum Indicator

```text
ΔPTCF = PTCF(t) - PTCF(t-1)
```

Interpretation:

```text
ΔPTCF < 0  → trust conditions improving
ΔPTCF ≈ 0  → stable
ΔPTCF > 0  → trust/coercion risk worsening
```

---

# 12. Acceleration

```text
A_PTCF =
[PTCF(t) - PTCF(t-1)]
-
[PTCF(t-1) - PTCF(t-2)]
```

If:

```text
PTCF > 50
AND
ΔPTCF > 0
AND
A_PTCF > 0
```

then:

```text
ESCALATION REVIEW = TRUE
```

---

# 13. Critical Red-Line Triggers

Trigger **CRITICAL REVIEW** when credible evidence shows one or more of:

- deliberate killing of peaceful demonstrators,
- mass arbitrary detention,
- enforced disappearance,
- collective punishment,
- punishment of relatives for another person's acts,
- deliberate use of civilians as human shields,
- hostage-taking,
- systematic denial of legal remedy,
- widespread internet shutdown specifically to conceal severe abuses,
- security forces operating without meaningful accountability.

These are not mathematically cancelled out by positive economic indicators.

---

# 14. Integration with NCGI

Suggested integration:

```text
NCGI_adjusted =
NCGI_base - 0.15 × max(0, PTCF - 40)
```

Cap the penalty:

```text
Maximum PTCF penalty = 10 NCGI points
```

Suggested interpretation:

```text
PTCF < 40
→ no national penalty

PTCF 40–59
→ mild warning

PTCF 60–79
→ significant warning

PTCF ≥ 80
→ maximum warning penalty + Critical Review
```

---

# 15. Pharaoh-Pattern Interpretation

```text
Leadership fears instability
        ↓
Control increases
        ↓
Institutional trust decreases
        ↓
People bypass formal channels
        ↓
Collective resistance increases
        ↓
Security response intensifies
        ↓
Public distrust deepens
        ↓
Leadership sees greater threat
        ↓
Control increases again
```

The indicator does not say:

> “This leader is Pharaoh.”

It says:

> “The measurable governance relationship is moving toward the Pharaoh-pattern feedback loop.”

---

# 16. Worked Example — Henan-Style Incident

This is an **illustrative example only**, not a final factual score for the July 30 event.

Suppose later-verified evidence supports:

```text
G = 7
D = 8
M = 7
C = 6
I = 6
T = 7
R = 5
V = 5
```

Then:

```text
PTCF =
10 × (
0.15×7 +
0.15×8 +
0.15×7 +
0.15×6 +
0.10×6 +
0.10×7 +
0.10×5 +
0.10×5
)

PTCF = 65
```

Interpretation:

```text
65 / 100
= High escalation risk
```

Suppose evidence remains incomplete:

```text
SourceQuality = 0.75
SourceDiversity = 0.60
EvidenceCoverage = 0.55
```

Then:

```text
Confidence =
0.75 × 0.60 × 0.55
= 24.75%
```

Responsible AI output:

> **PTCF: 65/100 — High risk, but only 25% confidence. Further independent evidence required.**

It must not say:

> “China is definitely entering rebellion.”

---

# 17. Fear Suppression Correction

Low visible protest does not necessarily mean high trust.

Define:

```text
FS = Fear Suppression score
```

0–10.

Then:

```text
Latent PTCF =
Observed PTCF + 2 × FS
```

Cap at 100.

Example:

```text
Observed PTCF = 45
Fear Suppression = 8

Latent PTCF = 61
```

This helps prevent the false conclusion:

> “There are few protests, therefore people trust the government.”

---

# 18. Human + AI Monitoring Output Format

```text
Country:
Date:

PTCF:
Trend:
Acceleration:
Fear Suppression:
Confidence:

Strongest signals:
1.
2.
3.

Counter-evidence:
1.
2.

Unknowns:
1.
2.

Critical red lines:
Yes / No

Interpretation:
Short evidence-based explanation.

Prediction:
Clearly labelled probability or scenario,
never presented as fact.
```

---

# 19. Anti-Propaganda Rule

For every high-risk event, AI must search for:

```text
Evidence supporting escalation
VS
Evidence supporting institutional resolution
```

Both must be evaluated.

This prevents the indicator from becoming a political confirmation machine.

---

# 20. Core Human + AI Principle

The desired direction is:

```text
Grievance
  ↓
Transparent investigation
  ↓
Independent review
  ↓
Fair remedy
  ↓
Public confidence restored
```

rather than:

```text
Grievance
  ↓
Distrust
  ↓
Suppression
  ↓
Resistance
  ↓
Greater suppression
```

---

# 21. Inheritance Principle

> **A stable society is not one in which nobody complains.**

> **A stable society is one in which people believe complaints can be heard without fear, investigated without manipulation, and corrected without violence.**

For the Human + AI era:

> **AI should not merely measure whether citizens obey. It should help measure whether institutions remain worthy of trust.**

And:

> **The strongest early-warning signal is often not anger itself, but the disappearance of belief that truth and justice can still be obtained through normal institutions.**
