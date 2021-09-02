package main
import "fmt"

func main() {
	prices := [3] int {10, 20, 30}
	fmt.Println("Nilai sebelum diubah ", prices)

	// ubah nilai index 2
	prices[2] = 50
	fmt.Println("Nilai setelah diubah ",prices)
}