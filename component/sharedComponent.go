package component

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var horizontalSpace string = "              " +
	"                                  " +
	"                                  " +
	"                                  "

func title(str string) fyne.CanvasObject {
	titleText := canvas.NewText(str, nil)
	titleText.TextSize = 50
	titleText.TextStyle.Bold = true
	space := canvas.NewText(" ", color.White)
	content := container.NewBorder(space, space, space, space, container.NewCenter(titleText))
	return content
}
