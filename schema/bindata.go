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

	info := bindataFileInfo{name: "gen.sh", size: 243, mode: os.FileMode(493), modTime: time.Unix(1521698190, 0)}
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

	info := bindataFileInfo{name: "rego-abstract.in", size: 467, mode: os.FileMode(420), modTime: time.Unix(1522433994, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoAbstractJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\x4d\x6f\xdb\x38\x13\xbe\xe7\x57\x0c\xdc\x02\x05\xde\xb7\xae\xbb\xc0\xa2\x40\x73\xcb\xa1\x87\x02\x7b\x28\x16\x7b\x6a\x90\x06\x23\x71\x64\xb1\xa5\x48\x75\x48\x39\xf6\x2e\xf2\xdf\x17\x24\x25\x59\x4e\xfc\x21\x5b\x8a\xdd\xed\xa5\x2e\x39\x7c\x38\xcf\x7c\x71\x48\xf5\x9f\x2b\x00\x80\xc9\x6b\x9b\xe6\x54\xe0\xe4\x1a\x26\xb9\x73\xe5\xf5\x6c\xf6\xdd\x1a\x3d\x8d\xa3\xef\x0c\xcf\x67\x82\x31\x73\xd3\xf7\x1f\x66\x71\xec\xd5\xe4\x6d\x5c\xe9\xa4\x53\xe4\xd7\xdd\x24\xd6\x31\xa6\xae\x99\x10\x64\x53\x96\xa5\x93\x46\x87\x69\x0d\x58\x4b\x00\x2a\x65\x1e\x2c\x38\x03\x82\x52\x85\x4c\x80\x60\xc9\x81\xc9\x20\x41\x4b\x80\xce\xb1\x4c\x2a\x47\x16\x50\x8b\x99\x61\x60\x52\xe8\x91\x2c\xb8\x1c\x1d\xa4\xa8\x21\x21\x60\xaa\x2c\x09\x90\xda\x19\x30\x2e\x27\x06\x5b\x52\x2a\x33\x99\x46\xe1\x56\xc7\x55\x19\x54\x34\xc9\x77\x5a\x2b\x88\x42\x48\x2f\x86\xea\x0b\x9b\x92\xd8\x49\xb2\x93\x6b\xc8\x50\x59\x7a\x7b\x15\x65\xca\xee\x4c\xb4\x55\x5c\xdb\x6a\xb8\x31\x1e\x6d\xc9\x94\xf9\xdd\x5e\xcd\x04\x65\x52\x87\x3d\xec\xac\xb3\xa2\x15\x7f\x0c\xbf\x1e\x5b\x83\xb5\xd2\xfd\x37\xdb\x4a\xad\x9d\x7d\xe2\x83\x3f\xa4\x0d\x46\x5e\x10\x5b\x69\x34\x89\x8e\xa5\xdf\x3d\x5d\xbb\xdf\x3c\x1b\xa2\x25\x3a\x47\xac\xbf\x6c\x37\x56\x2b\xf6\x6d\x71\xfb\x7e\xfa\xf1\xee\xff\xaf\xb7\x4e\x6f\xd0\x41\x66\x5c\x3d\xd1\xe8\x10\xab\x4e\xd4\x48\x0d\x08\x25\xb2\x93\x69\xa5\x90\x1b\xc2\xef\xe0\xb3\x83\x07\xa9\x14\x54\x3e\xca\x94\xea\xae\xc9\xd8\x14\xe0\x72\x82\x92\x69\x21\x4d\x65\x9b\x55\x21\x08\x01\x85\x00\xc3\x60\x16\xc4\x0f\x2c\x1d\x05\x51\xa3\xc9\x36\x31\x2c\x9e\x5a\xb0\xd5\x57\x3a\x2a\xb6\x9b\xa4\x9f\x1b\xf7\x91\xff\x93\x4a\x26\x4b\xda\x79\x2d\xd7\x74\xbc\x3d\xbc\x82\xdc\x4c\x93\x80\x88\xbd\x4b\xcb\x00\xce\xf4\xb3\x92\x4c\x62\x72\x0d\xb7\x3b\xa5\x82\xa4\xc6\x82\xf6\x20\xad\x49\x1d\x90\xe9\xd2\xd9\x29\x79\xb7\x47\xe5\xfe\x61\xba\xb1\x6c\x47\x62\x6f\xdf\xc2\x97\x2b\x12\xf7\x69\x8e\x7c\x58\xfc\x19\xad\xe0\xa5\x39\x2d\xcb\x58\xbb\x10\xac\x63\xa9\xe7\x1d\x67\x15\x95\x75\x90\x1b\x6d\xd8\xd7\xc4\x84\x60\x81\x4a\x8a\x03\xa6\x83\x6e\xdc\x44\xc8\xdd\x16\x84\x75\x9d\xe9\x41\xd3\xc8\xb4\x87\x5d\xb6\x11\xbd\x89\x18\x9e\x40\xe5\x73\xca\xf0\xf3\xb0\x5c\x95\x04\x9f\x74\x55\x1c\xc3\x6f\x5f\x3d\xd8\x58\x70\x38\xd7\xb6\xe2\xd7\xf6\x3b\xbc\x01\x74\xea\x9d\x5f\xf8\xed\xf6\x66\xfa\xf5\xb7\xe9\xc7\xbb\x5b\x9c\xfe\x7d\x33\xfd\xea\x0b\xdc\xff\x5e\xef\x77\x04\xb4\xa5\x7f\xe7\xec\x21\x57\x55\xce\xcc\x49\x13\xa3\x0b\xc9\x7a\xbc\xa3\xfe\xca\x3b\xc7\x2c\x48\x0b\x1b\x90\x90\xac\x42\xf9\x48\x30\xfd\x41\xfa\xa8\x48\x4c\x8c\x51\x84\x7b\x92\xb9\x0f\xbf\x94\x29\x9c\xe0\xf7\x46\xab\xd5\x08\xfc\x7c\xbf\xe0\xa1\x7c\x6e\xf9\x46\x43\x54\x21\x03\xc3\x36\x04\xbe\x12\x84\xed\xce\xce\x53\x50\x86\x95\x72\xf7\x86\x05\xf1\xc9\x3c\x43\xb2\x35\x15\x7f\xc3\xa9\xa1\x3d\x42\x7f\x44\x85\x7d\x20\xec\xe3\x99\xff\xa0\x5e\xd9\xf4\x22\x64\x83\xba\x27\x93\x6d\xa8\x6c\x27\x7d\x04\xa9\xdb\xa6\xa8\xac\xe9\xbd\x85\x89\xd4\x8e\xe6\xc4\xfe\xa7\xae\x94\x8a\x7f\x17\x49\x1c\x69\x4e\xe7\xb6\x5e\xdc\x0d\x34\x48\xc9\x94\x8e\x97\xc2\x1d\xbc\xf3\xfb\xb6\xab\xda\xa9\x9e\x6d\x87\x86\xf8\x75\x94\xb3\x90\x96\x58\x94\x8a\x06\xc4\xea\xa7\x88\x50\xc7\xa9\x3f\x08\xff\xbb\x81\x4a\xcb\xd2\xd8\xb1\xa2\xb4\xac\x12\x25\x53\xb5\x82\x06\xf5\x84\x58\xed\xb1\xa4\xae\x13\x7e\x95\xe3\x8a\x86\x85\x43\x26\x95\x23\xc6\x44\x9d\x5e\xb7\x36\x0f\xa3\x84\x62\x6d\xf6\x71\x11\xc1\x7d\x55\x2e\x2b\xf6\x26\xb1\x67\x4f\xdf\xcc\x30\xc9\xb9\xbe\xf7\xc7\xc2\x38\xdd\x04\xd4\x90\x17\x39\x69\x32\xc3\x05\xba\x93\x99\xc4\xe5\x27\xd5\x20\xf2\x7d\xed\xa1\xcb\xd3\x5a\x4f\xa6\x3e\xa8\x11\xb9\x40\xa9\x7a\x37\xa8\xb9\xd1\xbd\x91\x3f\x7f\x59\xfc\x7e\x84\xec\x87\xbe\xb2\xa9\x14\xdc\x57\x56\xa0\x3b\x90\xa2\xfe\xcf\xb0\x2a\x36\x27\xdf\xb5\x8f\x10\xdf\xe1\x05\x21\xc7\x05\x01\xc2\xba\x65\xae\xe1\xcf\x1d\xeb\x52\x90\x76\x32\x93\xa3\x30\xeb\xb4\x8c\x08\x05\x4a\x0d\x1d\xf8\x97\xa9\xd3\x95\x96\x3f\x2b\xfa\x5c\x5f\xdb\x7c\xa9\x1e\x66\x0e\x2d\x68\x39\xb2\x25\x34\x44\xd4\x73\xbb\xb6\xc0\xe5\xbd\x22\x3d\x77\xf9\xc9\x84\x0a\x5c\xca\xa2\x2a\x20\xc2\x3c\xbf\x23\x34\xf7\xf1\xde\xf7\xdf\x96\x61\xd3\x7f\x0c\x66\x38\xec\x3a\xd0\x10\xdc\x71\x07\x6a\xf8\xb5\xdd\xd2\xb9\x09\x4a\x3d\xd8\x85\x52\xff\xd2\x2e\x94\x7a\xa8\x0b\x6b\x82\xbf\xa8\x0b\xc3\x4b\xe7\xa9\xdc\xfc\xe2\x8b\x5f\x66\xc2\xad\xff\x85\x9a\xd7\xf6\x45\xe1\x62\xbd\x6b\xc9\xb2\x40\x5e\x8d\xd7\xbb\xc6\xba\x0f\x35\xee\x45\x1a\x58\x26\x14\x63\xbd\x7d\x49\x0b\x1e\x2e\xbc\x7f\x5d\x80\x48\xfb\x41\x61\x14\x1e\x35\xda\xb9\x69\x58\x4a\x99\x4e\xbf\x50\x6c\x90\x88\x58\xe1\x9b\x92\xcd\x4d\xa5\x84\xcf\x27\x41\xa8\x1c\x3c\x48\x97\x43\x8a\x7c\x54\x85\x18\x89\xe0\x8b\xf6\xc6\xf6\x32\xbd\xb1\x75\x66\xb4\xd8\xab\xb1\x5e\xfa\xb1\x62\x70\x03\x6c\xab\xa4\x56\x60\xd8\x89\x55\x60\x59\xfa\xba\xee\x0c\x94\x32\xfd\x11\x8a\xfd\x43\x4e\x3a\x1c\x65\xf1\x40\xb6\xe1\x2d\xdd\x19\x78\x43\x4b\x47\xac\x51\xbd\x39\xf7\xd1\xe6\x18\xb5\x95\xa4\x47\xca\xcd\x35\xdc\xb9\x43\x75\x90\xcf\x82\x3b\xba\x5f\xc0\x4f\x70\xc3\xe8\x8f\x1b\xc7\x7d\x52\xeb\xdf\xdd\x05\xf1\x4c\x19\xec\xe3\xa4\x20\xdc\x3f\x11\xa1\xa5\xd9\x53\x56\x49\xdb\x5b\x8d\x83\xdf\xf6\x37\xa4\x9d\x3c\xf8\x81\xbd\x95\x6d\xd2\x6f\xf0\x0b\xca\xce\xd9\xed\x33\xcf\x47\x37\x47\xd6\xff\xaa\xff\xa7\xcb\xd5\xe3\xd5\xbf\x01\x00\x00\xff\xff\x46\xe7\x7c\xa6\x71\x24\x00\x00")

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

	info := bindataFileInfo{name: "rego-abstract.json", size: 9329, mode: os.FileMode(420), modTime: time.Unix(1522434000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoAttributeIn = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\x4f\x6f\xdb\x36\x14\xbf\xf7\x53\x3c\x18\x05\x0a\x6c\x4d\xb1\x01\xc3\x80\xe6\x96\x43\x0f\x01\x76\x28\x86\x9d\x9a\xa6\xc1\x93\xf8\x6c\x71\xa1\x48\xf5\x91\x74\xec\x0d\xf9\xee\x03\x29\xc9\x96\x12\xdb\x92\x25\xd9\xee\x74\xb1\x21\x3d\xfe\xf8\x7e\x7c\x7f\x49\x42\xe3\x99\xb9\x75\x41\xb3\x6b\x98\x99\xe4\x6f\x4a\xdd\xec\xfd\x9b\xd6\x57\x41\x36\x65\x59\x38\x69\x74\x10\xfa\x43\x5a\x07\x66\x0e\x4b\x62\x2b\x8d\x26\x01\xe8\x1c\xcb\xc4\x3b\xb2\x1f\x5e\x8e\x45\x21\x64\x18\x88\xea\x33\x9b\x82\xd8\x49\xb2\xb3\x6b\x98\xa3\xb2\xf4\x42\xb4\x40\xe7\x88\x75\x4b\xee\xdf\x96\x48\x14\xfb\xb6\xbc\xfb\xe5\xea\xe3\xfd\xcf\x5f\xdf\xee\xfc\xde\xe2\x83\xcc\xb8\x7e\xa1\x52\x17\xad\x2d\x19\x90\x1a\x10\x0a\x64\x27\x53\xaf\x90\x6b\xc6\x1f\xe0\xd6\xc1\x93\x54\x0a\xbc\x25\x40\xa5\x9a\x63\xe6\x6c\x72\x70\x19\x41\xc1\xb4\x94\xc6\xdb\x7a\x94\x05\xd4\x02\x50\x08\x30\x0c\x66\x49\xfc\xc4\xd2\x51\x14\x35\x9a\x2c\x08\x4a\x15\x32\x89\x97\x4b\xb8\xd1\x57\x3a\xca\x77\xaf\x49\x3f\x3b\x1e\x22\xff\x27\x15\x4c\x96\xb4\x0b\x5a\x6e\xe9\x84\xf5\x08\x0a\x72\xfd\x99\x04\x94\xd8\xfb\xb4\x8c\xe0\x4c\xdf\xbd\x64\x12\xb3\x6b\xb8\xdb\x2b\x15\x25\x35\xe6\x74\x00\x69\x4b\xaa\x43\xa6\x49\x67\xaf\xe4\xfd\x01\x95\xfb\xfb\x69\x6b\x58\x71\xd8\x59\xdb\x53\x28\x65\x9e\x48\x3c\xa4\x19\x72\xb7\xf8\x2b\x5a\xd1\x4a\x0b\x5a\x15\xe0\x32\x74\x80\x60\x1d\x4b\xbd\x68\x18\x2b\xf7\xd6\x41\x66\xb4\x61\x70\x06\x12\x82\x25\x2a\x29\x3a\x96\x0e\x9a\x7e\x53\x42\xee\x5f\xc1\xf0\x3c\x77\x98\x62\x4b\xd3\xc8\xb4\xc7\xba\xec\x22\x7a\x53\x62\x04\x02\x3e\xc4\x94\xe1\xd7\x6e\xb9\x2e\x08\x3e\x69\x9f\x1f\xc3\xef\x50\x3e\x68\x0d\xe8\x8e\xb5\x9d\xf8\xd5\xfa\x75\x4f\x00\x8d\x84\x17\x06\x7e\xbb\xbb\xb9\xfa\xf2\xeb\xd5\xc7\xfb\x3b\xbc\xfa\xe7\xe6\xea\x4b\xc8\x70\x3f\x7d\x7d\x7b\xd8\x12\xe1\x79\x1e\x67\x2b\xef\xcc\x82\x34\x31\xba\x18\xad\xc7\x5b\xea\xaf\x8c\x1a\x76\x91\x16\x5a\x90\x90\xac\x63\xfe\x48\x30\x7d\x24\x7d\x94\x2b\x26\xc6\x28\xc2\x03\xd1\xdc\x87\x5f\xca\x84\x41\xd3\x07\xa3\xd5\x7a\x02\x7e\x29\x6a\x08\x50\x21\xb8\x2c\x39\x10\x3e\x86\x60\x9c\x86\x20\xa4\x82\x38\xdd\xd9\x79\x0a\x9a\xa3\x57\xee\xc1\xb0\x20\x1e\xcc\x33\x46\x5b\x9d\xf2\x5b\x46\xf5\x36\x54\xf9\x50\xa3\xe2\x3c\x10\xe7\x09\xcc\x1f\xa9\x57\x38\x9d\x84\x6c\x54\x77\x30\xd9\x9a\xca\x6e\xd2\x47\x90\xba\xab\xb3\xca\x96\xde\x7b\x98\x49\xed\x68\x41\x1c\xfe\x6a\xaf\x54\xf9\x9b\x27\xe5\x9b\xba\x3c\x6f\x12\xc6\xfd\xc8\x05\x29\x98\xd2\xe9\x42\xb8\x81\x77\x7e\xdb\x36\x55\x1b\x6a\xd9\xcd\xab\x31\x76\x9d\xa4\x18\xd2\x0a\xf3\x42\xd1\x08\x5f\xfd\x54\x22\x54\x7e\x1a\x2a\xe1\xff\xd7\x51\x69\x55\x18\x3b\x95\x97\x16\x3e\x51\x32\x55\x6b\xa8\x51\x07\xf8\x6a\x8f\x21\x55\x9e\x08\xa3\x1c\x7b\x1a\xe7\x0e\x73\xa9\x1c\x31\x26\x6a\x78\xde\x6a\x17\xa3\x84\xca\xdc\x1c\xfc\xa2\x04\x0f\x59\xb9\xf0\x1c\x96\xc4\x9e\x3d\x7c\xe7\x86\x49\x2e\xf4\x43\x28\x0b\xd3\x74\x13\x50\x41\x5e\xa4\xd2\xcc\x0d\xe7\xe8\x06\x33\x29\x87\x0f\xca\x41\x14\x1a\xdb\xae\xdd\xd3\x56\x4f\xa6\x3e\xa8\x25\x72\x8e\x52\xf5\xee\x50\x33\xa3\x7b\x23\xdf\x7e\x5e\xfe\x76\x84\xec\xef\x7d\x65\x53\x29\xb8\xaf\xac\x40\xd7\x11\xa2\xe1\x19\x97\xc5\x16\x14\xda\xf6\x09\xfc\x3b\x1e\x21\x64\xb8\x24\x40\xd8\xb6\xcc\x15\xfc\xb9\x7d\x5d\x0a\xd2\x4e\xce\xe5\x24\xcc\x1a\x2d\x23\x42\x8e\x52\x43\x03\xfe\x34\x79\xda\x6b\xf9\xdd\xd3\x6d\xb5\x6f\x0b\xa9\x7a\xdc\x72\x68\x41\xab\x89\x57\x42\x43\x89\x7a\x6e\xd3\xe6\xb8\x7a\x50\xa4\x17\x2e\x1b\x4c\x28\xc7\x95\xcc\x7d\x0e\x25\xcc\xeb\x3d\x42\xbd\x21\xef\xbd\x01\xde\x30\xac\xfb\x8f\xd1\x0c\xc7\x6d\x07\x6a\x82\x7b\xf6\x40\x35\xbf\x4d\xb7\x74\x6e\x82\x52\x8f\x36\xa1\xd4\x3f\xb4\x09\xa5\x1e\x6b\xc2\x8a\xe0\x0f\x6a\xc2\x78\xd4\x39\x94\x5b\x18\x7c\xf1\xcd\x4c\xdc\xf5\x9f\xa8\x79\xdd\x9c\x28\x5c\xac\x77\x2d\x58\xe6\xc8\xeb\xe9\x7a\xd7\x32\xef\x43\x85\x7b\x91\x06\x96\x09\xc5\x54\x67\x5f\xd2\x42\x80\x8b\xe7\x5f\x17\x20\xb2\xb9\x51\x98\x84\x47\x85\x76\x6e\x1a\x96\x52\xa6\xe1\x1b\x8a\x16\x89\x12\x2b\x5e\x2a\xd9\xcc\x78\x25\x42\x3c\x09\x42\xe5\xe0\x49\xba\x0c\x52\xe4\xa3\x32\xc4\x44\x04\x4f\xda\x1b\xdb\xcb\xf4\xc6\xd6\x99\xc9\x7c\xaf\xc2\x3a\xf5\x61\xc5\xe8\x06\xd8\xfa\xa4\x52\x60\x5c\xc5\xca\xb1\x28\x42\x5e\x77\x06\x0a\x99\x3e\xc6\x64\xff\x94\x91\x8e\xa5\xac\x2c\xc8\x36\x9e\xa5\x3b\x03\xef\x68\xe5\x88\x35\xaa\x77\xe7\x2e\x6d\x8e\x51\x5b\x49\x7a\xa2\xd8\xdc\xc2\x9d\xdb\x55\x47\xd9\x2c\x9a\xa3\x79\x05\x3e\xc0\x0c\x93\x1f\x6e\x1c\x77\xa7\xd6\xbf\xbb\x8b\xe2\x73\x65\xb0\x8f\x91\xa2\x70\xff\x40\x84\x0d\xcd\x9e\xb2\x4a\xda\xde\x6a\x74\x5e\xee\xb7\xa4\x9d\xec\xbc\x61\xdf\xc8\xd6\xe1\x37\xfa\x04\x65\xef\xd7\xdd\x5f\x5e\xbf\x6d\xbf\x79\x7e\xf3\x5f\x00\x00\x00\xff\xff\xde\xce\x05\x55\xae\x22\x00\x00")

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

	info := bindataFileInfo{name: "rego-attribute.in", size: 8878, mode: os.FileMode(420), modTime: time.Unix(1522433994, 0)}
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

	info := bindataFileInfo{name: "rego-info.json", size: 691, mode: os.FileMode(420), modTime: time.Unix(1521233328, 0)}
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

	info := bindataFileInfo{name: "rego-spec.in", size: 6065, mode: os.FileMode(420), modTime: time.Unix(1522433994, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoSpecJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5b\xdd\x6f\xdb\x38\x12\x7f\xcf\x5f\x31\x70\x0b\x14\xb8\x8b\x9b\x1e\x70\x28\xd0\xbc\x15\x77\x05\xb6\xc0\x2e\x50\x74\xf7\xa9\x41\x1a\xd0\xe2\xd8\x62\x43\x91\x2a\x49\x39\xf6\x16\xf9\xdf\x17\x24\xf5\x69\x7d\x5b\xb2\xdd\xe6\x61\xb7\x91\x86\xa3\xf9\xcd\x27\x87\x9c\xfc\xb8\x02\x00\x58\xbc\xd4\x41\x88\x11\x59\xdc\xc2\x22\x34\x26\xbe\xbd\xb9\xf9\xa6\xa5\x58\xfa\xa7\xaf\xa5\xda\xdc\x50\x45\xd6\x66\xf9\xe6\xed\x8d\x7f\xf6\x62\x71\xed\x57\x1a\x66\x38\xda\x75\x7f\xc6\x18\xb0\x35\x0b\x88\x61\x52\x64\x6f\x29\xea\x40\xb1\xd8\x3d\xba\x85\xc5\x67\x8c\x15\x6a\x14\x46\x03\x11\x20\x57\xdf\x30\x30\xf0\xc4\x4c\x08\x84\x73\x60\xf6\xb1\x31\x8a\xad\x12\x83\x96\x82\x82\x42\xee\xf8\xe9\xfc\x73\xfb\xd8\x7d\xcd\xaf\xcd\x9e\x12\x4a\x99\x25\x23\xfc\x93\x92\x31\x2a\xc3\x50\x2f\x6e\x61\x4d\xb8\xc6\x94\x44\xe1\xf7\x84\x29\xa4\x8b\x5b\xb8\x73\x4f\xdc\xd3\x48\x52\xe4\x0b\xf7\xfb\x7d\x4a\x18\x97\x39\xfc\x28\x48\x0b\xc9\x2a\xcf\xbd\xfa\x14\xae\xad\x54\x2f\x6e\x28\xae\x99\x70\xb2\xe8\x9b\xd2\x8a\x9c\xfc\xf9\xba\xe0\x58\x80\xab\x31\xcc\x60\x12\xa5\xc8\x7e\x71\x5d\x7d\x79\xa0\xd4\xdf\x99\x36\x20\xd7\x85\xae\x60\x85\xe6\x09\x51\x80\x09\x11\x54\xa6\x72\xa4\x99\xc2\xad\x62\xa5\x09\x51\x81\x14\xa8\x5f\x1f\xb2\x67\x06\xa3\xba\x48\x1d\x38\xb3\x0f\x2f\x2a\x0b\x9e\x1b\x31\x7b\x85\x0f\x54\x60\xc9\x3a\x05\xc7\xe7\xdc\xb5\x72\xc2\xaa\xa1\x5a\x3e\x91\x3b\xea\x1f\xee\x7d\xb7\x4e\x7f\x93\x9c\x6a\x58\x11\xcd\x02\x60\x62\x2d\x55\xe4\x10\x5a\x35\x37\x2b\xb5\xa6\xc5\x46\x47\xcd\xdf\xf6\x3b\x6c\xc9\x4b\x1a\x1c\x37\x7f\x1b\x93\xe0\x91\x6c\xf0\x80\x7f\xba\x50\x9b\x07\x41\xa2\xb6\x97\x32\x51\x01\xb6\x12\xa0\x30\xcc\xec\x5b\x5f\x97\xf5\x55\x79\x7b\x7f\x7d\x55\x95\xbf\x25\x9e\x0a\x55\x70\x46\x74\xcb\xcb\xda\xa7\x4a\xee\x9e\xae\x03\x23\x0b\x73\x00\x01\x5d\x4e\x43\x87\x56\xc9\x79\x76\xc5\x57\x4e\xd4\x1e\x08\x35\x3e\xda\x28\x26\x36\x8b\x46\xc2\xe7\xda\xd3\xe7\x06\x85\x06\x0a\x89\xc1\xc1\x5a\xf8\xab\x39\xb4\x03\x22\x60\x85\xe0\x99\xd1\x5e\xf8\x2b\x29\x39\x12\x51\x97\xbb\x49\x42\x8a\x1c\xc7\x49\xd8\x2e\xa0\xe7\x35\xbb\x80\x65\x01\x86\xeb\xb1\xf4\x68\x78\x80\xd7\x44\x6d\x73\x81\x26\x49\xcb\xc1\x35\x54\xd2\xff\x91\x08\xf9\x32\x20\x1a\x29\x6c\x51\xe9\x8a\xb4\x69\xa4\xcf\x2b\xe4\xce\xa0\xa0\xe3\x03\x73\x45\x34\x56\xc3\x10\x4c\xc8\xf4\xc1\xa3\x94\x3b\xfc\x52\x41\xba\x41\x33\x53\x84\x2a\x34\x8a\xe1\x76\xee\x10\xc8\xca\xc1\x18\x21\xad\xe3\x64\x9e\x94\xae\x3f\x53\x18\xc4\x8a\x6d\x67\x48\x7a\x4c\x43\xca\x69\x5e\x65\x56\x6b\xe4\x18\x21\x63\x9e\x28\xc2\xcf\x13\xa6\x45\x99\x1f\x23\x61\x22\xd8\xf7\x24\xb5\x7d\xae\x52\x26\x36\x4e\xd0\x13\x98\x5a\x49\x39\x39\x74\x48\x60\x1b\x04\xed\x75\x29\xa5\x19\x59\xeb\x47\xd9\x3e\x89\xe9\x7c\xe5\xd8\x33\x9b\x12\xea\x03\x76\xd6\x5d\xfd\x49\xf7\x66\xb4\x25\x8f\xa7\xde\x8b\xb4\xd4\x94\xd5\xb6\xb9\x23\x36\xb2\x31\x31\x06\x95\xf8\xd4\xb3\x1f\xfc\xba\xbd\x7b\xb3\x7c\x77\xff\xef\x97\xed\xca\x1f\x54\x18\xda\xb6\x8d\x45\x83\xc9\x04\x10\x88\x89\x32\x2c\x48\x38\x51\x19\xe0\xd7\xf0\xd1\x36\xa5\x9c\x43\xa2\xd1\x35\xa6\xa5\x35\x6b\x25\x23\x9f\x29\x15\x6e\x99\x4c\x74\xb6\xca\xf7\xab\x84\x52\x90\x0a\xe4\x16\xd5\x93\x62\x06\x7d\x38\x09\xd4\x40\x31\xe0\x44\x75\x38\xc1\x88\x42\xd6\x68\xc6\x2e\xf0\xd5\xbe\x3b\x87\x33\x3e\xcd\x43\x6f\x47\x52\xa1\x6c\xe9\x1e\xea\xa0\x7a\x68\x5a\xbb\x8d\xf2\xcf\x7d\x87\xc8\xc3\xdd\xb4\xb2\xac\xa7\x77\xa9\x7e\x82\x73\xf9\x84\xf4\x21\x08\x89\xea\x27\xaf\xc1\x72\x56\xda\xe0\x2e\x06\x13\x12\xd7\xc8\xb8\xec\x5a\x32\x56\x94\x68\x03\xa1\x14\x52\xd9\x9e\x67\x85\xb0\x25\x9c\xd1\x1e\xd5\xc1\xd0\x0d\x50\xf6\xd3\x90\x0c\x5b\x60\x4a\x16\x0c\xd0\x4b\x13\xd0\xf7\x9e\x87\x05\x90\xd8\x98\x92\xaa\xee\x96\xfb\x18\xe1\x83\x48\xa2\x31\xf8\xba\xf2\x41\x65\x41\x7f\xac\x35\xf2\x4f\xf5\xd7\xff\x01\x28\xe5\x3b\xbb\xf0\xeb\xdd\xfb\xe5\x97\xff\x2c\xdf\xdd\xdf\x91\xe5\xdf\xef\x97\x5f\x6c\x82\xfb\xd7\xcb\x6e\x43\x40\xe3\xae\xb4\xf2\xb6\xcf\x54\x89\x91\x1b\x14\xa8\x6c\xf1\x39\xca\x50\xb6\xac\x15\x66\x61\x1a\x2a\x2c\x61\xb5\x77\xe9\x63\x45\x82\x47\x14\xa3\x3c\xb1\xb5\xd2\x8d\xc1\xe7\xba\x5c\x26\xc5\x83\x14\x7c\x3f\x03\x3e\x5b\xad\x2d\x2b\x1b\x5b\x1a\x0d\xd0\xc4\x45\xa0\x6f\xa6\xc1\x66\x82\xf2\xd9\xe6\xf9\x70\x52\x5c\x93\x84\x9b\x07\xa9\x28\xaa\xa3\x71\xba\x60\xcb\x32\x7e\xc5\xa8\x89\x6d\x29\x89\x2d\x51\xee\x3b\xe0\xbe\x63\x91\x3f\xe2\xa0\x68\x3a\x09\x58\x27\xee\xd1\x60\x33\x28\xcd\xa0\x47\x80\xba\xcb\x92\x4a\x01\xef\x1a\x16\x4c\x18\xdc\xa0\xb2\xff\x14\x09\xe7\xfe\xff\xd1\xca\x3f\xc9\xaa\x73\x9e\x2f\xee\x27\x2a\x24\x56\x18\xcc\x17\xc2\x25\x7e\xe7\xb7\x6d\xff\x01\x4d\x1f\x96\x86\xc3\x9a\x63\xec\x3a\x4b\x2d\xc4\x1d\x89\x62\x8e\x13\x7c\xf5\x83\xe7\x90\xfa\xa9\x2d\x84\xbf\xae\xa3\xe2\x2e\x96\x7a\x2e\x2f\x8d\x93\x15\x67\x01\xdf\x43\xc6\xf5\x08\x5f\x1d\xb0\x24\xcd\x13\x76\x95\x51\x09\x4e\x73\x87\x35\xe3\x06\x15\x59\xf1\xe3\xf3\x56\xb5\x18\xd9\xd6\xd1\xe6\x66\xeb\x17\x9e\xb9\xcd\xca\x71\xa2\xac\x4a\xf4\xd9\xc3\x77\x2d\x15\xb2\x8d\x78\xb0\x65\x61\x9e\xdd\x04\xa4\x2c\x2f\x52\x69\xfc\xad\xce\xd1\x48\xfc\xf2\xa3\x72\x10\xda\x7d\x6d\x5f\xf3\x54\xc8\xa9\x70\x08\x57\xcf\x39\x22\xec\xf0\x6e\xab\x95\x38\x0e\xa5\x18\xcc\xf9\xe3\xa7\xed\x7f\x47\xd0\xbe\x1d\x4a\x1b\x30\xaa\x86\xd2\xba\x23\x99\x5e\xd2\x69\x59\x6c\x83\x76\xd7\x3e\x83\x7f\xbb\x13\x84\x90\x6c\x11\x08\x14\x5b\xe6\x94\xfd\xb9\x7d\x9d\x51\x14\x86\xad\xd9\x2c\xc8\x4a\x5b\x46\x02\x11\x61\x02\x4a\xec\x4f\x93\xa7\xfd\x71\xe5\xc7\xb4\x6d\xb3\xa9\x7a\x9a\x3a\x04\xc5\xdd\xcc\x9a\x10\xe0\xb9\x9e\xdb\xb4\x11\xd9\x3d\x70\x14\x1b\x13\x1e\x0d\x28\x22\x3b\x16\x25\x11\x78\x36\xf5\x1e\x21\xeb\xc7\x07\xf7\xbf\x39\xc2\x6c\xff\x31\x19\xe1\xb4\x76\x20\x03\xd8\xd2\x03\x65\xf8\xf2\xdd\xd2\xb9\x01\x32\x31\xd9\x84\x4c\xfc\xd4\x26\x64\x62\xaa\x09\x53\x80\x3f\xa9\x09\x3b\x2f\x60\xfa\xb0\x95\x6f\xe0\x2e\xd6\xcc\xb8\xae\xff\x44\x9b\xd7\xfc\x44\xe1\x62\x7b\xd7\x58\xb1\x88\xa8\xfd\x7c\x7b\x57\x9f\xf7\x21\xe5\x7b\x91\x0d\xac\x42\x42\xe7\x3a\xfb\x62\x1a\x2c\x3b\x77\xfe\x75\x01\x20\xf9\x85\xc2\x2c\x38\x52\x6e\xe7\x86\xa1\x31\x50\x1d\x03\x02\xa3\x40\x78\x5e\xee\x4e\x49\x87\x32\xe1\xd4\x4f\xcd\x10\x9e\xce\x4b\x06\x44\x8d\xca\x10\x33\x01\x3c\xe9\xde\x58\x5f\x66\x6f\xac\x8d\x9c\xcd\xf7\x52\x5e\xa7\x3e\xac\x98\xbc\x01\xd6\xc9\x2a\x15\x60\x5a\xc5\x8a\x48\x1c\xbb\xa9\x01\x09\x31\x0b\x1e\x5d\xb2\x7f\x0a\xd3\xd1\x53\x5f\x90\xb5\x3b\x4b\x37\x12\x5e\xe1\xce\xa0\x12\x84\xbf\x3a\x77\x69\x33\x8a\x08\xcd\x50\xcc\x14\x9b\x05\xbb\x73\xbb\xea\x24\x9b\x39\x73\x94\x6f\xc0\x8f\x30\xc3\xec\x87\x1b\xe3\xae\xd4\x86\xef\xee\x1c\xf9\x9a\x4b\x32\xc4\x48\x8e\x78\x78\x20\x42\x0e\x73\x20\x2d\x67\x7a\xb0\x18\xbd\x77\xfb\x15\x6a\xc3\x7a\x2f\xd8\x73\xda\x2c\xfc\x26\x9f\xa0\xb4\xbe\x6d\x7e\xd3\x30\x5a\x77\xd5\xfc\x5b\xd3\x70\x7c\xc7\x20\xf7\xe7\x8c\xa4\x7b\x9e\xc5\xdd\x36\xbb\x39\xe1\x74\x06\x03\x48\x3e\x2b\x5f\x8c\xca\x3f\xc9\xea\x38\x51\x7d\xd4\xe5\x4c\x13\xdd\x9d\x63\xdb\x25\x68\x7a\xe2\xdc\xf5\x51\x03\xc7\x9c\x14\x33\x22\xfa\x60\xda\xb8\x98\x8d\x09\x12\xa5\x50\x98\x86\xb9\x92\x0b\x0f\x24\xb7\x21\x48\xc7\x91\x2f\x80\xa0\xf7\x6e\xaf\x15\x85\xf5\xde\x21\xd7\x79\x93\x67\xa9\x3b\x26\x80\x9b\xff\x6c\xa3\xdd\x49\xdb\xbe\x36\x7e\xae\xb6\xd1\x8c\xf9\x50\xed\xd9\x0d\x39\x7e\x08\x33\x5f\x51\x3a\x4b\xe5\x4c\x3c\xda\x3c\x65\x6a\x10\xe7\x1c\xc5\x3c\x6a\xb4\xb1\x51\xdf\xe9\x5c\xe3\x69\xb5\x3d\xa0\x56\x74\x3a\x6b\xbd\x5e\x94\xef\x90\x35\xc4\xa8\x80\x04\x03\xea\xc8\xff\x65\x90\x44\x28\xcc\xc1\x5f\x02\xf1\xc6\xe9\xd3\xf9\xaa\xc5\xcc\x79\xbc\x8a\x22\xbb\x71\x3e\x9c\x71\x99\x75\xf6\x77\x64\x9e\x6e\x96\xd0\x33\xe9\x9d\xc2\x39\x4a\xc0\x31\xe9\xa7\x59\xba\x0d\x9a\xd3\x88\x36\x32\x58\x9b\xa5\xf3\x4c\x66\x10\xb0\x2d\x16\xaf\xfc\x7f\x9f\xaf\xfe\x09\x00\x00\xff\xff\xd8\x1e\x9d\x2a\x4f\x3a\x00\x00")

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

	info := bindataFileInfo{name: "rego-spec.json", size: 14927, mode: os.FileMode(420), modTime: time.Unix(1522434000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoTypeMappingJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\x3f\x6f\xdb\x3c\x10\xc6\x77\x7f\x8a\x03\xf3\x02\x2f\x50\x24\x56\xa7\x0e\x5e\x8b\x02\x5d\x0a\x74\xe8\x56\x74\x38\x8b\x27\xf9\x12\x8a\x64\x8f\xe7\xa4\x6a\xe1\xef\x5e\x50\x8a\x98\xc8\xb2\xe3\x00\x45\x36\xe9\xfe\xe8\x39\x3d\xbf\x23\xff\xac\x00\x00\xcc\x7f\xa9\xde\x51\x87\x66\x03\x66\xa7\x1a\x37\x55\x75\x9b\x82\xbf\x19\xa3\xeb\x20\x6d\x65\x05\x1b\xbd\x79\xff\xa1\x1a\x63\x57\xe6\x7a\xec\x54\x56\x47\xb9\xef\x5b\x1f\x09\xbe\x60\x8c\xec\xdb\x34\x65\x2d\xa5\x5a\x38\x2a\x07\x3f\xd4\xec\x38\x41\xc3\x8e\x80\x13\xec\x13\x59\xd0\x00\x1d\x46\xa0\x5f\x4a\xe2\xd1\x81\xf6\x91\x52\x8e\x3a\xf4\xed\x1e\x5b\x82\x14\xa9\xe6\x86\xeb\xb4\x2e\x92\x7d\x1c\x14\xc3\xf6\x96\x6a\x9d\xa2\x68\x2d\x67\x1d\x74\x5f\x25\x44\x12\x65\x4a\x66\x03\x0d\xba\x44\x8f\x25\x11\x35\xcb\xcc\xf2\xa3\x01\x43\x7a\xfd\x6e\xf6\x3e\x1a\x23\xd4\x64\xad\xab\xca\x52\xc3\x7e\x50\x48\x55\x8d\x4a\x6d\x90\xde\x94\xe2\xc3\xf0\x74\x28\xbf\x5d\x6a\xe7\x12\xa5\x71\x21\x54\x7c\xfc\x38\x95\x5c\xcf\x0b\x16\x56\x12\x4c\x5f\x2b\x6e\x6e\x7b\x68\xc9\x93\xa0\x06\x49\xc0\x1e\x82\x58\x92\x6c\xa7\x0a\xfa\xe4\x50\xe9\xd8\x6a\xf6\x39\xbb\x23\x16\x08\x0f\x1e\xb8\x8b\x8e\x3a\xf2\x8a\x59\x68\x7d\x3c\xc4\x49\xeb\x4b\xf6\x32\x82\x52\xfa\x32\x8a\x97\x90\x5c\x40\xd3\x8d\x1b\x68\x16\x3d\x87\xd5\xe9\xb7\xc3\xd3\x64\x66\x6a\x3e\x4f\xe7\xf9\x96\x5f\x24\xc4\x09\xee\xa8\x7f\xbe\xeb\x13\x06\x0e\x1e\xd0\xcf\x59\xfc\x9f\x20\xed\xb7\xf9\x09\x1a\x09\x1d\x60\x59\xfd\x37\x66\x21\xf4\x73\xcf\x42\xd6\x6c\xe0\xfb\x12\xc1\x20\x33\x0b\xff\x38\x66\x79\x01\xe2\xe3\xa0\x67\x30\x9e\xd8\xeb\xb2\xac\x76\x30\xe6\xf8\xcf\x17\x0e\x24\x95\xd3\xc8\x97\x7d\x26\x2f\xca\xab\x87\xf9\x1c\x1e\x32\xb5\x61\xb9\xd0\xf1\x6f\xca\x27\xe5\x0d\x66\xea\x62\x90\xd7\x4f\xf5\xe9\x9e\xbc\xee\xd1\xc1\x13\x63\x18\x3f\x01\x41\x20\x62\x7d\x97\xef\x4d\xdd\xa1\x42\x94\x70\xcf\x36\xdf\xa8\xff\x3c\xf7\xb9\xe3\x33\x5e\x7c\xab\xc3\xea\x6f\x00\x00\x00\xff\xff\x72\x98\x35\xd6\x4b\x06\x00\x00")

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

	info := bindataFileInfo{name: "rego-type-mapping.json", size: 1611, mode: os.FileMode(420), modTime: time.Unix(1521685528, 0)}
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

