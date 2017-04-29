package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"news-reader/websites"
	"os"
)

func main() {
	website := getSelectedWebsite()
	fmt.Println("Headlines from", website.Name, "\n")
	//read items with the rss library
	feed, _ := gofeed.NewParser().ParseURL(website.Url)
	for index, item := range feed.Items {
		fmt.Println(item.Title)
		if index > 9 {
			//stop after 10 items
			break
		}
	}
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
