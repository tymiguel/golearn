// first line should be the name of the package
// typically this would be the name of the go file
package main

// import any package needed
import (
	"errors"
	"fmt"
	"math"
)

func mymath() {
	fmt.Println("We are going to do some math!")
	// declare the variable var and name x as a data type int and assign the value 5
	var x int = 5

	// can shorthand the assignment and go will infer the type
	y := 9

	// we can use an opperand to add the terms together
	sum := x + y

	fmt.Println("x is:", x)
	fmt.Println("y is:", y)
	fmt.Println("x + y =", sum)

	// if/then control flow for odd/even check
	if sum%2 == 0 {
		fmt.Println("Sum is even!")
	} else {
		fmt.Println("Sum is odd!")
	}

	// if/elseif control flow for size of number
	if sum < 0 {
		fmt.Println("Sum is negative!")
	} else if sum < 10 {
		fmt.Println("Sum is single digit positive")
	} else {
		fmt.Println("Sum is 10 or larger.")
	}

}

func collections() {
	// data types can be combined to make datatypes that are collections

	// arrays that are fixed size of 5 and all ints
	// if we don't assign values, the array gets initialized with all 0s
	var first_array [5]int
	first_array[2] = 7 // can index an array and update input

	// shorthand arrary with assigned values
	second_array := [5]int{5, 2, 4, 5, 6}
	second_array[0] = 100

	// slice of ints don't need to use the size and we can append new values
	non_fixed_array := []int{5, 2, 4, 5, 6}
	non_fixed_array = append(non_fixed_array, 99)

	fmt.Println("Non fixed arround with non_fixed_array:", non_fixed_array)

	// map-key values; here keys are string and values are ints
	mymap := make(map[string]int)

	mymap["first_key"] = 1
	mymap["second_key"] = 2
	mymap["third_key"] = 3

	delete(mymap, "third_key") // remove key-value from mymap

	fmt.Println("This is the map:", mymap, "; and this is the first element:", mymap["first_key"])
}

func loops() {
	// the only loop available is the for loop
	fmt.Println("These are loops!")

	// counter; stopping conditions; increment/decrement
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// this is a while loop; remove the declaration and the increment/decrement
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// iterate over an array
	arr := []int{1, 19, 28, 36, 45}
	for index, value := range arr {
		fmt.Println("Index:", index, "Value:", value)
	}

	// iterate over a map
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 20

	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}
}

func do_some_math(a int, b int) int {
	// the arguments and types need to called out
	// return value is an int
	c := a + b
	return c
}

func return_different_values(a float64) (float64, error) {
	// the two returns are a float64 and an error data type
	// go doesn't have exceptions, so we can use errors
	// we don't define the function for negative values so we return 0 and an error
	if a < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	// because we return the actual value, we return nil as the error
	return math.Sqrt(a), nil
}

// a struct is a collection of fields
// groups them into a logical grouping
type person struct {
	// name of var and the type
	name string
	age  int
}

func inc(x int) {
	x++
}

func inc_accept_pointer(x *int) {
	// the *int allows the function to accept and be passed a pointer
	// we have to prefix x with * or else it will be trying to increment the memory address
	// here we deference x with * so it gets what is at that location and allows us to modify it
	*x++
}

func pointers() {
	i := 7

	// the & sign gives us the memory address of i
	// meaning, gives us a pointer to i
	fmt.Println("original i is", i, "at location", &i)

	// here we try to increment i
	// inc gets i pass by value and copy is discarded
	// so nothing happens to i
	inc(i)
	fmt.Println("after inc(i) i is", i, "at location", &i)

	// here we pass the pointer to i and increment the original value
	inc_accept_pointer(&i)
	fmt.Println("after inc_accept_pointer(&i) i is", i, "at location", &i)
}

// create a function
func main() {
	fmt.Println("Hello, world!")

	// calls the function math()
	mymath()

	// calls collections
	collections()

	// loops
	loops()

	// functions with args
	fmt.Println("This is function that adds", 10, "and", 12, "together, to get =", do_some_math(10, 12))

	// return two values
	result, err := return_different_values(10)
	result2, err2 := return_different_values(-10)
	fmt.Println("This is the sqrt of 10:", result, err)
	fmt.Println("This is the sqrt of -10:", result2, err2)

	// here we use a struct
	// we assign the values using colons :
	p := person{name: "Jimmy", age: 10}
	fmt.Println("This is the struct:", p, ". This is", p.name, ". He is", p.age)

	// pointers
	pointers()
}
