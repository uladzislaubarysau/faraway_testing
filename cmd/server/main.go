package main

import (
	"log"

	"github.com/uladzislaubarysau/faraway_testing/internal/server"
)

func main() {
	conf, err := loadConfig()
	if err != nil {
		log.Println("can't load config:", err)
	}

	s := server.New()
	s.Start(conf.Host + ":" + conf.Port)
}
