package main

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

func (c *Content) createMenuItems(win fyne.Window) {
	openItem := fyne.NewMenuItem("Open...", c.openFunc(win))
	saveItem := fyne.NewMenuItem("Save", c.saveFunc(win))
	content.SaveMenuItem = saveItem
	content.SaveMenuItem.Disabled = true
	saveAsItem := fyne.NewMenuItem("Save as...", c.saveAsFunc(win))

	fileMenu := fyne.NewMenu("File", openItem, saveItem, saveAsItem)
	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}

var mdExt = []string{".md", ".MD", ".mD", ".Md"}
var filter = storage.NewExtensionFileFilter(mdExt)

func (c *Content) openFunc(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if read == nil { // user cancelled
				return
			}

			defer read.Close()

			data, err := io.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			c.EditWidget.SetText(string(data))
			c.CurrentFile = read.URI()
			win.SetTitle(win.Title() + " - " + read.URI().Name())
			c.SaveMenuItem.Disabled = false
		}, win)

		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (c *Content) saveFunc(win fyne.Window) func() {
	return func() {
		if c.CurrentFile != nil {
			write, err := storage.Writer(c.CurrentFile)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			write.Write([]byte(c.EditWidget.Text))
			fmt.Println("executed write.Write")
			defer write.Close()
		}
	}
}

func (c *Content) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if write == nil { // user cancelled
				return
			}

			// Get file extension in lowercase for case-insensitive comparison
			ext := strings.ToLower(filepath.Ext(write.URI().Name()))
			if ext != ".md" {
				dialog.ShowInformation("Error", "Please save as a Markdown file (.md extension)", win)
				return
			}

			// save file
			write.Write([]byte(c.EditWidget.Text))
			c.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + " - " + write.URI().Name())
			content.SaveMenuItem.Disabled = false
		}, win)

		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}
