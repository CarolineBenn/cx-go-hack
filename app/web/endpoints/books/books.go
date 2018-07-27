package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

// Series struct which contains
// an array of series
type Series struct {
	Series []Serie `json:"series"`
}

// Serie struct which contains a author
// a title and a list of books
type Serie struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Books  []Book `json:"books"`
}

// Book struct which contains a title
// a release date and position in series
type Book struct {
	Title       string `json:"title"`
	ReleaseDate string `json:"releaseDate"`
	Position    int    `json:"position"`
	SeriesTitle string
}

// Return struct which contains a list of books
type Return struct {
	Latest []Book
}

// Next func
func Latest(c *gin.Context) {

	// Open our jsonFile
	jsonFile, err := os.Open("app/data/index.json")

	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"error": "Couldn't find any books"})
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var series Series

	json.Unmarshal(byteValue, &series)

	response := Return{}
	for i := 0; i < len(series.Series); i++ {
		end := len(series.Series[i].Books) - 1
		fmt.Println("Title: ", series.Series[i].Books[end].Title)
		fmt.Println("Release Date: ", series.Series[i].Books[end].ReleaseDate)
		book := series.Series[i].Books[end]
		book.SeriesTitle = series.Series[i].Title
		response.Latest = append(response.Latest, book)
	}

	c.JSON(200, gin.H{"Latest Books": response.Latest})
}
