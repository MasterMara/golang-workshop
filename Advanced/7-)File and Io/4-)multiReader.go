package main

import (
	"io"
	"os"
	"strings"
)

func main() {

	myReader1 := strings.NewReader("Hello")
	myReader2 := strings.NewReader("World")

	multiReader := io.MultiReader(myReader1, myReader2)

	io.Copy(os.Stdout, multiReader)

}

/* Some Notes By Mustafa Karacabey
1-) SomeTimes u want to merge your reader as a 1 reader
*/
