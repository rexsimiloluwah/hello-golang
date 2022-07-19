package main

import (
	"fmt"
	"sync"
	"time"
)

// What is the fork-join model?
// The fork-join model is a model that Go employs in managing concurrent programs
// The child process dissociates from the main process, and must join at the end of its execution.
// The fork point is created using the go keyword, while the join point can be created using Wait Groups and Channels
func main() {
	// now := time.Now()
	// //Execute the tasks
	// go task1()
	// go task2()
	// go task3()
	// go task4()
	// fmt.Println("Tasks completed, time elapsed: ", time.Since(now))

	// Using Wait Groups
	// Create a Wait Group
	now := time.Now()
	var wg sync.WaitGroup
	// if the counter becomes 0; all goroutines blocked by wait become released
	// The wg.Done() is used to decrement the counter
	wg.Add(1)
	// Using an anonymous function to run the goroutine
	go func() {
		defer wg.Done()
		work()
	}()
	wg.Wait()
	fmt.Println("elapsed: ", time.Since(now))
	fmt.Println("Done waiting, main exits.")

	// Using channels for running the tasks concurrently
	now = time.Now()
	done := make(chan struct{})
	// Execute the tasks asynchronously
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)

	// Join point
	<-done
	<-done
	<-done
	<-done

	fmt.Println("Time elapsed: ", time.Since(now))
	fmt.Println("Main process exits.")
}

// Tasks
func task1(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 1 Completed.")
	done <- struct{}{}
}

func task2(done chan struct{}) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Task 2 Completed.")
	done <- struct{}{}
}

func task3(done chan struct{}) {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Task 3 Completed.")
	done <- struct{}{}
}

func task4(done chan struct{}) {
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Task 4 Completed.")
	done <- struct{}{}
}

// random function
func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Printing some stuff...")
}
