package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// var (
// 	mainWindow *gtk.Window

// 	listParts        *gtk.ListBox
// 	buttonAddPart    *gtk.Button
// 	searchEntryPart  *gtk.SearchEntry
// 	buttonDeletePart *gtk.Button

// 	textViewLog   *gtk.TextView
// 	stackPartData *gtk.Stack

// 	entryName           *gtk.Entry
// 	comboBoxLanguage    *gtk.ComboBox
// 	textViewDescription *gtk.TextView

// 	entryUuid        *gtk.Entry
// 	buttonRandomUuid *gtk.Button
// 	textViewPartData *gtk.TextView
// )

var (
	stackLog         *gtk.Stack
	gtkLogTextBuffer *gtk.TextBuffer
)

func writeLog(format string, important bool, i ...interface{}) {
	gtkLogTextBuffer.Insert(
		gtkLogTextBuffer.GetEndIter(),
		fmt.Sprintf(format, i...),
	)

	if important {
		stackLog.SetVisibleChildName("1")
	}
}

func gtkInit() {
	var err error

	gtk.Init(nil)

	mainBuilder, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal(err)
	}

	//err = mainBuilder.AddFromFile("ui/main.glade")
	err = mainBuilder.AddFromString(gladeMain)
	if err != nil {
		panic(err)
	}

	getObject := func(s string) glib.IObject {
		obj, err := mainBuilder.GetObject(s)
		if err != nil {
			panic(err)
		}
		return obj
	}

	// Get objects
	windowMain := getObject("windowMain").(*gtk.Window)

	// listParts := getObject("listParts").(*gtk.ListBox)
	buttonAddPart := getObject("buttonAddPart").(*gtk.Button)
	searchEntryPart := getObject("searchEntryPart").(*gtk.SearchEntry)
	buttonDeletePart := getObject("buttonDeletePart").(*gtk.Button)

	// Log stuff
	stackLog = getObject("stackLog").(*gtk.Stack)
	textViewLog := getObject("textViewLog").(*gtk.TextView)
	gtkLogTextBuffer, err = textViewLog.GetBuffer()
	if err != nil {
		panic(err)
	}

	stackPartData := getObject("stackPartData").(*gtk.Stack)

	// entryName := getObject("entryName").(*gtk.Entry)
	// comboBoxLanguage := getObject("comboBoxLanguage").(*gtk.ComboBox)
	// textViewDescription := getObject("textViewDescription").(*gtk.TextView)

	// buttonRandomUuid := getObject("buttonRandomUuid").(*gtk.Button)
	// textViewPartData := getObject("textViewPartData").(*gtk.TextView)

	// Connect
	windowMain.Connect("destroy", func() {
		gtk.MainQuit()
	})

	buttonAddPart.Connect("clicked", func() {
		stackPartData.SetVisibleChildName("1")
	})

	buttonDeletePart.Connect("clicked", func() {
		stackPartData.SetVisibleChildName("0")
	})

	searchEntryPart.Connect("search-changed", func(self *gtk.SearchEntry) {
		log.Println(entryGetText(&self.Entry))
	})

	// UUID
	entryUuid := getObject("entryUuid").(*gtk.Entry)
	buttonRandomUuid := getObject("buttonRandomUuid").(*gtk.Button)

	entryUuid.Connect("changed", func(self *gtk.Entry) {
		text, err := self.GetText()
		if err != nil {
			writeLog("Error: %s\n", true, err.Error())
			return
		}

		if validateUuid(text) {
			writeLog("New UUID: %s\n", true, text)
		}
	})

	buttonRandomUuid.Connect("clicked", func() {
		entryUuid.SetText(newUuid4())
	})

	windowMain.ShowAll()
	gtk.Main()
}

// Helper functions

func entryGetText(entry *gtk.Entry) string {
	buf, err := entry.GetBuffer()
	if err != nil {
		return ""
	}

	text, err := buf.GetText()
	if err != nil {
		return ""
	}

	return text
}
