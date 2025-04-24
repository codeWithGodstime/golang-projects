package ui

import (
	"fmt"
	"io"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/codeWithGodstime/mini-postman/core"
)

var tabIndex = 0
var requestBodyEntry, responseBodyEntry *widget.Entry

func SmallSideBar() fyne.CanvasObject {

	sidebar := container.NewVBox(
		layout.NewSpacer(),
		widget.NewButton("Collection", func() {
			log.Println("new collection created")
		}),
		layout.NewSpacer(),
	)
	sidebarContainer := container.New(layout.NewHBoxLayout(), sidebar)
	sidebarContainer.Resize(fyne.NewSize(20, 0))

	return sidebarContainer
}

func ToolBar() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(
			theme.FileApplicationIcon(), func() {},
		),
		widget.NewToolbarAction(
			theme.SettingsIcon(), func() {},
		),
	)
	return toolbar
}

func RequestEntry() fyne.CanvasObject {

	requestTypeDropDownButton := widget.NewSelect([]string{"GET", "POST", "PATCH", "DELETE"}, nil)

	requestTypeDropDownButton.Selected = "GET"
	requestTypeDropDownButton.Refresh()

	entry := widget.NewEntry()
	sendButton := widget.NewButton("Send Request", func() {
		log.Println(requestTypeDropDownButton.Selected, entry.Text)
		method := requestTypeDropDownButton.Selected
		url := entry.Text

		go func() {
			log.Println(requestBodyEntry.Text)

			response, err := core.MakeRequestController(method, url, nil, requestBodyEntry.Text)
			if err != nil {
				fyne.Do(func() {
					responseBodyEntry.SetText(err.Error())
				})
				return
			}
			defer response.Body.Close()

			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				responseBodyEntry.SetText("Failed to read response body: " + err.Error())
				return
			}

			responseBody := string(bodyBytes)
			fyne.Do(func() {responseBodyEntry.SetText(responseBody)})		
		}()
	})

	wrapper := container.NewBorder(
		nil, nil,
		requestTypeDropDownButton,
		sendButton,
		entry,
	)
	return wrapper
}

func MainContent() fyne.CanvasObject {
	requestBodyEntry = widget.NewMultiLineEntry()
	requestBodyEntry.SetPlaceHolder("Enter request body (JSON)...")

	responseBodyEntry = widget.NewMultiLineEntry()
	responseBodyEntry.SetPlaceHolder("Response will appear here...")
	responseBodyEntry.Disable()

	bodySplit := container.NewVSplit(
		requestBodyEntry,
		responseBodyEntry,
	)

	content := container.NewBorder(
		ToolBar(),
		container.NewVBox(
			widget.NewSeparator(),
			RequestEntry(),
			widget.NewSeparator(),
		),
		nil,
		nil,
		container.NewStack(bodySplit), 
	)

	return content
}

func makeTabWithClose(tabs *container.AppTabs, name string, content fyne.CanvasObject) *container.TabItem {
	var tabItem *container.TabItem

	closeBtn := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		tabs.Remove(tabItem)
	})

	closeBtn.Importance = widget.LowImportance

	tabItem = container.NewTabItemWithIcon(name, nil, content)

	return tabItem
}

func newRequestTab(tabs *container.AppTabs) *container.TabItem {
	tabIndex++
	name := fmt.Sprintf("Request %d", tabIndex)
	tabContentWrapper := container.NewStack(MainContent())
	return makeTabWithClose(tabs, name, tabContentWrapper)
}

func RequestTabs() fyne.CanvasObject {
	tabs := container.NewAppTabs()

	tabs.Append(newRequestTab(tabs))

	plusTab := container.NewTabItemWithIcon("", theme.ContentAddIcon(), widget.NewLabel(""))
	tabs.Append(plusTab)

	tabs.OnSelected = func(tab *container.TabItem) {
		if tab == plusTab {
			newTab := newRequestTab(tabs)
			tabs.Items = append(tabs.Items[:len(tabs.Items)-1], newTab, plusTab)
			tabs.Select(newTab)
			tabs.Refresh()
		}
	}

	return tabs
}
