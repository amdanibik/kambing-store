package category

// user

type CategoryRegister struct {
	Name string `form:"name"`
}

type CategoryUpdater struct {
	Name string `form:"name"`
}
