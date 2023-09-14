package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User Model
type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;<-:create"`
	Username  string    `json:"username" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Role      string    `gorm:"type:varchar(255);not null"`
	IsNew     bool      // is new user, password not changed from default
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserInput struct {
	ID              string `json:"id" binding:"required,uuid4"`
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type SignIn struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	JWT    string `json:"token"`
	UserID string `json:"user_id"`
	IsNew  bool   `json:"is_new"`
}

type ChangeDefaultPassword struct {
	Email           string `json:"email" binding:"required,email"`
	OldPassword     string `json:"old_password" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

//func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
//	user.ID = uuid.New()
//	return
//}

func UserGetAll() (users []User, err error) {
	err = DB.Find(&users).Error

	return users, err
}

func UserGetById(uuid string) (user *User, err error) {
	err = DB.Where("id = ?", uuid).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, err
}

func UserGetByEmail(email string) (user *User, err error) {
	err = DB.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, err
}

func UserCreate(user *User) (_ *User, err error) {
	err = DB.Create(&user).Error

	return user, err
}

func UserSave(updatedUser *User) (_ *User, err error) {
	err = DB.Save(&updatedUser).Error

	return updatedUser, err
}

func UserDelete(user *User) (_ *User, err error) {
	err = DB.Delete(&user).Error

	return user, err
}
