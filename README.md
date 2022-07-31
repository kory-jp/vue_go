# vue_go

## コンテナログイン

```
docker exec -it go_container /bin/sh
```

## デバック

```
docker exec -it go_container /bin/sh
dlv debug main.go(ファイル名)
```


## go install

```
docker exec -it go_container /bin/sh
go install
```