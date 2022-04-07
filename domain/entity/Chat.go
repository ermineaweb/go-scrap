package entity

type Chat struct {
	Streamer           *Streamer
	NbMessageOverStart int
	NbMessageOverTime  int
	Speed              float32
	Buffer             int
}

func NewChat(streamerUsername string) *Chat {
	return &Chat{Streamer: NewStreamer(streamerUsername)}
}
