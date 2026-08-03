# Block 6 - Microservices and API Gateway (17:38-23:35)

## Triggers

More engineers hired, plus slow endpoints.

## Domain decomposition (Google Drive example)

- file service: uploads and downloads
- notification service: push to web and mobile
- auth service: authentication
- real-time service: synchronization across devices

One team per service. Warning from the source: microservices are a complex topic; they typically only make sense for large organizations.

## Why the load balancer alone is not enough

The question arises: who decides that POST /api/login goes to the auth service? And where does authentication live? This is where the API Gateway appears.

- It receives the request, analyzes it, and routes it to the correct service.
- It can aggregate responses from several services into one.
- It is the single point of entry to the system.

## VPC (virtual private network)

All services stay within a private network, without exposing ports externally. Only the gateway is accessible.

## Risk introduced

The gateway becomes a single point of failure, so it also needs to be load balanced and scaled.
