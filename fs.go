package bindataafero

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

type AssetFn func(name string) ([]byte, error)
type AssetInfoFn func(name string) (os.FileInfo, error)
type AssetDirFn func(name string) ([]string, error)

func writeAssetFile(fs afero.Fs, assetFn AssetFn, assetInfoFn AssetInfoFn, targetBaseDir, srcBaseDir, filePath string) error {
	srcFilePath := filepath.Join(srcBaseDir, filePath)
	targetFilePath := filepath.Join(targetBaseDir, filePath)
	data, err := assetFn(srcFilePath)
	if err != nil {
		return err
	}
	info, err := assetInfoFn(srcFilePath)
	if err != nil {
		return err
	}
	err = afero.WriteFile(fs, targetFilePath, data, info.Mode())
	if err != nil {
		return err
	}
	return fs.Chtimes(targetFilePath, info.ModTime(), info.ModTime())
}

// WriteAssetsInDirectory writes assets in `dirPath` to a afero.Fs directory `targetBaseDir`.
func WriteAssetsInDirectory(fs afero.Fs, assetFn AssetFn, assetInfoFn AssetInfoFn, assetDirFn AssetDirFn, targetBaseDir, srcBaseDir, dirPath string) error {
	srcDirPath := filepath.Join(srcBaseDir, dirPath)
	targetDirPath := filepath.Join(targetBaseDir, dirPath)

	children, err := assetDirFn(srcDirPath)
	if err != nil {
		// Should not happen, dirPath must exist
		panic(err)
	}

	if targetDirPath != "" {
		if err = fs.MkdirAll(targetDirPath, os.FileMode(0755)); err != nil {
			return err
		}
	}

	for _, child := range children {
		path := filepath.Join(dirPath, child)
		if _, err := assetDirFn(filepath.Join(srcBaseDir, path)); err == nil {
			// This is a directory
			if err := WriteAssetsInDirectory(fs, assetFn, assetInfoFn, assetDirFn, targetBaseDir, srcBaseDir, path); err != nil {
				return err
			}
		} else {
			// This is a file
			if err := writeAssetFile(fs, assetFn, assetInfoFn, targetBaseDir, srcBaseDir, path); err != nil {
				return err
			}
		}
	}

	return nil
}

// NewBindataFs creates a afero.Fs using specified assets. Errors may be returned.
func NewBindataFs(assetFn AssetFn, assetInfoFn AssetInfoFn, assetDirFn AssetDirFn) (afero.Fs, error) {
	fs := afero.NewMemMapFs()
	if err := WriteAssetsInDirectory(fs, assetFn, assetInfoFn, assetDirFn, "", "", ""); err != nil {
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
