package command

type CommandRequest struct {
	CommandType CommandType `json:"commandType"`
	TargetID    string      `json:"targetId"`
}

type CommandType int

const (
	RequestRoute CommandType = iota
	ResetRoute
)
