package main

import (
	"regexp"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	stackLog         *gtk.Stack
	gtkLogTextBuffer *gtk.TextBuffer

	currentLanguage = languageEnglish
)

type part struct {
	partList   *partList
	listBoxRow *gtk.ListBoxRow
	labelName  *gtk.Label
	labelUuid  *gtk.Label
	smPart     *smPart
}

func (self *part) setUuid(s string) {
	self.labelUuid.SetText(s)
	self.smPart.Uuid = s
}

func (self *part) setTitle(s string, lang string) {
	self.labelName.SetText(s)
	self.smPart.setTitle(s, currentLanguage)
}

func (self *part) destroy() {
	self.listBoxRow.Destroy()

	self.partList.removePart(self)
}

type partList struct {
	listBox          *gtk.ListBox
	buttonAddPart    *gtk.Button
	buttonDeletePart *gtk.Button
	searchEntryPart  *gtk.SearchEntry

	parts      map[*part]struct{}
	activePart *part

	partEditor *partEditor
}

func (pl *partList) init() {
	pl.buttonAddPart.Connect("clicked", func() {
		pl.createNewPart()
	})

	pl.searchEntryPart.Connect("search-changed", func(self *gtk.SearchEntry) {
		text, _ := self.GetText()
		pl.filterVisible(text)
	})
}

func (self *partList) createNewPart() {
	listBoxRow, _ := gtk.ListBoxRowNew()
	labelName, _ := gtk.LabelNew("")
	labelUuid, _ := gtk.LabelNew("")

	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	box.SetHomogeneous(true)
	box.Container.Add(labelName)
	box.Container.Add(labelUuid)

	listBoxRow.Add(box)

	part := &part{
		partList:   self,
		listBoxRow: listBoxRow,
		labelName:  labelName,
		labelUuid:  labelUuid,
		smPart: &smPart{
			Descriptions: map[string]*smPartDescription{},
		},
	}

	self.parts[part] = struct{}{}

	part.listBoxRow.Connect("activate", func() {
		self.partEditor.setEditorActive(true)
		part.partList.setActivePart(part)
		part.partList.partEditor.reloadValues()
	})

	self.listBox.Add(part.listBoxRow)
	part.listBoxRow.ShowAll()
}

func (self *partList) removePart(p *part) {
	p.listBoxRow.Destroy()

	delete(self.parts, p)
}

func (self *partList) filterVisible(s string) {
	if s == "" {
		for part, _ := range self.parts {
			part.listBoxRow.SetVisible(true)
		}
		return
	}

	r := regexp.MustCompile("(?i)" + regexp.QuoteMeta(s))

	for part, _ := range self.parts {
		part.listBoxRow.SetVisible(part.smPart.matches(r))
	}

	self.listBox.ShowAll()
}

func (self *partList) setActivePart(p *part) (ok bool) {
	_, ok = self.parts[p]
	if !ok {
		return
	}

	self.activePart = p

	return
}

func (self *partList) setUuidOfActive(s string) {
	if self.activePart != nil {
		self.activePart.setUuid(s)
	}
}

func (self *partList) setTitleOfActive(s string) {
	if self.activePart != nil {
		self.activePart.setTitle(s, currentLanguage)
	}
}

func (self *partList) setDescriptionOfActive(s string) {
	if self.activePart != nil {
		self.activePart.smPart.setDescription(s, currentLanguage)
	}
}

type partEditor struct {
	partList *partList

	stackPartData *gtk.Stack

	entryName           *gtk.Entry
	textViewDescription *gtk.TextView

	entryUuid        *gtk.Entry
	buttonRandomUuid *gtk.Button
	imageUuidStatus  *gtk.Image
	textViewPartData *gtk.TextView

	comboBoxLanguage *gtk.ComboBoxText
}

func (pe *partEditor) init() {
	pe.entryName.Connect("changed", func(self *gtk.Entry) {
		text, _ := self.GetText()
		pe.partList.setTitleOfActive(text)
	})

	pe.entryUuid.Connect("changed", func(self *gtk.Entry) {
		text, err := self.GetText()
		if err != nil {
			logError(err)
			return
		}

		if validateUuid(text) {
			pe.imageUuidStatus.SetProperty("stock", "gtk-ok")
		} else {
			pe.imageUuidStatus.SetProperty("stock", "gtk-no")
		}
	})

	pe.buttonRandomUuid.Connect("clicked", func() {
		pe.partList.setUuidOfActive(randomUuid())
	})

	pe.comboBoxLanguage.Connect("changed", func(self *gtk.ComboBoxText) {
		lang := self.GetActiveID()
		currentLanguage = lang

		pe.reloadValues()
	})
}

func (self *partEditor) reloadValues() {
	self.entryName.SetText(
		self.partList.activePart.smPart.getTitle(currentLanguage),
	)

	tb, _ := self.textViewDescription.GetBuffer()
	tb.SetText(
		self.partList.activePart.smPart.getDescription(currentLanguage),
	)

	self.entryUuid.SetText(
		self.partList.activePart.smPart.Uuid,
	)

	tb, _ = self.textViewPartData.GetBuffer()
	tb.SetText(
		self.partList.activePart.smPart.PartData,
	)
}

func (self *partEditor) setEditorActive(b bool) {
	if b {
		self.stackPartData.SetVisibleChildName("1")
	} else {
		self.stackPartData.SetVisibleChildName("0")
	}
}

func gtkInit() {
	var err error

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

		entryName:           getObject("entryName").(*gtk.Entry),
		textViewDescription: getObject("textViewDescription").(*gtk.TextView),

		entryUuid:        getObject("entryUuid").(*gtk.Entry),
		buttonRandomUuid: getObject("buttonRandomUuid").(*gtk.Button),
		imageUuidStatus:  getObject("imageUuidStatus").(*gtk.Image),
		textViewPartData: getObject("textViewPartData").(*gtk.TextView),

		comboBoxLanguage: getObject("comboBoxLanguage").(*gtk.ComboBoxText),
	}

	pl := &partList{
		listBox:          getObject("listParts").(*gtk.ListBox),
		buttonAddPart:    getObject("buttonAddPart").(*gtk.Button),
		buttonDeletePart: getObject("buttonDeletePart").(*gtk.Button),
		searchEntryPart:  getObject("searchEntryPart").(*gtk.SearchEntry),

		parts:      map[*part]struct{}{},
		partEditor: pe,
	}

	pe.partList = pl

	pe.init()
	pl.init()

	// Log stuff
	stackLog = getObject("stackLog").(*gtk.Stack)
	textViewLog := getObject("textViewLog").(*gtk.TextView)
	gtkLogTextBuffer, err = textViewLog.GetBuffer()
	if err != nil {
		panic(err)
	}

	windowMain := getObject("windowMain").(*gtk.Window)
	windowMain.Connect("destroy", func() {
		gtk.MainQuit()
	})

	windowMain.ShowAll()
	gtk.Main()
}
