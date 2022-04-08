package main

import (
	"fmt"
	"go-twitch/domain/entity"
	"sort"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
)

const MEASURE_INTERVAL = 5

func main() {
	// OR client := twitch.NewClient("yourtwitchusername", "oauth:123123123")
	client := twitch.NewAnonymousClient()

	chats := entity.NewChats().
		AddStreamer("zerator").
		AddStreamer("otplol_").
		AddStreamer("kameto").
		AddStreamer("hasanabi")

	channel := make(chan string)

	for _, chat := range chats {
		go func(chat *entity.Chat) {
			fmt.Println("create channel " + chat.Streamer.Name)
			client.OnPrivateMessage(func(m twitch.PrivateMessage) {
				channel <- m.Channel
			})
			client.Join(chat.Streamer.Name)
		}(chat)
	}

	go func() {
		t := time.NewTicker(MEASURE_INTERVAL * time.Second)
		defer t.Stop()

		for {
			select {
			case name := <-channel:
				fmt.Print("\033[H\033[2J")
				chats.IncreaseMessagesOverStart(name)
				sort.Sort(chats)
				for _, chat := range chats {
					fmt.Printf("%s\ntotal: %d\tspeed: %.1f msg/s\n\n",
						chat.Streamer.Name,
						chat.NbMessageOverStart,
						chat.Speed,
					)
				}

			case <-t.C:
				chats.IncreaseMessagesOverTime()
				chats.UpdateSpeed(MEASURE_INTERVAL)
			}
		}
	}()

	err := client.Connect()
	defer client.Disconnect()
	if err != nil {
		panic(err)
	}
}
