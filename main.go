package main

import (
	"github.com/jon-perrett/go-doc-search/config"
	"github.com/jon-perrett/go-doc-search/persist"
)

func main() {
	config := config.ReadYaml("config.yml")
	persist.NewStore(config.Store)
}
