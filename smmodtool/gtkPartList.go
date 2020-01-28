package main

import (
	"regexp"

	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
)

type partList struct {
	buttonSave    *gtk.Button
	buttonLoad    *gtk.Button
	buttonCompile *gtk.Button

	listBox          *gtk.ListBox
	buttonSortUp     *gtk.Button
	buttonSortDown   *gtk.Button
	buttonAddPart    *gtk.Button
	buttonDeletePart *gtk.Button
	searchEntryPart  *gtk.SearchEntry

	parts      map[*part]struct{}
	activePart *part

	partEditor *partEditor
}

func (pl *partList) init() {
	notImplemented := func() {
		dialogInfo("Sorry", "This function is not implemented yet.")
	}

	pl.buttonSortUp.SetOpacity(0.5)
	pl.buttonSortDown.SetOpacity(0.5)
	pl.buttonCompile.SetOpacity(0.5)

	pl.buttonSortUp.Connect("clicked", notImplemented)
	pl.buttonSortDown.Connect("clicked", notImplemented)
	pl.buttonCompile.Connect("clicked", notImplemented)

	pl.buttonSave.Connect("clicked", notImplemented)
	pl.buttonLoad.Connect("clicked", notImplemented)

	// It doesn't work otherwise. For some reason.
	pl.listBox.Connect("row-selected", func(self *gtk.ListBox, listBoxRow *gtk.ListBoxRow) {
		listBoxRow.Activate()
	})

	pl.buttonDeletePart.Connect("clicked", func() {
		if pl.activePart != nil && dialogYesNo("Delete?", "Do you not want to not delete this part?") {
			pl.activePart.destroy()
		}
	})

	pl.buttonAddPart.Connect("clicked", func() {
		part := pl.createNewPart(smPartNew(""))
		part.setUuid(randomUuid())
	})

	pl.searchEntryPart.Connect("search-changed", func(self *gtk.SearchEntry) {
		text, _ := self.GetText()
		pl.filterVisible(text)
	})
}

func (self *partList) clear() {
	for p := range self.parts {
		p.destroy()
	}
}

func (self *partList) createNewPart(smp *smPart) *part {
	// Building gtk.ListBoxRow
	listBoxRow, _ := gtk.ListBoxRowNew()
	labelName, _ := gtk.LabelNew("")
	labelUuid, _ := gtk.LabelNew("")

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
		smPart:     smp,
	}

	part.listBoxRow.Connect("activate", func() {
		self.partEditor.setEditorActive(true)
		part.partList.setActivePart(part)
		part.partList.partEditor.reloadFields()
	})

	self.parts[part] = struct{}{}

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
