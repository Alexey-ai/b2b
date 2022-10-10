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

	test := router.Group("/test")
	{
		test.POST("test", h.testPost)
		test.GET("/", h.testGet)
	}
	return router
}
