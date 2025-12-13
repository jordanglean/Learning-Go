package main

import (
	"fmt"
)

type floatMap map[string]float64

func (floatMap floatMap) output() {
	fmt.Println(floatMap)
}

func main() {
	website := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}

	fmt.Println(website)
	fmt.Println(website["Google"])

	website["Facebook"] = "https://facebook.com"
	fmt.Println(website)

	delete(website, "Google")
	fmt.Println(website)

	userNames := make([]string, 2, 5)
	userNames[0] = "Julie"
	userNames[1] = "Mike"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Mannuel")
	fmt.Println(userNames)

	courseRatings := make(floatMap, 5)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.2

	courseRatings.output()

	for key, item := range courseRatings {
		fmt.Println("Key: ", key)
		fmt.Println("Item: ", item)
	}
}
