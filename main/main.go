package main

import (
	logagent "dacast-log-agent"
	"dacast-log-agent/config"
)

func main() {
	logagent.SetConfig(&config.ClientConfig{ClientType: "kinesis", Config: map[string]string{"streamName": "poc-log-manager"}})
	err := logagent.Dispatch("id3", 123, "12345", "12345", "the action 3", "the context 3")
	if err != nil {
		return
	}
}
