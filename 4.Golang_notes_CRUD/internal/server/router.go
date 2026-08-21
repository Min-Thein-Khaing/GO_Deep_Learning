package server

import (
	"net/http"
	"notes-api/internal/notes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(database *mongo.Database) *gin.Engine {
	r := gin.Default()

	r.GET("healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"status": "Health",
		})
	})

	notes.RegisterRoutes(r, database)
	return r
}
