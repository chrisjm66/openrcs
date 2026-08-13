package command

type CommandRequest struct {
	CommandType string `json:"commandType"`
	TargetID    string `json:"targetId"`
}
