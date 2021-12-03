package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//open and defer close file
	filepath := "input.txt"
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	counts := make([][2]int, 26)
	length := 0
	for scanner.Scan() {
		line := scanner.Text()
		for k, v := range line {
			if v == '0' {
				counts[k][0]++
			} else {
				counts[k][1]++
			}
		}
		length = len(line)
	}

	gamma := 0
	epsilon := 0
	for i := 0; i < length; i++ {
		gamma <<= 1
		epsilon <<= 1
		//bitwise operator: <<= means Shift left by pushing zeros in from the right and let the leftmost bits fall off
		if counts[i][0] < counts[i][1] {
			gamma |= 1
		} else {
			epsilon |= 1
			//bitwise operator: |= means += for bits
		}
	}

	fmt.Println(gamma * epsilon)
}
