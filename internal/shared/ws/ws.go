package ws

import (
	//"github/grzegab/calendar/internal/to_delete"
	"math/rand"
	"net/http"

	"github.com/gorilla/websocket"
)

type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // @TODO: change this when going prod!
	},
}

// WsUpgrade upgrade connection to WS
func WsUpgrade(h *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade connection to websocket", http.StatusBadRequest)
		return
	}

	response := WsJsonResponse{
		MessageType: "welcome",
	}

	err = conn.WriteJSON(response)
	if err != nil {
		http.Error(w, "Could not write to websocket", http.StatusInternalServerError)
		return
	}

	client := &Client{
		Hub:  h,
		Conn: conn,
		Send: make(chan []byte, 256),
		OnMessage: func(msg []byte) {
			Handle(h, msg)
		},
	}
	client.SetUserID(rand.Int())

	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
