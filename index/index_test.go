package index

import (
	"testing"

	"github.com/jon-perrett/go-doc-search/document"
	"github.com/jon-perrett/go-doc-search/persist"
)

func TestDocumentIndexAndRetrieval(t *testing.T) {
	content := make(map[string]string)
	content["foo"] = "bar"
	d := document.Document{Content: content, DocumentId: "id", IndexFields: []string{"foo"}, Fields: []string{"foo"}}
	s := persist.NewStore(t.TempDir())
	s.Persist(&d)
	i := NewIndex(s)
	i.Add(&d)

	data := i.Search("bar")
	expected := "{\"foo\":\"bar\"}"
	if string(data[0][:]) != expected {
		t.Errorf("Expected %v, got %v", expected, string(data[0]))
	}

}
