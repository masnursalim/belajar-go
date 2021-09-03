package main
import "fmt"

type Person struct {
	fname string
	lname string
}

type Employee struct {
	person Person
	empId int
}

func(p Person) details(){
	fmt.Println(p, "I am person")
}

func (e Employee) details()  {
	fmt.Println(e, "I am a employee")
}

func main() {
	p1 := Person{"Nursalim","Al Farizi"}
	p1.details()

	e1 := Employee{person:p1, empId:123}
	e1.details()
}