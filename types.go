package main

import "github.com/gorilla/websocket"

type connection struct {
	ws *websocket.Conn

	send chan []byte
}

type subscription struct {
	conn *connection
	room string
}

type message struct {
	data []byte
	room string
}

type hub struct {
	// Registered Connections
	room       map[string]map[*connection]bool
	broadcast  chan message
	register   chan subscription
	unregister chan subscription
}
