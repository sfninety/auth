package main

import (
	"math/rand"
	"time"

	"github.com/sfninety/auth/internal/server"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	server.Init()
	server.Run()
}
