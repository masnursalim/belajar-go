package main
import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
}

func main() {
	nursalim := Person{Age:30, FirstName:"Nursalim", LastName:"Al Farizi"}
	fmt.Println(nursalim)
	fmt.Println(nursalim.LastName)
}