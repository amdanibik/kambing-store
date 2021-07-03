package product

type Product struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Category_id uint   `json:"category_id"`
	Kandang     string `json:"kandang"`
	Jumlah      int    `json:"jumlah"`
	Hargamin    int    `json:"hargamin"`
	Hargamax    int    `json:"hargamax"`
}
