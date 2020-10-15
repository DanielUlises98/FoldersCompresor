package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

//DataPath ... as
type DataPath struct {
	files    []os.FileInfo
	exitPath string
	cDir     string
	fName    string
}

// takeInaOuth ... Initialize global variables of the program

// This function sets the amount of goroutines that are going to proces the data
func initializeWorkers(nrw int, jobs chan DataPath, results chan DataPath) {

	for i := 0; i < nrw; i++ {
		go workers(i, jobs, results)
	}
}

// Collects all the results of workers
// and ensures that the goroutines have finished
func recibeAnswers(numbJobs int, results chan DataPath) {
	for i := 0; i < numbJobs; i++ {
		data := <-results
		fmt.Println(data.fName)
	}
}

// sendJobsF ... Sends a typeStruct to the goroutines that are running
func sendJobsF(jobs chan DataPath) {

	// Iterates on each child folder
	for _, folder := range allDirs {

		// childDir Is the child folder inside the parent folder
		childDir := f.InPath + folder.Name()

		// Get's the info of the current folder
		fi, _ := isAInfo(childDir)

		//Here asks if it's a file or a folder
		switch mode := fi.Mode(); {
		case mode.IsDir():
			{
				//Adds a slash to point inside the child directory
				childDir := childDir + "/"

				// Get's all the files inside of the given path
				filesInsideOf, _ := ioutil.ReadDir(childDir)

				//Builds the structure and  it sends them through a channel
				jobs <- DataPath{
					files:    filesInsideOf,
					exitPath: f.OutPath,
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
	// so the recivers knows it don't longer need to way
	// When the channel is closed the gorroutines are now free to start
	close(jobs)
}

// workers ... weasd
//Is used by the goroutines
// Takes a struct (DataPath) and uses it to create the zip files
// while fetching the files that are inside of the child folder
func workers(id int, jobs <-chan DataPath, results chan<- DataPath) {

	for data := range jobs {
		// Returns  a new file so i can defer the close  and new ZipWriter
		zipFile, zipWriter := newZipFile(data.exitPath + data.fName + ".zip")

		// // Needs to be open the whole process
		// // When the function ends the defer is called and the process is terminated
		defer zipFile.Close()
		defer zipWriter.Close()

		for _, file := range data.files {

			writeFiles(file.Name(), data.cDir, data.fName, zipWriter)
		}
		data.fName = strconv.Itoa(id) + ": " + data.fName
		results <- data
	}
}

func isAInfo(path string) (os.FileInfo, bool) {
	fi, err := os.Stat(path)
	if err != nil {
		log.Println("there it was an error while verifying the integrity of this file: [", path, "] ")
	}
	if fi.IsDir() {
		return fi, true
	}
	if fi.Mode().IsRegular() {
		return nil, true
	}
	return nil, true
}

func newZipFile(pathName string) (*os.File, *zip.Writer) {

	// Creates the zip file with the given name
	zipFile, err := os.Create(pathName)
	if err != nil {
		log.Fatal(err, "Couldn't create a pointer to a File")
	}
	// Takes the created file and makes into a zip file
	//Creates a new zip writer
	zipWriter := zip.NewWriter(zipFile)

	return zipFile, zipWriter
}

func writeFiles(fileName, cDir, fName string, zipWriter *zip.Writer) {
	fi, isFile := isAInfo(cDir + fileName)
	// Get's the whole file in a slice of bytes
	if isFile && fi == nil {
		dat, err := ioutil.ReadFile(cDir + fileName)
		if err != nil {
			log.Println(err)
		}

		// creates the file inside of the zip
		f, err := zipWriter.Create(fName + "/" + fileName)
		if err != nil {
			log.Println(err)
		}
		// Writes the bytes in the created file
		_, err = f.Write(dat)
		if err != nil {
			log.Println(err)
		}
	}
}
