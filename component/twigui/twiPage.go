package twigui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type TwiPage interface {
	GetOptionContainer() fyne.CanvasObject
}

// Draw returns the packaged components
func Draw(twipage TwiPage) *fyne.Container {
	opt := twipage.GetOptionContainer()
	return container.NewVBox(opt)
}

func BuildObj(str string) TwiPage {
	if str == "searchByID" {
		return &SearchByID{}
	}

	return &SearchByKeyword{}
}

var option = []string{"Single searching", "Multiple searching"}
