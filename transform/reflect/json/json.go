package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string `json:"title"`
	Age   int    `json:"time"`
}

func main() {
	myMovie := Movie{"aa", 2000}
	jsonStr, err := json.Marshal(myMovie)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", jsonStr)

	movie := Movie{}
	err = json.Unmarshal(jsonStr, &movie)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", movie)
}
