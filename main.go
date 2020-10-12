package main

import (
	"time"
)

func main() {

	defer timeTrack(time.Now(), "Program ")
	//Initialize main variables
	takeInaOuth()
	job := make(chan DataPath, numbJobs)
	results := make(chan DataPath, numbJobs)

	initializeWorkers(4, job, results)
	sendJobsF(job)
	recibeAnswers(numbJobs, results)
}
