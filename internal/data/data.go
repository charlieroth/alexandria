package data

import (
	"fmt"
	"sync"
)

type JsonMutexDB struct {
	BigLock   sync.Mutex
	Documents map[string]Document
}

func (db *JsonMutexDB) AddDocument(doc Document) error {
	db.Documents[doc.Id] = doc
	return nil
}

func (db *JsonMutexDB) GetDocument(id string) (Document, error) {
	if doc, ok := db.Documents[id]; ok {
		return doc, nil
	}

	return Document{}, fmt.Errorf("document with id %s already exists", id)
}

func NewJsonMutexDB() *JsonMutexDB {
	db := JsonMutexDB{}
	db.BigLock = sync.Mutex{}
	db.Documents = make(map[string]Document)
	return &db
}
