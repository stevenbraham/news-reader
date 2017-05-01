package main

type Website struct {
	Name, Url string
}

//list of default websites

var DefaultWebsites = []Website{
	{
		Name: "BBC",
		Url:  "http://feeds.bbci.co.uk/news/video_and_audio/world/rss.xml",
	},
	{
		Name: "CNN",
		Url:  "http://rss.cnn.com/rss/edition.rss",
	},
	{
		Name: "RT",
		Url:  "https://www.rt.com/rss/",
	},
	{
		Name: "AlJazeera",
		Url:  "http://www.aljazeera.com/xml/rss/all.xml",
	},
}
