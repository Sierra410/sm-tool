package main

import "github.com/gotk3/gotk3/gtk"

const (
	statusOk  = `• c •`
	statusErr = `OらO`
)

func getStatusString(b bool) string {
	if b {
		return statusOk
	}

	return statusErr
}

type partEditor struct {
	partList *partList

	stackPartData *gtk.Stack

	// Page 1
	entryName             *gtk.Entry
	textBufferDescription *gtk.TextBuffer
	// Page 2
	entryUuid        *gtk.Entry
	buttonRandomUuid *gtk.Button

	labelUuidStatus *gtk.Label

	textBufferPartData  *gtk.TextBuffer
	textBufferJsonError *gtk.TextBuffer

	comboBoxLanguage *gtk.ComboBoxText
}

func (pe *partEditor) init() {
	pe.entryName.Connect("changed", func(self *gtk.Entry) {
		text, _ := self.GetText()
		pe.partList.setTitleOfActive(text)
	})

	pe.entryUuid.Connect("changed", func(self *gtk.Entry) {
		text, _ := self.GetText()

		ok := validateUuid(text)
		if ok {
			pe.partList.setUuidOfActive(text)
		}

		pe.setUuidStatus(ok)
	})

	pe.buttonRandomUuid.Connect("clicked", func() {
		// SetUuidOfActive is called implicitly via "changed" signal
		pe.entryUuid.SetText(randomUuid())
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

		pe.partList.activePart.smPart.partData = text
		err := pe.partList.activePart.smPart.unmarshalPartData()

		pe.setJsonError(err)
	})
}

func (self *partEditor) reloadFields() {
	if self.partList.activePart == nil {
		return
	}

	self.entryName.SetText(
		self.partList.activePart.smPart.getTitle(currentLanguage),
	)

	self.textBufferDescription.SetText(
		self.partList.activePart.smPart.getDescription(currentLanguage),
	)

	self.entryUuid.SetText(
		self.partList.activePart.smPart.uuid,
	)

	self.textBufferPartData.SetText(
		self.partList.activePart.smPart.partData,
	)
}

func (self *partEditor) setEditorActive(b bool) {
	if b {
		self.stackPartData.SetVisibleChildName("1")
	} else {
		self.stackPartData.SetVisibleChildName("0")
	}
}

func (self *partEditor) setUuidStatus(ok bool) {
	self.labelUuidStatus.SetText(getStatusString(ok))
}

func (self *partEditor) setJsonError(err error) {
	var text string

	if err == nil {
		text = "No errors"
	} else {
		text = err.Error()
	}

	self.textBufferJsonError.SetText(text)
}
