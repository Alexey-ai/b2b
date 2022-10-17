package service

type M05Handler struct {
	SbpApiSender *SbpApiSender
	ctx          *EmulatorContext
}

func NewM05Handler(ctx *EmulatorContext) *M05Handler {
	return &M05Handler{ctx: ctx}
}

func (handler *M05Handler) Handle() {
	NewB2BApiSender().SendAsync(handler.ctx.PrepareMessage(NewB2BApiSenderMessage(handler.ctx.M06, "M06", handler.ctx)))
}
