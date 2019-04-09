# cpf

Check Plug-in Find: Minimalistic command line app to get information about Check_MK check plugins.

## Usage

Pipe `cpf` output into `fzf` to fuzzy find in Check_MK plugin meta data.

    cpf | fzf

## Build

Attention: This build procedure is quick and dirty, not reproducible.

    go get github.com/anaskhan96/soup
    go build main.go
    mv main cpf
    chmod +x cpf

