package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"news-reader/websites"
	"os"
)

func main() {
	loadWebsitesFromFile()
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

//build a lists of website from websites.yaml, if no file exists, it creates a default file
func loadWebsitesFromFile() {
	if _, err := os.Stat("websites.yaml"); os.IsNotExist(err) {
		//file does not exist, create default file
		encodedDefaultWebsites, _ := yaml.Marshal(DefaultWebsites)
		ioutil.WriteFile("websites.yaml", encodedDefaultWebsites, 0777)

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
