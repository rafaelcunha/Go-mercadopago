package webhooks

import (
	"encoding/json"
	"log"
	"net/http"
)

// WebhookJSONDecoder - Decodigica o webhook de pagamento.
func WebhookJSONDecoder(r *http.Request) (*Webhook, error) {
	decoder := json.NewDecoder(r.Body)
	var webhook Webhook
	err := decoder.Decode(&webhook)

	if err != nil {
		log.Println("webhookJsonDecoder - falha ao gerar estrutura do Webhook recebido.")
		return nil, err
	}

	return &webhook, nil
}

// TrataWebhook - Trata uma chamada de webhook recebida
/*func TrataWebhook(w http.ResponseWriter, r *http.Request) (err error) {

	parametros := r.URL.Query() // Returns a url.Values, which is a map[string][]string

	log.Println("Parametros recebidos: ")

	for key, value := range parametros {
		log.Println("Key: ", key, "Value: ", value)
	}

	tipo := parametros.Get("type")

	switch tipo {
	case "payment":

		break
	default:
		err = errors.New("Tipo de webhook sem tratamento definido. Tipo: " + tipo)
	}

	return err
}*/
