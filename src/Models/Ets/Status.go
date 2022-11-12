package Ets

type Status struct {
	Value           string `json:"value"`
	ChangedDateTime string `json:"changedDateTime"`
}

func NewStatus(value string, changedDateTime string) *Status {
	return &Status{
		Value:           value,
		ChangedDateTime: changedDateTime,
	}
}
