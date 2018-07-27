package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	Bought      bool   `json:"bought"`
	SeriesTitle string
}

// All func
func All(c *gin.Context) {
	series := getData(c)

	c.JSON(200, gin.H{"Books": series})
}

// Latest func
func Latest(c *gin.Context) {
	series := getData(c)

	response := Serie{}
	for i := 0; i < len(series.Series); i++ {
		end := len(series.Series[i].Books) - 1
		book := series.Series[i].Books[end]
		book.SeriesTitle = series.Series[i].Title
		response.Books = append(response.Books, book)
	}

	c.JSON(200, gin.H{"Latest Books": response.Books})
}

// getData func
func getData(c *gin.Context) Series {
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

	return series
}

// ToBuy func
func ToBuy(c *gin.Context) {
	series := getData(c)

	response := Series{}
	for i := 0; i < len(series.Series); i++ {
		group := series.Series[i]
		group.Books = nil
		response.Series = append(response.Series, group)
		for j := 0; j < len(series.Series[i].Books); j++ {
			if !series.Series[i].Books[j].Bought {
				book := series.Series[i].Books[j]
				response.Series[i].Books = append(response.Series[i].Books, book)
			}
		}
	}

	c.JSON(200, gin.H{"Books to buy": response})
}

func BookRoute(c *gin.Context) {
	// Open our jsonFile
	jsonFile, err := os.Open("app/data/index.json")

	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"books": "false"})
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var series Series
	json.Unmarshal(byteValue, &series)

	var serie Serie
	for i := 0; i < len(series.Series); i++ {
		var title = strings.ToLower(series.Series[i].Title)
		var titleWithHypens = strings.Replace(title, " ", "-", -1)
		var id = strings.ToLower(c.Param("id"))

		if (titleWithHypens == id) {
			serie = series.Series[i]
		}
	}

	fmt.Println(serie)
	if (serie.Author != "") {
		c.JSON(200, gin.H{"series": serie})
	} else {
		c.JSON(200, gin.H{"error":"404 – Not found"})
	}
}


