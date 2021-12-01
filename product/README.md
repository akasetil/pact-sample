# PACT 検証用の Provider 側の Web アプリ

テストのやり方などは[pact-js の example](https://github.com/pact-foundation/pact-js/tree/master/examples/e2e)を参考にしている

## アプリ起動

```
$ node src/providerService.js
```

## テスト実行

```
$ ./node_modules/.bin/mocha test/provider.spec.js
```
