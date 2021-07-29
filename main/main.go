package main

import (
	"log-agent"
)

func main() {
	config := make(map[string]interface{})
	config["clienttype"] = "kinesis"
	config["config"] = map[string]string{"streamName":"poc-log-manager", "awsRegion": "us-east-1"}

	logagent.SetConfig(config)
	err := logagent.Dispatch("id3", 123, "12345", "12345", "the action 4", "main-test", "the context 666")
	if err != nil {
		return
	}
}
