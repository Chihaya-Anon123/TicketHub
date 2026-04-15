package router

import (
	"github.com/Chihaya-Anon123/TicketHub/internal/api"
	"github.com/Chihaya-Anon123/TicketHub/internal/config"
	"github.com/Chihaya-Anon123/TicketHub/internal/middleware"
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
			authGroup.POST("/login", api.Login)
		}

		projectGroup := apiV1.Group("/projects")
		projectGroup.Use(middleware.JWTAuth(cfg))
		{
			projectGroup.POST("", api.CreateProject)
			projectGroup.GET("", api.ListProjects)

			projectGroup.POST("/:projectId/members", api.AddProjectMember)

		}
	}
	return r
}
