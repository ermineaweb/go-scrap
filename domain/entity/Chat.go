package entity

type Chat struct {
	Streamer           *Streamer
	NbMessageOverStart int
	NbMessageOverTime  int
	Buffer             int
	MsgPerMin          int
}

func NewChat(streamer *Streamer) *Chat {
	return &Chat{Streamer: streamer}
}
