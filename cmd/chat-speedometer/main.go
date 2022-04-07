package main

import (
	"fmt"
	"go-twitch/domain/entity"
	"sort"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
)

const MEASURE_INTERVAL = 5

type Chats []*entity.Chat

func main() {
	// client := twitch.NewClient("yourtwitchusername", "oauth:123123123")
	// or client := twitch.NewAnonymousClient() for an anonymous user (no write capabilities)
	client := twitch.NewAnonymousClient()

	chats := Chats{
		entity.NewChat("asmongold"),
		entity.NewChat("otplol_"),
		entity.NewChat("aminematue"),
		entity.NewChat("antoinedaniel"),
		entity.NewChat("andymilonakis"),
		entity.NewChat("shadowkekw"),
		entity.NewChat("gaules"),
	}

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
			case c := <-channel:
				fmt.Print("\033[H\033[2J")
				chats.updateTotalOverStart(c)
				sort.Sort(chats)
				for _, c := range chats {
					fmt.Printf("%s\ntotal: %d\tspeed: %.1f msg/s\n\n", c.Streamer.Name, c.NbMessageOverStart, c.Speed)
				}

			case <-t.C:
				chats.updateTotalOverTime()
				chats.updateSpeed(MEASURE_INTERVAL)
			}
		}
	}()

	err := client.Connect()
	defer client.Disconnect()
	if err != nil {
		panic(err)
	}
}

func (chats Chats) updateTotalOverStart(username string) {
	for _, c := range chats {
		if c.Streamer.Name == username {
			c.NbMessageOverStart = c.NbMessageOverStart + 1
			c.Buffer = c.Buffer + 1
		}
	}
}

func (chats Chats) updateTotalOverTime() {
	for _, c := range chats {
		c.NbMessageOverTime = c.Buffer
		c.Buffer = 0
	}
}

func (chats Chats) updateSpeed(duration int) {
	for _, c := range chats {
		c.Speed = float32(c.NbMessageOverTime) / float32(duration)
	}
}

// implement interface sortable
func (chats Chats) Len() int           { return len(chats) }
func (chats Chats) Less(i, j int) bool { return chats[i].Speed > chats[j].Speed }
func (chats Chats) Swap(i, j int)      { chats[i], chats[j] = chats[j], chats[i] }
