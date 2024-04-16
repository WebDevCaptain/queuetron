package main

import (
	"fmt"
	"sync"
)

type Task struct {
	ID      int
	Content string
}

var MAX_QUEUE_SIZE = 1000

var Queue = make(chan Task, MAX_QUEUE_SIZE)

var wg sync.WaitGroup = sync.WaitGroup{}

// workers
func worker(id int) {
	for task := range Queue {
		processTask(task, id)
	}
}

func startWorkers(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go worker(i)
	}
}

func processTask(task Task, workerId int) {
	fmt.Println("Using worker ID", workerId)
	fmt.Printf("Processing task with ID = %d: %s \n", task.ID, task.Content)
	wg.Done()
}

func addTaskToQueue(task Task) {
	wg.Add(1)
	Queue <- task
}

func main() {
	// background workers
	numWorkers := 3
	startWorkers(numWorkers)

	// Add dummy tasks
	task1 := Task{
		ID:      1,
		Content: "Task1 is easy",
	}
	addTaskToQueue(task1)

	task2 := Task{
		ID:      2,
		Content: "Task2 is time taking.",
	}
	// time.Sleep(2 * time.Second)
	addTaskToQueue(task2)

	for i := 0; i <= 100; i++ {
		task := Task{
			ID:      i,
			Content: fmt.Sprintf("Task%v is a regular task.", i),
		}

		addTaskToQueue(task)
	}

	// Keep waiting for all go routines to end.
	wg.Wait()

	fmt.Println("All tasks done")
}
