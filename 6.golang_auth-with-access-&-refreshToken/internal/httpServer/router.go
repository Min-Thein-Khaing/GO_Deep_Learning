package httpserver

import (
	"go_auth/internal/app"
	"go_auth/internal/middleware"
	"go_auth/internal/user"

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


	//unauth
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.POST("/refresh", handler.RefreshToken)


	//list 
	api := r.Group("/api")
	{	
		api.Use(middleware.AuthRequired(a.Config.JWT_SECRET))
		api.GET("/files",middleware.RequireRoles("user"), func(c *gin.Context) {
			userId := middleware.GetUserID(c)
			c.JSON(200, gin.H{
				"message": "hello",
				"user_id": userId,
			})
		})

		api.GET("/product")
	}
	return r
}
