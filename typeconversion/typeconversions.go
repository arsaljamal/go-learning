package main

import "fmt"

type Celcius float64
type Farenheit float64

func main() {
	var i int = 30
	var f float64 = float64(i)
	fmt.Println(f)

	var c int32 = 10
	var j int64 = int64(c)
	var k float32 = 3.14
	var l float64 = float64(k)

	fmt.Println(j)
	fmt.Println(l)

	s := "Hello"
	b := []byte(s)
	s2 := string(b)

	fmt.Println(b)
	fmt.Println(s2)

	var temCelcius Celcius = 100
	var temFarenheit Farenheit = toFarenheit(temCelcius)
	fmt.Println(temCelcius)
	fmt.Println(temFarenheit)
}

func toFarenheit(c Celcius) Farenheit {
	return Farenheit(c*9/5 + 32)
}