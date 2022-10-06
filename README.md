# dag-store
Data store based on ipfs merkle-dag

```shell
go get github.com/CUIT-CBI/dag-store
```
## Getting Started
Create or open a store instance:
```go
store := dagstore.New(".data")
defer store.Close()
```
Add data and get its cid
```go
c, err := store.Add(context.Background(), file)
```
Get data from the store
```go
writerTo, err := store.Get(context.Background(), c)
```