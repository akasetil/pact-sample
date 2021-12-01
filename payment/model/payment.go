package model

type Payment struct {
	UserID string  `json:"userID"`
	Order  []Order `json:"order"`
}

type Order struct {
	ProductID string `json:"productID"`
	Quantity  int    `json:"quantity"`
}

type Product struct {
	ProductID string `json:"productID"`
	Price     int    `json:"price" pact:"float"`
	Stock     int    `json:"stock"`
}
