package main

import "fmt"

func main() {
	var aValue float64 = 3.14
	//*float64 declares a variable that holds a pointer to a float64.
	var aPointer *float64 = &aValue //& operator return the address (&) can be read as the address operator, the address of the value
	fmt.Println("aPointer", aPointer)
	fmt.Println("*aPointer", *aPointer) //* takes a pointer and gets the value at the address the pointer points to
	v := 2.14
	halve(&v)
	fmt.Println("my number in 'main': ", v)

}

func halve(value *float64) {
	*value = *value / 2
	fmt.Println("my number in 'main': ", *value)

}

func First(slice []int) (int, error) {
	// YOUR CODE HERE
	if len(slice) > 0 {
		return slice[0], nil
	} else {
		return 0, fmt.Errorf("Slice is empty")
	}
}