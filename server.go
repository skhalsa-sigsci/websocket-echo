package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, homeTemplate)
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

var homeTemplate = `
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>WebSocket server
<p>
</body>
</html>
`

func main() {
	http.HandleFunc("/ws", ws)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
