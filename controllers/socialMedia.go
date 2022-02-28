package controllers

import (
	"encoding/json"
	"final-project/helpers"
	"final-project/modules/socialmedia"
	"final-project/structs"
	"strconv"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	_context "github.com/gorilla/context"
)

type SocialMediaController struct {
	beego.Controller
}

func (c *SocialMediaController) CreateSocialMedia() {
	var userId int
	valid := validation.Validation{}
	name := c.GetString("name")
	socialMediaUrl := c.GetString("social_media_url")
	num := _context.Get(c.Ctx.Request, "userLogin")
	data, _ := json.Marshal(num)
	_ = json.Unmarshal(data, &userId)
	request := structs.ValidationSocialMedia{
		UserId:         userId,
		Name:           name,
		SocialMediaUrl: socialMediaUrl,
	}
	if validate, _ := valid.Valid(&request); !validate {
		errValidation := helpers.ErrValidation((valid.Errors))
		response := helpers.DefResponse(nil, helpers.ErrResponse(errValidation), 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	dataResponse := socialmedia.CreateSocialMedia(request, c.Ctx.Request.Context())
	if dataResponse.Data == nil {
		c.Data["json"] = dataResponse
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	c.Data["json"] = dataResponse
	c.Ctx.Output.SetStatus(201)
	c.ServeJSON()
	return
}

func (c *SocialMediaController) GetSocialMedia() {
	var userId int
	num := _context.Get(c.Ctx.Request, "userLogin")
	data, _ := json.Marshal(num)
	_ = json.Unmarshal(data, &userId)
	dataResponse := socialmedia.GetSocialMedia(userId, c.Ctx.Request.Context())
	if dataResponse.Data == nil {
		c.Data["json"] = dataResponse
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	c.Data["json"] = dataResponse
	c.Ctx.Output.SetStatus(200)
	c.ServeJSON()
	return
}

func (c *SocialMediaController) UpdateSocialMedia() {
	valid := validation.Validation{}
	name := c.GetString("name")
	socialMediaUrl := c.GetString("social_media_url")
	num := c.Ctx.Input.Param(":id")
	socialMediaId, _ := strconv.Atoi(num)
	request := structs.ValidationSocialMediaUpdate{
		Name:           name,
		SocialMediaUrl: socialMediaUrl,
	}
	if validate, _ := valid.Valid(&request); !validate {
		errValidation := helpers.ErrValidation((valid.Errors))
		response := helpers.DefResponse(nil, helpers.ErrResponse(errValidation), 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	dataResponse := socialmedia.UpdateSocialMedia(request, socialMediaId, c.Ctx.Request.Context())
	if dataResponse.Data == nil {
		c.Data["json"] = dataResponse
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	c.Data["json"] = dataResponse
	c.Ctx.Output.SetStatus(200)
	c.ServeJSON()
	return
}

func (c *SocialMediaController) DeleteSocialMedia() {
	num := c.Ctx.Input.Param(":id")
	socialMediaId, _ := strconv.Atoi(num)
	dataResponse := socialmedia.DeleteSocialMedia(socialMediaId, c.Ctx.Request.Context())
	if dataResponse.Data == nil {
		c.Data["json"] = dataResponse
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	c.Data["json"] = dataResponse
	c.Ctx.Output.SetStatus(200)
	c.ServeJSON()
	return
}
