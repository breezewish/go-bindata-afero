package bindataafero

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

type AssetFn func(name string) ([]byte, error)
type AssetInfoFn func(name string) (os.FileInfo, error)
type AssetDirFn func(name string) ([]string, error)

func writeAssetFile(fs afero.Fs, assetFn AssetFn, assetInfoFn AssetInfoFn, filePath string) error {
	data, err := assetFn(filePath)
	if err != nil {
		return err
	}
	info, err := assetInfoFn(filePath)
	if err != nil {
		return err
	}
	err = afero.WriteFile(fs, filePath, data, info.Mode())
	if err != nil {
		return err
	}
	return fs.Chtimes(filePath, info.ModTime(), info.ModTime())
}

func writeAssetsInDirectory(fs afero.Fs, assetFn AssetFn, assetInfoFn AssetInfoFn, assetDirFn AssetDirFn, dirPath string) error {
	children, err := assetDirFn(dirPath)
	if err != nil {
		// Should not happen, dirPath must exist
		panic(err)
	}

	if dirPath != "" {
		if err = fs.MkdirAll(dirPath, os.FileMode(0755)); err != nil {
			return err
		}
	}

	for _, child := range children {
		path := filepath.Join(dirPath, child)
		if _, err := assetDirFn(path); err == nil {
			// This is a directory
			if err := writeAssetsInDirectory(fs, assetFn, assetInfoFn, assetDirFn, path); err != nil {
				return err
			}
		} else {
			// This is a file
			if err := writeAssetFile(fs, assetFn, assetInfoFn, path); err != nil {
				return err
			}
		}
	}

	return nil
}

// NewBindataFs creates a afero.Fs using specified assets. Errors may be returned.
func NewBindataFs(assetFn AssetFn, assetInfoFn AssetInfoFn, assetDirFn AssetDirFn) (afero.Fs, error) {
	fs := afero.NewMemMapFs()
	if err := writeAssetsInDirectory(fs, assetFn, assetInfoFn, assetDirFn, ""); err != nil {
		return nil, err
	}
	return fs, nil
}

// MustNewBindataFs creates a afero.Fs using specified assets. Errors will lead to panic.
func MustNewBindataFs(assetFn AssetFn, assetInfoFn AssetInfoFn, assetDirFn AssetDirFn) afero.Fs {
	fs, err := NewBindataFs(assetFn, assetInfoFn, assetDirFn)
	if err != nil {
		panic(err)
	}
	return fs
}
