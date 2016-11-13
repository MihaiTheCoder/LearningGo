package main

import "fmt"

func main() {
	myArray := [3]int{} //Array of size 3
	myArray[0] = 42
	myArray[1] = 27
	myArray[2] = 99

	fmt.Println(myArray)

	mySArray := [...]int{42, 77, 99} //Use array iniliazer to set the initial values, size is replaced with ... because we have the values
	fmt.Println(mySArray)
	fmt.Println(len(mySArray)) //len(array) -> Get the size/length of any collection

	mySlice := myArray[:] //Put the entire array in a slice (slice is a list, you can append elements to it)
	fmt.Println(mySlice)

	mySlice = append(mySlice, 100) //append to slice

	fmt.Println(mySlice)

	a := [5]int{1, 2, 3, 4, 5} //array
	s := a[1:4]                //slice with values 2,3,4
	s1 := a[2:]                // same as a[2 : len(a)] ->3,4,5
	s2 := a[:3]                // same as a[0 : 3] ->1,2,3
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)

	mySSlice := []float32{12., 15., 18.} //Declaring a slice (because we don't have the size, or the ...)
	fmt.Println(mySSlice)

	myOSlice := make([]float32, 100) //create a slice with an initial size of 100
	fmt.Println(myOSlice)

	myMap := make(map[int]string) //map(dictionary) -> key of type int, value of type string
	fmt.Println(myMap)
	myMap[42] = "Foo"
	myMap[12] = "Bar"

	fmt.Println(myMap)
	fmt.Println(myMap[10]) //Get[missingKey] ->Return default of string -> empty string
}
