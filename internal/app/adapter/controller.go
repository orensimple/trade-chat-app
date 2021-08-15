package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/orensimple/trade-chat-app/internal/app/adapter/mongodb"
	"github.com/orensimple/trade-chat-app/internal/app/adapter/repository"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

// Controller is a controller
type Controller struct {
	ChatRepository    repository.Chat
	MessageRepository repository.Message
}

// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	db := mongodb.Connection()

	// init prometheus metrics
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r)

	chatRepository := repository.NewChatRepo(db)
	messageRepository := repository.NewMessageRepo(db)

	ctrl := Controller{
		ChatRepository:    chatRepository,
		MessageRepository: messageRepository,
	}

	r.GET("/health", ctrl.health)

	api := r.Group("/api")
	{
		api.POST("/chat", ctrl.chatCreate)
		api.GET("/chat/:id", ctrl.chatGet)
		api.GET("/chats/search", ctrl.chatSearch)

		api.POST("/message", ctrl.messageCreate)
		api.GET("/message/:id", ctrl.messageGet)
		api.GET("/messages/search", ctrl.messageSearch)
	}

	return r
}
