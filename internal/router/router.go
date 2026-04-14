package router

import (
	"github.com/Chihaya-Anon123/TicketHub/internal/api"
	"github.com/Chihaya-Anon123/TicketHub/internal/config"
	"github.com/gin-gonic/gin"
)

// 设置路由
func SetupRouter(cfg config.JWTConfig) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")
	{
		authGroup := apiV1.Group("/auth")
		{
			authGroup.POST("/register", api.Register)
		}
	}
	return r
}
