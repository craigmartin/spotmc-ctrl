package data

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _bindata_go = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindata_go_bytes() ([]byte, error) {
	return bindata_read(
		_bindata_go,
		"bindata.go",
	)
}

func bindata_go() (*asset, error) {
	bytes, err := bindata_go_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1424610049, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _gen_go = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x49\x2c\x49\xe4\xe2\xd2\xd7\x4f\xcf\xb7\x4a\x4f\xcd\x4b\x2d\x4a\x2c\x49\x55\x48\xcf\xd7\x4d\xca\xcc\x03\xc9\x28\xe8\x16\x64\xa7\x83\xd5\x28\xe8\x71\x01\x02\x00\x00\xff\xff\x4a\xbc\x83\xf8\x33\x00\x00\x00")

func gen_go_bytes() ([]byte, error) {
	return bindata_read(
		_gen_go,
		"gen.go",
	)
}

func gen_go() (*asset, error) {
	bytes, err := gen_go_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "gen.go", size: 51, mode: os.FileMode(420), modTime: time.Unix(1424609850, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _start_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\xc9\x28\xc9\xcd\xb1\xe3\x52\x00\x02\x9b\x8c\xd4\xc4\x14\x08\x13\xcc\xcd\x4d\x2d\x49\x54\xc8\x28\x29\x29\xd0\x4d\x2d\x2c\xcd\x2c\xb3\x55\x2a\x4a\x4d\x2b\x4a\x2d\xce\x50\x52\x48\xce\xcf\x2b\x49\xcd\x2b\xb1\x55\x32\x34\xb0\x56\x28\x2d\xca\xb1\xd5\x57\x82\x1a\xa1\x8f\x30\xc3\x26\x29\x3f\xa5\x52\xa1\xb8\xa4\x32\x27\xd5\x56\x29\x0d\xa8\x43\xb7\x38\xb3\x2a\xd5\xca\xcc\xa4\xa0\x42\x09\x61\xcb\xe3\xc6\xb6\xc7\x8d\x53\x1e\x37\xf6\x3f\x6e\xea\x00\xb3\x57\x3f\x6e\x6c\x01\xa3\x7d\x8f\x1b\xa7\x3f\x6e\x9c\xff\xb8\xa1\xc9\x26\xa9\x08\x59\xc3\xcc\xc7\x8d\x93\xc1\x72\x5d\x60\x45\x0b\x41\x9a\x1b\x17\x3c\x6e\x9c\x0a\xd2\x86\xa4\xda\x46\x1f\xe4\x02\x3b\x2e\xa0\xa3\xc0\x7e\x04\x04\x00\x00\xff\xff\xbb\x8b\xe6\xdf\xeb\x00\x00\x00")

func start_html_bytes() ([]byte, error) {
	return bindata_read(
		_start_html,
		"start.html",
	)
}

func start_html() (*asset, error) {
	bytes, err := start_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "start.html", size: 235, mode: os.FileMode(420), modTime: time.Unix(1424609850, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _stop_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\xc9\x28\xc9\xcd\xb1\xe3\x52\x00\x02\x9b\x8c\xd4\xc4\x14\x08\x13\xcc\xcd\x4d\x2d\x49\x54\xc8\x28\x29\x29\xd0\x4d\x2d\x2c\xcd\x2c\xb3\x55\x2a\x4a\x4d\x2b\x4a\x2d\xce\x50\x52\x48\xce\xcf\x2b\x49\xcd\x2b\xb1\x55\x32\x34\xb0\x56\x28\x2d\xca\xb1\xd5\x57\x82\x1a\xa1\x8f\x30\xc3\x26\x29\x3f\xa5\x52\xa1\xb8\xa4\x32\x27\xd5\x56\x29\x0d\xa8\x43\xb7\x38\xb3\x2a\xd5\xca\xcc\xa4\xa0\x42\x09\x61\xcb\xe3\xc6\x15\x8f\x1b\xf7\x3d\x6e\xea\x7e\xdc\xd4\xf1\xb8\xb1\xed\x71\xe3\xea\xc7\x8d\x2d\x60\xb4\xef\x71\xe3\xf4\xc7\x8d\xf3\x1f\x37\x34\xd9\x24\x15\x21\x6b\x98\xf9\xb8\x71\x32\x58\xae\x0b\xac\x68\xe1\xe3\xc6\xfe\xc7\x8d\x0b\x1e\x37\x4e\x05\x69\x43\x52\x6d\xa3\x0f\x72\x81\x1d\x17\xd0\x51\x60\x3f\x02\x02\x00\x00\xff\xff\xd8\x02\x3f\xac\xeb\x00\x00\x00")

func stop_html_bytes() ([]byte, error) {
	return bindata_read(
		_stop_html,
		"stop.html",
	)
}

func stop_html() (*asset, error) {
	bytes, err := stop_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "stop.html", size: 235, mode: os.FileMode(420), modTime: time.Unix(1424609850, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _top_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\xc9\x28\xc9\xcd\xb1\xe3\x52\x00\x02\x9b\x8c\xd4\xc4\x14\x08\x13\xcc\xcd\x4d\x2d\x49\x54\xc8\x28\x29\x29\xd0\x4d\x2d\x2c\xcd\x2c\xb3\x55\x2a\x4a\x4d\x2b\x4a\x2d\xce\x50\x52\x48\xce\xcf\x2b\x49\xcd\x2b\xb1\x55\x32\x34\xb5\x56\x28\x2d\xca\xb1\xd5\x57\x82\x1a\xa1\x8f\x30\xc3\x26\x29\x3f\xa5\x52\xa1\xb8\xa4\x32\x27\xd5\x56\x29\x0d\xa8\x43\xb7\x38\xb3\x2a\xd5\xca\xcc\xa4\xa0\x42\x09\x61\x4b\x75\x75\x66\x9a\x82\x5e\x68\x41\x6d\x2d\x5c\x08\x21\x1c\x5c\x92\x58\x54\x92\x9a\x82\x26\xf7\xb8\xb1\xed\x71\xe3\x94\xc7\x8d\x2d\x8f\x1b\x97\x81\xc9\x7d\x8f\x1b\x67\xda\x24\x15\xd9\xa1\xa8\xb2\x01\xba\x1d\xe8\x60\x5b\x25\xfd\xe2\x92\xfc\x02\x25\xbb\xc7\x8d\x2b\x1e\x37\x35\x3e\x6e\xea\xb6\xd1\x4f\xb4\x43\xb3\x2b\x35\xa7\x38\x15\xc3\x92\x15\x20\x83\x9b\x7a\xc1\xb6\x01\xd9\xd3\xf1\xda\x06\x96\xdf\xf0\xb8\xa9\xf3\x71\x63\x3f\x58\xc9\x62\xb0\x72\x20\x7b\xc1\xe3\xc6\xa9\x40\x7d\x18\x3a\x80\xb6\xe6\x21\xfb\x0c\xab\x2b\x88\x09\x86\x5e\xb0\x1d\x34\x70\x1e\xce\x40\x81\xea\x26\x26\xe4\x81\x0e\x07\x05\x3d\xc4\xa5\xdd\x20\xd5\x58\x42\x1f\x3d\x1c\xe0\x7c\x1b\x7d\x50\x1a\xb2\xe3\x02\x26\x2b\x70\x2a\x05\x04\x00\x00\xff\xff\x77\x38\x34\x80\xad\x02\x00\x00")

func top_html_bytes() ([]byte, error) {
	return bindata_read(
		_top_html,
		"top.html",
	)
}

func top_html() (*asset, error) {
	bytes, err := top_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "top.html", size: 685, mode: os.FileMode(420), modTime: time.Unix(1424609850, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"bindata.go": bindata_go,
	"gen.go": gen_go,
	"start.html": start_html,
	"stop.html": stop_html,
	"top.html": top_html,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"bindata.go": &_bintree_t{bindata_go, map[string]*_bintree_t{
	}},
	"gen.go": &_bintree_t{gen_go, map[string]*_bintree_t{
	}},
	"start.html": &_bintree_t{start_html, map[string]*_bintree_t{
	}},
	"stop.html": &_bintree_t{stop_html, map[string]*_bintree_t{
	}},
	"top.html": &_bintree_t{top_html, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

