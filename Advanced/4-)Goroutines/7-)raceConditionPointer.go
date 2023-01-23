package main

import (
	"fmt"
	"sync"
)

type exampleStruct struct {
	val int
}

func increment(myPtr *exampleStruct, wg *sync.WaitGroup) {

	myPtr.val += 1
	wg.Done()
}

func incremetFixed(myPtr *exampleStruct, wg *sync.WaitGroup) {

	lock := sync.Mutex{}

	lock.Lock()
	myPtr.val += 1
	lock.Unlock()
	wg.Done()
}

func main() {

	myPtrRaceTest := exampleStruct{}

	wg := sync.WaitGroup{}
	wg.Add(10000)

	for i := 0; i < 10000; i++ {

		//go increment(&myPtrRaceTest, &wg)
		go incremetFixed(&myPtrRaceTest, &wg)
	}

	wg.Wait()

	fmt.Println(myPtrRaceTest)
}

/*Some Notes By Mustafa KARACABEY:





 */
