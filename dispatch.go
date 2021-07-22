package logagent

import (
	"dacast-log-agent/core/services/logsservice"
	"dacast-log-agent/infrastructure/storage/logsrepository"
)

func Dispatch(id string, createdAt int, accountId, userId, actionType, context string) error {
	logsRepository := logsrepository.NewKinesisClient()
	logsService := logsservice.New(logsRepository)
	err := logsService.Dispatch(id, createdAt, accountId, userId, actionType, context)
	if err != nil {
		return err
	}

	return nil
}
