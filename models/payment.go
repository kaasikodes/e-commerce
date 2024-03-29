package models

import "time"

type Payment struct {
	ID      string `json:"id"`
	OrderID string `json:"orderId"`
	Amount  int    `json:"amount"`
	Paid    bool   `json:"paid"`
	PaidAt  time.Time `json:"paidAt"`
	Method  string `json:"method"` //should be an enum
}