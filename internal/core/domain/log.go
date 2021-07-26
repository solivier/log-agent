package domain

type Log struct {
	Id         string `json:"id"`
	CreatedAt  int    `json:"created-at"`
	AccountId  string `json:"account-id"`
	UserId     string `json:"user-id"`
	ActionType string `json:"action-type"`
	Context    string `json:"context"`
}

func NewLog(id string, createdAt int, accountId, userId, actionType, context string) Log {
	return Log{
		Id:         id,
		CreatedAt:  createdAt,
		AccountId:  accountId,
		UserId:     userId,
		ActionType: actionType,
		Context:    context,
	}
}
