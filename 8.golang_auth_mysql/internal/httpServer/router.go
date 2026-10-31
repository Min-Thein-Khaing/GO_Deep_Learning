package httpserver

import (
	"go_auth_mysql/internal/app"
	"go_auth_mysql/internal/middleware"
	"go_auth_mysql/internal/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", health)

	repo := user.NewRepo(a.DB)
	svc := user.NewService(repo, a.Config.JWT_SECRET, a.Config.JWT_REFRESH_SECRET)
	handler := user.NewHandler(svc)

	// Public auth routes
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.POST("/refresh", handler.RefreshToken)

	// Protected routes
	api := r.Group("/api")
	{
		api.Use(middleware.AuthRequired(a.Config.JWT_SECRET))
		api.GET("/files", middleware.RequireRoles("user"), func(c *gin.Context) {
			userID := middleware.GetUserID(c)
			c.JSON(200, gin.H{
				"message": "hello from mysql auth",
				"user_id": userID,
			})
		})

		api.GET("/profile", func(c *gin.Context) {
			userID := middleware.GetUserID(c)
			c.JSON(200, gin.H{
				"message": "user profile accessed",
				"user_id": userID,
			})
		})
	}

	return r
}

