// Code generated for package bindataafero by go-bindata DO NOT EDIT. (@generated)
// sources:
// testdata/a/a/n.txt
// testdata/a/bar.txt
// testdata/a/foo.txt
// testdata/b/h.txt
// testdata/foo/foo.txt
// testdata/foo.txt
// testdata/root.txt
package bindataafero

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _testdataAANTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x4b\x2d\x2e\x49\x4d\x51\xc8\xcc\x53\x48\xd4\x4f\xe4\x02\x04\x00\x00\xff\xff\xc7\xaf\x90\x96\x0e\x00\x00\x00")

func testdataAANTxtBytes() ([]byte, error) {
	return bindataRead(
		_testdataAANTxt,
		"testdata/a/a/n.txt",
	)
}

func testdataAANTxt() (*asset, error) {
	bytes, err := testdataAANTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testdata/a/a/n.txt", size: 14, mode: os.FileMode(420), modTime: time.Unix(1597727847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testdataABarTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x4a\x2c\x52\xc8\xcc\x53\x48\xe4\x02\x04\x00\x00\xff\xff\x89\xb9\x23\x1f\x09\x00\x00\x00")

func testdataABarTxtBytes() ([]byte, error) {
	return bindataRead(
		_testdataABarTxt,
		"testdata/a/bar.txt",
	)
}

func testdataABarTxt() (*asset, error) {
	bytes, err := testdataABarTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testdata/a/bar.txt", size: 9, mode: os.FileMode(420), modTime: time.Unix(1597727802, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testdataAFooTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\xcb\xcf\x57\xc8\xcc\x53\x48\xe4\x02\x04\x00\x00\xff\xff\xb9\x07\x4c\x51\x09\x00\x00\x00")

func testdataAFooTxtBytes() ([]byte, error) {
	return bindataRead(
		_testdataAFooTxt,
		"testdata/a/foo.txt",
	)
}

func testdataAFooTxt() (*asset, error) {
	bytes, err := testdataAFooTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testdata/a/foo.txt", size: 9, mode: os.FileMode(420), modTime: time.Unix(1597727796, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testdataBHTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\xd0\x2b\xa9\x28\xe1\x02\x04\x00\x00\xff\xff\x3e\xe6\xad\x3e\x06\x00\x00\x00")

func testdataBHTxtBytes() ([]byte, error) {
	return bindataRead(
		_testdataBHTxt,
		"testdata/b/h.txt",
	)
}

func testdataBHTxt() (*asset, error) {
	bytes, err := testdataBHTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testdata/b/h.txt", size: 6, mode: os.FileMode(420), modTime: time.Unix(1597727831, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testdataFooFooTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\xcb\xcf\x57\xc8\xcc\x53\x48\xcb\xcf\xe7\x02\x04\x00\x00\xff\xff\xe5\x6c\xc1\x3b\x0b\x00\x00\x00")

func testdataFooFooTxtBytes() ([]byte, error) {
	return bindataRead(
		_testdataFooFooTxt,
		"testdata/foo/foo.txt",
	)
}

func testdataFooFooTxt() (*asset, error) {
	bytes, err := testdataFooFooTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testdata/foo/foo.txt", size: 11, mode: os.FileMode(420), modTime: time.Unix(1597727821, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testdataFooTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\xcb\xcf\x57\xc8\xcc\x53\x28\xca\xcf\x2f\xe1\x02\x04\x00\x00\xff\xff\x32\xb4\x50\xc6\x0c\x00\x00\x00")

func testdataFooTxtBytes() ([]byte, error) {
	return bindataRead(
		_testdataFooTxt,
		"testdata/foo.txt",
	)
}

func testdataFooTxt() (*asset, error) {
	bytes, err := testdataFooTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testdata/foo.txt", size: 12, mode: os.FileMode(420), modTime: time.Unix(1597727789, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testdataRootTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xca\xcf\x2f\x51\x48\xcb\xcc\x49\xe5\x02\x04\x00\x00\xff\xff\x63\x4e\x62\x6d\x0a\x00\x00\x00")

func testdataRootTxtBytes() ([]byte, error) {
	return bindataRead(
		_testdataRootTxt,
		"testdata/root.txt",
	)
}

func testdataRootTxt() (*asset, error) {
	bytes, err := testdataRootTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testdata/root.txt", size: 10, mode: os.FileMode(420), modTime: time.Unix(1597727784, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"testdata/a/a/n.txt":   testdataAANTxt,
	"testdata/a/bar.txt":   testdataABarTxt,
	"testdata/a/foo.txt":   testdataAFooTxt,
	"testdata/b/h.txt":     testdataBHTxt,
	"testdata/foo/foo.txt": testdataFooFooTxt,
	"testdata/foo.txt":     testdataFooTxt,
	"testdata/root.txt":    testdataRootTxt,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"testdata": &bintree{nil, map[string]*bintree{
		"a": &bintree{nil, map[string]*bintree{
			"a": &bintree{nil, map[string]*bintree{
				"n.txt": &bintree{testdataAANTxt, map[string]*bintree{}},
			}},
			"bar.txt": &bintree{testdataABarTxt, map[string]*bintree{}},
			"foo.txt": &bintree{testdataAFooTxt, map[string]*bintree{}},
		}},
		"b": &bintree{nil, map[string]*bintree{
			"h.txt": &bintree{testdataBHTxt, map[string]*bintree{}},
		}},
		"foo": &bintree{nil, map[string]*bintree{
			"foo.txt": &bintree{testdataFooFooTxt, map[string]*bintree{}},
		}},
		"foo.txt":  &bintree{testdataFooTxt, map[string]*bintree{}},
		"root.txt": &bintree{testdataRootTxt, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
