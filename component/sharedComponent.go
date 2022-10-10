package component

import (
	"image/color"
	"py-scraper-gui/component/twigui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var selects = []string{"searchByKeyword", "searchByID"}

// horizontalSpace serving as the border to split two neighboured objects
var horizontalSpace string = "           " +
	"                        " +
	"                        "

// title returns a canvasObject containing the label showing the current page's gist.
func Tittle(str string) fyne.CanvasObject {
	titleText := canvas.NewText(str, nil)
	titleText.TextSize = 50
	titleText.TextStyle.Bold = true
	space := canvas.NewText(" ", color.White)
	content := container.NewBorder(space, space, space, space, container.NewCenter(titleText))
	return content
}

// TwitterLayout returns the GUI components depending on the selection
func TwitterLayout() fyne.CanvasObject {
	title := Tittle("Twitter scraper")
	selected := "searchByKeyword"
	searchGUI := twigui.BuildObj(selected)
	gui := twigui.Draw(searchGUI)

	// Insert a selection
	selectContent := widget.NewSelect(selects, func(s string) {
		selected = s
		searchGUI = twigui.BuildObj(selected)
		gui.RemoveAll()
		gui.Add(twigui.Draw(searchGUI))
	})

	// Creat space
	space := canvas.NewText("", color.Opaque)
	
	// Create a container for the selection bar
	selectContainer := container.NewBorder(space, space, space,
		space, selectContent)

	content := container.NewVBox(title, selectContainer, gui)

	return content
}
