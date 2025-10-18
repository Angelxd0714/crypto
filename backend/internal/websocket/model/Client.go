package model

import (
	"gin-quickstart/internal/database"

	"github.com/google/uuid"
	websocket "github.com/gorilla/websocket"
)

type Client struct {
	broker *Broker
	conn   *websocket.Conn
	send   chan []byte
	UserID uuid.UUID
	Alerts []database.Alert
}

func (c *Client) readPump() {
	defer c.conn.Close() // Cierra la conexión si esta goroutine termina
	// ... lógica para leer pings ...
}

// writePump: Goroutine 2 (Escribe al cliente)
// Escucha en su propio channel 'send' y escribe en el WebSocket
func (c *Client) writePump() {
	defer c.conn.Close()
	for message := range c.send {
		c.conn.WriteMessage(websocket.TextMessage, message)
	}
}
