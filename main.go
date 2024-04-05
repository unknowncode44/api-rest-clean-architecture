package main

import (
	"github.com/unknowncode44/api-rest-clean-architecture/config"
	"github.com/unknowncode44/api-rest-clean-architecture/database"
	"github.com/unknowncode44/api-rest-clean-architecture/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewEchoServer(conf, db).Start()
}
