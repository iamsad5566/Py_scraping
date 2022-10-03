package main

import (
	"py-scraper-gui/component"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var width float32 = 1200
var height float32 = 800

func main() {
	openGUI()
}

func openGUI() {
	guiApp := app.New()
	window := guiApp.NewWindow("Twitter scraper GUI")
	window.Resize(fyne.NewSize(width, height))
	// window.SetFixedSize(true)
	window.SetContent(overLayout())
	window.ShowAndRun()
}

// overLayout contains all the components and make them be arranged appropriately
func overLayout() fyne.CanvasObject {
	caSize := fyne.Size{
		Width:  200,
		Height: 800,
	}

	category := component.CategoryLayout()
	twitter := component.TwitterLayout()
	content := container.NewHBox(container.NewGridWrap(caSize, category), twitter)
	return content
}
