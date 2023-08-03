package domain

type PaymentEvent struct {
	Event   string `json:"event"`
	Entity  string `json:"entity"`
	ID      string `json:"id"`
	Created int64  `json:"created_at"`
	Payload struct {
		Payment struct {
			Entity           string `json:"entity"`
			ID               string `json:"id"`
			Amount           int    `json:"amount"`
			Currency         string `json:"currency"`
			Status           string `json:"status"`
			OrderID          string `json:"order_id"`
			Method           string `json:"method"`
			Description      string `json:"description"`
			Email            string `json:"email"`
			Contact          string `json:"contact"`
			Fee              int    `json:"fee"`
			Tax              int    `json:"tax"`
			ErrorCode        string `json:"error_code"`
			ErrorDescription string `json:"error_description"`
			CardID           string `json:"card_id"`
			Bank             string `json:"bank"`
			Wallet           string `json:"wallet"`
			VPA              string `json:"vpa"`
			International    bool   `json:"international"`
		} `json:"payment"`
	} `json:"payload"`
}
