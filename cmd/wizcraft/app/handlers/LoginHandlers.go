package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/user"
	"MSaaS-Framework/MSaaS/pkg/base"

	"MSaaS-Framework/MSaaS/internal/security"
)

// LoginRequest 구조체 추가 (이메일과 비밀번호를 JSON Body에서 받기 위함)
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         *ent.User
}

// PostLogin gets a Login by ID
func PostLogin(c *gin.Context) {

	var req LoginRequest
	var respond LoginResponse

	// JSON Body 바인딩
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	client, err := base.GetDBClientFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get database client"})
		return
	}

	tx, err := client.Tx(c) // 트랜잭션 시작
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to start transaction"})
		return
	}

	// get the user from the database
	userFromDB, err := tx.User.Query().Where(user.EmailEQ(req.Email)).Only(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
		return
	}

	// check password
	check, err := security.VerifyPassword(req.Password, userFromDB.Password, userFromDB.Salt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error verifying password"})
		return
	}
	if !check {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	respond.User = userFromDB

	// access token & refresh token 생성

	// commit the transaction
	if err := tx.Commit(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to commit transaction"})
		return
	}

	// return the user
	c.JSON(http.StatusOK, respond)
}

// GetLogout gets a Logout by ID
func GetLogout(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add get logic
	c.String(http.StatusOK, "Get Logout with ID: %s", id)
}

// RegisterLoginRoutes registers the CRUD routes for Login
func RegisterLoginRoutes(router *gin.Engine) {
	router.GET("/logout/:id", GetLogout)
	router.POST("/login", PostLogin) // POST 요청으로 변경
}
