package risk


type TransactionsRequest struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ID            int `json:"id"`
	UserID        int `json:"user_id"`
	AmountUsCents int `json:"amount_us_cents"`
	CardID        int `json:"card_id"`
}

type RiskRatingsResponse struct {
	RiskRatings []string `json:"risk_ratings"`
}
