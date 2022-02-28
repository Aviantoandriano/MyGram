package middleware

import (
	"encoding/json"
	"final-project/helpers"
	"final-project/structs"
	"fmt"
	"strconv"

	"github.com/beego/beego/v2/server/web/context"
	_context "github.com/gorilla/context"
)

func AuthoPhoto(ctx *context.Context) {
	var db = helpers.DB
	var currentUser structs.Photo
	var userId int
	//user login
	userLogin := _context.Get(ctx.Request, "userLogin")
	data, err := json.Marshal(userLogin)
	if err != nil {
		fmt.Println(1)
		dataResponse := helpers.DefResponse(nil, "Internal Service Error", 500)
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}
	err = json.Unmarshal(data, &userId)
	fmt.Println(userId, "sini wey")
	if err != nil {
		fmt.Println(err)
		dataResponse := helpers.DefResponse(nil, "Internal Service Error", 500)
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}

	idParam := ctx.Input.Param(":id")
	photoId, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println(3)
		dataResponse := helpers.DefResponse(nil, "Internal Service Error", 500)
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}

	db = db.WithContext(ctx.Request.Context()).Where("id = ?", photoId).First(&currentUser)
	if db.RowsAffected <= 0 {
		dataResponse := helpers.DefResponse(nil, "Photo Not Found", 404)
		ctx.Output.SetStatus(404)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}
	if currentUser.UserId != userId {
		dataResponse := helpers.DefResponse(nil, "Access Forbiden", 403)
		ctx.Output.SetStatus(403)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}

}

func AuthoComment(ctx *context.Context) {
	var db = helpers.DB
	var currentUser structs.Comment
	var userId int

	userLogin := _context.Get(ctx.Request, "userLogin")
	data, err := json.Marshal(userLogin)
	if err != nil {
		dataResponse := helpers.DefResponse(nil, "Internal Service Error", 500)
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}

	err = json.Unmarshal(data, &userId)
	if err != nil {
		dataResponse := helpers.DefResponse(nil, "Internal Service Error", 500)
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}

	idParam := ctx.Input.Param(":id")
	idComment, err := strconv.Atoi(idParam)
	if err != nil {
		dataResponse := helpers.DefResponse(nil, "Internal Service Error", 500)
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}

	db = db.WithContext(ctx.Request.Context()).Where("id = ?", idComment).First(&currentUser)
	if db.RowsAffected <= 0 {
		dataResponse := helpers.DefResponse(nil, "Comment Not Found", 404)
		ctx.Output.SetStatus(404)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}
	if currentUser.UserId != userId {
		dataResponse := helpers.DefResponse(nil, "Access Forbiden", 403)
		ctx.Output.SetStatus(403)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}

}

func AuthoSocialMedia(ctx *context.Context) {
	var db = helpers.DB
	var currentUser structs.SocialMedia
	var userId int

	userLogin := _context.Get(ctx.Request, "userLogin")
	data, err := json.Marshal(userLogin)
	if err != nil {
		dataResponse := helpers.DefResponse(nil, "Internal Service Error", 500)
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}

	err = json.Unmarshal(data, &userId)
	if err != nil {
		dataResponse := helpers.DefResponse(nil, "Internal Service Error", 500)
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}

	paramId := ctx.Input.Param(":id")
	idSocialMedia, err := strconv.Atoi(paramId)
	if err != nil {
		dataResponse := helpers.DefResponse(nil, "Internal Service Error", 500)
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}
	db = db.WithContext(ctx.Request.Context()).Where("id = ?", idSocialMedia).First(&currentUser)
	if db.RowsAffected <= 0 {
		dataResponse := helpers.DefResponse(nil, "Social Media Not Found", 404)
		ctx.Output.SetStatus(404)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}
	if currentUser.UserId != userId {
		dataResponse := helpers.DefResponse(nil, "Access Forbiden", 403)
		ctx.Output.SetStatus(403)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}

}
