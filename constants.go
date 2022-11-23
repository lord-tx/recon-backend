package main

import "time"

const (
	writeWait = 10 * time.Second

	readWait = 10 * time.Second

	pingPeriod = (readWait * 9) / 10

	maxMessageSize = 512
)
