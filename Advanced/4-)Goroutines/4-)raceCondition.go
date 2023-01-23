package main

import "time"

var sharedInt int = 0
var unusedValue int = 0

func runSimpleReader() {
	for {
		var val int = sharedInt
		if val%10 == 0 {
			unusedValue = unusedValue + 1
		}
	}
}

func runSimpleWriter() {
	for {
		sharedInt = sharedInt + 1
	}
}

// 1.WAY EXAMPLE OF RACE CONDITION
func main() {
	go runSimpleReader()
	go runSimpleWriter()
	time.Sleep(10 * time.Second)
}

/*Some Notes By Mustafa KARACABEY:
1-)  Some NOtesIf you run this code, this will not result in a crash,
2-) If you run the code with the built-in race condition checker, the go compiler will complain about the problem. go run -race .
3-) A race condition in Go occurs when two or more goroutines have shared data and interact with it simultaneously. This is best explained with the help of an example
4-) Occurs when 2 values ​​want to reach memory at the same time


*/
