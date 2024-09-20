package main

import (
	"github.com/jon-perrett/go-doc-search/persist"
)

func main() {
	persist.NewStore("./mydir")
}
