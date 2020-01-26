package main

import (
	"regexp"
	"runtime"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
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

func (self *part) setTitle(s string) {
	self.labelName.SetText(s)
	self.smPart.setTitle(s, currentLanguage)
}

func (self *part) destroy() {
	delete(self.partList.parts, self)

	self.listBoxRow.Destroy()

	self.partList.removePart(self)
	self.partList = nil
	self.listBoxRow = nil
	self.smPart = nil
}

func (self *part) reloadLables() {
	self.labelName.SetText(
		self.smPart.getTitle(currentLanguage),
	)

	self.labelUuid.SetText(
		self.smPart.Uuid,
	)
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
	// It doesn't work otherwhise. For some reason.
	pl.listBox.Connect("row-selected", func(self *gtk.ListBox, listBoxRow *gtk.ListBoxRow) {
		listBoxRow.Activate()
	})

	pl.buttonDeletePart.Connect("clicked", func() {
		if pl.activePart != nil {
			pl.activePart.destroy()
		}
	})

	pl.buttonAddPart.Connect("clicked", func() {
		part := pl.createNewPart()
		part.setUuid(randomUuid())
	})

	pl.searchEntryPart.Connect("search-changed", func(self *gtk.SearchEntry) {
		text, _ := self.GetText()
		pl.filterVisible(text)
	})
}

func (self *partList) createNewPart() *part {
	listBoxRow, _ := gtk.ListBoxRowNew()
	labelName, _ := gtk.LabelNew("")
	labelUuid, _ := gtk.LabelNew("")

	labelUuid.SetEllipsize(pango.ELLIPSIZE_END)
	labelName.SetEllipsize(pango.ELLIPSIZE_END)

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
		part.partList.partEditor.reloadFields()
	})

	self.listBox.Add(part.listBoxRow)
	part.listBoxRow.ShowAll()

	return part
}

func (self *partList) removePart(p *part) {
	if self.activePart == p {
		self.activePart = nil
		self.partEditor.setEditorActive(false)
	}

	_, ok := self.parts[p]
	if ok {
		p.destroy()

		runtime.GC()
	}
}

func (self *partList) filterVisible(s string) {
	if s == "" {
		self.listBox.ShowAll()
		return
	}

	r := regexp.MustCompile("(?i)" + regexp.QuoteMeta(s))

	for part, _ := range self.parts {
		m := part.smPart.matches(r)
		part.listBoxRow.SetVisible(m)
	}
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
		self.activePart.setTitle(s)
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

	entryName             *gtk.Entry
	textViewDescription   *gtk.TextView
	textBufferDescription *gtk.TextBuffer

	entryUuid          *gtk.Entry
	buttonRandomUuid   *gtk.Button
	imageUuidStatus    *gtk.Image
	textViewPartData   *gtk.TextView
	textBufferPartData *gtk.TextBuffer

	comboBoxLanguage *gtk.ComboBoxText
}

func (pe *partEditor) init() {
	pe.entryName.Connect("changed", func(self *gtk.Entry) {
		text, _ := self.GetText()
		pe.partList.setTitleOfActive(text)
	})

	pe.entryUuid.Connect("changed", func(self *gtk.Entry) {
		text, _ := self.GetText()

		if validateUuid(text) {
			pe.imageUuidStatus.SetProperty("stock", "gtk-ok")
			pe.partList.setUuidOfActive(text)
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

		pe.reloadFields()

		for p := range pe.partList.parts {
			p.reloadLables()
		}
	})

	pe.textBufferDescription.Connect("changed", func(self *gtk.TextBuffer) {
		text, _ := self.GetText(
			self.GetStartIter(),
			self.GetEndIter(),
			true,
		)
		pe.partList.setDescriptionOfActive(text)
	})

	pe.textBufferPartData.Connect("changed", func(self *gtk.TextBuffer) {
		text, _ := self.GetText(
			self.GetStartIter(),
			self.GetEndIter(),
			true,
		)

		pe.partList.activePart.smPart.PartData = text
	})
}

func (self *partEditor) reloadFields() {
	self.entryName.SetText(
		self.partList.activePart.smPart.getTitle(currentLanguage),
	)

	self.textBufferDescription.SetText(
		self.partList.activePart.smPart.getDescription(currentLanguage),
	)

	self.entryUuid.SetText(
		self.partList.activePart.smPart.Uuid,
	)

	self.textBufferPartData.SetText(
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

	textViewDescription := getObject("textViewDescription").(*gtk.TextView)
	textBufferDescription, _ := textViewDescription.GetBuffer()

	textViewPartData := getObject("textViewPartData").(*gtk.TextView)
	textBufferPartData, _ := textViewPartData.GetBuffer()

	pe := &partEditor{
		stackPartData: getObject("stackPartData").(*gtk.Stack),

		entryName:             getObject("entryName").(*gtk.Entry),
		textViewDescription:   textViewDescription,
		textBufferDescription: textBufferDescription,

		entryUuid:          getObject("entryUuid").(*gtk.Entry),
		buttonRandomUuid:   getObject("buttonRandomUuid").(*gtk.Button),
		imageUuidStatus:    getObject("imageUuidStatus").(*gtk.Image),
		textViewPartData:   textViewPartData,
		textBufferPartData: textBufferPartData,

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
