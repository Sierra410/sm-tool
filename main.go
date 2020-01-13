package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

var (
	loginfo = log.New(os.Stderr, "", 0)
	logwarn = log.New(os.Stderr, "Warning: ", 0)
	logerr  = log.New(os.Stderr, "ERROR: ", 0)
)

// Images that represend 1 channel are reffered to as "ichans" (Image Channel)
// And yes, diffuse texture is called a "channel" too despite having 3 channels by itself.

func printHelp() {
	loginfo.Printf(regexp.MustCompile(`(?m)^\t+`).ReplaceAllString(`
		Scrap Mechanic Texture Compiler
		
		Usage: %s [DIR]...
		
		If no DIR(s) is(are) provided current working directory is used instead.
		
		The program will try to find all files that follow a specific naming pattern and combine them into two files: xxxx_dif.png and xxxx_asg.png
		
		ASG group:
		  xxxx_a.png   (Alpha)
		  xxxx_s.png   (Specular)
		  xxxx_g.png   (Glow)
		  xxxx_r.png   (Reflectivity)
		DIF group: 
		  xxxx_d.png   (Diffuse)
		  xxxx_da.png  (Diffuse alpha)
		  xxxx_ao.png  (Ambient Occlusion)
		
		Note:
		  All files must be in PNG format.
		  If there there's at least one file in any of the groups both DIF and ASG textures will be generated.
		  If there are no textures in a group an empty image with default resolution (%[2]dx%[2]dpx) will be generated instead.
		  All files in the same group must have the same dimensions and the dimensions must be a power of 2!
	`, ""), os.Args[0], defaultResolutionPx)
}

func parseArgs() (dirs []string) {
	args := os.Args[1:len(os.Args)]

	dirs = make([]string, 0, len(os.Args))

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-h", "--help":
			printHelp()
			os.Exit(1)
		default:
			rp, err := filepath.Abs(args[i])
			if err != nil || !dirExists(rp) {
				loginfo.Printf(`Argument "%s" is not an option or a directory\n`, args[i])
				os.Exit(1)
			}

			dirs = append(dirs, rp)
		}
	}

	return
}

func main() {
	dirs := parseArgs()
	// User current working dir if no other dirs are provided
	if len(dirs) == 0 {
		wd, err := os.Getwd()
		if err != nil {
			logerr.Panicln(err)
		}

		dirs = append(dirs, wd)
	}

	for _, x := range dirs {
		icfs, err := scanDirForIchanFiles(x)
		if err != nil {
			logerr.Printf("Error reading \"%s\"\n", x)
			continue
		}

		for _, icf := range icfs {
			ics, err := icf.load()
			if err != nil {
				logerr.Printf("[%s]\n%s\n", icf.name, err.Error())
				continue
			}

			loginfo.Println(ics.name)

			dif := ics.compileDif()
			difPath := path.Join(ics.dir, ics.name+"_dif.png")
			err = writeImageAsPngFile(difPath, dif)
			if err != nil {
				logwarn.Printf("writing to \"%s\" failed with error\n%s\n", difPath, err.Error())
			}

			asg := ics.compileAsgr()
			asgPath := path.Join(ics.dir, ics.name+"_asg.png")
			err = writeImageAsPngFile(asgPath, asg)
			if err != nil {
				logwarn.Printf("writing to \"%s\" failed with error\n%s\n", asgPath, err.Error())
			}
		}
	}
}
