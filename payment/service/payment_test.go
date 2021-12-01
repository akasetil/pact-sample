package service

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"log"
	"payment/model"
	"testing"
)

func TestConsumer(t *testing.T) {

	productID := "product0001"

	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Consumer: "payment",
		Provider: "product",
		Host:     "localhost",
		PactDir: "../../pactfile",
	}
	defer pact.Teardown()

	// Set up our expected interactions.
	pact.
		AddInteraction().
		Given(fmt.Sprintf("Product %s exists", productID)).
		UponReceiving(fmt.Sprintf("A request to get %s", productID)).
		WithRequest(dsl.Request{
			Method:  "GET",
			Path:    dsl.String(fmt.Sprintf("/product/%s", productID)),
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.Match(&model.Product{}),
		})

	// Run the test, verify it did what we expected and capture the contract
	if err := pact.Verify(func() (err error) {
		productClient := ProductClient{baseURL: fmt.Sprintf("http://localhost:%d", pact.Server.Port)}
		_, err = productClient.getProduct(productID)
		if err != nil {
			return
		}
		return
	}); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}