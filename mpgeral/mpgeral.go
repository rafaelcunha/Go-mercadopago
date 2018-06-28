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

// Identification - Estrutura com a informação de documento e tipo de documento
type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

// SearchParameter - estrutura com nome e valor dos parâmetros a serem passados para realização de buscas
type SearchParameter struct {
	Name  string
	Value string
}

// ErrorCause - estrutura para receber as informações de causas erros em consultas.
type ErrorCause struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

// ErrorResponse - estrutura para receber as informações de erros em consultas.
type ErrorResponse struct {
	Message string       `json:"message"`
	Error   string       `json:"error"`
	Status  int          `json:"status"`
	Cause   []ErrorCause `json:"cause"`
}

// Paging - estrutura para receber informações de paginação nas consultas.
type Paging struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// SearchResponse - Resultado de search em withdrawal
type SearchResponse struct {
	Paging  Paging `json:"paging"`
	Results []byte `json:"results"`
}

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
