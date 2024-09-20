package index

import (
	"github.com/jonperrett/go-doc-search/persist"
)

type Index struct {
	fieldToId map[string][]string
	store     *persist.Store
}

type Indexable interface {
	IndexKeys() []string
	GetDataAt(key string) string
	Id() string
}

func NewIndex(store *persist.Store) Index {
	mapping := make(map[string][]string)
	return Index{mapping, store}
}

func (idx *Index) Add(item Indexable) error {
	for _, k := range item.IndexKeys() {
		idx.fieldToId[item.GetDataAt(k)] = append(idx.fieldToId[k], item.Id())
	}
	return nil
}

func (idx *Index) Search(keyword string) [][]byte {
	nMatches := len(idx.fieldToId[keyword])
	results := make([][]byte, nMatches)
	for n := range nMatches {
		data, err := idx.store.Retrieve(idx.fieldToId[keyword][n])
		if err != nil {
			// TODO: handle
		}
		results[n] = data
	}
	return results
}
