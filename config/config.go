package config

import (
	"log"
	"os"

	"github.com/jon-perrett/go-doc-search/persist"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Store persist.StoreConfig
}

func ReadYaml(path string) *Config {
	f, err := os.Open(path)
	if err != nil {
		log.Default().Fatalf("Unable to read file %s: %s", path, err)
		panic(err)
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		log.Default().Fatalf("Unable to stat file %s", path)
		panic(err)
	}
	buf := make([]byte, stat.Size())
	f.Read(buf)
	var config Config
	yaml.Unmarshal(buf, &config)
	return &config
}
