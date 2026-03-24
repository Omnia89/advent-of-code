package util

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func GetRawTest(day string) string {
	// open "data.txt" file in current directory
	file, err := os.Open(fmt.Sprintf(`%s/test.txt`, day))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read file
	res, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(res)
}

func GetTestByRow(day string) []string {
	// get raw data
	rawData := GetRawTest(day)

	// split raw data by row
	dataByRow := strings.Split(rawData, "\n")

	// Check if last is empty
	if dataByRow[len(dataByRow)-1] == "" {
		dataByRow = dataByRow[:len(dataByRow)-1]
	}

	return dataByRow
}

func GetRawData(day string) string {
	// open "data.txt" file in current directory
	file, err := os.Open(fmt.Sprintf(`%s/data.txt`, day))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read file
	res, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(res)
}

func GetDataByRow(day string) []string {
	// get raw data
	rawData := GetRawData(day)

	// split raw data by row
	dataByRow := strings.Split(rawData, "\n")

	// Check if last is empty
	if dataByRow[len(dataByRow)-1] == "" {
		dataByRow = dataByRow[:len(dataByRow)-1]
	}

	return dataByRow
}

func ToInt(s string) int {
	t := strings.TrimSpace(s)
	r, _ := strconv.Atoi(t)
	return r
}

func StringToIntSlice(s string, separator string) []int {
	parts := strings.Split(s, separator)
	var numbers []int
	for _, s := range parts {
		if s != "" {
			numbers = append(numbers, ToInt(s))
		}
	}
	return numbers
}

func IntContains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func IntMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
