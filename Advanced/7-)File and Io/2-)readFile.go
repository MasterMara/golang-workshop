package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile(fileName string) {

	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Something Went wrong !!!", err)
	}

	fmt.Println(string(b))
}

func readFileWithBufio(fileName string) {

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Something went wrong !!!", err)
	}

	// Investigate this one Go Interface Magic
	// f has Read() method
	bufReader := bufio.NewReader(f)
	io.Copy(os.Stdout, bufReader)
}

func readFileSeek(fileName string) {

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Something went wrong !!!", err)
	}
	f.Seek(5, 1)
	readByte := make([]byte, 4)
	f.Read(readByte)

	fmt.Println(string(readByte))
}

func main() {

	readFile("FirstTestFile.txt")
	readFileWithBufio("SecondTestFile.txt")
	readFileSeek("FirstTestFile.txt")

}

/*Mustafa KARACABEY Notes

1-) Buffering : https://cloudbunny.net/buffer-teknolojisi-nedir-ne-ise-yarar-300.html

2-) pc'lerin gelen verilerin işleneceği kadar tutulabileceği yeterli ara bellek alanının oluşturulması hadisesi seklinde kısaca tanımlayabiliriz..
bu noktada göndereci  insanımızın veri aktarım hızını azaltmak gibi sabote edici bir girişimde bulunulmaz buffering eyleminde..aslında, ara belleğe almak,
verinin ulaşma hızındaki değişimleri en sorunsuz şekliyle halletmek için yaygın olarak kullanılan bir yöntemdir..bunun zaten kullanılmakta olan bır mantık oldugunu söylemekte mumkundur..
3-) when doint io operations you want to buffer some data

*/
