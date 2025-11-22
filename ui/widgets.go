package ui

import (
	"dicomdir-generator/backend"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// FolderPickerWidget creates a reusable folder picker
func FolderPickerWidget(parent fyne.Window, label string) (fyne.CanvasObject, func() string) {
	folderPathLabel := widget.NewLabel("")
	countLabel := widget.NewLabel("") // for showing DICOM count

	var selectedFolder string

	selectBtn := widget.NewButton("Select Folder", func() {
		fd := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				log.Println("Folder open error:", err)
				return
			}
			if uri == nil {
				return
			}

			path := uri.Path()
			selectedFolder = path
			folderPathLabel.SetText(path)

			// Count DICOM files
			count, err := backend.CountDcmFiles(path)
			if err != nil {
				dialog.ShowError(err, parent)
				return
			}

			if count == 0 {
				countLabel.SetText("No DICOM (.dcm) files found.")
			} else {
				countLabel.SetText(fmt.Sprintf("%d DICOM files found.", count))
			}

		}, parent)
		fd.Show()
	})

	box := container.NewVBox(
		widget.NewLabel(label),
		folderPathLabel,
		countLabel,
		selectBtn,
	)

	return box, func() string {
		return selectedFolder
	}
}

// StyledButton creates a simple styled button
func StyledButton(label string, onTapped func()) *widget.Button {
	return widget.NewButton(label, onTapped)
}

// ProgressLabel creates a label that can be updated dynamically
func ProgressLabel(text string) (*widget.Label, func(string)) {
	label := widget.NewLabel(text)
	update := func(newText string) {
		label.SetText(newText)
	}
	return label, update
}
