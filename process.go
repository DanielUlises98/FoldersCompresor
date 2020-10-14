package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const (
	goRoutines = 4
)

var (
	inPath, outPath string
	allDirs         []os.FileInfo
	numbJobs, nrs   int
)

//DataPath ... as
type DataPath struct {
	files    []os.FileInfo
	exitPath string
	cDir     string
	fName    string
}

// takeInaOuth ... Initialize global variables of the program
func takeInaOuth() {
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
	if os.IsNotExist(err) {
		fmt.Println("The output does not exist.\n making a new directory")
		os.MkdirAll(*output, os.ModePerm)
	} else if iPath.IsDir() && oPath.IsDir() {
		inPath = *input
		outPath = *output
		nrs = *routines

		allDirs, _ = ioutil.ReadDir(inPath)
		numbJobs = len(allDirs)
	}

	//defer fmt.Println("You can find your compressed files here: " + "[" + outPath + "]")
}

// This function sets the amount of goroutines that are going to proces the data
func initializeWorkers(nrw int, jobs chan DataPath, results chan DataPath) {
	for i := 0; i < nrw; i++ {
		go writeTheFiles(i, jobs, results)
	}
}

// Collects all the results of writeTheFiles
// and ensures that the goroutines have finished
func recibeAnswers(numbJobs int, results chan DataPath) {
	for i := 0; i < numbJobs; i++ {
		data := <-results
		fmt.Println(data.fName)
	}
}

// sendJobsF ... Sends a typeStruct to the goroutines that are running
func sendJobsF(jobs chan DataPath) {

	// Getas the name of every file in the current directory
	for _, folder := range allDirs {

		// Adds the number of current goroutines

		// It drecrements the number of goroutines by 1 after
		// the goroutine is done

		// Is the child folder inside the parent folder
		childDir := inPath + folder.Name()

		// Get the info of the current folder
		// So
		fi, err := os.Stat(childDir)
		if err != nil {
			fmt.Println(err)
		}
		//I can validate if it's a file or a directory
		switch mode := fi.Mode(); {
		case mode.IsDir():
			{
				//I add a slash to point inside the child directory
				childDir := childDir + "/"
				// Get's all the files inside of the given path
				filesInsideOf, _ := ioutil.ReadDir(childDir)

				//Builds the structure and  it sends them through a channel
				jobs <- DataPath{
					files:    filesInsideOf,
					exitPath: outPath,
					cDir:     childDir,
					fName:    folder.Name(),
				}
			}

		case mode.IsRegular():
			fmt.Println("Files without a parent directory cannot be compressed")
			numbJobs--
		}
	}
	//	Close the channel so no more values will be sent to it
	// so the recivers knows it don't longer need to wayt
	close(jobs)

}

// writeTheFiles ... weasd
//Is used by the goroutines
// Takes a struct (DataPath) and uses it to create the zip files
// while fetching the files that are inside of the child folder
func writeTheFiles(id int, jobs <-chan DataPath, results chan<- DataPath) {

	for data := range jobs {

		// Creates the zip file with the given name
		newZipFile, err := os.Create(data.exitPath + data.fName + ".zip")
		if err != nil {
			fmt.Println(err)
		}

		// Needs to be open the whole process
		// When the function ends the defer is called and the process is terminated
		defer newZipFile.Close()

		// takes the created file and makes into a zip file
		//Creates a new zip writer
		zipWriter := zip.NewWriter(newZipFile)

		// Needs to be open the whole process
		// When the function ends the defer is called and the process is terminated
		defer zipWriter.Close()

		for _, fileName := range data.files {
			// Get's the whole file in a slice of bytes
			dat, err := ioutil.ReadFile(data.cDir + fileName.Name())
			if err != nil {
				fmt.Println(err)
			}

			// creates the file inside of the zip
			f, err := zipWriter.Create(data.fName + "/" + fileName.Name())
			if err != nil {
				fmt.Println(err)
			}
			// Writes the bytes in the created file
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		}
		data.fName = strconv.Itoa(id) + ": " + data.fName
		results <- data
	}

}
