package composition

import "fmt"

type myEmptyInterface interface{} // It will accept all values

func checkValueType(any interface{}) {

	fmt.Printf("Type is : %T\t value is : %v\n", any, any)
}

func sendMeAnything(any ...interface{}) {

	fmt.Println("Send me Anything")
	for _, item := range any {
		fmt.Printf("Type is : %T\t value is : %v\n", item, item)
	}

}

func sendMeAnythingAndCheckWichTypeAmI(any interface{}) {
	fmt.Println("CHECK TYPE WITH TYPE SWITCH")

	switch x := any.(type) {
	case int:
		fmt.Println("I am integer", x)
	case string:
		fmt.Println("I am string", x)
	default:
		fmt.Println("I dont know which type ı am")
	}
}

func main() {

	var (
		name  = "Mustafa"
		age   = 18
		notes = []string{"1,2,3"}
	)

	checkValueType(name)
	checkValueType(age)
	checkValueType(notes)

	sendMeAnything(name, age, notes)
	sendMeAnythingAndCheckWichTypeAmI(name)
	sendMeAnythingAndCheckWichTypeAmI(age)

}

/*Some Notes By Mustafa KARACABEY:

1-) With empty interface type can have zero or more methods. Every type in go has zero or more methods that's all in empty interface
2-) If an interface has no methods, it is called an empty interface. All types implement the empty interface
3-) all type in go is empty interface definiton because empty intetface has 0 or more values :)

*/
