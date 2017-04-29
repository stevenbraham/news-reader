package main

import (
	"fmt"
	"news-reader/websites"
	"os"
)

func main() {
	//list all websites
	website := getSelectedWebsite()
	fmt.Println(website)
}

//Reads website from command line or displays an command line input tool
func getSelectedWebsite() websites.Website {
	if len(os.Args) == 2 {
		websiteName := os.Args[1]
		for _, website := range websites.Websites {
			if website.Name == websiteName {
				return website
			}
		}
		//looped over all options
		panic(websiteName + " is not a valid website")
	}
	//render input screen
	fmt.Println("Choose a news source:")
	for id, website := range websites.Websites {
		fmt.Println(id, website.Name)
	}
	fmt.Print("\nid:")
	index := 0
	fmt.Scan(&index)
	return websites.Websites[index] //index out of range will fire if the user selects an invalid id
}
