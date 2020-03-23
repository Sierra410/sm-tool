package main

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
)

type part struct {
	partList *partList
	index    int

	listBoxRow *gtk.ListBoxRow
	labelName  *gtk.Label
	labelUuid  *gtk.Label

	smPart *smPart
}

func newPart(smp *smPart) *part {
	// Building gtk.ListBoxRow
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	box.SetHomogeneous(true)

	labelName, _ := gtk.LabelNew("")
	labelName.SetEllipsize(pango.ELLIPSIZE_END)
	box.Container.Add(labelName)

	labelUuid, _ := gtk.LabelNew("")
	box.Container.Add(labelUuid)

	listBoxRow, _ := gtk.ListBoxRowNew()
	listBoxRow.Add(box)
	listBoxRow.SetName("")

	p := &part{
		listBoxRow: listBoxRow,
		labelName:  labelName,
		labelUuid:  labelUuid,
		smPart:     smp,
	}

	listBoxRow.Connect("activate", func() {
		p.partList.partEditor.setEditorActive(true)
		p.partList.setActivePart(p)
		p.partList.partEditor.reloadFields()
	})

	return p
}

func (self *part) setUuid(s string) {
	if self.smPart.setUuid(s) {
		self.reloadLables()
	}
}

func (self *part) setTitle(s string) {
	self.labelName.SetText(s)
	self.smPart.setTitle(s, currentLanguage)
}

func (self *part) destroy() {
	delete(self.partList.parts, self.listBoxRow.Native())

	self.listBoxRow.Destroy()

	self.partList = nil
	self.listBoxRow = nil
	self.smPart = nil
}

func (self *part) reloadLables() {
	self.labelName.SetText(
		self.smPart.getTitle(currentLanguage),
	)

	self.labelUuid.SetMarkup(
		"<span font_desc=\"mono\">" + self.smPart.uuid + "</span>",
	)
}
