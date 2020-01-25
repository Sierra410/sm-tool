package main

import (
	"regexp"

	"github.com/gotk3/gotk3/gtk"
)

// Parts

type gtkSmPart struct {
	listBoxRow *gtk.ListBoxRow
	smPart     *smPart
}

type gtkSmParts struct {
	list  *gtk.ListBox
	parts []*gtkSmPart
}

func (self *gtkSmParts) new() (ok bool) {
	return true
}

func (self *gtkSmParts) remove(n int) (ok bool) {
	return true
}

func (self *gtkSmParts) filter(s string) {
	if s == "" {
		for _, part := range self.parts {
			part.listBoxRow.SetVisible(true)
		}
		return
	}

	r := regexp.MustCompile("(?i)" + regexp.QuoteMeta(s))

	for _, part := range self.parts {
		part.listBoxRow.SetVisible(part.smPart.matches(r))
	}
}
