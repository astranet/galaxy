// Code generated by go-bindata.
// sources:
// templates/client_rpc_go.tpl
// templates/handler_rpc_go.tpl
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

var _templatesClient_rpc_goTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\xdf\x8f\xe2\x36\x10\x7e\x8e\xff\x8a\xb9\x48\x7b\x8a\x57\x51\xd2\x67\x24\x2a\x5d\xd9\xad\x7a\xad\xca\x22\x6e\xfb\x54\x55\x2b\x93\x0c\xc4\xdd\xc4\x0e\xb6\x03\x5d\x45\xf9\xdf\x2b\xff\x80\x70\x57\xb8\x05\xf5\x09\x66\xec\xf9\xf1\xcd\xf7\x79\x92\xe7\x30\x93\x25\xc2\x06\x05\x2a\x66\xb0\x84\xd5\x1b\x34\xa8\xab\xe5\x62\x96\xc1\xc3\x13\xcc\x9f\x9e\xe1\xf1\xe1\xf3\x73\x46\xf2\x1c\x3e\xd5\x35\x14\x15\x13\x1b\xd4\xd0\x74\xda\xc0\x0a\xa1\x94\x02\x81\x0b\x28\x3a\x6d\x64\x03\x45\xcd\x51\x18\x30\x15\x33\xa0\x2b\xd9\xd5\x25\x20\x37\x15\x2a\xc0\x66\x85\x25\x48\x05\x7b\xc5\x5a\x30\x15\xd7\x19\x21\x2d\x2b\x5e\xd9\x06\xa1\xef\xb3\x85\xff\x3b\x67\x0d\x0e\x03\x21\xbc\x69\xa5\x32\x90\x90\x28\x5e\xbd\x19\xd4\x31\x89\x62\x14\x85\x2c\xb9\xd8\xe4\x7f\x6b\x29\xac\x83\xcb\x9c\xcb\xce\xf0\xda\x1a\x02\x4d\x5e\x19\xd3\xc6\x84\x44\xf1\x86\x9b\xaa\x5b\x65\x85\x6c\xf2\xf6\x75\x93\xa3\x52\x52\xe9\x98\x50\x42\xcc\x5b\xeb\xea\xfd\x8c\xcc\x74\x0a\x17\x0a\xd7\xfc\x9f\x61\xf8\x82\x6a\xc7\x0b\x9c\xf9\xfe\xb9\x30\xa8\xd6\xac\x40\xe8\x49\x74\xf1\xb2\x3b\xfa\x55\x4b\xe1\xa3\x3e\x1f\x82\x7e\x92\xe5\xdb\x30\x90\xe1\xba\x62\x4f\xad\xe1\x52\x68\xd0\x46\x75\x85\x81\xde\xc6\xad\x3b\x51\x40\x51\x61\xf1\x7a\x5d\x70\x22\x5b\x03\xf7\xd7\xdd\xa5\xd7\x5e\xb4\xd8\xf9\x1a\x6c\xea\xe9\x14\x04\xaf\xad\x23\x72\x26\x7c\xbc\x2e\x45\x3f\x90\x68\x20\x91\x42\xd3\x29\x61\x33\x7d\x67\x28\xbf\x3c\x3f\x2f\xce\x8d\xff\x41\x26\x0a\xb7\x70\x6f\xb9\xcd\x96\xb8\xed\x50\x1b\x0a\xc9\xc1\xd6\xad\x14\x1a\x53\x70\x14\xd3\xe3\xf0\xe6\xb8\x7f\xa7\xc5\x84\x44\x36\x45\x28\xf9\xbd\x7e\x52\x12\xdd\x30\xdf\x94\xd0\x77\xf5\xd5\x1f\x67\x62\x07\xb9\x5c\xcc\xbc\x7f\xa1\xf8\x8e\x99\xf0\x06\xc2\xb0\x27\xb7\xea\x80\xa6\x24\x3a\x41\x36\x81\xf1\x7f\x6a\xd9\x38\x61\xe0\x7c\xe1\x51\x89\x0e\xf6\xb5\xb8\xaf\x1e\xa7\xed\xe0\xeb\x87\xd3\xb4\x35\x36\x28\x0c\xb3\x89\xc2\xeb\xf1\x34\x26\x2f\x61\xa3\xdc\x5f\x6c\x97\x42\x79\x5e\x21\x7f\xfe\x65\x37\xc7\x41\x19\x7e\xe6\xba\x75\x36\x4c\xa6\x10\x32\x67\x63\xd7\x99\x97\x1a\x75\xb2\xb7\xb7\x3e\x8c\xb2\xb7\xe6\xd4\xa7\xd2\xd9\xa3\xfd\x59\x27\xf1\xc5\x9e\x26\x70\xb7\x8b\x5d\x25\x4a\xa2\x03\xd5\x82\xd7\xce\x15\x9e\x84\x6e\x2d\xd2\x14\x5e\x6c\x33\x7e\x8d\x65\x4b\x64\xe5\xa7\xba\x4e\xec\x69\x66\x8f\x29\x89\x5e\x60\x0a\x47\x3b\x9b\xd5\x52\x63\xe2\x5b\x74\xde\x2f\x86\x99\x4e\xbb\x2d\xfe\x61\xea\xb8\x0e\xae\xa7\xdf\x5c\xe3\x7c\x0d\x35\x8a\xe4\x50\x8f\xc2\x8f\xf0\x83\x3b\x88\xc2\x1c\xbe\xc1\xa4\x3d\xb5\xde\x0d\x77\xe5\x04\xee\x74\x9c\x7e\x5b\x2b\xb5\x22\xe1\x62\x33\xe6\xb5\x40\xff\x8b\xd4\x42\xfd\x9f\x85\x4e\x1c\x97\x87\xe9\x7c\xe3\x4c\x05\xaf\xc9\x4d\x12\x12\xb8\xb7\x7a\x5c\xe2\x36\x69\xd0\x54\xb2\x0c\xf8\x52\x58\x0b\x7b\xe7\x68\xee\xc6\xe5\xd4\x0f\xf4\x6b\xcd\xd9\xb1\x96\xcc\xb0\xc0\xa9\xfd\x4a\x65\xbf\x33\xa5\x2b\x56\x27\x3b\x6a\xdb\xdc\x86\x23\x17\x35\xc7\x7d\x08\x0c\x35\x0f\xc5\x52\x70\x9f\x3c\x7f\x81\x95\xa8\x12\x9b\x95\xd2\x13\xa0\xdb\xdb\xe0\xb9\x15\xe2\x01\xea\xf6\x51\x29\x47\x1b\xf8\x07\x42\x03\x03\x3d\x89\x76\x4c\xc1\xb9\xd5\xec\x28\x3b\xac\x5a\xaf\x48\x07\xee\x0f\xd1\x04\x78\xfe\x65\x7d\x44\x2f\x4c\x2b\x38\xf4\x44\x53\xfb\xed\xf0\x82\x1b\x99\x3b\x25\x2d\xa8\x62\x8e\xfb\x63\x08\x19\xc8\xbf\x01\x00\x00\xff\xff\xa6\xe2\xde\x4b\x99\x08\x00\x00")

func templatesClient_rpc_goTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesClient_rpc_goTpl,
		"templates/client_rpc_go.tpl",
	)
}

func templatesClient_rpc_goTpl() (*asset, error) {
	bytes, err := templatesClient_rpc_goTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/client_rpc_go.tpl", size: 2201, mode: os.FileMode(420), modTime: time.Unix(1558685942, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHandler_rpc_goTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x54\xcf\x4e\xe3\x3c\x10\x3f\x7b\x9e\x62\xe4\xc3\xa7\x06\xf5\x4b\xee\x95\x7a\xd8\x05\x56\xb0\x12\x34\x82\xde\x10\x5a\x5c\x67\x9a\x78\x49\xec\xc8\x76\xca\x56\x56\xde\x7d\xe5\xb4\xd0\x6a\x45\xbb\xd9\x9b\x33\xb6\xe7\xf7\xcf\x93\x2c\xc3\x4b\x53\x10\x96\xa4\xc9\x0a\x4f\x05\xae\xb6\xd8\x90\xab\x1e\xf2\xcb\x14\xaf\x16\x78\xbf\x58\xe2\xf5\xd5\xed\x32\x85\x2c\xc3\x2f\x75\x8d\xb2\x12\xba\x24\x87\x4d\xe7\x3c\xae\x08\x0b\xa3\x09\x95\x46\xd9\x39\x6f\x1a\x94\xb5\x22\xed\xd1\x57\xc2\xa3\xab\x4c\x57\x17\x48\xca\x57\x64\x91\x9a\x15\x15\x68\x2c\xbe\x59\xd1\xa2\xaf\x94\x4b\x01\x5a\x21\x5f\x45\x49\x18\x42\x9a\xef\x96\xf7\xa2\xa1\xbe\x07\x50\x4d\x6b\xac\xc7\x09\x30\xbe\xda\x7a\x72\x1c\x18\x27\x2d\x4d\xa1\x74\x99\xfd\x74\x46\xc7\x82\x32\x1c\x80\xf1\x52\xf9\xaa\x5b\xa5\xd2\x34\x59\xa9\xf4\xff\xa5\xd1\x4a\xc6\x15\x87\x04\xc0\x6f\xdb\xa1\xff\x37\x12\xbe\xb3\x94\x5b\x5a\xab\x5f\x7d\xff\x90\x5f\xde\x08\x5d\xd4\x64\x51\x69\x4f\x76\x2d\x24\x61\x80\x10\xd2\xef\xce\xe8\xfd\xd6\xed\xfb\xce\x57\x53\x6c\xfb\x1e\x7a\x80\x8d\xb0\x67\xbb\x3d\xb6\x24\xcf\xc3\xcd\xf1\xbf\x10\xd2\x43\x21\xb7\x6a\x23\xfc\x5e\x78\xe8\x47\x30\x5e\xb4\x5e\x19\xed\xd0\x79\xdb\x49\x8f\x21\xf2\x5a\x77\x5a\xa2\xac\x48\xbe\x8e\xb8\x39\x31\xad\xc7\x8b\x11\x07\x93\x51\xa7\x30\x00\x53\x6b\x8c\x4d\xe7\x73\xd4\xaa\x8e\x05\x36\x7c\x0e\x5a\xff\x76\x3f\xf4\xc0\x7a\x60\x96\x7c\x67\x75\x6c\xf3\x21\xe8\x9e\xde\xce\xdd\x9f\x00\x73\x9b\xcf\xec\x7e\x24\xbb\x51\x92\xa6\xc0\xc6\x2a\x9d\x42\x72\x3e\xb6\xf0\x41\xf0\x5c\x7c\x3b\xdd\xb3\x7f\x4a\x22\x99\x02\x8b\x3a\x66\xe8\x36\x72\x1a\xad\x38\x7a\x04\x27\x80\x0e\xd9\x0f\x0e\x9c\xb6\x60\xe7\xc0\x28\x0b\x22\xec\x1f\xef\xbf\x69\x6b\x6a\x48\x7b\x11\x0f\xec\x87\xe0\x7d\x04\x4e\x30\xbb\x23\x5f\x99\xc2\xdd\x89\x16\xe7\xd8\x88\xf6\xc9\x79\xab\x74\xf9\xfc\xf4\xbc\x5b\x04\x60\xfc\x82\xcf\xf0\xe8\x9b\xf1\x7c\xf1\xb8\xe4\x51\xfa\xf4\x23\xfa\xc9\x8f\x81\xf4\x09\x94\x04\x6f\x96\xcb\xfc\x80\x35\x49\x3e\xc3\x3a\x0a\x6d\x0c\x61\x38\x3d\x7c\xd7\xd6\x1a\xfb\x40\xae\x35\xda\xd1\x91\xf9\x43\x1d\xf7\x68\x2f\xf1\xbf\x34\xe3\x14\x6b\xfc\x05\x7a\xf8\x1d\x00\x00\xff\xff\xc3\x0a\x65\x13\x5f\x05\x00\x00")

func templatesHandler_rpc_goTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHandler_rpc_goTpl,
		"templates/handler_rpc_go.tpl",
	)
}

func templatesHandler_rpc_goTpl() (*asset, error) {
	bytes, err := templatesHandler_rpc_goTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/handler_rpc_go.tpl", size: 1375, mode: os.FileMode(420), modTime: time.Unix(1558691650, 0)}
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
	"templates/client_rpc_go.tpl": templatesClient_rpc_goTpl,
	"templates/handler_rpc_go.tpl": templatesHandler_rpc_goTpl,
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
		"client_rpc_go.tpl": &bintree{templatesClient_rpc_goTpl, map[string]*bintree{}},
		"handler_rpc_go.tpl": &bintree{templatesHandler_rpc_goTpl, map[string]*bintree{}},
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

