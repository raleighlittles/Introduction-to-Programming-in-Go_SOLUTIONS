package main

import "fmt"

func main() {
	problem1()
	problem2(10) // should print False, since 10/2=5 and 5 is odd
	problem3(1,2,3,4,5,6,7,8,9,10) // should print 10, max of list is 10
	nextOdd := makeOddGenerator()
	fmt.Println("Getting the next few odd numbers...")
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println("The 5th fibonnaci number is: ", problem5(5))
	problem6()
}

func problem1() {
	fmt.Println(`func sum([]int slice ()`)
}

func problem2(my_int int) {

	half_of_int := my_int / 2

	if half_of_int % 2 == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

func problem3(args ...int) {
	max := 0

	for _, value := range args {
		if value > max {
			max = value
		}
	}

	fmt.Println("Maximum is: ", max)
}

func makeOddGenerator() func() uint {
 // Generator for making odd numbers
 i := uint(1)
 return func() (ret uint) {
	 ret = i
	 i += 2
	 return
 }
}

func problem5(num int) int {
	// fibonacci recursion
	if num == 0 {
		return 0;
	} else if num == 1 {
		return 1;
	} else {
		return problem5(num-1) + problem5(num-2)
	}
	
}

func problem6() {

	fmt.Println(`defer: allows you to schedule a function to be called immediately before the current 
	function returns.`)
	fmt.Println(`panic: throws a runtime exception, stops function execution.`)
	fmt.Println(`recover: stops the panic and returns the value that was passed in the call to the panic function`)

	fmt.Println(`To recover from a runtime panic, you must include the call to recover in a defer() block`)
}