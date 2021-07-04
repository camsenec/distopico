package api

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageDatabase interface {
	GetAllMessages() (interface{}, error)
	GetMessageById(id primitive.ObjectID) (interface{}, error)
}

type MessageAPI struct {
	DB MessageDatabase
}

func (a *MessageAPI) GetMessages(ctx *gin.Context) {
	res, err := a.DB.GetAllMessages()
	if err != nil {
		ctx.AbortWithError(400, errors.New("Object not found"))
	}
	ctx.JSON(200, res)

}

func (a *MessageAPI) GetMessageByID(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(400, errors.New("Invalid id"))
	}
	fmt.Println(id)
	res, err := a.DB.GetMessageById(id)
	if err != nil {
		ctx.AbortWithError(400, errors.New("Object not found"))
	}
	ctx.JSON(200, res)
}
