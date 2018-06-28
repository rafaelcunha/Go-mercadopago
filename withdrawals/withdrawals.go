package withdrawals

import (
	"encoding/json"
	"strings"

	"github.com/rafaelcunha/Go-mercadopago/mpgeral"
)

const withdrawalsURI string = mpgeral.MLURI + "withdrawals/"

// BankAccount - Estrutura contendo informação da conta bancária
type BankAccount struct {
	Alias          string                 `json:"alias"`
	Holder         string                 `json:"holder"`
	Type           string                 `json:"type"`
	Number         string                 `json:"number"`
	Branch         string                 `json:"branch"`
	BankID         string                 `json:"bank_id"`
	Identification mpgeral.Identification `json:"identification"`
}

// Withdrawal - Estrutura para receber o saque
type Withdrawal struct {
	ID            int         `json:"id"`
	UserID        int         `json:"user_id"`
	SiteID        string      `json:"site_id"`
	Status        string      `json:"status"`
	StatusDetail  string      `json:"status_detail"`
	Amount        float64     `json:"amount"`
	Fee           int         `json:"fee"`
	CurrencyID    string      `json:"currency_id"`
	BankAccount   BankAccount `json:"bank_account"`
	PayerBank     string      `json:"payer_bank"`
	Method        string      `json:"method"`
	ClientID      string      `json:"client_id"`
	CreatedFrom   string      `json:"created_from"`
	OwOperationID int64       `json:"ow_operation_id"`
	DateSent      string      `json:"date_sent"`
	DatePaid      string      `json:"date_paid"`
	DateConfirmed string      `json:"date_confirmed"`
	DateCreated   string      `json:"date_created"`
	LastModified  string      `json:"last_modified"`
}

// SearchResponse - Resultado de search em withdrawal
type SearchResponse struct {
	Paging  mpgeral.Paging `json:"paging"`
	Results []Withdrawal   `json:"results"`
}

//=========================================================================//
//===============================Funções===================================//
//=========================================================================//

// SearchWithdrawals - busca os saques conforme a lista de parâmetros passada.
func SearchWithdrawals(parameters []mpgeral.SearchParameter, accessToken string) (SearchResponse, error) {

	resposta, err := mpgeral.APISearch(parameters, withdrawalsURI, accessToken)

	var searchResponse SearchResponse

	if err != nil {
		return searchResponse, err
	}

	err = json.NewDecoder(strings.NewReader(resposta)).Decode(&searchResponse)

	if err != nil {
		return searchResponse, err
	}

	return searchResponse, nil
}
