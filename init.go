package main

import (
	"os"
	"path/filepath"
	"runtime"
)

func getDirPath() (string, string, error) {
	logDirPath, outputDirPath := ".", "."
	switch runtime.GOOS {
	case "windows":
		logDirPath = filepath.Join(os.Getenv("APPDATA"), "PDFManager", "logs")
		outputDirPath = filepath.Join(os.Getenv("APPDATA"), "PDFManager", "output")
	case "linux":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", "", err
		}
		logDirPath = filepath.Join(homeDir, ".local", "share", "PDFManager", "logs")
		outputDirPath = filepath.Join(homeDir, ".local", "share", "PDFManager", "outputs")
	case "darwin":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", "", err
		}
		logDirPath = filepath.Join(homeDir, "Library", "Logs", "PDFManager")
		outputDirPath = filepath.Join(homeDir, "Library", "Output", "PDFManager")
	}

	return logDirPath, outputDirPath, nil
}

func ensureOperationsDir() (string, string, error) {
	logDirPath, outputDirPath, err := getDirPath()
	if err != nil {
		return "", "", err
	}
	if err := os.MkdirAll(logDirPath, 0755); err != nil {
		return "", "", err
	}
	if err := os.MkdirAll(outputDirPath, 0755); err != nil {
		return "", "", err
	}
	return logDirPath, outputDirPath, nil
}
