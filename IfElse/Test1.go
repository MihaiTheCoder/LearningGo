package main

import (
    "strconv"
)

func main()  {
    foo := 2

    if foo == 1 {
        println("bar")
    } else {
        println("buz")
    }

    if foo2 := 2; foo2 ==1 {
     println("foo2 is 1")   
    } else {
        println("foo2 is not 1")
    }

    switch foo3 :=1; foo {
    case 1: 
        println(strconv.Itoa(foo3))        
    }

    foo4 :=1

    switch {
    case foo4 == 1:
        println("one")
    case foo > 2:
        println("two")
    }
}