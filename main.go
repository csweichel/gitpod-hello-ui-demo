package main

import (
	"github.com/andlabs/ui"
)

var mainwin *ui.Window

func setupUI() {
	mainwin = ui.NewWindow("libui Control Gallery", 640, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

    ui.MsgBox(mainwin, "Hello World", "Showing a GTK message box in Gitpod")
    ui.Quit()
}

func main() {
	ui.Main(setupUI)
}