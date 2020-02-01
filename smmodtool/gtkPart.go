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
	if self.smPart.setUuid(s) {
		self.reloadLables()
	}
}

func (self *part) setTitle(s string) {
	self.labelName.SetText(s)
	self.smPart.setTitle(s, currentLanguage)
}

func (self *part) destroy() {
	for _, x := range self.partList.parts[self.index+1:] {
		x.index = x.index - 1
	}

	self.partList.parts = append(
		self.partList.parts[:self.index],
		self.partList.parts[self.index+1:]...,
	)

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
