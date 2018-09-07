package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["one"] = "Good Morning"
	map1["two"] = "Bonjour"

	fmt.Println(map1)
	fmt.Println(map1["one"])

	map2 := map[string]int{"hello": 1, "world": 2}
	fmt.Println(map2)
	fmt.Println(map2["hello"])

	map1["one"] = "Good night"
	fmt.Println(map1)

	delete(map1, "one")
	fmt.Println(map1)

	val1, exists1 := map1["two"]
	fmt.Println(val1, exists1)

	val2, exists2 := map1["one"]
	fmt.Println(val2, exists2)

	rangeLoop := map[int]string{
		0: "hello",
		1: "world",
		2: "this",
		3: "is",
		4: "a",
		5: "test",
	}

	for key, val := range rangeLoop {
		fmt.Println(key, " - ", val)
	}
}
