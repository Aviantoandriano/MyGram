package structs

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ErrorValidation struct {
	Parameter string `json:"parameter"`
	Message   string `json:"message"`
}
