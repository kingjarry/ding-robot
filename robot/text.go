package robot

import (
	"encoding/json"
)

type textMessage struct {
	MsgType string `json:"msgtype"`
	text    `json:"text"`
	at      `json:"at"`
}

func (m textMessage) Send(webHook string) SendResult {
	bytes, _ := json.Marshal(m)
	content := string(bytes)
	return send(webHook, content)
}

type text struct {
	Content string `json:"content"`
}

type at struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type textMessageBuilder struct {
	message *textMessage
}

func (builder textMessageBuilder) Build() Message {
	return builder.message
}

func NewTextMessageBuilder() *textMessageBuilder {
	return &textMessageBuilder{message: &textMessage{MsgType: "text"}}
}

func (builder *textMessageBuilder) BuildContent(content string) *textMessageBuilder {
	builder.message.Content = content
	return builder
}

func (builder *textMessageBuilder) BuildAtMobiles(mobiles ...string) *textMessageBuilder {
	for _, mobile := range mobiles {
		builder.message.AtMobiles = append(builder.message.AtMobiles, mobile)
	}
	return builder
}

func (builder *textMessageBuilder) BuildIsAtAll(isAtAll bool) *textMessageBuilder {
	builder.message.IsAtAll = isAtAll
	return builder
}
