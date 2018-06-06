package main

import "fmt"
import "math"
import "testing"

func main() {
	problem1()
}

func problem1() {
	fmt.Println(`The average() function would throw a runtime error if an empty slice were passed into it,
	since it would be attempting to divide by zero at the end. One way to get around this is to check if the length
	of the slice is 0 before starting, and if it is, simply return an empty data type, throw an exception, or even print an error message.`)
}

// The average() function is duplicated below for convenience
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

type testpair struct {
	values []int8
	result int8
}

var min_tests = []testpair{
	{[]int8{1, 2, 3, 4, 5}, 1},
	{[]int8{2, 2, 2, 2, 2}, 2},
	{[]int8{5, 4, 3, 3, 3}, 3},
}

var max_tests = []testpair{
	{[]int8{1, 2, 3, 4, 5}, 5},
	{[]int8{2, 2, 2, 2, 2}, 2},
	{[]int8{1, 2, 3, 3, 3}, 3},
}

func TestMin(t *testing.T) {
	for _, pair := range min_tests {
		value := Min(pair.values)
		if value != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"instead got", value,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range max_tests {
		value := Max(pair.values)
		if value != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"instead got", value,
			)
		}
	}
}

// The min and max functions from before
// Finds the minimum of a slice of integers
func Min(xs []int8) int8 {
	var minimum int8
	minimum = math.MaxInt8

	for _, item := range xs {
		if item < minimum {
			minimum = item
		}
	}

	return minimum
}

// Finds the maximum of a slice of integers
func Max(xs []int8) int8 {
	var maximum int8
	maximum = math.MinInt8

	for _, item := range xs {
		if item > maximum {
			maximum = item
		}
	}

	return maximum
}
