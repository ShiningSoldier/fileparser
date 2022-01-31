package data

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

func GetFileData(pathToFile string, channel chan []string) {
	var result []string
	file, err := os.Open(pathToFile)
	CheckError(err)

	defer func() {
		err = file.Close()
	}()
	CheckError(err)

	csvLines, err := csv.NewReader(file).ReadAll()
	CheckError(err)

	for _, line := range csvLines {
		result = append(result, line[0])
	}

	channel <- result
}

func GetTimestamp() string {
	now := time.Now()
	return strconv.FormatInt(now.Unix(), 10)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
