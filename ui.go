package main

import (
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

func runUI(outputDir string) {
	OutputFilePath = outputDir
	a := app.New()

	w := newWindow(a)
	w.SetContent(homeView(w))

	w.ShowAndRun()
}

func newWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("PDFManager")
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	w.SetFixedSize(true)

	return w
}

func homeView(w fyne.Window) *fyne.Container {
	infoLabel := widget.NewLabel("Please select the files to be merged")
	infoLabel.Wrapping = fyne.TextWrapWord

	fileList := widget.NewList(
		func() int { return len(SelectedFiles) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(SelectedFiles[i])
		},
	)
	fileListCard := widget.NewCard("Selected Files", "", fileList)

	outputEntry := widget.NewEntry()
	outputEntry.SetPlaceHolder(OutputFilePath)
	outputEntry.Disable()

	outputBrowseButton := widget.NewButton("Browse", func() {})

	outputRow := container.NewBorder(nil, nil, nil, outputBrowseButton, outputEntry)

	browseButton := widget.NewButton("Browse", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				return
			}
			SelectedFiles = append(SelectedFiles, reader.URI().Name())
			fileList.Refresh()
		}, w)

		fd.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
		fd.Show()
	})
	submitButton := widget.NewButton("Submit", func() {})

	buttons := container.New(
		layout.CustomPaddedLayout{
			TopPadding:    10,
			BottomPadding: 10,
			LeftPadding:   15,
			RightPadding:  15,
		}, container.NewHBox(layout.NewSpacer(), browseButton, submitButton))

	return container.NewBorder(
		infoLabel,
		container.NewVBox(outputRow, buttons),
		nil, nil,
		fileListCard,
	)
}
