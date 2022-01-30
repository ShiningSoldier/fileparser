package main

import (
	"fileparser/compare"
	"fileparser/data"
	"fmt"
)


func main() {
	choices := getChoices()

	var choice int
	fmt.Print("Please, choose what do you want to do with csv files: \n")
	for key, item := range choices {
		fmt.Printf("%d: %s \n", key, item)
	}
	_, err := fmt.Scan(&choice)
	data.CheckError(err)

	switch choice {
	case 1:
		compare.GetDifferences()
	default:
		fmt.Print("Please, enter the correct number")
	}
}

func getChoices() map[int]string {
	return map[int]string{
		1: "Compare two files",
	}
}
