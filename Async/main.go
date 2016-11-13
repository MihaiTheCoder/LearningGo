package main

import (
    "runtime"
)

func main()  {
    runtime.GOMAXPROCS(4)

    ch :=make(chan string)
    doneCh := make(chan bool)

    go abcGen(ch)
    
    go printer(ch, doneCh)
    println("Le ciupi")
    
    <-doneCh
}

func abcGen(ch chan string) {
    for i := byte('a'); i <= byte('z'); i ++ {
        ch <- string(i)
    }
    
    close(ch)
}

func printer(ch chan string, doneCh chan bool) {
    for c := range ch{
        println(c)
    }

    doneCh <- true
}