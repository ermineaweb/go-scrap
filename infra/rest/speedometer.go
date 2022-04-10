package rest

import (
	"fmt"
	"go-twitch/domain/entity"
	"net/http"
)

func SpeedometerHandler(chats *entity.Chats) func(http.ResponseWriter, *http.Request) {
	fmt.Println("create the handler")
	return func(w http.ResponseWriter, r *http.Request) {
		chats.EncodeToJSON(w)
	}
}
