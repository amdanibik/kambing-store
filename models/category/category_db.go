package category

// user
type Category struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
}
