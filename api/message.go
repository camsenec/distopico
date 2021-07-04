package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/distopico/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageDatabase interface {
	GetAllMessages() (interface{}, error)
	GetMessageById(id primitive.ObjectID) (interface{}, error)
	CreateMessage(message *model.MessageModel) (interface{}, error)
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

func (a *MessageAPI) CreateMessage(ctx *gin.Context) {
	message := model.MessageModel{}
	if err := ctx.ShouldBind(&message); err == nil {
		message.ID = primitive.NewObjectID()
		message.CreatedDate = time.Now()
	} else {
		ctx.AbortWithError(500, errors.New("Bind Failure"))
	}
	_, err := a.DB.CreateMessage(&message)
	if err != nil {
		ctx.AbortWithError(500, errors.New("Object Creation Failure"))
	}

	ctx.JSON(201, message)
}
