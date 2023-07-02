# Go mock

## インストール

```
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
go install github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen
```

## mock 生成コマンド

interface が定義されているディレクトリ内に mock フォルダを作成
その後に以下のコマンドを実行

次に対象ファイルを-source に指定し、作成した mock の出力先を-destination に指定して mockgen コマンドを実行します。

```
mockgen -source=user.go -destination=./mock/mock_user.go
```

## 参考サイト

[gomock を完全に理解する](https://zenn.dev/sanpo_shiho/articles/01da627ead98f5)
