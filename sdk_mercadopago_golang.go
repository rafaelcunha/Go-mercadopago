package main

import (
	"fmt"

	"github.com/rafaelcunha/sdk_mercadopago_golang/payments"
)

func main() {
	/*pagamento, err := payments.GetPaymentByID(3678606946, "APP_USR-8387298048269828-041717-0512a631b5ef748a19c384d028af76e1-52192556")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", pagamento)*/

	/*var parametros = []payments.SearchParameter{
		payments.SearchParameter{
			Name:  "parametro1",
			Value: "valor1",
		},
		payments.SearchParameter{
			Name:  "parametro2",
			Value: "valor2",
		},
	}*/

	searchResponse, err := payments.SearchPayments(nil, "APP_USR-8387298048269828-041717-0512a631b5ef748a19c384d028af76e1-52192556")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v", searchResponse)

}
