package main

import "fmt"

func main() {
	problem1()
	problem2()
	problem3()

}

func problem1() {
	fmt.Println("Packages are used to help you organize your code, and because they make it easier to reuse your code.")
}

func problem2() {
	fmt.Println(`Functions with a capital letter are public, those with a lowercase letter are not private; meaning that
		other packages cannot use it.`)
}

func problem3() {
	fmt.Println(`A package alias lets you refer to a package with a shorter name.
		For example, say the full name of your package is: "golang/chapter11/samplepackage", then you can import it as:
		import sp "golang/chapter11/samplepackage"
		and now when you refer to functions inside of that package, you can just do:
		sp.<function>`)
}

// Begin problem 4 below, after Average function
// Func was taken from page 122
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Finds the minimum of a slice of integers
func Min(xs []float64) float64 {
	minimum := float64(xs[0])

	for _, item := range xs {
		if minimum < item {
			minimum = item
		}
	}

	return minimum
}

// Finds the maximum of a slice of integers
func Max(xs []float64) float64 {
	maximum := float64(xs[0])

	for _, item := range xs {
		if maximum > item {
			maximum = item
		}
	}

	return maximum
}

// The answer for problem #5 is the documentation for the two functions above.
// Sadly the godoc features don't work on CentOS Linux release 7.4.1708 (Core)
