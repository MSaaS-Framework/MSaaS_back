package app

import (
	"net/http"

	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/handlers"

	"github.com/gin-gonic/gin"
)

// HttpHandlers는 여러 핸들러를 관리하는 슬라이스입니다.
var HttpHandlers []HttpHandler

// HttpHandler 구조체는 각 HTTP 경로와 요청 방식을 관리합니다.
type HttpHandler struct {
	Method   string
	Path     string
	Function func(w http.ResponseWriter, r *http.Request)
}

func defaultRoutes(router *gin.Engine) {
	// 기본 경로 등록
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Wizcraft API 서버에 오신 것을 환영합니다.")
	})

	router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
}

// 등록할 핸들러를 추가하는 함수 예시
func RegisterRoutes(router *gin.Engine) {
	// 기본 경로 등록
	defaultRoutes(router)
	// 핸들러 등록
	handlers.RegisterProjectRoutes(router)
	handlers.RegisterAPIRoutes(router)
	handlers.RegisterDatabaseRoutes(router)
	handlers.RegisterServiceRoutes(router)
	handlers.RegisterLoginRoutes(router)
}
