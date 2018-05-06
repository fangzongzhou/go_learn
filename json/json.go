package main

import (
	"fmt"
	"encoding/json"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string

}

func marshal(movies []Movie)  {
	data,_:=json.Marshal(movies)
	fmt.Printf("%s\n",data)

}
func marshalIndent(movies []Movie)  {
	data,_:=json.MarshalIndent(movies,""," ")
	fmt.Printf("%s\n",data)

}
func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},

	}

	marshal(movies)
	marshalIndent(movies)




}
