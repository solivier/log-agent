package kinesis_repository

import (
	"dacast-log-agent/config"
	"dacast-log-agent/lib/core/adapters"
	"github.com/mmatagrin/ctxerror"
)

func GetRepository(config config.ClientConfig) (adapters.LogsRepository, error) {
	switch config.ClientType {
	case "kinesis":
		return NewKinesisClient(config), nil
	default:
		return nil, ctxerror.New("no repositories found for " + config.ClientType)
	}
}

