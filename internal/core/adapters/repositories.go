package adapters

import "dacast-log-agent/internal/core/domain"

type LogsRepository interface {
	Save(log domain.Log) error
}
