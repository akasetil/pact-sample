# PACT base App

PACT 検証用のベースアプリ

## 処理フロー

![](./plantuml/シーケンス図.png)

## 使い方

### PACT cli のインストール

mac の場合

```
brew tap pact-foundation/pact-ruby-standalone
brew install pact-ruby-standalone
```

### PACT の契約書ファイル作成

```
cd payment
go get
go test -v ./...
```

これで`pactfile/`ディレクトリに契約書ファイルが作成される

### provider 側の検証

```
cd product
npm install
// もしPermission deniedが発生したら, sudo npm install を実行する
./node_modules/.bin/mocha test/provider.spec.js
```

`pactfile/`ディレクトリにある契約書ファイルを読み込み、テストを mocha を利用したテストを実行している
返却する Product の構造体を変更すると pact の検証で失敗するようになる
