package handler

import (
	"github.com/Alexey-ai/b2b/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	t := router.Group("/:version/request/:operationtype")
	{
		t.POST("/:endpoint/:txId", h.B2BPost)
		t.POST("/:endpoint/:txId/:messageType", h.B2BPost)
		t.POST("/testPost", h.testPost)
		t.GET("/test", h.testGet)
	}
	return router
}
