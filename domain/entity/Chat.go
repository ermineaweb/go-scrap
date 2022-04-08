package entity

type Chat struct {
	Streamer           *Streamer
	NbMessageOverStart int
	NbMessageOverTime  int
	Buffer             int
	Speed              float32
}

func NewChat(streamer *Streamer) *Chat {
	return &Chat{Streamer: streamer}
}
