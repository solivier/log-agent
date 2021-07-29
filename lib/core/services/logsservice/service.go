package logsservice

import (
	"github.com/mmatagrin/ctxerror"
	"log-agent/lib/core/adapters"
	"log-agent/lib/core/domain"
)

func New(logsRepository adapters.LogsRepository) *Service {
	return &Service{
		logsRepository: logsRepository,
	}
}

type Service struct {
	logsRepository adapters.LogsRepository
}

func (srv *Service) Dispatch(id string, createdAt int, accountId, userId, actionType, serviceId, context string) error {
	ctxErr := ctxerror.SetContext(map[string]interface{}{
		"id":     id,
		"created-at": createdAt,
		"account-id": accountId,
		"user-id": userId,
		"action-type": actionType,
		"service-id": serviceId,
		"context": context,
	})

	log := domain.NewLog(id, createdAt, accountId, userId, actionType, serviceId, context)

	if err := srv.logsRepository.Save(log); err != nil {
		return ctxErr.Wrap(err, "Create log into repositories has failed")
	}

	return nil
}
