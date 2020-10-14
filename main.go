package main

import "time"

// things to implement
//1.-better validation almost
// 2.-more(better) flags almost

// 3.- better output(cleaner information presented to the user)
// test what if workers>jobs
// Rewrite the program so it uses chuncks of RAM

func main() {

	defer timeTrack(time.Now(), "Program ")

	takeInaOuth()
	job := make(chan DataPath, numbJobs)
	results := make(chan DataPath, numbJobs)

	initializeWorkers(nrs, job, results)
	sendJobsF(job)
	recibeAnswers(numbJobs, results)
}
