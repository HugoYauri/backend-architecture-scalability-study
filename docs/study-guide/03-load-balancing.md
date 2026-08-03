# Block 3 - Load Balancing (7:54-12:45)

## Trigger

The app becomes popular and users complain about slowness.

## The load balancer

The load balancer is the intermediate piece that decides which server handles each request, and routes only to healthy machines (health checks).

## Balancing algorithms

| Algorithm | How it works | When to use it |
|---|---|---|
| Round Robin | Alternates servers sequentially | Homogeneous traffic |
| Weighted Round Robin | Round robin with weights per capacity | Machines of different power |
| Sticky Sessions | Returns the user to the same server | When the server holds session state |
| Least Connections | Picks the server with fewest active connections | Requests of variable duration |

## Important nuance (from the source)

Not all requests cost the same. Uploading a file (POST /upload) consumes much more than downloading it (GET /file/123). Because of this, plain round robin can be unfair.

## Extra capability

The load balancer can also route by path, for example sending POST /upload to a specialized file-handling service.
