package main
import (
	"os"
	"fmt"
)

func main() {
	argumentWithPath := os.Args		// return all arguments including path
	argumentSlice := os.Args[1:]	// return all elements after path
	argumentAt2 := os.Args[2]		// return specified argument only
	fmt.Println(argumentWithPath)
	fmt.Println(argumentSlice)
	fmt.Println(argumentAt2)
}