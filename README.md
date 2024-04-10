A microservice API for a fictional e-commerce system to deal with orders.\
Developed in Go using go-chi for building the REST API, Redis as the datastore.

GET     /orders        "Lists the orders using pagination"\
POST    /orders        "Creates an order with provided body"\
GET     /orders/{id}   "Get order by ID"\
PUT     /orders/{id}   "Update order that has ID with provided body"\
DELETE  /orders/{id}   "Delete order by ID"\
```
type Order struct {
	OrderID     uint64     `json:"order_id"`
	CustomerID  uuid.UUID  `json:"customer_id"`
	LineItems   []LineItem `json:"line_items"`
	CreatedAt   *time.Time `json:"created_at"`
	ShippedAt   *time.Time `json:"shipped_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type LineItem struct {
	ItemID   uuid.UUID `json:"item_id"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
```
