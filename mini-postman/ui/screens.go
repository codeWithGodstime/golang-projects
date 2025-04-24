package ui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var tabIndex = 0

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

	requestTypeDropDownButton := widget.NewSelect([]string{"GET", "POST", "PATCH", "DELETE"}, func(s string) {})
	entry := widget.NewEntry()
	sendButton := widget.NewButton("Send Request", nil)

	wrapper := container.NewBorder(
		nil, nil,
		requestTypeDropDownButton,
		sendButton,
		entry,
	)
	return wrapper
}

func MainContent() fyne.CanvasObject {

	requestBodyEntry := widget.NewMultiLineEntry()
	requestBodyEntry.SetPlaceHolder("Enter request body(JSON)....")

	responseBodyEntry := widget.NewMultiLineEntry()
	responseBodyEntry.SetPlaceHolder("Response will appear here...")
	responseBodyEntry.Disable()

	bodySplit := container.NewVSplit(
		requestBodyEntry,
		responseBodyEntry,
	)
	bodySplitWrapper := container.NewStack(bodySplit)

	content := container.NewStack(
			container.NewVBox(
			ToolBar(),
			widget.NewSeparator(),
			RequestEntry(),
			widget.NewSeparator(),
			bodySplitWrapper,
		),
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
