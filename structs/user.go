package structs

import (
	"time"
)

type User struct {
	Id         int       `gorm:"primarykey;autoIncrement:true;"`
	Username   string    `gorm:"type:varchar(20);column:username;unique"`
	Email      string    `gorm:"type:varchar(100);column:email;unique"`
	Password   string    `gorm:"type:varchar(100);column:password"`
	Age        int       `gorm:"column:age"`
	Created_At time.Time `gorm:"column:created_at"`
	Updated_At time.Time `gorm:"column:updated_at"`
}

type UserResponse struct {
	Id       int    `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UserComment struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserGetResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdateResponse struct {
	Id         int       `gorm:"primarykey;column:id;autoIncrement:true"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Age        int       `json:"age"`
	Updated_At time.Time `gorm:"column:updated_at"`
}
