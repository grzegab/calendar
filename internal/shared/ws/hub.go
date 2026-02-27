package ws

import "fmt"

type Hub struct {
	clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	broadcast  chan []byte
	stop       chan struct{}
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		broadcast:  make(chan []byte),
		stop:       make(chan struct{}),
	}
}

func (h *Hub) Stop() {
	close(h.stop)
}

func (h *Hub) Broadcast(msg []byte) {
	h.broadcast <- msg
}

func (h *Hub) Run() {
	for {
		select {
		case <-h.stop:
			for client := range h.clients {
				delete(h.clients, client)
				close(client.Send)
			}
			return

		case client := <-h.Register:
			fmt.Println("registering client")
			h.clients[client] = true

		case client := <-h.Unregister:
			delete(h.clients, client)
			close(client.Send)

		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
		}
	}
}
