# go-bindata-afero

A bridge between [go-bindata](https://github.com/go-bindata/go-bindata) and [Afero](https://github.com/spf13/afero), converts go-bindata embedded data into a [Afero Memory FS](https://github.com/spf13/afero#memmapfs).

## Usage

```go
fs, err := NewBindataFs(Asset, AssetInfo, AssetDir)
```
