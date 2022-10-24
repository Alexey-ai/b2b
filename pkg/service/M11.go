package service

type M11Handler struct {
	SbpApiSender *SbpApiSender
	ctx          *EmulatorContext
}

func NewM11Handler(ctx *EmulatorContext) *M11Handler {
	return &M11Handler{ctx: ctx}
}

func (handler *M11Handler) Handle() {
	NewB2BApiSender().SendAsync(handler.ctx.PrepareMessage(NewB2BApiSenderMessage(handler.ctx.M12, "M12", handler.ctx)))
}
