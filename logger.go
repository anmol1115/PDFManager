package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Logger struct {
	File *os.File
}

func initLogger() (*Logger, error) {
	dirPath := "."
	switch runtime.GOOS {
	case "windows":
		dirPath = filepath.Join(os.Getenv("APPDATA"), "PDFManager", "logs")
	case "linux":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		dirPath = filepath.Join(homeDir, ".local", "share", "PDFManager", "logs")
	case "darwin":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		dirPath = filepath.Join(homeDir, "Library", "Logs", "PDFManager")
	}

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return nil, err
	}

	filePath := filepath.Join(dirPath, "app.log")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &Logger{File: file}, nil
}

func (l *Logger) Write(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logString := fmt.Sprintf("[%s] -- %s\n", timestamp, message)

	l.File.WriteString(logString)
}

func (l *Logger) Close() {
	l.File.Close()
}
