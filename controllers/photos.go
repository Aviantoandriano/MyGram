package controllers

import (
	"encoding/json"
	"final-project/helpers"
	"final-project/modules/photo"
	"final-project/structs"
	"strconv"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	_context "github.com/gorilla/context"
)

type PhotoController struct {
	beego.Controller
}

func (c *PhotoController) AddPhoto() {
	var userId int
	valid := validation.Validation{}
	Title := c.GetString("title")
	Caption := c.GetString("caption")
	PhotoUrl := c.GetString("photo_url")
	num := _context.Get(c.Ctx.Request, "userLogin")
	data, _ := json.Marshal(num)
	_ = json.Unmarshal(data, &userId)
	request := structs.ValidationPhoto{
		Title:    Title,
		Caption:  Caption,
		PhotoUrl: PhotoUrl,
		UserId:   userId,
	}
	if validate, _ := valid.Valid(&request); !validate {
		errValidation := helpers.ErrValidation((valid.Errors))
		response := helpers.DefResponse(nil, helpers.ErrResponse(errValidation), 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	dataResponse := photo.AddPhoto(request, c.Ctx.Request.Context())
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

func (c PhotoController) GetPhotos() {
	var userId int
	num := _context.Get(c.Ctx.Request, "userLogin")
	data, _ := json.Marshal(num)
	_ = json.Unmarshal(data, &userId)
	dataResponse := photo.GetPhotos(userId, c.Ctx.Request.Context())
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

func (c PhotoController) UpdatePhoto() {
	var userId int
	valid := validation.Validation{}
	title := c.GetString("title")
	caption := c.GetString("caption")
	photoUrl := c.GetString("photo_url")
	num := c.Ctx.Input.Param(":id")
	idPhoto, _ := strconv.Atoi(num)
	idJwt := _context.Get(c.Ctx.Request, "userLogin")
	data, _ := json.Marshal(idJwt)
	_ = json.Unmarshal(data, &userId)
	request := structs.ValidationPhoto{
		Title:    title,
		Caption:  caption,
		PhotoUrl: photoUrl,
		UserId:   userId,
	}
	if validate, _ := valid.Valid(&request); !validate {
		errValidation := helpers.ErrValidation((valid.Errors))
		response := helpers.DefResponse(nil, helpers.ErrResponse(errValidation), 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	dataResponse := photo.UpdatePhoto(request, idPhoto, c.Ctx.Request.Context())
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

func (c PhotoController) DeletePhoto() {
	num := c.Ctx.Input.Param(":id")
	idPhoto, _ := strconv.Atoi(num)
	dataResponse := photo.DeletePhoto(idPhoto, c.Ctx.Request.Context())
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
