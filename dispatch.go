package logagent

import (
	"log-agent/config"
	"log-agent/lib/core/services/logsservice"
	"log-agent/lib/infrastructure/storage/repositories"
	"sync"
)

var logsService *logsservice.Service
var mutex = &sync.Mutex{}
var clientConfig config.ClientConfig

func SetConfig(config config.ClientConfig) {
	clientConfig = config
}

func getService() (*logsservice.Service, error) {
	if logsService == nil{
		mutex.Lock()
		defer mutex.Unlock()

		logsRepository, err := kinesis_repository.GetRepository(clientConfig)
		if nil != err {
			return nil, err
		}
		logsService = logsservice.New(logsRepository)
	}

	return logsService, nil
}

func Dispatch(id string, createdAt int, accountId, userId, actionType, context string) error {
	logsService, err := getService()
	if nil != err {
		return err
	}

	err = logsService.Dispatch(id, createdAt, accountId, userId, actionType, context)
	if err != nil {
		return err
	}

	return nil
}
