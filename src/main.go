package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/thiago18l/restful-gin-api/src/server"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	server.Start()
}
