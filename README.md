# log-agent

## configuration

First chose the client type (kinesis for example). Then pass the required configuration for this client under client key.

	clientConfig := make(map[string]interface{})
	clientConfig["clienttype"] = "kinesis"
	clientConfig["config"] = map[string]string{"streamName":"xxx", "awsRegion": "xxx"}
	logagent.SetConfig(clientConfig)

Then you can call the Dispatch command to dispatch logs into the queue system.

	err = logagent.Dispatch("test", int(time.Now().UnixNano()), "1234", "1234", "xxx", "context")
	if err != nil {
		return nil, err
	}
