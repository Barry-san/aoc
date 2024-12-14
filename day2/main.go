// TODO :

package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	sarr := ReadFile("input.txt")
	sum := 0

	for _, i := range sarr {
		if imporvedGetSafety(i) {
			sum = sum + 1
		}
	}
	fmt.Println(sum)

}

func ReadFile(filename string) []string {
	bs, err := os.ReadFile(filename)

	if err != nil {
		println("Error reading from file. Error :", err)
	}
	return strings.Split(string(bs), "\n")
}

func getSafety(inp string) bool {
	nums := strings.Split(inp, " ")
	prev, _ := strconv.Atoi(nums[0])

	second, _ := strconv.Atoi(nums[1])

	direction := "inc"

	if second < prev {
		direction = "dec"
	}

	for _, elem := range nums[1:] {
		current, _ := strconv.Atoi(elem)

		//when adjacent levels are equal
		if prev == current {
			return false
		}

		//decreasing

		if direction == "dec" {
			if (prev - current) > 3 {
				return false
			}

			if prev < current {
				return false

			}
		}

		if direction == "inc" {
			if (current - prev) > 3 {
				return false
			}
			if current < prev {
				return false
			}
		}
		prev = current
	}

	return true
}

func imporvedGetSafety(inp string) bool {
	ssafety := getSafety(inp)

	if ssafety {
		return ssafety
	}

	arr := strings.Split(inp, " ")

	for ind := range arr {
		newInp := slices.Concat(arr[:ind], arr[ind+1:])
		s := getSafety(strings.Join(newInp, " "))
		if s {
			return true
		}
	}

	return false
}
