A simple CLI app with a web scraper. 

This app basically goes through the [Stardew Valley's List of All Gifts](https://stardewvalleywiki.com/List_of_All_Gifts) page, gets the data from the table found in that page, and creates a JSON file (`villagers.json`) containing all of the gathered data.

This project also serves as *my* practice for **Golang**. 

## How to use
1. `go build -o stardew-gifts-scraper.exe`
2. Run `./stardew-gifts-scraper.exe`

