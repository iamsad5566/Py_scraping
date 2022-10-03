package main

import (
	"image/color"
	"py-scraper-gui/component"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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
	window.SetContent(overallLayout())
	window.ShowAndRun()
}

// overLayout contains all the components and make them be arranged appropriately
func overallLayout() fyne.CanvasObject {
	// caSize := fyne.Size{
	// 	Width:  200,
	// 	Height: 800,
	// }

	category := component.ChoiceLayout()
	space := canvas.NewText(" ", color.Opaque)
	twitter := component.TwitterLayout()
	line := canvas.NewLine(color.Black)
	content := container.NewBorder(space, space, container.NewCenter(category), space, container.NewHBox(line, container.NewCenter(twitter)))

	return content
}
