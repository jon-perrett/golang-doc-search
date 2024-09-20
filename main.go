package main

import (
	"github.com/jonperrett/go-doc-search/persist"
)

func main() {
	persist.NewStore("./mydir")
}
