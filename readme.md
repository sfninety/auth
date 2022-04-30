# Authentication mechanisms we are implementing

1. Entropy Bucket - random bytes that we store alongside each JTI in the JWT to try and ensure the request is valid. Even if the signing key is leaked, an attacker wouldn't be able to create a valid JWT because they would need a whitelisted JTI and an Entropy Bucket value to pair with it.

    . JTI is going to be encryption(sub + random bytes)

2. Refresh token rotation - New refresh token on each refresh.

3. Token reuse detection - Most recently issued refresh token should be invalidated if the token before it is reused.

4. Both access token and refresh token should be hashed/encrypted.



# Auth flow

1. After successful authentication through the mobile app, an AUTHZ jwt is provided to the user. 

2. Upon sending this to the `/token` endpoint, two jwts will be returned, an ACCTOK and a REF. The access token can be used to make a request, and until its expiry it can be used to make requests.

3. We check if the access tokens JTI has been used before - check if it has been used before. If it has, invalidate the request. Then, check if the entropy value is correct. 

    - We store all JTIs that have been used.
    - We store token information separately, like `sub` and `exp`, alongside and entropy value.
    - A correct entropy value will map to a token, which will also tell us WHO the user is and if the token is expired.
    

4. Once the access token is expired, the user can refresh using the refresh token and the previous refresh token will be moved from the whitelist to the blacklist. A new pair of ACCTOK and REF will be returned.