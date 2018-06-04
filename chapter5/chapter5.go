package main

import "fmt"

func main() {
	problem1()
	problem2()
}

func problem1() {
	fmt.Println("The next line below should say 'Small'")
	i := 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}

func problem2() {
	fmt.Println("Beginning to print all numbers in [1,100] that are divisible by 3.")
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		} else {

		}
		
	}
}