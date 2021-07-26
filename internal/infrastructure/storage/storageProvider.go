package storage

import (
	"dacast-log-agent/internal/core/adapters"
	"dacast-log-agent/internal/infrastructure/config"
	"dacast-log-agent/internal/infrastructure/storage/kinesis_repository"
	"github.com/mmatagrin/ctxerror"
)

func GetRepository(config config.ClientConfig) (adapters.LogsRepository, error) {
	switch config.ClientType {
	case "kinesis":
		return kinesis_repository.NewKinesisClient(config), nil
	default:
		return nil, ctxerror.New("no repository found for " + config.ClientType)
	}
}

