package main

import "fmt"

type Minutes int //Custom type with underlying type of int.
type Weight float64
type Title string
type Answer bool
type Hours int

func main() {
	minutes := Minutes(37) //Conversion from int to Minutes
	weight := Weight(64.5) //Conversion from float64 to Weight
	title := Title("Here comes the sun")
	answer := Answer(false)

	fmt.Println(minutes, weight, title, answer)
	//The underlying type will determine what intrinsic operations your custom type supports.
	//For example, if you base the minutes types on an int, it'll support all the same operations that int does, addition,
	//subtraction, etc
	minutes++
	if weight > 26.7 { //A value custom type can also be compared to a value of the same type or it can be compared to a value of the underlying type
	// like in this case
		fmt.Println("Weight exceeded")
	}
	hours := Hours(2)
	//Values of different type cannot be compared even if the have the same underlying type
	//if minutes > hours { won't compile
		fmt.Println(hours)
	//}
	hours = 1
	fmt.Println(hours)


	//Initializing a struct
	monitor1 := Monitor{"high", "hdmi", 1080}//We can initialize fields without specifying which one,
	// The fields have to appear in order and we need to specified all values that the struct has
	fmt.Println(monitor1)
	monitor2 := Monitor{Value: 24.5, Resolution: "1080p"} //We can also specified which value we want to initialize, the rest will be initialized to theirs default values
	fmt.Println(monitor2)
	monitor3 := Monitor{} // We can also left all the value be initialed by default
	fmt.Println(monitor3)
}

/* We've also created an Increment function that takes an Hours value, adds 1 to it (wrapping it around to 0 if the result
is greater than 23), and returns the result as a new Hours value */
func (h *Hours) Increment() {
	*h = (*h + 1) % 24
}
/*A struct is a group of named elements, called fields, each of which has a name and a type.*/
type Monitor struct {
	Resolution string
	Connector  string
	Value      float64
}

