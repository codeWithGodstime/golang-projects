package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/codeWithGodstime/mini-postman/core"
)

var tabIndex = 0
var requestBodyEntry, responseBodyEntry *widget.Entry

// utility function
func formatJSON(input string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(input), "", "  ")
	if err != nil {
		return input // fallback to raw if JSON is invalid
	}
	return out.String()
}

func convertStructToMapForTree(collections []core.Collection) (map[string][]string, map[string]interface{}) {
	treeData := make(map[string][]string)
	idLookup := make(map[string]interface{})

	for _, col := range collections {
		colID := col.Name
		treeData[""] = append(treeData[""], colID)
		idLookup[colID] = col

		// Top-level requests
		for _, req := range col.Requests {
			reqID := colID + "/" + req.Name
			treeData[colID] = append(treeData[colID], reqID)
			idLookup[reqID] = req
		}

		// Folders
		for _, folder := range col.Folders {
			folderID := colID + "/" + folder.Name
			treeData[colID] = append(treeData[colID], folderID)
			idLookup[folderID] = folder

			for _, req := range folder.Requests {
				reqID := folderID + "/" + req.Name
				treeData[folderID] = append(treeData[folderID], reqID)
				idLookup[reqID] = req
			}
		}
	}

	return treeData, idLookup
}

func SmallSideBar(collections []core.Collection) (fyne.CanvasObject, fyne.CanvasObject) {

	treeData, idLookup := convertStructToMapForTree(collections)

	tree := widget.NewTree(
		func(uid string) []string {
			return treeData[uid]
		},
		func(uid string) bool {
			_, isBranch := treeData[uid]
			return isBranch
		},
		func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Node")
		},
		func(uid string, branch bool, obj fyne.CanvasObject) {
			parts := strings.Split(uid, "/")
			obj.(*widget.Label).SetText(parts[len(parts)-1])
		},
	)

	tree.OnSelected = func(uid string) {
		obj := idLookup[uid]
		switch v := obj.(type) {
		case core.Request:
			fmt.Println("Request selected:", v.Method, v.URL)
		case core.Folder:
			fmt.Println("Folder:", v.Name)
		case core.Collection:
			fmt.Println("Collection:", v.Name)
		}
	}

	sidebarContainer := container.New(layout.NewStackLayout(), tree)

	return tree, sidebarContainer
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
			fmt.Println(response)

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
			fyne.Do(func() { responseBodyEntry.SetText(responseBody) })

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

	requestBodyEntry := widget.NewMultiLineEntry()
	requestBodyEntry.SetPlaceHolder("Enter request body (JSON)...")

	responseBodyEntry := widget.NewMultiLineEntry()
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

	headerFields := container.NewVBox()

	addHeaderField := func() {
		keyEntry := widget.NewEntry()
		keyEntry.SetPlaceHolder("Header Key")
		keyEntryContainer := container.NewStack(keyEntry)

		valueEntry := widget.NewEntry()
		valueEntry.SetPlaceHolder("Header Value")
		valueEntryContainer := container.New(layout.NewStackLayout(), valueEntry)

		headerRow := container.NewGridWithRows(
			1,
			keyEntryContainer,
			valueEntryContainer, 
		)
		headerFields.Add(headerRow)
	}

	addHeaderField()

	addHeaderButton := widget.NewButton("Add Header", func() {
		addHeaderField()
		headerFields.Refresh()
	})

	formContainer := container.NewVBox(
		headerFields,
		addHeaderButton,
	)

	tabs := container.NewAppTabs()
	headersTabContainer := container.NewTabItem("Headers", formContainer)
	bodyTabContainer := container.NewTabItem("Body", content)

	tabs.Append(headersTabContainer)
	tabs.Append(bodyTabContainer)

	return tabs
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
