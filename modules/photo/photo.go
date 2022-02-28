package photo

import (
	"context"
	"final-project/helpers"
	"final-project/structs"
	"time"

	"gorm.io/gorm/clause"
)

func AddPhoto(request structs.ValidationPhoto, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var newPhoto structs.Photo
	var response structs.PhotoResponse
	newPhoto = structs.Photo{
		Title:      request.Title,
		Caption:    request.Caption,
		PhotoUrl:   request.PhotoUrl,
		Created_At: time.Now(),
		UserId:     request.UserId,
	}
	db = db.WithContext(ctx).Create(&newPhoto)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Add Photo Error", 400)
		return
	}
	response = structs.PhotoResponse{
		Id:         newPhoto.Id,
		Title:      newPhoto.Title,
		Caption:    newPhoto.Caption,
		PhotoUrl:   newPhoto.PhotoUrl,
		UserId:     newPhoto.UserId,
		Created_At: newPhoto.Created_At,
	}
	dataResponse = helpers.DefResponse(response, nil, 201)
	return
}

func GetPhotos(userId int, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var photos []structs.Photo
	var store structs.PhotoGetResponse
	var response []structs.PhotoGetResponse
	db = db.WithContext(ctx).Preload(clause.Associations).Where("user_id = ?", userId).Find(&photos)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Photo Not Found", 404)
		return
	}
	for _, val := range photos {
		store = structs.PhotoGetResponse{
			ID:        val.Id,
			Title:     val.Title,
			Caption:   val.Caption,
			PhotoUrl:  val.PhotoUrl,
			UserId:    val.UserId,
			CreatedAt: val.Created_At,
			UpdatedAt: val.Updated_At,
			User: structs.UserGetResponse{
				Username: val.User.Username,
				Email:    val.User.Email,
			},
		}
		response = append(response, store)
	}
	dataResponse = helpers.DefResponse(response, nil, 200)
	return
}

func UpdatePhoto(request structs.ValidationPhoto, idPhoto int, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var photo structs.Photo
	var response structs.PhotoUpdateResponse
	db = db.WithContext(ctx).Where("id = ?", idPhoto).First(&photo)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Photo Not Found", 404)
		return
	}
	photo = structs.Photo{
		Title:      request.Title,
		Caption:    request.Caption,
		PhotoUrl:   request.PhotoUrl,
		Updated_At: time.Now(),
	}
	db = db.WithContext(ctx).Where("id = ?", idPhoto).Updates(&photo)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Update Photo error", 400)
		return
	}
	response = structs.PhotoUpdateResponse{
		Id:         idPhoto,
		Title:      photo.Title,
		Caption:    photo.Caption,
		PhotoUrl:   photo.PhotoUrl,
		UserId:     request.UserId,
		Updated_At: photo.Updated_At,
	}
	dataResponse = helpers.DefResponse(response, nil, 200)
	return
}

func DeletePhoto(photoId int, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var response map[string]string
	db = db.WithContext(ctx).Debug().Exec("delete from photos where id = ?", photoId)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Photo Not Found", 404)
		return
	}
	response = map[string]string{
		"message": "Your Photo Has Been Successfully Deleted",
	}
	dataResponse = helpers.DefResponse(response, nil, 200)
	return
}
