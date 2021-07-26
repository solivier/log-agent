package logagent

import (
	"dacast-log-agent/internal/core/services/logsservice"
	"dacast-log-agent/internal/infrastructure/storage"
	"sync"
)

var logsService *logsservice.Service
var mutex = &sync.Mutex{}
var clientConfig ClientConfig

func SetConfig(config *ClientConfig) {
	clientConfig = *config
}

func getService() (*logsservice.Service, error) {
	if logsService == nil{
		mutex.Lock()
		defer mutex.Unlock()

		logsRepository, err := storage.GetRepository(clientConfig)
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
