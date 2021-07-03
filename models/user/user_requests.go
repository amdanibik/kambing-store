package user

// user

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}

type UserUpdater struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Mobile   string `form:"mobile"`
}
