package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
)

type comparatorFunc func(x, y int) bool

const dir = "inputs/"

func readFile(fileName string) []string {
	lines := make([]string, 0)
	openedFile, err := os.OpenFile(path.Join(dir, fileName), os.O_RDONLY, 0)
	if err != nil {
		log.Fatalf("Error opening the file %s\n", err)
		return lines
	}
	scanr := bufio.NewScanner(openedFile)
	for scanr.Scan() {
		lines = append(lines, scanr.Text())
	}
	if err := scanr.Err(); err != nil {
		fmt.Printf("The error when reading the file is %s\n", err)
		return lines
	}
	return lines
}

func main() {
	processFile()
}
func splitLineToWords(line string) []string {
	re := regexp.MustCompile(`\s+`)
	return re.Split(line, -1)
}
func convertStringArrayToInts(words []string) []int {
	intWords := make([]int, 0)
	for _, word := range words {
		intWord, err := strconv.Atoi(word)
		if err != nil {
			log.Fatalf("Input contains non int lines %s\n", err)
		}
		intWords = append(intWords, intWord)
	}
	return intWords
}

func processFile() {
	lines := readFile("day2.txt")

	fmt.Printf("The safelines are %d\n", calculateSafeLines(lines, false))
	fmt.Printf("The safelines with tolerance are : %d\n", calculateSafeLines(lines, true))
}
func calculateSafeLines(lines []string, tolerance bool) int {
	safeLines := 0
	for _, line := range lines {
		ints := convertStringArrayToInts(splitLineToWords(line))

		strictlyMonotonic, breakPos := checkLevelSafety(ints)
		if tolerance && !strictlyMonotonic {
			if ok, removeElem := checkLevelSafetyWithSubs(ints, breakPos); ok {
				newSlice := removeElementSlice(ints, removeElem)
				strictlyMonotonic, _ = checkLevelSafety(newSlice)
			}
		}
		if strictlyMonotonic {
			safeLines++
		}
	}
	return safeLines
}
func checkLevelSafetyWithSubs(ints []int, breakPos int) (bool, int) {
	//Calc sublist to avoid looping entire list
	startPos, endPos := 0, 0
	if len(ints) < 4 {
		return false, -1
	} else if breakPos == len(ints)-2 {
		startPos = breakPos - 3
		endPos = len(ints)
	} else if breakPos == 0 {
		startPos = 0
		endPos = 4
	} else {
		startPos = breakPos - 1
		endPos = breakPos + 3
	}
	subList := ints[startPos:endPos]
	for i := 0; i <= len(subList); i++ {
		subList1 := removeElementSlice(subList, i)
		if ok, _ := checkLevelSafety(subList1); ok {
			return true, startPos + i
		}
	}
	return false, -1
}
func removeElementSlice(ints []int, elementIndex int) []int {
	newSlice := make([]int, 0)
	newSlice = append(newSlice, ints[:elementIndex]...)
	if elementIndex < len(ints) {
		newSlice = append(newSlice, ints[(elementIndex+1):]...)
	}
	return newSlice
}
func checkLevelSafety(ints []int) (bool, int) {
	_, establishCompFunc := returnComparatorFunc(ints[0], ints[1])
	strictlyMonotonic := true
	i := 0
	for ; i < len(ints)-1; i++ {
		if isSafe(ints[i], ints[i+1], establishCompFunc) {
			strictlyMonotonic = false
			break
		}
	}
	return strictlyMonotonic, i
}
func isSafe(i, j int, compFunc comparatorFunc) bool {
	diff := AbsValue(i - j)
	return !compFunc(i, j) || diff > 3 || diff < 1
}
func AbsValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

const less = "less"
const more = "more"

func returnComparatorFunc(x, y int) (string, comparatorFunc) {
	if x < y {
		return less, func(a, b int) bool {
			return a < b
		}
	}
	return more, func(a, b int) bool {
		return a > b
	}
}
