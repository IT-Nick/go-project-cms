package main

import (
	"fmt"
	"github.com/IT-Nick/go-project-cms/cmd/rcms/config"
)

func main() {
	// подгружаем файлы конфигурации
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	fmt.Println(config.Config.ConfigVar)
}
