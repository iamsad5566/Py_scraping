package component

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func FacebookLayout() fyne.CanvasObject {
	title := title("Facebook scrper")
	marginVertical := canvas.NewText(" ", color.Opaque)
	marginHorizontal := canvas.NewText(horizontalSpace, color.Opaque)
	content := container.NewBorder(marginVertical, marginVertical, marginHorizontal, nil, container.NewCenter(title))
	return content
}
