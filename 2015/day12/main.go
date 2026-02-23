package main

import (
	"encoding/json"
	"fmt"
	"regexp"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day12")
	// data := util.GetTestByRow("day12")

	part1(data[0])
	part2(data[0])
}

func part1(data string) {
	counter := 0

	numberReg := regexp.MustCompile(`-?\d+`)

	numbers := numberReg.FindAllString(data, -1)

	for _, n := range numbers {
		counter += util.ToInt(n)
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func redFilter(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, val := range v {
			if key == "red" || val == "red" {
				return nil
			}
			v[key] = redFilter(val)
		}
		return v
	case []interface{}:
		temp := []interface{}{}
		for _, val := range v {
			temp = append(temp, redFilter(val))
		}
		return temp
	default:
		return v
	}
}

func part2(data string) {
	counter := 0

	numberReg := regexp.MustCompile(`-?\d+`)

	var jsonData interface{}

	json.Unmarshal([]byte(data), &jsonData)

	filteredData := redFilter(jsonData)

	filteredByte, _ := json.Marshal(filteredData)

	filteredString := string(filteredByte)

	numbers := numberReg.FindAllString(filteredString, -1)

	for _, n := range numbers {
		counter += util.ToInt(n)
	}

	fmt.Printf("Part 2: %d\n", counter)
}
