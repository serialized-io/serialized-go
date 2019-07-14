package main

type OrderEvent interface{}

type OrderPlacedEvent struct {
	OrderID    OrderID    `json:"orderId"`
	CustomerID CustomerID `json:"customerId"`
	Amount     Amount     `json:"amount"`
}

type PaymentReceivedEvent struct {
	OrderID    OrderID    `json:"orderId"`
	CustomerID CustomerID `json:"customerId"`
	AmountPaid Amount     `json:"amountPaid"`
}

type OrderFullyPaidEvent struct {
	OrderID    OrderID    `json:"orderId"`
	CustomerID CustomerID `json:"customerId"`
}

type OrderCancelledEvent struct {
	OrderID    OrderID    `json:"orderId"`
	CustomerID CustomerID `json:"customerId"`
	Reason     string     `json:"reason"`
}

type OrderShippedEvent struct {
	OrderID        OrderID        `json:"orderId"`
	CustomerID     CustomerID     `json:"customerId"`
	TrackingNumber TrackingNumber `json:"trackingNumber"`
}
