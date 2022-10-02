package guicomponent

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var width float32 = 600
var height float32 = 800

type selection struct {
	single   bool
	multiple bool
}

var checkedOpt selection

func Open() {
	guiApp := app.New()
	window := guiApp.NewWindow("Twitter scraper GUI")
	window.Resize(fyne.NewSize(width, height))
	window.SetFixedSize(true)
	window.SetContent(checkBox())
	window.ShowAndRun()
}

func checkBox() fyne.CanvasObject {
	singleBox := widget.NewCheck("Single searching", func(b bool) {
		checkedOpt.single = b
	})
	multipleBox := widget.NewCheck("Multiple searching", func(b bool) {
		checkedOpt.multiple = b
	})
	content := container.NewVBox(singleBox, multipleBox)
	return content
}
