package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// Series struct which contains
// an array of series
type Series struct {
	Series []Serie `json:"series"`
}

// User struct which contains a name
// and an array of series
type User struct {
	Name   string  `json:"name"`
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

// getData func
func getData(c *gin.Context) Users {
	// Open our jsonFile
	jsonFile, err := os.Open("app/data/index.json")

	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"error": "Couldn't find any books"})
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	users := Users{}

	json.Unmarshal(byteValue, &users)

	return users
}

// getSeries func
func getSeries(c *gin.Context) Series {
	users := getData(c)
	series := Series{}

	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Name == c.Query("name") {
			return Series{users.Users[i].Series}
		}

		for j := 0; j < len(users.Users[i].Series); j++ {
			series.Series = append(series.Series, users.Users[i].Series[j])
		}
	}

	return series
}

// All func
func All(c *gin.Context) {
	series := getSeries(c)

	c.JSON(200, gin.H{"Books": series})
}

// Latest func
func Latest(c *gin.Context) {
	series := getSeries(c)

	response := Serie{}
	for i := 0; i < len(series.Series); i++ {
		end := len(series.Series[i].Books) - 1
		book := series.Series[i].Books[end]
		book.SeriesTitle = series.Series[i].Title
		response.Books = append(response.Books, book)
	}

	c.JSON(200, gin.H{"Latest Books": response.Books})
}

// ToBuy func
func ToBuy(c *gin.Context) {
	series := getSeries(c)
	response := Series{}

	for i := 0; i < len(series.Series); i++ {
		group := series.Series[i]
		group.Books = nil
		response.Series = append(response.Series, group)
		for j := 0; j < len(series.Series[i].Books); j++ {
			releaseDate, _ := time.Parse(time.RFC822, series.Series[i].Books[j].ReleaseDate)
			if !series.Series[i].Books[j].Bought && !inFuture(releaseDate) {
				book := series.Series[i].Books[j]
				response.Series[i].Books = append(response.Series[i].Books, book)
			}
		}
	}

	c.JSON(200, gin.H{"Books to buy": response})
}

// Upcoming func
func Upcoming(c *gin.Context) {
	series := getSeries(c)
	response := Series{}

	for i := 0; i < len(series.Series); i++ {
		group := series.Series[i]
		group.Books = nil
		response.Series = append(response.Series, group)
		for j := 0; j < len(series.Series[i].Books); j++ {
			releaseDate, _ := time.Parse(time.RFC822, series.Series[i].Books[j].ReleaseDate)
			if inFuture(releaseDate) {
				book := series.Series[i].Books[j]
				response.Series[i].Books = append(response.Series[i].Books, book)
			}
		}
	}

	c.JSON(200, gin.H{"Books to buy": response})
}

// inFuture func
func inFuture(check time.Time) bool {
	date := time.Now().UTC()
	return check.After(date)
}

// BookRoute func
func BookRoute(c *gin.Context) {
	series := getSeries(c)

	var serie Serie
	for i := 0; i < len(series.Series); i++ {
		var title = strings.ToLower(series.Series[i].Title)
		var titleWithHypens = strings.Replace(title, " ", "-", -1)
		var id = strings.ToLower(c.Param("id"))

		if titleWithHypens == id {
			serie = series.Series[i]
		}
	}

	fmt.Println(serie)
	if serie.Author != "" {
		c.JSON(200, gin.H{"series": serie})
	} else {
		c.JSON(200, gin.H{"error": "404 – Not found"})
	}
}
