## やりたいこと

model packageのstructからdef packageのstructへの変換を生成

不足していること

- 適切にtypeを変換
- 何らかの変換処理を間に挟む。(e.g. bson.ObjectIdをstringにするために.Hex()を呼ぶ)
- packageのimport(fullpathを渡せるようにする)
- tagを見て分岐
- ネストした属性のサポート
- 埋め込みのサポート


## setup

```sh
$ make clean
$ make install
$ make setup
```

## run

```sh
$ make default # json + convert
make default
go-structjson --target model/model.go | jq . -S > output/model.json
go-structjson --target def/def.go | jq . -S > output/def.json
python convert.py --src output/model.json --dst output/def.json > convert/convert.go
gofmt -w convert/convert.go
$ make run
go run main.go
{"path":"/index","title":"index page"}
```