package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type dataPath struct {
	Files    []os.FileInfo
	exitPath string
	cDir     string
	fName    string
}

// takeInaOuth ... asd
func takeInaOuth() {
	cDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	// Get's the path with folders that are going to be compressed
	input := flag.String("i", "", "the path to the directory with the folders,\n current dir is the default one")

	// Get's the path where the program is going to output the compressed zip files
	output := flag.String("o", "", "Were the  zip files are going to be created,\n current dir is the default one ")
	flag.Parse()

	// Some sort of validations
	// It can be better...
	var inPath = *input
	if inPath == "." {
		inPath = cDir
	}

	var outPath = *output
	if outPath[1:2] != ":" {
		outPath = cDir + outPath
	}
	GetDirectories(inPath, outPath)
	fmt.Println("You can find your compressed files here: " + "[" + outPath + "]")
}

// GetDirectories ... asd
func GetDirectories(parentDir string, outPath string) {

	// Gets all the folders inside of the given path
	allDirs, _ := ioutil.ReadDir(parentDir)

	// Getas the name of every file in the current directory
	for _, folder := range allDirs {

		childDir := parentDir + folder.Name()

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

				// If it's a directory , takes all the files from that directory an compress
				// them into a single zip file
				dt := dataPath{
					Files:    filesInsideOf,
					exitPath: outPath,
					cDir:     childDir,
					fName:    folder.Name(),
				}
				WriteTheFiles(dt)
			}

		case mode.IsRegular():
			fmt.Println("Files without a parent directory cannot be compressed")
		}
		fmt.Println("The folder: " + GetTheNames(folder.Name()) + " was compressed SUCCESFULLY")
	}
}

// WriteTheFiles ... weasd
func WriteTheFiles(data dataPath) error {

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

	for _, fileName := range data.Files {
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

// GetTheNames ... asd
// aaa [aaa]
func GetTheNames(name string) string {
	i := strings.Index(name, "[")
	if i >= 0 {
		j := strings.Index(name[i:], "]")
		if j >= 0 {
			return name[i+1 : j+i]
		}
	}
	return ""
}

func main() {

	takeInaOuth()

}
