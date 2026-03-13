package main

import "log"

const WINDOW_HEIGHT float32 = 500
const WINDOW_WIDTH float32 = 700

func main() {
	logger, err := initLogger()
	if err != nil {
		log.Fatalf("Unable to load logger, %v", err)
	}
	defer logger.Close()

	logger.Write("Running main UI loop")
	runUI()
	logger.Write("Exiting main UI loop")
}
