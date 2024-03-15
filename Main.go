package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Stardew Valley Gifts Finder!")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("-> ")
	scanner.Scan()

menu:
	for {
		fmt.Println("Menu")
		fmt.Println("1 - Start scraper (default)\n2 - Start scraper (Vercel)\n3 - Display villager details\n4 - Display all villager details\n0 - EXIT")
		fmt.Print("-> ")
		scanner.Scan()
		input := scanner.Text()

		switch {
		case input == "1":
			fmt.Println("Starting web scraper...")
			ScrapeGifts(1)
		case input == "2":
			fmt.Println("Starting web scraper...")
			ScrapeGifts(2)
		case input == "0":
			break menu
		default:
			break menu
		}
	}
}
