package port

type LogsService interface {
	Dispatch(id string, createdAt int, accountId, userId, actionType, context string) error
}
