# Block 4 - Nginx Demonstration (12:45-14:58)

## Setup described in the source

Three identical services in Go behind an Nginx proxy, with two upstream blocks:

- Route /rr: round robin, alternating between server 1, 2, and 3 on refresh.
- Route /sticky: session affinity, a session cookie always returns you to the same server; deleting the cookie causes reassignment.

## Concept to retain

The session cookie (which could be your authentication token) is what enables session affinity.

## Relation to this repository's reproducible demo

src/01-load-balancing-demo/ reproduces this idea with multiple stateless Go servers behind Nginx. See its README for exact routes and how it differs, in scope, from the original conceptual demo (for example, this repository does not assume any specific commercial sticky-session module; see that demo's README for the implementation approach actually used and its limitations).
