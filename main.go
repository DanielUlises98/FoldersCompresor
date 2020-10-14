package main

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/DanielUlises98/FoldersCompresor/flags"
)

var (
	f        = flags.InitFlags()
	allDirs  []os.FileInfo
	numbJobs int
)

// things to implement
//1.-better validation almost
// 2.-more(better) flags almost

// 3.- better output(cleaner information presented to the user)
// test what if workers>jobs
// Rewrite the program so it uses chuncks of RAM

func main() {

	defer timeTrack(time.Now(), "Program ")

	allDirs, _ = ioutil.ReadDir(f.InPath)
	numbJobs = len(allDirs)

	job := make(chan DataPath, numbJobs)
	results := make(chan DataPath, numbJobs)

	initializeWorkers(f.NumbRoutines, job, results)
	sendJobsF(job)
	recibeAnswers(numbJobs, results)
}
