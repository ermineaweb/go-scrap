package services

import (
	"go-twitch/domain/entity"
	"go-twitch/domain/messaging"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
)

type SpeedometerService struct {
	Chats               entity.Chats
	twitchClient        *twitch.Client
	startStreamConsumer messaging.Consumer
	stopStreamConsumer  messaging.Consumer
	measureInterval     int
	streamChannel       chan string
}

func NewSpeedometerService(
	twitchClient *twitch.Client,
	startStreamConsumer messaging.Consumer,
	stopStreamConsumer messaging.Consumer,
	measureInterval int,
) *SpeedometerService {
	return &SpeedometerService{
		twitchClient:        twitchClient,
		Chats:               entity.NewChats(),
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
		refreshUpdate := time.NewTicker(time.Duration(s.measureInterval) * time.Second)
		defer refreshUpdate.Stop()

		for {
			select {
			case name := <-s.streamChannel:
				s.Chats = s.Chats.IncreaseMessagesOverStart(name)

			case <-refreshUpdate.C:
				s.Chats = s.Chats.IncreaseMessagesOverTime().UpdateSpeed(s.measureInterval)
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
	s.Chats = s.Chats.AddStreamer(string(message))
	s.twitchClient.OnPrivateMessage(func(m twitch.PrivateMessage) {
		s.streamChannel <- m.Channel
	})
	s.twitchClient.Join(string(message))
}

// when notifier sends a stop stream notif, the chat is removed from
// the list and the client is removed from the chat
func (s *SpeedometerService) streamStopCallback(message []byte) {
	s.Chats = s.Chats.RemoveStreamer(string(message))
	s.twitchClient.Depart(string(message))
}
