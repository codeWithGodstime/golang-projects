package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/codeWithGodstime/mini-postman/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("")

	layoutSplit := container.NewHSplit(
		ui.SmallSideBar(),
		// ui.MainContent(),
		ui.RequestTabs(),
	)

	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New Collection", func() {
				// new doc logic
			}),
			fyne.NewMenuItem("Open Collection", func() {
				// new doc logic
			}),
			fyne.NewMenuItem("Save Collection", func() {
				// new doc logic
			}),
			fyne.NewMenuItem("New Request", func() {
				// open logic
			}),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Undo", nil),
		),
	))

	w.SetContent(layoutSplit)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
