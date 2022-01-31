package compare

import (
	"encoding/csv"
	"fileparser/data"
	"os"
)

import (
	"fmt"
)

func GetDifferences() {
	var firstFilePath string
	var secondFilePath string

	fmt.Print("Please, enter the path to the first file:")
	_, err := fmt.Scan(&firstFilePath)
	data.CheckError(err)
	fmt.Print("Now - the path to the second file:")
	_, err = fmt.Scan(&secondFilePath)
	data.CheckError(err)

	channel := make(chan []string)

	go data.GetFileData(firstFilePath, channel)
	go data.GetFileData(secondFilePath, channel)

	difference := getDifferenceBetweenSlices(channel)

	if len(difference) > 0 {
		createCsvWithDifferences(difference)
		fmt.Print("The file with the differences generated")
	} else {
		fmt.Print("There is no difference")
	}

}

func getDifferenceBetweenSlices(chan1 chan []string) []string {
	firstSlice, secondSlice := <-chan1, <-chan1
	mb := make(map[string]struct{}, len(secondSlice))

	for _, i := range secondSlice {
		mb[i] = struct{}{}
	}

	var diff []string

	for _, i := range firstSlice {
		if _, found := mb[i]; !found {
			diff = append(diff, i)
		}
	}

	return diff
}

func createCsvWithDifferences(diffSlice []string) {
	sec := data.GetTimestamp()

	filename := "difference" + sec + ".csv"

	csvFile, err := os.Create(filename)
	data.CheckError(err)

	csvWriter := csv.NewWriter(csvFile)

	for _, row := range diffSlice {
		err := csvWriter.Write([]string{row})
		data.CheckError(err)
	}
}
