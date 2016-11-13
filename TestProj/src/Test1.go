package main

const (
    first = "the first"
    second = "the second"
    third = "the third"
    firtNumber = 1 << (10 * iota)
    secondNumber
    thirdNumber
)

const (
    forthNumber = iota
)

func main()  {
    //println("Hello Go, is it me your looking for")
    var myInt int
    myInt = 42
    println(myInt)
    var myFloat32 float32 = 42.
    println(myFloat32)

    myString:= "Hello Go"
    println(myString);

    myComplex := complex(2, 3)
    println(myComplex)
    println(real(myComplex))
    println(imag(myComplex))

    println(firtNumber)
    println(secondNumber)
    println(thirdNumber)

    println(forthNumber)
}