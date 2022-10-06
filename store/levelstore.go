// Inplemented batching store
package levelstore

import (
	"context"
	"fmt"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelStore struct {
	path string
	leveldb.DB
}

func New(path string) (datastore.Batching, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelStore{path, *db}, nil
}

func (s *LevelStore) Get(ctx context.Context, key datastore.Key) (value []byte, err error) {
	fmt.Println("get", key.String())
	return s.DB.Get(key.Bytes(), nil)
}

func (s *LevelStore) Has(ctx context.Context, key datastore.Key) (exists bool, err error) {
	return s.DB.Has(key.Bytes(), nil)
}

func (s *LevelStore) GetSize(ctx context.Context, key datastore.Key) (size int, err error) {
	if value, err := s.Get(ctx, key); err != nil {
		return -1, err
	} else {
		return len(value), nil
	}
}

func (s *LevelStore) Query(ctx context.Context, q query.Query) (query.Results, error) {
	//TODO implement me
	panic("implement query")
}

func (s *LevelStore) Put(ctx context.Context, key datastore.Key, value []byte) error {
	fmt.Println("put:", key.String())
	return s.DB.Put(key.Bytes(), value, nil)
}

func (s *LevelStore) Delete(ctx context.Context, key datastore.Key) error {
	//TODO implement me
	panic("implement delete")
}

func (s *LevelStore) Sync(ctx context.Context, prefix datastore.Key) error {
	//TODO implement me
	panic("implement sync")
}

func (s *LevelStore) Batch(ctx context.Context) (datastore.Batch, error) {
	return nil, nil
}
