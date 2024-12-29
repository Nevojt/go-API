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

type UserUpdate struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

// LoginRequest Структура для передачі логіна
type LoginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"password"`
}

// TokenResponse Структура для токену
type TokenResponse struct {
	Token string `json:"token"`
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

func GetUserById(id string) (*UserResponse, error) {
	var user Users
	result := db.Where("id = ?", id).First(&user) // Використання First замість Find
	if result.Error != nil {
		return nil, result.Error // Повертаємо помилку, якщо щось пішло не так
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound // Перевірка чи був знайдений запис
	}

	return &UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Role:     user.Role,
		IsActive: user.IsActive,
	}, nil
}

func GetUserByIdFull(id string) (*Users, error) {
	var user Users
	result := db.Where("id = ?", id).First(&user) // Використання First замість Find
	if result.Error != nil {
		return nil, result.Error // Повертаємо помилку, якщо щось пішло не так
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound // Перевірка чи був знайдений запис
	}

	return &user, nil
}

func GetUserByEmailFull(email string) (*Users, error) {
	var user Users
	result := db.Where("email = ?", email).First(&user) // Використання First замість Find
	if result.Error != nil {
		return nil, result.Error // Повертаємо помилку, якщо щось пішло не так
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound // Перевірка чи був знайдений запис
	}

	return &user, nil
}

func UpdateUser(user *Users) error {

	if err := db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id string) error {
	var user Users
	result := db.Where("id =?", id).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
