package kinesis_repository

import (
	"github.com/mmatagrin/ctxerror"
	"log-agent/config"
	"log-agent/lib/core/adapters"
)

func GetRepository(config config.ClientConfig) (adapters.LogsRepository, error) {
	switch config.ClientType {
	case "kinesis":
		return NewKinesisClient(config), nil
	default:
		return nil, ctxerror.New("no repositories found for " + config.ClientType)
	}
}

