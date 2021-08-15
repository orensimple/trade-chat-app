package adapter

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/orensimple/trade-chat-app/internal/app/adapter/mongodb/model"
	"github.com/orensimple/trade-chat-app/internal/app/application/usecase"
	"github.com/orensimple/trade-chat-app/internal/app/domain"
	"github.com/prometheus/common/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ctrl Controller) health(c *gin.Context) {
	c.JSON(http.StatusOK, domain.SimpleResponse{Status: "OK"})
}

func (ctrl Controller) chatCreate(c *gin.Context) {
	var req domain.CreateChatRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	userUUIDs := make([]uuid.UUID, 0, len(req.UserIDs))

	for _, id := range req.UserIDs {
		uuid, err := uuid.Parse(id)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

			return
		}

		userUUIDs = append(userUUIDs, uuid)
	}

	newChat := model.Chat{
		ID:        primitive.NewObjectID(),
		UserIDs:   userUUIDs,
		Label:     req.Label,
		CreatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
	}

	res, err := usecase.CreateChat(ctrl.ChatRepository, &newChat)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed create chat"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) chatGet(c *gin.Context) {
	id := c.Param("id")
	chatID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

		return
	}

	res, err := usecase.GetChat(ctrl.ChatRepository, &model.Chat{ID: chatID})
	if err != nil && err.Error() != "chat not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get chat"})

		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, domain.SimpleResponse{Status: "chat not found"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) chatSearch(c *gin.Context) {
	var req domain.SearchChatRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	userUUIDs := make([]uuid.UUID, 0, len(req.UserIDs))

	for _, id := range req.UserIDs {
		uuid, err := uuid.Parse(id)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

			return
		}

		userUUIDs = append(userUUIDs, uuid)
	}

	res, err := usecase.SearchChat(ctrl.ChatRepository, &model.Chat{UserIDs: userUUIDs})
	if err != nil && err.Error() != "chat not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get chat"})

		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, domain.SimpleResponse{Status: "chat not found"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) messageCreate(c *gin.Context) {
	var req domain.CreateMessageRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong user uuid"})

		return
	}

	chatID, err := primitive.ObjectIDFromHex(req.ChatID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong chat id"})

		return
	}

	newMessage := model.Message{
		ID:        primitive.NewObjectID(),
		ChatID:    chatID,
		UserID:    userID,
		Body:      req.Body,
		CreatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
	}

	res, err := usecase.CreateMessage(ctrl.MessageRepository, &newMessage)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed create message"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) messageGet(c *gin.Context) {
	id := c.Param("id")
	messageID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

		return
	}

	res, err := usecase.GetMessage(ctrl.MessageRepository, &model.Message{ID: messageID})
	if err != nil && err.Error() != "chat not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get message"})

		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, domain.SimpleResponse{Status: "message not found"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) messageSearch(c *gin.Context) {
	var req domain.SearchMessageRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong user id"})

		return
	}

	chatID, err := primitive.ObjectIDFromHex(req.ChatID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong chat id"})

		return
	}

	res, err := usecase.SearchMessage(ctrl.MessageRepository, &model.Message{UserID: userID, ChatID: chatID})
	if err != nil && err.Error() != "chat not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get message"})

		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, domain.SimpleResponse{Status: "message not found"})

		return
	}

	c.JSON(http.StatusOK, res)
}
