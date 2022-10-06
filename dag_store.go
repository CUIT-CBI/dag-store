package dagstore

import (
	"context"
	levelstore "dagstore/store"
	"fmt"
	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	chunk "github.com/ipfs/go-ipfs-chunker"
	offline "github.com/ipfs/go-ipfs-exchange-offline"
	"github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-unixfs/importer/helpers"
	"github.com/ipfs/go-unixfs/importer/trickle"
	uio "github.com/ipfs/go-unixfs/io"
	"io"
)

type DagStore struct {
	//store   datastore.Batching
	service blockservice.BlockService
}

func New(dir string) DagStore {
	ds, err := levelstore.New(dir)
	if err != nil {
		fmt.Println(err)
	}
	bs := blockstore.NewBlockstore(ds)
	exch := offline.Exchange(bs) // TODO: block交换策略
	bserv := blockservice.New(bs, exch)
	return DagStore{bserv}
}

func (ds *DagStore) Close() error {
	if err := ds.service.Close(); err != nil {
		return err
	}
	return nil
}

func (ds *DagStore) Add(ctx context.Context, reader io.Reader) (cid.Cid, error) {
	dagService := merkledag.NewDAGService(ds.service)
	params := helpers.DagBuilderParams{ // TODO:
		Maxlinks:   10,
		RawLeaves:  true,
		CidBuilder: cid.V0Builder{},
		Dagserv:    dagService,
		NoCopy:     false,
	}

	chunker, err := chunk.FromString(reader, "")
	if err != nil {
		return cid.Cid{}, err
	}
	db, err := params.New(chunker)
	if err != nil {
		return cid.Cid{}, err
	}
	root, err := trickle.Layout(db)
	if err != nil {
		return cid.Cid{}, err
	}
	return root.Cid(), ds.service.AddBlock(ctx, root)
}

func (ds *DagStore) Get(ctx context.Context, c cid.Cid) (io.WriterTo, error) {
	dagService := merkledag.NewDAGService(ds.service)
	node, err := dagService.Get(context.Background(), c)
	if err != nil {
		return nil, err
	}
	return uio.NewDagReader(ctx, node, dagService)
}
