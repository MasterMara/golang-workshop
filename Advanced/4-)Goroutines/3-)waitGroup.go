package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		fmt.Println("I am inside Goroutines")
		wg.Done()
	}()

	wg.Wait() // It will wait until done inside; Blocking Daha önce çalıştırılmış goroutineleri bekliyor
	fmt.Println("I Am outside goroutine")
}

/*Some Notes By Mustafa KARACABEY:



 */
