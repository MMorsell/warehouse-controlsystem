package webServer

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"syscall"

	"github.com/gorilla/websocket"
	botClientService "gits-15.sys.kth.se/Gophers/walle/theHive/api"
	serviceContract "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var webSubPool *[]botClientService.WebSub

func SetupWebServer(subs *[]botClientService.WebSub) {
	log.Printf("Starting Web server!")

	webSubPool = subs
	http.HandleFunc("/websocketConnection", wsConnection)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}

func wsConnection(w http.ResponseWriter, r *http.Request) {
	//Upgrade connection from http to webSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	log.Printf("Received subscribe request from browser gui")

	// Save the subscriber stream according to the given client ID
	channel := make(chan serviceContract.GridPositions)
	clientClosedConnection := false
	sub := botClientService.WebSub{Channel: &channel, ClosedConnection: &clientClosedConnection}
	*webSubPool = append(*webSubPool, sub)
	for {
		select {

		case update := <-*sub.Channel:
			//Convert to json
			b, err := json.Marshal(update)
			if err != nil {
				log.Fatal("encode error:", err)
			}

			if err = conn.WriteMessage(websocket.TextMessage, b); err != nil {
				if errors.Is(err, syscall.EPIPE) {
					// just ignore.
					clientClosedConnection = true
					log.Printf("Client Disconnected, removing connection")
					return
				} else {
					// Not "broken pipe" error, log message
					log.Fatal(err)
					return
				}
			}

		case <-r.Context().Done():
			log.Printf("Client has disconnected")
			return
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../web/static/index.html")
}
