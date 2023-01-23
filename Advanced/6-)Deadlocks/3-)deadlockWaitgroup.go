package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2) // It will cause deadlock case

	go func() {
		fmt.Println("Hello Goroutine")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Here is the Finished the main goroutine.") //Never comes here
}

/*Some Notes By Mustafa Karacabey
 */
