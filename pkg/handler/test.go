package handler

import (
	"net/http"

	"github.com/Alexey-ai/b2b"
	"github.com/gin-gonic/gin"
)

func (h *Handler) testPost(c *gin.Context) {

	var input b2b.Request

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Value == "" {
		newErrorResponse(c, http.StatusInternalServerError, "check value")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    input.Id,
		"Value": input.Value,
	})
}

func (h *Handler) testGet(c *gin.Context) {
	newTestResponse(c, http.StatusAccepted, "it work's")
}
