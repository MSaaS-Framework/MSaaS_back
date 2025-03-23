package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"MSaaS-Framework/MSaaS/internal/security"
	"MSaaS-Framework/MSaaS/pkg/base"
	"MSaaS-Framework/MSaaS/pkg/object"
)

func SetUserOptions() {
	// user options
}

func CreateUser(c *gin.Context) {

	// user 객체 생성
	user := object.User{}

	// get project data from request
	c.BindJSON(&user)

	// create salt and hash password
	hash, salt, err := security.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	// Save user to database by using Ent
	// Context에서 DBClient를 가져옴
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

	_, err = tx.User.Create(). // User를 생성하기 위해 ent.User 객체 생성
					SetEmail(user.Email).
					SetPassword(hash).
					SetSalt(salt).
					Save(c)
	if err != nil {
		tx.Rollback() // 실패 시 롤백
		c.JSON(500, gin.H{"error": "Failed to create User"})
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.String(http.StatusOK, "Create User")

}
