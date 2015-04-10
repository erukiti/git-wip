git-wip
=======

Github 及び Github:enterprise 向けの git-wip flow 用ツール。

前提
----

コマンドラインの、git, hub コマンドをそれぞれインストールしてる必要あり

インストール
------------

```sh
$ go build git-wip.go
$ sudo cp git-wip ~/bin
```

たとえば、~/bin にコピーする。

設定
----

環境変数に設定を行う

環境変数名|内容
----------|----------------------------
GITHUB_TOKEN | hub コマンドが利用
GITWIP_TEMPLATE_CMD | テンプレートを取ってくるコマンドを登録する


使い方
------

```sh
$ git wip [-T] <branch>
```

* -T が指定されていれば GITWIP_TEMPLATE_CMD を実行して結果をテンプレートとして使う
* 現在のブランチから分岐した branch ブランチが生成される
* 空のコミットが生成され origin に push される
* プルリクエストが発行される
