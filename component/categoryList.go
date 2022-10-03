package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"Twitter", "News", "Facebook"}

func CategoryLayout() fyne.CanvasObject {
	content := container.NewGridWithColumns(1, categoryList())
	return content
}

func categoryList() fyne.CanvasObject {
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("Category", func() {

			})
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Button).SetText(data[lii])
		},
	)

	content := container.NewGridWithColumns(1, list)
	return content
}
