package boilerplate

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
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

var _functions_hello_index_js = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\x41\xaa\x83\x30\x10\x87\xf1\x7d\x4e\xf1\xdf\x3c\x92\x40\xf0\x00\xc2\x3b\x8c\xc6\xa9\x5a\x86\x19\xc9\x8c\xad\x20\xde\xbd\xb8\x29\x74\xfd\x7d\xfc\xaa\x8a\x29\x53\xc7\x3a\xa7\x68\x3e\x34\x5f\x65\xc6\x63\x97\xea\xab\x4a\xcc\x81\x8e\x4d\x9b\x5b\xb7\x0c\x32\x31\xe1\xff\xdb\x12\x15\x54\x3f\x0a\xea\x98\x71\x06\xe0\x87\xda\x9a\x56\x32\xbb\x31\x7a\x91\x78\x8f\xbf\x67\x2c\xa0\x7c\x8f\x63\x92\x9d\xb9\xe0\xc4\x42\xcc\xda\x23\xbe\xb5\xf1\x14\x71\xe5\x70\x85\x4f\x00\x00\x00\xff\xff\x92\xa4\x17\xbd\x93\x00\x00\x00")

func functions_hello_index_js() ([]byte, error) {
	return bindata_read(
		_functions_hello_index_js,
		"functions/hello/index.js",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"functions/hello/index.js": functions_hello_index_js,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"functions": &_bintree_t{nil, map[string]*_bintree_t{
		"hello": &_bintree_t{nil, map[string]*_bintree_t{
			"index.js": &_bintree_t{functions_hello_index_js, map[string]*_bintree_t{
			}},
		}},
	}},
}}
