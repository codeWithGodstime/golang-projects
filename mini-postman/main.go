package main

import (
	"encoding/json"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	"github.com/codeWithGodstime/mini-postman/core"
	"github.com/codeWithGodstime/mini-postman/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("")
	var appData core.AppData

	file, err := os.Open("db.json")
	if err != nil {
		panic("Could not read system files")
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&appData)

	if err != nil {
		fmt.Println(err)
	}

	_, sideBarContainer := ui.SmallSideBar(appData.Collections)

	layoutSplit := container.NewHSplit(
		sideBarContainer,
		ui.RequestTabs(),
	)
	layoutSplit.SetOffset(0.2)

	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New Collection", func() {
				dialog.NewEntryDialog("Enter a name for your collection", "", func(name string) {
					if name != "" {
						newCol := core.Collection{Name: name, Requests: []core.Request{}, Folders: []core.Folder{}}
						appData.Collections = append(appData.Collections, newCol)

						// data, _ := json.MarshalIndent(appData, "", "  ")
						// os.WriteFile("db.json", data, 0644)
						_, sideBarContainer = ui.SmallSideBar(appData.Collections)
						// layoutSplit.Objects[0] = sideBarContainer
						layoutSplit.Refresh()
					}
				}, w).Show()
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
