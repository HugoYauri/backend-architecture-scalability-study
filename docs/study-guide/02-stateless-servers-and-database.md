# Block 2 - Stateless Servers and Database (5:26-7:54)

## Solution

Decouple the data from the server by moving it to an external relational database.

## What is gained

- Persistence: the server can die without losing data.
- Horizontal scalability: multiple instances share a single source of truth.
- The data-synchronization problem between instances is eliminated.

## Separation of responsibilities (first appearance)

- The server is responsible for serving the user.
- The database is responsible for storing the data.
- The server knows about the database; the database does not know about the server.
