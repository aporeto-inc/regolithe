// Code generated by go-bindata.
// sources:
// templates/spec-md.gotpl
// templates/toc-md.gotpl
// DO NOT EDIT!

package static

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesSpecMdGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\xc1\x6a\xc3\x30\x10\x44\xef\xf9\x8a\x25\xea\xa1\x3d\x44\x1f\x10\x68\xa1\x90\x1c\xdb\x43\x9b\x0f\x88\x22\x4f\xeb\x05\xc7\x56\xa5\x0d\x24\x08\xfd\x7b\x91\xe4\xd8\xb8\xd0\x9b\x3d\x7a\xb3\x9a\x59\x29\x45\x31\x92\xfe\x74\xb0\xfa\x6d\x68\xd0\xe9\x7d\x2f\x2c\xb7\x77\x73\x06\xa5\xb4\x5a\xc5\x48\xe2\xf9\x1c\x9c\xb1\x58\x70\x3b\x04\xeb\xd9\x09\x0f\xfd\x1d\x7c\x18\x5c\xa0\xed\x33\x0d\x0e\xde\xe4\x83\x50\x1d\xa4\x3f\xd0\x55\xa1\x65\x97\x45\x48\xf6\xc4\x48\xfc\x55\x5c\x29\xad\x94\x52\x34\x61\xf3\xb8\x8a\xa1\x6f\xc6\x4b\x36\xd9\x52\x73\xec\xaf\x6e\x08\x68\x5e\x45\x3c\x9f\x2e\x82\x02\x97\x39\xb3\x54\x06\x79\xd3\x7f\xe3\x5f\xd3\xa6\x5e\xae\xe8\x98\x37\x31\x16\xa7\xc7\x1a\x0e\x3f\xa4\x0f\x37\x07\x5a\xe3\x2a\xf0\xbd\xe9\xd6\x94\x52\x06\x8b\x9a\xd2\xb6\xac\xef\x72\x1a\x7f\x73\xd6\x2e\x60\xc9\x4c\x05\x9e\x8e\x63\x67\xbd\x83\xf3\xb0\x46\xd0\xa4\xf4\x42\x87\x96\x03\x99\x7b\x24\xe2\x40\xcd\x74\x3e\xb7\x5f\x3e\xc5\x9f\xfd\xc7\x48\xb6\x35\xde\x58\x81\xe7\x20\x6c\x03\xe9\xaa\x6f\x66\xff\xf4\xf9\x1b\x00\x00\xff\xff\x1e\x6b\xd2\xde\xf7\x01\x00\x00")

func templatesSpecMdGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSpecMdGotpl,
		"templates/spec-md.gotpl",
	)
}

func templatesSpecMdGotpl() (*asset, error) {
	bytes, err := templatesSpecMdGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/spec-md.gotpl", size: 503, mode: os.FileMode(420), modTime: time.Unix(1521690895, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTocMdGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcd\x31\x6b\xc3\x30\x10\x05\xe0\x5d\xbf\xe2\xaa\x2e\xed\x20\x51\x4a\xbb\x94\x62\x30\xf5\x62\x4a\xc0\x10\xc8\x1a\x14\xf9\x9c\x1c\x91\xee\x8c\x2c\x13\x82\xd1\x7f\x0f\x36\x59\xde\xf0\xe0\x7b\xef\xf7\xc5\x18\x88\x2e\x5d\x7b\xb9\x71\x20\xce\xa6\xa7\xc9\x9d\x02\xc2\xae\xf9\xf8\xfc\xda\xf2\x1b\x8c\xa9\x94\x7a\x85\x65\x81\x4c\x39\x20\xbc\xd9\x3f\xe1\x81\xce\x73\x72\x99\x84\xed\x3f\xde\x41\x47\x61\x09\x94\x2f\xa8\x41\x8f\x49\xfa\xd9\xe7\x23\xbb\x88\xfa\x1d\x4a\x81\xba\x6b\xa1\x11\x3f\x47\xe4\xbc\x21\xa5\x2a\x38\x60\x9a\x48\xf8\x67\x5d\xb6\x75\xd7\xb6\x3c\x88\x7d\x96\x50\x8a\x52\xeb\xa3\x78\xb0\xfb\x11\x3d\x0d\xe4\x37\x39\x81\x29\x45\x3d\x02\x00\x00\xff\xff\x2f\xc9\x8f\x8f\xba\x00\x00\x00")

func templatesTocMdGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesTocMdGotpl,
		"templates/toc-md.gotpl",
	)
}

func templatesTocMdGotpl() (*asset, error) {
	bytes, err := templatesTocMdGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/toc-md.gotpl", size: 186, mode: os.FileMode(420), modTime: time.Unix(1521249819, 0)}
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
	"templates/spec-md.gotpl": templatesSpecMdGotpl,
	"templates/toc-md.gotpl": templatesTocMdGotpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"spec-md.gotpl": &bintree{templatesSpecMdGotpl, map[string]*bintree{}},
		"toc-md.gotpl": &bintree{templatesTocMdGotpl, map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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

