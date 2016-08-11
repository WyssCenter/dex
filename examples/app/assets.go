// Code generated by go-bindata.
// sources:
// data/index.html
// DO NOT EDIT!

package main

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

var _dataIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x93\xcf\x8a\xe3\x30\x0c\xc6\xef\x7d\x0a\xe1\x7b\x37\xf7\xc5\x29\xec\x0e\xbd\x0d\x14\xe6\x05\x8a\x63\xab\x89\xa9\xff\x61\x2b\x43\x4b\xe9\xbb\x8f\x53\x37\x61\x52\xd2\xa1\x73\x93\xd1\x27\x7d\x3f\x49\x98\x77\x64\xcd\x66\x05\xc0\x1b\xaf\xce\x43\x90\xc3\x83\x8f\x16\x84\x24\xed\x5d\xcd\x2a\xe3\x5b\xed\x58\x49\x0d\xd9\x30\x85\x00\xff\x7a\xea\xd0\x91\x96\x82\x10\x72\xd9\x5f\xae\x5d\xe8\x09\xe8\x1c\xb0\x66\x84\x27\x62\xe0\x84\xcd\xb1\x8c\x3e\xa5\xbd\x34\x3a\xcb\x19\x04\x23\x24\x76\xde\x28\x8c\x39\xe5\xad\x15\xeb\x84\x41\xc4\xdc\x46\x81\xd1\x89\xc0\x1f\xa0\x88\xd7\x5a\xa5\x6f\xee\x55\x58\x26\xd9\x9e\x28\x0a\x48\xd2\x07\x4c\xcf\x29\x70\x50\xed\x8b\xea\x45\x8a\xbb\x78\x4e\x70\xb9\x80\x3e\xc0\x9f\xdd\xee\x3f\x5c\xaf\x13\xc4\xcc\x36\xf5\x8d\xd5\xd9\xf8\x53\x98\x3e\x3f\xdf\x6f\x5b\x1c\x76\x64\x49\xc4\x16\xa9\x66\xfb\xc6\x08\x77\x64\xb7\x6e\x68\x12\xfe\xb2\x55\xa9\x73\x6a\x2c\xe3\xd5\xd0\x7c\xb3\x5a\x80\x7b\xb8\xa8\x14\xc6\x34\x42\x1e\x19\x58\xa4\xce\xab\x9a\x65\x9e\xa1\x61\xf1\x7e\xf3\x0a\x57\x0b\x18\xb3\x73\x66\xcd\x04\x34\x2d\x87\x37\x71\xb3\x54\xf9\x30\xc0\xf6\x24\x3b\xe1\x5a\xbc\x39\x8d\xbe\x23\xfe\x7c\xa8\xfb\x30\xce\xd3\x4f\x03\x45\x6c\xf3\xb5\x30\xb2\x57\xdc\x3f\x46\x31\x40\xf5\xdc\x9a\x57\xe5\x43\xf0\xaa\xfc\x90\xaf\x00\x00\x00\xff\xff\x9c\x89\xe2\x28\x29\x03\x00\x00")

func dataIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataIndexHtml,
		"data/index.html",
	)
}

func dataIndexHtml() (*asset, error) {
	bytes, err := dataIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/index.html", size: 809, mode: os.FileMode(436), modTime: time.Unix(1468620773, 0)}
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
	"data/index.html": dataIndexHtml,
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
	"data": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{dataIndexHtml, map[string]*bintree{}},
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