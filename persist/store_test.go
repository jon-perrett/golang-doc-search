package persist

import (
	"fmt"
	"os"
	"testing"
)

type persistMe struct {
	b  []byte
	id string
}

func (p *persistMe) Id() string {
	return p.id
}
func (p *persistMe) Bytes() []byte {
	return p.b
}

func TestPersistItem(t *testing.T) {
	path := t.TempDir()
	store := NewStore(StoreConfig{path})
	p := persistMe{b: []byte{1, 2}, id: "foo"}
	store.Persist(&p)
	fi, err := os.Stat(fmt.Sprintf("%s/foo", path))
	if err != nil {
		t.Errorf("Couldn't stat file")
		t.FailNow()
	}
	if fi.Size() != 2 {
		t.Errorf("Size of file incorrect")
	}
}
