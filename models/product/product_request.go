package product

// user

type ProductRegister struct {
	Category_id uint   `form:"category_id"`
	Kandang     string `form:"kandang"`
	Jumlah      int    `form:"jumlah"`
	Hargamin    int    `form:"hargamin"`
	Hargamax    int    `form:"hargamax"`
}

type ProductUpdater struct {
	Category_id uint   `form:"category_id"`
	Kandang     string `form:"kandang"`
	Jumlah      int    `form:"jumlah"`
	Hargamin    int    `form:"hargamin"`
	Hargamax    int    `form:"hargamax"`
}
