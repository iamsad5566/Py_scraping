package twigui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SearchByID struct {
	Choice     string
	LabelText  string
	EntryInput string
	RunTime    string
}

var _ TwiGUI = (*SearchByID)(nil)
var option = []string{"Single searching", "Multiple searching"}

func (s *SearchByID) SetChoice(str string) {
	if len(str) > 0 {
		s.Choice = str
	} else {
		s.Choice = ""
		s.EntryInput = ""
		s.RunTime = ""
	}
}

func (s *SearchByID) SetLabel(str string) {
	label := ""
	if str == "Single searching" {
		label = "ID:"
	} else if str == "Multiple searching" {
		label = "IDs:"
	}
	s.LabelText = label
}

func (s *SearchByID) SetOption() fyne.CanvasObject {
	label := canvas.NewText("Select a mode to execute:", color.NRGBA{60, 80, 255, 255})
	label.TextSize = 20
	space := canvas.NewText(" ", color.Opaque)
	space.TextSize = 10

	entryBox := container.NewCenter(space)
	choice := widget.NewRadioGroup(option, func(str string) {
		entryBox.RemoveAll()
		s.SetChoice(str)

		if len(str) > 0 {
			entryBox.Add(s.SetEntries(str))
		}

		entryBox.Refresh()
	})

	size := fyne.Size{Width: 200, Height: 60}
	content := container.NewVBox(container.NewCenter(container.NewVBox(label, space,
		container.NewCenter(container.NewGridWrap(size, choice)))), entryBox)
	return content
}

func (s *SearchByID) SetEntries(str string) fyne.CanvasObject {
	s.SetLabel(str)
	label := canvas.NewText(s.LabelText, nil)
	label.TextSize = 15
	space := canvas.NewText(" ", color.Opaque)
	sizeEntry := fyne.Size{Width: 250, Height: 40}

	entryID := widget.NewEntry()
	s.EntryInput = entryID.Text

	labelRun := canvas.NewText("Run:", nil)
	labelRun.TextSize = 15
	entryRun := widget.NewEntry()
	s.RunTime = entryRun.Text

	content := container.NewVBox(space, label, container.NewGridWrap(sizeEntry, entryID), labelRun, container.NewGridWrap(sizeEntry, entryRun))
	return content
}

func (s *SearchByID) Layout() fyne.CanvasObject {
	tmp := container.NewVBox(canvas.NewText("", color.Opaque))
	return tmp
}
