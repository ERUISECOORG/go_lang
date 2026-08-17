package routines

import (
	"fmt"
	"math/rand"
	"time"
	"sync" /*add sync library to use the waitgroup*/
)

/*When multiple tasks are modifying the same memory location this might cause data lose*/
/*For that, we need to add a mutex to control the paralel writing*/
var m = sync.Mutex{}
/*
It also exists the RWMutex{}
Example of usage:
	m.Rlock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()

This method will verify if the content is locked before reading
And will lock it from being modified it is beign read, this does allow paralel read
We must use this to allow multiple task to read while avoiding other task to modify what is beign read
and prevent the read task to read while the content is beign modified
*/

var wg = sync.WaitGroup{} /*This enables Lock and Unlock methods, that is a full lock*/
var dbData = []string{"id1","id2", "id3", "id4", "id5"}
var results = []string{}
func routines(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1) /*This adds 1 to the task counter*/
		go dbCall(i) 
		/*
			"go" key word makes tells the program to start the task without waiting for it to complete
			without a wait group, this will cause the program to finish without waiting for the task, so it won't do nothing
		*/
	}
	wg.Wait() /*This will verify that the task counter is 0 before continuing with the rest of the code*/
	fmt.Printf("\nTotal Execution Time: %v", time.Since(t0))
	fmt.Printf("\nTotal Results: %v", results)
}


func dbCall(i int) {
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	m.Lock() /*Everything inside this task will not act with parallelism, instead it will only have space for one task to perform the code inside the lock at the time*/
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done() /*This will decrease the counter of each task, every time each function parallel execution completes*/
}