package main

import (
	"os/exec"
	"runtime"

	"github.com/gotk3/gotk3/gtk"
)

func dialogYesNo(title, message string) bool {
	messageDialog := gtk.MessageDialogNew(
		windowMain,
		0,
		gtk.MESSAGE_QUESTION,
		gtk.BUTTONS_YES_NO,
		message,
	)
	messageDialog.SetTitle(title)
	response := (messageDialog.Run() == gtk.RESPONSE_YES)
	messageDialog.Destroy()
	return response
}

func dialogInfo(title, message string) {
	messageDialog := gtk.MessageDialogNew(
		windowMain,
		0,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_OK,
		message,
	)
	messageDialog.SetTitle(title)
	messageDialog.Run()
	messageDialog.Destroy()
}

func dialogDir(title string) string {
	fcnd, _ := gtk.FileChooserNativeDialogNew(
		title,
		windowMain,
		gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER,
		"Select",
		"Cancel",
	)

	if fcnd.Run() == int(gtk.RESPONSE_ACCEPT) {
		return fcnd.GetFilename()
	} else {
		return ""
	}
}

func openUrl(url string) {
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	}
}
