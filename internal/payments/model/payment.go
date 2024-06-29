package model

type Payment struct {
	AccountId   int     `json:"account_id"`
	PhoneNumber string  `json:"phone_number"`
	Amount      float64 `json:"amount"`
	Narration   string  `json:"narration"`
	CallbackURL string  `json:"callback_url"`
	Channel     string  `json:"channel"`
	//CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	//UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PaymentResponse struct {
	ResponseCode         string `json:"responseCode"`
	Message              string `json:"message"`
	MerchantRequestID    string `json:"merchantRequestID"`
	TransactionReference string `json:"transactionReference"`
	PaymentGateway       string `json:"paymentGateway"`
	CheckoutRequestID    string `json:"checkoutRequestID"`
	CustomerMessage      string `json:"customerMessage"`
}
