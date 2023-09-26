package main

import "fmt"

func main() {
	var numbers [5]int // fixed size array
	fmt.Println(numbers)

	var numbersWithData = [5]int{1,2,3,4,5}
	fmt.Println(numbersWithData)

	var numberSlices []int
	fmt.Println(numberSlices);

	numberSlices1 := make([]int, 5) 
	fmt.Println(numberSlices1)

	existingSlice := numberSlices1[1:4]
	fmt.Println(existingSlice)

	numbersAppend := []int{1,2,3}
	numbersAppend = append(numbersAppend,4,5)
	fmt.Println(numbersAppend)

	src := []int{1,2,3}
	dst := make([]int,len(src))
	copy(dst,src)
	fmt.Println(dst)

	changeLenghtOfSlice := []int{1,2,3,4,5}
	changeLenghtOfSlice = changeLenghtOfSlice[:3:4]
	fmt.Println(changeLenghtOfSlice)

	for index,value := range changeLenghtOfSlice {
		changeLenghtOfSlice[index] = index + value + 3
		fmt.Printf("Number[%d] = %d\n", index, value)
	}
}