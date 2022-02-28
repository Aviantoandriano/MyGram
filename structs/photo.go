package structs

import "time"

type Photo struct {
	Id         int       `gorm:"primarykey;autoIncrement:true;column:id"`
	Title      string    `gorm:"type:varchar(100);column:title"`
	Caption    string    `gorm:"type:varchar(100);column:caption"`
	PhotoUrl   string    `gorm:"type:varchar(100);column:photo_url"`
	UserId     int       `gorm:"column:user_id"`
	Created_At time.Time `gorm:"column:created_at"`
	Updated_At time.Time `gorm:"column:updated_at"`
	User       User      `gorm:"foreignKey:UserId"`
}

type PhotoGetResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserGetResponse
}

type PhotoResponse struct {
	Id         int       `gorm:"primarykey;column:id;autoIncrement:true"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	PhotoUrl   string    `json:"photo_url"`
	UserId     int       `json:"user_id"`
	Created_At time.Time `json:"created_at"`
}

type PhotoResponseCom struct {
	Id       int    `json:"id,omitempty"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}

type PhotoUpdateResponse struct {
	Id         int       `gorm:"primarykey;column:id;autoIncrement:true"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	PhotoUrl   string    `json:"photo_url"`
	UserId     int       `json:"user_id"`
	Updated_At time.Time `json:"updated_at"`
}
