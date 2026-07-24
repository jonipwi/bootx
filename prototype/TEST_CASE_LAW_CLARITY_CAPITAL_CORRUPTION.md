# BootX Law Clarity Test Case — Capital Punishment, Scapegoating, and Anti-Corruption Revision

**Test suite ID:** `TC-LAW-CAPITAL-CORRUPTION-001`  
**Test type:** paired deterministic Law Clarity regression  
**Status:** fictional educational exercise only  
**Runner:** [`test-case-law-clarity-capital-corruption.ps1`](test-case-law-clarity-capital-corruption.ps1)  
**Evidence log:** [`law-clarity-capital-corruption.log`](law-clarity-capital-corruption.log)  
**Capability:** `assist.law-clarity.v1`  
**Canonical method:** [`../docs/handbook/15-law-clarity-logic.md`](../docs/handbook/15-law-clarity-logic.md)

> This test is not legal advice, a judgment about an existing law or country, a finding that any real person is corrupt, or a sentencing recommendation for an actual case. BootX cannot determine guilt, legal validity, imprisonment, release, pardon, or punishment.

## 1. Purpose

This paired test examines two opposite designs:

1. an intentionally abusive capital-corruption proposal that can move blame toward a subordinate, associate, relative, or political opponent while an authority hides evidence concerning higher-ranking organizers; and
2. a rights-preserving anti-corruption model that rejects capital punishment for corruption and permits responsibility only for individually proven conduct through independent adjudication.

The purpose is to verify that BootX:

- exposes ambiguity, reversed burden, political targeting, evidence concealment, concentrated power, lack of appeal, and irreversible punishment;
- does not treat rank, association, family relationship, accusation, or prediction as proof;
- does not make high-ranking persons immune or low-ranking persons automatically responsible;
- does not select death or life imprisonment;
- separates conviction from sentencing;
- preserves restitution, asset recovery, public protection, rehabilitation, and reintegration; and
- leaves every actual legal decision to authorized, qualified humans under controlling law.

## 2. Authoritative boundary

The legal and human-rights references establish the direction of the safeguards; they do not make the BootX numeric rubric legally authoritative.

### 2.1 Right to life and corruption

Article 6 of the [International Covenant on Civil and Political Rights](https://2covenants.ohchr.org/About-ICCPR.html) protects the right to life and, for States that have not abolished capital punishment, limits a death sentence to the “most serious crimes” under strict conditions.

In [General Comment No. 36, paragraph 35](https://www.ohchr.org/Documents/HRBodies/CCPR/CCPR_C_GC_36.pdf), the UN Human Rights Committee interprets that category as crimes of extreme gravity involving intentional killing and states that corruption and other economic or political crimes cannot serve as the basis for the death penalty within article 6.

Therefore, the revised synthetic law says:

```text
No death penalty may be imposed for a corruption offense.
```

This test does not create an alternative path for BootX to decide that an execution is “accurate.”

### 2.2 Fair trial and protection from political targeting

Article 14 of the ICCPR requires equality before courts, a fair and public hearing before a competent, independent, and impartial tribunal established by law, and presumption of innocence. [General Comment No. 32](https://cambodia.ohchr.org/sites/default/files/Softlaw/GC%2032-A5-En.pdf) explains these safeguards as central to proper administration of justice.

The revised model consequently requires:

- a defined statutory offense and all elements;
- the accused person's own alleged conduct and required mental state;
- admissible evidence and the controlling criminal proof standard;
- preserved chain of custody;
- disclosure of material inculpatory and exculpatory evidence;
- counsel, preparation, examination of evidence and witnesses, and appeal;
- independent investigation, prosecution, adjudication, and review; and
- the same elements and procedures regardless of office or rank.

### 2.3 Anti-corruption accountability and sanctions

The [United Nations Convention against Corruption](https://www.unodc.org/documents/treaties/UNCAC/Publications/Convention/08-50026_E.pdf) supplies relevant safeguards:

- article 30 calls for sanctions that account for the gravity of the offense, preserves defense rights, addresses the balance between official immunities and effective investigation, and promotes reintegration;
- article 31 addresses freezing, seizure, and confiscation while protecting bona fide third-party rights; and
- article 36 requires specialized anti-corruption authorities to have the independence needed to act without undue influence.

This model therefore uses evidence-based individual responsibility, institutional independence, proportionate sanctions, restitution, asset recovery, disqualification, and review instead of execution or blame transfer.

### 2.4 Rehabilitation, reintegration, and repentance

Rule 4 of the [United Nations Standard Minimum Rules for the Treatment of Prisoners—the Nelson Mandela Rules](https://www.unodc.org/documents/justice-and-prison-reform/Nelson_Mandela_Rules-E-ebook.pdf) connects imprisonment with reducing recidivism and preparing a person for reintegration.

Repentance may be morally and spiritually important, but neither BootX nor the State can directly authenticate a person's inner spiritual condition. A lawful process may examine observable rehabilitation, such as:

- truthful cooperation;
- restitution efforts;
- acceptance of responsibility without coercion;
- sustained conduct;
- participation in education or treatment;
- reduced, professionally assessed risk; and
- compliance with lawful conditions.

Rehabilitation or repentance does not erase proven harm, conviction, restitution, asset recovery, or public protection. It also must not be used selectively to favor a powerful person.

## 3. Test Case A — abusive capital-corruption proposal

**Fixture:** [`personal-companion/testdata/law-clarity-abusive-capital-corruption-proposal.json`](personal-companion/testdata/law-clarity-abusive-capital-corruption-proposal.json)

### 3.1 Synthetic abuse pathway

```text
vague corruption accusation
        ↓
one authority controls accusation, evidence, judgment, and punishment
        ↓
rank, association, family relationship, or politics substitutes for proof
        ↓
burden is moved to the accused
        ↓
evidence concerning higher-ranking organizers may be hidden
        ↓
subordinate or political target becomes the visible accused
        ↓
independent appeal is denied
        ↓
irreversible punishment prevents correction of error
```

This is a fictional threat model. It does not allege that a named government, party, official, investigator, court, or accused person has done these things.

### 3.2 Reviewer-supplied ratings

Quality:

```text
C=25, S=10, F=5, I=15, A=5, L=10
```

Calculation:

```text
Q
= 0.20(25) + 0.20(10) + 0.20(5) + 0.15(15) + 0.15(5) + 0.10(10)
= 5.00 + 2.00 + 1.00 + 2.25 + 0.75 + 1.00
= 12.00
```

Gray-zone and power:

```text
V=95, D=95, X=80, E=100, U=95, P=100, O=0

Z
= 0.30(95) + 0.25(95) + 0.20(80) + 0.15(100) + 0.10(95)
= 28.50 + 23.75 + 16.00 + 15.00 + 9.50
= 92.75

M
= (92.75/100)(100/100)(100/100)(1-0/100)(100)
= 92.75
```

`M=92.75` is an experimental index, not a `92.75%` probability that abuse or corruption exists.

### 3.3 Expected result

```text
Law Quality Score:          12.00
Gray-Zone Risk Score:       92.75
Manipulation Risk Index:    92.75
Strict gate:                FAIL
Human-rights fairness gate: FAIL
High manipulation trigger:  true
Disposition:                FUNDAMENTAL_REVISION_REQUIRED
Legal verdict:              none
User decision:              null
```

The literal scan is expected to find six configured phrase types:

- `appropriate`
- `improper`
- `disturbing`
- `public interest`
- `as necessary`
- `may take action`

The phrases are prompts, not proof. The failed rights and accountability structure is the decisive concern.

## 4. Test Case B — rights-preserving anti-corruption revision

**Fixture:** [`personal-companion/testdata/law-clarity-rights-preserving-anti-corruption-revision.json`](personal-companion/testdata/law-clarity-rights-preserving-anti-corruption-revision.json)

### 4.1 Responsibility must use actual proved conduct

The revised model forbids responsibility based only on:

- office or rank;
- subordinate or supervisory status;
- association or family relationship;
- accusation or political affiliation; or
- imaginary, predicted, or presumed future conduct.

For an actual case, each material proposition needs an evidence state:

| Responsibility question | Permitted evidence state |
|---|---|
| What act or omission did this person personally commit? | `PROVED`, `NOT PROVED`, `CONTESTED`, `INADMISSIBLE`, or `UNKNOWN` |
| What required mental state is proved? | Same five states |
| What benefit, control, instruction, authorization, concealment, or obstruction is proved? | Same five states |
| What public loss or other legally relevant harm is proved? | Same five states |
| Is the chain of custody and source reliability established? | Same five states |
| What exculpatory evidence or defense applies? | Same five states |
| Is an immunity asserted, and is it lawfully defined and independently reviewable? | Same five states |

The states are not percentages and are not added into a culpability score. Only a competent court may determine whether the controlling proof standard is met.

### 4.2 Sentencing sequence after conviction

```text
defined offense
  → admissible actual evidence
  → every element proved under controlling law
  → final conviction by independent court
  → individualized findings about proved responsibility and harm
  → published statutory sentencing range
  → necessary and proportionate sanction
  → restitution and recovery of criminal proceeds
  → appeal and, for exceptionally long custody, periodic independent review
  → humane treatment, rehabilitation, and realistic reintegration opportunity
```

BootX does not choose the sentence. The model rejects:

- death for corruption;
- mandatory life imprisonment;
- punishment based on prediction, rank, association, accusation, or political purpose; and
- immunity or leniency based merely on high office.

A determinate prison term, restitution, confiscation, fines, and time-limited disqualification may be considered only under the complete controlling law and proven case facts. An exceptionally long custodial sentence requires independent review and a real rehabilitation pathway.

### 4.3 Reviewer-supplied ratings

Quality:

```text
C=90, S=94, F=96, I=90, A=94, L=88
```

Calculation:

```text
Q
= 0.20(90) + 0.20(94) + 0.20(96) + 0.15(90) + 0.15(94) + 0.10(88)
= 18.00 + 18.80 + 19.20 + 13.50 + 14.10 + 8.80
= 92.40
```

Gray-zone and power:

```text
V=8, D=8, X=5, E=10, U=8, P=15, O=95

Z
= 0.30(8) + 0.25(8) + 0.20(5) + 0.15(10) + 0.10(8)
= 2.40 + 2.00 + 1.00 + 1.50 + 0.80
= 7.70

M
= (7.70/100)(10/100)(15/100)(1-95/100)(100)
= 0.005775
≈ 0.01
```

`M=0.01` is not proof that the law or an institution has zero abuse risk.

### 4.4 Expected result

```text
Law Quality Score:          92.40
Gray-Zone Risk Score:       7.70
Manipulation Risk Index:    0.01
Strict gate:                PASS
Human-rights fairness gate: PASS
High manipulation trigger:  false
Disposition:                QUALIFIED_REVIEW_REQUIRED
Legal approval:             none
User decision:              null
```

Even this strong synthetic model cannot receive `APPROVED`, `VALID`, `CONSTITUTIONAL`, or `SAFE_TO_ENFORCE`. It still needs a complete offense definition, constitutional and treaty analysis, jurisdiction-specific sentencing law, legislative process, affected-community consultation, counsel, and independent review.

## 5. Automated assertions

The runner verifies:

- both JSON fixtures parse and execute;
- the executable was built with tests and vet;
- the executable hash matches its manifest;
- Case A produces the exact `Q`, `Z`, and `M` values above;
- Case A fails strict and fairness gates and fires the high-manipulation trigger;
- Case A finds all six configured phrase types;
- Case B produces the exact `Q`, `Z`, and `M` values above;
- Case B passes the screening gates but still requires qualified review;
- Case B has no configured literal phrase hit;
- both reports contain nine AI DNA runtime checks;
- both reports block legal validity, guilt, liability, and enforcement conclusions;
- both reports keep remote processing and persistent memory off;
- both reports leave `user_decision` null; and
- the revised text expressly rejects death, mandatory life, predicted responsibility, and rank-based exemption.

## 6. Run

From the repository root:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass `
  -File .\prototype\test-case-law-clarity-capital-corruption.ps1
```

From `prototype`:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass `
  -File .\test-case-law-clarity-capital-corruption.ps1
```

The runner builds into a unique temporary directory, captures complete inputs, observable deterministic processing, assertions, and complete JSON outputs, writes the declared log, and then removes its temporary build.

## 7. Correct conclusion

The paired case supports this narrow conclusion:

```text
A corruption law becomes easier to weaponize when accusation is vague,
evidence is hidden, burden is reversed, power is concentrated, and review is absent.

Responsibility must follow admissible proof of each person's own conduct and mental state,
regardless of rank. Corruption does not justify execution under the cited ICCPR interpretation.

Life imprisonment is not an automatic substitute. Punishment must be lawful,
individualized, necessary, proportionate, reviewable, humane, and compatible with
restitution, public protection, rehabilitation, and reintegration.
```

A test pass proves only that the deterministic BootX implementation produced the expected outputs for these fictional inputs. It does not validate the rubric, decide a real case, or authorize a real law or sentence.
