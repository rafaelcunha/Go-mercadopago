package mpgeral

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
