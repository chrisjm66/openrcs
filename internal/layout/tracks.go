package layout

type Track struct {
	Length int
}

type TrackId string

type TrackPosition struct {
	TrackId TrackId
	Offset  int
}

type TrackEndpoint struct {
	TrackId  TrackId
	Endpoint TrackEnd
}

type TrackEnd int

const (
	Start TrackEnd = iota
	End
)
