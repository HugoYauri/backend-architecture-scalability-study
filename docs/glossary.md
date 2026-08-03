# Glossary

Short, plain-language definitions of terms used throughout docs/study-guide/. These reflect how the terms are used in the source material, not a formal academic definition.

- Stateful server: a server that stores application data (for example, files) inside itself, so that losing the server means losing the data.
- Stateless server: a server that keeps no durable application data of its own; all durable data lives in an external store (for example, a database), so any instance can serve any request.
- Load balancer: a component that receives incoming requests and distributes them across multiple backend servers, typically also checking server health.
- Health check: a periodic check the load balancer performs to confirm a backend server is alive before routing traffic to it.
- Sticky session (session affinity): a load-balancing behavior where a given client is consistently routed to the same backend server, usually via a cookie.
- Horizontal scaling: adding more machines to a cluster to handle more load.
- Vertical scaling: increasing the resources (CPU, RAM, disk) of an existing machine.
- Autoscaling: automatically provisioning or removing instances based on observed load.
- Microservice: an independently deployable service responsible for a narrow part of a system's domain.
- API Gateway: a single entry point that receives client requests, and routes them to the correct backend microservice, optionally aggregating responses.
- VPC (virtual private network): a private network boundary in which internal services are not directly reachable from the public internet, except through a controlled entry point such as a gateway.
- Authentication: verifying who a user is.
- Authorization: verifying what an authenticated user is allowed to do.
- JWT (JSON Web Token): a signed token format commonly used to carry authentication information without a database lookup on every request.
- Presigned URL: a temporary, expiring URL that grants direct upload or download access to an object storage bucket without routing the file's bytes through the application server.
- Message broker: an intermediary system (for example Kafka or RabbitMQ) that lets services communicate asynchronously through published events instead of direct calls.
- Dead letter queue: a holding area for messages that could not be delivered or processed successfully, usually paired with alerting.
- Fan-out: distributing a single event to multiple independent subscribers.
- Cache-aside pattern: a caching strategy where the application checks the cache first, and on a miss reads from the source of truth and then populates the cache.
- CDN (Content Delivery Network): a geographically distributed network of servers that cache and serve large or frequently requested assets close to the end user.
- Rate limiting: restricting how many requests a client can make in a given time window, to protect a system from overuse or abuse.
