package ui

import (
	"dicomdir-generator/backend"
	"errors"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func RunApp() {
	a := app.NewWithID("com.example.dicomdirgenerator")
	w := a.NewWindow("DICOMDIR Generator")
	w.Resize(fyne.NewSize(520, 320))

	// Folder picker UI
	folderBox, getFolder := FolderPickerWidget(w, "Please select a DICOM folder:")

	// Progress label
	progressLabel, updateProgress := ProgressLabel("Select a folder and click Generate")

	// Generate button
	generateBtn := StyledButton("Generate DICOMDIR", func() {
		path := getFolder()
		if path == "" {
			dialog.ShowError(errors.New("Please select a folder first"), w)
			return
		}

		updateProgress("Generating DICOMDIR...")

		err := backend.GenerateDicomDir(path)
		if err != nil {
			updateProgress("Error occurred.")
			dialog.ShowError(err, w)
			return
		}

		updateProgress("DICOMDIR generated successfully!")
		dialog.ShowInformation("Success", "DICOMDIR generated in the same folder!", w)
	})

	githubURL, _ := url.Parse("https://github.com/rumiani/dicomdir-generator")
	label := widget.NewLabel("⭐ Give me a star on")

	githubLink := widget.NewHyperlink("GitHub", githubURL)
	githubRow := container.NewHBox(label, githubLink)

	// If you want it clickable:
	// footer := widget.NewHyperlink("Give me a star on GitHub", parseURL("https://github.com/your/repo"))

	// Main layout (content + footer)
	content := container.NewVBox(
		folderBox,
		progressLabel,
		generateBtn,
		layout.NewSpacer(),
		githubRow,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
