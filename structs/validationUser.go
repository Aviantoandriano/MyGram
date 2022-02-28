package structs

type ValidatedUser struct {
	UserName string `valid:"Required"`
	Email    string `valid:"Required;Email"`
	Password string `valid:"Required;MinSize(6)"`
	Age      int    `valid:"Required;Min(8)"`
}

type ValidateLogin struct {
	Email    string `valid:"Required;Email"`
	Password string `valid:"Required;MinSize(6)"`
}

type ValidateUserUpdate struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username" valid:"Required"`
	Email    string `json:"email" valid:"Required; Email"`
}
