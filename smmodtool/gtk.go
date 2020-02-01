package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	windowMain   *gtk.Window
	mainPartList *partList
	logger       *gtkLogger

	currentLanguage = languageEnglish
)

func gtkInit(onReady func()) chan bool {
	gtk.Init(nil)

	mainBuilder, _ := gtk.BuilderNew()
	mainBuilder.AddFromString(gladeMain)

	getObject := func(s string) glib.IObject {
		obj, err := mainBuilder.GetObject(s)
		if err != nil {
			panic(err)
		}
		return obj
	}

	pe := &partEditor{
		stackPartData: getObject("stackPartData").(*gtk.Stack),

		entryName:             getObject("entryName").(*gtk.Entry),
		textBufferDescription: getObject("textBufferDescription").(*gtk.TextBuffer),

		entryUuid:           getObject("entryUuid").(*gtk.Entry),
		buttonRandomUuid:    getObject("buttonRandomUuid").(*gtk.Button),
		checkButtonIsBlock:  getObject("checkButtonIsBlock").(*gtk.CheckButton),
		labelUuidStatus:     getObject("labelUuidStatus").(*gtk.Label),
		textBufferPartData:  getObject("textBufferPartData").(*gtk.TextBuffer),
		textBufferJsonError: getObject("textBufferJsonError").(*gtk.TextBuffer),

		comboBoxLanguage: getObject("comboBoxLanguage").(*gtk.ComboBoxText),
	}

	pl := &partList{
		buttonSave:    getObject("buttonSave").(*gtk.Button),
		buttonLoad:    getObject("buttonLoad").(*gtk.Button),
		buttonCompile: getObject("buttonCompile").(*gtk.Button),

		listBox:          getObject("listParts").(*gtk.ListBox),
		buttonSortUp:     getObject("buttonSortUp").(*gtk.Button),
		buttonSortDown:   getObject("buttonSortDown").(*gtk.Button),
		buttonAddPart:    getObject("buttonAddPart").(*gtk.Button),
		buttonDeletePart: getObject("buttonDeletePart").(*gtk.Button),
		searchEntryPart:  getObject("searchEntryPart").(*gtk.SearchEntry),

		parts:      []*part{},
		partEditor: pe,
	}

	pe.partList = pl

	pe.init()
	pl.init()

	mainPartList = pl

	// Log stuff

	stackLog := getObject("stackLog").(*gtk.Stack)
	logger = gtkLoggerNew(
		getObject("textViewLog").(*gtk.TextView),
		func() {
			stackLog.SetVisibleChildName("1")
		},
	)

	windowMain = getObject("windowMain").(*gtk.Window)
	windowMain.Connect("destroy", func() {
		gtk.MainQuit()
	})

	windowMain.ShowAll()

	chanFinished := make(chan bool, 1)

	gtk.MainIterationDo(true)

	onReady()

	go func() {
		gtk.Main()

		chanFinished <- true
	}()

	return chanFinished
}
