package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func serveWs(w http.ResponseWriter, r *http.Request, roomId string) {
	fmt.Println(roomId)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{send: make(chan []byte), ws: ws}
	s := subscription{c, roomId}
	h.register <- s
	go s.readPump()
	go s.writePump()
}

func main() {
	// Initialize Hub Service
	go h.run()

	router := gin.New()

	router.GET("/ws/:room_id", func(ctx *gin.Context) {
		roomId := ctx.Param("room_id")
		// Initialize WebSocket Service
		serveWs(ctx.Writer, ctx.Request, roomId)
	})

	router.Run("localhost:8080")
}
