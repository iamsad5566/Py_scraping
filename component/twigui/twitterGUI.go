package twigui

import "fyne.io/fyne/v2"

type TwiPage interface {
	SetLabel(str string)
	SetChoice(str string)
	GetOptionContainer() fyne.CanvasObject
	GetButton() fyne.CanvasObject
}
