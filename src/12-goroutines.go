package main 

import (
	"fmt"
	"time"
	"runtime"
)

func main(){
	// To add parallelism to the concurrent execution, we need to have multiple cores
	runtime.GOMAXPROCS(4)  // This scales up the application cores to 4 to enable parallel processing 

	start := time.Now()
	// Sequential execution (executes in the same thread)
	// func(){
	// 	for i:=0; i<3; i++{
	// 		fmt.Println(i)
	// 	}
	// }()

	// func(){
	// 	for i:=0; i<3; i++{
	// 		fmt.Println(i)
	// 	}
	// }()

	// Concurrent execution (Adding the go keyword creates a new (green or goroutine) thread to execute that function)
	// The created thread is managed by the GO runtime and not the Operating system 
	go func(){
		for i:=0; i<3; i++{
			fmt.Println(i)
		}
	}()

	go func(){
		for i:=0; i<3; i++{
			fmt.Println(i)
		}
	}()
	
	time.Sleep(10*time.Millisecond)
	elapsedTime:= time.Since(start)
	fmt.Println("Total time elapsed: ", elapsedTime.String())

	//time.Sleep(time.Second)
}