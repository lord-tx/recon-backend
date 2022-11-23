package main

import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var h = hub{
	broadcast:  make(chan message),
	room:       make(map[string]map[*connection]bool),
	register:   make(chan subscription),
	unregister: make(chan subscription),
}
