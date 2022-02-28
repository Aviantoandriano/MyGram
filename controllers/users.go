package controllers

import (
	"encoding/json"
	"final-project/helpers"
	"final-project/modules/user"
	"final-project/structs"
	"strconv"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	_context "github.com/gorilla/context"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Register() {
	valid := validation.Validation{}
	Age, _ := c.GetInt("age")
	request := structs.ValidatedUser{
		UserName: c.GetString("username"),
		Email:    c.GetString("email"),
		Password: c.GetString("password"),
		Age:      Age,
	}
	if validate, _ := valid.Valid(&request); !validate {
		errValidation := helpers.ErrValidation((valid.Errors))
		response := helpers.DefResponse(nil, helpers.ErrResponse(errValidation), 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	dataResponse := user.Register(request, c.Ctx.Request.Context())
	if dataResponse.Data == nil {
		response := helpers.DefResponse(nil, dataResponse.Error, 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	c.Data["json"] = dataResponse
	c.Ctx.Output.SetStatus(201)
	c.ServeJSON()
	return
}

func (c *UserController) Login() {
	valid := validation.Validation{}
	request := structs.ValidateLogin{
		Email:    c.GetString("email"),
		Password: c.GetString("password"),
	}
	if validate, _ := valid.Valid(&request); !validate {
		errValidation := helpers.ErrValidation((valid.Errors))
		response := helpers.DefResponse(nil, helpers.ErrResponse(errValidation), 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	dataResponse := user.Login(request, c.Ctx.Request.Context())
	if dataResponse.Data == nil {
		response := helpers.DefResponse(nil, dataResponse.Error, 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	c.Data["json"] = dataResponse
	c.Ctx.Output.SetStatus(200)
	c.ServeJSON()
	return
}

func (c UserController) UpdateUsers() {
	valid := validation.Validation{}
	num := c.Ctx.Input.Param(":id")
	idUser, _ := strconv.Atoi(num)
	request := structs.ValidateUserUpdate{
		Id:       idUser,
		Username: c.GetString("username"),
		Email:    c.GetString("email"),
	}
	if validate, _ := valid.Valid(&request); !validate {
		errValidation := helpers.ErrValidation((valid.Errors))
		response := helpers.DefResponse(nil, helpers.ErrResponse(errValidation), 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	dataResponse := user.UpdateUser(request, c.Ctx.Request.Context())
	if dataResponse.Data == nil {
		response := helpers.DefResponse(nil, dataResponse.Error, 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	c.Data["json"] = dataResponse
	c.Ctx.Output.SetStatus(200)
	c.ServeJSON()
	return
}

func (c UserController) DeleteUser() {
	var userId int
	num := _context.Get(c.Ctx.Request, "userLogin")
	data, _ := json.Marshal(num)
	_ = json.Unmarshal(data, &userId)
	dataResponse := user.DeleteUser(userId, c.Ctx.Request.Context())
	if dataResponse.Data == nil {
		response := helpers.DefResponse(nil, dataResponse.Error, 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	c.Data["json"] = dataResponse
	c.Ctx.Output.SetStatus(200)
	c.ServeJSON()
	return
}
