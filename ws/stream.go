package ws 
const (
	publish     = "publish"
	subscribe   = "subscribe"
	unsubscribe = "unsubscribe"
)

type Server struct {
	topic string
	subscription Subscription
}

func (ser *Server) Subscribe(topic string,conn *websocket.Conn) {
    append(ser.subscription[topic],)
}

// Unsubscribe the client to the particular topic   
func (ser *Server) Unsubscribe(topic string) {
	delete(ser[topic],conn) 
}

func (s *Server) Send(conn *websocket.Conn, message string) {
	conn.WriteMessage(websocket.TextMessage, []byte(message))
}


func (s *Server) SendWithWait(conn *websocket.Conn, message string, wg *sync.WaitGroup) {
	conn.WriteMessage(websocket.TextMessage, []byte(message))
	wg.Done()
}


func (s *Server) ProcessMessage(conn *websocket.Conn,clientID string,msg byte[]) {
    m := Message{}
	if err := json.Unmarshal(msg, &m); err != nil {
		s.Send(conn, errInvalidMessage)
	}
    action := strings.TrimSpace(strings.ToLower(m.action))
	switch action {
	case publish:
		s.Publish(m.Topic, []byte(m.Message))
	case subscribe:
		s.Subscribe(conn, clientID, m.Topic)
	case unsubscribe:
		s.Unsubscribe(clientID, m.Topic)
	default:
		s.Send(conn, errActionUnrecognizable)
	}
}