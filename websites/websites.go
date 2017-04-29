package websites

type Website struct {
	Name, Url string
}

//all news website that my program can use
var Websites = []Website{
	{
		Name: "BCC",
		Url:  "http://feeds.bbci.co.uk/news/video_and_audio/world/rss.xml",
	},
	{
		Name: "CNN",
		Url:  "http://rss.cnn.com/rss/edition.rss",
	},
	{
		Name: "RT",
		Url:  "http://feeds.bbci.co.uk/news/video_and_audio/world/rss.xml",
	},
	{
		Name: "AlJazeera",
		Url:  "http://www.aljazeera.com/xml/rss/all.xml",
	},
}
