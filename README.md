# Epub Ruby Remover

Epub形式のファイルからRubyを削除するスクリプトです。

## Usage

ルビを削除したいepubファイルを引数に指定して実行します。

```
go run ./main.go <epub file>
```

出力時は `<epub file>_no_ruby.epub` というファイル名で出力されます。

ディレクトリを一括で処理したい場合は以下のようにします。

```
go run ./main.go dir <epub directory>
```