package test

import (
	"encoding/json"
	"github.com/kingjarry/ding-robot/robot"
	"log"
	"testing"
)

func Test_New(t *testing.T) {
	marshal, _ := json.Marshal(robot.NewTextMessageBuilder().Build())
	log.Println(string(marshal))
	message := robot.NewTextMessageBuilder().BuildAtMobiles("123", "345", "456").BuildContent("test message testing").BuildIsAtAll(true).Build()
	bytes, _ := json.Marshal(message)
	log.Println(string(bytes))
}
