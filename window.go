package fyne_extend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

var (
	myTheme = NewTheme(0)
)

func NewWindow(app fyne.App, title string, content fyne.CanvasObject) fyne.Window {
	window := app.NewWindow(title)

	window.SetContent(content)
	app.Settings().SetTheme(myTheme)

	increaseFont := desktop.CustomShortcut{KeyName: fyne.KeyEqual, Modifier: fyne.KeyModifierShift | fyne.KeyModifierSuper}

	window.Canvas().AddShortcut(&increaseFont, func(shortcut fyne.Shortcut) {
		myTheme.factor += 0.1
		content.Refresh()
	})

	decreaseFont := desktop.CustomShortcut{KeyName: fyne.KeyMinus, Modifier: fyne.KeyModifierShift | fyne.KeyModifierSuper}

	window.Canvas().AddShortcut(&decreaseFont, func(shortcut fyne.Shortcut) {
		myTheme.factor -= 0.1
		content.Refresh()
	})

	return window
}
