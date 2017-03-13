// Code generated by go-bindata.
// sources:
// assets/s/main.css
// assets/templates/.transactions.html.swp
// assets/templates/familypage.html
// assets/templates/forms.html
// assets/templates/index.html
// assets/templates/layout.html
// assets/templates/newaccount.html
// assets/templates/newuser.html
// assets/templates/transactions.html
// assets/templates/userhome.html
// assets/templates/view.html
// assets/templates/viewac.html
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

var _assetsSMainCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8b\xb1\xae\xc2\x30\x0c\x45\xf7\x7c\x85\xa5\xb7\x3e\xb1\xa0\x0c\x4d\xbf\x26\xc6\x75\x14\xd1\xd8\x91\x49\x24\x10\xe2\xdf\x29\xa1\x2c\x70\xc7\x7b\xce\x41\xa5\xdb\xdd\xc1\x36\x8c\xa7\x73\x32\xed\x42\xe1\x6f\x9a\x98\x11\x67\xf7\x70\xee\x50\x16\xe9\xf0\xab\x78\xcf\x1c\xe3\x50\x58\xad\x7c\x0c\x35\x5a\x2c\xf8\x7a\x85\x8b\xae\x99\x00\xd7\xad\x99\x07\x2b\xd1\x52\x96\x17\x1b\x55\x96\xda\x1b\xfc\x03\xf6\xd6\x54\xf6\xfe\xcb\x69\xb4\xff\x35\x12\x65\x49\xe1\xf8\x06\xcf\x00\x00\x00\xff\xff\x54\x41\xbd\xbe\xb6\x00\x00\x00")

func assetsSMainCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsSMainCss,
		"assets/s/main.css",
	)
}

func assetsSMainCss() (*asset, error) {
	bytes, err := assetsSMainCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/s/main.css", size: 182, mode: os.FileMode(420), modTime: time.Unix(1489057121, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesTransactionsHtmlSwp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\xda\x4f\x8b\x13\x31\x18\x06\xf0\xec\x82\xde\xb4\x8b\xff\x0e\x1e\x24\xd6\xa2\x97\xdd\xc9\x5a\x05\x45\xc6\x82\x20\x2b\x7b\x10\xf6\xa0\xe2\x35\x9d\x79\xed\xd4\x6d\x92\x21\x79\x5b\x29\x43\xf5\xe0\xd7\xf0\x22\xf8\x59\x3c\x7a\x10\xf4\x53\xf8\x05\xbc\xf9\x76\x6d\xa1\xae\x42\x6f\x2e\xe2\xf3\x83\x87\x4c\xf2\xa6\xd3\x37\xc7\x81\xf4\x77\x9f\xed\x3f\xd6\x77\xb2\xdb\x4a\x6c\x29\xf5\xf9\xcb\xc7\xe7\xe9\xd4\xf6\xe6\xf7\xcb\x4a\x39\xcb\x5c\xd1\x2b\xb5\xce\x23\x8a\x3c\xed\xae\xdd\xa6\x5e\x2f\x5e\x68\x52\x11\x87\x35\x27\x33\x08\x23\xeb\x07\x26\xc5\xc2\x0c\x86\x5c\x8d\xfb\x59\x11\x9c\x29\x42\x49\xb1\x08\x7e\x12\xa6\xa6\x0e\xc5\x21\xb1\x0b\x9e\xa6\xc6\xa6\x44\xf2\x23\x26\x57\x8f\x2c\x93\x3c\x45\xeb\x93\x2d\x78\x18\x7c\xca\x2a\x76\xa3\xf5\x2d\x00\xfc\x9f\xc6\xfc\x62\xe7\xee\x19\x75\xab\x7b\x73\x77\x3e\xbd\xd6\xbe\xaa\xcf\x9f\x7b\x7a\xd2\x5d\x01\x00\x00\x00\x00\xc0\x5f\xc4\xf5\x86\x7a\x23\xe3\xe6\x62\xbe\xb5\x18\x37\x8e\x8d\x00\x00\x00\x00\x00\x00\x00\xf0\xef\xb2\xa5\x52\x1f\xce\xca\x43\x4b\x1d\x7d\xfc\x2f\xbf\xff\xbf\xc9\xfc\xab\xe4\x93\xe4\xbd\xe4\x9d\xe4\xad\xe4\xa5\x64\x5f\xb2\x23\xb9\x21\xb9\x2e\xb9\x22\xb9\x24\xb9\x28\xb9\xd0\xfa\xf9\xae\xd3\xad\x93\x3b\x13\x00\x00\x00\x00\x00\x00\xc0\x71\x4d\xb3\xbc\x4f\xae\xdb\x7d\x9b\xa8\x3d\x9b\xc9\x6a\xce\xb6\x3f\xa2\x9e\x54\xc9\x97\xf3\x95\xdc\x70\xec\x29\x2d\x72\x2e\x7b\x4d\xd3\x99\x64\x07\xe3\x58\x87\x44\xb3\x99\xd4\xca\xd5\xda\xd1\x6d\x75\x2d\x3b\x1e\xb8\x30\xf6\xfc\xfb\x06\x29\x3d\xa4\xf4\xe7\xc2\x5e\x0c\x6e\x59\xc8\xe7\xff\xd9\x34\xd1\xfa\x01\xe9\xce\xe1\x76\x67\xa2\xef\xdd\xd7\xd9\x9e\x75\xd9\x01\xc5\x61\x28\xb3\x27\x2b\x37\xe0\xb5\xf4\xb9\xec\x3b\xaf\xba\xbd\xd5\x5a\x6e\x64\xe1\x97\xb3\x56\x64\xcb\xf9\x59\x7f\x04\x00\x00\xff\xff\x54\x9b\x2b\xf3\x00\x30\x00\x00")

func assetsTemplatesTransactionsHtmlSwpBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesTransactionsHtmlSwp,
		"assets/templates/.transactions.html.swp",
	)
}

func assetsTemplatesTransactionsHtmlSwp() (*asset, error) {
	bytes, err := assetsTemplatesTransactionsHtmlSwpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/.transactions.html.swp", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1489425875, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesFamilypageHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x53\xc1\x6e\x9c\x30\x10\xbd\xe7\x2b\x46\x73\x4a\xa4\x06\x1a\xa2\x5c\x22\x40\xaa\x22\x45\xca\x21\xd5\x4a\x55\x3f\xc0\xe0\xd9\x60\xd5\x36\xd4\x36\xda\x5d\x59\xfc\x7b\x0d\x86\x65\xe1\x52\x2e\xe0\xe1\xbd\x37\xf3\x9e\x6d\xef\x1d\xa9\x4e\x32\x47\x80\x0d\x31\x8e\xc3\x70\x07\xe1\xc9\x9b\xa7\xf2\x9d\x29\x21\x2f\x70\x60\x5f\xf4\x0a\xde\x27\x61\x9d\xc4\xda\x4f\xa6\x68\x18\xf2\x34\x80\x22\xba\x97\xf1\xc3\x7b\xc3\xf4\x17\xc1\x84\xfd\x24\x55\x91\xb1\xb3\xe2\x84\x93\xa2\xbc\x2e\x22\x3e\xf9\x6d\xc9\xe8\x49\x0f\x1e\x1f\x43\xa1\xfa\x45\x92\x6a\x07\xc9\x81\x19\xd2\x0e\x30\xbe\x11\xf0\xad\x11\x72\x1c\x30\xaf\xcc\xaa\x92\xa7\x8b\xa6\xf7\xa4\xf9\x32\x7e\xba\x4c\x34\x82\xaf\x04\xef\xc5\x11\xee\xa7\xe1\x3e\xec\xac\x9f\xbc\x2b\x52\x0f\x0b\xef\xd8\x1a\x05\xac\x76\xa2\xd5\x05\x32\xce\xd5\xe4\x01\x41\x91\x6b\x5a\x5e\x60\xd7\x5a\x87\x25\xdc\xad\x06\x82\x62\x70\xba\x71\xd9\x3c\x43\x2d\x99\xb5\x05\x2a\xb2\x36\xc4\x87\x65\x30\x1a\x40\x30\x85\xf6\x5c\xde\xd0\xd7\x99\x23\x35\x2b\x7f\x70\x0e\x73\xf4\x31\xc1\x40\xc9\x56\xca\x98\x3d\xbc\x42\x2e\x74\xd7\x3b\x70\x97\x8e\x0a\x74\x74\x0e\x01\x8d\x29\x16\xd8\xcf\x79\xe2\x4a\xf9\xb0\x10\xcd\xee\x68\x75\x43\xf5\x9f\xaa\x3d\x2f\xd4\x2e\x26\x5d\x6e\x02\x3e\x04\x23\xa7\xd6\xf0\x7d\xcf\x6e\xae\x23\x74\xcc\xb9\xd0\xb3\xc0\xc4\xbf\x7c\xcb\xbe\x0f\x08\x86\xfe\xf6\xc2\x10\x5f\x74\x4f\xfc\x69\xa7\x7a\xff\x29\xb4\x50\xbd\x82\x17\x90\x34\xd2\xed\xc3\xe6\xff\x5b\xab\x8f\x22\x6c\xc5\xff\xbb\x5f\x5b\x64\xbb\x16\x1b\xbc\xed\x2b\x25\xdc\x9c\x49\x9e\x8e\xdb\xbc\x3d\x35\xb7\x17\xa1\x62\x96\xc6\x8b\xf0\x2f\x00\x00\xff\xff\x93\x9f\x2a\x90\x1d\x03\x00\x00")

func assetsTemplatesFamilypageHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesFamilypageHtml,
		"assets/templates/familypage.html",
	)
}

func assetsTemplatesFamilypageHtml() (*asset, error) {
	bytes, err := assetsTemplatesFamilypageHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/familypage.html", size: 797, mode: os.FileMode(420), modTime: time.Unix(1489400313, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesFormsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x55\x4d\x6e\xdb\x3c\x10\xdd\xfb\x14\x04\x91\xc5\xf7\x01\x81\xe2\xa4\x05\x0a\x18\xb6\x81\xc0\x89\x57\x4d\x62\xb4\x29\xba\x34\x18\x73\x1c\x0b\x11\x49\x95\xa4\xec\x0a\x82\x4e\xd3\x9b\xf4\x64\x1d\x52\x94\x2c\x39\x8e\xed\x02\x59\x55\x8b\x84\x3f\x33\x8f\xf3\xde\xfc\xb8\x28\x38\x2c\x63\x09\x84\x2e\xb5\x98\x33\xce\xe7\x6c\xb1\x50\x99\xb4\xb4\x2c\x7b\x04\xbf\xe1\x52\x69\x41\xd8\xc2\xc6\x4a\x8e\x28\x1a\xd4\xf7\x44\x80\x5d\x29\x3e\xa2\xa9\x32\x96\x8e\x2b\xe3\xd5\x87\xf1\x35\xe7\xe4\xba\xb2\x19\x5e\xe0\xbe\xba\x88\x65\x9a\x59\x62\xf3\x14\x46\xd4\xc2\x4f\x74\x97\x4c\xe0\x3a\x33\xa0\xdd\x8a\x12\x0d\x8c\x2b\x99\xe4\x64\xcd\x92\x0c\x6f\x8a\x22\x9a\x0a\x10\x65\x49\xc9\x2a\xe6\x1c\xe4\x11\xa4\x10\x57\x05\x96\x32\x6b\x11\x78\x44\xa3\xe2\xe3\xf9\x55\xbf\x74\xf0\x3f\xb2\x58\x03\xdf\x83\x62\xb2\x27\x11\x23\x8e\x43\x1b\x51\x8c\x9f\xd6\x31\xb8\x75\x70\xb8\x70\x3a\x8c\x7b\x45\x01\x92\xa3\x34\xbd\x96\x70\x22\x9d\x8b\x3c\x3c\x6f\x1a\xdd\x0c\x24\xb0\xb0\x21\xb8\xa5\x56\x22\x20\xb9\xaf\x28\x34\x93\xcf\x40\xce\x5e\xc8\x39\x39\x5b\x93\xc1\x88\xfc\x17\x4d\x99\x88\x3e\xc7\xc6\x7e\xd7\xb1\x05\x54\x10\x8c\x21\x5e\x82\xff\x03\xa4\x87\x55\xa9\xcb\xc4\x56\xa4\xb3\x75\xf4\x2d\x68\x58\x96\x03\xbf\xbf\xf7\x6b\x3a\x3e\x70\x39\x70\x31\x08\x25\x21\xc7\xe7\xa3\x5b\x47\x69\x78\x51\x41\xb7\xa3\xac\xb8\x56\xfc\x2b\x3e\x5b\x05\xba\x02\xb0\x24\x39\xac\x80\x55\x7b\xf9\x9f\x57\xec\x3d\xf9\x19\xe8\x58\xf1\x28\x14\x8f\x79\x7f\xd6\x7f\x4b\xb1\x93\x65\x6c\x8f\x94\xe5\xfb\xdb\xc2\x5d\xec\xeb\x87\xba\x27\xee\xd8\x0b\x10\x46\x66\x2c\x17\xf0\xee\x6d\x41\x88\xc7\x9a\x62\x89\x0d\x90\x90\x05\x91\x26\xcc\xbe\x2e\x4c\x12\x85\xd8\x1f\xd5\x6b\xbb\x76\xfe\x9c\xe1\xf0\x49\x57\x21\x5e\x0b\x77\x88\xf5\xf2\xfb\x57\x27\x5a\x99\x89\x27\xd0\x94\x18\x0b\xe9\x88\xf6\xa3\xfe\x25\x2a\x10\xa3\x16\xfd\xa6\x21\xbd\x27\x1d\x37\x50\xb3\x4c\xa3\x36\x30\x38\xc0\x3b\xad\x4c\x5a\x4e\x47\xba\x15\x25\x25\xf7\x6a\xd3\x74\x6c\xbd\x3f\xde\xb5\x98\x4f\x63\x99\xe4\xb1\x7c\x7e\x73\xd6\x35\x06\x27\x24\xf7\x6b\xb0\x25\x0f\x9a\x83\xde\xe6\xf8\x1f\xcc\x8b\xf7\x42\xbe\xda\x92\x1b\x8c\x74\xd0\xf5\xe3\x78\x54\xfb\x19\x67\x44\xb7\xb5\xeb\xee\xbe\x4c\x27\xae\x4f\x9b\xe7\x1f\x57\x20\xc9\xed\x1a\x74\xbe\x03\x54\x53\xf1\xe1\x5f\xd6\x90\x1c\x12\xd7\x6d\x01\xf2\x53\x9d\xe9\xce\xb0\xf1\x36\x73\x87\xd2\xce\x55\x77\x88\x70\x96\x1b\x3a\xbe\xc1\xbf\xaf\xc7\xc2\x8e\x29\x0e\x4a\xbb\x42\xe3\x3b\xff\xbf\x6b\xde\x4c\x8d\x86\xcf\x29\x65\x3b\xc1\xb6\x76\x2a\x85\x07\xc2\xf6\xa4\xa2\x5d\xac\x52\x66\xfc\x90\xed\x96\x6b\x38\x7f\xeb\x67\x79\xb2\xf2\x13\x77\x86\x36\x1b\xa5\xf9\xb6\x3e\x1f\x12\xde\x9c\x92\x9d\x0c\xa4\xe1\xbc\xd6\x5e\x25\x3c\xdd\xf0\x56\xee\xee\x61\x43\xd2\x13\x9d\xd1\xf3\xb2\xe5\x3a\x51\x72\x19\x63\xfc\xc7\xbd\xae\x5a\x5e\x87\x34\xf5\x0c\xb7\x9a\x56\xdb\x71\x6f\x57\xcf\x3f\x01\x00\x00\xff\xff\x45\xda\x4b\x8d\xf3\x08\x00\x00")

func assetsTemplatesFormsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesFormsHtml,
		"assets/templates/forms.html",
	)
}

func assetsTemplatesFormsHtml() (*asset, error) {
	bytes, err := assetsTemplatesFormsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/forms.html", size: 2291, mode: os.FileMode(420), modTime: time.Unix(1489423765, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\xcd\xce\x9b\x30\x10\xbc\xf3\x14\x2b\x9f\x5a\xa9\xcd\x0f\x52\x2e\x15\x20\x55\x51\x7b\x0a\x55\x2e\x7d\x00\x03\x4b\xb1\x8a\x6d\x6a\x9b\xa6\x11\xe2\xdd\x6b\xc7\x90\x90\x84\x34\xcd\x97\x13\xbb\x9e\x19\xaf\x66\x36\xee\x3a\x83\xbc\xa9\xa9\x41\x20\x15\xd2\x82\xf4\x7d\x00\xc3\x2f\xaa\xd6\xc9\x5e\xe6\x3f\xd1\x40\x2a\x05\x1e\xa3\xa5\x6d\x04\xe7\xe3\xae\x63\x25\x2c\xa6\xf8\x2c\xe9\xba\x45\x8a\xba\xef\xa3\x65\x96\x4c\x80\x28\x0a\x8b\x0b\x2e\xdc\xa8\x94\x8a\x03\xcd\x0d\x93\x22\x26\xb5\xfc\xc1\x04\x01\x8e\xa6\x92\x45\x4c\x1a\xa9\x0d\xb9\xd0\xfd\x28\x61\xb2\x73\x28\x3b\x43\x78\x7d\xf4\x95\x72\x56\x1f\x3f\x41\xc4\x44\xd3\x1a\x30\xc7\x06\x63\x62\xf0\x8f\x21\x20\x28\xb7\xdf\xe5\x09\x40\xe0\x37\xad\x5b\x5b\xda\x11\x3d\xa5\xef\x49\x12\x65\xea\x5a\x2d\x45\x9e\xa1\x82\x6f\x96\x09\xff\xd0\x6c\x35\x2a\xf7\x35\x55\xfd\x3e\xf4\x66\x75\xf7\x54\xeb\x83\x54\xc5\xad\x68\x33\xf4\x47\xe1\xe6\x50\xcc\xb0\xaf\x28\xba\xcd\x38\xb3\x93\xb8\x79\x62\xb2\xf3\xde\x0d\x63\xf8\xea\x42\x8e\x96\xce\xe8\xe4\xa1\xf3\x02\x0f\xa3\x3d\x4f\xdc\xdf\x2a\x74\x4b\xe2\xad\x83\xcf\x79\x2e\x5b\x61\x1e\xa5\xf1\xd4\x3f\x7f\xe9\xc9\xc1\x99\x0c\x28\x13\xe0\xdc\xfc\xaf\x04\xee\xf9\x5f\x38\x65\xf5\x2d\x17\x5d\x73\x24\xfb\xe2\x2d\x29\x35\xd4\x18\x7b\x71\x4c\x16\xdd\xe6\x43\xb8\xea\x09\x28\xfc\xd5\x32\x85\x05\x18\x66\x6a\x8b\xdd\x7c\x0c\x57\x90\x57\x54\x59\x8f\x51\xe9\x49\xb0\xeb\x99\x1b\xdf\xa5\x4c\x30\xde\x72\xd8\x40\x8d\x4e\x5a\xbf\xbf\xc3\x6c\xa5\x28\x99\x0d\xed\x85\x1d\x0a\x5f\x5b\x22\x1f\xee\x79\x8b\x86\x72\xf2\x5f\x1f\xf7\x68\xfa\x5c\x64\x54\xa3\x7d\x2e\x46\x48\xf0\x37\x00\x00\xff\xff\xdf\x16\x24\x30\x4b\x04\x00\x00")

func assetsTemplatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesIndexHtml,
		"assets/templates/index.html",
	)
}

func assetsTemplatesIndexHtml() (*asset, error) {
	bytes, err := assetsTemplatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/index.html", size: 1099, mode: os.FileMode(420), modTime: time.Unix(1489069089, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesLayoutHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x51\x3d\x73\xb3\x30\x0c\xde\xf9\x15\x7e\xb5\xbf\xf5\xda\xc1\x78\xe9\xc7\xd4\x5e\x18\xb2\x74\x54\x40\x04\x5f\x8c\x9d\x43\xa2\x57\x8e\xcb\x7f\xaf\xc1\x6d\x43\x87\xc6\x8b\xbe\x9e\x47\x96\xf4\xcc\x73\x43\xad\x0b\xa4\xa0\x23\x6c\xe0\x72\x29\xcc\xbf\xc7\xdd\xc3\xfe\xad\x7a\x52\x9d\xf4\xde\x16\x26\x1b\x95\x9e\x59\x30\xd9\x5d\xc3\x9e\x04\x55\xdd\xe1\xc0\x24\x25\x8c\xd2\xfe\xbf\x87\x4d\xd9\xbb\x70\x52\x03\xf9\x12\x58\x26\x4f\xdc\x11\x09\x28\x99\xce\x54\x82\xd0\x87\xe8\x9a\x19\x54\x37\x50\x5b\x82\x66\xdd\xa3\x0b\x77\x4b\x6a\xd3\x42\x9c\x78\xb2\x55\xac\x4f\x24\xea\x35\x06\x9a\x8c\xce\xb9\x3c\x90\xbe\x4e\x64\x0e\xb1\x99\x36\xd4\xc6\xbd\xab\xda\x23\x73\x09\x3d\x85\x71\xd3\x75\x2d\xe3\xf7\xc7\x2d\xf6\xce\x4f\x60\x9f\x57\x6b\x34\xfe\x05\x3c\xd3\xc0\x31\xa0\x07\x5b\x7d\x79\x37\xc0\x32\x60\x60\xac\xc5\xc5\x90\x16\xda\x6f\xa2\x1b\x24\x1f\x8f\x71\x14\xb0\x2f\xf1\xa8\x76\xa3\xfc\x42\x1a\x9d\x16\xb2\xc5\x3c\x53\x68\x92\x4a\xe9\x5d\xa5\x3b\x20\xd3\x22\x5d\xc6\xe5\x3b\xa4\xd3\xac\xba\xfd\x10\x3e\x03\x00\x00\xff\xff\x7a\x8a\x17\xb6\xea\x01\x00\x00")

func assetsTemplatesLayoutHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesLayoutHtml,
		"assets/templates/layout.html",
	)
}

func assetsTemplatesLayoutHtml() (*asset, error) {
	bytes, err := assetsTemplatesLayoutHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/layout.html", size: 490, mode: os.FileMode(420), modTime: time.Unix(1488553008, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesNewaccountHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xe2\xb2\x81\xd0\x0a\x40\x60\x93\x91\x9a\x98\x02\x61\x82\xb9\x25\x99\x25\x39\xa9\x76\x1e\xa9\x39\x39\xf9\x36\xfa\x10\x0e\x42\x32\x37\xb5\x24\x51\x21\x39\x23\xb1\xa8\x38\xb5\xc4\x56\xa9\xb4\x24\x4d\xd7\x42\x09\x6a\x8c\x3e\xc2\x1c\x9b\xa4\xfc\x94\x4a\x24\x5d\x19\x86\x76\x7e\xa9\xe5\x0a\xa1\xc5\xa9\x45\x40\x65\x86\x08\x99\xea\xea\x82\xa2\xcc\xbc\x92\x34\x05\x25\xd5\x62\x25\x05\xbd\xda\x5a\xa8\x51\x10\xfd\x40\xb5\x10\xd7\x02\x02\x00\x00\xff\xff\xd6\x4d\xae\x17\xbf\x00\x00\x00")

func assetsTemplatesNewaccountHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesNewaccountHtml,
		"assets/templates/newaccount.html",
	)
}

func assetsTemplatesNewaccountHtml() (*asset, error) {
	bytes, err := assetsTemplatesNewaccountHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/newaccount.html", size: 191, mode: os.FileMode(420), modTime: time.Unix(1488373318, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesNewuserHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xe2\xb2\x81\xd0\x0a\x40\x60\x93\x91\x9a\x98\x02\x61\x82\xb9\x25\x99\x25\x39\xa9\x76\x1e\xa9\x39\x39\xf9\x36\xfa\x10\x0e\x42\x32\x37\xb5\x24\x51\x21\x39\x23\xb1\xa8\x38\xb5\xc4\x56\xa9\xb4\x24\x4d\xd7\x42\x09\x6a\x8c\x3e\xc2\x1c\x9b\xa4\xfc\x94\x4a\x24\x5d\x19\x86\x76\x7e\xa9\xe5\x0a\xa1\xc5\xa9\x45\x40\x65\x86\x08\x99\xea\x6a\xbd\x50\xbf\xc4\xdc\xd4\xda\x5a\x9b\xa4\x22\x24\x0d\x89\x0a\x19\x45\xa9\x69\xb6\x4a\xfa\x40\x6d\x8e\xc9\xc9\xf9\xa5\x79\x25\x4a\x60\x23\xa0\x1c\x1b\xfd\x44\x98\xad\x10\xab\x80\xc6\x42\x3c\x06\x08\x00\x00\xff\xff\xa6\xb7\x94\xc4\xea\x00\x00\x00")

func assetsTemplatesNewuserHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesNewuserHtml,
		"assets/templates/newuser.html",
	)
}

func assetsTemplatesNewuserHtml() (*asset, error) {
	bytes, err := assetsTemplatesNewuserHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/newuser.html", size: 234, mode: os.FileMode(420), modTime: time.Unix(1488373372, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesTransactionsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8f\xcd\xaa\xc2\x30\x10\x85\xf7\x79\x8a\x50\xba\xbc\xa4\x70\x97\x97\xdc\x80\x20\x5d\x77\xe1\x0b\x4c\xcd\x60\x8b\x4d\x52\x92\x69\x41\x42\xde\xdd\xa9\x22\x2a\x9a\x45\x60\xce\x37\x3f\xe7\xe4\x4c\xe8\xe6\x09\x08\x65\x35\x20\xd8\xaa\x14\xa1\x87\x5f\x73\x88\xe0\x13\x1c\x69\x0c\x3e\xe9\x86\x05\xa1\x09\xfa\x09\x8d\xc8\x99\xd1\x09\x65\x7d\xfe\xa9\x57\xf9\xf7\x2f\x55\x0b\x4e\x75\x18\xc7\x60\xd5\xeb\x98\xdc\x56\x51\x34\x42\xf2\xd3\x64\x4d\xce\xf5\xaa\xda\x18\x5c\x29\xba\xe1\xfa\x1d\xec\x31\xd1\x27\x70\xc1\xe3\x45\x32\xde\xb9\xb0\xf8\x2f\x0d\x8c\xba\x25\xce\x21\xe1\x83\xf1\xcf\x47\xd9\x27\x7a\x7b\xb3\x70\xf7\xbd\x29\xcf\xac\x3d\x24\xdc\xb2\x5e\x03\x00\x00\xff\xff\x71\x3a\xcc\xd7\x00\x01\x00\x00")

func assetsTemplatesTransactionsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesTransactionsHtml,
		"assets/templates/transactions.html",
	)
}

func assetsTemplatesTransactionsHtml() (*asset, error) {
	bytes, err := assetsTemplatesTransactionsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/transactions.html", size: 256, mode: os.FileMode(420), modTime: time.Unix(1489425875, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesUserhomeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x92\x4f\x6f\x82\x40\x10\xc5\xef\x7e\x8a\xc9\xc6\x78\x30\x0d\x1b\xeb\xcd\x2e\x34\x3d\xb4\x87\x26\x6d\x4c\x4c\xcf\x66\x84\x51\x48\xd9\x85\xb2\xa8\x35\x1b\xbe\x7b\x67\x17\xa9\xa6\x7f\xb8\x90\x9d\xdf\x9b\x37\xf3\x58\x9c\x6b\x49\xd7\x25\xb6\x04\x22\x27\xcc\x44\xd7\x8d\x54\x7e\x9b\xbc\x59\x6a\x60\x89\x3b\x82\x05\x38\x17\x3d\x69\xd2\x5d\xa7\x24\x13\xc6\xf3\xe4\x21\x4d\xab\xbd\x69\x2d\x57\xe6\xc9\xc8\xb9\x62\x0b\xd1\x0b\xd9\xd0\x3c\x87\xb4\x44\x6b\x63\xa1\xc9\x5a\x76\x10\x09\x1b\x04\x38\xa8\xc9\x64\xac\x74\x4e\x4e\x57\x79\x75\x84\xc1\x6c\x2a\xb9\x3a\x52\xfb\xd2\x6b\xc6\x7e\x24\x2c\x62\x38\xcf\xe6\x52\x83\x86\xf7\x19\xbf\xdf\xc0\xf8\xd0\x13\xd4\xd1\x92\x9a\xa2\xca\xa2\xc1\x83\x85\xc0\x4f\xd8\x88\x3e\x58\x18\xf9\x24\x06\x35\x37\x9e\x8d\x3c\x57\x65\x91\x28\x84\xbc\xa1\x6d\x2c\xe4\xa1\xa0\xe3\xfd\xde\x8b\x62\x1e\x7c\x69\xe9\xba\x09\xa6\x7d\xe9\x35\x1c\x7d\x94\xef\x83\x92\x98\x4c\xcc\xc6\xd6\x77\xce\xe9\xca\xd0\xc9\x4f\x7b\xf4\xd1\x94\x64\xfb\xf3\x1e\x43\xd6\xfe\xad\xa4\x4f\xe7\xe3\x7d\xce\x66\x21\xc3\x8a\xda\xe7\x6a\x03\xc2\xb6\x68\x32\x71\x95\xf6\x72\x2f\x01\x15\x66\xc7\xd4\x7f\xa0\x6b\xb4\x6d\xf4\x1a\xb3\x6c\x8d\x7d\xfc\x7f\x14\x35\x9e\x7a\xf2\x13\x5c\x39\xc3\x1f\x38\xcd\x6b\xbe\xc8\xdf\xad\x1b\xb4\xe4\x7f\x94\xaf\x00\x00\x00\xff\xff\x71\x4b\xac\x40\x3d\x02\x00\x00")

func assetsTemplatesUserhomeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesUserhomeHtml,
		"assets/templates/userhome.html",
	)
}

func assetsTemplatesUserhomeHtml() (*asset, error) {
	bytes, err := assetsTemplatesUserhomeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/userhome.html", size: 573, mode: os.FileMode(420), modTime: time.Unix(1489423578, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesViewHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8d\xcf\x8a\xc3\x20\x10\xc6\xef\x3e\xc5\x10\x72\x5c\x0c\xec\x71\x71\x85\x85\x25\x87\xbd\xec\xa1\x4f\x60\x70\x9a\xa6\x8d\xa6\xa8\x09\x2d\x92\x77\xef\x4c\xfe\x41\x05\x05\x7f\xf3\xcd\xef\xcb\xd9\xe2\xb9\xf3\x08\x45\x4c\xc6\xdb\xce\xb7\xc5\x3c\x8b\x9c\xcb\x6b\x84\xaf\x6f\x90\x7f\x43\xb3\x8d\x76\x9e\x98\xdf\x7b\x7c\x00\x67\xf8\xca\xda\xa1\x83\x75\xdc\x77\x71\x09\xec\xb6\xe6\x39\x46\x0c\x14\x31\x4e\x9e\x36\xb6\x2e\x50\x5e\x5d\x3e\xf5\x01\xff\x83\xc5\x10\x55\x45\x4c\xa8\x64\x9a\x1e\x35\x09\x83\xf1\x2d\x42\x79\xfb\x80\x72\x62\xef\x52\xc0\xab\x29\x68\x01\x74\x54\xb2\x9a\x7a\x27\x59\x87\x81\x9c\xaa\xa2\xff\xc1\x7e\x91\xc3\xef\xec\xc7\x0d\xa3\xdf\xa9\xa0\x37\x70\x0f\x7a\xcb\xd6\xea\x28\x5e\xc1\x2b\x00\x00\xff\xff\x90\xda\x0c\x23\x1f\x01\x00\x00")

func assetsTemplatesViewHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesViewHtml,
		"assets/templates/view.html",
	)
}

func assetsTemplatesViewHtml() (*asset, error) {
	bytes, err := assetsTemplatesViewHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/view.html", size: 287, mode: os.FileMode(420), modTime: time.Unix(1489250289, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesViewacHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x52\xd1\x6a\x83\x30\x14\x7d\xef\x57\x84\xe2\xc3\x06\xc3\x42\x1f\x47\x26\x14\x46\x61\x63\x6c\x63\x74\x1f\x10\xf5\xb6\xca\x4c\xe2\x92\x28\x14\xf1\xdf\x77\x93\xe8\x8c\x3a\xd6\x97\x5e\xcf\xb9\xf7\x9c\x93\xe4\x76\x9d\x01\x5e\x57\xcc\x00\xd9\x16\xc0\xf2\x6d\xdf\x6f\x08\xfe\xba\x2e\x32\xb2\x26\xf7\x0f\x24\x9e\x10\x96\x39\xe0\x59\xa6\x64\xcb\x32\xdb\x4a\x8b\x7d\xe2\x88\xbe\xa7\x3b\xac\x37\xd4\xb0\xb4\x82\xc4\x4d\x50\xa3\x7c\xe1\x3f\x8a\xe4\xa3\x11\xa2\x14\x17\xba\xc3\x7a\x46\x3c\xbd\xae\xb1\xb7\xcf\xd3\x1f\xa0\x29\x40\xad\xe1\xf7\x46\xd5\x52\xc3\x44\x60\x35\x78\x63\xba\x97\x52\x1b\x17\xfc\xc8\x78\x7c\xc8\xb2\x86\x37\xf6\xc0\x27\xc5\x84\x66\x99\x29\xa5\xd0\xc4\x1d\x61\x18\x40\xfc\x02\x24\xfa\x22\x77\x51\x6b\xe7\x9c\xc0\xc0\x2e\xce\x94\x27\x34\xc5\x0b\xe0\x52\xc0\x95\x44\x6d\x7c\x38\x1b\x50\xf6\x2e\xd2\x04\x23\xe4\x53\x2b\xde\xf3\xb5\x86\xc0\x26\x04\xdb\xf8\xa8\x24\x9f\x11\xe5\x99\xdc\xc0\xf7\xde\x0e\x8c\xfc\x6d\xd0\xe0\x9c\xad\x81\x2d\x42\x7f\x2e\x1b\x61\x6c\x80\x5f\x0e\xd1\x47\xd0\x23\x16\x58\x40\xa5\x61\x21\xf9\x9f\xd2\x5c\xd2\x07\x5e\x49\x8a\x7c\xa5\x88\xcd\xc3\xf3\x2c\xfb\x3d\x9f\xdb\xdd\xb3\x21\xf1\x7f\xec\x58\xbd\xa1\x17\x46\xc4\x6f\xd7\x26\x5c\xdb\x94\x69\xb0\xbb\xf8\x13\x00\x00\xff\xff\xc8\x79\xf6\xac\xcb\x02\x00\x00")

func assetsTemplatesViewacHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesViewacHtml,
		"assets/templates/viewac.html",
	)
}

func assetsTemplatesViewacHtml() (*asset, error) {
	bytes, err := assetsTemplatesViewacHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/viewac.html", size: 715, mode: os.FileMode(420), modTime: time.Unix(1489232941, 0)}
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
	"assets/s/main.css": assetsSMainCss,
	"assets/templates/.transactions.html.swp": assetsTemplatesTransactionsHtmlSwp,
	"assets/templates/familypage.html": assetsTemplatesFamilypageHtml,
	"assets/templates/forms.html": assetsTemplatesFormsHtml,
	"assets/templates/index.html": assetsTemplatesIndexHtml,
	"assets/templates/layout.html": assetsTemplatesLayoutHtml,
	"assets/templates/newaccount.html": assetsTemplatesNewaccountHtml,
	"assets/templates/newuser.html": assetsTemplatesNewuserHtml,
	"assets/templates/transactions.html": assetsTemplatesTransactionsHtml,
	"assets/templates/userhome.html": assetsTemplatesUserhomeHtml,
	"assets/templates/view.html": assetsTemplatesViewHtml,
	"assets/templates/viewac.html": assetsTemplatesViewacHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"s": &bintree{nil, map[string]*bintree{
			"main.css": &bintree{assetsSMainCss, map[string]*bintree{}},
		}},
		"templates": &bintree{nil, map[string]*bintree{
			".transactions.html.swp": &bintree{assetsTemplatesTransactionsHtmlSwp, map[string]*bintree{}},
			"familypage.html": &bintree{assetsTemplatesFamilypageHtml, map[string]*bintree{}},
			"forms.html": &bintree{assetsTemplatesFormsHtml, map[string]*bintree{}},
			"index.html": &bintree{assetsTemplatesIndexHtml, map[string]*bintree{}},
			"layout.html": &bintree{assetsTemplatesLayoutHtml, map[string]*bintree{}},
			"newaccount.html": &bintree{assetsTemplatesNewaccountHtml, map[string]*bintree{}},
			"newuser.html": &bintree{assetsTemplatesNewuserHtml, map[string]*bintree{}},
			"transactions.html": &bintree{assetsTemplatesTransactionsHtml, map[string]*bintree{}},
			"userhome.html": &bintree{assetsTemplatesUserhomeHtml, map[string]*bintree{}},
			"view.html": &bintree{assetsTemplatesViewHtml, map[string]*bintree{}},
			"viewac.html": &bintree{assetsTemplatesViewacHtml, map[string]*bintree{}},
		}},
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

