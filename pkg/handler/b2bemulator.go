package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Alexey-ai/b2b"
	"github.com/Alexey-ai/b2b/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var M06 string
var M07 string
var M21 string

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
	m06, err := os.Open("messages/M06.xml")
	defer m06.Close()
	m07, err := os.Open("messages/M07.xml")
	defer m07.Close()
	m21, err := os.Open("messages/M21.xml")
	defer m21.Close()
	if err != nil {
		println(err)
	}

	file, err := io.ReadAll(m21)
	M21 = string(file)
	file, err = io.ReadAll(m06)
	M06 = string(file)
	file, err = io.ReadAll(m07)
	M07 = string(file)

	defer c.Request.Body.Close()
	ctx := service.NewEmulatorContext()
	ctx.ContentType = c.Request.Header.Get("Content-Type")
	body, err := io.ReadAll(c.Request.Body)
	txId := c.Param("txId")
	if len(txId) > 0 {
		ctx.TxId = txId
	} else {
		ctx.TxId = strings.Replace(uuid.New().String(), "-", "", -1)
	}
	//httpResponse.Header()["X-SBP-TRN-NUM"] = []string{ctx.TxId}
	//httpResponse.Write([]byte("OK"))
	//if strings.Contains(ctx.ContentType, "stream") {
	//	body = *Cryptography.NewCryptographyService(cfg).Decrypt(&body)
	//}

	if err != nil {
		fmt.Printf("error parse: %v", err)
		return
	}

	ctx.OriginalMessageId = string(ctx.Regex.FindSubmatch(body)[1])
	ctx.M21 = M21
	ctx.M07 = M07
	ctx.M06 = M06

	switch c.Param("messageType") {
	case "M05":
		service.NewM05Handler(ctx).Handle()
	case "M08":
		service.NewM08Handler(ctx).Handle()
	default:
		println("Unsupported request " + c.Request.URL.String())
	}
}
