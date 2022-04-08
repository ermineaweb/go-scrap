package entity

type Chats []*Chat

// return a new list of chats with optionals streamers
func NewChats(streamers ...*Streamer) Chats {
	var chats Chats
	for _, streamer := range streamers {
		chats = append(chats, NewChat(streamer))
	}
	return chats
}

func (c Chats) AddStreamer(name string) Chats {
	return append(c, NewChat(NewStreamer(name)))
}

// update the total message since the start of the scrap
func (c Chats) IncreaseMessagesOverStart(username string) {
	for _, chat := range c {
		if chat.Streamer.Name == username {
			chat.NbMessageOverStart = chat.NbMessageOverStart + 1
			chat.Buffer = chat.Buffer + 1
		}
	}
}

// update the total message since the last delay
func (c Chats) IncreaseMessagesOverTime() {
	for _, chat := range c {
		chat.NbMessageOverTime = chat.Buffer
		chat.Buffer = 0
	}
}

// update the speed of messages, in messages/second
func (c Chats) UpdateSpeed(duration int) {
	for _, chat := range c {
		chat.Speed = float32(chat.NbMessageOverTime) / float32(duration)
	}
}

// implement interface sortable
func (c Chats) Len() int           { return len(c) }
func (c Chats) Less(i, j int) bool { return c[i].Speed > c[j].Speed }
func (c Chats) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
