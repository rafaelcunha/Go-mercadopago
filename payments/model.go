package payments

import (
	"github.com/rafaelcunha/Go-mercadopago/mpgeral"
)

type Phone struct {
	AreaCode  string `json:"area_code"`
	Number    string `json:"number"`
	Extension string `json:"extension"`
}

type Payer struct {
	EntityType     string                 `json:"entity_type"`
	Type           string                 `json:"type"`
	ID             int64                  `json:"id,string"`
	Email          string                 `json:"email"`
	Identification mpgeral.Identification `json:"identification"`
	Phone          Phone                  `json:"phone"`
	FirstName      string                 `json:"first_name"`
	LastName       string                 `json:"last_name"`
}

type Collector struct {
	EntityType     string                 `json:"entity_type"`
	Type           string                 `json:"type"`
	ID             int64                  `json:"id"`
	Email          string                 `json:"email"`
	Identification mpgeral.Identification `json:"identification"`
	Phone          Phone                  `json:"phone"`
	FirstName      string                 `json:"first_name"`
	LastName       string                 `json:"last_name"`
}

type Order struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type TransactionDetails struct {
	TotalPaidAmount          float64     `json:"total_paid_amount"`
	AcquirerReference        interface{} `json:"acquirer_reference"`
	PaymentMethodReferenceID interface{} `json:"payment_method_reference_id"`
	NetReceivedAmount        float64     `json:"net_received_amount"`
	FinancialInstitution     interface{} `json:"financial_institution"`
	PayableDeferralPeriod    interface{} `json:"payable_deferral_period"`
	InstallmentAmount        float64     `json:"installment_amount"`
	ExternalResourceURL      string      `json:"external_resource_url"`
	OverpaidAmount           float64     `json:"overpaid_amount"`
}

type Cardholder struct {
	Identification mpgeral.Identification `json:"identification"`
	Name           string                 `json:"name"`
}

type Card struct {
	ID              int64      `json:"id,string"`
	FirstSixDigits  string     `json:"first_six_digits"`
	ExpirationMonth int        `json:"expiration_month"`
	Cardholder      Cardholder `json:"cardholder"`
	DateLastUpdated string     `json:"date_last_updated"`
	DateCreated     string     `json:"date_created"`
	ExpirationYear  int        `json:"expiration_year"`
	LastFourDigits  string     `json:"last_four_digits"`
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
	CounterCurrency      interface{} `json:"counter_currency"`
	Amount               float64     `json:"amount"`
	ID                   int64       `json:"id"`
	Source               Source      `json:"source"`
	Status               string      `json:"status"`
	CollectorMovementID  int64       `json:"collector_movement_id"`
	DateCreated          string      `json:"date_created"`
	GtwRefundID          int64       `json:"gtw_refund_id"`
	PaymentID            int64       `json:"payment_id"`
	PayerMovementID      int64       `json:"payer_movement_id"`
	UniqueSequenceNumber int64       `json:"unique_sequence_number"`
	Metadata             interface{} `json:"metadata"`
}

// Payment estrutua para receber o Json de um pagamento
type Payment struct {
	AcquirerReconciliation    []interface{}      `json:"acquirer_reconciliation"`
	StatementDescriptor       string             `json:"statement_descriptor"`
	Captured                  bool               `json:"captured"`
	DateLastUpdated           string             `json:"date_last_updated"`
	MerchantAccountID         interface{}        `json:"merchant_account_id"`
	PayerID                   int                `json:"payer_id"`
	IssuerID                  string             `json:"issuer_id"`
	Description               string             `json:"description"`
	TransactionAmount         int                `json:"transaction_amount"`
	Card                      Card               `json:"card"`
	TransactionDetails        TransactionDetails `json:"transaction_details"`
	ClientID                  string             `json:"client_id"`
	CouponAmount              int                `json:"coupon_amount"`
	Metadata                  interface{}        `json:"metadata"`
	MoneyReleaseSchema        interface{}        `json:"money_release_schema"`
	CollectorID               int                `json:"collector_id"`
	Status                    string             `json:"status"`
	FinancingType             interface{}        `json:"financing_type"`
	ProcessingMode            string             `json:"processing_mode"`
	StatusDetail              string             `json:"status_detail"`
	TransactionID             string             `json:"transaction_id"`
	Installments              int                `json:"installments"`
	InternalMetadata          interface{}        `json:"internal_metadata"`
	Refunds                   []Refund           `json:"refunds"`
	PaymentTypeID             string             `json:"payment_type_id"`
	CounterCurrency           interface{}        `json:"counter_currency"`
	ProfileID                 string             `json:"profile_id"`
	PayerTags                 []interface{}      `json:"payer_tags"`
	ReserveID                 interface{}        `json:"reserve_id"`
	CouponID                  int                `json:"coupon_id"`
	ShippingAmount            int                `json:"shipping_amount"`
	FeeDetails                []FeeDetails       `json:"fee_details"`
	Acquirer                  interface{}        `json:"acquirer"`
	DateCreated               string             `json:"date_created"`
	ID                        int64              `json:"id"`
	Collector                 Collector          `json:"collector"`
	DateOfExpiration          string             `json:"date_of_expiration"`
	MoneyReleaseDays          interface{}        `json:"money_release_days"`
	Order                     Order              `json:"order"`
	ExternalReference         string             `json:"external_reference"`
	AvailableActions          []string           `json:"available_actions"`
	ApplicationID             int                `json:"application_id"`
	Marketplace               string             `json:"marketplace"`
	MerchantNumber            int                `json:"merchant_number"`
	CallForAuthorizeID        int                `json:"call_for_authorize_id"`
	RiskExecutionID           int64              `json:"risk_execution_id"`
	APIVersion                string             `json:"api_version"`
	CurrencyID                string             `json:"currency_id"`
	SponsorID                 int64              `json:"sponsor_id"`
	DeductionSchema           interface{}        `json:"deduction_schema"`
	PaymentMethodID           string             `json:"payment_method_id"`
	AdditionalInfo            AdditionalInfo     `json:"additional_info"`
	SiteID                    string             `json:"site_id"`
	BinaryMode                bool               `json:"binary_mode"`
	OperationType             string             `json:"operation_type"`
	DifferentialPricingID     interface{}        `json:"differential_pricing_id"`
	MoneyReleaseDate          string             `json:"money_release_date"`
	Payer                     Payer              `json:"payer"`
	NotificationURL           string             `json:"notification_url"`
	TransactionAmountRefunded float64            `json:"transaction_amount_refunded"`
	CollectorTags             []string           `json:"collector_tags"`
	AuthorizationCode         string             `json:"authorization_code"`
	DateApproved              string             `json:"date_approved"`
	LiveMode                  bool               `json:"live_mode"`
}

type paymentSearch struct {
	ID int64 `json:"id"`
}

// SearchResponse - Resposta à search em payments apenas com os ids dos pagamentos
type SearchResponse struct {
	Paging  mpgeral.Paging  `json:"paging"`
	Results []paymentSearch `json:"results"`
}
