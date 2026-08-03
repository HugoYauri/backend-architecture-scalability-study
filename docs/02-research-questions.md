# Research Questions

Each question below is grounded in a specific block of the study guide (referenced in parentheses) and is answerable, at least partially, using the material already present in docs/study-guide/.

1. RQ1 - Architecture and scale mismatch (Block 0). Why can an architecture that works for a startup with 10 users be the wrong one for the same product at 10,000 users? What early over-engineering risks does the guide warn against?

2. RQ2 - State decoupling (Blocks 1-2). What specific failure modes does keeping state inside a single server introduce (data loss on crash, inability to scale horizontally), and how does moving state to an external database resolve them? What new responsibilities does this split create?

3. RQ3 - Load balancing algorithm selection (Blocks 3-4). Given that not all requests have the same cost (for example, file upload vs. file download), how should the choice among round robin, weighted round robin, sticky sessions, and least connections be justified?

4. RQ4 - Scaling strategy (Block 5). Under what conditions is horizontal scaling preferable to vertical scaling, and how do autoscaling and serverless approaches change the operational responsibilities of a team?

5. RQ5 - Microservices decomposition (Block 6). What complexity (operational, organizational) does splitting a monolith into microservices behind an API gateway introduce, and under what organizational conditions does the guide suggest it becomes justified?

6. RQ6 - Authentication vs. authorization boundary (Block 7). How does the guide distinguish who are you (authentication) from what can you do (authorization), and what are the trade-offs of validating a JWT at the gateway without a network call per request?

7. RQ7 - Large file handling (Block 8). Why is routing large file uploads through the application server and relational database considered an anti-pattern, and how does the presigned-URL pattern change where load is placed in the system?

8. RQ8 - Asynchronous communication guarantees (Block 9). What guarantees (availability, durability, retries, dead-letter queues, fan-out) does introducing a message broker provide over direct synchronous calls between services?

9. RQ9 - Cache vs. CDN (Block 10). What is the qualitative difference between what should be cached in an in-memory store (for example, Redis) versus what should be served through a CDN, and why?

10. RQ10 - Rate limiting as a reuse of caching infrastructure (Block 11). How can the same caching layer used for performance also be reused to implement rate limiting, and what does this protect against?

Note on scope: these questions are answered at the level of the conceptual source material (a synthesis or documentary level), not through new empirical experiments. See docs/03-methodology.md and docs/04-limitations.md.
