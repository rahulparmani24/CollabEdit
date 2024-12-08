package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan *Message
}

func HandleConnections(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan *Message, 256),
	}

	hub.register <- client

	go client.readMessages()
	go client.writeMessages()
}

func (c *Client) readMessages() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		var msg Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}
		c.hub.broadcast <- &msg
	}
}

func (c *Client) writeMessages() {
	defer c.conn.Close()
	for msg := range c.send {
		err := c.conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}
