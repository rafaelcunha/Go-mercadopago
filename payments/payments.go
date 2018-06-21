package payments

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const mpURI string = "https://api.mercadopago.com/v1/"
const paymentsURI string = mpURI + "payments/"

type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Phone struct {
	AreaCode  string `json:"area_code"`
	Number    string `json:"number"`
	Extension string `json:"extension"`
}

type Payer struct {
	EntityType     string         `json:"entity_type"`
	Type           string         `json:"type"`
	ID             string         `json:"id"`
	Email          string         `json:"email"`
	Identification Identification `json:"identification"`
	Phone          Phone          `json:"phone"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
}

type Collector struct {
	EntityType     string         `json:"entity_type"`
	Type           string         `json:"type"`
	ID             int            `json:"id"`
	Email          string         `json:"email"`
	Identification Identification `json:"identification"`
	Phone          Phone          `json:"phone"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
}

type Order struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type TransactionDetails struct {
	FinancialInstitution     string  `json:"financial_institution"`
	NetReceivedAmount        float32 `json:"net_received_amount"`
	TotalPaidAmount          float32 `json:"total_paid_amount"`
	InstallmentAmount        float32 `json:"installment_amount"`
	OverpaidAmount           float32 `json:"overpaid_amount"`
	ExternalResourceURL      string  `json:"external_resource_url"`
	PaymentMethodReferenceID string  `json:"payment_method_reference_id"`
	PayableDeferralPeriod    string  `json:"payable_deferral_period"`
	AcquirerReference        int     `json:"acquirer_reference"`
}

type Cardholder struct {
	Name           string         `json:"name"`
	Identification Identification `json:"identification"`
}

type Card struct {
	ID              string     `json:"id"`
	LastFourDigits  string     `json:"last_four_digits"`
	FirstSixDigits  string     `json:"first_six_digits"`
	ExpirationYear  int        `json:"expiration_year"`
	ExpirationMonth int        `json:"expiration_month"`
	DateCreated     string     `json:"date_created"`      // Formatar corretamente a data
	DateLastUpdated string     `json:"date_last_updated"` // Formatar corretamente a data
	Cardholder      Cardholder `json:"cardholder"`
}

type Item struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PictureURL  string `json:"picture_url"`
	CategoryID  string `json:"category_id"`
	Quantity    string `json:"quantity"`
	UnitPrice   string `json:"unit_price"`
}

type AdditionalInfoPhone struct {
	AreaCode string `json:"area_code"`
	Number   string `json:"number"`
}

type AdditionalInfoAddress struct {
	ZipCode      string `json:"zip_code"`
	StreetName   string `json:"street_name"`
	StreetNumber int    `json:"street_number"`
}

type AdditionalInfoPayer struct {
	FirstName        string                `json:"first_name"`
	LastName         string                `json:"last_name"`
	Phone            AdditionalInfoPhone   `json:"phone"`
	Address          AdditionalInfoAddress `json:"address"`
	RegistrationDate string                `json:"registration_date"` // Formatar corretamente a data

}

type ReceiverAddress struct {
	ZipCode      string `json:"zip_code"`
	StreetName   string `json:"street_name"`
	StreetNumber int    `json:"street_number"`
	Floor        string `json:"floor"`
	Apartment    string `json:"apartment"`
}

type Shipments struct {
	ReceiverAddress ReceiverAddress `json:"receiver_address"`
}

type AdditionalInfo struct {
	IPAddress string              `json:"ip_address"`
	Items     []Item              `json:"items"`
	Payer     AdditionalInfoPayer `json:"payer"`
	Shipments Shipments           `json:"shipments"`
}

type Barcode struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

type FeeDetails struct {
	Type     string  `json:"type"`
	Amount   float32 `json:"amount"`
	FeePayer string  `json:"fee_payer"`
}

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Refund struct {
	ID        int     `json:"id"`
	PaymentID int     `json:"PaymentID"`
	Amount    float32 `json:"amount"`
	//Metadata   string `json:"metadata"`
	DateCreated          string `json:"date_created"`
	UniqueSequenceNumber int    `json:"unique_sequence_number"`
	Status               string `json:"status"`
	PayerMovementID      int    `json:"payer_movement_id"`
	CollectorMovementID  int    `json:"collector_movement_id"`
	GtwRefundID          int    `json:"gtw_refund_id"`
	CounterCurrency      int    `json:"counter_currency"`
	Source               Source `json:"source"`
}

// Payment estrutua para receber o Json de um pagamento
type Payment struct {
	ID                int    `json:"id"`
	DateCreated       string `json:"date_created"`       // Formatar corretamente a data
	DateApproved      string `json:"date_approved"`      // Formatar corretamente a data
	DateLastUpdated   string `json:"date_last_updated"`  // Formatar corretamente a data
	DateOfExpiration  string `json:"date_of_expiration"` // Formatar corretamente a data
	MoneyReleaseDate  string `json:"money_release_date"` // Formatar corretamente a data
	OperationType     string `json:"operation_type"`
	IssuerID          string `json:"issuer_id"`
	CollectorID       int    `json:"collector_id"`
	Payer             Payer  `json:"payer"`
	BinaryMode        bool   `json:"binary_mode"`
	LiveMode          bool   `json:"live_mode"`
	Order             Order  `json:"order"`
	ExternalReference string `json:"external_reference"`
	Description       string `json:"description"`
	//Metadata                  string             `json:"metadata"`  "metadata": {},
	/*InternalMetadata                  string             `json:"internal_metadata"`
		"internal_metadata": {
	        "subtype": "store",
	        "type": "qr"
	    },*/
	CurrencyID                string             `json:"currency_id"`
	TransactionAmount         float32            `json:"transaction_amount"`
	TransactionAmountRefunded float32            `json:"transaction_amount_refunded"`
	CouponAmount              float32            `json:"coupon_amount"`
	CampaignID                int                `json:"campaign_id"`
	CouponCode                string             `json:"coupon_code"`
	TransactionDetails        TransactionDetails `json:"transaction_details"`
	FeeDetails                []FeeDetails       `json:"fee_details"`
	DifferentialPricingID     int                `json:"differential_pricing_id"`
	ApplicationFee            float32            `json:"application_fee"`
	Status                    string             `json:"status"`
	StatusDetail              string             `json:"status_detail"`
	Capture                   bool               `json:"capture"`
	Captured                  bool               `json:"captured"`
	CallForAuthorizeID        string             `json:"call_for_authorize_id"`
	PaymentMethodID           string             `json:"payment_method_id"`
	PaymentTypeID             string             `json:"payment_type_id"`
	Token                     string             `json:"token"`
	Card                      Card               `json:"card"`
	StatementDescriptor       string             `json:"statement_descriptor"`
	Installments              int                `json:"installments"`
	NotificationURL           string             `json:"notification_url"`
	CallbackURL               string             `json:"callback_url"`
	AdditionalInfo            AdditionalInfo     `json:"additional_info"`
	Barcode                   Barcode            `json:"barcode"`
	PayerID                   int                `json:"payer_id"`
	SponsorID                 int                `json:"sponsor_id"`
	AuthorizationCode         string             `json:"authorization_code"`
	MoneyReleaseSchema        int                `json:"money_release_schema"`
	CounterCurrency           int                `json:"counter_currency"`
	SiteID                    string             `json:"site_id"`
	Marketplace               string             `json:"marketplace"`
	ApplicationID             int                `json:"application_id"`
	TransactionID             string             `json:"transaction_id"`
	CouponID                  int                `json:"coupon_id"`
	RiskExecutionID           int                `json:"risk_execution_id"`
	AvailableActions          []string           `json:"available_actions"`
	ProfileID                 string             `json:"profile_id"`
	APIVersion                string             `json:"api_version"`
	ClientID                  string             `json:"client_id"`
	ShippingAmount            int                `json:"shipping_amount"`
	ReserveID                 int                `json:"reserve_id"`
	ProcessingMode            string             `json:"processing_mode"`
	MerchantAccountID         int                `json:"merchant_account_id"`
	Acquirer                  int                `json:"acquirer"`
	MerchantNumber            int                `json:"merchant_number"`
	//AcquirerReconciliation    int                `json:"acquirer_reconciliation"`
	DeductionSchema  int       `json:"deduction_schema"`
	MoneyReleaseDays int       `json:"money_release_days"`
	Collector        Collector `json:"collector"`
	Refunds          []Refund  `json:"refunds"`
}

// GetPaymentByID busca um pagamento pelo seu id.
func GetPaymentByID(id int, accessToken string) (payment Payment, err error) {

	var uri = paymentsURI + strconv.Itoa(id) + "?access_token=" + accessToken
	fmt.Println("Buscando pagamento na uri: " + uri)

	response, err := http.Get(uri)

	if err != nil {
		return
	}

	data, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(data, &payment)

	return
}
