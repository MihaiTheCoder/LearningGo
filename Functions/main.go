package main

func main() {
	sayHello()

	message := "Hello"
	sayHelloV2(message)

	//In sayHelloV2 we modify message variable
	//but here we still have the unmodified version
	//because message is passed as value
	println(message)

	//Pass variable address
	sayHelloV3(&message)

	//New message will be shown
	println(message)

	sayHelloV4("Hello", "from", "the", "other", "side")

	println(add(1, 2, 3, 4, 5))

	numTerms, sum := addV2(1, 3, 5, 8)

	println("Added: ", numTerms, " terms for total of ", sum)

	numTerms, sum = addV3(1, 3, 5, 9)

	println("Added: ", numTerms, " terms for total of ", sum)

    addFunc := func(terms ...int) (sum int) {
        for _, term := range terms {
            sum += term
        }
        return
    }

    println(addFunc(1,5,9))
}

func sayHello() {
	println("Hello")
}

func sayHelloV2(message string) {
	println(message)
	message = "El ciupi"
}

func sayHelloV3(message *string) {
	println(*message)
	*message = "Hello Go"
}

func sayHelloV4(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}

func add(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}

	return result
}

func addV2(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return len(terms), result
}

func addV3(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)
	return
}
