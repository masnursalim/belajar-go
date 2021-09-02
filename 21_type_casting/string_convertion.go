package main
import (
	"fmt"
	"strconv"
)

func main() {
	var str1 = "101"
	var str2 = "10.123"

	newInt, _ := strconv.ParseInt(str1, 0, 64)
	fmt.Println(newInt)

	newFloat, _ := strconv.ParseFloat(str2, 64)
	fmt.Println(newFloat)

}