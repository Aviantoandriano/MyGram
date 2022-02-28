package controllers

import (
	"encoding/json"
	"final-project/helpers"
	"final-project/modules/comment"
	"final-project/structs"
	"strconv"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	_context "github.com/gorilla/context"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) AddComment() {
	var userId int
	valid := validation.Validation{}
	message := c.GetString("message")
	photoId, _ := c.GetInt("photo_id")
	num := _context.Get(c.Ctx.Request, "userLogin")
	data, _ := json.Marshal(num)
	_ = json.Unmarshal(data, &userId)
	request := structs.ValidationComment{
		UserId:  userId,
		PhotoId: photoId,
		Message: message,
	}
	if validate, _ := valid.Valid(&request); !validate {
		errValidation := helpers.ErrValidation((valid.Errors))
		response := helpers.DefResponse(nil, helpers.ErrResponse(errValidation), 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	dataResponse := comment.CreateComment(request, c.Ctx.Request.Context())
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

func (c *CommentController) GetComments() {
	var userId int
	num := _context.Get(c.Ctx.Request, "userLogin")
	data, _ := json.Marshal(num)
	_ = json.Unmarshal(data, &userId)
	dataResponse := comment.GetComment(userId, c.Ctx.Request.Context())
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

func (c *CommentController) UpdateComment() {
	valid := validation.Validation{}
	message := c.GetString("message")
	num := c.Ctx.Input.Param(":id")
	commentId, _ := strconv.Atoi(num)
	request := structs.ValidationUpdateComment{
		Message: message,
	}
	if validate, _ := valid.Valid(&request); !validate {
		errValidation := helpers.ErrValidation((valid.Errors))
		response := helpers.DefResponse(nil, helpers.ErrResponse(errValidation), 400)
		c.Data["json"] = response
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	dataResponse := comment.UpdateComment(request, commentId, c.Ctx.Request.Context())
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

func (c *CommentController) DeleteComment() {
	num := c.Ctx.Input.Param(":id")
	commentId, _ := strconv.Atoi(num)
	dataResponse := comment.DeleteComment(commentId, c.Ctx.Request.Context())
	if dataResponse.Data == nil {
		// response := helpers.DefResponse(nil, dataResponse.Error, 400)
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
