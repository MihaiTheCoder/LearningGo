package main

import (
	"strconv"
)

func main() {
	foo := 2

	if foo == 1 {
		println("bar")
	} else if foo == 2 {
		println("buz")
	} else {
		println("el patron")
	}

	if foo2 := 2; foo2 == 1 { // you can declare variable inside of if
		println("foo2 is 1")
	} else {
		println("foo2 is not 1")
	}

	switch foo3 := 1; foo { //you can declare variable inside of switch statement
	case 1:
		println(strconv.Itoa(foo3))
	}

	foo4 := 1

	switch { //You can make a switch where every condition might use a different condition with different variables
	case foo4 == 1:
		println("one")
	case foo > 2:
		println("two")
	}
}
