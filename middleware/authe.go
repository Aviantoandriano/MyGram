package middleware

import (
	"encoding/json"
	"final-project/helpers"
	"final-project/structs"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
	_context "github.com/gorilla/context"
)

func Authe(ctx *context.Context) {
	fmt.Println("masuk")
	var db = helpers.DB
	var response map[string]interface{}
	var currentUser structs.User

	jwtToken := ctx.Request.Header.Get("Authorization")
	fmt.Println(jwtToken)
	if !strings.Contains(jwtToken, "Bearer") {
		responData := helpers.DefResponse(nil, "Unauthorized", 401)
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(responData, true, true)
		return
	}

	token, err := helpers.VerifyToken(ctx)
	if err != nil {
		fmt.Println(err)
		responData := helpers.DefResponse(nil, "Unauthorized", 401)
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(responData, true, true)
		return
	}

	data, err := json.Marshal(token)
	if err != nil {
		responData := helpers.DefResponse(nil, "Internal Server Error", 500)
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(responData, true, true)
		return
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		responData := helpers.DefResponse(nil, "Internal Server Error", 500)
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(responData, true, true)
		return
	}

	userId := response["id"]
	id := int(userId.(float64))

	db = db.WithContext(ctx.Request.Context()).Where("id = ?", id).First(&currentUser)
	if db.RowsAffected <= 0 {
		dataResponse := helpers.DefResponse(nil, "access denied", 401)
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(dataResponse, true, true)
		return
	}
	_context.Set(ctx.Request, "userLogin", currentUser.Id)
}
