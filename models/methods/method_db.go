package methods

// user
type Methods struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
}
