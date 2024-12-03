package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	println(getDirection([]string{"1", "2", "7", "8", "9"}))

}

func getDirection(input []string) bool {
	direction := "inc"
	if input[0] > input[1] {
		direction = "dec"
	}

	last := input[0]
	for _, inp := range input[1:] {
		fmt.Println(inp)
		fmt.Println(last)
		fmt.Println(direction)
		if direction == "inc" {
			if inp < last {
				return false
			}

			last = inp
		} else {
			if inp > last {
				return false
			}
			last = inp
		}
	}
	return true
}

func readValFromFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func getAnswer() int {
	file, err := readValFromFile("./input.txt")
	sum := 0
	if err == nil {
		strFile := string(file)
		input := strings.Split(strFile, "\n")

		for _, inp := range input {
			strinp := strings.Split(inp, " ")
			res := getDirection(strinp)

			if res {
				sum += 1
			}
		}
	}
	return sum
}
