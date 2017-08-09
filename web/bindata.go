// Code generated by go-bindata.
// sources:
// .build/assets/close.svg
// .build/assets/edit.css
// .build/assets/edit.html
// .build/assets/edit.js
// .build/assets/index.js
// .build/assets/links.css
// .build/assets/links.html
// DO NOT EDIT!

package web

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _closeSvg = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8e\xc1\x8a\x84\x30\x0c\x86\x5f\x25\xe4\xbe\x69\x9b\x62\xcd\x2e\xd6\xc3\x9e\xdd\x87\x58\x50\x5a\xc1\x19\x65\x2c\x56\xe6\xe9\x87\x56\x2f\x3f\xe1\xfb\xbf\x84\x74\xfb\x11\xe0\x7c\x2c\xcf\xdd\x63\x4c\x69\xfb\x51\x2a\xe7\x4c\xd9\xd2\xfa\x0a\x8a\xb5\xd6\x6a\x3f\x02\x42\x9e\xc7\x14\x3d\x5a\x87\x10\xa7\x39\xc4\x74\xcd\xc7\x3c\xe5\xdf\xf5\xf4\xa8\x41\x83\x75\x60\x1d\xf6\xdd\xf6\x9f\x22\x8c\x1e\xff\x58\xa8\x81\x6f\x72\x3c\xb0\x23\x2b\xd0\x52\x03\x46\xc0\x34\x24\x52\x79\x25\xed\x2d\xdd\xdc\x5c\x5e\xdd\x58\x98\x0c\x43\x89\xc1\x08\xb0\x26\xc3\x8b\x94\x4b\x35\x0a\xff\xaa\x65\x6d\xc0\xc8\x1b\x55\xdf\x95\x87\xfb\x4f\x00\x00\x00\xff\xff\xe5\x79\xbd\x91\xd8\x00\x00\x00"

func closeSvgBytes() ([]byte, error) {
	return bindataRead(
		_closeSvg,
		"close.svg",
	)
}

func closeSvg() (*asset, error) {
	bytes, err := closeSvgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "close.svg", size: 216, mode: os.FileMode(420), modTime: time.Unix(1502237806, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _editCss = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\xdf\x6e\xe2\x3c\x10\xc5\xef\xbf\xa7\xb0\xca\x45\x41\xc2\xa9\xa1\x85\x7e\xd8\xd2\x3e\xc4\xde\xed\xa5\xe3\x4c\xc2\xa8\x8e\xc7\xb2\x1d\x08\x8d\x78\xf7\x55\xfe\x40\xb3\xdb\xae\x56\xab\x08\x14\xec\x1f\x9e\x99\x73\x8e\x73\x2a\x2e\x5d\xae\xcd\x5b\x15\xa8\x71\x85\x5c\x94\x65\xa9\x4a\x72\x89\x97\xba\x46\x7b\x91\x8f\xdf\xb5\x85\xb3\xbe\x3c\xae\x59\xd4\x2e\xf2\x08\x01\x27\x22\xe2\x3b\xc8\x97\xad\x6f\xc7\x9f\x67\xc0\xea\x98\xe4\xb3\x10\xd7\x92\x42\xdd\x25\x68\x13\xd7\x16\x2b\x27\x0d\xb8\x04\xe1\xba\xc8\x75\xe8\xce\x58\xa4\xa3\xdc\xef\xfa\xff\xd5\x3a\x54\xe8\xa4\x60\xba\x49\xa4\x3c\x45\x4c\x48\x4e\x06\xb0\x3a\xe1\x09\xae\x8b\x26\xd8\xee\x5f\xba\x79\xfe\xa2\x1b\x35\x55\x14\xc2\xb7\xca\xeb\xa2\x40\x57\xc9\xed\xce\xb7\xca\x90\xa5\x20\x17\x87\xc3\x41\xe5\x14\x0a\x08\x3c\xe8\x02\x9b\x28\x5f\x7c\x3b\xad\xc8\x8d\x6f\x59\x24\x8b\x05\x5b\x18\x63\x14\x35\xc9\xa2\x03\xe9\xc8\x81\xca\xa9\xe5\xf1\xa8\x0b\x3a\x4b\xc1\xb6\xbe\x65\x7b\xdf\xb2\x50\xe5\x7a\x29\xd6\xc3\x93\x6d\x57\xc3\x0c\xb2\x24\xd3\xc4\xee\xf3\x91\xe2\x50\x8e\x80\xe4\x67\xc8\xdf\x30\x71\x74\xbe\x49\xdc\x5b\x6d\xe0\x48\xb6\x80\xd0\x4d\x5d\x16\x45\x71\x43\x6b\x7a\xff\x23\x61\x6a\xdf\xfd\x32\xe4\x38\xfd\x6e\x2f\x3e\xe9\x3d\x33\xc8\x42\x99\x66\x32\x6e\x37\xbe\x55\x29\x68\x37\x19\x32\xbc\xf6\xa6\xb2\xad\x10\x75\x64\xa0\x23\x70\x74\x9c\x9a\xa4\xee\x7b\x32\x1a\x6d\xe1\xc7\x52\xac\x3e\xd6\x38\x05\xec\x2b\x26\xf2\x6c\x0c\xc1\x4d\xe9\x9c\x52\xa2\x9a\xf7\x85\x3f\xab\x7e\xdb\x0d\xbd\x87\x7f\x35\x65\x18\x3a\xb3\xe8\xde\x66\x49\xe6\x93\x28\xe5\xbe\x7f\x46\xa4\x6c\xcc\x97\x08\xc0\x2c\x0a\x03\xfa\x4d\xdf\x44\x15\x87\x72\x14\xaa\x00\x43\x41\x0f\x72\xf4\xe6\x8f\x58\x76\x74\x69\x26\xff\x48\x4e\x91\xe8\x9b\xec\x3f\x82\xf5\x97\x6a\xc6\x97\x96\x74\x92\xc3\x6c\xd7\x85\xb1\xb1\xbb\x07\x5f\xe7\x91\x6c\x93\x40\x25\xf2\x52\xa8\x01\x91\x42\x8d\x6a\xc8\x7b\x92\x07\x25\x3e\xc6\xc0\x5a\x57\x20\x9b\x60\x97\x0f\x4f\xf1\xc9\x58\x8a\x90\xc5\x53\xf5\xb0\x9a\x43\xf7\x1a\xa3\x0d\x77\x37\x3e\x88\x00\x1e\x74\x92\x8e\xa6\xb7\xf9\xde\x78\xd5\xff\xf7\x2d\xeb\xbf\x94\x69\x42\xa4\x20\x3d\xe1\x70\xc8\x2c\x29\xe4\xb5\xc1\x74\x99\x72\x62\x9a\x1c\x0d\xcf\xe1\x1d\x21\x2c\x45\xf6\xfa\xba\x66\x62\xcd\x44\xb6\x79\xdd\xad\xd9\x66\xa5\x26\x5a\x0a\x35\x1d\xc5\xe1\x04\x2e\xc5\x9b\xc2\x36\x66\x27\x8c\xdd\x1d\xcb\x9e\x7f\x07\xd1\x1d\x21\xe0\xa8\xa3\x3c\xd2\x09\xc2\x8c\xde\x5f\xff\xfb\x19\x00\x00\xff\xff\x56\xb7\x20\x4a\xe0\x04\x00\x00"

func editCssBytes() ([]byte, error) {
	return bindataRead(
		_editCss,
		"edit.css",
	)
}

func editCss() (*asset, error) {
	bytes, err := editCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "edit.css", size: 1248, mode: os.FileMode(420), modTime: time.Unix(1502237806, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _editHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\xb1\x72\xe3\x20\x10\xed\xfd\x15\x7b\xd4\x67\xa3\x99\xbb\xe2\xc6\x07\x4a\xe1\x78\x52\x26\x93\x71\x93\x12\xa3\x95\x21\x41\x82\xc0\xca\x89\xfe\x3e\x83\x50\x6c\x65\xd2\x85\x86\xdd\x7d\xec\x7b\x8f\x05\xf1\xeb\xf6\x7e\x77\x78\x7a\xd8\x83\xa1\xce\xd5\x2b\x51\x36\x00\x61\x50\x35\x39\x00\x10\x64\xc9\x61\x7d\xe7\x05\x2f\x51\xa9\x76\x48\x0a\x0c\x51\x58\xe3\xeb\x60\xcf\x92\xed\x7c\x4f\xd8\xd3\xfa\x30\x06\x64\xa0\x4b\x26\x19\xe1\x3b\xf1\x4c\xfb\x1f\xb4\x51\x31\x21\xc9\x81\xda\xf5\x3f\xc6\x67\x22\x67\xfb\x17\x30\x11\x5b\xc9\x78\xe2\xd8\x58\xda\xe8\x94\xd8\x04\xe6\x15\xd1\x49\x96\x68\x74\x98\x0c\x22\x5d\x01\x1a\x03\xce\xfc\xb9\xe1\x3b\x5d\x76\xb7\xe5\xbc\xf5\x3d\xa5\xcd\xc9\xfb\x93\x43\x15\x6c\xda\x68\xdf\xe5\x8e\x9b\x56\x75\xd6\x8d\xf2\x51\x39\x7c\x53\xe3\xf6\x6f\x55\xfd\xfe\x53\x55\x3f\x53\x16\xfc\x73\x62\xe2\xe8\x9b\x71\x36\xd3\xfa\xd8\x81\x1a\xc8\x6b\xdf\x05\x87\x84\x92\xf9\xb6\x9d\xad\x02\x88\xc6\x9e\xc1\x36\x92\x1d\x55\xbc\x14\x17\x65\xed\x12\xab\x05\x6f\xec\x79\x01\xda\x3e\x0c\xb4\xb0\xc0\xa6\xa3\x43\x74\x0c\x82\x53\x1a\x8d\x77\x0d\x46\xc9\xf6\x3d\x61\x04\x32\x08\x43\x74\x40\x1e\x92\xf1\x91\xb0\xcf\x8c\x13\xc7\xc5\xc5\x52\xe0\xaa\xdd\x85\x2f\xda\x82\xe7\xcb\xd4\xab\x92\x24\x1d\x6d\x20\x48\x51\x5f\x5f\xed\x79\x32\x5b\x90\x32\x92\x32\x09\xc1\xcb\xaf\xfa\x08\x00\x00\xff\xff\x3a\x69\xb4\x3e\x6d\x02\x00\x00"

func editHtmlBytes() ([]byte, error) {
	return bindataRead(
		_editHtml,
		"edit.html",
	)
}

func editHtml() (*asset, error) {
	bytes, err := editHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "edit.html", size: 621, mode: os.FileMode(420), modTime: time.Unix(1502237806, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _editJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x51\x53\xdb\xb8\x13\x7f\xe7\x53\x50\xfd\x67\x18\x69\x10\x0a\xb4\x2f\x1d\xfc\xf7\x30\x94\xe6\xae\x73\xd3\x1e\x1d\xe0\xe1\x6e\x7a\x7d\x58\x5b\x6b\x5b\x87\x22\x19\x49\x86\xe4\x12\x7f\xf7\x1b\xd9\x4e\x70\x20\x94\xeb\x43\x3c\x92\x76\xf5\xdb\x5d\xed\x6f\x77\x73\x0f\x6e\x5f\xda\x59\x42\x8b\xc6\xe4\x41\x59\x43\x73\xb6\xcc\xc5\x5d\xba\xd9\x67\x6c\xe9\x30\x34\xce\xec\x4b\x9b\x37\x33\x34\x41\xdc\x35\xe8\x16\xd7\xa8\x31\x0f\xd6\xd1\x8c\xb5\x49\x2e\xee\xe0\xbf\x5f\x39\xd7\x7a\xb8\x95\xff\xf0\x52\xee\x10\x02\x4e\x35\xc6\xdd\xfa\x86\xf7\xa3\x3b\x1c\xb8\x64\xcb\x4c\xf8\xb0\xd0\x28\x3c\x86\xaf\xce\xd6\xe8\xc2\x82\x02\x97\x9c\x10\xd6\xb6\x8c\x4a\x3b\x5b\xad\xe2\x37\x5d\xb6\x8c\x25\x31\xe4\x79\xe5\x92\xbd\xad\x98\xe3\x69\xf6\x88\xcc\x96\xeb\xe5\x3e\x50\xd9\x8b\x21\x0d\x95\xf2\x49\xfc\x88\x79\xe5\x52\xd9\x2f\xa5\x35\xf8\x8b\xf1\xe9\xb7\xef\xfd\x1e\x9d\xb3\x6e\x38\x90\xc2\x1a\x6d\x41\x8e\x81\x7b\x4b\x52\x38\xf4\xb5\x35\x1e\x6f\x70\x1e\x78\x9e\x4a\xe1\x03\x84\xc6\x27\xb0\x46\x14\x85\x75\x53\xc8\xab\x47\x3f\x81\x2d\x81\x66\x3c\x67\x2d\x6b\x3b\xec\xce\xd6\x18\x1c\x36\xe6\x77\xde\x1e\x5e\x18\x68\x44\x68\x41\xd4\xce\x06\x1b\x16\x35\x0a\x6b\x3e\x5a\x83\xe9\x58\x79\x1c\x9d\xa8\x1b\x5f\x51\x60\xc9\x80\x10\x65\x6d\xb2\x0d\x30\xdd\x76\x66\x8d\xb0\x71\xe8\x55\x88\x07\x15\xaa\x4f\x08\x12\xc7\x28\x3c\x1b\x70\xe6\x95\x8b\x19\xbe\xc2\xbb\x06\x7d\xe8\xf5\x3a\xf1\xcb\x80\x1e\x8d\xfc\xed\xfa\xf2\xf7\xe7\x4e\x3d\x9a\xa2\xe4\xc2\x9a\x80\x26\x1c\xdd\x2c\x6a\x24\x7c\x8f\x40\x5d\x6b\x95\x43\xd4\x9f\xfc\xed\xad\x49\xf2\x0a\x9c\xc7\x90\x36\xa1\x78\x4f\x58\x32\xf2\xc6\x48\x1a\xf1\x85\x0f\x4e\x99\x52\x15\x0b\x0a\xec\x15\x7f\x9e\xfb\xb2\x81\x7a\xfa\x36\xeb\x6c\xb5\x94\x25\xb9\xb8\xc2\xbb\x34\x8b\x15\xd0\x55\xc5\xf8\x85\x06\xfa\xca\xd4\xe0\xc3\xfe\x1f\x5f\x3e\x7f\x0a\xa1\x1e\x9e\x29\xb2\xa4\xc6\x4e\x89\xbf\x39\xde\xe0\x47\xc5\x8c\xca\xae\xa2\x4a\x0c\xe9\x0e\x8e\xac\x0d\x51\xf2\xeb\xf4\x86\x70\xe8\x74\x6b\xeb\x5f\x51\xfe\x7a\x79\xdd\x6b\xb7\x8c\xce\x2b\xb7\x5a\xc5\xef\x63\xe1\x95\xf6\xb5\xba\xeb\x2b\xad\x8c\x6e\x7d\xb0\x8d\x91\xca\x94\x17\x5a\xa1\x09\x57\x98\x07\xca\x12\x69\x67\xb1\x09\xd0\x92\x93\x19\xb8\x52\x99\xa3\x60\x6b\xc2\x1f\x94\x91\xf6\x41\x28\x63\xd0\x7d\x42\x55\x56\x61\xf2\xee\x08\x44\xd5\x2f\xdf\x1e\x92\x7a\x4e\x58\xcb\xe1\xb9\x29\x5a\x88\x7b\xd0\x0d\xae\x56\x84\x30\x11\x9c\x9a\x51\x96\xc0\x9b\x74\x76\x70\x40\x67\x29\xf0\x5b\xca\x38\x9c\x55\x22\xd7\xe0\xfd\x67\xe5\x83\x00\x29\x29\xb9\x57\x9e\xb0\xd3\xf1\xb1\xc3\x99\xbd\xc7\x41\xc2\x5a\xbe\x9d\xea\xc8\x03\xbc\x47\x13\x3e\x62\x01\x8d\x8e\xa1\x40\xaa\x6d\xcf\x33\x51\x43\xa8\x0c\xcc\x50\xf8\x26\xeb\xc9\x44\x4f\x98\xf0\xb5\x56\x81\x92\x09\x61\xdf\x4e\xbe\x27\xfd\x53\xed\x74\x37\x52\x28\xe6\x86\x92\x09\xd4\x6a\xd2\x38\x3d\x21\x87\xc0\x36\x05\x40\x97\x8d\xd3\xa7\x59\xcb\x86\x3a\xa7\xdb\xf5\x05\x69\xc7\xe2\x3a\xd2\x3c\xb2\x10\x84\xbd\x3d\xa3\x90\x82\x70\xb6\x09\xc8\xce\x68\x96\x82\x88\xfe\x45\xab\x1c\x44\xe3\xf4\xc1\x01\xad\x94\x0f\xd6\x2d\x84\xc3\x5a\x43\x8e\xd7\x21\x32\x60\xd9\x72\xd3\x68\xcd\xc9\x04\xa5\x0a\x13\x72\x98\x31\x9e\x3d\x06\x6a\x9d\x2a\x95\x39\xdc\x23\x51\xc2\x51\x04\x9c\x87\xa1\x00\x53\x42\x38\xee\x78\xcf\xa2\xc9\x6f\x09\xdb\x12\x75\x19\xd0\xca\xc4\x73\x48\x3b\x4a\x50\x02\x71\x13\x5b\xc4\x79\x08\x4e\x65\x4d\xa4\x63\xe5\xb0\x20\x3c\x8b\x82\xb1\xa5\x68\x19\xea\x1a\x8d\xbc\xa8\x94\x8e\x95\xc7\xb3\x35\x8c\xaf\xc1\x10\xc6\xb3\xa7\xe6\x2a\x13\xba\xe3\x31\xce\xf1\xff\x53\x03\xf7\xaa\x84\x60\x9d\x68\x3c\xba\xf3\x32\x0e\x2e\x65\x24\xce\x2f\x0b\x4a\xbe\x40\xae\x4c\xb0\xbe\x22\xec\x8c\xfc\xd5\xbc\x7d\x77\xf2\xfe\xe8\x82\x9c\x92\x8b\xe0\xf4\xd1\x05\x79\xe2\x45\xc6\xf8\x9a\xdd\xc8\x49\x70\x60\x7c\x61\xdd\x8c\x70\xe2\x73\xd0\xf8\x27\x3d\x61\x84\xf1\x12\x43\x3f\x48\x3b\x12\xc7\x78\x3f\x80\xc7\x73\x23\xa7\xf3\xe8\x13\x05\x7e\xcc\x81\x9f\x30\xc6\x4e\x6f\x29\x3b\xd5\x74\x98\x09\xb1\xe5\x73\xf3\x7c\x12\xfd\x14\x07\x77\xb2\x6f\x38\x4b\x09\x49\x80\xb2\x44\x15\x34\x63\x91\x91\xeb\x9e\xf0\x71\xfa\x79\x7a\x33\x25\x7c\x4c\xce\xbd\xac\x67\x27\xfd\x09\x4a\xae\x56\xdb\xc1\xdc\x8e\x83\xf9\xf1\xcb\x1d\xb3\x58\xfd\x7a\xab\x22\x9f\x92\x2f\xd9\x45\xbe\x9e\x64\xc9\x33\xf2\xf5\xa4\x1c\x6a\x72\x8b\x39\xc9\x36\x45\xc8\xf4\xea\xea\xf2\xea\x74\x9f\x1c\x42\xf2\x34\xdd\xc9\xeb\xe9\x6e\x79\xd9\xc1\xdf\x51\xd2\x09\x19\xc7\xf5\xfe\x7f\xf9\xac\x26\x8c\x57\x8f\x7b\xed\x09\xe3\xc5\x66\xdf\x38\x4d\x18\x1f\xff\xb9\x63\xcb\x8c\xb2\x64\xe8\x95\x20\xe5\x34\x36\xa4\x18\x14\x9a\x38\x09\x1d\x7a\xf5\x0f\x12\x9e\xf1\x37\x27\x2c\x29\x77\x68\xf8\x26\x9b\xa9\x40\xb8\xec\x34\x8a\x1d\x1a\xb7\xb8\x68\x6a\xc2\xe1\x45\x85\x1a\x7c\x88\x73\xf6\x65\x8d\xbc\x02\x53\xe2\x1a\xa3\xda\xa5\xa1\x55\x7e\x4b\xb8\xe9\x14\x62\x0e\xf2\x9f\xa3\xb1\x2a\x68\xde\x51\xb4\xc4\xed\x9e\x99\xbf\xc8\xca\xf8\xb7\x6b\x99\x8d\x59\x99\xb1\xe4\xed\xf1\x71\x9a\xe6\x07\x07\xeb\xb2\x48\xb3\xbe\x63\xc6\x06\xd9\xb5\xca\x42\x14\x36\x6f\x7c\x1c\x21\x94\xb1\x96\x25\xa8\x3d\xee\x6f\x4e\x5b\xd6\xfd\x4a\xbb\x5a\xd1\xd2\xf6\x63\x72\xef\xdf\x00\x00\x00\xff\xff\x55\x24\x10\x3b\x94\x0b\x00\x00"

func editJsBytes() ([]byte, error) {
	return bindataRead(
		_editJs,
		"edit.js",
	)
}

func editJs() (*asset, error) {
	bytes, err := editJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "edit.js", size: 2964, mode: os.FileMode(420), modTime: time.Unix(1502237810, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\xdb\x6e\xe3\x36\x13\xbe\xf7\x53\xcc\x8f\x3f\x58\x52\x58\x85\x76\xb6\x77\x31\xdc\x62\xd7\x49\xd1\xc3\xb6\x29\x92\xf4\xa2\x28\x7a\xc1\x95\x46\x32\xd7\x14\x29\xf0\x10\xc7\xcd\xfa\x0d\xfa\x00\x7d\xbe\x3e\x49\xc1\x83\x2c\xc7\x4d\xd2\x02\xbd\x69\x2e\x62\x71\x38\x9c\xf9\xe6\xf4\x91\xb4\xf1\xaa\x72\x42\x2b\x5a\xc0\xc3\x64\x72\xc7\x0d\x9c\x34\xa6\x83\x05\x9c\x50\xd2\x68\xd3\x91\xa2\x9c\x00\x00\x9c\x54\x5d\x9f\xa4\xff\xaf\xba\x7e\x94\x4a\x3b\x48\xa5\xdd\x4b\xbd\x91\x59\xea\x8d\x1c\xa4\x95\xee\xb7\xdf\xe2\x16\x16\xa0\xf8\x9d\x68\xb9\xd3\x86\x79\x8b\xe6\x6d\x8b\xca\x31\xa1\x6a\xbc\xbf\x6a\x28\xf9\x8e\x57\x42\x39\x6d\x57\xa4\x80\xcf\x17\x30\x83\x2f\x80\xfc\xf1\xdb\xef\xa7\x4b\x02\xe7\x40\x96\xce\xc8\xd3\x25\x49\x16\x25\xb7\xee\x47\x23\xe7\x19\x37\xc2\x02\xf6\xe1\x28\xde\x61\x08\x09\xc0\xa0\xf3\x46\xc1\x09\xad\x75\xe5\xbb\xe0\xab\x32\xc8\x1d\x5e\x4a\x0c\xab\xa4\x59\xcc\x27\xbb\x6c\xc7\xa0\x15\xbf\x3e\xb2\x95\xec\xa4\xbd\xca\x85\xc8\x1a\xd3\xb1\x16\x1d\x9d\x15\xe1\xe7\x9d\xf6\xaa\x16\xaa\x5d\x4a\x81\xca\x5d\x63\xe5\x68\x31\x9f\x40\x52\xab\xac\xa5\xa4\xe3\xa6\x15\xea\xd4\xe9\x9e\x94\xb0\x11\xaa\xd6\x1b\x26\x94\x42\xf3\x15\x8a\x76\xe5\xa6\x9f\xc1\x69\xb4\xcd\x56\x69\xfd\xe6\x00\x4f\xc0\xf7\xa5\xd1\xdd\x21\x22\x6f\xc4\x08\xaa\xe7\xc6\x85\x2a\x78\x23\x98\xf5\x1f\xac\x33\x42\xb5\xf4\xac\x60\xb6\x97\xc2\x51\x32\x25\x11\x4d\xce\x43\xd4\xfe\xf9\xec\x97\xd1\x81\xd4\xbc\x7e\x3a\xdc\xe0\x3a\xd6\x2b\x21\xa0\x52\x57\x3c\xa8\xb0\x9e\xbb\x55\xcc\x5b\x0c\x93\xf1\x8f\xfc\x9e\x3e\xc4\x9a\x78\x23\xcf\x81\x4c\x79\x2f\xa6\xde\xc8\x29\x81\xd7\xf1\x78\x2a\x58\xcd\x1d\xbf\xdd\xf6\x78\x0e\xe4\xa3\xd5\x8a\x4c\x00\x76\x05\xe3\x72\xc3\xb7\x76\xec\xc4\xa0\x95\x30\x00\x88\x06\xe8\xff\x82\x80\xe9\xf5\x20\x03\xb0\x2b\xbd\xb9\x34\x46\x9b\xa8\xcb\x30\x7c\x46\x28\xe1\x2f\xc5\x99\x56\xbb\x49\xfc\x89\xa5\xd3\xde\x85\x60\xe2\x89\xb8\x28\xf3\x81\x08\x1a\x16\x59\xe3\xd5\xab\xf4\xc1\x82\xf0\xd3\x27\x20\x64\xbe\x6f\x6a\x76\xc7\x25\xf5\x46\x16\xac\xd1\x95\xb7\x34\xfb\xf4\x46\x5e\x88\x7a\xb9\xe2\xaa\xc5\x24\xdb\x85\xfa\xa5\xec\x06\xac\xef\x85\x5a\x3f\xdd\x9c\x31\xff\x71\x73\x9f\x5b\x6d\x44\x2b\x14\xbc\x06\xb2\xcf\xde\x3c\x84\x11\x26\x90\x39\xbc\x77\x94\x90\x22\xfa\x65\x06\x3b\x7d\x87\x4b\xc9\x43\x8f\x35\xbe\x5a\x0f\x1b\xbc\xae\xb3\x54\x0a\xb5\x0e\xf5\xcf\xbe\x78\x9c\xcb\x67\x26\x81\x70\x52\x0c\x06\x9c\x33\x94\xac\x0c\x36\xa4\x0c\x00\xb3\x38\xba\x1f\x97\xbc\xef\x51\xd5\xb7\x9a\x06\x6c\xa3\x93\xd5\x8b\x4e\x6c\xcf\xd5\xe8\x67\x0f\x74\xa5\x1c\x39\xf4\x92\xe9\xe2\x79\x4f\x31\x1f\x71\xba\x9c\xe1\xca\x46\xb2\x2a\x81\xd8\x8a\x4b\xfc\x89\x9e\x15\x39\xea\x16\xdd\x0d\x4a\xcc\xad\xcd\x2c\xba\x77\xdc\xe2\x5b\x55\x5f\xde\xbb\x80\x87\xe7\x41\x2e\x61\x56\xc2\xb8\x38\x3b\x98\xc0\x7d\xbb\x1d\xd6\xb0\x43\x6b\x79\x9b\xcb\xf8\x77\xb5\x49\x55\x38\x0e\x39\x55\x2c\x05\x83\x43\x62\xb2\x99\xcb\xeb\xeb\xab\xeb\x73\x08\x1d\x30\x78\xfa\x57\x99\x18\x62\x59\x89\x1a\x2f\x0c\xdf\xa0\xf9\xeb\xc8\xbf\x6c\x67\xf6\xc8\xce\x61\xcf\x3f\x4d\x1e\xf9\x12\x18\xc6\xa6\x60\xce\x88\x2e\x8d\x47\x98\xea\xb8\xbd\x18\x38\x7c\x98\xed\x71\x78\xe3\xe8\xe6\xdd\x48\x6e\x32\xc6\x3a\xe2\x7f\x64\x6a\x38\x1f\x2e\x24\xd6\xf0\x1a\xbf\x56\xf4\xcd\x6c\x96\x86\x11\x50\x5a\x3c\x56\xb8\xf2\x6e\xd4\x88\x61\x45\xba\xd6\x8a\x12\xeb\x3f\x74\xc2\x91\x72\x8c\x2a\x97\x19\x59\x6f\xf0\x0e\x95\xbb\xc0\x86\x7b\x99\x69\xfe\x9f\x31\xe5\xc0\x36\xcf\xa6\xe5\x98\x49\x5d\x22\xca\x1f\xae\x6e\x6e\xf3\x7d\x17\x8e\xbe\x4c\xae\x70\x0e\xdf\xdc\x5c\x7d\xcf\xd2\x1d\x20\x9a\x2d\x7d\xc8\xa7\xc2\xff\x5d\xf1\x98\x86\xe1\xbf\xc1\xc3\xf3\xd1\x43\x5c\x8f\xf6\x8f\x8b\xfd\x82\xd9\x03\x0a\x1f\x99\x7b\x64\xf8\x5c\x9f\xb4\x1f\x17\x07\xd4\x7e\xd4\x43\xc1\xaf\x75\xda\x6c\x99\xc1\x5e\xf2\x0a\x6f\x1c\x77\x48\x1f\x76\x25\x28\x2f\x65\x09\x64\x8a\xb5\x70\xfb\xf4\xef\xa1\x0d\x54\x4f\x0f\xa4\xbb\xe1\x36\x08\x15\x8e\x55\x0f\x1d\xb6\xc6\x6d\xad\x37\x8a\x94\x8f\xe6\x28\x4f\x78\x50\xe8\xb9\x75\xf8\xfc\x76\x15\x05\xc7\xfb\xc1\x43\x68\xef\xa8\x21\x45\xb5\x7e\xa2\x85\xf7\x8d\x47\xc8\xf8\x46\x49\x0d\x9f\x72\x7c\x7c\x99\x45\xe0\xf9\xc9\xc2\xeb\xfa\x32\xb4\xff\x7b\x61\x1d\x2a\x34\x94\xa4\xc7\x12\x29\xf3\xab\xa9\x98\x4f\xd2\x47\x38\x79\x6c\x29\xbc\x33\x62\xa3\xef\x8a\xf0\xf3\x67\x00\x00\x00\xff\xff\x02\x84\xe0\x0d\x7b\x0a\x00\x00"

func indexJsBytes() ([]byte, error) {
	return bindataRead(
		_indexJs,
		"index.js",
	)
}

func indexJs() (*asset, error) {
	bytes, err := indexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.js", size: 2683, mode: os.FileMode(420), modTime: time.Unix(1502237810, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _linksCss = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\x41\x6e\xe3\x30\x0c\x45\xf7\x73\x0a\x03\x59\x64\x33\x0a\x94\x38\x28\x5a\xfa\x16\xbd\x01\x63\x49\x36\x11\x46\x14\x24\xba\xb6\x2b\xf4\xee\x45\x90\xb8\x40\x16\x5c\x10\x7c\x7c\xf8\xff\x22\x6e\xad\x17\xec\xaf\x43\x96\x29\x3a\xd8\x85\x10\xba\x20\x51\x4d\xc0\x1b\xf1\x0a\xfb\x4f\x64\x3f\xe3\xba\xff\xdf\x14\x8c\xc5\x14\x9f\xe9\x49\x14\xfa\xf6\x70\x3e\xa5\xe5\xb1\xce\x9e\x86\x51\xa1\xb5\xf6\xe7\xc0\x14\xaf\xa5\xce\xe4\x74\x84\x77\x6b\xd3\xd2\xdd\x30\x0f\x14\xc1\x36\x38\xa9\x3c\x81\x06\x6b\x2f\x2c\x19\x76\xf6\x23\x74\xea\x17\x35\xce\xf7\x92\x51\x49\x22\x44\x89\xfe\x0f\x84\x51\xbe\x7c\xae\x92\xb0\x27\x5d\xc1\x1e\xde\xb6\xd3\x21\x4c\xcc\x66\xca\xbc\xb9\x9c\x73\x0f\x57\x19\xd1\xc9\x0c\xc7\xb4\x34\xf7\xb1\xcd\xbd\xdc\xf6\x36\x1e\x37\xbe\x6d\xdb\x67\x3a\x73\x11\x55\xb9\xc1\xd9\xa6\x65\xe3\x26\xae\x09\x9d\xa3\x38\x80\xed\x98\x8a\x9a\xa2\x2b\x7b\xa3\x6b\xf2\x2f\x11\x27\x6e\x98\xea\xab\xe8\x74\x17\xfd\xfb\x0d\x00\x00\xff\xff\xb8\xc8\x00\x8a\x64\x01\x00\x00"

func linksCssBytes() ([]byte, error) {
	return bindataRead(
		_linksCss,
		"links.css",
	)
}

func linksCss() (*asset, error) {
	bytes, err := linksCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "links.css", size: 356, mode: os.FileMode(420), modTime: time.Unix(1502237811, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _linksHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x4d\x8f\x94\x40\x10\xbd\xef\xaf\x28\x3b\x7b\x5c\x28\x36\x7a\x30\xd8\x8d\x31\xab\xf1\xb2\x89\x66\x33\x1e\x3c\xf6\x40\x01\x1d\x1a\x7a\xa4\x0b\x94\x10\xfe\xbb\x69\x60\xe2\xcc\x98\x98\x58\x97\x4e\xd5\xab\x8f\xf7\x5e\x5a\xbe\xfa\xf8\xe5\xe9\xf0\xfd\xeb\x27\xa8\xb9\xb5\xd9\x9d\x3c\x3f\xa4\x8b\xec\x0e\x00\x40\xb2\x61\x4b\xd9\x67\x07\x69\x0a\x1f\x72\x36\x23\xc1\xb3\xe9\x1a\x2f\x71\x43\xb6\xae\x96\x58\x43\xcd\x7c\x8a\xe8\xc7\x60\x46\x25\x9e\x5c\xc7\xd4\x71\x74\x98\x4e\x24\x20\xdf\x32\x25\x98\x7e\x31\x86\x23\xef\x20\xaf\x75\xef\x89\xd5\xc0\x65\xf4\x56\xe0\xbe\xc8\x9a\xae\x81\xba\xa7\x52\x09\xf4\x18\x32\x1f\xe7\xde\x8b\x15\x0d\xd1\x93\x55\xc2\xf3\x64\xc9\xd7\x44\x2c\xfe\x9e\x0b\x34\x52\xc4\xd2\x75\xec\xe3\xca\xb9\xca\x92\x3e\x19\x1f\xe7\xae\xc5\xdc\xfb\xf7\xa5\x6e\x8d\x9d\xd4\x8b\xb6\xf4\x53\x4f\xe9\x9b\x24\x79\x78\x9d\x24\xff\xba\x20\x71\x33\x44\x1e\x5d\x31\xed\x07\x0b\x33\x42\x6e\xb5\xf7\x4a\xac\x2c\x77\x22\x2b\x56\x3f\x66\xbb\x55\x76\xb3\xaa\x7e\xbc\x40\x07\xfb\x27\x09\x31\xcf\xd0\xeb\xae\x22\xb8\x6f\x68\x7a\x80\xfb\xde\x0d\x4c\x90\x2a\x88\x61\x59\xae\x3a\xa5\x35\xd7\xa3\x6b\x51\xef\xba\xe7\x79\x9f\x8d\xbf\xbd\x3c\xc3\xb2\x88\xac\x72\x18\x8a\x0d\x4d\xb0\x2c\x12\x75\x26\x8f\x3d\xe0\xff\xac\x38\x4b\x2c\x07\x6b\xa3\xa1\xb7\x22\xbb\x6d\x09\x6b\xaf\x49\xe2\x2d\xcb\x79\x06\xea\x8a\x4b\x31\x12\xcf\x26\x48\x2c\xcc\x18\x1c\xde\xac\x95\xb8\xfd\xc0\xdf\x01\x00\x00\xff\xff\xc6\x38\x2f\x69\x99\x02\x00\x00"

func linksHtmlBytes() ([]byte, error) {
	return bindataRead(
		_linksHtml,
		"links.html",
	)
}

func linksHtml() (*asset, error) {
	bytes, err := linksHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "links.html", size: 665, mode: os.FileMode(420), modTime: time.Unix(1502237810, 0)}
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
	"close.svg": closeSvg,
	"edit.css": editCss,
	"edit.html": editHtml,
	"edit.js": editJs,
	"index.js": indexJs,
	"links.css": linksCss,
	"links.html": linksHtml,
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
	"close.svg": &bintree{closeSvg, map[string]*bintree{}},
	"edit.css": &bintree{editCss, map[string]*bintree{}},
	"edit.html": &bintree{editHtml, map[string]*bintree{}},
	"edit.js": &bintree{editJs, map[string]*bintree{}},
	"index.js": &bintree{indexJs, map[string]*bintree{}},
	"links.css": &bintree{linksCss, map[string]*bintree{}},
	"links.html": &bintree{linksHtml, map[string]*bintree{}},
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

