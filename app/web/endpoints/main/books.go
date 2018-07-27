package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Users struct which contains
// an array of users
type Series struct {
	Series []Serie `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type Serie struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Books  []Book `json:"books"`
}

// Social struct which contains a
// list of links
type Book struct {
	Title       string `json:"title"`
	ReleaseDate string `json:"releaseDate"`
	Position    int    `json:"position"`
}

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("./../../../data/index.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened index.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var series Series

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &series)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(series.Series); i++ {
		fmt.Println("Series Title: " + series.Series[i].Title)
	}

}
