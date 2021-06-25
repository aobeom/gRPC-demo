package configs

import (
	"log"
)

const (
	BindAddress = "localhost:50051"
)

func StartInfo(isServer bool) {
	if isServer {
		log.Printf("Server is Listening: %s\n", BindAddress)
	} else {
		log.Printf("Connected Server: %s\n", BindAddress)
	}
}
