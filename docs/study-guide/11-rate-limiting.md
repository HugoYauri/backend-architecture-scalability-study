# Block 11 - Rate Limiting (44:51-end)

## Objective

Prevent malicious users from exhausting your resources, driving up costs, and degrading the experience for everyone else.

## How it works (reusing the existing cache)

1. A layer at the gateway identifies the user (by IP or another identifier).
2. It writes a counter in Redis: user:123 -> 5 requests in the last minute.
3. If it exceeds the threshold (for example 10 per minute), it responds with 429 Too Many Requests.

## Connection to the real world (from the source)

Daily token limits on services such as Claude or ChatGPT are a variant of this same idea: expensive-to-compute resources, rationed per user.

## Additional use of the cache

Storing the results of costly computations (compilations, transformations) instead of recalculating them.
