package main

import "fmt"

func main() {
	problem1()
	problem2()
	problem3()
	problem4()
}

func problem1() {
	my_arr := [7]float64 {1, 2, 3, 5, 8, 13, 21}
	fmt.Println("The 4th element in the Fibonacci sequence is", my_arr[4])
}

func problem2() {
	my_slice := make([]int, 3, 9) // 3 is the size, 9 is the max capacity
	fmt.Println("The number below should be 3.")
	fmt.Println(len(my_slice))
}

func problem3() {
	x := [6]string{"a","b", "c", "d", "e", "f"}
	fmt.Println("x[2:5] should give me the letters: c,d,e")
	fmt.Println(x[2:5]) // non-inclusive range
}

func problem4() {
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
		}

	min := x[0]

	for i := 0; i < len(x); i++ {
	   // For the love of god why is there not a built in function for this
	   if x[i] < min {
		   min = x[i]
	   }
	}

	fmt.Println("Min of x is: ", min)
}