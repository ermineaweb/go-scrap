package entity

type Chat struct {
	Streamer           *Streamer `json:"streamer"`
	NbMessageOverStart int       `json:"nb_msg_over_start"`
	NbMessageOverTime  int       `json:"-"`
	MsgPerMin          int       `json:"msg_per_min"`
	Buffer             int       `json:"-"`
}

func NewChat(streamer *Streamer) *Chat {
	return &Chat{Streamer: streamer}
}
