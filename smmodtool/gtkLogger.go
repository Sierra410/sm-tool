package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

type gtkLogger struct {
	logger     *log.Logger
	textView   *gtk.TextView
	textBuffer *gtk.TextBuffer
	grabFocus  func()
}

func (self *gtkLogger) printlnImportant(i ...interface{}) {
	self.println(i...)
	self.grabFocus()
}

func (self *gtkLogger) printfImportant(format string, i ...interface{}) {
	self.printf(format, i...)
	self.grabFocus()
}

func (self *gtkLogger) println(i ...interface{}) {
	self.printf(fmt.Sprintln(i...))
}

func (self *gtkLogger) printf(format string, i ...interface{}) {
	str := fmt.Sprintf(format, i...)

	self.logger.Printf("LOG: %s", str)

	self.textBuffer.Insert(
		self.textBuffer.GetEndIter(),
		str,
	)
}

func (self *gtkLogger) Error(err error) {
	self.printfImportant("Error: %s\n", err.Error())
}

func gtkLoggerNew(tv *gtk.TextView, f func()) *gtkLogger {
	b, _ := tv.GetBuffer()

	return &gtkLogger{
		logger:     log.New(os.Stderr, "LOG:", log.Ltime),
		textView:   tv,
		textBuffer: b,
		grabFocus:  f,
	}
}
