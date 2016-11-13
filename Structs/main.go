package main

import "fmt"

func main() {
	foo := myStruct{} // create struct on stack
	foo.myField = "Bar"
	fmt.Println(foo.myField)

	//Composite literal
	fooV2 := myStruct{"barus"} //Must match the order of the fields
	fmt.Println(fooV2.myField)

	//create struct on heap instead of stack
	//new doesn't allow composite literal
	//Go does some magic with Garbage collection and you don't need to dereference the pointer
	//fooV3 is a pointer-> go will detect and dereference it if needed
	fooV3 := new(myStruct)
	fooV3.myField = "Baron"
	fmt.Println(foo.myField)

	fooV4 := &myStruct{"bar"} //Alternative to new(myStruct)
	fmt.Println(fooV4.myField)
	//& and new is the same thing in this case

	cFoo := newMyCompStruct() // invoke constructor function (normal function that creates a struct)
	cFoo.myMap["Ciupa"] = "Capra"
	fmt.Println(cFoo)

	mp := messagePrinter{"foo"}
	mp.printMessage()

	emp := messagePrinter{}
	emp.message = "foo"
	emp.printMessage()
}

type myStruct struct { //declare struct with name myStruct
	myField string //one field per line, myfield of type string
}

//constructor function -> pattern "new" + name of struct
func newMyCompStruct() *myCompStruct {
	result := myCompStruct{}
	result.myMap = map[string]string{}
	return &result
}

type myCompStruct struct {
	myMap map[string]string
}

type messagePrinter struct {
	message string
}

//First method -> Method on messagePrinter struct
func (mp *messagePrinter) printMessage() {
	println(mp.message)
}

type enhancedMessagePrinter struct {
	messagePrinter
}
