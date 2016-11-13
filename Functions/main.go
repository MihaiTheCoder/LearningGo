package main

func main() {
	sayHello() //invoke function

	message := "Hello"
	sayHelloV2(message) //invoke function with value of message aka "Hello"

	//In sayHelloV2 we modify message variable
	//but here we still have the unmodified version
	//because message is passed as value
	println(message)

	//Pass variable address
	sayHelloV3(&message) //invoke function with the address of the message so that function may modify it

	//New message will be shown
	println(message)

	sayHelloV4("Hello", "from", "the", "other", "side") //invoke of variadic function

	println(add(1, 2, 3, 4, 5)) //invokation of function that retuns a single value

	numTerms, sum := addV2(1, 3, 5, 8) //invokation of a function that retuns multiple values

	println("Added: ", numTerms, " terms for total of ", sum)

	numTerms, sum = addV3(1, 3, 5, 9)

	println("Added: ", numTerms, " terms for total of ", sum)

	addFunc := func(terms ...int) (sum int) { //anonymous Function(lambda)
		for _, term := range terms {
			sum += term
		}
		return
	}

	println(addFunc(1, 5, 9))
}

func sayHello() { //basic functions that returns nothing, and has no parameters
	println("Hello")
}

func sayHelloV2(message string) { //Returns nothing, has one parameter of type string (passed by value)
	println(message)
	message = "El ciupi"
}

func sayHelloV3(message *string) { //Get parameter reference, so that we can update the original variable
	println(*message)
	*message = "Hello Go"
}

func sayHelloV4(messages ...string) { //Variadic func with variable number of parameters
	for _, message := range messages {
		println(message)
	}
}

func add(terms ...int) int { //variadic func that returns int
	result := 0
	for _, term := range terms {
		result += term
	}

	return result
}

func addV2(terms ...int) (int, int) { //func with multiple return values -> Returns 2 integers
	result := 0
	for _, term := range terms {
		result += term
	}
	return len(terms), result //len(terms) is the first value, result the second value returned
}

func addV3(terms ...int) (numTerms int, sum int) { //func with named return parameters
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)
	return // we already said in the signatures that numTerms and sum will be returned
}
