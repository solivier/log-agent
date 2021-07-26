
package main

import (
	logagent "dacast-log-agent"
	"dacast-log-agent/infrastructure/config"
)

func main() {
	logagent.SetConfig(config.NewKinesisClientConfig("poc-log-manager"))
	err := logagent.Dispatch("id2", 123, "12345", "12345", "the action 2", "the context 2")
	if err != nil {
		return
	}
}
