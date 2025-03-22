package handlers

import (
	"MSaaS-Framework/MSaaS/pkg/object"
	uuid "github.com/google/uuid"

	"net/http"

	"github.com/gin-gonic/gin"

	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// RegisterUserRoutes registers the CRUD routes for User
func (h *UserHandler) RegisterUserRoutes(router *gin.Engine) {
	router.POST("/user", h.CreateUser)
	router.GET("/user/:id", h.GetUser)
	router.PUT("/user/:id", h.UpdateUser)
	router.DELETE("/user/:id", h.DeleteUser)
}

// CreateUser creates a new User
func (h *UserHandler) CreateUser(c *gin.Context) {
	userFromClient := object.User{}

	if err := c.ShouldBindJSON(&userFromClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	user, err := h.service.CreateUser(c, userFromClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUser gets a User by ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	user, err := h.service.GetUserByID(c, uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser updates an existing User by ID
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	userFromClient := object.User{}

	if err := c.ShouldBindJSON(&userFromClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	user, err := h.service.UpdateUser(c, uuid, userFromClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes an existing User by ID
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.service.DeleteUser(c, uuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.String(http.StatusOK, "Delete User with ID: %s", id)
}
