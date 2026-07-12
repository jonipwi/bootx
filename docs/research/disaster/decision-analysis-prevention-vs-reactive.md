# Decision Analysis Study

## Prevention vs. Reactive Rescue in Disaster Management

**Version:** 1.0

**Status:** Analytical model, not an event-outcome report<br>
**Evidence maturity:** E1 — hypothesis and decision framework<br>
**Known limitation:** numerical inputs must be estimated for a defined population, hazard, action, and time window

> **Purpose:** Compare two policy approaches using decision theory,
> expected-value reasoning, and public-interest principles.

## Executive Summary

This study compares:

-   **Strategy A -- No additional advance precaution:** retain baseline
    readiness and respond after local effects become observable.
-   **Strategy B -- Selected preventive preparedness:** add a specified,
    proportionate action before impact based on credible evidence.

This is a policy comparison, **not** a claim that prevention should
always override every other concern. Effective policy seeks a
proportional balance between protecting lives, limiting unnecessary
disruption, and communicating uncertainty honestly.

------------------------------------------------------------------------

# Simplified Expected Loss Model

Let \(p\) be the probability of a severe event and \(1-p\) the probability
of a non-severe event. Let \(L_1,L_0\) be losses without the proposed
preparation in those states; \(R_1,R_0\) the corresponding residual losses
with preparation; and \(C\) the preparation cost incurred regardless of the
state.

$$
EL_{reactive}=pL_1+(1-p)L_0
$$

$$
EL_{preventive}=C+pR_1+(1-p)R_0
$$

Preparation has lower modeled expected loss when:

$$
C<p(L_1-R_1)+(1-p)(L_0-R_0)
$$

The earlier expression \(C<p(L-R)\) is only the special case where losses in
the non-severe state are assumed equal and normalized to zero. The expanded
form matters because false alarms, evacuation injury, lost income, warning
fatigue, privacy intrusion, and unequal burdens can occur even when the severe
event does not.

All terms must use a compatible consequence scale or a documented
multi-criteria method. Fundamental rights and prohibited actions remain
constraints; they are not simply priced away by an expected-value result.

------------------------------------------------------------------------

# Comparison

  -----------------------------------------------------------------------
  Dimension               Reactive ("Rescue After Preventive Preparedness
                          Harm")                  
  ----------------------- ----------------------- -----------------------
  Primary objective       Recover after losses    Reduce losses before
                                                  impact

  Human life              Baseline modeled risk   Lower only if the action is
                                                  effective and does not add
                                                  offsetting harm

  Infrastructure          Baseline modeled loss   Reduced loss only for hazards
                                                  the measure actually mitigates

  Economic continuity     Larger disruption after Up-front disruption,
                          disaster                potentially smaller
                                                  total disruption

  Public perception       May appear cheaper      May be criticized if
                          initially               impacts are less severe
                                                  than forecast

  Scientific basis        Responds after evidence Uses forecasts and risk
                          is visible              analysis

  Long-term resilience    Context-dependent       May improve through learning
                                                  and capacity; may decline if
                                                  repeated actions are harmful

  Trust requirement       Moderate                High (requires clear
                                                  communication of
                                                  uncertainty)
  -----------------------------------------------------------------------

------------------------------------------------------------------------

# Risk Matrix

  Scenario                      Reactive Risk          Preventive Risk
  ---------------------- -------------------- ------------------------
  False alarm              Low immediate cost   Moderate inconvenience
  Missed disaster          Potentially very high   Lower only if warning and
                                                  action are effective
  Mass casualties          Hazard-dependent        May reduce probability or
                                                  exposure; not guaranteed
  Long-term resilience     Context-dependent       Context-dependent

------------------------------------------------------------------------

# Leadership Evaluation

Leaders should not be rewarded or punished solely based on whether the
worst-case scenario occurred.

A stronger evaluation asks:

-   Was the decision based on the best available evidence?
-   Was the response proportionate to the risk?
-   Was uncertainty communicated honestly?
-   Were policies updated as new information arrived?
-   Were unnecessary harms minimized?

------------------------------------------------------------------------

# Typhoon Bavi research hypothesis

A warning followed by limited damage can reflect several
possibilities:

-   the storm weakened,
-   the track shifted,
-   preparations reduced consequences,
-   or the highest-risk areas were spared.

A smaller-than-feared impact does **not** by itself prove that preparedness
was unnecessary, effective, or causal. Those alternatives require forecast,
action, hazard, exposure, and outcome evidence. The 2026 event record remains
preliminary; see [Typhoon Bavi (2026)](typhoon-bavi-aftermath-summary.md).

------------------------------------------------------------------------

# Common Good Principle

The goal is not to maximize warnings or minimize warnings.

The goal is to **minimize expected harm while preserving public trust**.

Good governance combines:

-   evidence,
-   proportional action,
-   transparency,
-   continuous learning,
-   accountability.

------------------------------------------------------------------------

# Conclusion

Preventive preparedness can provide greater societal resilience when based on
credible evidence, effective measures, proportionate decisions, and review of
false-alarm and distributional harms.

However, maintaining trust requires explaining uncertainty, reviewing
outcomes after each event, and improving future forecasts.

> **Power without wisdom is dangerous.**
>
> **Knowledge with humility becomes a blessing.**
