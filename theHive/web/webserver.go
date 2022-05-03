package webServer

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"syscall"

	"github.com/gorilla/websocket"
	serviceContract "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var webSubPool *[]chan serviceContract.GridPositions

func SetupWebServer(subs *[]chan serviceContract.GridPositions) {
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

	log.Printf("Received subscribe request from ID: %s", r.Context())

	// Handle subscribe request
	fin := make(chan bool)
	// Save the subscriber stream according to the given client ID
	channel := make(chan serviceContract.GridPositions, 2)
	*webSubPool = append(*webSubPool, channel)
	ctx := r.Context()

	for {
		select {
		case update := <-channel:
			//Convert to json
			b, err := json.Marshal(update)
			if err != nil {
				log.Fatal("encode error:", err)
			}

			//TODO: Handle client disconnect issue
			if err = conn.WriteMessage(websocket.TextMessage, b); err != nil {
				if errors.Is(err, syscall.EPIPE) {
					// just ignore.
					return
				} else {
					// Here is not "broken pipe" error.
					log.Fatal(err)
				}
			}
			// if err == io.EOF {
			// 	log.Printf(("Banaaaan"))
			// }
			// if err == syscall.EPIPE {
			// 	log.Printf(("Client forcefully disconnected"))
			// } else if err != nil {
			// 	log.Fatal(err)
			// }
		case <-fin:
			log.Printf("Closing stream for client ID: %s", r.Context())
			return
		case <-ctx.Done():
			log.Printf("Client ID %s has disconnected", r.Context())
			return
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../web/static/index.html")
}
