package main

import "fmt"

func main() {
	scores := [5]float64{2.2, 1.2, 6.7, 8, 10.0}
	fmt.Println(averageScore(scores))
	fmt.Println(doesPetExist("Basil"))
	fmt.Println(addGroceriesToSlice("Tomato", "Plum", "Cheese"))
}

func averageScore(scores [5]float64) (average float64) {
	for _, value := range scores {
		average += value
	}
	return average / float64(len(scores))
}


var petNames = map[string]string {
	"Basil": "Rat",
	"Hansel": "Cat",
}

func doesPetExist(petName string) bool {
	_, pet := petNames[petName]
		return pet
}

var grocerySlice = []string{"wine", "grapes"}

func addGroceriesToSlice(groceries ...string) []string {
	food := grocerySlice
	for _, grocery := range groceries {
		food = append(food, grocery)
	}
	return food
}

