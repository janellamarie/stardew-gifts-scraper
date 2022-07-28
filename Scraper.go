/* Stardew Valley Wiki: https://stardewvalleywiki.com/List_of_All_Gifts
 * Tutorials: https://oxylabs.io/blog/golang-web-scraper, -https://blog.logrocket.com/web-scraping-with-go-and-colly/, https://blog.logrocket.com/building-web-scraper-go-colly/
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

// Villager struct to contain information
type Villager struct {
	Name     string   `json:"name"`
	Birthday string   `json:"birthday"`
	Loves    []string `json:"loves"`
	Likes    []string `json:"likes"`
	Neutral  []string `json:"neutral"`
	Dislikes []string `json:"dislikes"`
	Hates    []string `json:"hates"`
}

var villagers map[string]Villager // map of Villagers

func scrapeGifts() {
	url := "https://stardewvalleywiki.com/List_of_All_Gifts" // URL to scrape

	villagers = make(map[string]Villager) // initialize the map

	c := colly.NewCollector() // a collector makes HTTP request + traverses HTML pages

	// colly.AllowedDomains was making my code not work -> why?

	// runs every time Colly makes a request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Currently visiting: ", r.URL.String())
	})

	// runs every time Colly receives a response
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response: ", r.StatusCode)
	})

	// runs whenever Colly encounters an error while making the request
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	// parse the HTML elements
	c.OnHTML("table.wikitable", func(element *colly.HTMLElement) {
		element.ForEach("tr", func(i int, row *colly.HTMLElement) {
			name := ""
			temp := Villager{} // create Villager struct
			row.ForEach("td", func(j int, data *colly.HTMLElement) {
				switch j {
				case 0: // name
					data.Text = data.Text[:len(data.Text)-1] // slice the string by "extending" the length of the slice
					name = data.Text
					temp.Name = name
				case 1: // birthday
					temp.Birthday = data.Text[:len(data.Text)-1]
				case 2: // loves
					temp.Loves = parseList(data.Text[:len(data.Text)-1])
				case 3: // likes
					temp.Likes = parseList(data.Text[:len(data.Text)-1])
				case 4: // neutral
					temp.Neutral = parseList(data.Text[:len(data.Text)-1])
				case 5: // dislikes
					temp.Dislikes = parseList(data.Text[:len(data.Text)-1])
				case 6: // hates
					temp.Hates = parseList(data.Text[:len(data.Text)-1])
				}
			})
			villagers[name] = temp
		})
	})

	// go to website
	c.Visit(url) // starts the scraper

	displayAllVillagers(villagers)
	exportToJSON(villagers) // export the villagers map into a JSON file
}

func displayAllVillagers(villagers map[string]Villager) {
	for i, villager := range villagers {
		fmt.Println(i, ": ", villager.Birthday)
		fmt.Println("loves: ", villager.Loves, "total: ", len(villager.Loves))
		fmt.Println("likes: ", villager.Likes, "total: ", len(villager.Likes))
		fmt.Println("neutral: ", villager.Neutral, "total: ", len(villager.Neutral))
		fmt.Println("dislikes: ", villager.Dislikes, "total: ", len(villager.Dislikes))
		fmt.Println("hates: ", villager.Hates, "total: ", len(villager.Hates))
	}
}

func parseList(list string) []string {
	final := strings.Split(list, "\n")
	for i, item := range final {
		item = strings.ReplaceAll(item, "\n", "")
		var regex = regexp.MustCompile(`^\s`)
		res := regex.ReplaceAllLiteralString(item, "")
		final[i] = res
	}
	return final[1:]
}

func exportToJSON(villagers map[string]Villager) {
	ctr := 0                                 // ctr just to check if it's the first JSON object
	file, err := os.Create("villagers.json") // create the villagers.json file
	_, err = file.WriteString("[")           // write the opening square bracket for the JSON file
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, villager := range villagers { // iterate through the map of villagers
		v, err := json.MarshalIndent(villager, "", "\t") // format the object into a JSON supported text
		if err != nil {
			fmt.Println(err)
			return
		}

		if ctr > 0 {
			toWrite := ", " + string(v)
			_, err = file.Write([]byte(toWrite))
		} else {
			_, err = file.Write(v)
			ctr++
		}

		if err != nil {
			fmt.Println(err)
		}
	}

	_, err = file.WriteString("]") // add closing bracket
}
