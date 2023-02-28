package main

import (
	"fmt"
	"time"
)

const golangBirthDate = 2007

type golangStruct struct {
	age int
}

func calculateGolangAge(time time.Time) int {
	return time.Year() - golangBirthDate
}

func (g *golangStruct) destroyStruct() {
	fmt.Println("Destroy Struct...")
	g = nil
}

func main() {

	myStr := golangStruct{
		age: calculateGolangAge(time.Now()),
	}

	fmt.Printf("beforeDestroy myStr address:%p\t myStr value:%v\n", &myStr, myStr.age)

	myStr.destroyStruct()

	fmt.Printf("afterDestroy myStr address:%p\t myStr value:%v\n", &myStr, myStr.age)
}
