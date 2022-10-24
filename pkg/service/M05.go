package service

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type M05Handler struct {
	SbpApiSender *SbpApiSender
	ctx          *EmulatorContext
}

func NewM05Handler(ctx *EmulatorContext) *M05Handler {
	return &M05Handler{ctx: ctx}
}

func (handler *M05Handler) Handle() {
	fmt.Println("M05 Handle")
	NewB2BApiSender().SendAsync(handler.ctx.PrepareMessage(NewB2BApiSenderMessage(handler.ctx.M06, "M06", handler.ctx)))
	if viper.GetBool("sendm07") {
		time.Sleep(150)
		NewB2BApiSender().SendAsync(handler.ctx.PrepareMessage(NewB2BApiSenderMessage(handler.ctx.M07, "M07", handler.ctx)))
	}
}
