package document

import (
	"encoding/json"
)

type Document struct {
	Content     map[string]string
	DocumentId  string
	IndexFields []string
	Fields      []string
}

func (d *Document) Id() string {
	return d.DocumentId
}
func (d *Document) IndexKeys() []string {
	return d.IndexFields
}

func (d *Document) Bytes() []byte {
	b, err := json.Marshal(d.Content)
	if err != nil {
		return []byte{} // TODO: be better...
	}
	return b
}

func (d *Document) GetDataAt(key string) string {
	return d.Content[key]
}
