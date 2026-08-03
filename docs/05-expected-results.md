# Expected Results

## Current status

No empirical results currently exist in this repository. The source material (the study guide) is conceptual and pedagogical: it explains mechanisms and trade-offs, and uses illustrative numbers to make ideas concrete, but it does not report a measured experiment. Accordingly, results/ contains only placeholders and instructions, not data.

## What running the demos would produce, if instrumented

If the demos in src/ were extended with a load-testing tool (see docs/06-future-work.md), the following artifacts would be the expected outputs. None of these exist yet; this section documents the plan, not a result.

1. Load-balancing demo: request distribution counts per backend instance under round-robin routing versus a hashed/sticky routing strategy, for example as a bar chart saved to results/figures/ and a raw counts table saved to results/metrics/.
2. Cache-aside demo: cache hit ratio and average response latency for repeated reads of the same key, comparing a cold cache to a warm cache.
3. Rate-limiting demo: the request/response sequence showing when responses switch from 200 OK to 429 Too Many Requests as the configured threshold is crossed, and how the counter resets after the time window elapses.

## How to interpret any numbers you generate yourself

If you run the demos and collect your own measurements, please store them under data/raw/ and results/metrics/ with a clear description of your hardware, network conditions, and how many repetitions you ran, so that anyone reading results/ can judge how much weight to give the numbers. Do not present single, unrepeated runs on a laptop as general performance claims.
