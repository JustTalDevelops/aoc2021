package common

import (
	"io/ioutil"
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

// InputsAsInts reads the input.txt file and returns it as a []int.
func InputsAsInts() []int {
    inputs := Inputs()
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
