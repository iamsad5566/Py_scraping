package component

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// horizontalSpace serving as the border to split two neighboured objects
var horizontalSpace string = "           " +
	"                        "

// title returns a canvasObject containing the label showing the current page's gist.
func tittle(str string) fyne.CanvasObject {
	titleText := canvas.NewText(str, nil)
	titleText.TextSize = 50
	titleText.TextStyle.Bold = true
	space := canvas.NewText(" ", color.White)
	content := container.NewBorder(space, space, space, space, container.NewCenter(titleText))
	return content
}

// tmp
// TwitterLayout contains all the elements belonging to twitter related manipulations
func TwitterLayout() fyne.CanvasObject {
	title := tittle("Twitter scraper")

	// Insert a selection
	selectContent := widget.NewSelect(selects, func(s string) {
		selection = &s
	})

	selectContainer := container.NewBorder(canvas.NewText("", color.Opaque),
		canvas.NewText("", color.Opaque), canvas.NewText(horizontalSpace, color.Opaque),
		canvas.NewText(horizontalSpace, color.Opaque), selectContent)

	options := twiOption()
	button := twiButton()

	content := container.NewVBox(title, selectContainer, options, button)

	return content
}
