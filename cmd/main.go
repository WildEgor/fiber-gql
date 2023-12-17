package main

import (
	"fmt"

	server "github.com/WildEgor/fibergo-microservice-boilerplate/internal"
	log "github.com/sirupsen/logrus"
)

func main() {
	server, _ := server.NewServer()
	log.Fatal(server.Listen(fmt.Sprintf(":%v", "8888")))
}
