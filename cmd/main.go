package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		address: ":8080",
		db:      dbConfig{},
	}

	api := application{
		config: cfg,
	}

	h := api.mount()
	error := api.run(h)
	if error != nil {
		log.Println("server has failed to start, err:%s",error)
		os.Exit(1)
	}

}