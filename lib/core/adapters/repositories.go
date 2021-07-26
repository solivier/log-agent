package adapters

import "dacast-log-agent/lib/core/domain"

type LogsRepository interface {
	Save(log domain.Log) error
}
