# Block 7 - Authentication and Authorization (23:35-27:14)

## Validation flow (normal request)

1. The client sends the Authorization header with a JWT.
2. The gateway validates the signature and expiration without network calls.
3. If invalid, it returns 401 Unauthorized.
4. If valid, it routes to the destination service.

## Issuance flow (login)

1. The user sends credentials.
2. The gateway asks the user service to verify they exist.
3. The auth service signs a token with a private key.
4. It returns the token to the client.

## Critical distinction

| Concept | Question it answers |
|---|---|
| Authentication | Who are you? Do you have access to the application? |
| Authorization | What can you do? Do you have permission to upload files? |

Authorization can live distributed across each service, with its own permission levels.
