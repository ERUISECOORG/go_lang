package channels

import (
	"fmt"
	"sync"
	"time"
)

func multsends() {
	var c = make(chan int, 5)
	var wg sync.WaitGroup

	//We send 3 routines
	for w := 1; w <= 3; w++ {
		wg.Add(1) // We increment the wait group before sending the routine
		go function(c, &wg, w)
	}

	// this go routine will close the channel once the wait group is done waiting for the routines
	// it is in his own function so the wait only blocks this function but it doesn't blocks the rest of the main code.
	go func() {
		wg.Wait()
		close(c)
	}()

	// The range finishes thanks to the close in the routine
	// It doesn't matter if the channel is closed before the reading is complete
	// You can read from the channel until it is empty, it just closes from writing
	// You can't open a channel again once it is closed
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}

	fmt.Println("All workers done")
}

func function(c chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done() // This will clear the routine from the wait group whatever happens during the execution

	for i := range 5 {
		c <- i * id
	}

	fmt.Println("Exiting Process", id)
}