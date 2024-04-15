package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID      int
	Content string
}

var MAX_QUEUE_SIZE = 1000

var Queue = make(chan Task, MAX_QUEUE_SIZE)

// workers
func worker(id int) {
	for task := range Queue {
		processTask(task)
	}
}

func startWorkers(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go worker(i)
	}
}

func processTask(task Task) {
	fmt.Printf("Processing task with ID = %d: %s \n", task.ID, task.Content)
}

func addTaskToQueue(task Task) {
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
	time.Sleep(2 * time.Second)
	addTaskToQueue(task2)

	// Keep waiting for all go routines to end.
	time.Sleep(time.Second * 10)
}
