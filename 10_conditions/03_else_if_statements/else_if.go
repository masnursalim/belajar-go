package main
import "fmt"

func main() {
	time := 22

	if time < 10 {
		fmt.Println("Good Morning")
	}else if time < 20 {
		fmt.Println("Good Day")
	}else {
		fmt.Println("Good Evening")
	}
}