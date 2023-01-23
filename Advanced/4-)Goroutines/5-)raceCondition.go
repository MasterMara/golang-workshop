package main

import (
	"fmt"
	"sync"
)

func main() {

	counter := 0
	const num = 1000

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(num)

	for i := 0; i < num; i++ {

		go func() {
			mu.Lock()
			temp := counter
			temp++
			counter = temp
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter", counter)
}

/*Some Notes By Mustafa KARACABEY:

1-) Good Resource : https://www.educative.io/answers/what-are-race-conditions-in-golang
2-) Note: We cannot have a race condition if we only have a single goroutine.



*/
