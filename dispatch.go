package logagent

import (
	"dacast-log-agent/internal/core/services/logsservice"
	"dacast-log-agent/internal/infrastructure/config"
	"dacast-log-agent/internal/infrastructure/storage/logsrepository"
	"sync"
)

var logsService *logsservice.Service
var mutex = &sync.Mutex{}
var kinesisConfig config.KinesisClientConfig

func SetConfig(config *config.KinesisClientConfig) {
	kinesisConfig = *config
}

func getService() *logsservice.Service {
	if logsService == nil{
		mutex.Lock()
		defer mutex.Unlock()

		logsRepository := logsrepository.NewKinesisClient(kinesisConfig)
		logsService = logsservice.New(logsRepository)
	}

	return logsService
}

func Dispatch(id string, createdAt int, accountId, userId, actionType, context string) error {
	logsService = getService()

	err := logsService.Dispatch(id, createdAt, accountId, userId, actionType, context)
	if err != nil {
		return err
	}

	return nil
}
