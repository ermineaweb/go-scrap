package entity

type Chats []*Chat

// return a new list of chats with optionals streamers
func NewChats() Chats {
	return Chats{}
}

func (c Chats) AddStreamer(name string) Chats {
	streamer := NewStreamer(name)
	if c.contains(streamer) {
		return c
	}
	return append(c, NewChat(NewStreamer(name)))
}

func (c Chats) RemoveStreamer(name string) Chats {
	for k, s := range c {
		if s.Streamer.Name == name {
			return append(c[:k], c[k+1:]...)
		}
	}
	return c
}

func (c Chats) contains(streamer *Streamer) bool {
	for _, chat := range c {
		if chat.Streamer.Name == streamer.Name {
			return true
		}
	}
	return false
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

// update the speed of messages, in messages/min
func (c Chats) UpdateSpeed(duration int) {
	for _, chat := range c {
		chat.MsgPerMin = chat.NbMessageOverTime * 60 / duration
	}
}

// implement interface sortable
func (c Chats) Len() int           { return len(c) }
func (c Chats) Less(i, j int) bool { return c[i].MsgPerMin > c[j].MsgPerMin }
func (c Chats) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
