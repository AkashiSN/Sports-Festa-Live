# Sports-Festa-Live

## 開発環境

- golang
- dep

(Linux)

```bash
$ make dep
$ make deps
```

MacOSでbrewを使っている人は下記コマンドでdepをインストールして下さい。

(MacOS)
```bash
$ brew install dep
$ make deps
```
## ビルド

```bash
$ make build
```

## 実行

```bash
$ ./Sports-Festa-Live
```

`Unix Domain Socket`でListenする場合は、

```bash
$ ./Sports-Festa-Live -s
```
