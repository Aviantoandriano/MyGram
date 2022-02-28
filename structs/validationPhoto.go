package structs

type ValidationPhoto struct {
	Title    string `valid:"Required"`
	Caption  string
	PhotoUrl string `valid:"Required"`
	UserId   int
}

type ValidatePhotoUpdate struct {
	Title    string `valid:"Required"`
	Caption  string
	PhotoUrl string `valid:"Required"`
}
