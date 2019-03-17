package main

import "fmt"

func main() {
	ages := map[string]float64{} //Define an empty map
	ages["Alice"] = 12 //This is a equivalent to ages := map[string]float64{"Alice": 12}
	ages["Bob"]	  = 9
	fmt.Println(ages)
	fmt.Println(ages["Alice"], ages["Bob"])

	//Another way of map initialization
	ageMap := map[string]float64{"Alice": 12, "Carol": 14, "Bob": 5}
	for name, value := range ageMap {
		fmt.Println(name, value)
	}

	for name := range ageMap { //We can ignore the value
		fmt.Println(name)
	}

}

func HalfPriceSale(prices map[string]float64) map[string]float64 {
	// YOUR CODE HERE
	halves := map[string]float64{}
	for name, value := range prices {
		halves[name] = value / 2
	}
	return halves
}