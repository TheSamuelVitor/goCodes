package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// creating the object
type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	JsonUnmarshal()
}

func JsonMarshal() {

	// creating a new cachorro
	rex := cachorro{
		Nome:  "Rex",
		Raca:  "PitBull",
		Idade: 3,
	}

	// making the object a json
	cachorroEmJson, error := json.Marshal(rex)

	// error management
	if error != nil {
		log.Fatal(error)
	}

	// showing the json correctly
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

}

func JsonUnmarshal() {

	// creating a struct
	cachorro := cachorro{}

	// data received
	cachorroEmJson := `{"nome":"Rex","raca":"PitBull","idade":3}`

	// converting the data received in a declared struct
	// error handling 
	if error := json.Unmarshal([]byte(cachorroEmJson), &cachorro); error != nil {
		log.Fatal(error)
	}

	// showing the data
	fmt.Println(cachorro.Nome)
	fmt.Println(cachorro.Raca)
	fmt.Println(cachorro.Idade)
}
