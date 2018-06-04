package main

import "fmt"

func main(){
	problem1()
	problem2()
	problem3()
	problem4()
	problem5()
	problem6()
}

func problem1() {
var my_var = "New Variable"
my_second_var := "New Variable"
	fmt.Println(my_var, "and", my_second_var, "were declared differently.")
}

func problem2(){
  x := 5
  x += 1
 // should print 6
 fmt.Println(x)
}

func problem3(){
	fmt.Println(`
		Scope refers to the location of where the variable is defined. For example, if you define
		a variable in a function, the variable's scope is said to be that function. The easiest way
		to determine a variable's scope is to see where in the code it was defined, the closest program
		structure (i.e. function, if statement, etc.) is the scope for the variable.`)

}

func problem4(){
	fmt.Println(`
		The main difference between var and const is that variables declared with const cannot be
		redefined, if you try to redefine a const variable the problem will fail at compile-time.`)
}


func problem5(){
	fmt.Print("Enter a temperature in Fahrenheit: ")
	var input_temperature float64
	fmt.Scanf("%f", &input_temperature)

	output_temperature := (input_temperature - 32) * (5.0/9.0)
	fmt.Println("Your temperature in Celsius is: ", output_temperature)
}

func problem6(){
	fmt.Print("Enter a length in feet: " )
	var input_length float64
	fmt.Scanf("%f", &input_length)

	conversion_factor := 0.3048

	output_length := input_length * conversion_factor
	fmt.Println("Your length in meters is: ", output_length)
}