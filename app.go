package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/distopico/database"
	"github.com/distopico/router"
)

func main() {
	db := database.ConnectDB()
	defer db.CloseDB()

	engine := router.Create(&db)

	httpHandler := engine

	addr := fmt.Sprintf("%s:%d", "127.0.0.1", 9000)
	fmt.Println("Started Listening for plain HTTP connection on " + addr)
	server := &http.Server{Addr: addr, Handler: httpHandler}

	log.Fatal(server.ListenAndServe())

}
