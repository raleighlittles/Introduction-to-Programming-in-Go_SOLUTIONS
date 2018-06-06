package main

import "fmt"

func main() {
	problem1()
	problem2()
	problem3()
	problem4()
	problem5()
}

func problem1() {
	fmt.Println(`Integers are stored on computers in bytes, where each byte is composed of 8 bits.`)
}

func problem2() {
	fmt.Println(`The largest 8 digit number in binary is '11111111', which is 255 in decimal.`)
}

func problem3() {
	fmt.Println(`321325 x 424521 =`, 321325*424521)
}

func problem4() {
	fmt.Println(`A string is essentially a sequence of characters. To find the length of a string, do:
		len(my_string)`)
}

func problem5() {
	var result bool
	result = (true && false) || (false && true) || !(false && false)
	// simplifies to: FALSE || FALSE || TRUE , which should be true
	fmt.Println("Expression result is: ", result)
}
