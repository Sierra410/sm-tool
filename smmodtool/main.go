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

	flag.Parse()

	loadMod := flag.Arg(0)

	if runtime.GOOS == "windows" {
		if os.Getenv("GTK_THEME") == "" {
			os.Setenv("GTK_THEME", "Matcha-sea")
		}
	}

	finished := gtkInit(func() {
		logger.printlnImportant(betaWarning)

		if loadMod != "" {
			mainPartList.setModDir(loadMod)
			mainPartList.loadFromModDir()
		}
	})

	<-finished
}

func mapDeepcopy(from, to map[string]interface{}) {
	for k, v := range from {
		switch v.(type) {
		case map[string]interface{}:
			m := map[string]interface{}{}
			mapDeepcopy(v.(map[string]interface{}), m)
			to[k] = m
		default:
			to[k] = v
		}
	}
}
