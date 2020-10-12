package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	inPath, outPath string
	allDirs         []os.FileInfo
	numbJobs        int
)

//DataPath ... as
type DataPath struct {
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
	inPath = *input
	if inPath == "." {
		inPath = currentDir
	}

	outPath = *output
	if outPath[1:2] != ":" {
		outPath = currentDir + outPath
	}
	allDirs, _ = ioutil.ReadDir(inPath)
	numbJobs = len(allDirs)

	//defer fmt.Println("You can find your compressed files here: " + "[" + outPath + "]")
}
func initializeWorkers(nrw int, jobs chan DataPath, results chan DataPath) {
	for i := 0; i < nrw; i++ {
		go writeTheFiles(i, jobs, results)
	}
}

// sendJobsF ... asd
func sendJobsF(jobs chan DataPath) {

	// Creates a new waiting group for the goroutines
	// Gets all the folders inside of the given path

	//SEND JOBS
	// Getas the name of every file in the current directory
	for _, folder := range allDirs {

		// Adds the number of current goroutines

		// It drecrements the number of goroutines by 1 after
		// the goroutine is done

		// Is the child folder inside the parent folder
		childDir := inPath + folder.Name()

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

				jobs <- DataPath{
					files:    filesInsideOf,
					exitPath: outPath,
					cDir:     childDir,
					fName:    folder.Name(),
				}

				// If it's a directory , takes all the files from that directory an compress
				// them into a single zip file

				//WORKER
				//writeTheFiles(dt)
			}

		case mode.IsRegular():
			fmt.Println("Files without a parent directory cannot be compressed")
			numbJobs--
		}
		//Needs to be used in other place
		//fmt.Println(i, " The folder: "+getTheNames(folder.Name())+" was compressed SUCCESFULLY")
	}
	close(jobs)
	//wait for the group of goroutines to end

}

func recibeAnswers(numbJobs int, results chan DataPath) {
	for i := 0; i < numbJobs; i++ {
		data := <-results
		fmt.Println(data.fName)
	}
}

// writeTheFiles ... weasd
func writeTheFiles(id int, jobs <-chan DataPath, results chan<- DataPath) {

	// Creates the file with the given name

	for data := range jobs {
		newZipFile, err := os.Create(data.exitPath + data.fName + ".zip")
		if err != nil {
			fmt.Println(err)
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
				fmt.Println(err)
			}
			f, err := zipWriter.Create(data.fName + "/" + fileName.Name())
			if err != nil {
				fmt.Println(err)
			}
			// Writes the file into the zip file
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		}
		data.fName = strconv.Itoa(id) + ": " + data.fName
		results <- data
	}

}
