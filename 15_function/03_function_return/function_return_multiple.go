package main
import "fmt"

func myFunc(x int, y string) (result int, txt1 string)  {
	result = x + x
	txt1 = y + " World!"

	return
}

func swap(x, y string) (string, string){
	return y, x
}

func main() {
	fmt.Println(myFunc(10, "Hello"))

	a, b := myFunc(30, "Welcome to my ")
	fmt.Println(a, b)

	_, c := myFunc(30, "Welcome to my ")
	fmt.Println(c)

	d, _ := myFunc(30, "Welcome to my ")
	fmt.Println(d)

	x, y := swap("Nursalim", "Al Farizi")
	fmt.Println(x, y)
}