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

login

```
curl -c cookie.txt -H  'Content-Type:application/json' -X POST -d '{"email":"test@exm.com", "password":"password"}' http://localhost:8000/api/login
```

logout

```
curl -H "Authorization: Bearer  <cookie.txtのTOKEN貼り付け>" -H 'Content-Type:application/json' -X DELETE http://localhost:8000/api/v1/logout
```

curl -H "Authorization: eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjo5LCJlbWFpbCI6InRlc3RAZXhtLmNvbSIsImV4cCI6MTY5MjAxNzI4NiwiaWF0IjoxNjkyMDE1NDg2LCJpc3MiOiJnaXRodWIuY29tL2tvcnktanAvdnVlX2dvL2FwaSIsImp0aSI6ImIyZmE5Y2QwLThlNGQtNGRkNi04Y2Y0LWJkMmVlZDNkNTc5NiIsInN1YiI6ImFjY2Vzc190b2tlbiJ9.m06JvUYRmwJFnqkfWFccegKTWJR_wg3j_Qd34AFbW-dMxXeKvt22m67mEOhMP-pv3gGPU90yijReenY78gTVHF2Ger-g_p7vrAbaKDPuIANl2h03i0N7LAWgthsqCdxc8KX19k_8G-XibwJiLEVVfbAeCAggP4OFEWGLu9dJYobdGPDQXvOanbgZ01aIhErWvLfRexiVL17pzouMVY_24adT2qpwNhA6KnfjNrW8tab-JFJMMUJwAvthcXmIL3d_0KCIa4dNHLR3cw9anRfh1GrEMDy6H1S49JFTYWrYOU-PsmjDgT98lFpdrAkTOU9rUkz1pr66fefDafTtvLA42ccJuLeDhXwSz2kP3II8MAiYDXlhascemnR8nWqC4OCsj2KItBGQyPInaBMqnWrrtPwf7B-4hU4_ZH9HhgY9f5pcVBB00-zjA5fds7TgDKwxOYDLVMrubfZxJ-8CtrBoOAu5LQfX0W369equkfzt6BbzCeaeQ6Fa1zxOnWzwQ76l9U_EWzuBIFuwlBwVnZ-vdxvbZdq2q24uvnKvQb34AV_p6utnQVdB-EEj18GC_AUcYreTuB2Iy14iA1Yq8yDmjfexSwwk_zKJy54nL9_tbF9KyW0QGSbzmjcTIPdw9hAgegEMurUh1XescdCt9VrLiS23kXbeSx-O6lE7_JxgxTg" -H 'Content-Type:application/json' -X DELETE http://localhost:8000/api/v1/logout

## Task

```
curl -H "Authorization: Bearer <cookie.txtのTOKEN貼り付け>" -H 'Content-Type:application/json' -X GET http://localhost:8000/api/v1/tasks
```

curl -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjo5LCJlbWFpbCI6InRlc3RAZXhtLmNvbSIsImV4cCI6MTY5MjAxNzI4NiwiaWF0IjoxNjkyMDE1NDg2LCJpc3MiOiJnaXRodWIuY29tL2tvcnktanAvdnVlX2dvL2FwaSIsImp0aSI6ImIyZmE5Y2QwLThlNGQtNGRkNi04Y2Y0LWJkMmVlZDNkNTc5NiIsInN1YiI6ImFjY2Vzc190b2tlbiJ9.m06JvUYRmwJFnqkfWFccegKTWJR_wg3j_Qd34AFbW-dMxXeKvt22m67mEOhMP-pv3gGPU90yijReenY78gTVHF2Ger-g_p7vrAbaKDPuIANl2h03i0N7LAWgthsqCdxc8KX19k_8G-XibwJiLEVVfbAeCAggP4OFEWGLu9dJYobdGPDQXvOanbgZ01aIhErWvLfRexiVL17pzouMVY_24adT2qpwNhA6KnfjNrW8tab-JFJMMUJwAvthcXmIL3d_0KCIa4dNHLR3cw9anRfh1GrEMDy6H1S49JFTYWrYOU-PsmjDgT98lFpdrAkTOU9rUkz1pr66fefDafTtvLA42ccJuLeDhXwSz2kP3II8MAiYDXlhascemnR8nWqC4OCsj2KItBGQyPInaBMqnWrrtPwf7B-4hU4_ZH9HhgY9f5pcVBB00-zjA5fds7TgDKwxOYDLVMrubfZxJ-8CtrBoOAu5LQfX0W369equkfzt6BbzCeaeQ6Fa1zxOnWzwQ76l9U_EWzuBIFuwlBwVnZ-vdxvbZdq2q24uvnKvQb34AV_p6utnQVdB-EEj18GC_AUcYreTuB2Iy14iA1Yq8yDmjfexSwwk_zKJy54nL9_tbF9KyW0QGSbzmjcTIPdw9hAgegEMurUh1XescdCt9VrLiS23kXbeSx-O6lE7_JxgxTg" -H 'Content-Type:application/json' -c cookie.txt -X GET http://localhost:8000/api/v1/tasks
