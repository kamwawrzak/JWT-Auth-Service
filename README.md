# JWT-Auth-Service

## Run locally
```
make build && make run
```

## Examples
**Signup new user** 
```
curl -i -X POST localhost:8089/signup -d '{"email": "test@example.com", "password": "123", "password_repeat": "123"}'
```
**Login user (get JWT)**  
```
curl -i -X POST localhost:8089/login -d '{"email": "test@example.com", "password": "123"}'
```
