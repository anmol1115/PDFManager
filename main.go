package main

import "log"

func main() {
	logger, err := initLogger()
	if err != nil {
		log.Fatalf("Unable to load logger, %v", err)
	}
	defer logger.Close()
}
