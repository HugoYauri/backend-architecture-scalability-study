# Block 8 - File Upload and Delivery (27:14-32:33)

## Anti-pattern

Uploading files directly to your application server or to the relational database. A 200 MB video saturates the server, causes timeouts, and opens the door to resource-exhaustion attacks.

## Correct pattern: presigned URL

1. The client requests to upload a file, sending only metadata (name, size, type).
2. Gateway -> file service -> relational database stores the metadata row (ID, name, size, type); never the file itself.
3. The service asks the object storage (for example S3 or Google Cloud Storage) for a temporary upload URL.
4. That URL has a short expiration and a size limit (for example 50 MB).
5. The client uploads the file directly to the bucket, without touching your infrastructure.

## Why it works

Heavy traffic never crosses your servers. This solves the original slowness problem in uploads.

## New complication

When the file reaches the bucket, actions need to be triggered (generate a thumbnail, notify, synchronize). The bucket cannot know about all your services.
