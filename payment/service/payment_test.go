package service

import (
	"fmt"
	"github.com/pact-foundation/pact-go/v2/consumer"
	"github.com/pact-foundation/pact-go/v2/matchers"
	. "github.com/pact-foundation/pact-go/v2/sugar"
	"log"
	"testing"
)

func TestConsumer(t *testing.T) {

	productID := "product0001"

	// Create Pact connecting to local Daemon
	pact, err := consumer.NewV3Pact(consumer.MockHTTPProviderConfig{
		Consumer: "payment",
		Provider: "product",
		Host:     "127.0.0.1",
		PactDir: "../../pactfile",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Set up our expected interactions.
	pact.
		AddInteraction().
		Given(ProviderStateV3{Name: fmt.Sprintf("Product %s exists", productID)}).
		UponReceiving(fmt.Sprintf("A request to get %s", productID)).
		WithRequest("GET", matchers.String(fmt.Sprintf("/product/%s", productID))).
		WillRespondWith(200).
		WithHeader("Content-Type", Regex("application/json; charset=utf-8", "application\\/json; charset=utf-8")).
		WithJSONBody(Map{
			"productID": S("product0001"),
			"price": Integer(100),
			"stock": Integer(10),
		})

	// Run the test, verify it did what we expected and capture the contract
	if err := pact.ExecuteTest(t, func(config MockServerConfig) (err error) {
		log.Println(config.Port)
		productClient := ProductClient{baseURL: fmt.Sprintf("http://localhost:%d", config.Port)}
		_, err = productClient.getProduct(productID)
		if err != nil {
			return
		}
		return
	}); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}
