package payments

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/rafaelcunha/Go-mercadopago/mpgeral"
)

const paymentsURI string = mpgeral.MPURI + "payments/"

// SearchPayments busca os pagamentos conforme a lista de parâmetros passada.
func SearchPayments(parameters []mpgeral.SearchParameter, accessToken string) (SearchResponse, error) {

	resposta, err := mpgeral.APISearch(parameters, paymentsURI, accessToken)

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

// GetPaymentByID busca um pagamento pelo seu id.
func GetPaymentByID(id int64, accessToken string) (Payment, error) {

	var payment Payment

	// TODO: Verificar uma forma de converter int64 para inteiro
	var uri = paymentsURI + strconv.Itoa(int(id)) + "?access_token=" + accessToken

	response, err := http.Get(uri)

	if err != nil {
		return payment, err
	}

	data, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode >= 200 && response.StatusCode < 300 {

		err = json.Unmarshal(data, &payment)

		return payment, nil
	}
	var erroResposta mpgeral.ErrorResponse

	err = json.Unmarshal(data, &erroResposta)

	if err != nil {
		return payment, err
	}

	return payment, errors.New(erroResposta.Message)
}
