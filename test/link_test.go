package test

import (
	"encoding/json"
	"fmt"
	"github.com/kingjarry/ding-robot/robot"
	"testing"
)

func Test_NewLinkMessage(t *testing.T) {
	builder := robot.NewLinkMessageBuilder().BuildMessageUrl("message url test").BuildText("Text test").BuildTitle("title test").BuildPicUrl("pic url test")
	message := builder.Build()
	bytes, _ := json.Marshal(message)
	fmt.Println(string(bytes))
}
