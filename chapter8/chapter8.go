package main

import "fmt"

func main(){
	problem1(7)
	problem2()
	problem3()
	problem4()
	fmt.Println("x = 1, y = 2")
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("x = ", x, "y = ", y) // should now say: x=2, y = 1
}

func problem1(new_variable uint) {
	fmt.Println("The address of the variable new_variable is", &new_variable)
}

func problem2() {
	fmt.Println(`To assign a value to a pointer, use the new() operator, or assign it directly with the equals 
	sign: i.e., *my_ptr = 20; or my_ptr = new(int)`)

}

func problem3() {
	// Not sure why this question was asked after the previous question?
	fmt.Println(`To create a pointer, use the pointer operator (*),
	i.e.: *my_ptr;`)
}

func problem4() {
	x := 1.5
	square(&x)
	fmt.Println(x)
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	var temporary int
	temporary = *y
	*y = *x
	*x = temporary
}