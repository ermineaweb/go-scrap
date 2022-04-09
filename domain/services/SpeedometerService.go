package services

import (
	"fmt"
	"go-twitch/domain/entity"
	"go-twitch/domain/messaging"
	"sort"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
)

type SpeedometerService struct {
	twitchClient        *twitch.Client
	chats               entity.Chats
	startStreamConsumer messaging.Consumer
	stopStreamConsumer  messaging.Consumer
	measureInterval     int
	streamChannel       chan string
}

func NewSpeedometerService(twitchClient *twitch.Client, startStreamConsumer messaging.Consumer, stopStreamConsumer messaging.Consumer, measureInterval int) *SpeedometerService {
	return &SpeedometerService{
		twitchClient:        twitchClient,
		chats:               entity.NewChats(),
		startStreamConsumer: startStreamConsumer,
		stopStreamConsumer:  stopStreamConsumer,
		measureInterval:     measureInterval,
		streamChannel:       make(chan string),
	}
}

func (s *SpeedometerService) Run() {
	go s.startStreamConsumer.Consume(s.streamStartCallback)
	go s.stopStreamConsumer.Consume(s.streamStopCallback)

	go func() {
		t := time.NewTicker(time.Duration(s.measureInterval) * time.Second)
		refresh := time.NewTicker(1 * time.Second)
		defer t.Stop()

		for {
			select {
			case name := <-s.streamChannel:
				s.chats.IncreaseMessagesOverStart(name)

			case <-t.C:
				s.chats.IncreaseMessagesOverTime()
				s.chats.UpdateSpeed(s.measureInterval)

			case <-refresh.C:
				fmt.Print("\033[H\033[2J")
				sort.Sort(s.chats)
				for _, chat := range s.chats {
					fmt.Printf("%s\ntotal:\t%d\nspeed:\t%d msg/min\n\n",
						chat.Streamer.Name,
						chat.NbMessageOverStart,
						chat.MsgPerMin,
					)
				}
			}
		}
	}()

	err := s.twitchClient.Connect()
	defer s.twitchClient.Disconnect()
	if err != nil {
		panic(err)
	}
}

// when notifier sends a start stream notification, the chat is appended to the list,
// an event listener is added and messages are sended on the Speedometer's chan
func (s *SpeedometerService) streamStartCallback(message []byte) {
	s.chats = s.chats.AddStreamer(string(message))
	s.twitchClient.OnPrivateMessage(func(m twitch.PrivateMessage) {
		s.streamChannel <- m.Channel
	})
	s.twitchClient.Join(string(message))
}

// when notifier sends a stop stream notif, the chat is removed from
// the list and the client is removed from the chat
func (s *SpeedometerService) streamStopCallback(message []byte) {
	s.chats = s.chats.RemoveStreamer(string(message))
	s.twitchClient.Depart(string(message))
}
