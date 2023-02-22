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

// Bufferred Channellar belli bir alandan sonra bloklanırlar. Deadlocka düşer. Buffer Channel Full!!!
// Fakat aynı anda bu kanaldan data okunursa Deadlocak düşmez
// If you try to send even if close it will cause pani
// Closing twice also cause panic
