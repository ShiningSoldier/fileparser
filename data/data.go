package data

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

func GetFileData(pathToFile string) []string {
	var result []string
	file, err := os.Open(pathToFile)
	defer file.Close()
	CheckError(err)

	csvLines, err := csv.NewReader(file).ReadAll()
	CheckError(err)

	for _, line := range csvLines {
		result = append(result, line[0])
	}

	return result
}

func GetTimestamp() string {
	now := time.Now()
	return strconv.FormatInt(now.Unix(), 10)
}

func CheckError(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}
