package main

import (
	"image/color"
	"py-scraper-gui/component"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var width float32 = 1200
var height float32 = 800

var data = []string{"Twitter", "News", "Facebook"}

func main() {
	openGUI()
}

func openGUI() {
	guiApp := app.New()
	window := guiApp.NewWindow("Scraper GUI")
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

	choice, _ := choiceLayout()
	// space := canvas.NewText(" ", color.Opaque)
	// line := canvas.NewLine(color.Black)
	// content := container.NewBorder(space, space, container.NewCenter(choice), space, container.NewHBox(line, container.NewCenter(component.FacebookLayout())))
	content := container.NewHBox(choice)
	return content
}

func choiceLayout() (fyne.CanvasObject, fyne.CanvasObject) {
	option, chosen := choiceList()
	space := canvas.NewText(" ", color.Opaque)
	content := container.NewGridWithRows(3, space, option, space)
	return content, chosen
}

func choiceList() (fyne.CanvasObject, fyne.CanvasObject) {
	chosen := component.TwitterLayout()
	choice := widget.NewRadioGroup(data, func(s string) {
		if s == "Facebook" {
			chosen = component.FacebookLayout()
		}
		chosen.Refresh()
	})

	space := canvas.NewText("        ", color.Opaque)
	content := container.NewGridWithColumns(3, space, container.NewHBox(choice, chosen))
	return content, chosen
}
