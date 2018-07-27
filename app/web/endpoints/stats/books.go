package stats

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
}

// Books func
func Books(c *gin.Context) {

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

	for i := 0; i < len(series.Series); i++ {
		fmt.Println(series.Series[i].Title + " by " + series.Series[i].Author + " (" , len(series.Series[i].Books) , "books in series)")
	}

	c.JSON(200, gin.H{"series": series.Series})
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

