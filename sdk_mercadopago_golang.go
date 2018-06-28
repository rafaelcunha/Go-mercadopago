package main

import (
	"fmt"

	"github.com/rafaelcunha/Go-mercadopago/mpgeral"
	"github.com/rafaelcunha/Go-mercadopago/payments"
)

func main() {

	// Teste do get em payments
	/*pagamento, err := payments.GetPaymentByID(3678606946, "ADM-601-062722-89883dd252118f749958836f3bac8052-rcunha-62867623")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", pagamento)*/

	// Teste do search de payments
	var parametros = []mpgeral.SearchParameter{
		mpgeral.SearchParameter{
			Name:  "collector.id",
			Value: "52192556",
		},
	}

	searchResponse, err := payments.SearchPayments(parametros, "ADM-601-062722-89883dd252118f749958836f3bac8052-rcunha-62867623")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", searchResponse)

	// Teste o search de withdrawals
	/*
		var parametros = []mpgeral.SearchParameter{
			mpgeral.SearchParameter{
				Name:  "user_id",
				Value: "217979967",
			},
			mpgeral.SearchParameter{
				Name:  "limit",
				Value: "2",
			},
		}

		searchResponse, err := withdrawals.SearchWithdrawals(parametros, "ADM-601-062722-89883dd252118f749958836f3bac8052-rcunha-62867623")

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%+v\n", searchResponse)
	*/
}
