package main
import "fmt"

func main() {
	// Mencetak 1 s/d 10
	for i:=1; i <= 10; i++ {
		fmt.Println("Nilai i :",i)
	}

	// Mencetak kelipatan 10 dari 0 s/d 100
	for i:=0; i <= 100; i+=10 {
		fmt.Println(i)
	}
}