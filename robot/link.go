package robot

import (
	"encoding/json"
)

type linkMessage struct {
	MsgType string `json:"msgtype"`
	link    `json:"link"`
}

type link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageUrl string `json:"messageUrl"`
	PicUrl     string `json:"picUrl"`
}

func (m linkMessage) Send(webHook string) SendResult {
	bytes, _ := json.Marshal(m)
	content := string(bytes)
	return send(webHook, content)
}

type linkMessageBuilder struct {
	message *linkMessage
}

func NewLinkMessageBuilder() *linkMessageBuilder {
	return &linkMessageBuilder{message: &linkMessage{MsgType: "link"}}
}

func (builder *linkMessageBuilder) BuildTitle(title string) *linkMessageBuilder {
	builder.message.Title = title
	return builder
}

func (builder *linkMessageBuilder) BuildText(text string) *linkMessageBuilder {
	builder.message.Text = text
	return builder
}

func (builder *linkMessageBuilder) BuildMessageUrl(messageUrl string) *linkMessageBuilder {
	builder.message.MessageUrl = messageUrl
	return builder
}

func (builder *linkMessageBuilder) BuildPicUrl(picUrl string) *linkMessageBuilder {
	builder.message.PicUrl = picUrl
	return builder
}

func (builder linkMessageBuilder) Build() Message {
	return builder.message
}
