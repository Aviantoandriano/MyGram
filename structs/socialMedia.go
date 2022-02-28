package structs

import "time"

type SocialMedia struct {
	Id             int       `gorm:"primarykey;autoIncrement:true;column:id"`
	Name           string    `gorm:"type:varchar(20);column:name"`
	SocialMediaUrl string    `gorm:"type:varchar(100);column:social_media_url"`
	UserId         int       `gorm:"column:user_id"`
	Created_At     time.Time `gorm:"column:created_at"`
	Updated_At     time.Time `gorm:"column:updated_at"`
	User           User      `gorm:"foreignKey:UserId"`
}

type SocialMediaResponse struct {
	Id             int       `gorm:"primarykey;column:id;autoIncrement:true"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	Created_At     time.Time `gorm:"column:created_at"`
}

type SocialMediaUpdateResponse struct {
	Id             int       `gorm:"primarykey;column:id;autoIncrement:true"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	Updated_At     time.Time `gorm:"column:updated_at"`
}

type SocialMediaGetResponse struct {
	Id             int       `gorm:"primarykey;autoIncrement:true;column:id"`
	Name           string    `gorm:"type:varchar(20);column:name"`
	SocialMediaUrl string    `gorm:"type:varchar(100);column:social_media_url"`
	UserId         int       `gorm:"column:user_id"`
	Created_At     time.Time `gorm:"column:created_at"`
	Updated_At     time.Time `gorm:"column:updated_at"`
	User           UserUpdateResponse
}
