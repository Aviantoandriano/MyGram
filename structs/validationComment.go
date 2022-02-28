package structs

type ValidationComment struct {
	UserId  int
	PhotoId int    `valid:"Required"`
	Message string `valid:"Required"`
}

type ValidationUpdateComment struct {
	Message string `valid:"Required"`
}
