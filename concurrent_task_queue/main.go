package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Task is a simple struct with an ID
type Task struct {
	ID int
}

// Worker simulates processing tasks from the jobs channel
func worker(id int, jobs <-chan Task, results chan<- string) {
	for task := range jobs {
		fmt.Printf("Worker %d: Started Task %d\n", id, task.ID)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // simulate time
		fmt.Printf("Worker %d: Finished Task %d\n", id, task.ID)
		results <- fmt.Sprintf("Task %d done by Worker %d", task.ID, id)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Task, numJobs)
	results := make(chan string, numJobs)

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send tasks
	for j := 1; j <= numJobs; j++ {
		jobs <- Task{ID: j}
	}
	close(jobs)

	// Collect results
	for r := 1; r <= numJobs; r++ {
		fmt.Println("Result:", <-results)
	}
}
