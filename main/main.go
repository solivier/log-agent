package main

import (
	"fmt"
	"log-agent"
)

func main() {
	config := make(map[string]interface{})
	config["clienttype"] = "kinesis"
	config["config"] = map[string]string{"streamName":"poc-log-manager", "awsRegion": "us-east-1"}

	logagent.SetConfig(config)
	err := logagent.Dispatch(123, "test2", "test", "the action 666", "second-test", "the context 666888")
	if err != nil {
		fmt.Println(err)
	}
}
