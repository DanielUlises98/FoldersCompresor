package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type dataPath struct {
	files    []os.FileInfo
	exitPath string
	cDir     string
	fName    string
}

// takeInaOuth ... asd
func takeInaOuth() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	// Get's the path with folders that are going to be compressed
	input := flag.String("i", "", "the path to the directory with the folders,\n current dir is the default one")

	// Get's the path where the program is going to output the compressed zip f
	output := flag.String("o", "", "Were the  zip files are going to be created,\n current dir is the default one ")
	flag.Parse()

	// Some sort of validations
	// It can be better...
	var inPath = *input
	if inPath == "." {
		inPath = currentDir
	}

	var outPath = *output
	if outPath[1:2] != ":" {
		outPath = currentDir + outPath
	}
	defer fmt.Println("You can find your compressed files here: " + "[" + outPath + "]")
	getDirectories(inPath, outPath)
}

// getDirectories ... asd
func getDirectories(parentDir string, outPath string) {
	defer timeTrack(time.Now(), "getdirectories")
	// Creates a new waiting group for the goroutines
	wg := new(sync.WaitGroup)
	// Gets all the folders inside of the given path
	allDirs, _ := ioutil.ReadDir(parentDir)

	// Adds the number of current goroutines
	wg.Add(len(allDirs))

	// Getas the name of every file in the current directory
	for _, folder := range allDirs {
		go func(folderName string) {

			// It drecrements the number of goroutines by 1 after
			// the goroutine is done
			defer wg.Done()

			// Is the child folder inside the parent folder
			childDir := parentDir + folderName

			// Get the info of the current folder
			fi, err := os.Stat(childDir)
			if err != nil {
				fmt.Println(err)
			}
			//I validate if it's a file or a directory
			switch mode := fi.Mode(); {
			case mode.IsDir():
				{

					childDir := childDir + "/"

					// Get's all the files inside of the given path
					filesInsideOf, _ := ioutil.ReadDir(childDir)

					dt := dataPath{
						files:    filesInsideOf,
						exitPath: outPath,
						cDir:     childDir,
						fName:    folderName,
					}

					// If it's a directory , takes all the files from that directory an compress
					// them into a single zip file
					writeTheFiles(dt)
				}

			case mode.IsRegular():
				fmt.Println("Files without a parent directory cannot be compressed")
			}
			fmt.Println("The folder: " + getTheNames(folderName) + " was compressed SUCCESFULLY")
		}(folder.Name())
	}
	//wait for the group of goroutines to end
	wg.Wait()
}

// writeTheFiles ... weasd
func writeTheFiles(data dataPath) error {

	// Creates the file with the given name
	newZipFile, err := os.Create(data.exitPath + data.fName + ".zip")
	if err != nil {
		return err
	}
	// Needs to be open the whole process
	// When it ends the defer is called and the process is terminated
	defer newZipFile.Close()

	// takes the created file and makes into a zip file
	zipWriter := zip.NewWriter(newZipFile)
	// Needs to be open the whole process
	// When it ends the defer is called and the process is terminated
	defer zipWriter.Close()

	for _, fileName := range data.files {
		// Get's the whole file
		dat, err := ioutil.ReadFile(data.cDir + fileName.Name())
		if err != nil {
			return err
		}

		f, err := zipWriter.Create(data.fName + "/" + fileName.Name())
		if err != nil {
			return err
		}
		// Writes the file into the zip file
		_, err = f.Write(dat)
		if err != nil {
			return err
		}
	}
	return nil
}
