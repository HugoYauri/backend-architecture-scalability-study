# Future Work

Ordered roughly from highest to lowest priority.

1. Identify the primary source. Locate the original video's title, author or channel, platform, and URL, and add a proper entry to references/references.bib. This is the single most important missing piece for treating this repository as citable.
2. Add automated load testing. Instrument the three demos in src/ with a tool such as k6, Locust, or wrk, define a fixed test protocol (duration, concurrency, hardware/environment), and store the raw output in data/raw/ and summarized metrics in results/metrics/, with figures in results/figures/.
3. Build a fuller microservices reference implementation. Extend the load-balancing demo into a small multi-service system (for example file service, notification service, auth service) behind a real API gateway, closer to Block 6 of the study guide, following the spirit of projects such as ftgo-application (see references/consulted-repositories.md).
4. Add a formal literature review. Complement the practitioner-oriented study guide with citations to peer-reviewed distributed systems literature and standard references (for example on consistency models, queueing theory for load balancing, and caching theory).
5. Threat model the authentication and gateway block. Block 7 of the guide describes JWT validation at the gateway; a future addition could analyze token revocation, key rotation, and gateway-as-single-point-of-failure risks in more depth.
6. Measure cache behavior empirically. Instrument the cache-aside demo to report real hit/miss ratios and latency distributions under a repeatable workload, rather than only demonstrating the control flow.
7. Independent technical review. Ask a second reviewer, ideally with backend/distributed-systems experience, to check docs/study-guide/ against the original source material for translation or interpretation errors.
8. Add a rendered architecture diagram. Convert figures/architecture-evolution-overview.md (currently a text/Mermaid description) into a rendered image once the diagramming approach is confirmed with the user.
