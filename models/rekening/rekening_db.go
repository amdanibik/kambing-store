package rekening

type Rekening struct {
	ID             uint   `gorm:"primarykey" json:"id"`
	Method_id      uint   `json:"methods_id"`
	Nama           string `json:"nama"`
	Nomor_rekening string `json:"nomor_rekening"`
	Bank           string `json:"bank"`
}
