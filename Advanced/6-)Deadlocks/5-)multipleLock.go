package main

import (
	"fmt"
	"sync"
)

func main() {

	lock := sync.Mutex{}

	lock.Lock()
	lock.Lock() // Two Lock causes deadlock

	lock.Unlock()
	fmt.Println("Locked in unlocked !!!")
}

/* Some notes by Mustafa karacbaey
 */
