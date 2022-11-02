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

	jsn_res, err := json.Marshal(pArr)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(jsn_res))

}
