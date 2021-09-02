package main
import "fmt"

func getAverage (arr []int) float32  {
	var avg float32
	var sum int
	var size = len(arr)

	for i := 0; i < size; i++ {
		sum = sum + arr[i]
	}

	avg = float32(sum/size)

	return avg
}

func main() {
	var bil = []int {1000, 500, 2, 31, 50}
	
	var avg = getAverage(bil)

	fmt.Printf("Nilai rata-rata : %.2f\n", avg)
}