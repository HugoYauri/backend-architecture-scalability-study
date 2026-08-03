# Block 10 - Caching and CDN (37:36-44:51)

## Prior warning (from the source)

Do not optimize prematurely. This block comes last for a reason.

## Problem

500 users open the same file every day, repeating the entire path: gateway -> file service -> database -> object storage.

## Cache (Redis)

A key-value store held in RAM, much faster than going to disk.

### What to cache and what not to

- Cache: file metadata (name, properties) since it changes rarely or never.
- Do not cache: the file's bytes (RAM is expensive and scarce; Redis is not built for large blobs).

Exception noted in the source: some companies (for example Netflix) do cache content in RAM for initial spikes around premieres.

### Cache-aside pattern (worth memorizing)

Does the key "file:123" exist in the cache?

- Yes: return it.
- No: fetch it from the database, then store it in the cache, then return it.

## CDN

A network of hundreds of points of presence distributed geographically, outside your cluster, specialized in large assets.

- The first user in Lisbon triggers the full round trip and the CDN caches the file.
- The next 499 users receive it from the nearby node: about 20 ms instead of about 1 second.
- A user in the United States avoids the roughly 300 ms of latency to Lisbon.

## Key difference (from the source)

Redis caches small data inside the system; the CDN caches large assets at the edge, close to the user.
