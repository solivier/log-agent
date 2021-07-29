package domain

type Log struct {
	Id         string `json:"id"`
	CreatedAt  int    `json:"created-at"`
	AccountId  string `json:"account-id"`
	UserId     string `json:"user-id"`
	ActionType string `json:"action-type"`
	ServiceId  string `json:"service-id"`
	Context    string `json:"context"`
}

func NewLog(id string, createdAt int, accountId, userId, actionType, serviceId, context string) Log {
	return Log{
		Id:         id,
		CreatedAt:  createdAt,
		AccountId:  accountId,
		UserId:     userId,
		ActionType: actionType,
		ServiceId:  serviceId,
		Context:    context,
	}
}
