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

正常成功

```
./node_modules/.bin/mocha test/provider.spec.js
Product Service listening on http://localhost:8081


  Pact Verification
[2021-12-01 15:51:00.011 +0000] WARN (29319 on akasetomoyanoMacBook-Pro.local): pact@10.0.0-beta.54: VerifierV3 is now deprecated

  You no longer need to import the verifier from @pact-foundation/pact/v3
  You may now update your imports to:

      import { Verifier } from '@pact-foundation/pact'

  Thank you for being part of pact-js beta!
[2021-12-01 15:51:00.013 +0000] INFO (29319 on akasetomoyanoMacBook-Pro.local): pact@10.0.0-beta.54: Verifying provider
[2021-12-01 15:51:00.017 +0000] INFO (29319 on akasetomoyanoMacBook-Pro.local): pact-core@13.3.0: Verifying Pacts.
[2021-12-01 15:51:00.017 +0000] INFO (29319 on akasetomoyanoMacBook-Pro.local): pact-core@13.3.0: Verifying Pact Files
[2021-12-01 15:51:00.104 +0000] WARN (29319 on akasetomoyanoMacBook-Pro.local): pact-core@13.3.0: Ignoring option 'enablePending' because it is invalid without 'pactBrokerUrl' also being set. This may indicate an error in your configuration
[2021-12-01 15:51:00.107 +0000] WARN (29319 on akasetomoyanoMacBook-Pro.local): pact-core@13.3.0: The Verifier option 'disableSslVerification' was was explicitly set to undefined and will be ignored. This may indicate an error in your config. Remove the option entirely to prevent this warning

Verifying a pact between payment and product
[2021-12-01T15:51:00Z INFO  pact_verifier] Running provider state change handler 'Product product0001 exists' for 'A request to get product0001'
  Given Product product0001 exists
[2021-12-01T15:51:00Z INFO  pact_verifier] Running provider verification for 'A request to get product0001'
[2021-12-01T15:51:00Z INFO  pact_verifier::provider_client] Sending request to provider at http://localhost:64738/
[2021-12-01T15:51:00Z INFO  pact_verifier::provider_client] Received response: Response ( status: 200, headers: Some({"access-control-allow-origin": ["*"], "content-type": ["application/json; charset=utf-8"], "x-powered-by": ["Express"], "connection": ["close"], "content-length": ["51"], "date": ["Wed", "01 Dec 2021 15:51:00 GMT"]}), body: Present(51 bytes, application/json;charset=utf-8) )
[2021-12-01T15:51:00Z INFO  pact_matching] comparing to expected response: Response ( status: 200, headers: Some({"Content-Type": ["application/json; charset=utf-8"]}), body: Present(50 bytes) )
[2021-12-01T15:51:00Z INFO  pact_verifier] Running provider state change handler 'Product product0001 exists' for 'A request to get product0001'
  A request to get product0001
    returns a response which
      has status code 200 (OK)
      includes headers
        "Content-Type" with value "application/json; charset=utf-8" (OK)
      has a matching body (OK)


[2021-12-01 15:51:00.430 +0000] INFO (29319 on akasetomoyanoMacBook-Pro.local): pact-core@13.3.0: Verification successful
Pact Verification Complete!
finished: 0
    ✓ validates the expectations of Product Service (422ms)


  1 passing (425ms)
```

currencyCode などを追加したらエラーになった

```
./node_modules/.bin/mocha test/provider.spec.js
Product Service listening on http://localhost:8081


  Pact Verification
[2021-12-01 15:50:33.518 +0000] WARN (29295 on akasetomoyanoMacBook-Pro.local): pact@10.0.0-beta.54: VerifierV3 is now deprecated

  You no longer need to import the verifier from @pact-foundation/pact/v3
  You may now update your imports to:

      import { Verifier } from '@pact-foundation/pact'

  Thank you for being part of pact-js beta!
[2021-12-01 15:50:33.520 +0000] INFO (29295 on akasetomoyanoMacBook-Pro.local): pact@10.0.0-beta.54: Verifying provider
[2021-12-01 15:50:33.523 +0000] INFO (29295 on akasetomoyanoMacBook-Pro.local): pact-core@13.3.0: Verifying Pacts.
[2021-12-01 15:50:33.523 +0000] INFO (29295 on akasetomoyanoMacBook-Pro.local): pact-core@13.3.0: Verifying Pact Files
[2021-12-01 15:50:33.617 +0000] WARN (29295 on akasetomoyanoMacBook-Pro.local): pact-core@13.3.0: Ignoring option 'enablePending' because it is invalid without 'pactBrokerUrl' also being set. This may indicate an error in your configuration
[2021-12-01 15:50:33.620 +0000] WARN (29295 on akasetomoyanoMacBook-Pro.local): pact-core@13.3.0: The Verifier option 'disableSslVerification' was was explicitly set to undefined and will be ignored. This may indicate an error in your config. Remove the option entirely to prevent this warning

Verifying a pact between payment and product
[2021-12-01T15:50:33Z INFO  pact_verifier] Running provider state change handler 'Product product0001 exists' for 'A request to get product0001'
  Given Product product0001 exists
[2021-12-01T15:50:33Z INFO  pact_verifier] Running provider verification for 'A request to get product0001'
[2021-12-01T15:50:33Z INFO  pact_verifier::provider_client] Sending request to provider at http://localhost:64721/
[2021-12-01T15:50:33Z INFO  pact_verifier::provider_client] Received response: Response ( status: 200, headers: Some({"date": ["Wed", "01 Dec 2021 15:50:33 GMT"], "content-type": ["application/json; charset=utf-8"], "x-powered-by": ["Express"], "connection": ["close"], "content-length": ["75"], "access-control-allow-origin": ["*"]}), body: Present(75 bytes, application/json;charset=utf-8) )
[2021-12-01T15:50:33Z INFO  pact_matching] comparing to expected response: Response ( status: 200, headers: Some({"Content-Type": ["application/json; charset=utf-8"]}), body: Present(50 bytes) )
[2021-12-01T15:50:33Z INFO  pact_verifier] Running provider state change handler 'Product product0001 exists' for 'A request to get product0001'
  A request to get product0001
    returns a response which
      has status code 200 (OK)
      includes headers
        "Content-Type" with value "application/json; charset=utf-8" (OK)
      has a matching body (FAILED)


Failures:

1) Verifying a pact between payment and product Given Product product0001 exists - A request to get product0001
    1.1) has a matching body
           $.price -> Expected '1000.45' to be an integer value

There were 1 pact failures

[2021-12-01 15:50:33.999 +0000] ERROR (29295 on akasetomoyanoMacBook-Pro.local): pact-core@13.3.0: Verification unsuccessful
    1) validates the expectations of Product Service


  0 passing (488ms)
  1 failing

  1) Pact Verification
       validates the expectations of Product Service:
     Error: Verfication failed
      at /Users/akasetomoya/Desktop/SD執筆/product/node_modules/@pact-foundation/pact-core/src/verifier/nativeVerifier.js:41:32
      at Object.<anonymous> (node_modules/ffi-napi/lib/_foreign_function.js:115:9)
```
