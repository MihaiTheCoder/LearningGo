package main

import "fmt"

func main()  {
    myArray := [3]int{}
    myArray[0] = 42
    myArray[1] = 27
    myArray[2] = 99

    fmt.Println(myArray)

    mySArray := [...]int{42, 77, 99}
    fmt.Println(mySArray)
    fmt.Println(len(mySArray))

    mySlice := myArray[:]
    fmt.Println(mySlice)

    mySlice = append(mySlice, 100);

    fmt.Println(mySlice)

    mySSlice := []float32{12.,15.,18.}
    fmt.Println(mySSlice)

    myMap := make(map[int]string)
    fmt.Println(myMap)
    myMap[42] = "Foo"
    myMap[12] = "Bar"

    fmt.Println(myMap)
    fmt.Println(myMap[10])
}