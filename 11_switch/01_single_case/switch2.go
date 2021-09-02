package main
import "fmt"

func main() {
	fmt.Print("Masukkan bilangan: ")
	var input int
	fmt.Scanln(&input)

	switch input {
	case 10:
		fmt.Println("Nilai bilangan 10")
	case 20:
		fmt.Println("Nilai bilangan 20")
	case 30:
		fmt.Println("Nilai bilangan 30")
	case 40:
		fmt.Println("Nilai bilangan 40")
	default:
		fmt.Println("Bilangan di luar 10, 20, 30, 40")				

	}
}