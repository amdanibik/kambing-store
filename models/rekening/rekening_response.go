package rekening

type RekeningResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []Rekening `json:"data"`
}

type RekeningResponseSingle struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    Rekening `json:"data"`
}
