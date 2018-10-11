package payments

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/rafaelcunha/Go-mercadopago/conexaomp/mpgeral"
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

		data = []byte(strings.Replace(string(data), "\\\\", "\\", -1)) // Tive que fazer isso porque o MP está retornando \\ onde deveria ser apenas \. Sebe-se dus lá por que.

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

// AdicionaPagamento - Adiciona um novo pagamento no vetor.
func (pagamentos *ListaPagamentos) AdicionaPagamento(pagamento *Payment) error {
	pagamentos.wait.Add(1)
	defer pagamentos.wait.Done()

	pagamentos.mutex.Lock()
	defer pagamentos.mutex.Unlock()

	for _, pgto := range pagamentos.Pagamentos {
		if pgto.ID == pagamento.ID {
			return errors.New("payments.adicionaPagamento - Pagamento já existente")
		}
	}

	pagamentos.Pagamentos = append(pagamentos.Pagamentos, pagamento)

	return nil
}

// AdicionaPagamentoID - Atualiza a informação do pagamento via GET em Payments e o adiciona no vetor
func (pagamentos *ListaPagamentos) AdicionaPagamentoID(pagamentoID int64, accessToken string) (*Payment, error) {
	pagamentos.wait.Add(1)
	defer pagamentos.wait.Done()

	pagamentos.mutex.Lock()
	defer pagamentos.mutex.Unlock()

	for _, pgto := range pagamentos.Pagamentos {
		if pgto.ID == pagamentoID {
			return nil, errors.New("payments.adicionaPagamento - Pagamento já existente")
		}
	}

	pagamentos.mutex.Unlock()

	novoPagamento, err := GetPaymentByID(pagamentoID, accessToken)

	if err != nil {
		log.Println("payments.adicionaPagamentoID - Erro ao adicionar o pagamento.")
		return nil, err
	}

	pagamentos.mutex.Lock()

	pagamentos.Pagamentos = append(pagamentos.Pagamentos, &novoPagamento)

	log.Println("Encontrado pagamento: ", novoPagamento.ID)

	return &novoPagamento, nil
}

// RemovePagamentoID - Remove um pagamento do vetor de pagamentos.
func (pagamentos *ListaPagamentos) RemovePagamentoID(pagamentoID int64) *Payment {
	pagamentos.wait.Add(1)
	defer pagamentos.wait.Done()

	pagamentos.mutex.Lock()
	defer pagamentos.mutex.Unlock()

	for index, pgto := range pagamentos.Pagamentos {
		if pgto.ID == pagamentoID {
			if index == 0 {
				pagamentos.Pagamentos = append(pagamentos.Pagamentos[index+1:])
			} else if index >= (len(pagamentos.Pagamentos) - 1) {
				pagamentos.Pagamentos = append(pagamentos.Pagamentos[:index-1])
			} else {
				pagamentos.Pagamentos = append(pagamentos.Pagamentos[:index-1], pagamentos.Pagamentos[index+1:]...)
			}
			return pgto
		}
	}

	return nil
}

// Wait - agaurda sincronização completa do vetor.
func (pagamentos *ListaPagamentos) Wait() {
	pagamentos.wait.Wait()
}
