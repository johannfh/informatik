package api

type Hub struct {
	clients map[*Client]bool

	broadcast chan []byte

	register chan *Client

	unregister chan *Client
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			delete(h.clients, client)
		case msg := <-h.broadcast:
			for client := range h.clients {
				select {
				// Try sending the message to the connection
				case client.send <- msg:
				// Close the connection if channel is full
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
