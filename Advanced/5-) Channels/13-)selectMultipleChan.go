package main

import "fmt"

func main() {
	chan1 := make(chan int, 1)
	chan1 <- 1

	chan2 := make(chan int, 2)
	chan2 <- 2

	////this will not print the second chan value
	//NOT DETERMENISTIC
	//select {
	//	case c1Val := <-chan1:
	//		fmt.Println(c1Val)
	//	case c2Val := <-chan2:
	//		fmt.Println(c2Val)
	//}
	//
	////this will not end
	//DEADLOCK !!!
	//for {
	//	select {
	//		case c1Val := <-chan1:
	//			fmt.Println(c1Val)
	//		case c2Val := <-chan2:
	//			fmt.Println(c2Val)
	//	}
	//}
	//
	////this will not end
	//INFINITE LOOP
	//for {
	//	select {
	//		case c1Val := <-chan1:
	//			fmt.Println(c1Val)
	//		case c2Val := <-chan2:
	//			fmt.Println(c2Val)
	//		default:
	//			break
	//	}
	//}

	var done bool
	for !done {
		//selects randomly from the channels
		//output order is non-deterministic
		select {
		case c1Val := <-chan1:
			fmt.Println(c1Val)
		case c2Val := <-chan2:
			fmt.Println(c2Val)
		default:
			done = true
		}
	}
}

/* Some Notes By Mustafa Karacabey

1-) Order does not matter at all
2-) First operaiton that is non-blocking will be choosen and if all cases are blocking then defalut case will be chosen
3-) Investigate, Fan-out, Funnel, Turn-out
4-) Investiage Spinning CAS
*/
