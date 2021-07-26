package adapters

import "log-agent/lib/core/domain"

type LogsRepository interface {
	Save(log domain.Log) error
}
