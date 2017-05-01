# About
This is a small a CLI program that I created so that I can quickly see newsheadlines in my terminal

# Required libraries

Make sure to `go get` the following libraries:

* [mmcdole/gofeed](github.com/mmcdole/gofeed)

# How to use

Run `go build` to create a local executable or `go install news-reader`.

You now have an executable binary program. You can directly run the program from the terminal.

It will ask from which you site you want have news and fetches the 10 latest headlines. You can give a news site as a argument, for example `./news-reader BBC`.

Currently the following news sites are supported:

* BBC
* CCN
* RT
* AlJazeera