package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Firstname string
}

func main() {
	// p1 := person{
	// 	Firstname: "Jamie",
	// }

	// p2 := person{
	// 	Firstname: "Jenny",
	// }

	// // Marshal
	// xp := []person{p1, p2}

	// bs, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(string(bs))

	// // Unmarshal
	// xp2 := []person{}
	// err = json.Unmarshal(bs, &xp2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("go data structure", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)

	http.ListenAndServe(":8081", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		Firstname: "Jamie",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoder got a bad data", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decoded bad data", err)
	}
	log.Println("person: ", p1)
}
