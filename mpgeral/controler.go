package mpgeral

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// MPURI - URL base para acesso às APIs Mercado Pago
const MPURI string = "https://api.mercadopago.com/v1/"

// MLURI - URL base para acesso às APIs Mercado Pago
const MLURI string = "https://api.mercadolibre.com/"

// APISearch - busca em alguma API conforme a lista de parâmetros passada.
func APISearch(parameters []SearchParameter, uriBase string, accessToken string) ([]byte, error) {

	var uri = uriBase + "search?"

	if parameters != nil {
		for _, element := range parameters {
			uri += element.Name + "=" + element.Value + "&"
		}
	}

	uri += "access_token=" + accessToken

	var data []byte

	response, err := http.Get(uri)

	if err != nil {
		return data, err
	}

	data, _ = ioutil.ReadAll(response.Body)

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return data, nil
	}
	var erroResposta ErrorResponse

	err = json.NewDecoder(strings.NewReader(string(data))).Decode(&erroResposta)

	if err != nil {
		return data, err
	}

	return data, errors.New(erroResposta.Message)
}
