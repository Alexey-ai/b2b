package service

import (
	"time"
)

type M08Handler struct {
	SbpApiSender *SbpApiSender
	ctx          *EmulatorContext
}

func NewM08Handler(ctx *EmulatorContext) *M08Handler {
	return &M08Handler{ctx: ctx}
}

func (handler *M08Handler) Handle() {
	NewB2BApiSender().SendAsync(handler.ctx.PrepareMessage(NewB2BApiSenderMessage(handler.ctx.M21, "M21", handler.ctx)))
	time.Sleep(300)
	NewB2BApiSender().SendAsync(handler.ctx.PrepareMessage(NewB2BApiSenderMessage(handler.ctx.M23, "M23", handler.ctx)))
}
