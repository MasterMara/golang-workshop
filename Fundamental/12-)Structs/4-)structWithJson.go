package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
	Like []Like `json:"Like"`
}

type Like struct {
	id   int
	name string
}

type Person struct {
	First string
	Last  string
}

func main() {

	myUser := user{
		Id:   5,
		Name: "Mustafa",
		Like: []Like{
			{id: 4, name: "Mustafa"},
			{id: 6, name: "Kadir"},
		},
	}

	arr, _ := json.Marshal(myUser)
	fmt.Println(string(arr))

	//1.WAY  This will marshal the JSON into []bytes
	p1 := Person{"alice", "bob"}
	bs, _ := json.Marshal(p1) //Convert Data to []bytes
	fmt.Println(string(bs))

	//2.WAY This will unmarshal the JSON from []bytes
	var p2 Person
	bs = []byte(`{"First":"alice","Last":"bob"}`)
	json.Unmarshal(bs, &p2) //Convert  Json to data
	fmt.Println(p2)

}

/*Some Notes By Mustafa KARACABEY:

1-) Marshal => String
2-) Encode => Stream


*/
