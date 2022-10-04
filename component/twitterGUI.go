package component

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var option = []string{"Single searching", "Multiple searching"}
var keyword *string

func TwitterLayout() fyne.CanvasObject {
	title := title("Twitter scraper")
	options := twiOption()
	button := twiButton()
	content := container.NewVBox(title, options, button)
	return content
}

func twiOption() fyne.CanvasObject {
	label := canvas.NewText("Select a mode to execute:", color.NRGBA{60, 80, 255, 255})
	label.TextSize = 20
	horizontal := canvas.NewText(horizontalSpace, color.Opaque)
	space := canvas.NewText(" ", color.Opaque)
	space.TextSize = 10

	keywordBox := container.NewBorder(space, space, horizontal, horizontal, space)
	choice := widget.NewRadioGroup(option, func(s string) {
		keywordBox.RemoveAll()
		keywordBox.Add(searchKeyword(s))
		keywordBox.Refresh()
	})
	content := container.NewVBox(container.NewCenter(container.NewVBox(label, space,
		container.NewCenter(container.NewVBox(choice)))), keywordBox)
	return content
}

func searchKeyword(str string) fyne.CanvasObject {
	labelText := ""
	if str == "Single searching" {
		labelText = "Keyword:"
	} else if str == "Multiple searching" {
		labelText = "Keywords (list):"
	}
	label := canvas.NewText(labelText, nil)
	label.TextSize = 15
	space := canvas.NewText(" ", color.Opaque)
	entry := widget.NewEntry()
	size := fyne.Size{Width: 200, Height: 60}
	keyword = &entry.Text

	content := container.NewVBox(space, label, container.NewGridWrap(size, entry))
	return content
}

func twiButton() fyne.CanvasObject {
	button := widget.NewButton("                   Go!                   ", func() {
		fmt.Println(*keyword)
	})
	space := canvas.NewText(" ", color.Opaque)
	horizontal := canvas.NewText(horizontalSpace, color.Opaque)
	content := container.NewBorder(space, space, horizontal, horizontal, button)
	return content
}
