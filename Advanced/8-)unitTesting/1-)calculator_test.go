package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	// Arrange
	number1, number2 := 3, 5
	expected := 8
	calculator := Calculator{}

	// Act
	actual := calculator.Add(number1, number2)

	// Assert
	if actual != expected {
		t.Error()
	}
}

/* Notes By Mustafa Karacabey
1-) _test olarak naming convention yapılmaktadır.
2-) yazılan test fonksyionlarının Test ismi ile başlaması gerekmektedir.
3-) go test _test.go dosyası çalıştırılabilir.
*/
