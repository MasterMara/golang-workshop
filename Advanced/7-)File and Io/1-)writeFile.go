package main

import (
	"fmt"
	"os"
)

func writeOsWrite() {

	// Case 1 : If does not exist it wil create
	err := os.WriteFile("FirstTextFile.txt", []byte("Hello Txt World"), os.ModePerm)
	if err != nil {
		fmt.Println("Something Went Wrong !!", err)
	}
}

func writeFile() {

	// First Create
	f, err := os.Create("SecondTestFile.txt")
	if err != nil {
		fmt.Println("Something Went Wrong creatating file !!!", err)
	}

	// Then Write file
	_, err = f.Write([]byte("Second Hello World\n"))
	if err != nil {
		fmt.Println("SomeThing Went Wrong !!!", err)
	}

	f.Close()
}

func main() {

	//writeOsWrite()
	writeFile()
}

/*Mustafa KARACABEY Notes

1-) \n --> gives me line Break

*/
