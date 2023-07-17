package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	app := gin.New()
	app.GET("/echo", func(c *gin.Context) {
		conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)

		for {
			// read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	// http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
	// 	conn, _ := upgrader.Upgrade(w, r, nil)

	// 	for {
	// 		// read message from browser
	// 		msgType, msg, err := conn.ReadMessage()
	// 		if err != nil {
	// 			return
	// 		}

	// 		// Print the message to the console
	// 		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

	// 		// Write message back to browser
	// 		if err = conn.WriteMessage(msgType, msg); err != nil {
	// 			return
	// 		}
	// 	}
	// })

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "websockets.html")
	// })

	app.GET("/", func(c *gin.Context) {
		c.File("websockets.html")
	})

	// http.ListenAndServe(":8080", nil)
	app.Run(":8080")
}
