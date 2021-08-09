package adapters

import "github.com/solivier/log-agent/lib/core/domain"

type LogsRepository interface {
	Save(log domain.Log) error
}
