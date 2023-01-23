package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file1, err := os.Create("one.txt")
	if err != nil {
		fmt.Println("SomeThing Went wrong!!!", err)
	}

	file2, err2 := os.Create("two.txt")
	if err2 != nil {
		fmt.Println("SomeThing Went wrong!!!", err2)
	}

	multiWriter := io.MultiWriter(os.Stdout, file1, file2)

	n, err := multiWriter.Write([]byte("WORLD!!!"))
	file1.Close()
	file2.Close()

	fmt.Println(n, err)

}

/*
1-) Ayno veriyi 2 yere birden yazmak istiyorsunuz !!! o zaman multiWriter kullanabilirisin
*/
