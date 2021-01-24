package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api"

// Planet is a planet type
type Planet struct {
	Name string `json:"name"`
	Terrain string `json:"terrain"`
	Population string `json:"population"`
}

// Person is a person type
type Person struct {
	Name string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld Planet
}

// AllPeople is a collection of Person types
type AllPeopleData struct {
	People []Person `json:"results"`
}

func (p *Person) getHomeworld() {
	res, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Print("Error fetching homeworld", err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("Error reading response body", err)
	}

	json.Unmarshal(bytes, &p.Homeworld)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(BaseURL + "/people/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request star wars people")
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse req body")
	}

	fmt.Println(string(bytes))

	var peopleData AllPeopleData
	if err := json.Unmarshal(bytes, &peopleData); err != nil {
		fmt.Println("Error parsing json", err)
	}

	fmt.Println(peopleData)

	for _, pers := range peopleData.People {
		pers.getHomeworld()
		fmt.Println(pers)
	}

}

func main() {
	fmt.Println(BaseURL)
	http.HandleFunc("/people/", getPeople)
	log.Fatal(http.ListenAndServe(":8040", nil))
}
