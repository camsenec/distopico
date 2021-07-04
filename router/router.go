package router

import (
	"github.com/distopico/api"
	"github.com/distopico/database"
	"github.com/gin-gonic/gin"
)

func Create(db *database.MongoDB) *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger(), gin.Recovery())

	messageHandler := api.MessageAPI{DB: db}

	message := g.Group("/message")
	{
		message.GET("", messageHandler.GetMessages)
		message.GET("/:id", messageHandler.GetMessageByID)
		message.POST("", messageHandler.CreateMessage)
	}

	return g
}
