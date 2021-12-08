package common

import (
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Inputs reads the input.txt file and returns it as a []string.
func Inputs() []string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(string(b), "\n")
	for i := 0; i < len(inputs); i++ {
		inputs[i] = strings.TrimSpace(inputs[i])
	}
	return inputs
}

// InputsSplit returns the input.txt file as a [][]string.
func InputsSplit(s string) [][]string {
	inputs := Inputs()
	result := make([][]string, len(inputs))
	for i := 0; i < len(inputs); i++ {
		result[i] = strings.Split(inputs[i], s)
	}
	return result
}

// InputsAsInts reads the input.txt file and returns it as a []int.
func InputsAsInts() []int {
	return StringsToInts(Inputs())
}

// CopyStrings copies the given string slice.
func CopyStrings(input []string) []string {
	return append(make([]string, 0, len(input)), input...)
}

// CopyInts copies the given int slice.
func CopyInts(input []int) []int {
	return append(make([]int, 0, len(input)), input...)
}

// Abs returns the absolute value of the given int.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Median returns the median of the given int slice.
func Median(nums []int) int {
	sort.Ints(nums)

	size := float64(len(nums))
	median := (float64(nums[int(size/2)]) + float64(nums[int(size-1)/2])) / 2
	if median != math.Trunc(median) {
		return 0
	}
	return int(median)
}

// Average returns the average of the given int slice.
func Average(nums []int) int {
	return int(math.Floor(float64(Sum(nums)) / float64(len(nums))))
}

// Sum returns the sum of the given int slice.
func Sum(input []int) int {
	var result int
	for _, i := range input {
		result += i
	}
	return result
}

// IntFromBinary converts a binary string to an integer.
func IntFromBinary(binary string) int {
	i, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

// StringsToInts converts a []string to a []int.
func StringsToInts(inputs []string) []int {
	ints := make([]int, len(inputs))
	for i, input := range inputs {
		ints[i] = Atoi(input)
	}
	return ints
}

// Atoi converts a string to an int, otherwise it panics.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
