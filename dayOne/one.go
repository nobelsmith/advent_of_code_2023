package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Digit struct {
	index int
	digit string
}

func main() {

	// examples := []string{
	// 	"two1nine",
	// 	"eightwothree",
	// 	"abcone2threexyz",
	// 	"xtwone3four",
	// 	"4nineeightseven2",
	// 	"zoneight234",
	// 	"7pqrstsixteen",
	// }

	spelledDigits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	file, err := os.Open("dayOne/puzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	calibrationValues := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		// for _, line := range examples {

		digits := []Digit{}
		for k, v := range spelledDigits {
			subStr := strings.Contains(line, k)
			if subStr {
				strIdx := strings.Index(line, k)
				digit := Digit{strIdx, v}
				digits = append(digits, digit)
			}
			// fmt.Println(digits)
		}

		for i, c := range line {
			if unicode.IsDigit(c) {
				digits = append(digits, Digit{i, string(c)})
			}
		}
		sort.Slice(digits, func(i, j int) bool {
			return digits[i].index < digits[j].index
		})
		// fmt.Println("sorted digits", digits)
		calibrationValues = append(calibrationValues, digits[0].digit+digits[len(digits)-1].digit)
		fmt.Println("calibration values", digits[0].digit+digits[len(digits)-1].digit)
	}

	var calibrationValueSum int
	for _, v := range calibrationValues {
		intVar, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		calibrationValueSum += intVar
		fmt.Println(intVar, calibrationValueSum)
	}
	fmt.Println(calibrationValueSum)
}
