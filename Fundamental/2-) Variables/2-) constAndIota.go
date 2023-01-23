package main

import "fmt"

func main() {

	// 1 WAY --> Single constant declaration  without Initialize and with Initialize
	const myTeam = "Fenerbahçe"
	const centerCity string = "İstanbul"

	// 2.Way --> Multiple Declaration  without Initialize and with Initialize
	const (
		name string = "Mustafa"
		age         = 15
	)

	// IOTA
	const (
		optionId_1 = iota
		optionId_2
		optionId_3
	)

	fmt.Println("My Team is:", myTeam)
	fmt.Println("Center city is : ", centerCity)
	fmt.Println("Your name is : ", name)
	fmt.Println("Your age is : ", age)
	fmt.Println("My optionId1 is : ", optionId_1)
	fmt.Println("My optionId2 is : ", optionId_2)
	fmt.Println("My optionId3 is : ", optionId_3)

}

/*Some Notes By Mustafa KARACABEY:

1-) Constant values are immuatable Which means cannot change after initialize
2-) Constants can not declare without initializing == The value of a constant must be assigned when you declare it
3-) Const --> Need to know compile time
4-) var --> Need to know run time.

*/
