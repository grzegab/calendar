package ws

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub       *Hub
	Conn      *websocket.Conn
	Send      chan []byte
	userID    int
	OnMessage func([]byte)
}

func (c *Client) SetUserID(id int) {
	c.userID = id
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Printf("received msg from user %d: %s\n", c.userID, message)

		if c.OnMessage != nil {
			c.OnMessage(message)
		} else {
			// fallback to simple broadcast
			c.Hub.broadcast <- message
		}
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(30 * time.Second)

	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {

		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, nil)
				return
			}

			c.Conn.WriteMessage(websocket.TextMessage, message)

		case <-ticker.C:
			c.Conn.WriteMessage(websocket.PingMessage, nil)
		}
	}
}
