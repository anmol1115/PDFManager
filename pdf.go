package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func mergePDF(selectedFiles []string, outputFilePath string) error {
	fileName := fmt.Sprintf("%d.pdf", time.Now().Unix())
	filePath := filepath.Join(outputFilePath, fileName)
	err := api.MergeCreateFile(selectedFiles, filePath, false, nil)
	return err
}
