# Synthesis and Suggested Study Path

## Synthesis: the connecting thread

| Stage | Piece added | Problem it solves |
|---|---|---|
| 1 | Database | Persistence and shared state |
| 2 | Load balancer | Traffic distribution |
| 3 | Horizontal/vertical scaling | Capacity |
| 4 | Microservices | Organizational scalability (teams) |
| 5 | API Gateway + VPC | Single entry point and security |
| 6 | Object storage + presigned URLs | Handling heavy traffic (downloads) |
| 7 | Event broker | Resilience in service-to-service communication |
| 8 | Cache + CDN | Latency and cost |
| 9 | Rate limiting | Protecting resources |

Note: this numbering follows the source guide's own synthesis table; it does not map one-to-one to the Block 0-11 numbering used elsewhere in docs/study-guide/, since the synthesis groups some blocks together.

## Author's conclusion (from the source)

Architecture is not only infrastructure. It involves algorithms, product decisions, and the choice of database type. It means understanding the problem and the domain before building the solution.

## Suggested study path (from the source)

1. Fundamentals (Blocks 1-3): set up a server with a database and an Nginx instance balancing two instances. This is the most cost-effective exercise if you are just starting out.
2. Scaling (Blocks 4-5): understand when to scale horizontally versus vertically before touching microservices.
3. Distributed systems (Blocks 6-9): do not go further here without mastering the previous stages; complexity grows quickly.
4. Optimization (Blocks 10-11): last, and only once you have metrics that justify the optimization.
