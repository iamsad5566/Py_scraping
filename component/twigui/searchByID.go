package twigui

import (
	"image/color"
	"log"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SearchByID struct {
	Choice          string
	LabelText       string
	EntryInput      string
	RunTime         string
	HorizontalSpace string
}

var _ TwiPage = (*SearchByID)(nil)
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

func (s *SearchByID) GetOptionContainer() fyne.CanvasObject {
	label := canvas.NewText("Select a mode to execute:", color.NRGBA{60, 80, 255, 255})
	label.TextSize = 20
	space := canvas.NewText(" ", color.Opaque)
	space.TextSize = 10

	entryBox := container.NewCenter(space)
	choice := widget.NewRadioGroup(option, func(str string) {
		entryBox.RemoveAll()
		s.SetChoice(str)

		if len(str) > 0 {
			entryBox.Add(s.getEntriesContainer(str))
		}

		entryBox.Refresh()
	})

	size := fyne.Size{Width: 200, Height: 60}
	content := container.NewVBox(container.NewCenter(container.NewVBox(label, space,
		container.NewCenter(container.NewGridWrap(size, choice)))), entryBox)
	return content
}

func (s *SearchByID) getEntriesContainer(str string) fyne.CanvasObject {
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

func (s *SearchByID) GetButton() fyne.CanvasObject {
	notification := canvas.NewText("", color.Opaque)
	notification.TextSize = 30
	notification.TextStyle.Bold = true

	button := widget.NewButton("              Go!             ", func() {
		if s.Choice == "" {
			notification.Text = "Please select a mode!"
		} else if s.EntryInput == "" {
			notification.Text = "Please fill in the ID!"
		} else {
			notification.Text = ""
			notification.Color = color.Opaque
			notification.Refresh()
			s.twiExec()
			return
		}
		notification.Color = color.RGBA{255, 0, 0, 255}
		notification.Refresh()

	})

	space := canvas.NewText(" ", color.Opaque)
	horizontal := canvas.NewText(s.HorizontalSpace, color.Opaque)
	content := container.NewVBox(container.NewBorder(space, space, horizontal, horizontal, button), container.NewCenter(notification))
	return content

}

func (s *SearchByID) twiExec() {
	if s.EntryInput == "" || s.Choice == "" {
		return
	}

	s.EntryInput = strings.ReplaceAll(s.EntryInput, " ,", ",")
	s.EntryInput = strings.ReplaceAll(s.EntryInput, ", ", ",")

	var args = []string{"twitter-scraping/main.py", s.Choice}
	if s.RunTime != "" && s.RunTime != "0" {
		args = append(args, s.RunTime)
	}

	IDs := strings.Split(s.EntryInput, ",")
	args = append(args, IDs...)

	cmd := exec.Command("venv/bin/python3", args...)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
