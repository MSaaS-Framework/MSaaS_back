package handlers

import (
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/user"
	"MSaaS-Framework/MSaaS/internal/security"
	"MSaaS-Framework/MSaaS/pkg/base"
	"MSaaS-Framework/MSaaS/pkg/object"

	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser creates a new User
func CreateUser(c *gin.Context) {

	userFromClient := object.User{}

	if err := c.ShouldBindJSON(&userFromClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	client, err := base.GetDBClientFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get database client"})
		return
	}

	tx, err := client.Tx(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to start transaction"})
		return
	}

	// check if user already exists
	_, err = tx.User.Query().Where(user.EmailEQ(userFromClient.Email)).Only(c)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, salt, err := security.HashPassword(userFromClient.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	respond, err := tx.User.Create().
		SetEmail(userFromClient.Email).
		SetPassword(hashedPassword).
		SetSalt(salt).
		Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, respond)
}

// GetUser gets a User by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add get logic
	c.String(http.StatusOK, "Get User with ID: %s", id)
}

// UpdateUser updates an existing User by ID
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add update logic
	c.String(http.StatusOK, "Update User with ID: %s", id)
}

// DeleteUser deletes an existing User by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add delete logic
	c.String(http.StatusOK, "Delete User with ID: %s", id)
}

// RegisterUserRoutes registers the CRUD routes for User
func RegisterUserRoutes(router *gin.Engine) {
	router.POST("/user", CreateUser)
	router.GET("/user/:id", GetUser)
	router.PUT("/user/:id", UpdateUser)
	router.DELETE("/user/:id", DeleteUser)
}
