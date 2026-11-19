package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Bluetooth Hunter")

	title := widget.NewLabel("Bluetooth Hunter")
	status := widget.NewLabel("Ready to scan")

	scanButton := widget.NewButton("Start Scanning", func() {
		status.SetText("Scanner will start soon...")
	})

	content := container.NewVBox(
		title,
		status,
		scanButton,
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(600, 400))
	window.ShowAndRun()
}
