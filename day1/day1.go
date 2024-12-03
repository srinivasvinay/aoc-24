package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
)

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

func processFile() {
	sortedLeftArray, sortedRightArray := populateArrays(readFile("day1.txt"))
	fmt.Printf("The distance is %d\n", findDistance(sortedLeftArray, sortedRightArray))
	fmt.Printf("The similarty socre is: %d\n", calculateSimilarityScore(sortedLeftArray, sortedRightArray))
}

func findDistance(leftArray, rightArray []int) int {
	sum := 0
	for index, leftArrayValue := range leftArray {
		sum += AbsValue(leftArrayValue - rightArray[index])
	}
	return sum
}
func AbsValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func populateArrays(lines []string) ([]int, []int) {
	leftArray := make([]int, 0)
	rightArray := make([]int, 0)
	for _, line := range lines {
		leftNum, rightNum := splitAndReturnInts(line)
		leftArray = append(leftArray, leftNum)
		rightArray = append(rightArray, rightNum)
	}
	sort.Ints(leftArray)
	sort.Ints(rightArray)

	return leftArray, rightArray
}
func splitAndReturnInts(line string) (int, int) {
	words := splitLineToWords(line)
	leftNum, err := strconv.Atoi(words[0])
	if err != nil {
		log.Fatalf("Encountered a non int line %s", err)
	}
	rightNum, err := strconv.Atoi(words[1])
	if err != nil {
		log.Fatalf("Encountered a non int line %s", err)
	}
	return leftNum, rightNum

}
func calculateSimilarityScore(leftSortedArray, rightSortedArray []int) int {
	simScore := 0
	rightArrayIndex := 0
	freqMap := make(map[int]int)
	for _, leftNum := range leftSortedArray {
		if _, ok := freqMap[leftNum]; !ok {
			freqMap[leftNum] = 0
		}

		freq := 0
		for rightArrayIndex < len(rightSortedArray) {
			rightNum := rightSortedArray[rightArrayIndex]
			if leftNum == rightNum {
				freq++
				freqMap[leftNum] = freq
			} else if leftNum < rightNum {
				break
			}
			rightArrayIndex++
		}
		simScore += leftNum * freqMap[leftNum]
	}
	return simScore
}
