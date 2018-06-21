// Code generated by go-bindata.
// sources:
// gen.sh
// rego-abstract.in
// rego-abstract.json
// rego-attribute.in
// rego-info.json
// rego-spec.in
// rego-spec.json
// rego-type-mapping.json
// DO NOT EDIT!

package schema

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

var _genSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xcf\x31\x0a\x02\x31\x10\x85\xe1\x3e\xa7\x18\x57\x21\x5a\x24\x39\x80\x20\x28\x58\xd8\xca\x5a\x87\x49\x8c\xd9\xa8\x9b\x84\xcc\x78\x7f\x71\xd9\x42\x6c\x2d\xdf\x5f\x7c\xf0\x96\x0b\xe3\x52\x36\x0e\x69\x10\xa2\x86\xf6\x04\x55\x03\x48\x32\xd6\xee\xfb\xfe\x7c\x3a\x5c\xfa\xa3\xb5\x46\x76\xab\xb5\x47\x86\x16\x62\x51\xc8\xdc\x92\x7b\x71\xd0\x29\x6f\x3a\x69\xa2\x9c\xbb\x23\x6e\xe8\x59\xa7\x0c\xbb\x9f\x74\xa7\x92\xff\xf5\xa9\x06\xff\x65\x4f\x73\x72\x45\x1b\x41\xdd\x40\x7f\xae\x5c\x91\x51\xc7\xb2\x15\xb1\xa8\x79\x82\xaa\x8f\x08\xe4\x87\x30\x22\x68\xf1\x0e\x00\x00\xff\xff\x19\x53\xac\x2b\xf3\x00\x00\x00")

func genShBytes() ([]byte, error) {
	return bindataRead(
		_genSh,
		"gen.sh",
	)
}

func genSh() (*asset, error) {
	bytes, err := genShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gen.sh", size: 243, mode: os.FileMode(493), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoAbstractIn = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xb1\x4e\xc4\x30\x10\x44\xfb\x7c\xc5\x2a\x47\x79\x87\xa9\x28\xd2\x1d\x12\x05\x1d\x42\xa1\x8e\x36\xf6\x9a\xf8\x64\x6c\x6b\x77\x4f\x08\xa1\xfc\x3b\x4a\x0c\xe1\xa0\xa0\xb3\x76\xde\x78\x46\xf3\xd1\x00\x00\xb4\x57\x62\x27\x7a\xc5\xb6\x83\x76\x52\x2d\x9d\x31\x27\xc9\xe9\x50\xaf\xd7\x99\x5f\x8c\x63\xf4\x7a\xb8\xb9\x35\xf5\xb6\x6b\xf7\xd5\xa9\x41\x23\x2d\xbe\xe3\x28\xca\x68\xf5\x5b\x70\x24\x96\x43\xd1\x90\xd3\x2a\x27\xc0\x2f\x02\x30\xc6\xfc\x26\xa0\x19\x1c\xd9\x88\x4c\x80\x20\xa4\x90\x3d\x8c\x28\x04\xa8\xca\x61\x3c\x2b\x09\x60\x72\x26\x33\x30\x45\x5c\x7e\x12\xd0\x09\x15\x2c\x26\x18\x09\x98\xce\x42\x0e\x42\xd2\x0c\x59\x27\x62\x90\x42\x36\xf8\x60\x2b\xbc\x75\x7c\x2f\x6b\xc5\x3c\x9e\xe8\xa7\x20\x3a\x17\x16\x0c\xe3\x23\xe7\x42\xac\x81\xa4\xed\xc0\x63\x14\xda\x37\x95\x29\x97\x4a\xdd\xaa\x7a\xb7\x86\xbf\xee\x75\x4b\x26\xbf\xa4\xed\x8c\x23\x1f\xd2\x9a\x21\xe6\xc2\xb1\xe1\xf3\xfa\x9a\xb7\xc1\x36\xfa\xdf\xb0\x61\x38\xf6\xfd\xd3\xc3\xdd\x73\x7f\x3f\x0c\x7f\xff\x6a\xe6\xe6\x33\x00\x00\xff\xff\x33\x37\x6b\x77\xd3\x01\x00\x00")

func regoAbstractInBytes() ([]byte, error) {
	return bindataRead(
		_regoAbstractIn,
		"rego-abstract.in",
	)
}

func regoAbstractIn() (*asset, error) {
	bytes, err := regoAbstractInBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-abstract.in", size: 467, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoAbstractJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\x4d\x6f\xdb\x38\x13\xbe\xe7\x57\x0c\xdc\x02\x05\xde\x37\xae\xbb\xc0\xa2\x40\x73\xcb\xa1\x87\x00\x7b\x28\x16\x7b\x6a\x90\x06\x23\x71\x64\xb1\xa5\x48\x75\x48\x39\xf6\x2e\xf2\xdf\x17\x24\x25\x59\x4e\xfc\x21\x5b\x8a\xd3\xed\xa5\x8e\x38\x1c\xce\x33\x1f\x0f\x87\x94\xfe\xb9\x00\x00\x98\xbc\xb5\x69\x4e\x05\x4e\xae\x60\x92\x3b\x57\x5e\xcd\x66\xdf\xad\xd1\xd3\xf8\xf4\xbd\xe1\xf9\x4c\x30\x66\x6e\xfa\xe1\xe3\x2c\x3e\x7b\x33\xb9\x8c\x33\x9d\x74\x8a\xfc\xbc\xeb\xc4\x3a\xc6\xd4\x35\x03\x82\x6c\xca\xb2\x74\xd2\xe8\x30\xac\x01\x6b\x09\x40\xa5\xcc\x83\x05\x67\x40\x50\xaa\x90\x09\x10\x2c\x39\x30\x19\x24\x68\x09\xd0\x39\x96\x49\xe5\xc8\x02\x6a\x31\x33\x0c\x4c\x0a\xbd\x26\x0b\x2e\x47\x07\x29\x6a\x48\x08\x98\x2a\x4b\x02\xa4\x76\x06\x8c\xcb\x89\xc1\x96\x94\xca\x4c\xa6\x51\xb8\xb5\x71\x55\x06\x13\x4d\xf2\x9d\xd6\x06\xa2\x10\xd2\x8b\xa1\xfa\xc2\xa6\x24\x76\x92\xec\xe4\x0a\x32\x54\x96\x2e\x2f\xa2\x4c\xd9\x1d\x89\xbe\x8a\x73\x5b\x0b\x37\x9e\x47\x5f\x32\x65\x7e\xb5\x37\x33\x41\x99\xd4\x61\x0d\x3b\xeb\xcc\x68\xc5\x1f\xc3\xaf\xc7\xd6\x61\xad\x74\xff\xc5\xb6\x42\x6b\x47\x9f\xc4\xe0\x0f\x69\x83\x93\x17\xc4\x56\x1a\x4d\xa2\xe3\xe9\xf7\x4f\xe7\xee\x77\xcf\x86\x68\x89\xce\x11\xeb\x2f\xdb\x9d\xd5\x8a\x7d\x5b\xdc\x7e\x98\x7e\xba\xfb\xff\xdb\xad\xc3\x1b\x70\x90\x19\x57\x4f\x2c\x3a\x84\xaa\x93\x35\x52\x03\x42\x89\xec\x64\x5a\x29\xe4\x06\xf0\x7b\xb8\x71\xf0\x20\x95\x82\xca\x67\x99\x52\xdd\x39\x19\x9b\x02\x5c\x4e\x50\x32\x2d\xa4\xa9\x6c\x33\x2b\x24\x21\xa0\x10\x60\x18\xcc\x82\xf8\x81\xa5\xa3\x20\x6a\x34\xd9\x26\x87\xc5\x53\x0f\xb6\xf6\x4a\x47\xc5\x76\x97\xf4\x0b\xe3\x3e\xf0\x7f\x52\xc9\x64\x49\x3b\x6f\xe5\x1a\x8e\xf7\x87\x37\x90\x9b\x61\x12\x10\x75\xef\xb2\x32\x28\x67\xfa\x59\x49\x26\x31\xb9\x82\xdb\x9d\x52\x41\x52\x63\x41\x7b\x34\xad\x41\x1d\x90\xe9\xc2\xd9\x29\x79\xb7\xc7\xe4\xfe\x69\xba\x31\x6d\x47\x61\x6f\x5f\xc2\xd3\x15\x89\xfb\x34\x47\x3e\x2c\xfe\x0c\x56\x88\xd2\x9c\x96\x65\xe4\x2e\x04\xeb\x58\xea\x79\x27\x58\x45\x65\x1d\xe4\x46\x1b\xf6\x9c\x98\x10\x2c\x50\x49\x71\xc0\x75\xd0\xcd\x9b\xa8\x72\xb7\x07\x61\xcd\x33\x3d\x60\x1a\x99\xf6\xf0\xcb\x36\xa0\xd7\x51\x87\x07\x50\xf9\x9a\x32\xfc\x3c\x2d\x57\x25\xc1\x67\x5d\x15\xc7\xe0\xdb\xc7\x07\x1b\x13\x0e\xd7\xda\x56\xfd\xb5\xff\x0e\x2f\x00\x1d\xbe\xf3\x13\xbf\xdd\x5e\x4f\xbf\xfe\x36\xfd\x74\x77\x8b\xd3\xbf\xaf\xa7\x5f\x3d\xc1\xfd\xef\xed\xfe\x40\x40\x4b\xfd\x3b\x47\x0f\x85\xaa\x72\x66\x4e\x9a\x18\x5d\x28\xd6\xe3\x03\xf5\x57\xde\xd9\x66\x41\x5a\xd8\x50\x09\xc9\x2a\xd0\x47\x82\xe9\x0f\xd2\x47\x65\x62\x62\x8c\x22\xdc\x53\xcc\x7d\xf0\xa5\x4c\x61\x07\xbf\x37\x5a\xad\x46\xc0\xe7\xfb\x05\xaf\xca\xd7\x96\x6f\x34\x44\x15\x2a\x30\x2c\x43\xe0\x99\x20\x2c\x77\x76\x9c\x82\x32\xac\x94\xbb\x37\x2c\x88\x4f\xc6\x19\x8a\xad\x61\xfc\x8d\xa0\x86\xf6\x08\xfd\x16\x15\xd6\x81\xb0\x8e\x47\xfe\x83\x7a\x55\xd3\x8b\x80\x0d\xe6\x9e\x0c\xb6\x81\xb2\x1d\xf4\x11\xa0\x6e\x1b\x52\x59\xc3\xbb\x84\x89\xd4\x8e\xe6\xc4\xfe\xa7\xae\x94\x8a\xff\x17\x49\x7c\xd2\xec\xce\x2d\x5f\xdc\x0d\x74\x48\xc9\x94\x8e\x57\xc2\x1d\x7d\xe7\x8f\x6d\xd7\xb4\x53\x23\xdb\x3e\x1a\x12\xd7\x51\xf6\x42\x5a\x62\x51\x2a\x1a\x90\xab\x9f\xa3\x86\x3a\x4f\xfd\x46\xf8\xdf\x4d\x54\x5a\x96\xc6\x8e\x95\xa5\x65\x95\x28\x99\xaa\x15\x34\x5a\x4f\xc8\xd5\x1e\x53\x6a\x9e\xf0\xb3\x1c\x57\x34\x2c\x1d\x32\xa9\x1c\x31\x26\xea\x74\xde\xda\xdc\x8c\x12\x8a\xdc\xec\xf3\x22\x2a\xf7\xac\x5c\x56\xec\x5d\x62\xcf\x5e\xbe\x99\x61\x92\x73\x7d\xef\xb7\x85\x71\xba\x09\xa8\x55\xbe\xca\x4e\x93\x19\x2e\xd0\x9d\x8c\x24\x4e\x3f\x89\x83\xc8\xf7\xb5\x87\x0e\x4f\x6b\x3b\x99\xfa\x68\x8d\x9a\x0b\x94\xaa\x77\x83\x9a\x1b\xdd\x5b\xf3\xcd\x97\xc5\xef\x47\xc8\x7e\xec\x2b\x9b\x4a\xc1\x7d\x65\x05\xba\x03\x25\xea\xff\x0d\x63\xb1\x39\xf9\xae\x7d\x84\xfc\x0e\x37\x08\x39\x2e\x08\x10\xd6\x2d\x73\xad\xfe\xdc\xb9\x2e\x05\x69\x27\x33\x39\x0a\xb2\x4e\xcb\x88\x50\xa0\xd4\xd0\x51\xff\x32\x3c\x5d\x69\xf9\xb3\xa2\x9b\xfa\xd8\xe6\xa9\x7a\x98\x3b\xb4\xa0\xe5\xc8\x9e\xd0\x10\xb5\x9e\x3b\xb4\x05\x2e\xef\x15\xe9\xb9\xcb\x4f\x06\x54\xe0\x52\x16\x55\x01\x51\xcd\xf3\x33\x42\x73\x1e\xef\x7d\xfe\x6d\x11\x36\xfd\xc7\x60\x84\xc3\x8e\x03\x0d\xc0\x1d\x67\xa0\x06\x5f\xdb\x2d\x9d\x1b\xa0\xd4\x83\x43\x28\xf5\x2f\x1d\x42\xa9\x87\x86\xb0\x06\xf8\x8b\x86\x30\xdc\x74\x9e\x8a\xcd\x4f\x7e\xf5\xc3\x8c\x29\xa4\xbb\xa7\xa2\x74\xa7\x35\x77\x37\x59\xa0\xe5\xcb\x27\x71\xa9\xdb\x58\xc1\xa6\x2c\x49\x80\xcc\x20\x2c\x01\x86\x21\x1e\x48\xce\xcc\x96\xe1\x6e\xe3\x85\x5a\xf4\xf6\xde\xe4\xd5\x3a\xf4\x92\x65\x81\xbc\x1a\xaf\x43\x8f\xbb\x1b\xd4\x7a\x5f\xa5\x4d\x67\x42\x31\xd6\x0d\x9f\xb4\xe0\xd5\x85\x5b\xbe\x57\x00\xd2\xbe\x36\x19\x05\x47\xad\xed\xdc\x30\x2c\xa5\x4c\xa7\x1f\x9b\x36\x40\x44\x5d\xe1\xcd\x99\xcd\x4d\xa5\x44\xe0\x0a\x42\xe5\xe0\x41\xba\x1c\x52\xe4\xa3\x78\x70\x24\x80\x2f\x7a\x02\xb0\xaf\x73\x02\xb0\xce\x8c\x96\x7b\xb5\xae\x97\xbe\x92\x19\xdc\xe6\xdb\x2a\xa9\x0d\x18\xb6\x2f\x17\x58\x96\x9e\xd7\x9d\x81\x52\xa6\x3f\x02\xd9\x3f\xe4\xa4\xc3\x5e\x17\xdb\x0e\x1b\xde\x18\x38\x03\xef\x68\xe9\x88\x35\xaa\x77\xe7\xde\xc0\x1d\xa3\xb6\x92\xf4\x48\xb5\xb9\x56\x77\xee\x54\x1d\x14\xb3\x10\x8e\xee\x7b\xfe\x13\xc2\x30\xfa\x15\xce\x71\x2f\x0e\xfb\xf7\xb0\x41\x3c\x53\x06\xfb\x04\x29\x08\xf7\x2f\x44\x68\x61\xf6\x94\x55\xd2\xf6\x36\xe3\xe0\x17\x0c\x1b\xd2\x4e\x1e\xfc\x8c\xa0\x95\x6d\xca\x6f\xf0\x3d\xd1\xce\xd1\xed\x23\xcf\x9f\x6e\x3e\x59\xff\x55\x7f\xcf\x73\xf1\x78\xf1\x6f\x00\x00\x00\xff\xff\xe1\x5e\xb6\xd9\x57\x25\x00\x00")

func regoAbstractJsonBytes() ([]byte, error) {
	return bindataRead(
		_regoAbstractJson,
		"rego-abstract.json",
	)
}

func regoAbstractJson() (*asset, error) {
	bytes, err := regoAbstractJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-abstract.json", size: 9559, mode: os.FileMode(420), modTime: time.Unix(1529624640, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoAttributeIn = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\x4f\x6f\xdb\x36\x14\xbf\xf7\x53\x3c\x18\x05\x0a\x6c\x4d\xb1\x01\xc3\x80\xe6\x96\x43\x0f\x01\x76\x28\x86\x9d\x9a\xa6\xc6\xb3\xf8\x64\x71\xa1\x48\xf6\x91\x72\xec\x0d\xf9\xee\x03\x29\xc9\x96\x12\xdb\x92\x25\xc5\xee\x74\x49\x20\x3d\xfe\xf4\x7e\x7c\xff\x29\x43\xe3\x9a\xf9\x8d\xa5\xd9\x35\xcc\xcc\xe2\x6f\x4a\xfc\xec\xfd\x9b\xd6\x53\x41\x2e\x61\x69\xbd\x34\x3a\x08\xfd\x21\x9d\x07\x93\xc2\x8a\xd8\x49\xa3\x49\x00\x7a\xcf\x72\x51\x78\x72\x1f\x9e\xaf\x45\x21\x64\x58\x88\xea\x33\x1b\x4b\xec\x25\xb9\xd9\x35\xa4\xa8\x1c\x3d\x13\xb5\xe8\x3d\xb1\x6e\xc9\xfd\xdb\x12\x89\x62\xdf\x56\x77\xbf\x5c\x7d\xbc\xff\xf9\xeb\xdb\xbd\xcf\x5b\x7c\x90\x19\x37\xcf\x54\xea\xa2\xb5\x23\x03\x52\x03\x82\x45\xf6\x32\x29\x14\x72\xcd\xf8\x03\xdc\x7a\x78\x94\x4a\x41\xe1\x08\x50\xa9\xe6\x9a\x94\x4d\x0e\x3e\x23\xb0\x4c\x2b\x69\x0a\x57\xaf\x72\x80\x5a\x00\x0a\x01\x86\xc1\xac\x88\x1f\x59\x7a\x8a\xa2\x46\x93\x03\x41\x89\x42\x26\xf1\x7c\x0b\xb7\xfa\x4a\x4f\xf9\xfe\x3d\xe9\x67\xc7\x63\xe4\xff\x24\xcb\xe4\x48\xfb\xa0\xe5\x8e\x4e\xd8\x8f\xa0\x20\xd7\x8f\x49\x40\x89\x7d\x48\xcb\x08\xce\xf4\xbd\x90\x4c\x62\x76\x0d\x77\x07\xa5\xa2\xa4\xc6\x9c\x8e\x20\xed\x48\x75\xc8\x34\xe9\x1c\x94\xbc\x3f\xa2\x72\x7f\x3f\x6d\x2d\xb3\xc7\x9d\xb5\xfd\x0a\xa5\xcc\x23\x89\x79\x92\x21\x77\x8b\xbf\xa0\x15\xad\xb4\xa4\xb5\x05\x9f\xa1\x07\x04\xe7\x59\xea\x65\xc3\x58\x79\xe1\x3c\x64\x46\x1b\x06\x6f\x60\x41\xb0\x42\x25\x45\xc7\xd6\x41\xd3\x6f\x4a\xc8\xc3\x3b\x18\xae\xa7\x0e\x53\xec\x68\x1a\x99\xf4\xd8\x97\x7d\x44\x6f\x4a\x8c\x40\xa0\x08\x31\x65\xf8\xa5\x5b\x6e\x2c\xc1\x27\x5d\xe4\xa7\xf0\x3b\x96\x0f\x5a\x0b\xba\x63\x6d\x2f\x7e\xb5\x7f\xdd\x2f\x80\x46\xc2\x0b\x0b\xbf\xdd\xdd\x5c\x7d\xf9\xf5\xea\xe3\xfd\x1d\x5e\xfd\x73\x73\xf5\x25\x64\xb8\x9f\xbe\xbe\x3d\x6e\x89\x70\x3d\x8d\xb3\x55\xe1\xcd\x92\x34\x31\xfa\x18\xad\xa7\x5b\xea\xaf\x8c\x1a\x76\x91\x0e\x5a\x90\xb0\xd8\xc4\xfc\xb1\xc0\xe4\x81\xf4\x49\xae\xb8\x30\x46\x11\x1e\x89\xe6\x3e\xfc\x12\x26\x0c\x9a\xce\x8d\x56\x9b\x09\xf8\x25\xa8\x21\x40\x85\xe0\x72\xe4\x41\x14\x31\x04\xe3\x6b\x08\x42\x2a\x88\xaf\x3b\x3b\x4f\x41\x29\x16\xca\xcf\x0d\x0b\xe2\xc1\x3c\x63\xb4\xd5\x29\xbf\x65\xd4\xc2\x85\x2a\x1f\x6a\x54\x7c\x0f\xc4\xf7\x04\xe6\x0f\xd4\x2b\x9c\x5e\x85\x6c\x54\x77\x30\xd9\x9a\xca\x7e\xd2\x27\x90\xba\xab\xb3\xca\x8e\xde\x7b\x98\x49\xed\x69\x49\x1c\xfe\xd5\x85\x52\xe5\xdf\x7c\x51\xde\xa9\xcb\xf3\x36\x61\xdc\x8f\xdc\x10\xcb\x94\x4c\x17\xc2\x0d\xbc\xf3\xdb\xb6\xa9\xda\x50\xcb\x6e\x6f\x8d\xb1\xeb\x24\xc5\x90\xd6\x98\x5b\x45\x23\x7c\xf5\x53\x89\x50\xf9\x69\xa8\x84\xff\x5f\x47\xa5\xb5\x35\x6e\x2a\x2f\xb5\xc5\x42\xc9\x44\x6d\xa0\x46\x1d\xe0\xab\x3d\x96\x54\x79\x22\xac\xf2\x5c\xd0\x38\x77\x48\xa5\xf2\xc4\xb8\x50\xc3\xf3\x56\xbb\x18\x2d\xa8\xcc\xcd\xc1\x2f\x4a\xf0\x90\x95\x6d\xc1\x61\x4b\xdc\xd9\xc3\x37\x35\x4c\x72\xa9\xe7\xa1\x2c\x4c\xd3\x4d\x40\x05\x79\x91\x4a\x93\x1a\xce\xd1\x0f\x66\x52\x2e\x1f\x94\x83\x28\x34\xb6\x5d\xd3\xd3\x4e\x4f\xa6\x3e\xa8\x25\x72\x8e\x52\xf5\xee\x50\x33\xa3\x7b\x23\xdf\x7e\x5e\xfd\x76\x82\xec\xef\x7d\x65\x13\x29\xb8\xaf\xac\x40\xdf\x11\xa2\xe1\x1a\x97\xc5\x96\x14\xda\xf6\x09\xfc\x3b\x1e\x21\x64\xb8\x22\x40\xd8\xb5\xcc\x15\xfc\xb9\x7d\x5d\x0a\xd2\x5e\xa6\x72\x12\x66\x8d\x96\x11\x21\x47\xa9\xa1\x01\xff\x3a\x79\xba\xd0\xf2\x7b\x41\xb7\xd5\xdc\x16\x52\xf5\xb8\xed\xd0\x82\xd6\x13\xef\x84\x86\x12\xf5\xdc\xa6\xcd\x71\x3d\x57\xa4\x97\x3e\x1b\x4c\x28\xc7\xb5\xcc\x8b\x1c\x4a\x98\x97\x33\x42\x3d\x90\xf7\x1e\x80\xb7\x0c\xeb\xfe\x63\x34\xc3\x71\xe3\x40\x4d\xf0\xc0\x0c\x54\xf3\xdb\x76\x4b\xe7\x26\x28\xf5\x68\x13\x4a\xfd\x43\x9b\x50\xea\xb1\x26\xac\x08\xfe\xa0\x26\x8c\x47\x9d\x43\xb9\x85\xc5\x17\x1f\x66\x4c\x2e\xfd\x9c\x72\xeb\x87\x35\x77\xb7\x69\x4c\xcb\xef\x9f\xd9\xa5\x6a\x63\x05\x1b\x6b\x49\x80\x4c\x21\xbe\x02\x0c\x43\x39\x90\x9c\x39\x5b\xc6\xb3\x8d\x57\x6a\xd1\xb7\xe7\x26\x17\xeb\xd0\x2d\xcb\x1c\x79\x33\x5d\x87\x5e\x56\x37\xa8\x70\x2f\xd2\xa6\x33\xa1\x98\xea\x84\x4f\x3a\x08\x70\xf1\x94\xef\x02\x44\xb6\xdf\x4d\x26\xe1\x51\xa1\x9d\x9b\x86\xa3\x84\x69\xf8\xd8\xd4\x22\x51\x62\xc5\x4f\x67\x2e\x33\x85\x12\x31\x57\x10\x2a\x0f\x8f\xd2\x67\x90\x20\x9f\x94\x07\x27\x22\xf8\xaa\x13\x80\xbb\xcc\x04\xe0\xbc\x99\xcc\xf7\x2a\xac\xd7\x3e\x92\x19\xdd\xe6\xbb\x62\x51\x29\x30\xae\x2e\xe7\x68\x6d\xc8\xeb\xde\x80\x95\xc9\x43\x4c\xf6\x8f\x19\xe9\x58\xeb\xca\xb6\xc3\xc5\x2f\x06\xde\xc0\x3b\x5a\x7b\x62\x8d\xea\xdd\xb9\x0b\xb8\x67\xd4\x4e\x92\x9e\x28\x36\x77\x70\xe7\x76\xd5\x51\x36\x8b\xe6\x68\x7e\xe8\x1f\x60\x86\xc9\x8f\x70\x4e\xfb\x72\xd8\xbf\x87\x8d\xe2\xa9\x32\xd8\xc7\x48\x51\xb8\x7f\x20\xc2\x96\x66\x4f\x59\x25\x5d\x6f\x35\x3a\x7f\xc2\xd0\x92\xf6\xb2\xf3\x77\x04\x5b\xd9\x3a\xfc\x46\x9f\x13\x1d\x7c\xba\xff\xc9\xcb\xbb\xed\x3b\x4f\x6f\xfe\x0b\x00\x00\xff\xff\x7f\x2e\xce\xf9\x94\x23\x00\x00")

func regoAttributeInBytes() ([]byte, error) {
	return bindataRead(
		_regoAttributeIn,
		"rego-attribute.in",
	)
}

func regoAttributeIn() (*asset, error) {
	bytes, err := regoAttributeInBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-attribute.in", size: 9108, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoInfoJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\xbd\x4e\xc3\x40\x0c\xc7\xf7\x3c\x85\x75\x30\xb6\x0d\x13\x43\x36\xc6\x0e\xa0\x4e\x2c\x88\xc1\x4d\x7c\xad\xab\x34\x77\xd8\x0e\x02\x55\x7d\x77\x74\xb9\xa6\x4a\x3f\x36\x96\x44\xf9\xd9\xff\x8f\xf8\x50\x00\x00\xb8\x47\xad\xb7\xb4\x47\x57\x81\xdb\x9a\xc5\xaa\x2c\x77\x1a\xba\x79\xa6\x8b\x20\x9b\xb2\x11\xf4\x36\x7f\x7a\x2e\x33\x7b\x70\xb3\xac\x34\xb6\x96\x92\x6e\xd9\xf9\x30\xc2\x86\xb4\x16\x8e\xc6\xa1\x4b\xa3\x57\x32\x04\xee\x7c\x90\x3d\x26\x06\xb8\x0e\xbd\x01\x82\x92\x41\xf0\xa0\x91\x6a\xf6\x5c\x0f\x43\x5d\x9c\xad\x7f\xe3\xe0\x1c\xd6\x3b\xaa\x6d\xa4\xd8\x34\x9c\xf6\xb0\x5d\x49\x88\x24\xc6\xa4\xae\x02\x8f\xad\xd2\x69\x45\xe8\xab\x67\xa1\xc6\x55\xf0\x31\x90\x81\x46\x21\xcf\x3f\x27\x9b\xbc\x17\x82\x4d\xbf\xbf\x49\x34\x55\x1e\xc8\xe7\xac\x28\x4e\xba\x49\xcc\xe1\xc6\x6f\xca\x2e\x6a\xab\x09\x77\x9b\x89\xff\xbd\xd3\xac\x06\x13\xb0\x00\xbd\x52\x7a\xd5\xd8\xb6\x80\x91\xd5\x9d\x75\xc7\xeb\xca\xff\x8b\x7c\xc3\x3d\xa5\xab\xdb\x96\x20\xd9\x5d\x9e\xff\x56\xec\xb1\x6f\x53\x66\xce\xbe\xdb\x6a\x3c\xdc\x4d\xb1\xab\xe8\xf7\xbc\x37\xa6\xbf\xac\x96\xd7\x71\xe3\x9f\x70\x67\xb4\x21\x99\xc4\x15\xf9\x79\x2c\xfe\x02\x00\x00\xff\xff\xb8\x06\xde\x01\xb3\x02\x00\x00")

func regoInfoJsonBytes() ([]byte, error) {
	return bindataRead(
		_regoInfoJson,
		"rego-info.json",
	)
}

func regoInfoJson() (*asset, error) {
	bytes, err := regoInfoJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-info.json", size: 691, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoSpecIn = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\xcd\x6e\xe4\x36\x0c\xbe\xe7\x29\x08\x6f\x8f\xc9\x4e\x4f\x3d\xcc\x6d\xfb\x03\xb4\x40\x0b\x14\x69\x7a\x2a\x16\x03\x8d\x44\xcf\x68\x57\x23\x79\x29\x3a\x69\xb0\x98\x77\x2f\xe4\x7f\x8f\x65\xc7\xde\x38\x09\x7a\x09\x10\x8a\xa4\x3e\x7e\x24\x4d\x6a\xbe\x5e\x01\x00\x24\xdf\x79\x79\xc4\x93\x48\xb6\x90\x1c\x99\xb3\xed\x66\xf3\xc9\x3b\x7b\x53\x4a\xdf\x3b\x3a\x6c\x14\x89\x94\x6f\xbe\xff\x61\x53\xca\xde\x25\xd7\xa5\x25\x6b\x36\x18\xec\xfe\xca\x50\xea\x54\x4b\xc1\xda\xd9\xfa\x54\xa1\x97\xa4\xb3\x42\xb4\x85\xe4\x16\x33\x42\x8f\x96\x3d\x08\x0b\x6e\xff\x09\x25\xc3\x83\xe6\x23\x08\x63\x40\x07\x31\x33\xe9\x7d\xce\x18\x34\x14\x10\x9a\xc2\x9f\x6f\xae\x7b\xcc\x8a\xdb\x4a\xdb\x5a\x2a\x94\xd2\x41\x4d\x98\x3f\xc9\x65\x48\xac\xd1\x27\x5b\x48\x85\xf1\x58\xa9\x10\x7e\xc9\x35\xa1\x4a\xb6\xf0\x4f\x21\x29\xa4\x27\xa7\xd0\x24\xc5\xff\x1f\x2b\xc5\xac\xeb\xe1\x6b\xab\xda\x22\xeb\xc9\x4b\xfa\x08\xd3\x80\xea\xdd\x46\x61\xaa\x6d\x81\xc5\x6f\x3a\x16\x8d\xfa\xf9\xba\xf5\xd8\x06\x37\x70\x58\x87\x29\x88\xc4\x63\x72\xdd\x3f\xbc\x20\xf5\x77\xed\x19\x5c\xda\x72\x05\x7b\xe4\x07\x44\x0b\x7c\x44\xa0\x9a\x72\x54\x35\xe1\x81\x58\xc7\x47\x24\x70\x16\xfd\xfb\x4b\xf7\x9a\xf1\x34\x84\x34\x11\x67\x7d\x71\xd2\x33\x38\x47\x63\x2e\x09\x9f\x49\x60\x27\x3b\xad\xc7\x73\x53\x5a\x8d\x62\x3f\x51\x23\x57\x34\x85\xfa\x47\x71\x3e\xcd\xe9\xaf\xce\x28\x0f\x7b\xe1\xb5\x04\x6d\x53\x47\xa7\x22\xc2\x40\x73\x9c\xd4\x01\x8b\xd1\x42\x6d\x4e\x9f\x2e\xd8\x4e\x95\x44\x0a\xb7\x39\xcd\x84\xfc\x2c\x0e\x78\xe1\xbf\x32\xf4\xbc\xb3\xe2\x34\x76\xe8\x72\x92\x38\xaa\x80\x96\x35\x3f\x8e\x1e\x77\xf9\xea\x9d\x7e\xbc\xbe\xea\xe3\x1f\xe9\xa7\x96\x0a\xa3\x85\x1f\x39\x1c\x5c\xd5\x29\xf7\xca\x0e\xd8\xb5\xe9\x00\x01\xbe\xfb\x19\xba\xcc\x4a\xe3\x73\xaa\xbf\x1a\xa5\xf1\x46\x18\xf8\xf1\x4c\xda\x1e\x92\xa8\xe2\x79\x20\x3d\x47\x08\x95\x84\x82\x71\x36\x0b\x77\xf1\xd6\x96\xc2\xc2\x1e\xa1\x74\xa6\x9e\x0c\x7f\xef\x9c\x41\x61\x87\xb8\x63\x08\x15\x1a\x5c\x86\x70\x1c\x60\xe9\x6b\x75\x80\x5d\x00\xf3\x79\xec\x88\xe6\x37\xf8\x00\xea\x58\x09\xc4\x90\x76\x9b\x6b\x2e\xd2\x9f\xc4\x09\xcd\x8d\x14\x1e\x15\xdc\x23\xf9\x1e\xda\xaa\xd3\xd7\x05\xf9\x2f\xa3\x55\xcb\x1b\x73\x2f\x3c\xf6\xdb\x10\xf8\xa8\xfd\x85\xa8\xf2\x0e\xff\xab\x26\x3d\x20\xaf\xd4\xa1\x84\x4c\x1a\xef\xd7\x6e\x81\x7a\x1c\x2c\x01\x19\x0a\xa7\xae\xa4\xca\xfe\x95\xda\x20\x23\x7d\xbf\xc2\x47\x4f\x7b\xa8\x3c\xad\x4b\x66\x7f\x46\x2e\x01\x99\x99\x9c\x84\x79\x9d\x36\x6d\xc7\xfc\x12\x84\xb9\xd5\x5f\xf2\x2a\xf7\x0d\xa5\xda\x1e\x0a\xa0\x2f\x90\x6a\x72\xee\xd9\xad\x23\x64\x78\x20\xf8\x92\x4b\xe7\x78\xe1\xac\x5f\x94\xfb\x3c\x53\xeb\x8d\xe3\xd2\xd9\x73\x5a\x7d\xc6\x66\x7d\xf1\x3e\xd9\xed\x3e\xdc\xdd\xdd\xfe\xf6\xe3\xdf\x77\xbf\xec\x76\x93\xcf\x8f\x89\x55\xf9\xb6\x56\x99\xde\x96\x3f\x18\xe3\x1e\x8a\x4d\x4c\xa1\x34\x82\x10\x44\xf3\x1a\x69\x1f\x23\x0f\xae\x9f\xb0\xe1\xcb\xe3\x95\x76\xe6\xc9\xc5\xb8\x13\x9a\x7f\xe6\x66\xfb\x4d\x2b\x9d\x11\x6d\xfd\xf8\x8b\x7d\x0e\x52\x72\xa7\xa2\xfc\x65\x4e\x14\xf6\xdd\x61\xcd\xbd\xf1\xca\x37\x16\x41\xb5\xf0\xbd\x41\x04\x19\xa1\x0c\x88\x96\x47\x11\xaa\x57\x7b\xe8\xb8\x78\xa1\x6d\x75\x62\xc7\x8a\x3f\x8c\xc7\x8b\x74\xec\xb6\xe5\x9b\x4b\x34\x8d\xcd\xda\xf2\xea\x89\x5c\x3e\xe6\x1a\x0b\xc8\xc3\xce\x2c\x3c\x08\x30\xda\x7e\x0e\xdf\x29\x1e\x84\xb8\xe6\xb0\xfb\xa6\xe1\x11\xe5\xbb\x9a\x1c\x2f\xcb\xf6\x8c\xe1\x32\x59\xac\xc3\x79\xd1\x7d\x52\x79\xc8\x90\xc2\xec\x7e\x7a\x8e\xfc\xec\x64\x7e\x42\xcb\x17\xbf\xb5\x98\xe8\x7c\x5f\x6f\x5a\xac\xfc\x1d\xef\x47\x91\x3a\x2a\x13\x57\x38\x81\x70\xd3\xbc\x75\x65\x49\xc1\x2d\xfc\x4e\xc7\x11\x96\x4e\x5a\x84\x6b\x02\x5c\xf2\xf9\x89\xa3\x3b\x20\xbf\x0c\xb4\x85\xcd\x1a\x47\x57\x3a\x59\x01\xe0\x58\x2f\x5e\x95\x7f\xcf\x57\xff\x05\x00\x00\xff\xff\x88\xdb\xad\xc1\xb1\x17\x00\x00")

func regoSpecInBytes() ([]byte, error) {
	return bindataRead(
		_regoSpecIn,
		"rego-spec.in",
	)
}

func regoSpecIn() (*asset, error) {
	bytes, err := regoSpecInBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-spec.in", size: 6065, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoSpecJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5b\x5b\x6f\xdb\x3a\x12\x7e\xcf\xaf\x18\xb8\x05\x0a\xec\xc6\x4d\x17\x58\x14\x68\xde\x8a\xdd\x02\x1b\x60\x17\x28\xba\xe7\xa9\x41\x1a\xd0\xe2\xc8\x66\x43\x91\x2a\x49\x39\xf6\x29\xf2\xdf\x0f\x48\xea\x6a\xdd\x2d\xd9\x6e\xfb\x70\x4e\x2b\x0d\x47\xf3\xcd\x95\x43\x8e\x7f\x5e\x01\x00\x2c\x5e\xeb\x60\x83\x11\x59\xdc\xc2\x62\x63\x4c\x7c\x7b\x73\xf3\x5d\x4b\xb1\xf4\x4f\xdf\x4a\xb5\xbe\xa1\x8a\x84\x66\xf9\xee\xfd\x8d\x7f\xf6\x6a\x71\xed\x57\x1a\x66\x38\xda\x75\xff\x8f\x31\x60\x21\x0b\x88\x61\x52\x64\x6f\x29\xea\x40\xb1\xd8\x3d\xba\x85\xc5\x17\x8c\x15\x6a\x14\x46\x03\x11\x20\x57\xdf\x31\x30\xf0\xcc\xcc\x06\x08\xe7\xc0\xec\x63\x63\x14\x5b\x25\x06\x2d\x05\x05\x85\xdc\xf1\xd3\xf9\xe7\xf6\xb1\xfb\x9a\x5f\x9b\x3d\x25\x94\x32\x4b\x46\xf8\x67\x25\x63\x54\x86\xa1\x5e\xdc\x42\x48\xb8\xc6\x94\x44\xe1\x8f\x84\x29\xa4\x8b\x5b\xb8\x77\x4f\xdc\xd3\x48\x52\xe4\x0b\xf7\xef\x87\x94\x30\x2e\x73\xf8\x59\x90\x16\x92\x55\x9e\x7b\xf5\x29\x0c\xad\x54\xaf\x6e\x28\x86\x4c\x38\x59\xf4\x4d\x69\x45\x4e\xfe\x72\x5d\x70\x2c\xc0\xd5\x18\x66\x30\x89\x52\x64\xbf\xb8\xae\xbe\x3c\x50\xea\x7f\x99\x36\x20\xc3\x42\x57\xb0\x42\xf3\x8c\x28\xc0\x6c\x10\x54\xa6\x72\xa4\x99\xc2\xad\x62\xa5\xd9\xa0\x02\x29\x50\xbf\x3d\x64\xcf\x0c\x46\x75\x91\x3a\x70\x66\x1f\x5e\x54\x16\xbc\x34\x62\xf6\x0a\x1f\xa8\xc0\x92\x75\x0a\x8e\x2f\xb9\x6b\xe5\x84\x55\x43\xb5\x7c\x22\x77\xd4\xff\xb9\xf7\xdd\x3a\xfd\x8f\xe4\x54\xc3\x8a\x68\x16\x00\x13\xa1\x54\x91\x43\x68\xd5\xdc\xac\xd4\x9a\x16\x1b\x1d\x35\x7f\xdb\xef\xb0\x25\x2f\x69\x70\xdc\xfc\x6d\x4c\x82\x27\xb2\xc6\x03\xfe\xe9\x42\x6d\x1e\x05\x89\xda\x5e\xca\x44\x05\xd8\x4a\x80\xc2\x30\xb3\x6f\x7d\x5d\xd6\x57\xe5\xed\xc3\xf5\x55\x55\xfe\x96\x78\x2a\x54\xc1\x19\xd1\x2d\x2f\x6b\x9f\x2a\xb9\x7b\xba\x0e\x8c\x2c\xcc\x01\x04\x74\x39\x0d\x1d\x5a\x25\xe7\xd9\x15\x5f\x39\x51\x7b\x20\xd4\xf8\x68\xa3\x98\x58\x2f\x1a\x09\x5f\x6a\x4f\x5f\x1a\x14\x1a\x28\x24\x06\x07\x6b\xe1\x8f\xe6\xd0\x0e\x88\x80\x15\x82\x67\x46\x7b\xe1\xaf\xa4\xe4\x48\x44\x5d\xee\x26\x09\x29\x72\x1c\x27\x61\xbb\x80\x9e\xd7\xec\x02\x96\x05\x18\xae\xc7\xd2\xa3\xe1\x01\x5e\x13\xb5\xcd\x05\x9a\x24\x2d\x07\xd7\x50\x49\xff\x45\x22\xe4\xcb\x80\x68\xa4\xb0\x45\xa5\x2b\xd2\xa6\x91\x3e\xaf\x90\x3b\x83\x82\x8e\x0f\xcc\x15\xd1\x58\x0d\x43\x30\x1b\xa6\x0f\x1e\xa5\xdc\xe1\xb7\x0a\xd2\x35\x9a\x99\x22\x54\xa1\x51\x0c\xb7\x73\x87\x40\x56\x0e\xc6\x08\x69\x1d\x27\xf3\xa4\x74\xfd\x99\xc2\x20\x56\x6c\x3b\x43\xd2\x63\x1a\x52\x4e\xf3\x2a\xb3\x5a\x23\xc7\x08\x19\xf3\x44\x11\x7e\x9e\x30\x2d\xca\xfc\x18\x09\x13\xc1\x7e\x24\xa9\xed\x73\x95\x32\xb1\x76\x82\x9e\xc0\xd4\x4a\xca\xc9\xa1\x43\x02\xdb\x20\x68\xaf\x4b\x29\xcd\xc8\x5a\x3f\xca\xf6\x49\x4c\xe7\x2b\xc7\x9e\xd9\x94\x50\x1f\xb0\xb3\xee\xea\x4f\xba\x37\xa3\x2d\x79\x3c\xf5\x5e\xa4\xa5\xa6\xac\xb6\xcd\x1d\xb1\x91\x8d\x89\x31\xa8\xc4\xe7\x9e\xfd\xe0\xb7\xed\xfd\xbb\xe5\x87\x87\xbf\xbf\x6e\x57\xfe\xa0\xc2\xd0\xb6\x6d\x2c\x1a\x4c\x26\x80\x40\x4c\x94\x61\x41\xc2\x89\xca\x00\xbf\x85\x3b\xdb\x94\x72\x0e\x89\x46\xd7\x98\x96\xd6\x84\x4a\x46\x3e\x53\x2a\xdc\x32\x99\xe8\x6c\x95\xef\x57\x09\xa5\x20\x15\xc8\x2d\xaa\x67\xc5\x0c\xfa\x70\x12\xa8\x81\x62\xc0\x89\xea\x70\x82\x11\x85\xac\xd1\x8c\x5d\xe0\xab\x7d\x77\x0e\x67\x7c\x9a\x87\xde\x8e\xa4\x42\xd9\xd2\x3d\xd4\x41\xf5\xd0\xb4\x76\x1b\xe5\x3f\x0f\x1d\x22\x0f\x77\xd3\xca\xb2\x9e\xde\xa5\xfa\x09\xce\xe5\x33\xd2\xc7\x60\x43\x54\x3f\x79\x0d\x96\xb3\xd2\x1a\x77\x31\x98\x0d\x71\x8d\x8c\xcb\xae\x25\x63\x45\x89\x36\xb0\x91\x42\x2a\xdb\xf3\xac\x10\xb6\x84\x33\xda\xa3\x3a\x18\xba\x01\xca\xfe\x34\x24\xc3\x16\x98\x92\x05\x03\xf4\xd2\x04\xf4\xa3\xe7\x61\x01\x24\x36\xa6\xa4\xaa\xbb\xe5\x3e\x46\xf8\x24\x92\x68\x0c\xbe\xae\x7c\x50\x59\xd0\x1f\x6b\x8d\xfc\x53\xfd\xf5\x7f\x00\x4a\xf9\xce\x2e\xfc\x76\xff\x71\xf9\xf5\x1f\xcb\x0f\x0f\xf7\x64\xf9\xe7\xc7\xe5\x57\x9b\xe0\xfe\xf6\xba\xdb\x10\xd0\xb8\x2b\xad\xbc\xed\x33\x55\x62\xe4\x1a\x05\x2a\x5b\x7c\x8e\x32\x94\x2d\x6b\x85\x59\x98\x86\x0a\x4b\x58\xed\x5d\xfa\x58\x91\xe0\x09\xc5\x28\x4f\x6c\xad\x74\x63\xf0\xb9\x2e\x97\x49\xf1\x28\x05\xdf\xcf\x80\xcf\x56\x6b\xcb\xca\xc6\x96\x46\x03\x34\x71\x11\xe8\x9b\x69\xb0\x99\xa0\x7c\xb6\x79\x3e\x9c\x14\x43\x92\x70\xf3\x28\x15\x45\x75\x34\x4e\x17\x6c\x59\xc6\xaf\x18\x35\xb1\x2d\x25\xb1\x25\xca\x7d\x07\xdc\x77\x2c\xf2\x27\x1c\x14\x4d\x27\x01\xeb\xc4\x3d\x1a\x6c\x06\xa5\x19\xf4\x08\x50\xf7\x59\x52\x29\xe0\x5d\xc3\x82\x09\x83\x6b\x54\xf6\xaf\x22\xe1\xdc\xff\x3f\x5a\xf9\x27\x59\x75\xce\xf3\xc5\xc3\x44\x85\xc4\x0a\x83\xf9\x42\xb8\xc4\xef\xfc\xb6\xed\x3f\xa0\xe9\xc3\xd2\x70\x58\x73\x8c\x5d\x67\xa9\x85\xb8\x23\x51\xcc\x71\x82\xaf\x7e\xf2\x1c\x52\x3f\xb5\x85\xf0\xf7\x75\x54\xdc\xc5\x52\xcf\xe5\xa5\x71\xb2\xe2\x2c\xe0\x7b\xc8\xb8\x1e\xe1\xab\x03\x96\xa4\x79\xc2\xae\x32\x2a\xc1\x69\xee\x10\x32\x6e\x50\x91\x15\x3f\x3e\x6f\x55\x8b\x91\x6d\x1d\x6d\x6e\xb6\x7e\xe1\x99\xdb\xac\x1c\x27\xca\xaa\x44\x9f\x3d\x7c\x43\xa9\x90\xad\xc5\xa3\x2d\x0b\xf3\xec\x26\x20\x65\x79\x91\x4a\xe3\x6f\x75\x8e\x46\xe2\x97\x1f\x95\x83\xd0\xee\x6b\xfb\x9a\xa7\x42\x4e\x85\x43\xb8\x7a\xce\x11\x61\x87\x77\x5b\xad\xc4\xf1\x46\x8a\xc1\x9c\xef\x3e\x6f\xff\x39\x82\xf6\xfd\x50\xda\x80\x51\x35\x94\xd6\x1d\xc9\xf4\x92\x4e\xcb\x62\x6b\xb4\xbb\xf6\x19\xfc\xdb\x9d\x20\x6c\xc8\x16\x81\x40\xb1\x65\x4e\xd9\x9f\xdb\xd7\x19\x45\x61\x58\xc8\x66\x41\x56\xda\x32\x12\x88\x08\x13\x50\x62\x7f\x9a\x3c\xed\x8f\x2b\xef\xd2\xb6\xcd\xa6\xea\x69\xea\x10\x14\x77\x33\x6b\x42\x80\xe7\x7a\x6e\xd3\x46\x64\xf7\xc8\x51\xac\xcd\xe6\x68\x40\x11\xd9\xb1\x28\x89\xc0\xb3\xa9\xf7\x08\x59\x3f\x3e\xb8\xff\xcd\x11\x66\xfb\x8f\xc9\x08\xa7\xb5\x03\x19\xc0\x96\x1e\x28\xc3\x97\xef\x96\xce\x0d\x90\x89\xc9\x26\x64\xe2\x97\x36\x21\x13\x53\x4d\x98\x02\xfc\x45\x4d\xd8\x79\x01\xd3\x87\xad\x7c\x03\x77\xb1\x66\x46\x46\xcc\x3c\x62\x14\x9b\xe3\x36\x77\x77\xa1\x4b\xcb\xd7\x07\x76\xc9\xee\xfb\x95\x8c\x63\xa4\xc0\x42\x70\x9f\x00\xa9\xc0\x37\x24\x67\xce\x96\xee\x6c\xe3\x44\x5b\xf4\xfc\xdc\xe4\x62\x3b\xf4\x58\xb1\x88\xa8\xfd\x7c\x3b\x74\x5f\xdd\x20\xe5\x7b\x91\x6d\xba\x42\x42\xe7\x3a\xe1\x63\x1a\x2c\x3b\x77\xca\x77\x01\x20\xf9\xb5\xc9\x2c\x38\x52\x6e\xe7\x86\xa1\x31\x50\x1d\x63\x10\xa3\x40\x78\x5e\xee\xe6\x4c\x6f\x64\xc2\xa9\x9f\x0d\x22\x3c\x9d\x0a\x0d\x88\x1a\x95\x07\x67\x02\x78\xd2\x0e\x40\x5f\xa6\x03\xd0\x46\xce\xe6\x7b\x29\xaf\x53\x1f\xc9\x4c\xde\xe6\xeb\x64\x95\x0a\x30\xad\x2e\x47\x24\x8e\xdd\x6c\x84\x84\x98\x05\x4f\x2e\xd9\x3f\x6f\xd2\x01\x5b\xbf\xed\xd0\xee\xc6\xc0\x48\x78\x83\x3b\x83\x4a\x10\xfe\xe6\xdc\x05\xdc\x28\x22\x34\x43\x31\x53\x6c\x16\xec\xce\xed\xaa\x93\x6c\xe6\xcc\x51\xbe\xe7\x3f\xc2\x0c\xb3\x1f\xe1\x8c\xbb\x38\x1c\xbe\x87\x75\xe4\x21\x97\x64\x88\x91\x1c\xf1\xf0\x40\x84\x1c\xe6\x40\x5a\xce\xf4\x60\x31\x7a\x27\x18\x2a\xd4\x86\xf5\x8e\x11\xe4\xb4\x59\xf8\x4d\x3e\x27\x6a\x7d\xdb\xfc\xa6\x61\x80\xf0\xaa\xf9\x5f\x4d\x3f\x01\xe8\x18\x57\xff\x92\x91\x74\x4f\xed\xb8\x3b\x75\x37\x0d\x9d\x4e\x9a\x00\xc9\x7f\x11\x50\xfc\x20\xe0\x59\x56\x87\xa6\xea\x03\x3d\x67\x9a\x5b\xef\x1c\x4e\x2f\x41\xd3\x13\xa7\xcb\x8f\x1a\xab\xe6\xa4\x98\x84\xd1\x07\x33\xd5\xc5\x04\x50\x90\x28\x85\xc2\x34\x4c\xcf\x5c\x78\xec\xba\x0d\x41\x3a\x74\x7d\x01\x04\xbd\x37\x98\xad\x28\xac\xf7\x0e\xb9\xb4\x9c\x3c\x31\xde\x31\xe7\xdc\xfc\xe3\x94\x76\x27\x6d\xfb\xda\xf8\xe9\xe1\x46\x33\xe6\xa3\xc3\x67\x37\xe4\xf8\x51\xd3\x7c\x45\xe9\xc4\x98\x33\xf1\x64\xf3\x94\xa9\x41\x9c\x73\xe0\xf4\xa8\x01\xce\x46\x7d\xa7\xd3\x9b\xa7\xd5\xf6\x80\x5a\xd1\xe9\xac\xf5\x7a\x51\xbe\x29\xd7\x10\xa3\x02\x12\x0c\xa8\x23\xff\x96\x41\x12\xa1\x30\x07\xbf\x77\xe2\x8d\x33\xb6\xf3\x55\x8b\x99\xf3\x78\x15\x45\x76\xaf\x7e\x38\xc9\x33\xeb\x84\xf3\xc8\x3c\xdd\x2c\xa1\x67\xd2\x3b\x6b\x74\x94\x80\x63\xd2\x4f\xb3\x74\x6b\x34\xa7\x11\x6d\x64\xb0\x36\x4b\xe7\x99\xcc\x20\x60\x5b\x2c\x5e\xf9\xff\xbe\x5c\xfd\x15\x00\x00\xff\xff\xb9\x1d\x5d\x75\x35\x3b\x00\x00")

func regoSpecJsonBytes() ([]byte, error) {
	return bindataRead(
		_regoSpecJson,
		"rego-spec.json",
	)
}

func regoSpecJson() (*asset, error) {
	bytes, err := regoSpecJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-spec.json", size: 15157, mode: os.FileMode(420), modTime: time.Unix(1529624640, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoTypeMappingJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\xc1\x8e\xd3\x30\x10\x86\xef\x7d\x8a\x51\x16\x09\x09\xb1\x0d\x27\x0e\xbd\x23\x71\x41\xe2\xc0\x0d\x71\x98\xc6\x93\x74\x76\x1d\xdb\x8c\x27\x5d\x0a\xea\xbb\x23\x3b\x6d\xba\x49\xb3\xed\x82\xd4\x9e\x62\xcf\x8c\xe7\xf7\xf7\x4f\xfd\x67\x01\x00\x50\xbc\x89\xd5\x86\x5a\x2c\x56\x50\x6c\x54\xc3\xaa\x2c\x1f\xa2\x77\xf7\xfd\xee\xd2\x4b\x53\x1a\xc1\x5a\xef\x3f\x7c\x2c\xfb\xbd\xbb\xe2\x7d\x5f\xa9\xac\x96\x52\xdd\xb7\x5d\x20\xf8\x82\x21\xb0\x6b\xe2\x31\x6a\x28\x56\xc2\x41\xd9\xbb\x9c\xb3\xe1\x08\x35\x5b\x02\x8e\xd0\x45\x32\xa0\x1e\x5a\x0c\x40\xbf\x94\xc4\xa1\x05\xdd\x05\x8a\x69\xd7\xa2\x6b\x3a\x6c\x08\x62\xa0\x8a\x6b\xae\xe2\x72\x68\xb9\x0b\xb9\xa3\x5f\x3f\x50\xa5\xc7\x5d\x34\x86\x53\x1f\xb4\x5f\xc5\x07\x12\x65\x8a\xc5\x0a\x6a\xb4\x91\x0e\x29\x01\x35\xb5\x19\xc5\x7b\x00\x39\xbc\x7c\x37\x5a\xf7\x60\x84\xea\xd4\xeb\xae\x34\x54\xb3\xcb\x1d\x62\x99\x15\x0c\x89\xfb\xfc\xb5\x1f\xae\x3c\xe4\x8d\x8f\x3f\xc8\x9e\x34\x18\xf1\x3b\xdc\x65\x08\xce\xe1\x7b\xa4\xdd\x73\x7a\x2a\xe8\xa2\xc5\x94\x02\xe8\xc6\x20\xdf\x46\x88\xdd\x3a\x7d\x41\x2d\xbe\x05\x1c\x60\xe6\xfc\xe5\xb4\xdd\x2c\xd8\x21\x7a\x1d\xf0\x90\x7a\x19\xf4\x25\xe0\x57\xc0\xb7\xfd\x7c\x15\x67\x35\xfb\xc5\xfc\x6a\x7f\x52\x56\x1c\x8b\x5f\xf6\xe0\x30\xbe\x57\x6d\x20\xa8\x50\xa9\xf1\x72\xb2\x62\xbd\x83\x86\x1c\x09\xaa\x97\x08\xec\xc0\x8b\x21\x79\xee\x10\x4d\xa7\x9c\x5d\x8a\x6e\x88\x05\xfc\x93\x03\x6e\x83\xa5\x96\x9c\xde\xd8\x1c\xa1\x9f\x1d\x0b\x99\x62\x05\xdf\xcf\x3d\x19\x8f\x76\xfa\xfd\x98\x9a\x7b\xc5\xd5\xd9\x39\xbf\xc4\x72\x00\x64\x32\x98\xe9\xcd\xcf\x08\x44\x95\xf9\x19\x38\xaf\x2b\xd2\xe4\xbc\x5a\xcc\x67\xff\x94\x0c\xcb\xd3\x86\x96\x7f\x53\x72\xe7\x06\x9a\xda\xe0\xe5\xf5\xaa\x3e\x6d\xc9\x69\x87\x16\x4e\x1e\x43\x7f\x04\x78\x81\x80\xd5\x63\x7a\x26\x75\x83\x0a\x41\xfc\x96\x4d\x7a\x40\x6f\xa1\x7b\x8b\x96\x0d\x1e\x64\xfd\xa3\x76\x87\x2d\x81\xaf\x01\xa1\xea\xa2\xfa\x16\x4e\x87\x41\xdd\xb9\x2a\x17\xfc\xaf\xd6\x97\xfe\xfb\xfd\xbb\xbc\xd8\x2f\xfe\x06\x00\x00\xff\xff\x74\x2b\x12\xf9\xe6\x06\x00\x00")

func regoTypeMappingJsonBytes() ([]byte, error) {
	return bindataRead(
		_regoTypeMappingJson,
		"rego-type-mapping.json",
	)
}

func regoTypeMappingJson() (*asset, error) {
	bytes, err := regoTypeMappingJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-type-mapping.json", size: 1766, mode: os.FileMode(420), modTime: time.Unix(1529622242, 0)}
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
	"gen.sh": genSh,
	"rego-abstract.in": regoAbstractIn,
	"rego-abstract.json": regoAbstractJson,
	"rego-attribute.in": regoAttributeIn,
	"rego-info.json": regoInfoJson,
	"rego-spec.in": regoSpecIn,
	"rego-spec.json": regoSpecJson,
	"rego-type-mapping.json": regoTypeMappingJson,
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
	"gen.sh": &bintree{genSh, map[string]*bintree{}},
	"rego-abstract.in": &bintree{regoAbstractIn, map[string]*bintree{}},
	"rego-abstract.json": &bintree{regoAbstractJson, map[string]*bintree{}},
	"rego-attribute.in": &bintree{regoAttributeIn, map[string]*bintree{}},
	"rego-info.json": &bintree{regoInfoJson, map[string]*bintree{}},
	"rego-spec.in": &bintree{regoSpecIn, map[string]*bintree{}},
	"rego-spec.json": &bintree{regoSpecJson, map[string]*bintree{}},
	"rego-type-mapping.json": &bintree{regoTypeMappingJson, map[string]*bintree{}},
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

