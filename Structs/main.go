package main

import "fmt"

func main()  {
    foo := myStruct{}
    foo.myField = "Bar"
    fmt.Println(foo.myField)

    //Composite literal
    fooV2 :=myStruct{"barus"}
    fmt.Println(fooV2.myField)

    //create struct on heap instead of stack
    //new doesn't allow composite literal    
    //Go does some magic with Garbage collection and you don't need to dereference the pointer
    //fooV3 is a pointer
    fooV3 := new(myStruct)
    fooV3.myField = "Baron"
    fmt.Println(foo.myField)

    fooV4 := &myStruct{"bar"}
    fmt.Println(fooV4.myField)
    //& and new is the same thing in this case

    cFoo := newMyCompStruct()
    cFoo.myMap["Ciupa"] = "Capra"
    fmt.Println(cFoo)

    mp := messagePrinter{"foo"}
    mp.printMessage()

    emp := messagePrinter{}
    emp.message = "foo"
    emp.printMessage()
}

type myStruct struct {
    myField string
}

//constructor function pattern new + name of struct
func newMyCompStruct() *myCompStruct {
    result := myCompStruct{}
    result.myMap = map[string]string{}
    return &result;
}

type myCompStruct struct {
    myMap map[string]string
}

type messagePrinter struct {
    message string
}

func (mp *messagePrinter) printMessage() {
    println(mp.message)
}

type enhancedMessagePrinter struct {
    messagePrinter
}