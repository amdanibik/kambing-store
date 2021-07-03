package rekening

// user

type RekeningRegister struct {
	Method_id      uint   `form:"methods_id"`
	Nama           string `form:"nama"`
	Nomor_rekening string `form:"nomor_rekening"`
	Bank           string `form:"bank"`
}

type RekeningUpdater struct {
	Method_id      uint   `form:"methods_id"`
	Nama           string `form:"nama"`
	Nomor_rekening string `form:"nomor_rekening"`
	Bank           string `form:"bank"`
}
