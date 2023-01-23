package main

import "fmt"

func destroyx(something **int) {
	fmt.Printf("address: %p, value : %v\n", something, *something)
	something = nil
	fmt.Printf("address: %p, value : %v\n", something, *something)
}

func problem3() {
	answer := new(int)
	*answer = 42
	destroyx(&answer)
	fmt.Printf("address:%p, value:%d", answer, *answer)
}
