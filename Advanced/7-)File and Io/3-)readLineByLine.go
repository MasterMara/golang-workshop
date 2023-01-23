package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLineByLine(fileName string) {

	// Case Create File
	err := os.WriteFile(fileName, []byte("line1\nline2\nline3"), os.ModePerm)
	if err != nil {
		fmt.Println("Something Went Wrong !!!", err)
	}
	//Open File
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Something Went Wrong !!!", err)
	}

	// Create New Bufio Scanner
	scanner := bufio.NewScanner(f)

	// Scan until there is nothing to scan
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	readLineByLine("line.txt")

}

/* Some Notes By Mustafa KARACABEY
 */
