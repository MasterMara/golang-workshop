package main

import "fmt"

func main() {

	var (
		i    = 45
		name = "Mustafa"
	)

	// 1.WAY Examples of this.
	switch {

	case i < 10:
		fmt.Println("i is less than 10")

	case i < 50:
		fmt.Println("i is less than 50")
		fallthrough

	case i > 100:
		fmt.Println("i is less than 100")
	}

	switch name {

	case "kadir":
		fmt.Println(name)

	case "Mustafa":
		fmt.Println(name)
		fallthrough
	case "zeynep": //Just run here because of fallthrough
		fmt.Println("wrong step  but still works !!!")
	case "asdd":
		fmt.Println("afşdsf")
	}

}

/*Some Notes By Mustafa KARACABEY:

1-) Third case is wrong  normally does not run. With fallthrough third case is running.
2-) after fallthrough there is no coming anything. Fallthrough is final.
3-) fallthrough will execute the body of the next case, no checking that next case for a match!, you should explicitly mention fallthrough.
4-) Fallthroug is some kind of finally in c sharp
*/
