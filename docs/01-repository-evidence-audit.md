# Repository Evidence Audit

**Audit date:** 2026-07-12  
**Scope:** all files present in the repository root at the audit date

## 1. Purpose

This audit prevents an important category error: treating a vision document, illustrative case score, screenshot, or compiled artifact as proof of a complete and validated system. It records what the current repository actually supports and identifies what future work must verify.

## 2. Artifact inventory and status

### Boot artifacts

| Artifact | Observed fact | Evidence status |
|---|---|---|
| `bootx.img` | 1,542,656-byte raw artifact; begins with x86 boot code; bytes 510–511 are the BIOS signature `55 AA`; embedded strings name stage1, stage2, kernel v0.1, FAT12, shell, and network behavior | Directly inspected artifact |
| `bootx.iso` | 3,325,952-byte artifact containing ISO `CD001` identification and BootX strings | Directly inspected artifact |
| `bootx.vdi` | 3,145,728-byte artifact containing the BootX payload and VirtualBox disk container data | Directly inspected artifact |
| `bootx.jpg` | Screenshot showing BootX Kernel v0.1 in VirtualBox, E1000 initialization, DHCP exchange, ARP, and ping reply | Visual evidence of one reported run, not a reproducible test |

SHA-256 hashes recorded during the audit:

| File | SHA-256 |
|---|---|
| `bootx.img` | `53AD76FDB61F12FBF9167777FDE92A75076695FCED11FF87B8B027E475DBF6AD` |
| `bootx.iso` | `76C751870FCAE1904B25820F1978033A514DF4ED16D430D33C7F89757F30600D` |
| `bootx.vdi` | `45551492AAD3F00EBD24069000E8E4701CF98A3E32F146198155CBA94D299218` |
| `bootx.jpg` | `B24C1237D24B7C26BF40C49E61DC75A56225C143950F50915F004CA7C1DD68CC` |

These hashes identify the audited artifacts; they do not certify that the binaries are secure.

### Missing reproducibility inputs

The README and install guide refer to files such as `Makefile`, `scripts/`, `bootloader/stage1/boot.asm`, blueprint and progress plans, a kernel, and a UEFI loader. Those inputs are not present in this checkout. Consequently:

- the binaries cannot be rebuilt from the repository as supplied;
- compiler, assembler, dependency, and build provenance cannot be independently verified;
- claims about internal memory safety, network correctness, and source-level behavior cannot be audited here;
- README build commands are not currently reproducible;
- the supplied images must be treated as opaque experimental artifacts, not trusted deployment media.

The highest technical priority is to restore source, build scripts, dependency versions, license information, tests, and a reproducible build record.

## 3. Conceptual corpus map

| Theme | Primary source documents | Contribution | Current limitation |
|---|---|---|---|
| Self-correction and power | `100-year-study-pharaoh-style.md` | warns against concentrated authority, suppressed correction, and narrative capture | philosophical pattern; needs operational governance indicators |
| Human capability | `intelligence-benchmark.md` | combines reasoning, emotion, adaptability, agency, and wisdom | additive expression is conceptual, not a validated psychometric scale |
| Trustworthy companionship | `ai-survival-plan.md`, `trust-model-plan.md`, `trustworthy-ai-plan.md`, `practical-and-scale.md` | establishes guardian, memory, autonomy, humility, and local trust concepts | overlaps heavily; lacks requirements, acceptance tests, and evidence |
| Cybersecurity | `cybersecurity-reformed.md`, `local-hybrid-cybersecurity.md` | proposes explanation, human approval, shared intelligence, and local control | approval for every event can cause fatigue; architecture and threat model are incomplete |
| Civilization balance | `ACGI-case-formula.md`, `civilization-ai-survival.md` | frames capability, protection, trust, autonomy, and extraction risk | metaphor and score thresholds are unvalidated; geopolitical symbolism should not be treated as causal science |
| Disaster decisions | five Typhoon Bavi documents | introduces probability, prevention, expected loss, uncertainty communication, and outcome review | event claims require sourced verification; scores are explicitly illustrative and cannot validate AI DNA |
| Space hypothesis | `WATER-MAGNETIC-FEASIBILITY.md` | carefully distinguishes water as resource/energy carrier and magnetism as enabling technology | broad long-range hypothesis; engineering claims require literature review and experiments |
| Safety framing | `DISCLAIMER.md` | emphasizes research status, correction, and human responsibility | disclaimers do not replace actual safety controls, governance, or validation |

## 4. Strong ideas worth preserving

- Truth must remain discoverable and correction must be rewarded.
- Intelligence without direction and accountability can cause harm.
- Protection and freedom are a design tension, not mutually exclusive slogans.
- Local ownership and global/shared intelligence can be combined through bounded interfaces.
- Forecasts and warnings should be judged by decision quality under uncertainty, not outcome alone.
- The system should explain limitations and enable informed human action.
- The most vulnerable users should be first-class design participants.
- Metaphors can inspire research, but physics, causality, and measurement must remain explicit.

## 5. Claims requiring caution or revision

### ACGI

The original formula multiplies several 0–10 ratings and divides by others. It has no validated measurement model, can divide by zero, is highly sensitive to a single rating, mixes unlike constructs, and assigns thresholds without empirical calibration. It is useful as a discussion map, not as a civilization prediction or decision rule. A safer replacement is developed in the mathematics document.

### AI DNA scores

Scores such as 93.3% or 96/100 are arithmetic applied to expert-like judgments, not empirical accuracy measures. Weights, raters, evidence rules, uncertainty, and inter-rater reliability were not established. Future reports must publish those elements and retain the dimension-level profile rather than implying biological or moral precision.

### Disaster case material

The documents contain public-event claims without citations in the repository. They also use a counterfactual—what would have happened without preparation—that cannot be directly observed. The decision-theory lesson is valuable, but factual timelines, impacts, dates, and attributions must be source-verified before academic or policy use.

### “Trust” and “common good”

These are not single measurable substances. Trust is context-specific willingness to accept vulnerability based on evidence and recourse. The common good involves contestable values and distributional effects. Both require plural participation and cannot be defined solely by developers or an AI.

### Emotional support

An “emotional stabilizer” may help reflection, but it also creates dependency, manipulation, crisis-response, and boundary risks. The system must not impersonate human intimacy or substitute for qualified care.

## 6. Evidence maturity labels

Every future document and feature should carry one label:

| Label | Meaning | Required language |
|---|---|---|
| E0 — Vision | value or desired future | “We intend…” |
| E1 — Hypothesis | falsifiable proposed relationship | “We hypothesize…” |
| E2 — Prototype | mechanism demonstrated in a limited environment | “The prototype demonstrated…” |
| E3 — Controlled evidence | evaluated against a comparator with predefined measures | “Under these test conditions…” |
| E4 — Field evidence | evaluated with representative users and real workflows | “In this population and context…” |
| E5 — Operational assurance | monitored deployment with audit, incident response, and independent review | “Operational evidence shows…” |

Do not use “proven,” “safe,” “trusted,” or “benefits humanity” without identifying the population, context, comparator, measures, time period, uncertainty, and remaining harms.

## 7. Immediate repository actions

1. Restore all source and reproducible build inputs for the BootX artifacts.
2. Add licenses for software, documentation, and third-party components.
3. Add a provenance manifest and signed release process.
4. Preserve original vision papers as historical design inputs; use this handbook for operational definitions.
5. Add citations and source dates to empirical case studies.
6. Correct legacy character-encoding corruption in a reviewed, traceable migration.
7. Require research protocols and preregistered metrics before presenting new numerical scores.
8. Establish issue, correction, incident, and decision-log processes.

