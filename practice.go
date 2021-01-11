package main

import "fmt"

func main() {
	var scoresArray [5]float64
	fmt.Println(scoresArray)
	//fmt.Println(reflect.TypeOf(float64(x) * 5.5))
	scores := [5]float64{5.5, 2.3, 7, 8.8, 11.1}
	for _, value := range scores {
		fmt.Println(value)
	}
} 

