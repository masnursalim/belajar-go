package main
import "fmt"

func main() {
	var (
		sum1 = 100 + 50
		sum2 = sum1 + 250
		sum3 = sum1 + sum2
	)

	fmt.Println(sum3)
}