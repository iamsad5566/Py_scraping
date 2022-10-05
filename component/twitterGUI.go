package component

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

var option = []string{"Single searching", "Multiple searching"}
var keyword *string
var selected *string

// TwitterLayout contains all the elements belonging to twitter related manipulations
func TwitterLayout() fyne.CanvasObject {
	title := title("Twitter scraper")
	options := twiOption()
	button := twiButton()
	content := container.NewVBox(title, options, button)
	return content
}

// twiOption returns a canvasObject containing the options and keyword entry
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

// searchKeyword returns a canvas object containing an input box that allows the user to key in the keyword.
func searchKeyword(str string) fyne.CanvasObject {
	labelText := ""
	selection := "0"
	if str == "Single searching" {
		labelText = "Keyword:"
		selection = "1"
		selected = &selection
	} else if str == "Multiple searching" {
		selection = "2"
		selected = &selection
		labelText = "Keywords (list):"
	}
	label := canvas.NewText(labelText, nil)
	label.TextSize = 15
	space := canvas.NewText(" ", color.Opaque)
	entry := widget.NewEntry()
	size := fyne.Size{Width: 250, Height: 40}
	keyword = &entry.Text

	content := container.NewVBox(space, label, container.NewGridWrap(size, entry))
	return content
}

// twiButton returns a canvasOject containing the "GO!" button, the scraping program will be executed after pressing this button.
func twiButton() fyne.CanvasObject {
	notification := canvas.NewText("", color.Opaque)
	notification.TextSize = 30
	notification.TextStyle.Bold = true

	button := widget.NewButton("                   Go!                   ", func() {
		if selected == nil {
			notification.Text = "Please select a mode!"
		} else if *keyword == "" {
			notification.Text = "Plaese fill in the keyword!"
		} else {
			notification.Text = ""
			notification.Color = color.Opaque
			notification.Refresh()
			twiExec(keyword, selected)
			return
		}

		notification.Color = color.RGBA{255, 0, 0, 255}
		notification.Refresh()
	})
	space := canvas.NewText(" ", color.Opaque)
	horizontal := canvas.NewText(horizontalSpace, color.Opaque)
	content := container.NewVBox(container.NewBorder(space, space, horizontal, horizontal, button), container.NewCenter(notification))
	return content
}

// twiExec is responsible for triggering the python scraping program to be executed.
func twiExec(keyword *string, selection *string) {
	if *keyword == "" || *selection == "" {
		return
	}

	query := strings.ReplaceAll(*keyword, " ,", ",")
	query = strings.ReplaceAll(query, ", ", ",")
	args := strings.Split(query, ",")
	args = append([]string{"twitter-scraping/main.py", *selection}, args...)

	cmd := exec.Command("venv/bin/python3", args...)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
