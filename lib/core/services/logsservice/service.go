package logsservice

import (
	"dacast-log-agent/lib/core/adapters"
	"dacast-log-agent/lib/core/domain"
	"github.com/mmatagrin/ctxerror"
)

func New(logsRepository adapters.LogsRepository) *Service {
	return &Service{
		logsRepository: logsRepository,
	}
}

type Service struct {
	logsRepository adapters.LogsRepository
}

func (srv *Service) Dispatch(id string, createdAt int, accountId, userId, actionType, context string) error {
	ctxErr := ctxerror.SetContext(map[string]interface{}{
		"id":     id,
		"created-at": createdAt,
		"account-id": accountId,
		"user-id": userId,
		"action-type": actionType,
		"context": context,
	})

	log := domain.NewLog(id, createdAt, accountId, userId, actionType, context)

	if err := srv.logsRepository.Save(log); err != nil {
		return ctxErr.Wrap(err, "Create log into repositories has failed")
	}

	return nil
}
