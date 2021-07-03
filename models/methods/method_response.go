package methods

type MethodsResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Data    []Methods `json:"data"`
}

type MethodsResponseSingle struct {
	Status  bool    `json:"status"`
	Message string  `json:"message"`
	Data    Methods `json:"data"`
}
