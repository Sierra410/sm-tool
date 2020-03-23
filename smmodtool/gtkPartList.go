package main

import (
	"regexp"
	"strconv"
	"time"

	"github.com/gotk3/gotk3/gtk"
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

	parts      map[uintptr]*part
	activePart *part

	partEditor *partEditor

	modDir modDirectory
}

func (pl *partList) init() {
	pl.buttonCompile.SetOpacity(0.5)

	pl.buttonSortUp.Connect("clicked", func() {
		if pl.activePart != nil && pl.activePart.index != 0 {
			pl.swapParts(
				pl.listBox.GetRowAtIndex(pl.activePart.index-1).Native(),
				pl.activePart.listBoxRow.Native(),
			)
			pl.listBox.InvalidateSort()
		}
	})

	pl.buttonSortDown.Connect("clicked", func() {
		if pl.activePart != nil && pl.activePart.index != (len(pl.parts)-1) {
			pl.swapParts(
				pl.listBox.GetRowAtIndex(pl.activePart.index+1).Native(),
				pl.activePart.listBoxRow.Native(),
			)
			pl.listBox.InvalidateSort()
		}
	})

	// TODO
	pl.buttonCompile.Hide()

	pl.buttonSave.Connect("clicked", func() {
		if pl.modDir.path == "" {
			pl.modDir.path = dialogDir("Select mod directory")
		}

		if pl.modDir.path == "" {
			return
		}

		err := pl.modDir.saveParts(pl.getSmParts())
		if err != nil {
			logger.printlnImportant(err)
			return
		}

		dialogInfo("Saved", "Saved")
	})

	pl.buttonLoad.Connect("clicked", func() {
		dir := dialogDir("Select mod directory")
		if dir == "" {
			return
		}

		pl.setModDir(dir)
		pl.loadFromModDir()
	})

	pl.listBox.SetSortFunc(func(row1, row2 *gtk.ListBoxRow, u uintptr) int {
		return pl.parts[row1.Native()].index - pl.parts[row2.Native()].index
	}, 0)

	go func() {
		for {
			time.Sleep(time.Second / 10)
			for _, x := range pl.parts {
				x.labelName.SetText(strconv.Itoa(x.index))
			}
		}
	}()

	// It doesn't work otherwise. For some reason.
	pl.listBox.Connect("row-selected", func(self *gtk.ListBox, listBoxRow *gtk.ListBoxRow) {
		listBoxRow.Activate()
	})

	pl.buttonDeletePart.Connect("clicked", func() {
		if pl.activePart != nil && dialogYesNo("Delete?", "Do you not want to not delete this part?") {
			pl.activePart.destroy()
			pl.partEditor.setEditorActive(false)
			pl.activePart = nil
		}
	})

	pl.buttonAddPart.Connect("clicked", func() {
		smp := smPartNew("New part", "")

		smp.partDataJson = defaultPartDataJson
		smp.unmarshalPartData()

		part := pl.createNewPart(smp)
		part.setUuid(randomUuid())
	})

	pl.searchEntryPart.Connect("search-changed", func(self *gtk.SearchEntry) {
		text, _ := self.GetText()
		pl.filterVisible(text)
	})
}

func (self *partList) setModDir(dir string) error {
	logger.println("Loading:", dir)

	self.modDir = modDirectory{}
	return self.modDir.setDir(dir)
}

func (self *partList) loadFromModDir() {
	smParts := self.modDir.loadParts()
	logger.printf("Loaded %d parts\n", len(smParts))
	self.partEditor.setEditorActive(false)
	self.reloadLables()

	self.clear()
	for _, p := range smParts {
		self.createNewPart(p)
	}
}

func (self *partList) reloadLables() {
	for _, p := range self.parts {
		p.reloadLables()
	}
}

func (self *partList) clear() {
	for _, p := range self.parts {
		p.listBoxRow.Destroy()

		p.partList = nil
		p.listBoxRow = nil
		p.smPart = nil
	}

	self.parts = map[uintptr]*part{}
}

func (self *partList) createNewPart(smp *smPart) *part {
	p := newPart(smp)

	self.pushPart(p)

	return p
}

func (self *partList) pushPart(p *part) {
	p.index = len(self.parts)
	p.partList = self

	self.parts[p.listBoxRow.Native()] = p

	self.listBox.Add(p.listBoxRow)
	p.listBoxRow.ShowAll()
}

func (self *partList) getSmParts() []*smPart {
	smps := make([]*smPart, 0, len(self.parts))

	for _, x := range self.parts {
		smps = append(smps, x.smPart)
	}

	return smps

	// self.listBox.GetChildren().Foreach(func(i interface{}) {
	// 	smps = append(smps, self.parts[i.(*gtk.Widget).GObject].smPart)
	// })
}

func (self *partList) isPartInList(p *part) bool {
	if p.listBoxRow == nil {
		return false
	}
	_, ok := self.parts[p.listBoxRow.Native()]
	return ok
}

func (self *partList) filterVisible(s string) {
	if s == "" {
		self.listBox.ShowAll()
		return
	}

	r := regexp.MustCompile("(?i)" + regexp.QuoteMeta(s))

	for _, part := range self.parts {
		m := part.smPart.matches(r)
		part.listBoxRow.SetVisible(m)
	}
}

func (self *partList) setActivePart(p *part) {
	if !self.isPartInList(p) {
		return
	}

	self.activePart = p
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

func (self *partList) setKindOfActive(b bool) {
	if self.activePart != nil {
		self.activePart.smPart.kind = b
	}
}
