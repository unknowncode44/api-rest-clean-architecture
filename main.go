package main

import (
	"fmt"

	"github.com/unknowncode44/api-rest-clean-architecture/config"
	"github.com/unknowncode44/api-rest-clean-architecture/database"
	utils "github.com/unknowncode44/api-rest-clean-architecture/internal/bcrautils"
	"github.com/unknowncode44/api-rest-clean-architecture/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	exchange := utils.GetExchange(conf)
	fmt.Println(exchange)
	server.NewEchoServer(conf, db).Start()
}
