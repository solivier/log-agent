package config

type KinesisClientConfig struct {
	StreamName string
}

func NewKinesisClientConfig(streamName string) *KinesisClientConfig {
	return &KinesisClientConfig{StreamName: streamName}
}
