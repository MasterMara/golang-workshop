package composition

import "fmt"

type studentx struct {
	name    string
	surname string
}

type lesson struct { // Composition. All the small piece make up whole.
	name       string
	credential int
	students   []studentx // Like Here
}

func main() {

	var students []studentx

	mustafa := studentx{
		name:    "Mustafa",
		surname: "Karacabey",
	}
	kadir := studentx{
		name:    "Kadir",
		surname: "Karacabey",
	}

	students = append(students, mustafa, kadir)
	math := lesson{
		name:       "Math",
		credential: 6,
		students:   students,
	}

	fmt.Println(math)
}

/* some Notes By Mustafa Karacabey
-  Go Programming language supports compositon instead of inheriatence. For example Car, wheel, Steering Wheel
*/
