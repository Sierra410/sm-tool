package main

import (
	"fmt"
	"log"
)

func logWritelnImportant(format string, i ...interface{}) {
	logWriteImportant(format+"\n", i...)
}

func logWriteImportant(format string, i ...interface{}) {
	logWrite(format, i...)
	stackLog.SetVisibleChildName("1")
}

func logWriteln(format string, i ...interface{}) {
	logWrite(format+"\n", i...)
}

func logWrite(format string, i ...interface{}) {
	str := fmt.Sprintf(format, i...)

	log.Printf("LOG: %s", str)

	gtkLogTextBuffer.Insert(
		gtkLogTextBuffer.GetEndIter(),
		str,
	)
}

func logError(err error) {
	logWriteImportant("Error: %s\n", err.Error())
}
