package main

func main() {
	for i := 1; i < 5; i++ { //classic
		println(i)
	}

	i := 0
	for { //infinite for loop
		i++
		println(i)
		if i >= 5 { //break the infinite -> Get out of loop
			break
		}
	}

	s := []string{"foo", "bar", "buz"} //slice
	for idx, v := range s {            //Foreach: range returns 2 results, index/value
		println(idx, v)
	}

	m := make(map[string]string)
	m["first"] = "foo"
	m["second"] = "bar"
	m["third"] = "buz"

	for k, v := range m { //Foreach over map, range returns key/value pair
		println(k, v)
	}
}
