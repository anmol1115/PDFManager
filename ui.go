package main

import (
	"errors"
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var SelectedFiles []string
var OutputFilePath string

func runUI(outputDir string, l *Logger) {
	OutputFilePath = outputDir
	a := app.New()

	w := newWindow(a)
	w.SetContent(homeView(a, w, l))

	w.ShowAndRun()
}

func newWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("PDFManager")
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	w.SetFixedSize(true)

	return w
}

func getOutputBrowseButton(outputEntry *widget.Entry, w fyne.Window) *widget.Button {
	return widget.NewButton("Browse", func() {
		dialog.ShowFolderOpen(func(reader fyne.ListableURI, err error) {
			if err != nil || reader == nil {
				return
			}
			OutputFilePath = reader.Path()
			outputEntry.SetPlaceHolder(OutputFilePath)
		}, w)
	})
}

func getBrowseFileButton(fileList *widget.List, w fyne.Window) *widget.Button {
	return widget.NewButton("Browse", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				return
			}
			SelectedFiles = append(SelectedFiles, reader.URI().Path())
			fileList.Refresh()
		}, w)

		fd.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
		fd.Show()
	})
}

func getSubmitButton(a fyne.App, w fyne.Window, l *Logger) *widget.Button {
	return widget.NewButton("Submit", func() {
		err := mergePDF(SelectedFiles, OutputFilePath)
		if err != nil {
			l.Write(fmt.Sprintf("Some error occurred while merging: %v", err))
			dialog.ShowError(errors.New("Error: Check log file for detailed response"), w)
		} else {
			var d dialog.Dialog
			quitBtn := widget.NewButton("OK", func() {
				d.Hide()
				a.Quit()
			})
			content := container.NewVBox(
				widget.NewLabel("PDFs merged successfully!"),
				container.NewHBox(layout.NewSpacer(), quitBtn),
			)
			d = dialog.NewCustomWithoutButtons("Success", content, w)
			d.Resize(fyne.NewSize(300, 150))
			d.Show()
		}
	})
}

func homeView(a fyne.App, w fyne.Window, l *Logger) *fyne.Container {
	infoLabel := widget.NewLabel("Please select the files to be merged")
	infoLabel.Wrapping = fyne.TextWrapWord

	fileList := widget.NewList(
		func() int { return len(SelectedFiles) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(filepath.Base(SelectedFiles[i]))
		},
	)

	fileListCard := widget.NewCard("Selected Files", "", fileList)

	outputEntry := widget.NewEntry()
	outputEntry.SetPlaceHolder(OutputFilePath)
	outputEntry.Disable()

	outputBrowseButton := getOutputBrowseButton(outputEntry, w)
	outputRow := container.NewBorder(nil, nil, nil, outputBrowseButton, outputEntry)

	browseFileButton := getBrowseFileButton(fileList, w)
	submitButton := getSubmitButton(a, w, l)

	buttons := container.New(
		layout.CustomPaddedLayout{
			TopPadding:    10,
			BottomPadding: 10,
			LeftPadding:   15,
			RightPadding:  0,
		}, container.NewHBox(layout.NewSpacer(), browseFileButton, submitButton))

	return container.NewBorder(
		infoLabel,
		container.NewVBox(outputRow, buttons),
		nil, nil,
		fileListCard,
	)
}
