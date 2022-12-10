package main

import (
	"github.com/cockscomb/tinyurl/web"
	"log"
)

func main() {
	server := web.NewServer()
	log.Fatalln(server.Start(":8080"))
}
