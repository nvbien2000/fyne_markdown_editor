package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Content contains contents of window from markdown editor app
type Content struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var content Content

func main() {
	a := app.New()
	win := a.NewWindow("Markdown")

	// get UI contents
	edit, preview := content.makeUI()
	content.createMenuItems(win)

	// set contents to the window
	win.SetContent(container.NewHSplit(edit, preview))

	// show window & run app
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (c *Content) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	c.EditWidget = edit
	c.PreviewWidget = preview
	edit.OnChanged = preview.ParseMarkdown
	return edit, preview
}
