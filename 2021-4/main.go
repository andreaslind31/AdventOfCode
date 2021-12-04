package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	bingoNrs := []int{
		27, 14, 70, 7, 85, 66, 65, 57, 68, 23, 33, 78, 4, 84, 25, 18, 43, 71, 76, 61, 34, 82, 93, 74, 26, 15, 83, 64, 2, 35, 19, 97, 32, 47, 6, 51, 99, 20, 77, 75, 56, 73, 80, 86, 55, 36, 13, 95, 52, 63, 79, 72, 9, 10, 16, 8, 69, 11, 50, 54, 81, 22, 45, 1, 12, 88, 44, 17, 62, 0, 96, 94, 31, 90, 39, 92, 37, 40, 5, 98, 24, 38, 46, 21, 30, 49, 41, 87, 91, 60, 48, 29, 59, 89, 3, 42, 58, 53, 67, 28,
	}

	//open and defer close file
	filepath := "input.txt"
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	bingoCharts := ReadInts(scanner, 2500)

	boardSize := 5
	var j int
	for i := 0; i < len(bingoCharts); i += boardSize {
		j += boardSize
		if j > len(bingoCharts) {
			j = len(bingoCharts)
		}
		fmt.Println(bingoCharts[i:j])
	}
	fmt.Printf("bingo numbers: %v", bingoNrs)

}
func ReadInts(s *bufio.Scanner, total int) []int {
	s.Split(bufio.ScanWords)
	a := make([]int, total)
	for i := 0; i < total; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}
	return a
}
