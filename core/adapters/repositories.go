package adapters

import "dacast-log-agent/core/domain"

type LogsRepository interface {
	Save(log domain.Log) error
}
