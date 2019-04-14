# cpf

Check Plug-in Find: Minimalistic command line app to get information about Check_MK check plugins.

## Usage

Pipe `cpf` output into `fzf` to fuzzy find in Check_MK plugin meta data.

    cpf | fzf

## Installation

Download executable into user specific directory for executables and make it executable:

    wget -O ~/.local/bin/cpf https://github.com/fkromer/cpf/releases/tag/0.1.0
    chmod +x ~/.local/bin/cpf

## Usage

    cpf

## Build

Attention: This build procedure is quick and dirty, not reproducible.

    go get gopkg.in/urfave/cli.v1
    go get github.com/gocolly/colly
    go build main.go
    mv main cpf
    chmod +x cpf

## Wishlist

- packaging
- refactor command line interface to [cobra](https://github.com/spf13/cobra)

## Development resources

- [cli (github README)](https://github.com/urfave/cli/blob/master/README.md)
- [colly (website docs)](http://go-colly.org/docs/)
- [colly (GoDoc)](https://godoc.org/github.com/gocolly/colly)
- [Scraping the Web in Golang with Colly and Goquery](https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/)
