package main

import (
	"fmt"

	"github.com/rafaelcunha/Go-mercadopago/payments"
)

func main() {
	/*pagamento, err := payments.GetPaymentByID(3678606946, "APP_USR-...")

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

	searchResponse, err := payments.SearchPayments(nil, "APP_USR-...")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v", searchResponse)

}
