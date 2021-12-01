const { VerifierV3 } = require("@pact-foundation/pact/v3");
const chai = require("chai");
const chaiAsPromised = require("chai-as-promised");
chai.use(chaiAsPromised);
const { server, createProduct0001 } = require("../src/provider.js");
const path = require("path");

// 検証用のサーバープロセスを起動する
server.listen(8081, () => {
  console.log("Product Service listening on http://localhost:8081");
});

// Provider側の検証を行う
describe("Pact Verification", () => {
  it("validates the expectations of Product Service", () => {
    let opts = {
      provider: "product",
      logLevel: "INFO",
      providerBaseUrl: "http://localhost:8081",
      pactUrls: [
        path.resolve(__dirname, "../../pactfile/payment-product.json"),
      ],

      stateHandlers: {
        "Product product0001 exists": () => {
          createProduct0001();
          return Promise.resolve(`Product product0001 exists`);
        },
      },

      enablePending: true,
    };

    return new VerifierV3(opts).verifyProvider().then((output) => {
      console.log("Pact Verification Complete!");
      console.log(output);
    });
  });
});
