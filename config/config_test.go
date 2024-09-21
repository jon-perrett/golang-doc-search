package config

import "testing"

func TestLoadStoreConfig(t *testing.T) {
	config := ReadYaml("./config.yml")
	if config.Store.Path != "./docstore" {
		t.Errorf("Did not load config correctly")
	}
}
