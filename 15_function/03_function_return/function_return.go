package main
import "fmt"

func myFunc(x int, y int) int  {
	return x+y
}

func max(num1 int, num2 int) int {
	var result int

	if(num1 > num2){
		result = num1
	}else {
		result = num2
	}

	return result
}

func main() {
	fmt.Println(myFunc(10, 5)) 

	result := myFunc(10,5)
	fmt.Println("Result", result)

	fmt.Println(max(5, 10))
	result = max(30,30)
	fmt.Println(result)
}