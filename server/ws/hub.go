package ws

type Room struct {
	ID      string             `json:"id"`
	name    string             `json:"name"`
	clients map[string]*Client `json:"clients"`
}
type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				if _, ok := r.clients[cl.ID]; !ok {
					r.clients[cl.ID] = cl
				}
			}
		case cl := <-h.Unregister:

			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].clients[cl.ID]; ok {

					if len(h.Rooms[cl.RoomID].clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}
					delete(h.Rooms[cl.RoomID].clients, cl.ID)
					close(cl.Message)
				}
			}
		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range h.Rooms[m.RoomID].clients {
					cl.Message <- m
				}
			}
		}
	}
}
