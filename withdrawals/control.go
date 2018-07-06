package withdrawals

import (
	"encoding/json"

	"github.com/rafaelcunha/Go-mercadopago/mpgeral"
)

const withdrawalsURI string = mpgeral.MLURI + "withdrawals/"

// SearchWithdrawals - busca os saques conforme a lista de parâmetros passada.
func SearchWithdrawals(parameters []mpgeral.SearchParameter, accessToken string) (SearchResponse, error) {

	resposta, err := mpgeral.APISearch(parameters, withdrawalsURI, accessToken)

	var searchResponse SearchResponse

	if err != nil {
		return searchResponse, err
	}

	err = json.Unmarshal(resposta, &searchResponse)

	if err != nil {
		return searchResponse, err
	}

	return searchResponse, nil
}
