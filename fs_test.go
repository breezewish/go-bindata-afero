package bindataafero

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func walk(t *testing.T, fs afero.Fs, root string) string {
	walked := ``
	err := afero.Walk(fs, root, func (path string, info os.FileInfo, err error) error {
		assert.Nil(t, err)
		if info.IsDir() {
			walked += fmt.Sprintf("%s: (Directory)\n", path)
		} else {
			content, err := afero.ReadFile(fs, path)
			assert.Nil(t, err)
			walked += fmt.Sprintf("%s: %s\n", path, strings.TrimSpace(string(content)))
		}
		return nil
	})
	assert.Nil(t, err)
	return strings.TrimSpace(walked)
}

func TestNewBindataFs(t *testing.T) {
	fs, err := NewBindataFs(Asset, AssetInfo, AssetDir)
	assert.Nil(t, err)

	c := walk(t, fs, "")
	assert.Equal(t, strings.TrimSpace(`
: (Directory)
testdata: (Directory)
testdata/a: (Directory)
testdata/a/a: (Directory)
testdata/a/a/n.txt: nested in a/a
testdata/a/bar.txt: bar in a
testdata/a/foo.txt: foo in a
testdata/b: (Directory)
testdata/b/h.txt: h.txt
testdata/foo: (Directory)
testdata/foo/foo.txt: foo in foo
testdata/foo.txt: foo in root
testdata/root.txt: root file
`), c)

	c = walk(t, fs, "testdata/a")
	assert.Equal(t, strings.TrimSpace(`
testdata/a: (Directory)
testdata/a/a: (Directory)
testdata/a/a/n.txt: nested in a/a
testdata/a/bar.txt: bar in a
testdata/a/foo.txt: foo in a
`), c)
}


func TestWriteAssetsInDirectory(t *testing.T) {
	fs := afero.NewMemMapFs()
	err := WriteAssetsInDirectory(fs, Asset, AssetInfo, AssetDir, "", "", "")
	assert.Nil(t, err)
	c := walk(t, fs, "")
	assert.Equal(t, strings.TrimSpace(`
: (Directory)
testdata: (Directory)
testdata/a: (Directory)
testdata/a/a: (Directory)
testdata/a/a/n.txt: nested in a/a
testdata/a/bar.txt: bar in a
testdata/a/foo.txt: foo in a
testdata/b: (Directory)
testdata/b/h.txt: h.txt
testdata/foo: (Directory)
testdata/foo/foo.txt: foo in foo
testdata/foo.txt: foo in root
testdata/root.txt: root file
`), c)

	fs = afero.NewMemMapFs()
	err = WriteAssetsInDirectory(fs, Asset, AssetInfo, AssetDir, "foo/bar", "", "")
	assert.Nil(t, err)
	c = walk(t, fs, "")
	assert.Equal(t, strings.TrimSpace(`
: (Directory)
foo: (Directory)
foo/bar: (Directory)
foo/bar/testdata: (Directory)
foo/bar/testdata/a: (Directory)
foo/bar/testdata/a/a: (Directory)
foo/bar/testdata/a/a/n.txt: nested in a/a
foo/bar/testdata/a/bar.txt: bar in a
foo/bar/testdata/a/foo.txt: foo in a
foo/bar/testdata/b: (Directory)
foo/bar/testdata/b/h.txt: h.txt
foo/bar/testdata/foo: (Directory)
foo/bar/testdata/foo/foo.txt: foo in foo
foo/bar/testdata/foo.txt: foo in root
foo/bar/testdata/root.txt: root file
`), c)

	fs = afero.NewMemMapFs()
	err = WriteAssetsInDirectory(fs, Asset, AssetInfo, AssetDir, "foo/bar", "testdata/b", "")
	assert.Nil(t, err)
	c = walk(t, fs, "")
	assert.Equal(t, strings.TrimSpace(`
: (Directory)
foo: (Directory)
foo/bar: (Directory)
foo/bar/h.txt: h.txt
`), c)

	fs = afero.NewMemMapFs()
	err = WriteAssetsInDirectory(fs, Asset, AssetInfo, AssetDir, "foo/bar", "", "testdata/b")
	assert.Nil(t, err)
	c = walk(t, fs, "")
	assert.Equal(t, strings.TrimSpace(`
: (Directory)
foo: (Directory)
foo/bar: (Directory)
foo/bar/testdata: (Directory)
foo/bar/testdata/b: (Directory)
foo/bar/testdata/b/h.txt: h.txt
`), c)

	fs = afero.NewMemMapFs()
	err = WriteAssetsInDirectory(fs, Asset, AssetInfo, AssetDir, "/", "", "testdata/b")
	assert.Nil(t, err)
	c = walk(t, fs, "/")
	assert.Equal(t, strings.TrimSpace(`
/: (Directory)
/testdata: (Directory)
/testdata/b: (Directory)
/testdata/b/h.txt: h.txt
`), c)

	c = walk(t, fs, "/testdata")
	assert.Equal(t, strings.TrimSpace(`
/testdata: (Directory)
/testdata/b: (Directory)
/testdata/b/h.txt: h.txt
`), c)
}
