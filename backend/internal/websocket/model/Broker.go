package model

type Broker struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func (b *Broker) run() {
	for {
		select {
		case client := <-b.Register:
			b.Clients[client] = true
		case client := <-b.Unregister:
			if _, ok := b.Clients[client]; ok {
				delete(b.Clients, client)
				close(client.send)
			}
		case message := <-b.Broadcast:
			// "Fan-out": Envía el mensaje a CADA cliente
			for client := range b.Clients {
				client.send <- message // Envía al channel del cliente, NO directamente a la conexión
			}
		}
	}
}
