# API

## Ping

```
curl localhost:8000/ping
```

## Account

Post

```
curl -H 'Content-Type:application/json' -X POST -d '{"name":"test", "email":"test@exm.com", "password":"password"}' http://localhost:8000/api/register
```

## Auth

Post

```
curl -H  'Content-Type:application/json' -X POST -d '{"email":"test@exm.com", "password":"password"}' http://localhost:8000/api/login
```

## Task

```
curl -H 'Content-Type:application/json' -X GET http://localhost:8000/api/v1/tasks
```
