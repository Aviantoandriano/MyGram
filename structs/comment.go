package structs

import "time"

type Comment struct {
	Id         int       `gorm:"primarykey;autoIncrement:true;column:id"`
	UserId     int       `gorm:"column:user_id"`
	PhotoId    int       `gorm:"column:photo_id"`
	Message    string    `gorm:"type:varchar(100);column:message"`
	Created_At time.Time `gorm:"column:created_at"`
	Updated_At time.Time `gorm:"column:updated_at"`
	User       User      `gorm:"foreignKey:UserId"`
	Photo      Photo     `gorm:"foreignKey:PhotoId"`
}

type CommentResponse struct {
	Id         int       `gorm:"primarykey;column:id;autoIncrement:true"`
	UserId     int       `json:"user_id"`
	PhotoId    int       `json:"photo_id"`
	Message    string    `json:"message"`
	Update_At  time.Time `json:"updated_at"`
	Created_At time.Time `json:"created_at"`
}

type CommentUpdateResponse struct {
	Id         int       `gorm:"primarykey;column:id;autoIncrement:true"`
	UserId     int       `json:"user_id"`
	PhotoId    int       `json:"photo_id"`
	Message    string    `json:"message"`
	Updated_At time.Time `json:"updated_at"`
}

type CommentRequest struct {
	Id      int    `json:"id,omitempty"`
	UserId  int    `json:"user_id"`
	PhotoId int    `json:"photo_id"`
	Message string `json:"message" valid:"Required"`
}

type CommentGetResponse struct {
	Id         int       `gorm:"primarykey;autoIncrement:true;column:id"`
	UserId     int       `gorm:"column:user_id"`
	PhotoId    int       `gorm:"column:photo_id"`
	Message    string    `gorm:"type:varchar(100);column:message"`
	Created_At time.Time `gorm:"column:created_at"`
	Updated_At time.Time `gorm:"column:updated_at"`
	User       UserComment
	Photo      PhotoResponseCom
}
