package service

import (
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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
	M14               string `json:"m_14,omitempty"`
	M21               string `json:"m_21,omitempty"`
	M23               string `json:"m_23,omitempty"`
	OriginalMessageId string `json:"original_message_id"`
	MessageId         string `json:"message_id,omitempty"`
	TxId              string `json:"tx_id,omitempty"`
	ContentType       string `json:"content-type,omitempty"`
	Regex             *regexp.Regexp
}

var matched, _ = regexp.Compile(`MsgId>[Aa-z-Z\d]{32}<`)

func NewB2BApiSenderMessage(msg string, msgType string, c *EmulatorContext) *B2BApiSenderMessage {
	return &B2BApiSenderMessage{MessageType: msgType, Message: msg, Ctx: c}
}
func (c *EmulatorContext) PrepareMessage(msg *B2BApiSenderMessage) *B2BApiSenderMessage {
	msg.Message = strings.ReplaceAll(msg.Message, "{OrgMsgId}", strings.ToUpper(c.OriginalMessageId))
	msg.Message = strings.ReplaceAll(msg.Message, "{MsgId}", strings.ToUpper(c.MessageId))
	msg.Message = strings.ReplaceAll(msg.Message, "{TxId}", strings.ToUpper(c.TxId))
	return msg
}

type IPSEnvelope struct {
	OriginalMessageId string `xml:"Document>FIToFICstmrCdtTrf>GrpHdr>MsgId"`
}

func NewEmulatorContext() *EmulatorContext {
	return &EmulatorContext{MessageId: strings.Replace(uuid.New().String(), "-", "", -1), Regex: matched}
}

func GetCtxbyType(msg string) *EmulatorContext {

	e := NewEmulatorContext()
	switch msg {
	case "M05":
		m06, errOpen06 := os.Open("pkg/messages/M06.xml")
		defer m06.Close()
		m07, errOpen07 := os.Open("pkg/messages/M07.xml")
		defer m07.Close()
		file, errRead := io.ReadAll(m06)
		e.M06 = string(file)
		file, _ = io.ReadAll(m07)
		e.M07 = string(file)
		if errOpen06 != nil || errRead != nil || errOpen07 != nil {
			log.Error(errOpen06, errOpen07, errRead)
		}
		return e

	case "M08":
		m21, errOpen21 := os.Open("pkg/messages/M21.xml")
		defer m21.Close()
		m23, errOpen23 := os.Open("pkg/messages/M23.xml")
		defer m23.Close()
		file, errRead := io.ReadAll(m21)
		e.M21 = string(file)
		file, _ = io.ReadAll(m23)
		e.M23 = string(file)
		if errOpen21 != nil || errRead != nil || errOpen23 != nil {
			log.Error(errOpen21, errOpen23, errRead)
		}
		return e
	case "M11":
		m12, errOpen := os.Open("pkg/messages/M12.xml")
		defer m12.Close()
		file, errRead := io.ReadAll(m12)
		e.M12 = string(file)
		if errOpen != nil || errRead != nil {
			log.Error(errOpen, errRead)
		}
		return e
	case "M13":
		m14, errOpen := os.Open("pkg/messages/M14.xml")
		defer m14.Close()
		file, errRead := io.ReadAll(m14)
		e.M14 = string(file)
		if errOpen != nil || errRead != nil {
			log.Error(errOpen, errRead)
		}
		return e
	default:
		log.Info("no need context for" + msg)
		return e
	}
}
