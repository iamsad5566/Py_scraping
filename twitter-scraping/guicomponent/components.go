package guicomponent

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var width float32 = 600
var height float32 = 800

type selection struct {
	single   bool
	multiple bool
}

var checkedOpt selection = selection{false, false}

func Open() {
	guiApp := app.New()
	window := guiApp.NewWindow("Twitter scraper GUI")
	window.Resize(fyne.NewSize(width, height))
	window.SetFixedSize(true)
	window.SetContent(overallLayout())
	window.ShowAndRun()
}

func overallLayout() fyne.CanvasObject {
	title := getTitle()
	checkBoxes := getCheckBox()
	button := getButton()
	top := canvas.NewText(" ", color.Opaque)
	top.TextSize = 50
	content := container.NewBorder(top, nil, nil, nil, container.NewVBox(title, checkBoxes, button))
	return content
}

func getTitle() fyne.CanvasObject {
	title := canvas.NewText("Twitter Scraper", nil)
	title.TextSize = 50
	title.TextStyle.Bold = true
	space := canvas.NewText(" ", color.White)
	content := container.NewBorder(space, space, space, space, container.NewCenter(title))
	return content
}

func getCheckBox() fyne.CanvasObject {
	label := canvas.NewText("Select a mode to execute:", color.NRGBA{60, 80, 255, 255})
	label.TextSize = 20
	singleBox := widget.NewCheck("Single searching", func(b bool) {
		checkedOpt.single = b
	})
	multipleBox := widget.NewCheck("Multiple searching", func(b bool) {
		checkedOpt.multiple = b
	})
	content := container.NewCenter(container.NewVBox(label,
		container.NewCenter(container.NewVBox(singleBox, multipleBox))))
	return content
}

func getButton() fyne.CanvasObject {
	button := widget.NewButton("Go!", func() {
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
	horizontal := canvas.NewText("                ", color.Opaque)
	content := container.NewBorder(space, space, horizontal, horizontal, button)
	return content
}
