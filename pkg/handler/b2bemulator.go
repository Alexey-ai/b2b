package handler

import (
	"io"
	"net/http"

	//"os"
	"regexp"
	"strings"

	"github.com/Alexey-ai/b2b"
	"github.com/Alexey-ai/b2b/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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

func (h *Handler) B2BPost(c *gin.Context) {
	defer c.Request.Body.Close()
	body, err := io.ReadAll(c.Request.Body)
	var requestBody string = string(body)
	r, _ := regexp.Compile(`MsgId>[A-z0-9]{32}<`)
	n, _ := regexp.Compile(`[A-z0-9]{32}`)
	var msgId string = r.FindString(requestBody)
	if c.Param("txId") == "M05" {
		c.AddParam("messageType", "M05")
	}
	log.Info("got " + c.Param("messageType"))
	txId := c.Param("txId")
	ctx := service.GetCtxbyType(c.Param("messageType"))
	ctx.ContentType = c.Request.Header.Get("Content-Type")
	ctx.OriginalMessageId = n.FindString(msgId)

	if len(txId) > 0 && txId != "M05" {
		ctx.TxId = txId
	} else {
		ctx.TxId = strings.Replace(uuid.New().String(), "-", "", -1)
	}
	c.Writer.Header().Add("X-SBP-TRN-NUM", ctx.TxId)
	c.AbortWithStatus(200)
	//if strings.Contains(ctx.ContentType, "stream") {
	//	body = *Cryptography.NewCryptographyService(cfg).Decrypt(&body)
	//}

	if err != nil {
		log.Error("error parse: %v", err)
		return
	}

	switch c.Param("messageType") {
	case "M05":
		go service.NewM05Handler(ctx).Handle()
	case "M08":
		go service.NewM08Handler(ctx).Handle()
	case "M11":
		go service.NewM11Handler(ctx).Handle()
	case "M13":
		go service.NewM13Handler(ctx).Handle()
	case "M22":
		log.Info("got M22")
	case "M24":
		log.Info("got M24")
	default:
		log.Info("Unsupported request " + c.Request.URL.String())
	}
}
