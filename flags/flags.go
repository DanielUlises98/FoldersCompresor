package flags

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	goRoutines = 4
)

type flags struct {
	InPath, OutPath string
	NumbRoutines    int
}

func InitFlags() flags {

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	// Get's the path with folders that are going to be compressed
	input := flag.String("i", currentDir, "the path to the directory with the folders,\n current dir is the default one")
	// Get's the path where the program is going to output the compressed zip f
	output := flag.String("o", currentDir, "Were the  zip files are going to be created,\n current dir is the default one ")

	routines := flag.Int("gr", goRoutines, "The amount of goroutines you want the program to use")
	flag.Parse()

	iPath, err := os.Stat(*input)
	if os.IsNotExist(err) {
		log.Fatal("The path : [", *input, "] does not exist")
	}
	oPath, err := os.Stat(*output)
	if os.IsNotExist(err) && oPath == nil {
		fmt.Println("The output: [", *output, "] does not exist.\n making a new output directory")
		err := os.MkdirAll(*output, os.ModePerm)
		if err != nil {
			log.Fatal("Couldn't create the folder")
		}
	}
	if iPath.IsDir() {
		f := flags{InPath: *input,
			OutPath:      *output,
			NumbRoutines: *routines}
		//	f.allDirs, _ = ioutil.ReadDir(f.inPath)
		return f
	}
	return flags{}
	//defer fmt.Println("You can find your compressed files here: " + "[" + outPath + "]")
}
