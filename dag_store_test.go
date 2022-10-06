package dagstore

import (
	"context"
	"fmt"
	"github.com/ipfs/go-cid"
	"os"
	"testing"
)

func add(store *DagStore) (cid.Cid, error) {
	file, err := os.OpenFile(".pdf", os.O_RDWR, 0755)
	defer file.Close()
	if err != nil {
		return cid.Cid{}, err
	}

	c, err := store.Add(context.Background(), file)
	if err != nil {
		return cid.Cid{}, err
	}
	return c, nil
}

func get(store *DagStore, c cid.Cid) error {
	outputFile, err := os.Create("test.pdf")
	defer outputFile.Close()
	if err != nil {
		return err
	}

	writerTo, err := store.Get(context.Background(), c)
	if err != nil {
		return err
	}
	writerTo.WriteTo(outputFile)
	return nil
}

func TestAdd(t *testing.T) {
	store := New("./.data")
	defer store.Close()

	c, err := add(&store)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(c)
}

func TestGet(t *testing.T) {
	store := New("./.data")
	defer store.Close()

	c, err := cid.Decode("QmVDQS58P34wT9MQNBBRTaNMmeTUXiaaGoztmx3wdbk5Ac")
	if err != nil {
		t.Fatal(err)
	}
	if err := get(&store, c); err != nil {
		t.Fatal(err)
	}
}

func TestAddGet(t *testing.T) {
	store := New("./.data")
	defer store.Close()

	c, err := add(&store)
	if err != nil {
		t.Fatal(err)
	}

	if err := get(&store, c); err != nil {
		t.Fatal(err)
	}
}
