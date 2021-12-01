package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"payment/model"
)

type PaymentService struct{}

func (PaymentService) Payment(payment *model.Payment) error {

	productClient := ProductClient{baseURL: "http://127.0.0.1:8000"}
	userClient := UserClient{baseURL: "http://127.0.0.1:8001"}

	var totalAmount int
	for _, order := range payment.Order {
		product, err := productClient.getProduct(order.ProductID)
		if err != nil {
			return err
		}
		err = productClient.reduceProductStock(order)
		if err != nil {
			return err
		}
		totalAmount += product.Price * order.Quantity
	}

	err := userClient.reduceUserBalance(payment.UserID, totalAmount)
	if err != nil {
		return err
	}

	return nil
}

type ProductClient struct {
	baseURL string
}

func (c *ProductClient) getProduct(productID string) (*model.Product, error) {
	url := fmt.Sprintf("%s/product/%s", c.baseURL, productID)
	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res model.Product
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// dummy
func (c *ProductClient) reduceProductStock(order model.Order) error {
	return nil
}

type UserClient struct {
	baseURL string
}

// dummy
func (c *UserClient) reduceUserBalance(userID string, amount int) error {
	return nil
}
