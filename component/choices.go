package component

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"Twitter", "News", "Facebook"}

func ChoiceLayout() fyne.CanvasObject {
	space := canvas.NewText(" ", color.Opaque)
	content := container.NewGridWithRows(3, space, choiceList(), space)
	return content
}

func choiceList() fyne.CanvasObject {
	choice := widget.NewRadioGroup(data, func(s string) {
		fmt.Println(s)
	})

	space := canvas.NewText("        ", color.Opaque)

	content := container.NewGridWithColumns(3, space, container.NewCenter(choice))
	return content
}
