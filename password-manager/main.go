package main

import (
	"crypto/sha256"
	"encoding/hex"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Vault")
	w.Resize(fyne.NewSize(500, 500))

	title := canvas.NewText("The Vault", color.Opaque)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	generatePasswordScreen := func(masterPassword string) fyne.CanvasObject {
		emailEntry := widget.NewEntry()
		emailEntry.SetPlaceHolder("Enter your email")

		serviceEntry := widget.NewEntry()
		serviceEntry.SetPlaceHolder("Enter service name or URL")

		passwordLabel := widget.NewLabel("")

		generateBtn := widget.NewButton("Generate Password", func() {
			email := emailEntry.Text
			service := serviceEntry.Text

			if email == "" || service == "" {
				passwordLabel.SetText("Please enter both email and service.")
				return
			}

			// Call your password generation logic here
			password := generatePassword(masterPassword, email, service)
			passwordLabel.SetText(password)
		})

		return container.NewPadded(
			container.NewVBox(
				widget.NewLabel("Generate Password for Service"),
				emailEntry,
				serviceEntry,
				generateBtn,
				passwordLabel,
			),
		)
	}

	masterPasswordScreen := func() fyne.CanvasObject {
		entry := widget.NewEntry()
		return container.NewPadded(
			container.NewCenter(
				container.NewVBox(
					title,
					entry,
					layout.NewSpacer(),
					widget.NewButton("Unlock With Passphrase", func() {
						w.SetContent(generatePasswordScreen(entry.Text))
					}),
				),
			),
		)
	}

	w.SetContent(masterPasswordScreen())
	w.ShowAndRun()
}

func generatePassword(masterPassword, email, service string) string {
	input := masterPassword + email + service
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])[:16]
}
