package main

func main() {
	//channel initialization
	bufferedChan := make(chan int, 5)

	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	bufferedChan <- 4
	bufferedChan <- 5
	bufferedChan <- 6 // Capacity is full and this is a deadlock case !!!

}

// Bufferred Channellar belli bir alandan sonra bloklanırlar. Deadlocka düşer
