package main
import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r *http.Request) bool {return true},
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000",nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		messageType, p, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(p))
		if err := socket.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
		}
	}
}
