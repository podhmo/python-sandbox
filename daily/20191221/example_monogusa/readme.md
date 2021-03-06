# monogusa

- https://github.com/podhmo/monogusa

## 最近のmonogusa

```
$ git ls-files ../../ | xargs dirname | grep -oP ".*example_monogusa$" | sort -u | tac
```

- [../../20191218/example_monogusa](../../20191218/example_monogusa)
- [../../20191217/example_monogusa](../../20191217/example_monogusa)
- [../../20191216/example_monogusa](../../20191216/example_monogusa)
- [../../20191215/example_monogusa](../../20191215/example_monogusa)
- [../../20191214/example_monogusa](../../20191214/example_monogusa)
- [../../20191211/example_monogusa](../../20191211/example_monogusa)

## また少しずつ調整をしていく

たしか前回でかんたんなweb用のコードは出力できるようになった。
ただ出力関係がまだ標準出力ベースでこれは後々きつくなっていくことが予想される感じだった。

とりあえずprestring.output経由で出力して複数ファイルの出力ができるようにしておきたい。

## 未来の話

ちょっとだけ未来の話。

(追記: しかしwebばかりに閉じた話になってしまっているけれど。web interfaceを作りたいだけではないのだよなー)

### on_startup, on_shutdown event

あと不足している部分はon_startup, on_shutdown的なevent (setup, teardownと見做しても良い)。なんで欲しくなるかというと例えばcomponentのconnect,disconnectのような処理を間にはさみたいから。コレがあるとCRUD的なインターフェイスのモジュールをそのままコマンドとして提供できるようになる。

crud.py

```python
async def read_notes(db: Database, *, completed_only:bool=True) -> List[Note]
    query = notes.select()
    return await db.fetch_all(query)
```

現状はそれがないのでCLI用のコマンドを作るために以下の様なコードが必要になる。

```python
async def list(db: Database) -> None:
    await db.connect()  # TODO: startup (lifecycle)
    print(await crud.read_notes(db))
```

これは、web上のAPIとしての公開を考えた時に嬉しくない。

```python
@app.get("/notes/", response_model=t.List[Note])
async def read_notes():
    return await crud.read_notes(db)

# 現状ではすべてcommand likeな処理のシェル実行をシミュレートしているので以下
from fastapi import Depends
from monogusa.web import runtime

@app.get("/notes/", response_model=runtime.CommandOutput)
async def read_notes(db: Database = Depends()):
    with runtime.handle() as s:
        await crud.read_notes(db)
        return s.dict()
```

### callback action

加えてcallback action (presenter) 的なものも本当は用意できておくと良い。コレがあると今までのようなcommand likeな関数 (write系) だけでなく、query likeな関数 (read系) も雑に公開できる余地が生まれる。とくにCLIとwebの相性的な話で。

雑に考えるならdefaultはprintで良いかもしれない。コレがあると完全にcrud.py的なものを作った時点で公開できるようになる。

これは実質crud.read_notesで良い :tada:

```python
async def list(db: Database, *, completed_only:bool=True) -> List[Note]
    return await crud.read_notes(db, completed_only=completed_only)
```

あるいは `--format=json` みたいなオプションを勝手にはやしても良いかもしれない。

ただ、暗黙のデフォルト動作を増やしすぎて挙動の理解が困難になるのは避けたい所。全ての動作が記述されたコードが生成される `--expose` 的なオプションがあっても良いのかもしれない。


### bulk action

あと地味に嫌なのがimportが重い系の処理を複数の対象に対して実行する必要が出てきた時。処理時間のほとんどはimportなどのload time。かなしい。

これを良い感じに実行する機能があると嬉しい。どういうインターフェイスにするかは決めあぐねていて。現状の思いつきのアイデアとしては標準入力を使うというもの (just my x cent 的なやつ)。

commands.txt

```
add "brush my teeth" --completed
add "read book"
add "goto bed"
```

こんな感じで実行する。dbへの接続/切断とかなんども繰り返したくないよね。。
(あとpandasのimportとか。。)

```console
$ cat commands.txt | python -m monogusa.cli crud.py
```

parseだけするようなdry-runオプションもあっても良いかもしれない。

まぁ本音を言えばIPython Kernel的なものを立ち上げてRPCしたい。

## prestring.output

組み込み自体はできたけれど。コード生成周りのコードが綺麗にならないな。。

### refactoring

少しstarlette的な感じを意識してコードを書いてみるか。
ようやく組み込めた。まぁ綺麗じゃないけど許容範囲内くらいの感じ。

## web with DI

fastAPIのDI的なものへの変換を書く必要があった。忘れてた。
あと、せっかくだしcliの方もコード生成しても良いような気がしてきた。
あと、どう認識されたかのscanができると嬉しいかもしれない。

### depends

コレ手軽にやる方法ないかな。後一部挙動を勘違いしているかもしれない？
ContextVarを使う例辺りを参考にしてみると良さそう。

- https://fastapi.tiangolo.com/tutorial/sql-databases-peewee/

hmm..そういえばyield fixtureに対応していないな。。
そしてpytestのscope的なものが欲しいかも？

- session -> on_startup, on_shutdown
- module -> ?
- function -> depends (each request) with contextvars

とりあえずやること

- 同期で全部関数を定義し直す
- 同期で引数0個のものは直接使う
- importエラーを起こさない様に対応する
- 非同期対応

### 追記

とりあえず強引だけど対応した。

### 追記

細かい調整が必要になった。長さが0のときとか。


## componentのキーワード引数

componentのキーワード引数もパラメーターとして露出できると便利なんだろうか？
