package entity

import "time"

type UserProfile struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Gender    string    `json:"gender" gorm:"type:varchar(20);not null"`
	Phone     string    `json:"phone" gorm:"type:varchar(20);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type User struct {
	UID        string      `json:"uid" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null;unique"`
	Email      string      `json:"email" gorm:"type:varchar(100);not null"`
	Password   string      `json:"password" gorm:"type:varchar(60);not null"`
	Profile_id int         `json:"profile_id" gorm:"type:int;not null"`
	Profile    UserProfile `json:"profile" gorm:"foreignKey:Profile_id"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserAddress struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	User_id   string    `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:User_id"`
	Address   string    `json:"address" gorm:"type:varchar(200);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type RequestUser struct {
	Email    string `form:"email" binding:"required" validate:"email,required"`
	Password string `form:"password" binding:"required" validate:"required"`
	Name     string `form:"name" binding:"required" validate:"required"`
	Phone    string `form:"phone" binding:"required" validate:"required"`
	Gender   string `form:"gender" binding:"required" validate:"required"`
}

type ResponseUser struct {
	UID       string    `json:"uid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
}

type RequestLoginUser struct {
	Email    string `form:"email" binding:"required" validate:"email,required"`
	Password string `form:"password" binding:"required" validate:"required"`
}

type ResponseUserLogin struct {
	Token string `json:"token"`
}

type RequestUserAddress struct {
	User_id string `form:"user_id" binding:"required" validate:"required"`
	Address string `form:"address" binding:"required" validate:"required"`
}

type ResponseUserAddress struct {
	ID        int       `json:"id"`
	User_id   string    `json:"user_id"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}
