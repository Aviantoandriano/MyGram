package routers

import (
	"final-project/controllers"
	"final-project/middleware"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/users/register", &controllers.UserController{}, "post:Register")
	beego.Router("/users/login", &controllers.UserController{}, "post:Login")
	beego.Router("/users/update/:id", &controllers.UserController{}, "put:UpdateUsers")
	beego.Router("/users/detele", &controllers.UserController{}, "delete:DeleteUser")

	beego.Router("/photos/create", &controllers.PhotoController{}, "post:AddPhoto")
	beego.Router("/photos", &controllers.PhotoController{}, "get:GetPhotos")
	beego.Router("/photos/update/:id", &controllers.PhotoController{}, "put:UpdatePhoto")
	beego.Router("/photos/delete/:id", &controllers.PhotoController{}, "delete:DeletePhoto")

	beego.Router("/comments/create", &controllers.CommentController{}, "post:AddComment")
	beego.Router("/comments", &controllers.CommentController{}, "get:GetComments")
	beego.Router("/comments/update/:id", &controllers.CommentController{}, "put:UpdateComment")
	beego.Router("/comments/delete/:id", &controllers.CommentController{}, "delete:DeleteComment")

	beego.Router("/socialmedias/create", &controllers.SocialMediaController{}, "post:CreateSocialMedia")
	beego.Router("/socialmedias", &controllers.SocialMediaController{}, "get:GetSocialMedia")
	beego.Router("/socialmedias/update/:id", &controllers.SocialMediaController{}, "put:UpdateSocialMedia")
	beego.Router("/socialmedias/delete/:id", &controllers.SocialMediaController{}, "delete:DeleteSocialMedia")

	// pemasangan middleware
	//users
	beego.InsertFilter("/users/update/:id", 1, middleware.Authe)
	beego.InsertFilter("/users/delete", 1, middleware.Authe)

	//photos
	beego.InsertFilter("/photos/create", 1, middleware.Authe)
	beego.InsertFilter("/photos", 1, middleware.Authe)
	beego.InsertFilter("/photos/update/:id", 1, middleware.Authe)
	beego.InsertFilter("/photos/update/:id", 2, middleware.AuthoPhoto)
	beego.InsertFilter("/photos/delete/:id", 1, middleware.Authe)
	beego.InsertFilter("/photos/delete/:id", 2, middleware.AuthoPhoto)

	//comments
	beego.InsertFilter("/comments/create", 1, middleware.Authe)
	beego.InsertFilter("/comments", 1, middleware.Authe)
	beego.InsertFilter("/comments/update/:id", 1, middleware.Authe)
	beego.InsertFilter("/comments/update/:id", 2, middleware.AuthoComment)
	beego.InsertFilter("/comments/delete/:id", 1, middleware.Authe)
	beego.InsertFilter("/comments/delete/:id", 2, middleware.AuthoComment)

	//social media
	beego.InsertFilter("/socialmedias/create", 1, middleware.Authe)
	beego.InsertFilter("/socialmedias", 1, middleware.Authe)
	beego.InsertFilter("/socialmedias/update/:id", 1, middleware.Authe)
	beego.InsertFilter("/socialmedias/update/:id", 2, middleware.AuthoSocialMedia)
	beego.InsertFilter("/socialmedias/delete/:id", 1, middleware.Authe)
	beego.InsertFilter("/socialmedias/delete/:id", 2, middleware.AuthoSocialMedia)
}
