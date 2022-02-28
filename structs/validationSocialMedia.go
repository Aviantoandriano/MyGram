package structs

type ValidationSocialMedia struct {
	Name           string `valid:"Required"`
	SocialMediaUrl string `valid:"Required"`
	UserId         int
}

type ValidationSocialMediaUpdate struct {
	Name           string `valid:"Required"`
	SocialMediaUrl string `valid:"Required"`
}
