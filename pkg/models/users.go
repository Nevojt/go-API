package models

import (
	"api/pkg/config"
	"api/pkg/util"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Users struct {
	gorm.Model
	UserName string `gorm:"unique;not null" json:"userName"` // Унікальне ім'я користувача.
	Email    string `gorm:"unique;not null" json:"email"`    // Унікальний емейл.
	Password string `json:"password"`                        // Пароль.
	Role     string `json:"role"`                            // Роль користувача.
	IsActive bool   `json:"isActive"`                        // Статус активності користувача.
}

type UserResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	IsActive bool   `json:"isActive"`
}

func init() {
	var err error
	err = config.Connect()
	if err != nil {
		return
	}
	db = config.GetDB()
	err = db.AutoMigrate(&Users{})
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
}

func CreateUser(user *Users) (*UserResponse, error) {

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	if err = db.Create(user).Error; err != nil {
		return nil, err
	}
	return &UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Role:     user.Role,
		IsActive: user.IsActive,
	}, nil
}

func GetAllUsers() []Users {
	var users []Users
	db.Find(&users)
	return users
}

func GetAllUsersResponse() []UserResponse {
	users := GetAllUsers()
	response := make([]UserResponse, len(users))
	for i, u := range users {
		response[i] = UserResponse{
			ID:       u.ID,
			UserName: u.UserName,
			Email:    u.Email,
			Role:     u.Role,
			IsActive: u.IsActive,
		}
	}
	return response
}
