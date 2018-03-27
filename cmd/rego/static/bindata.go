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

var _templatesSpecMdGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\xcb\x4e\xc3\x30\x10\xbc\xe7\x2b\x56\x0d\x07\x38\xd4\x1f\x50\x09\x24\xa4\xf6\x06\x1c\xa0\xe2\x1c\xd7\x59\xa8\x51\x1a\x1b\xef\x56\x4a\x15\xf9\xdf\x91\x1f\x75\x92\x8a\x9b\x3d\x9e\x9d\xd9\x59\x6f\x5d\xc3\x38\x82\xf8\xb0\xa8\xc4\xab\x69\xb1\x13\xbb\x9e\x35\x5f\xde\xe4\x09\xc1\xfb\x6a\x1c\xe1\xae\x93\x8c\xc4\x9f\xe8\x48\x9b\x1e\x36\x8f\x99\xfe\xb2\x80\x13\x97\x9d\x3e\x91\x95\x0a\x17\x92\x5b\x24\xe5\xb4\xe5\xcc\x8b\xa2\x38\xc8\x93\xed\x30\xc8\x5d\x8f\xb1\xe4\xd6\x2e\xe9\xea\xaf\xa9\xc2\xfb\xaa\xae\x6b\xd8\xa5\x6b\x55\x35\x4d\xf3\x43\xa6\x5f\xa8\x7a\x1f\xe0\x00\x61\xdf\x16\x4f\x63\x29\xf8\x19\x8b\x4e\x86\x66\x28\x5b\x8a\x77\xec\x12\x70\xd4\x36\x80\xc8\x33\x5f\x63\x29\x5b\x16\xda\x24\x97\x68\x73\x13\xc9\xec\xf4\xe1\xcc\x48\xd3\xa8\x76\x83\x35\x84\xed\xf3\xf4\xf4\x5f\xca\x75\xb4\x9b\xd5\x07\xc9\xe0\x3b\xd5\x45\x0b\x27\xfb\x6f\x5c\x10\xd7\xa9\xc1\x1a\x9a\xf0\x99\xf9\xef\xe0\x3e\x05\xc0\x5f\x10\xfb\x8b\x45\x58\xe1\xc0\xe8\x7a\xd9\xad\xc0\xfb\x40\x8c\xa8\xf7\x9b\xb8\x01\xe7\x43\xbe\x86\x3c\x1d\xe1\x92\x53\x42\x3e\x34\x79\x2e\x62\x8b\xd6\xa1\x92\x8c\x31\xfb\x13\xec\x8f\x9a\xa0\x34\x05\x9a\xa0\x2d\x8c\xd9\x90\x96\x5b\x72\xb3\x1a\xe3\x08\xea\x28\x9d\x54\x8c\x4e\x13\x6b\x45\x20\xae\xc3\x29\xf5\xe5\xf8\x17\x00\x00\xff\xff\xab\x76\x97\x9d\xbd\x02\x00\x00")

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

	info := bindataFileInfo{name: "templates/spec-md.gotpl", size: 701, mode: os.FileMode(420), modTime: time.Unix(1522180044, 0)}
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

