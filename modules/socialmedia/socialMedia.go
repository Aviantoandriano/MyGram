package socialmedia

import (
	"context"
	"final-project/helpers"
	"final-project/structs"
	"time"

	"gorm.io/gorm/clause"
)

func CreateSocialMedia(request structs.ValidationSocialMedia, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var newSocialMedia structs.SocialMedia
	var response structs.SocialMediaResponse
	newSocialMedia = structs.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
		UserId:         request.UserId,
		Created_At:     time.Now(),
	}
	db = db.WithContext(ctx).Create(&newSocialMedia)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Add Social Media Error", 400)
		return
	}
	response = structs.SocialMediaResponse{
		Id:             newSocialMedia.Id,
		Name:           newSocialMedia.Name,
		SocialMediaUrl: newSocialMedia.SocialMediaUrl,
		UserId:         request.UserId,
		Created_At:     newSocialMedia.Created_At,
	}
	dataResponse = helpers.DefResponse(response, nil, 201)
	return
}

func GetSocialMedia(userId int, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var socialMedia []structs.SocialMedia
	var store structs.SocialMediaGetResponse
	var response []structs.SocialMediaGetResponse
	db = db.WithContext(ctx).Preload(clause.Associations).Where("user_id = ?", userId).Find(&socialMedia)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Social Media Not Found", 404)
		return
	}
	for _, val := range socialMedia {
		store = structs.SocialMediaGetResponse{
			Id:             val.Id,
			Name:           val.Name,
			SocialMediaUrl: val.SocialMediaUrl,
			UserId:         userId,
			Created_At:     val.Created_At,
			Updated_At:     val.Updated_At,
			User: structs.UserUpdateResponse{
				Id:       val.User.Id,
				Username: val.User.Username,
				Age:      val.User.Age,
			},
		}
		response = append(response, store)
	}
	dataResponse = helpers.DefResponse(response, nil, 200)
	return
}

func UpdateSocialMedia(request structs.ValidationSocialMediaUpdate, socialMediaId int, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var socialMedia structs.SocialMedia
	var response structs.SocialMediaUpdateResponse
	db = db.WithContext(ctx).Where("id = ?", socialMediaId).First(&socialMedia)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Social Media Not Found", 404)
		return
	}
	socialMedia.Name = request.Name
	socialMedia.SocialMediaUrl = request.SocialMediaUrl
	socialMedia.Updated_At = time.Now()
	db = db.WithContext(ctx).Where("id = ?", socialMediaId).Updates(&socialMedia)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Update Social Media error", 400)
		return
	}
	response = structs.SocialMediaUpdateResponse{
		Id:             socialMedia.Id,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		Updated_At:     socialMedia.Updated_At,
	}
	dataResponse = helpers.DefResponse(response, nil, 200)
	return
}

func DeleteSocialMedia(socialMediaId int, ctx context.Context) (dataResponse structs.Response) {
	var db = helpers.DB
	var response map[string]string
	db = db.WithContext(ctx).Debug().Exec("delete from social_media where id = ?", socialMediaId)
	if db.RowsAffected <= 0 {
		dataResponse = helpers.DefResponse(nil, "Social Media Not Found", 404)
		return
	}
	response = map[string]string{
		"message": "Your Social Media Has Been Successfully Deleted",
	}
	dataResponse = helpers.DefResponse(response, nil, 200)
	return
}
