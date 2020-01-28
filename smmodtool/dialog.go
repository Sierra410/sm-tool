package main

import (
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
