package main

// func workerT(id int, jobs <-chan int, result chan<- int) {
// 	for j := range jobs {
// 		fmt.Println("worker ", id, " Started job ", j)
// 		//time.Sleep(time.Second)
// 		fmt.Println("worker ", id, "finished job ", j)
// 		result <- j * 2
// 	}
// }
// func initializeWorker(jobs chan int, results chan int) {
// 	for i := 1; i <= 3; i++ {
// 		go workerT(i, jobs, results)
// 	}
// }
// func sendJobs(numbJobs int, jobs chan int) {
// 	for j := 1; j <= numbJobs; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)
// }
// func recibeAnswers2(numbJobs int, results chan int) {

// 	for a := 1; a <= numbJobs; a++ {
// 		<-results
// 	}
// }

func main() {

	takeInaOuth()
	job := make(chan DataPath, numbJobs)
	results := make(chan DataPath, numbJobs)

	initializeWorkers(4, job, results)
	sendJobsF(job)
	recibeAnswers(numbJobs, results)

	// const numbJobs = 100

	// jobs := make(chan int, numbJobs)
	// results := make(chan int, numbJobs)

	// initializeWorker(jobs, results)
	// sendJobs(numbJobs, jobs)
	// recibeAnswers2(numbJobs, results)
	//takeInaOuth()
}
