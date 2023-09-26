package main

import "fmt"

func main() {
	var emptyMap = make(map[string]int)
	fmt.Println(emptyMap)

	var mapWithData = map[string]int {
		"Alice" : 12,
		"Job" : 11,
	}
	fmt.Println(mapWithData)

	aliceAge := mapWithData["Alice"]
	fmt.Println(aliceAge)

	mapWithData["Alice"] = 13
	mapWithData["Job"] = 9
	fmt.Println(mapWithData)

	delete(mapWithData, "Alice")
	fmt.Println(mapWithData)

	aliceAge, ok := mapWithData["Alice"]
	if ok {
		fmt.Println("Found: " , ok)
	} else {
		fmt.Println("Not Found")
	}
	
	fmt.Println("Age: " ,aliceAge)

	mapWithData["nancy"] = 12
	fmt.Println(mapWithData)

	for name, age := range mapWithData {
		fmt.Printf("Name : %s and age %d\n", name, age)
	}

	fmt.Println(len(mapWithData))
}