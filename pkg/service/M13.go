package service

type M13Handler struct {
	SbpApiSender *SbpApiSender
	ctx          *EmulatorContext
}

func NewM13Handler(ctx *EmulatorContext) *M13Handler {
	return &M13Handler{ctx: ctx}
}

func (handler *M13Handler) Handle() {
	NewB2BApiSender().SendAsync(handler.ctx.PrepareMessage(NewB2BApiSenderMessage(handler.ctx.M14, "M14", handler.ctx)))
}
