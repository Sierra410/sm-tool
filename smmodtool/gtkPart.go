package main

import "github.com/gotk3/gotk3/gtk"

type part struct {
	partList   *partList
	listBoxRow *gtk.ListBoxRow
	labelName  *gtk.Label
	labelUuid  *gtk.Label
	smPart     *smPart
	index      int
}

func (self *part) setUuid(s string) {
	self.labelUuid.SetText(s)
	self.smPart.uuid = s
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
		self.smPart.uuid,
	)
}
