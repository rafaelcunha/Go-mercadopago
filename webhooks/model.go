package webhooks

var m map[string]int

// Data - Estrutura do campo data do webhook
type Data struct {
	ID string `json:"id"`
}

// Webhook - Estrutura do webhook
type Webhook struct {
	ID          int64  `json:"id"`
	LiveMode    bool   `json:"live_mode"`
	Type        string `json:"type"`
	DateCreated string `json:"date_created"`
	UserID      int64  `json:"user_id"`
	APIVersion  string `json:"api_version"`
	Action      string `json:"action"`
	Data        Data   `json:"data"`
}
