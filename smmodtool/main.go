package main

import (
	"flag"
	"math/rand"
	"os"
	"runtime"
	"time"
)

var betaWarning = `This software is in BETA and comes with NO WARRANTY whatsoever.
BACK UP YOUR SHIT to prevent data loss.`

func main() {
	rand.Seed(time.Now().Unix())

	var (
		loadMod = ""
	)

	flag.StringVar(&loadMod, "load", loadMod, "path to a mod directory")
	flag.Parse()

	if runtime.GOOS == "windows" {
		if os.Getenv("GTK_THEME") == "" {
			os.Setenv("GTK_THEME", "Matcha-sea")
		}
	}

	finished := gtkInit(func() {
		logger.printlnImportant(betaWarning)

		if loadMod != "" {
			mainPartList.loadMod(loadMod)
		}
	})

	<-finished
}
