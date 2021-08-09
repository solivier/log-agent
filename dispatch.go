package logagent

import (
	"github.com/mitchellh/mapstructure"
	uuid "github.com/satori/go.uuid"
	"github.com/solivier/log-agent/config"
	"github.com/solivier/log-agent/lib/core/services/logsservice"
	"github.com/solivier/log-agent/lib/infrastructure/storage/repositories"
	kinesis_repository "github.com/solivier/log-agent/lib/infrastructure/storage/repositories"
	"sync"
)

var logsService *logsservice.Service
var mutex = &sync.Mutex{}
var clientConfig config.ClientConfig

func SetConfig(logAgentConfig map[string]interface{}) error {
	result := config.ClientConfig{}
	err := mapstructure.Decode(logAgentConfig, &result)
	if err != nil {
		return err
	}

	clientConfig = result

	return nil
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

func Dispatch(createdAt int, accountId, userId, actionType, serviceId, context string) error {
	logsService, err := getService()
	if nil != err {
		return err
	}
	id := uuid.NewV4()

	err = logsService.Dispatch(id.String(), createdAt, accountId, userId, actionType, serviceId, context)
	if err != nil {
		return err
	}

	return nil
}
