package main

var betaWarning = `This software is in BETA and comes with NO WARRANTY whatsoever.
BACK UP YOUR SHIT to prevent data loss.`

func main() {
	finished := gtkInit()

	logger.printlnImportant(betaWarning)

	<-finished
}
