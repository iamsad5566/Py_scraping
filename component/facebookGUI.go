package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func FacebookLayout() fyne.CanvasObject {
	title := title("Facebook scrper")
	content := container.NewVBox(title)
	return content
}
