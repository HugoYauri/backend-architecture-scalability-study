# Methodology

This study combines two complementary methods.

## 1. Documentary synthesis

1. The full original document (Spanish) was read in its entirety.
2. Each of the 12 blocks (0-11), the synthesis table, and the suggested study path were translated into English and reorganized as individual Markdown files under docs/study-guide/, preserving: all concepts, definitions, and rules to remember stated in the source; all tables (problem/consequence, algorithm/behavior/use case, etc.); all illustrative examples and numbers, explicitly labeled as illustrative, not empirical.
3. No content was added that is not present in, or a direct logical consequence of, the source material.

## 2. Reproducible hands-on demonstrations

For the three blocks of the guide that describe a concrete, implementable mechanism (load balancing, caching, rate limiting), a minimal Docker-based demo was built under src/.

| Demo | Guide block(s) | What it reproduces |
|---|---|---|
| 01-load-balancing-demo | Blocks 3-4 | Multiple identical stateless servers behind an Nginx reverse proxy, comparing round-robin-style distribution against a hashed/sticky routing strategy |
| 02-cache-aside-demo | Block 10 | The cache-aside pattern: check cache, on miss read from source of truth, populate cache, return |
| 03-rate-limiting-demo | Block 11 | A per-client request counter with a time window, returning 429 Too Many Requests past a threshold |

These demos are illustrative, not benchmarks. They are designed to let a reader observe the qualitative behavior described in the guide (for example, a sticky client keeps hitting the same backend), not to produce statistically valid performance numbers. Producing real performance numbers would require a load-testing tool (for example k6, Locust, wrk), a defined hardware and environment specification, and multiple repeated runs, none of which exist yet in this repository (see docs/06-future-work.md).

## Threats to validity and honesty notes

- The translation from Spanish to English was done manually by an AI assistant under human supervision; any translation error should be reported via an issue.
- No independent expert review of the technical content has been performed yet.
