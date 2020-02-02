package main

import (
	"os/exec"
	"regexp"
	"runtime"

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

	parts      []*part
	activePart *part

	partEditor *partEditor

	modDir modDirectory
}

func (pl *partList) init() {
	pl.buttonCompile.SetOpacity(0.5)

	pl.buttonSortUp.Connect("clicked", func() {
		if pl.activePart != nil && pl.activePart.index != 0 {
			pl.swapParts(pl.activePart.index-1, pl.activePart.index)
		}
	})

	pl.buttonSortDown.Connect("clicked", func() {
		if pl.activePart != nil && pl.activePart.index != (len(pl.parts)-1) {
			pl.swapParts(pl.activePart.index+1, pl.activePart.index)
		}
	})

	pl.buttonCompile.Connect("clicked", func() func() {
		c := 0

		return func() {
			switch c {
			case 0:
				dialogInfo("Sorry", "This function is not implemented yet.")
			case 1:
				dialogInfo("Sorry", "This function is not implemented.")
			case 2:
				dialogInfo("Sorry", "This function is still not implemented.")
			case 3:
				dialogInfo("Sorry", "This function is STILL not implemented!")
			case 4:
				dialogInfo("-_-", "No. STILL not implemented.")
			case 5:
				dialogInfo("-_-", "STILL not implemented.")
			case 6:
				dialogInfo("Seriously?", "Nope.")
			case 7:
				dialogInfo("...", "You DO realize that's not how it works right?.")
			case 8:
				dialogInfo("...", "...")
			case 9:
				dialogInfo("...", "Stop it. You know, I'm a program. I can do a lot of stuff. I could install viruses on your computer.")
			case 10:
				dialogInfo("...", "...")
			case 11:
				dialogInfo("...", "Do you really want me to do this?")
			case 12:
				dialogInfo("...", "...")
			case 13:
				url := "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

				switch runtime.GOOS {
				case "linux":
					exec.Command("xdg-open", url).Start()
				case "windows":
					exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
				case "darwin":
					exec.Command("open", url).Start()
				}

				pl.buttonCompile.Hide()
			}

			c++
		}
	}())

	pl.buttonSave.Connect("clicked", func() {
		if pl.modDir.path == "" {
			pl.modDir.path = dialogDir("Select mod directory")
		}

		if pl.modDir.path == "" {
			return
		}

		err := pl.modDir.saveParts(pl.parts)
		if err != nil {
			logger.printlnImportant(err)
			return
		}

		dialogInfo("Saved", "Saved")
	})

	pl.buttonLoad.Connect("clicked", func() {
		dir := dialogDir("Select mod directory")
		pl.loadMod(dir)

	})

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
		part := pl.createNewPart(smPartNew("New part", ""))
		part.setUuid(randomUuid())
	})

	pl.searchEntryPart.Connect("search-changed", func(self *gtk.SearchEntry) {
		text, _ := self.GetText()
		pl.filterVisible(text)
	})
}

func (self *partList) loadMod(dir string) {
	if dir == "" {
		return
	}

	logger.println("Loading:", dir)

	self.modDir = modDirectory{}
	err := self.modDir.setDir(dir)
	if err != nil {
		logger.printlnImportant(err)
		return
	}

	parts := self.modDir.loadParts()
	logger.printf("Loaded %d parts\n", len(parts))

	self.clear()
	for _, p := range parts {
		self.createNewPart(p)
	}

	self.partEditor.setEditorActive(false)
	self.reloadLables()
}

func (self *partList) reloadLables() {
	for _, p := range self.parts {
		p.reloadLables()
	}
}

func (self *partList) swapParts(a, b int) {
	partA := self.parts[a]
	partB := self.parts[b]

	self.parts[b] = partA
	partA.index = b

	self.parts[a] = partB
	partB.index = a

	self.listBox.Remove(partA.listBoxRow)
	self.listBox.Insert(partA.listBoxRow, partA.index)

	self.listBox.Remove(partB.listBoxRow)
	self.listBox.Insert(partB.listBoxRow, partB.index)
}

func (self *partList) clear() {
	for _, p := range self.parts {
		p.listBoxRow.Destroy()

		p.partList = nil
		p.listBoxRow = nil
		p.smPart = nil
	}

	self.parts = []*part{}
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

	self.pushPart(part)

	self.listBox.Add(part.listBoxRow)
	part.listBoxRow.ShowAll()

	return part
}

func (self *partList) pushPart(p *part) {
	p.index = len(self.parts)
	self.parts = append(self.parts, p)
}

func (self *partList) isPartInList(p *part) bool {
	return p.index < len(self.parts) && self.parts[p.index] == p
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
