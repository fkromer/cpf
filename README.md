# cpf

Check Plug-in Find: Minimalistic command line app to get information about Check_MK check plugins.

## Usage

Pipe `cpf` output into `fzf` to fuzzy find in Check_MK plugin meta data.

    cpf | fzf

## Installation

Download executable into user specific directory for executables and make it executable:

    wget -O ~/.local/bin/cpf https://github.com/fkromer/cpf/releases/download/0.1.0/cpf
    chmod +x ~/.local/bin/cpf

## Usage

    cpf

## Build

To trigger a reproducible build execute:

    go build main.go
    mv main cpf
    chmod +x cpf

`go build` will execute `go mod download` implicitly to get the dependencies specified in `go.mod`
if required. This is the case when building the first time or if the versions of dependencies change.

## Wishlist

- packaging
- refactor command line interface to [cobra](https://github.com/spf13/cobra)
- fuzzy find in check metadata with e.g. [fuzzy](https://github.com/sahilm/fuzzy) and show check metadata dependent on selection

## Development resources

- [cli (github README)](https://github.com/urfave/cli/blob/master/README.md)
- [colly (website docs)](http://go-colly.org/docs/)
- [colly (GoDoc)](https://godoc.org/github.com/gocolly/colly)
- [Scraping the Web in Golang with Colly and Goquery](https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/)
- [blog.golang.org - Using Go modules](https://blog.golang.org/using-go-modules)
