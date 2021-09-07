# GO fiber jwt mysql authentication example
```
docker-compose up --build -d
```
## Register
```
curl localhost:8080/register \
    -i \
    -H "content-type:application/json" \
    -d '{"name":"test","email":"a@b.com","password":"password"}'

```
## Login
```
curl localhost:8080/login \
    -iv \
    -c - 'http://localhost:8080/login' \
    -H "content-type:application/json" \
    -d '{"email":"tko@nida.ac.th","password":"password"}' \
```
## Get User with Cookies
```
curl localhost:8080/user \
    -iv \
    --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzExMTA2MjUsImlzcyI6IjYifQ.jgwlkjI6L-TSUXxx0zxFT4yeWGM3JMiwQPG6PoBUDjs" \
```
## Logout with cookies
```
curl localhost:8080/logout \
    -X POST \
    -iv \
    --cookie "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzExMTA2MjUsImlzcyI6IjYifQ.jgwlkjI6L-TSUXxx0zxFT4yeWGM3JMiwQPG6PoBUDjs" \
```
