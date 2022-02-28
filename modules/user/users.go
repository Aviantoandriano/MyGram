package user

import (
	"context"
	"final-project/helpers"
	"final-project/structs"
	"time"
)

type JwtToken struct {
	Token string
}

func Register(request structs.ValidatedUser, Ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var newUser structs.User
	var response structs.UserResponse

	db = helpers.DB.WithContext(Ctx).Model(&newUser).
		Where("username=?", request.UserName).
		Find(&newUser)
	if db.RowsAffected > 0 {
		dataResponse = helpers.DefResponse(nil, "Username Already Exist", 400)
		return
	}
	db = helpers.DB.WithContext(Ctx).
		Where("email=?", request.Email).
		Find(&newUser)
	if db.RowsAffected > 0 {
		dataResponse = helpers.DefResponse(nil, "Email Already Exist", 400)
		return
	}
	newUser = structs.User{
		Username:   request.UserName,
		Email:      request.Email,
		Password:   helpers.HashPass(request.Password),
		Age:        request.Age,
		Created_At: time.Now(),
	}
	db = helpers.DB.WithContext(Ctx).Create(&newUser)
	if db.Error != nil {
		dataResponse = helpers.DefResponse(nil, "Register Error", 400)
		return
	}
	response = structs.UserResponse{
		Age:      newUser.Age,
		Email:    newUser.Email,
		Id:       newUser.Id,
		Username: newUser.Username,
	}
	dataResponse = helpers.DefResponse(&response, nil, 201)
	return
}

func Login(request structs.ValidateLogin, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var currentUser structs.User
	db = db.WithContext(ctx).Where("email = ?", request.Email).First(&currentUser)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Email/Password salah", 400)
		return
	}
	if passCheker := helpers.PassCheck(currentUser.Password, request.Password); !passCheker {
		dataResponse = helpers.DefResponse(nil, "Email/Password salah", 400)
		return
	}
	token := helpers.GenerateToken(currentUser.Id, currentUser.Email)

	dataResponse = helpers.DefResponse(JwtToken{
		Token: token,
	}, nil, 200)
	return
}

func UpdateUser(request structs.ValidateUserUpdate, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var updatedUser structs.User

	db = db.WithContext(ctx).Where("id = ?", request.Id).First(&updatedUser)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Data Not Found", 404)
		return
	}
	updatedUser.Username = request.Username
	updatedUser.Email = request.Email
	updatedUser.Updated_At = time.Now()

	db = db.WithContext(ctx).Where("id = ?", request.Id).Updates(&updatedUser)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "UserUpdateError", 400)
	}
	result := structs.UserUpdateResponse{
		Id:         updatedUser.Id,
		Username:   updatedUser.Username,
		Email:      updatedUser.Email,
		Age:        updatedUser.Age,
		Updated_At: updatedUser.Updated_At,
	}
	dataResponse = helpers.DefResponse(result, nil, 200)
	return
}

func DeleteUser(userId int, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var response map[string]string
	db = db.WithContext(ctx).Debug().Exec("delete from users where id = ?", userId)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Users Not Found", 404)
		return
	}
	response = map[string]string{
		"message": "Your Account Has Been Successfully Deleted",
	}
	dataResponse = helpers.DefResponse(response, nil, 200)
	return
}
