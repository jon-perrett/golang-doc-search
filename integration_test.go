package main

import (
	"testing"

	"github.com/jonperrett/go-doc-search/document"
	"github.com/jonperrett/go-doc-search/persist"
)

func TestCanStoreDocument(t *testing.T) {
	d := document.Document{}
	store := persist.NewStore(t.TempDir())
	store.Persist(&d)
}
