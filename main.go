package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {

	p1 := Person{FirstName: "Sandy", LastName: "John"}
	p2 := Person{FirstName: "Mandy", LastName: "Peter"}

	pArr := []Person{p1, p2}

	jsn_marshl, err := json.Marshal(pArr)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(jsn_marshl))

	pArr2 := []Person{}

	err = json.Unmarshal(jsn_marshl, &pArr2)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(pArr2)

}
