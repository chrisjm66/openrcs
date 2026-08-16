package command

type CommandRequest struct {
	CommandType CommandType `json:"commandType"`
	TargetId    string      `json:"targetId"`
}

type CommandType int

const (
	RequestRoute CommandType = iota
	ResetRoute
)
