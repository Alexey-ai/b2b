package service

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type SbpApiSender struct {
	b2bApiUrl string
}

func (s *SbpApiSender) SendAsync(data *B2BApiSenderMessage) {
	url := strings.Replace(strings.Replace(s.b2bApiUrl, "{TxId}", data.Ctx.TxId, 1), "{MsgType}", data.MessageType, 1)
	resp, err := http.Post(url, "application/xml", bytes.NewReader([]byte(data.Message)))
	if err != nil {
		println(err.Error())
		return
	}
	defer resp.Body.Close()
}

func NewB2BApiSender() *SbpApiSender {
	return &SbpApiSender{b2bApiUrl: viper.GetString("b2bhost")}
}
