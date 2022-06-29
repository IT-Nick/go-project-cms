package main

import (
	"fmt"
	"github.com/IT-Nick/go-project-cms/cmd/rcms/config"
	"github.com/IT-Nick/go-project-cms/cmd/rcms/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	// подгружаем файлы конфигурации
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	fmt.Println(config.Config.ConfigVar)

	router := httprouter.New()

	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))
}
