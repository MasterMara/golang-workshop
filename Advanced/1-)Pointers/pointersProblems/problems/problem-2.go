package main

import "fmt"

func destroy(something *int) {
	fmt.Printf("address: %p, value : %d\n", something, *something)
	something = nil
	fmt.Printf("address: %p, value : %d\n", something, *something)
}

func problem2() {
	answer := new(int)
	*answer = 42
	destroy(answer)
	fmt.Printf("address:%p, value:%d", answer, *answer)
}
