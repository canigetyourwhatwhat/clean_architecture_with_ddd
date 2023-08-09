package request

type CartItem struct {
	ProductCode string `json:"productCode"`
	Quantity    int    `json:"quantity"`
}

type ListCartItem struct {
	CartItems []CartItem `json:"records"`
}
