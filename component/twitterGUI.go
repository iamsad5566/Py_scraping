package component

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type selection struct {
	single   bool
	multiple bool
}

var checkedOpt selection = selection{false, false}

func TwitterLayout() fyne.CanvasObject {
	title := title("Twitter scraper")
	checkBoxes := twiCheckBox()
	button := twiButton()
	content := container.NewVBox(title, checkBoxes, button)
	return content
}

func twiCheckBox() fyne.CanvasObject {
	label := canvas.NewText("Select a mode to execute:", color.NRGBA{60, 80, 255, 255})
	label.TextSize = 20
	space := canvas.NewText(" ", color.Opaque)
	space.TextSize = 10
	singleBox := widget.NewCheck("Single searching", func(b bool) {
		checkedOpt.single = b
	})
	multipleBox := widget.NewCheck("Multiple searching", func(b bool) {
		checkedOpt.multiple = b
	})
	content := container.NewCenter(container.NewVBox(label, space,
		container.NewCenter(container.NewVBox(singleBox, multipleBox))))
	return content
}

func twiButton() fyne.CanvasObject {
	button := widget.NewButton("                   Go!                   ", func() {
		if checkedOpt.multiple && checkedOpt.single {
			fmt.Println("run multiple and single!")
		} else if !checkedOpt.multiple && checkedOpt.single {
			fmt.Println("run single!")
		} else if checkedOpt.multiple && !checkedOpt.single {
			fmt.Println("run multiple!")
		} else {
			fmt.Println("Please select a mode!")
		}
	})
	space := canvas.NewText(" ", color.Opaque)
	horizontal := canvas.NewText(horizontalSpace, color.Opaque)
	content := container.NewBorder(space, space, horizontal, horizontal, button)
	return content
}
