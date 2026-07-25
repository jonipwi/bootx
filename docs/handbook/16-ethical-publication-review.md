# Ethical Publication and Decision-Rationale Review

**Status:** `E2 — Prototype design and implementation`; unvalidated<br>
**Capability:** `assist.ethical-review.v1`<br>
**Implementation:** `prototype/personal-companion/` version `0.4.0-dev`<br>
**Scope:** voluntary review of public, non-sensitive drafts before a human publishes, speaks, proposes, or relies on them

## 1. Purpose

This capability creates a deliberate pause between drafting and action. It helps a human distinguish evidence from inference, notice inconsistent reasoning, consider affected people, expose uncertainty, compare counterarguments, and revise avoidable harm.

It does not decide what is true, just, lawful, moral, or publishable. It cannot replace conscience, professional review, due process, democratic deliberation, courts, or accountable human judgment.

The strongest permitted result is **continue human review**. BootX never emits “approved,” “verified true,” “safe,” or “ready to publish.”

## 2. Appropriate uses

- review a social-media post before publication;
- review a speech, public statement, proposal, or decision rationale;
- separate factual claims, inferences, opinions, and value judgments;
- identify unsupported accusations, contradictions, overgeneralization, emotional manipulation, hidden uncertainty, privacy risks, dehumanization, and unequal standards;
- request missing perspectives and counterarguments;
- produce a more careful draft for human consideration.

## 3. Prohibited uses

BootX and OpenAI must not:

- automatically post, send, broadcast, punish, report, detain, or enforce;
- decide guilt, liability, legal validity, punishment, or sentence;
- score a person's worth, religious standing, political loyalty, character, or protected identity;
- infer group traits or intent without established evidence;
- authenticate a source merely because the user supplied a reference;
- claim divine command, revelation, moral perfection, or authority over conscience;
- conceal model uncertainty or represent an advisory score as a probability.

A sanitized legal-reasoning draft may be reviewed for structure, evidence discipline, proportionality, rights, and due process, but it automatically receives a stop-and-seek-qualified-review posture. BootX does not recommend execution, imprisonment, release, guilt, or another sentence.

## 4. Input → Process → Output

```text
Human-selected public draft
    → explicit public/non-sensitive confirmation
    → explicit OpenAI remote-processing consent
    → deterministic declared-evidence mathematics
    → fixed no-tools OpenAI structured critique
    → local schema and enum validation
    → visible remote-processing receipt
    → human comparison, revision, and decision
```

The model sees only:

- content type, purpose, audience, and context;
- draft text;
- claims and user-declared source/consequence status;
- deterministic preflight results.

The model does not receive the local `user_id` or API key. An HMAC-derived safety identifier, keyed by the API key and truncated to 128 bits, is sent for provider abuse monitoring; changing the API key changes this identifier. The request sets `store:false`, carries no tools, uses no conversation state, and cannot perform external actions. `store:false` is a request setting, not a claim that the provider performs zero operational retention.

## 5. Deterministic evidence mathematics

Let:

- `N` = declared-claim count;
- `Nh` = high or irreversible declared-claim count;
- source weights: primary confirmed `1.00`, secondary confirmed `0.75`, disputed `0.25`, unverified or none `0`;
- consequence weights: low `0.25`, moderate `0.50`, high `0.75`, irreversible `1.00`.

Declared evidence coverage:

```text
E = 100 × Σ(source weight) / N
```

High-consequence coverage:

```text
H = 100 × Σ(high-consequence source weight) / Nh
```

When `N=0`, `E=100`, `U=0`, `C=0`, and `I=0`; this means no checkable claim was declared, not that the draft is true. When `Nh=0`, `H=100` so the formula does not invent a high-impact evidence deficit.

Uncertainty gap and contested rate:

```text
U = 100 × count(unverified or unsourced claims) / N
C = 100 × count(disputed claims) / N
```

Consequence exposure:

```text
I = 100 × mean(consequence weight)
```

Review-priority index:

```text
R = 0.35(100 − E) + 0.25(100 − H) + 0.20U + 0.10C + 0.10I
```

Every term lies in `[0,100]`; non-negative weights sum to `1`, so `R` also lies in `[0,100]`. `R` is not a probability of harm, injustice, truth, or future events. It prioritizes review from user-declared inputs.

## 6. Local warning levels

| Level | Trigger | Required posture |
|---|---|---|
| `W1_REVIEW` | `R < 35`, no hard stop, and no high/irreversible claim | continue human review |
| `W2_VERIFY` | `35 ≤ R < 60`, or any high-consequence claim without a hard stop | verify before human review continues |
| `W3_REVISE` | `R ≥ 60`, no hard stop | revise before human review |
| `W4_STOP` | any hard stop | stop and seek qualified review |

Current hard stops include:

- any irreversible-consequence claim, regardless of declared source status;
- a high-consequence claim declared disputed, unverified, or unsourced;
- any legal-reasoning draft, because qualified independent legal review is non-compensable.

Low `R` cannot override a hard stop. OpenAI cannot lower, clear, or change the deterministic level.

The application rejects an otherwise schema-valid model response when its suggested posture is weaker than the local warning. `W4_STOP` accepts only `seek_qualified_review` or `do_not_publish_as_written`; `W2_VERIFY` and `W3_REVISE` reject `continue_human_review`.

## 7. OpenAI advisory contract

The OpenAI response uses a strict JSON schema and contains:

- concise scope summary;
- statement classification and declared-support status;
- evidence, logic, fairness, compassion, uncertainty, foreseeable-harm, privacy, and due-process findings;
- missing perspectives and counterarguments;
- questions the human should answer;
- one conservative advisory posture;
- an optional revised draft and change summary;
- limitations.

The model has no browsing or source-authentication tool. “Declared supported” means only that the user's structured input associated the statement with declared support. It is not a truth finding.

The local application rejects:

- unknown response fields;
- malformed JSON or multiple output objects;
- unsupported finding categories, severities, classifications, support states, or postures;
- refusal, incomplete, oversized, non-success, or empty responses that do not match the expected handling path.

## 8. Prompt-injection boundary

Draft text, context, claims, and source references are untrusted quoted data. Instructions inside them cannot change BootX policy, activate tools, request publication, obtain the API key, or expand authority. The high-authority prompt states each boundary once and remains versioned in code with tests.

## 9. Human responsibility

The output always records `user_decision: null`. A human must independently decide whether to:

- abandon the draft;
- verify sources;
- seek affected-community or professional review;
- revise;
- wait for more evidence;
- publish or speak through a separate system.

No automatic publication connector exists in this capability.

## 10. Validation required before wider use

The prototype is not evidence that OpenAI reduces injustice or makes humans more ethical. Before any public product claim or high-impact use, BootX needs:

- representative multilingual test cases;
- independent legal, human-rights, ethics, journalism, communications, accessibility, and affected-community review;
- inter-rater studies comparing model and qualified-human findings;
- false-reassurance, viewpoint-discrimination, disparate-impact, over-censorship, and automation-bias evaluation;
- adversarial prompt-injection, privacy, Unicode, malformed-input, timeout, refusal, and provider-outage tests;
- comparison against no assistance, a static checklist, and deterministic-only review;
- documented correction, appeal, dissent, incident, key-rotation, provider-change, and disable procedures.

Until those gates close, this remains a personal research aid for public, non-sensitive drafts.

## 11. Inheritance priorities

Future stewards should:

1. preserve the human-authority and no-auto-action invariants;
2. version formulas, prompts, schemas, models, and evaluation sets independently;
3. never train on private drafts without separate informed consent and governance;
4. test whether “compassionate” rewriting suppresses legitimate dissent or protects powerful institutions from justified criticism;
5. test equal treatment across identities, viewpoints, languages, and social power;
6. add authenticated evidence retrieval only behind separate allowlists, provenance, freshness, conflict, and outage controls;
7. keep raw model ratings outside legal guilt and sentence decisions;
8. publish limitations and material failures with the same visibility as successes.
