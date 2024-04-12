# Example client  

## Run test client
1. `export JWT_SECRET_KEY=<secret key using for JWT generation>`
2. `make start-client`

**Example request** 
1.  Get JWT from `/login` endpoint of JWT-Auth-Service
2. `export $JWT=<jwt value>`
3. Send query to protected resource  
```
curl -H 'Authorization: Bearer $JWT' localhost:9999/protected
```
