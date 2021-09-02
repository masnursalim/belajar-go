package main
import "fmt"

func main() {
	// create a map
	countryMap := map[string]string{"France":"Paris", "Italy":"Rome", "Japan":"Tokyo"}

	// print map using key
	for country := range countryMap {
		fmt.Println("Capital of", country,"is", countryMap[country])
	}

	for country, capital := range countryMap {
		fmt.Println("Capital of",country,"is",capital)
	 }
}