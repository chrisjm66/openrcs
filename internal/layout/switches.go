package layout

type SwitchId string

type Switch struct {
	Common  TrackEndpoint
	Normal  TrackEndpoint
	Reverse TrackEndpoint
}
