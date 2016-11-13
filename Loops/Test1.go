package main

func main()  {
    for i := 1; i < 5; i++ {
        println(i)
    }

    i := 0
    for{
        i++
        println(i)
        if(i >= 5) {
            break
        }
    }

    s := []string{"foo", "bar", "buz"}
    for idx, v:= range s {
        println(idx, v)
    }

    m := make(map[string]string)
    m["first"] = "foo"
    m["second"] ="bar"
    m["third"] = "buz"

    for k, v := range m {
        println(k, v)
    }
}