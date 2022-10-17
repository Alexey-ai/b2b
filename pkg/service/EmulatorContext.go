package service

import (
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type B2BApiSenderMessage struct {
	MessageType string
	Message     string
	Ctx         *EmulatorContext
}

type EmulatorContext struct {
	M05               string `json:"m_05,omitempty"`
	M06               string `json:"m_06,omitempty"`
	M07               string `json:"m_07,omitempty"`
	M08               string `json:"m_08,omitempty"`
	M11               string `json:"m_11,omitempty"`
	M12               string `json:"m_12,omitempty"`
	M21               string `json:"m_21,omitempty"`
	M23               string `json:"m_23,omitempty"`
	OriginalMessageId string `json:"original_message_id"`
	MessageId         string `json:"message_id,omitempty"`
	TxId              string `json:"tx_id,omitempty"`
	ContentType       string `json:"content-type,omitempty"`
	Regex             *regexp.Regexp
}

var matched, err = regexp.Compile(`MsgId>([Aa-z-Z\d]{32})<`)

func NewB2BApiSenderMessage(msg string, msgType string, c *EmulatorContext) *B2BApiSenderMessage {
	return &B2BApiSenderMessage{MessageType: msgType, Message: msg, Ctx: c}
}
func (c *EmulatorContext) PrepareMessage(msg *B2BApiSenderMessage) *B2BApiSenderMessage {
	msg.Message = strings.ReplaceAll(msg.Message, "{OrgMsgId}", c.OriginalMessageId)
	msg.Message = strings.ReplaceAll(msg.Message, "{MsgId}", c.MessageId)
	msg.Message = strings.ReplaceAll(msg.Message, "{TxId}", c.TxId)
	return msg
}

type IPSEnvelope struct {
	OriginalMessageId string `xml:"Document>FIToFICstmrCdtTrf>GrpHdr>MsgId"`
}

func NewEmulatorContext() *EmulatorContext {
	return &EmulatorContext{MessageId: strings.Replace(uuid.New().String(), "-", "", -1), Regex: matched}
}
