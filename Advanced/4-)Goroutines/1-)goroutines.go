package main

import "fmt"

func main() {

	// 1.WAY Hello Goroutines --> Output will be nonDeterministic which means maybe u will see result or not
	go func() {
		fmt.Println("I am in Goroutine")
	}()

	fmt.Println("Hello From Main Goroutine")

}

/*Some Notes By Mustafa KARACABEY:

1-) Concurrency : we can have a multiple thread executing code if 1 thread blocks another one is picked up	and worked on
2-) Paralelism : multiple thread executed at the exact same time . REQUUİRES MULTİPLE CPU
3-) Concurency is not paralelism
4-) As a default go tries 1 cpu
5-) Goroutine is  a lightweight thread that’s managed by go runtime.
*/
