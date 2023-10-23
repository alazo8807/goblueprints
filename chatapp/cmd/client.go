package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	// send is a channel on which messages are sent from the room.
	send chan []byte
	// room is the room this client is chatting in.
	room *room
}

// read will continually read from the websocket(web-client for example), if a msg is reaceived,
// it will forward it to the room channel.
func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}

		// Forward the message received to the channel.
		// TODO: Handle sending to another chat client.
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
