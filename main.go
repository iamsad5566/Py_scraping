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

// openGUI open a new windows which contains all the elements.
func openGUI() {
	guiApp := app.New()
	guiApp.Settings().SetTheme(&myTheme{})
	window := guiApp.NewWindow("Scraper GUI")
	window.Resize(fyne.NewSize(width, height))
	window.SetFixedSize(false)
	window.SetContent(overallLayout())
	window.ShowAndRun()
}

// overLayout contains all the components and make them be arranged appropriately
func overallLayout() fyne.CanvasObject {
	var guiMap = map[string]fyne.CanvasObject{"Twitter": component.TwitterLayout(), "Facebook": component.FacebookLayout(), "News": canvas.NewText("", color.Opaque)}
	chosen := container.NewCenter(guiMap[data[0]])
	choice := widget.NewRadioGroup(data, func(s string) {
		chosen.RemoveAll()
		chosen.Add(guiMap[s])
		chosen.Refresh()
	})

	space := canvas.NewText("        ", color.Opaque)
	choiceSpace := canvas.NewText(" ", color.Opaque)
	line := canvas.NewLine(color.Black)

	choiceList := container.NewGridWithRows(3, space, container.NewGridWithColumns(3, choiceSpace, choice, choiceSpace), space)
	content := container.NewBorder(space, space, container.NewCenter(choiceList), space, container.NewHBox(line, chosen))
	return content
}
