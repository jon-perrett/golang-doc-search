package persist

import (
	"fmt"
	"log"
	"os"
)

type Store struct {
	Path string
}

type Persistable interface {
	Bytes() []byte
	Id() string
}

func NewStore(path string) *Store {
	err := os.MkdirAll(path, 0744)
	if err != nil {
		log.Println(err)
	}
	return &Store{Path: path}
}

func (s *Store) Persist(item Persistable) (err error) {
	f, err := os.Create(fmt.Sprintf("%s/%s", s.Path, item.Id()))
	if err != nil {
		return err
	}
	defer func() error {
		if err := f.Close(); err != nil {
			return err

		}
		return nil
	}()
	if _, err := f.Write(item.Bytes()); err != nil {
		return err
	}
	return nil
}

func (s *Store) Retrieve(id string) ([]byte, error) {
	f, err := os.Open(fmt.Sprintf("%s/%s", s.Path, id))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	fmt.Printf("%d", stat.Size())
	buf := make([]byte, stat.Size())
	_, err = f.Read(buf)
	return buf, err
}
