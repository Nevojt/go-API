package controllers

import (
	"api/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	user := new(models.Users)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser, err := models.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"ID":       newUser.ID,
		"userName": newUser.UserName,
		"email":    newUser.Email,
		"role":     newUser.Role,
		"isActive": newUser.IsActive,
	})

}

// GetAllUsers Get all users
func GetAllUsers(c *gin.Context) {
	users := models.GetAllUsersResponse()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
