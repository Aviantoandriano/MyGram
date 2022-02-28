package comment

import (
	"context"
	"final-project/helpers"
	"final-project/structs"
	"time"

	"gorm.io/gorm/clause"
)

func CreateComment(request structs.ValidationComment, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var photo structs.Photo
	var newPhoto structs.Comment
	var response structs.CommentResponse
	db = db.WithContext(ctx).Where("id = ?", request.PhotoId).First(&photo)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Photo Not Found", 404)
		return
	}
	newPhoto = structs.Comment{
		Message:    request.Message,
		PhotoId:    request.PhotoId,
		UserId:     request.UserId,
		Created_At: time.Now(),
	}
	db = db.WithContext(ctx).Create(&newPhoto)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Add Comment Error", 400)
		return
	}
	response = structs.CommentResponse{
		Id:         newPhoto.Id,
		Message:    newPhoto.Message,
		PhotoId:    request.PhotoId,
		UserId:     request.UserId,
		Created_At: newPhoto.Created_At,
	}
	dataResponse = helpers.DefResponse(response, nil, 201)
	return
}

func GetComment(userId int, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var comments []structs.Comment
	var store structs.CommentGetResponse
	var response []structs.CommentGetResponse
	db = db.WithContext(ctx).Preload(clause.Associations).Where("user_id = ?", userId).Find(&comments)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Comment Not Found", 404)
		return
	}
	for _, val := range comments {
		store = structs.CommentGetResponse{
			Id:         val.Id,
			Message:    val.Message,
			PhotoId:    val.PhotoId,
			UserId:     val.UserId,
			Updated_At: val.Updated_At,
			Created_At: val.Created_At,
			User: structs.UserComment{
				Id:       val.User.Id,
				Email:    val.User.Email,
				Username: val.User.Username,
			},
			Photo: structs.PhotoResponseCom{
				Id:       val.Photo.Id,
				Title:    val.Photo.Title,
				Caption:  val.Photo.Caption,
				PhotoUrl: val.Photo.PhotoUrl,
				UserId:   val.Photo.UserId,
			},
		}
		response = append(response, store)
	}
	dataResponse = helpers.DefResponse(response, nil, 200)
	return
}

func UpdateComment(request structs.ValidationUpdateComment, commentId int, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var comment structs.Comment
	var response structs.CommentUpdateResponse
	db = db.WithContext(ctx).Where("id = ?", commentId).First(&comment)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "comment Not Found", 404)
		return
	}
	comment.Message = request.Message
	comment.Updated_At = time.Now()
	db = db.WithContext(ctx).Where("id = ?", commentId).Updates(&comment)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Update comment error", 400)
		return
	}
	response = structs.CommentUpdateResponse{
		Id:         commentId,
		UserId:     comment.UserId,
		PhotoId:    comment.PhotoId,
		Message:    comment.Message,
		Updated_At: comment.Updated_At,
	}
	dataResponse = helpers.DefResponse(response, nil, 200)
	return
}

func DeleteComment(commentId int, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var response map[string]string
	db = db.WithContext(ctx).Debug().Exec("delete from comments where id = ?", commentId)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Comment Not Found", 404)
		return
	}
	response = map[string]string{
		"message": "Your Comment Has Been Successfully Deleted",
	}
	dataResponse = helpers.DefResponse(response, nil, 200)
	return
}
