package layout

type SwitchId string

type Switch struct {
	Common  EdgeEnd `json:"common"`
	Normal  EdgeEnd `json:"normal"`
	Reverse EdgeEnd `json:"reverse"`
}
