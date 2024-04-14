package types

import "time"

type VerifyPaystackTransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID              int         `json:"id"`
		Domain          string      `json:"domain"`
		Status          string      `json:"status"`
		Reference       string      `json:"reference"`
		ReceiptNumber   interface{} `json:"receipt_number"`
		Amount          int         `json:"amount"`
		GatewayResponse string      `json:"gateway_response"`
		PaidAt          time.Time    `json:"paid_at"`
		CreatedAt       string      `json:"created_at"`
		Channel         string      `json:"channel"`
		Currency        string      `json:"currency"`
		IPAddress       string      `json:"ip_address"`
		Metadata        string      `json:"metadata"`
		Log             interface{} `json:"log"`
		Fees            interface{} `json:"fees"`
		FeesSplit       interface{} `json:"fees_split"`
		Authorization   struct{}    `json:"authorization"`
		Customer        struct {
			ID                       int         `json:"id"`
			FirstName                interface{} `json:"first_name"`
			LastName                 interface{} `json:"last_name"`
			Email                    string      `json:"email"`
			CustomerCode             string      `json:"customer_code"`
			Phone                    interface{} `json:"phone"`
			Metadata                 interface{} `json:"metadata"`
			RiskAction               string      `json:"risk_action"`
			InternationalFormatPhone interface{} `json:"international_format_phone"`
		} `json:"customer"`
		Plan               interface{} `json:"plan"`
		Split              struct{}    `json:"split"`
		OrderID            interface{} `json:"order_id"`
		PaidAt2            interface{} `json:"paidAt"`
		CreatedAt2         string      `json:"createdAt"`
		RequestedAmount    int         `json:"requested_amount"`
		POSTransactionData interface{} `json:"pos_transaction_data"`
		Source             interface{} `json:"source"`
		FeesBreakdown      interface{} `json:"fees_breakdown"`
		Connect            interface{} `json:"connect"`
		TransactionDate    string      `json:"transaction_date"`
		PlanObject         struct{}    `json:"plan_object"`
		Subaccount         struct{}    `json:"subaccount"`
	} `json:"data"`
}
