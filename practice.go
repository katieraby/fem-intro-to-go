package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 4
	fmt.Println(reflect.TypeOf(float64(x) * 5.5))
} 

