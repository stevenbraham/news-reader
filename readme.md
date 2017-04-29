# About
This is a small a CLI program that I created so that I can quickly see newsheadlines in my terminal

# How to use

1. Make sure you have the [GoFeed library](https://github.com/mmcdole/gofeed) installed with `go get`
2. Run `go build`

You now have an executable binary program. You can directly run the program from the terminal.

It will ask from which you site you want have news and fetches the 10 latest headlines. You can give a news site as a argument, for example `./news-reader BBC`.

Currently the following news sites are supported:

* BBC
* CCN
* RT
* AlJazeera