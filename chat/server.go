package main

import (
	"log"
	"net/http"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

var listman map[string]string
var listman2 string

func getlist() string {
	listman3 := ""
	for key, _ := range listman {
		listman3 += key + "<br>"
	}
	return listman3
}

func main() {
	listman = make(map[string]string)
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	type Message struct {
		// 클라이언트 부분과 json 이름이 동일한지 잘 확인 하셔야 합니다.
		Clist   string `json:"clist"`
		Message string `json:"message"`
	}

	//handle connected
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("New client connected")
		//join them to room
		c.Join("chat")
		listman[c.Id()] = ""

		listman2 = getlist()

		str := Message{Clist: listman2, Message: "--- <font color='red'>" + c.Id() +
			"</font> 님이 들어오셨습니다.<br>"}
		c.BroadcastTo("chat", "message", str)
	})

	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		//caller is not necessary, client will be removed from rooms
		//automatically on disconnect
		//but you can remove client from room whenever you need to
		c.Leave("chat")
		delete(listman, c.Id())

		listman2 = getlist()

		str := Message{Clist: listman2, Message: "--- <font color='blue'>" + c.Id() +
			"</font> 님이 나가셨습니다.<br>"}
		c.BroadcastTo("chat", "message", str)
		log.Println("Disconnected")
	})

	//error catching handler
	server.On(gosocketio.OnError, func(c *gosocketio.Channel) {
		log.Println("Error occurs")
	})

	//handle custom event
	server.On("send", func(c *gosocketio.Channel, msg Message) string {

		msg.Clist = listman2
		msg.Message = c.Id() + " :: <font color='green'>" + msg.Message + "</font>"

		// 전체가 아닌 접속한 소켓 아이디 한 사람에게만 메세지
		// channel, _ := server.GetChannel(c.Id())
		// channel.Emit("message", msg)

		// 방과 상관없이 접속한 모든 사람한테 메세지
		// server.BroadcastToAll("message", msg)

		// 특정 방에 있는 사람들에게 메세지
		c.BroadcastTo("chat", "message", msg)
		return "OK"
	})

	//setup http server
	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)
	serveMux.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at http://localhost:8080 ...")
	log.Panic(http.ListenAndServe(":8080", serveMux))
}
