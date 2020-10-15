package main

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/DanielUlises98/FoldersCompresor/flags"
	"github.com/DanielUlises98/FoldersCompresor/tracker"
)

var (
	f        = flags.InitFlags()
	allDirs  []os.FileInfo
	numbJobs int
)

// better output(cleaner information presented to the user)
// Rewrite the program so it uses chuncks of RAM
// validate if the zipfile already exists

func main() {

	defer tracker.TimeTrack(time.Now(), "Program ")

	allDirs, _ = ioutil.ReadDir(f.InPath)
	numbJobs = len(allDirs)

	job := make(chan DataPath, numbJobs)
	results := make(chan DataPath, numbJobs)

	initializeWorkers(f.NumbRoutines, job, results)
	sendJobsF(job)
	recibeAnswers(numbJobs, results)
}
