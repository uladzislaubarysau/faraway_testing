package main

import (
	"log"
	"time"

	"github.com/uladzislaubarysau/faraway_testing/internal/client"
)

func main() {
	conf, err := loadConfig()
	if err != nil {
		log.Println("can't load config:", err)
	}

	for {
		if msg, err := client.New(conf.ServerHost + ":" + conf.ServerPort).GetWow(); err != nil {
			log.Println("can't get word of wisdom:", err)
			time.Sleep(conf.TimeForRetry)
		} else {
			log.Println("word of wisdom:", msg)
			return
		}
	}
}
