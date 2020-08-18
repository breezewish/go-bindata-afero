# go-bindata-afero

A bridge between [go-bindata](https://github.com/go-bindata/go-bindata) and [Afero](https://github.com/spf13/afero), converts go-bindata embedded data into a [Afero Memory FS](https://github.com/spf13/afero#memmapfs).

## Usage

```go
fs := bindataafero.MustNewBindataFs(Asset, AssetInfo, AssetDir)
```

You can also write assets to an existing Afero FS, see [`fs_test.go`](./fs_test.go) for details.
