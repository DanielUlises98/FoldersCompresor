package main

import (
	"time"
)

// things to implement
//1.-better validation
// 2.-more(better) flags
// 3.- better output(cleaner information presented to the user)

func main() {

	defer timeTrack(time.Now(), "Program ")

	takeInaOuth()
	job := make(chan DataPath, numbJobs)
	results := make(chan DataPath, numbJobs)

	initializeWorkers(4, job, results)
	sendJobsF(job)
	recibeAnswers(numbJobs, results)
}
