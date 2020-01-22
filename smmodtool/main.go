package main

import (
	"errors"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	mainBuilder *gtk.Builder
	mainWindow  *gtk.Window

	buttonAdd    *gtk.Button
	buttonDelete *gtk.Button
)

var (
	errNotButton = errors.New("Object isn't button")
)

func getButton(name string) (*gtk.Button, error) {
	obj, err := mainBuilder.GetObject(name)
	if err != nil {
		return nil, err
	}

	switch obj.(type) {
	case *gtk.Button:
		return obj.(*gtk.Button), nil
	default:
		return nil, errNotButton
	}
}

func main() {
	var (
		obj glib.IObject
		err error
	)

	gtk.Init(nil)

	mainBuilder, err = gtk.BuilderNew()
	if err != nil {
		log.Fatal(err)
	}

	err = mainBuilder.AddFromString(gladeMain)
	//err = mainBuilder.AddFromFile("ui/main.glade")
	if err != nil {
		log.Fatal(err)
	}

	// Window
	obj, err = mainBuilder.GetObject("window")
	window := obj.(*gtk.Window)
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// PARTS

	// Buttons
	buttonAddPart, err := getButton("buttonAddPart")
	if err != nil {
		log.Fatal(err)
	}
	buttonDeletePart, err := getButton("buttonDeletePart")
	if err != nil {
		log.Fatal(err)
	}

	// obj, err = mainBuilder.GetObject("listParts")
	// listParts := obj.(*gtk.ListBox)
	// listParts.Connect("")

	buttonAddPart.Connect("clicked", func() {
		log.Println("buttonAddPart")
	})

	buttonDeletePart.Connect("clicked", func() {
		log.Println("buttonDeletePart")
	})

	window.ShowAll()
	gtk.Main()
}
