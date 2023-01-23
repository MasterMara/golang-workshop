package main

type Calculate interface {
	Add(x, y int) int
	Subtract(x, y int) int
	Multiply(x, y int) int
	Divide(x, y int) float64
}

type Calculator struct {
}

func (c *Calculator) Add(x, y int) int {
	return x + y
}

func (c *Calculator) Subtract(x, y int) int {
	return x - y
}

func (c *Calculator) Multiply(x, y int) int {
	return x * y
}

func (c *Calculator) Divide(x, y int) int {
	return x / y
}

/* Notes By Mustafa Karacabey
1-) _test olarak naming convention yapılmaktadır.
2-) yazılan test fonksyionlarının Test ismi ile başlaması gerekmektedir.
3-)go test -cover (Unit Testing Coverage bilgisini bize verecektir.)
4-) go test -cover -coverprofile=cover.out (Bize tuhaf bir çıktı verir)
5-) go tool cover -html=cover.out -o coverage.html (Bize coverage çıktısını verecektir.)
*/
