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
	Choice    string
	LabelText string
}

// Checked if the SearchByID implements the TwiPage interface
var _ TwiPage = (*SearchByID)(nil)

// SetChoice sets the Choice variale for the SearchByID object
func (s *SearchByID) SetChoice(str string) {
	s.Choice = str
}

// SetLabel sets the LabelText Choice variable for the SearchByID object
func (s *SearchByID) SetLabel(str string) {
	label := ""
	if str == "Single searching" {
		label = "ID:"
	} else if str == "Multiple searching" {
		label = "IDs:"
	}
	s.LabelText = label
}

// GetOptionContainer returns the fyne.CanvasObject that shows the executed mode of the scraper
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
			entryBox.Add(s.getEntriesAndBtnContainer(str))
		}

		entryBox.Refresh()
	})

	size := fyne.Size{Width: 200, Height: 60}
	content := container.NewVBox(container.NewCenter(container.NewVBox(label, space,
		container.NewCenter(container.NewGridWrap(size, choice)))), entryBox)
	return content
}

// getEntruesAndBtnContainer returns the entry and button depending on the chosen option
func (s *SearchByID) getEntriesAndBtnContainer(str string) fyne.CanvasObject {
	s.SetLabel(str)
	label := canvas.NewText(s.LabelText, nil)
	label.TextSize = 15
	space := canvas.NewText(" ", color.Opaque)
	sizeEntry := fyne.Size{Width: 250, Height: 40}

	entryID := widget.NewEntry()

	labelRun := canvas.NewText("Run:", nil)
	labelRun.TextSize = 15
	entryRun := widget.NewEntry()

	btn := s.getButton(&entryID.Text, &entryRun.Text)
	main := container.NewCenter(container.NewVBox(space, label,
		container.NewGridWrap(sizeEntry, entryID), labelRun,
		container.NewGridWrap(sizeEntry, entryRun)))

	content := container.NewVBox(main, btn)

	return content
}

// getButton returns the button so that the user can manipulate the scraper by simply providing some parameters and pressing the button
func (s *SearchByID) getButton(IDstr *string, RunStr *string) fyne.CanvasObject {
	notification := canvas.NewText("", color.Opaque)
	notification.TextSize = 20
	notification.TextStyle.Bold = true

	button := widget.NewButton("              Go!             ", func() {
		if *IDstr == "" {
			notification.Text = "Please fill in the ID!"
		} else {
			notification.Text = ""
			notification.Color = color.Opaque
			notification.Refresh()
			s.twiExec(IDstr, RunStr)
			return
		}
		notification.Color = color.RGBA{255, 0, 0, 255}
		notification.Refresh()

	})

	space := canvas.NewText(" ", color.Opaque)
	content := container.NewVBox(container.NewBorder(space, space, space, space, button), container.NewCenter(notification))
	return content

}

// twiExec is responsible for receiving the parameters and executing the Python program
func (s *SearchByID) twiExec(IDstr *string, RunStr *string) {
	if *IDstr == "" {
		return
	}

	searchQ := strings.ReplaceAll(*IDstr, " ,", ",")
	searchQ = strings.ReplaceAll(searchQ, ", ", ",")
	searchQ = strings.ReplaceAll(searchQ, "@", "")

	var args = []string{"twitter-scraping/mainID.py", strings.Split(s.Choice, " ")[0]}
	if RunStr != nil && *RunStr != "" && *RunStr != "0" {
		args = append(args, *RunStr)
	} else {
		args = append(args, "1")
	}

	IDs := strings.Split(searchQ, ",")
	args = append(args, IDs...)

	cmd := exec.Command("venv/bin/python3", args...)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
