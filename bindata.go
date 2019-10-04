// Code generated by go-bindata.
// sources:
// prom-templates/deployment.yaml
// prom-templates/kustomization.yaml
// prom-templates/route.yaml
// prom-templates/service.yaml
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

var _promTemplatesDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\x4d\x73\xd3\x40\x0c\xbd\xe7\x57\xe8\xc0\xd5\x89\xd3\x0e\x87\xee\x8d\xa1\x33\xe5\x40\x21\x53\x32\xdc\x95\xb5\x1a\x2f\xf5\x7e\x20\x69\xd3\x84\xb6\xff\x9d\x71\xdd\xd8\x1b\x28\x97\xce\xb0\xa7\xcd\x7b\xfb\x9e\xa4\x17\x19\x93\xfb\x4e\x2c\x2e\x06\x03\xb4\x57\x0a\xfd\x55\x16\xbb\xe5\x86\x14\x97\xb3\x3b\x17\x1a\x03\x97\x94\xba\x78\xf0\x14\x74\xe6\x49\xb1\x41\x45\x33\x03\x08\xe8\xc9\x40\xe2\xe8\x67\x92\xc8\xf6\x10\x53\xea\x9c\x45\x31\xb0\x9c\x01\x28\xf9\xd4\xa1\x52\xcf\x00\x94\xd2\xfe\x14\xf2\x01\xe8\x70\x43\x9d\x1c\x69\x00\x4c\xa9\xe0\x8f\x25\xfa\xe3\x82\xd3\x8f\x31\x28\xba\x40\x3c\x2a\xaa\x17\x4b\xeb\xaa\x5b\x52\xdb\x12\x8f\x56\xce\xe3\x96\x0c\x30\x6d\x9d\x28\x1f\xe6\xb7\xd4\x44\xc6\xc4\xf1\x07\x59\x9d\x47\xde\x2e\x06\xc4\x9c\xd7\xa3\xc6\x46\xef\x31\x34\x53\x3f\x15\x2c\x36\x2e\x2c\x36\x28\x6d\x81\x55\xb6\xf8\xf1\x38\xde\x01\x84\x14\xaa\xbc\x8f\x90\x5c\xa2\x5b\x74\x5d\xc1\x65\x8f\x72\x07\x75\x5d\xd7\x05\x68\x33\x77\x50\xc9\x67\x78\xf7\xb0\xba\xf9\x7a\xbd\xfe\x70\xf3\x04\x8f\xa0\xc8\xb0\xdf\xfd\x82\xca\x8f\x4f\xef\x23\xdf\xb9\xb0\xbd\x74\x6c\x60\xd1\x07\x44\xda\x52\x96\xc5\xf8\x60\x17\xbb\xec\xe9\x3a\xe6\xa0\x52\xf6\x3f\x45\x3e\x28\x2a\xd1\xc8\xb8\xa5\x6a\x10\x14\xbd\xf8\x5e\xbb\x42\x6d\x5f\xad\x60\xff\x99\xfd\xf4\xf6\xcf\xec\x7b\xa6\xb0\x32\xbb\xb3\xf9\xf2\x6c\x3e\xcd\x9f\x22\xbf\xd6\xec\x3d\x6d\xb2\x2b\x43\x3a\x96\x5e\x45\x56\x03\x17\xf5\xc5\x64\xc1\x84\x8d\x0b\x24\xb2\xe2\xb8\x21\x53\xa8\x5a\xd5\x74\x45\x5a\x42\x00\x69\x18\xef\x14\xfb\xdb\xf5\xf9\xbf\xb4\x2d\xf5\xdd\x7c\x5a\xaf\x57\x05\xa1\xce\x53\xcc\xfa\x8d\x6c\x0c\xcd\xcb\xd2\x8f\x4e\xc4\x2e\x36\x13\x55\x1a\x4a\xb6\x96\x44\xd6\x2d\x93\xb4\xb1\x6b\x4e\x95\xfd\xae\x64\xa6\x82\x3d\x2f\x46\x94\x98\xd9\x92\x94\xa3\x30\xfd\xcc\x24\x2a\xa7\xe3\xd9\x94\x0d\xbc\xaf\x6b\x7f\x82\x7a\xf2\x91\x0f\x06\xce\xae\xca\x54\x3b\xe7\xdd\xdb\xf5\xff\x7b\xdd\xa4\x45\xa6\x15\xc7\x3e\xb3\x2f\xe8\x49\x12\x5a\x32\xa0\x9c\x8f\x1e\x83\xe1\x1b\x8b\x93\x4f\x7a\x78\xfe\x98\x1e\x9e\x66\xbf\x03\x00\x00\xff\xff\xc7\xb5\xad\x6b\x11\x05\x00\x00")

func promTemplatesDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_promTemplatesDeploymentYaml,
		"prom-templates/deployment.yaml",
	)
}

func promTemplatesDeploymentYaml() (*asset, error) {
	bytes, err := promTemplatesDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prom-templates/deployment.yaml", size: 1297, mode: os.FileMode(420), modTime: time.Unix(1570199404, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _promTemplatesKustomizationYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\x49\x4f\xc3\x30\x10\x85\xef\xfe\x15\xa3\xde\x0d\xe2\xea\x5b\x80\xde\x12\x52\xb5\xd0\x6b\x35\x38\xaf\xc5\xc2\x9b\x6c\x27\xa2\x12\x3f\xbe\xca\xa2\x2e\x73\x7b\xef\xf3\xf2\x69\x74\x70\x2e\xf8\x9a\xbf\x61\x73\xe5\xbb\x1d\x2c\x74\x09\x29\x2b\x41\xc4\x31\x2a\x7a\x6b\x9b\xa6\xfd\x38\xd4\xd5\xeb\xba\x16\x9e\x1d\x36\x09\x47\xf3\xf7\x08\xa4\x48\xc8\xa1\x4f\x1a\xd3\x45\x49\x1d\xa2\x0d\x67\x07\x5f\x9e\xce\xec\xec\xd4\xa5\xd0\x17\xdc\x62\x46\x1a\x8c\x5e\x8a\xc8\x45\xff\x20\xef\x4a\xe2\x82\x93\xd1\x0d\xd2\x09\x4a\x48\xfa\x97\x93\x88\xd9\x23\x65\x13\xbc\x1a\xa5\xf2\xf3\xf0\x22\x88\x7e\x8d\xef\x14\xbd\x5f\x7f\x12\x44\x0e\x85\x3b\x2e\x3c\x4a\x10\x8d\xb6\x8a\x62\x0a\x4e\x10\xe5\x08\x3d\xd7\x05\x2e\x5a\x2e\x98\xd3\x3d\x19\x47\x07\x5f\xd8\x78\xcc\x2b\x98\x47\x2e\x4f\x69\x23\x8f\x18\x45\xd3\x15\x11\xc1\x0f\xea\x2e\xde\x4e\x6f\xb6\x6d\xf3\x59\x6d\x1f\x18\xd1\xc0\xb6\x87\xa2\xd5\x42\x0f\xfb\xaa\xfe\x5a\xaf\xc4\x25\x00\x00\xff\xff\x36\x0d\x22\xd5\x88\x01\x00\x00")

func promTemplatesKustomizationYamlBytes() ([]byte, error) {
	return bindataRead(
		_promTemplatesKustomizationYaml,
		"prom-templates/kustomization.yaml",
	)
}

func promTemplatesKustomizationYaml() (*asset, error) {
	bytes, err := promTemplatesKustomizationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prom-templates/kustomization.yaml", size: 392, mode: os.FileMode(420), modTime: time.Unix(1570197840, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _promTemplatesRouteYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xc1\x4e\x03\x31\x0c\x44\xef\xf9\x0a\x7f\x41\xd9\x1e\x9b\x3b\x57\x54\x15\xc4\xdd\x4a\x86\xd4\x62\x13\x47\x8e\xdb\x8a\xbf\x47\xd9\x05\xc1\x25\xd2\xbc\x37\x19\xf9\x53\x5a\x8e\x74\xd1\x9b\x23\x70\x97\x77\xd8\x10\x6d\x91\x6c\x92\x83\x76\xb4\x71\x95\x0f\x3f\x88\x3e\xdd\x8f\xa1\xc2\x39\xb3\x73\x0c\x44\x8d\x2b\x22\x75\xd3\x1a\x46\x47\x9a\xc8\x75\xbe\x44\xfb\xe8\x2b\xec\x2e\x09\x1b\xf9\x57\x9e\xf1\x01\x29\x57\x8f\x74\x5c\x96\x40\xd4\xd5\x7c\xff\xe8\x6c\x05\x7e\x9e\x99\x4e\xcb\x69\x4a\x5f\xc7\x8f\x83\x55\x69\xec\xdb\x79\xc8\x65\x1f\x96\x36\x90\x6e\x86\xe7\x5c\xf0\xf6\xd7\x38\xeb\x2a\xe9\x2b\xd2\x05\x59\x0c\xc9\x03\xd1\x43\xd6\x9c\xd8\xf2\xaf\x7a\xd1\x86\xf0\x1d\x00\x00\xff\xff\x84\xf0\x9a\x83\xff\x00\x00\x00")

func promTemplatesRouteYamlBytes() ([]byte, error) {
	return bindataRead(
		_promTemplatesRouteYaml,
		"prom-templates/route.yaml",
	)
}

func promTemplatesRouteYaml() (*asset, error) {
	bytes, err := promTemplatesRouteYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prom-templates/route.yaml", size: 255, mode: os.FileMode(420), modTime: time.Unix(1570192158, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _promTemplatesServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcc\x31\xae\xc2\x30\x0c\xc6\xf1\x3d\xa7\xf8\x2e\xf0\xa4\xbe\xb1\x5e\xb9\x00\x12\x88\xdd\x4d\x3d\x44\x24\xb1\x95\x98\x72\x7d\x94\x92\x89\xd1\x3f\x7f\xfa\xb3\xa5\x87\xb4\x9e\xb4\x12\x8e\xff\xf0\x4c\x75\x27\xdc\xa4\x1d\x29\x4a\x28\xe2\xbc\xb3\x33\x05\xa0\x72\x11\x82\x35\x2d\x01\xc8\xbc\x49\xee\x83\x01\x36\x9b\xde\x4d\xe2\x30\xd3\xe6\xf3\xf9\x77\x1e\x84\x75\x59\x97\x13\x30\xa6\xae\x51\x33\xe1\x7e\xb9\x4e\xfb\xc6\xdf\xb2\xbd\x52\x00\xba\x64\x89\xae\xed\xb7\xff\x09\x00\x00\xff\xff\x6d\xe1\xdb\x34\xac\x00\x00\x00")

func promTemplatesServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_promTemplatesServiceYaml,
		"prom-templates/service.yaml",
	)
}

func promTemplatesServiceYaml() (*asset, error) {
	bytes, err := promTemplatesServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prom-templates/service.yaml", size: 172, mode: os.FileMode(420), modTime: time.Unix(1570192162, 0)}
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
	"prom-templates/deployment.yaml": promTemplatesDeploymentYaml,
	"prom-templates/kustomization.yaml": promTemplatesKustomizationYaml,
	"prom-templates/route.yaml": promTemplatesRouteYaml,
	"prom-templates/service.yaml": promTemplatesServiceYaml,
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
	"prom-templates": &bintree{nil, map[string]*bintree{
		"deployment.yaml": &bintree{promTemplatesDeploymentYaml, map[string]*bintree{}},
		"kustomization.yaml": &bintree{promTemplatesKustomizationYaml, map[string]*bintree{}},
		"route.yaml": &bintree{promTemplatesRouteYaml, map[string]*bintree{}},
		"service.yaml": &bintree{promTemplatesServiceYaml, map[string]*bintree{}},
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

