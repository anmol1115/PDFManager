package main

import "log"

const WINDOW_HEIGHT float32 = 500
const WINDOW_WIDTH float32 = 700

func main() {
	logDir, outputDir, err := ensureOperationsDir()
	if err != nil {
		log.Fatalf("Unable to initialize operations directory: %v", err)
	}

	logger, err := initLogger(logDir)
	if err != nil {
		log.Fatalf("Unable to load logger, %v", err)
	}
	defer logger.Close()

	logger.Write("Running main UI loop")
	runUI(outputDir)
	logger.Write("Exiting main UI loop")
}
