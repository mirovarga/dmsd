package lib

import (
	"encoding/gob"
	"errors"
	"io/fs"
	"os"
)

const defaultDataFile = "dmsd.db"

type DB struct {
	file string
}

func (db DB) All() (items Items, err error) {
	var f *os.File
	f, err = os.Open(db.file)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return Items{}, nil
		}
		return nil, err
	}
	defer f.Close()

	err = gob.NewDecoder(f).Decode(&items)
	return
}

func (db DB) Store(items Items, replaceTags bool) error {
	all, err := db.All()
	if err != nil {
		return err
	}

	for _, i := range items {
		if existing, ok := all[i.FullPath]; !ok || ok && replaceTags {
			all[i.FullPath] = i
		} else {
			i.Tag(existing.Tags...)
			all[i.FullPath] = i
		}
	}

	f, err := os.Create(db.file)
	if err != nil {
		return err
	}
	defer f.Close()

	return gob.NewEncoder(f).Encode(all)
}

func NewDB(file string) DB {
	return DB{file: file}
}

func NewDefaultDB() DB {
	return NewDB(defaultDataFile)
}

func NewOrDefaultDB(file string) DB {
	if file == "" {
		return NewDefaultDB()
	}
	return NewDB(file)
}
