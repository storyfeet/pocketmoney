// Code generated by go-bindata.
// sources:
// assets/s/main.css
// assets/templates/.familypage.html.swp
// assets/templates/familypage.html
// assets/templates/index.html
// assets/templates/layout.html
// assets/templates/newaccount.html
// assets/templates/newuser.html
// assets/templates/transactions.html
// assets/templates/userhome.html
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

var _assetsSMainCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\xcb\x31\x0e\xc2\x30\x0c\x40\xd1\xdd\xa7\xb0\xc4\x8a\xd8\x32\xb4\x3d\x4d\x8c\xeb\x2a\xa2\xb1\x2b\x13\x4b\x20\xc4\xdd\x09\x15\x4c\xfd\xf3\xfb\x64\xfc\x7c\x01\xf6\x28\x5f\x6f\x8b\x5b\x28\x8f\xa7\x61\x10\x21\x9a\xe0\x0d\x70\xa9\xb3\x06\x1e\x49\x4a\x22\x39\xef\x44\xcc\xeb\x5f\x98\xf3\xec\x63\xda\x1e\x78\xb7\xb5\x30\xd2\xda\x9f\x5d\x15\xdd\xa2\xe1\x19\x29\x5a\x33\xfd\xf9\x9a\x7d\x29\xfa\xf5\x13\x74\xf4\x09\x00\x00\xff\xff\x2d\xd0\xa6\xa1\x8e\x00\x00\x00")

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

	info := bindataFileInfo{name: "assets/s/main.css", size: 142, mode: os.FileMode(420), modTime: time.Unix(1488548494, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesFamilypageHtmlSwp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9a\x3d\x6f\xd3\x40\x18\xc7\xaf\x65\x60\xa1\x14\x81\x84\xd4\xed\x70\x97\x56\x6a\xec\x26\xa5\x02\xaa\x34\x12\xaa\xd4\xaa\x43\xaa\x4a\xa8\x15\xeb\xd9\x7e\x12\x5b\xf5\xd9\xe6\xee\x4c\x1b\x59\x81\x99\x01\xc4\xcc\xe7\x80\xb5\x33\x1f\x81\x1d\x89\x89\x0d\x21\x56\x1e\xbf\x24\x69\xc2\x90\x05\x81\x10\xcf\x4f\xfa\xe7\x7c\x2f\xcf\xab\x9d\x21\x8a\xdd\xcd\xd3\xc3\x2e\x7f\x60\xdf\x67\xc8\x2d\xc6\x06\x9f\x2e\x9f\xbe\xb9\xbe\xb1\xf8\x6d\x85\x31\x29\x8c\x09\xe0\x9c\xcd\xe3\x00\x94\x19\xb4\xe6\x1e\x63\x2f\x6a\x87\x8e\xf6\x54\x98\x1a\xed\xf4\x93\x48\xc4\x7d\x47\x2b\xcf\xe9\x87\x26\xc8\x5c\xdb\x4b\xa4\xe3\x25\x3e\x28\x2f\x89\x9f\x27\x03\x27\x4d\xbc\x33\x30\x32\x89\x61\xe0\x08\xad\x01\x8d\x0c\xc8\x34\x12\x06\xb4\xd3\x13\x32\x8c\x06\xa9\xe8\x83\x1d\x18\x19\xcd\x4f\x80\x20\xfe\x57\x32\xd3\x6b\x3c\x5c\x62\x5b\xad\xe6\x66\x31\x5d\xb5\xee\xf1\x3b\xb7\x4f\xfe\x76\x56\x04\x41\x10\x04\x41\x10\x04\x41\xfc\x41\x4c\xba\xc0\x5e\xe2\xb8\x58\xcf\x57\xea\x71\x61\x66\x24\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\xe2\xdf\x45\xf8\x8c\x79\x37\x18\xfb\x8c\x2a\xfe\xff\x1f\xfd\xfe\xff\xba\xcc\xd8\x25\xea\x03\xea\x1d\xea\x35\xea\x00\xf5\x08\xd5\x40\xad\xa2\xee\xa2\xbe\xdf\x64\xec\x23\xea\x3d\xea\x15\x6a\x80\x3a\x45\xad\xa2\xbe\x2c\x31\xf6\x16\x95\xa2\x1a\xa8\x65\xd4\x35\xd4\x8f\x3a\x5e\x19\x93\x20\x08\x82\x20\x08\x82\x20\x08\xe2\xf7\x91\xe7\xa3\x57\x86\xb9\xe5\x0a\x0d\xd6\x70\xc8\x38\x92\xe7\x10\xfb\xf5\x75\xdb\xe9\x25\x4a\x76\xca\xeb\x72\x1e\xc6\x69\x66\xb8\x19\xa4\xb0\x6b\xe9\xcc\x95\xa1\xb1\x26\xbb\x7b\x49\xdc\x0b\x95\xe4\xc7\x42\xeb\xf3\x44\xf9\x7c\x67\xda\x20\xad\xd7\x2d\x1e\x0b\x59\xcc\xcf\xfd\x96\xd5\x69\xbb\x6a\xe2\x62\xad\x1b\xc6\xa1\xcc\x24\xdf\xe6\x11\x18\x03\x4a\xaf\x4f\xed\xcf\x77\x9d\x8a\xc2\x2c\xde\xb5\xec\x7c\x7b\xa3\xb5\x39\xb4\xb8\x82\x67\x59\xa8\xc0\x9f\x44\x6d\xce\x44\x3d\xd4\xe8\x58\x41\x6c\x66\xbc\x7a\x01\x78\x67\x6e\x72\x31\x4e\xb8\x3c\x34\x63\x7c\x84\x5b\xb3\xe9\x18\xb8\x30\x23\xa3\x4c\x63\x3a\x78\x75\xa5\x51\xed\xa0\xd5\x79\xec\xfb\x7c\xbf\x7c\x53\x9b\x77\x41\xba\xa0\xda\x0e\xae\x8e\x8f\x5c\xbd\x0b\x95\xc9\x16\xf7\x22\xac\x72\xd7\x92\xa0\xb5\xe8\xa3\xbf\x3c\xb7\xbb\xa0\xf9\x70\x88\xa6\x5b\x57\x4d\xc3\x1e\x2f\x76\x46\x37\xb1\xb8\x87\x5c\x78\x26\x4c\xb0\x2d\xc2\xf7\x65\x19\xcf\xe2\x12\x4c\x90\xf8\x58\x56\xa2\xb1\x28\xce\xc6\xc6\x6b\xa1\xae\x4a\xe5\xf6\x3e\x1e\xc6\x4f\x21\xd7\x47\xde\xb0\xf6\x71\xfd\x6d\x27\x8b\x3a\xbf\x3c\x35\xd5\x4e\x14\x4e\x52\xaa\x4e\xd8\x27\x75\x2b\x86\x43\xde\x68\xe0\x82\xfb\x04\x22\xf0\x30\x4a\xd5\x7d\x6e\x55\xa3\xc5\xad\xbd\x20\x8c\x7c\x7c\x22\xa7\x3a\xdd\x1e\xb9\xcc\x73\x25\xe2\x3e\x94\x79\xd9\x55\xf7\xc6\xc5\x8e\x12\x6a\x07\xcd\x4e\xdd\xdf\x63\xec\xd6\x4e\x11\xbf\x38\x5e\xad\x1d\x95\x59\x60\xdf\x9a\x9d\xa9\x2f\x42\x00\xa2\x08\xcb\x7e\x06\x00\x00\xff\xff\x55\x08\x11\x01\x00\x30\x00\x00")

func assetsTemplatesFamilypageHtmlSwpBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesFamilypageHtmlSwp,
		"assets/templates/.familypage.html.swp",
	)
}

func assetsTemplatesFamilypageHtmlSwp() (*asset, error) {
	bytes, err := assetsTemplatesFamilypageHtmlSwpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/.familypage.html.swp", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1488902512, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesFamilypageHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x93\xc1\x6e\x9c\x30\x10\x86\xef\x79\x8a\xd1\x9c\x12\xa9\x81\x86\x28\x97\x08\x90\xaa\x48\x91\x7a\x48\xb5\x52\xd5\x07\x30\x78\x36\x58\xb5\x0d\xb5\x8d\x76\x57\x16\xef\x5e\x83\x61\x59\xb8\x64\x2f\x8b\xcd\xff\xff\x33\xf3\x19\x7b\xef\x48\x75\x92\x39\x02\x6c\x88\x71\x1c\x86\x3b\x08\xbf\xbc\x79\x2a\xdf\x99\x12\xf2\x02\x07\xf6\x49\xaf\xe0\x7d\x12\xd6\x49\xdc\xfb\xc5\x14\x0d\x43\x9e\x06\x51\x54\xf7\x32\x3e\x78\x6f\x98\xfe\x24\x98\xb4\x1f\xa4\x2a\x32\x76\x4e\x9c\x74\x52\x94\xd7\x45\xd4\x27\x7f\x2c\x19\x3d\xe5\xc1\xe3\x63\xd8\xa8\x7e\x93\xa4\xda\x41\x72\x60\x86\xb4\x03\x8c\xff\x08\xf8\xd6\x08\x39\x36\x98\x57\x66\x4d\xc9\xd3\x25\xd3\x7b\xd2\x7c\x69\x3f\x5d\x3a\x1a\xc5\x57\x83\xf7\xe2\x08\xf7\xc2\x76\x31\x3a\x79\x57\xa4\xa6\x5e\x1f\x16\xdf\xb1\x35\x0a\x58\xed\x44\xab\x0b\x64\x9c\xab\x69\x06\x04\x45\xae\x69\x79\x81\x5d\x6b\x1d\x96\x70\xb7\x0e\x10\x12\xc3\xa4\x9b\x29\x9b\x67\xa8\x25\xb3\xb6\x40\x45\xd6\x06\x7c\x58\x86\x41\x83\x08\x26\x68\xcf\xe5\x8d\x7d\xed\x39\x5a\xb3\xf2\x07\xe7\x30\xa3\x8f\x04\x83\x25\x5b\x2d\x23\x7b\x78\x85\x5c\xe8\xae\x77\xe0\x2e\x1d\x15\xe8\xe8\x1c\x00\x8d\x14\x0b\xec\x67\x9e\xb8\x5a\x7e\x5a\x88\x10\x77\xb6\xba\xa1\xfa\x6f\xd5\x9e\x17\x6b\xc4\x82\xe5\x06\xf0\x21\x0c\x72\x6a\x0d\xdf\xd7\xec\xe6\x7d\x84\x8e\x39\x17\x6a\x16\x98\xf8\x97\x6f\xd9\xf7\x01\xc1\xd0\xbf\x5e\x18\xe2\x4b\xee\x89\x3f\xed\x52\xef\x3f\x84\x16\xaa\x57\xf0\x02\x92\x46\xbb\x7d\xd8\xbc\x7f\x6b\xf5\x51\x84\xa3\xf8\xba\xfa\xb5\x44\xb6\x2b\xb1\xd1\xdb\xbe\x52\xc2\xcd\x4c\xf2\x74\x3c\xe6\xed\x57\x73\x7b\x11\x2a\x66\x69\xbc\x08\xff\x03\x00\x00\xff\xff\x3b\xe2\xc7\x9e\x1d\x03\x00\x00")

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

	info := bindataFileInfo{name: "assets/templates/familypage.html", size: 797, mode: os.FileMode(420), modTime: time.Unix(1488902512, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x93\xcf\x6e\xa3\x30\x10\xc6\xef\x3c\xc5\xc8\xa7\x5d\x69\x37\x7f\x90\x72\x59\x01\xd2\x2a\xda\x3d\x2d\xab\x5c\xfa\x00\x06\x86\x62\x15\xdb\xd4\x1e\x9a\x46\x88\x77\xaf\x09\xa1\x21\x09\x69\xda\x70\xc2\xe3\xef\xfb\x6c\xfd\x66\xdc\x34\x84\xb2\x2a\x39\x21\xb0\x02\x79\xc6\xda\xd6\x83\xc3\x17\x14\xcb\x68\xa3\xd3\x27\x24\x88\xb5\xc2\x5d\x30\x77\x05\xef\x7d\xbb\x69\x44\x0e\xb3\xb1\x3e\x89\x9a\x66\x16\xa3\x6d\xdb\x60\x9e\x44\x23\x21\xaa\xcc\xe9\xbc\xa3\x37\xc8\xb5\x91\xc0\x53\x12\x5a\x85\xac\xd4\x8f\x42\x31\x90\x48\x85\xce\x42\x56\x69\x4b\xec\x68\xef\xaf\xe2\x47\xff\x3a\x95\xbb\x83\x7f\xba\xf5\x97\x4b\x51\xee\x7e\x41\x20\x54\x55\x13\xd0\xae\xc2\x90\x11\xbe\x12\x03\xc5\xa5\xfb\xcf\xf7\x02\x06\x2f\xbc\xac\xdd\xd2\x5d\xb1\xb7\xb4\x2d\x8b\x82\xc4\x9c\xa6\xc5\x28\x13\x34\xf0\xdf\x39\xe1\x83\xcc\xda\xa2\xe9\xfe\xc6\xa9\x0f\x87\xda\x64\xee\x86\x5b\xbb\xd5\x26\x3b\x0f\xad\x0e\xf5\x21\xb8\xda\x66\x13\xee\x13\x8b\xad\x13\x29\xc6\x80\x82\x79\x47\x33\xba\x8a\x57\xe1\x76\x60\x70\x03\xf1\xda\x60\x37\x09\x3d\x1f\xf8\x9d\xa6\xba\x56\x74\x0d\xf9\x4d\x48\xfd\xa1\x7b\x4c\x13\xa0\xb9\x50\xd0\x21\xfb\x14\xe6\x4b\xff\x1f\xc9\x45\x79\xee\xc5\xae\x38\x98\xfb\xc5\x3d\xad\xa8\x38\x91\x3b\x38\x64\xb3\x66\xf5\xc3\x5f\xb4\x0c\x0c\x3e\xd7\xc2\x60\x06\x24\xa8\x74\xda\xd5\x4f\x7f\x01\x69\xc1\x8d\x63\x8c\xc6\x8e\xba\xb7\x9c\x38\xf1\x5b\x2c\x94\x90\xb5\x84\x15\x94\xd8\x45\xdb\xef\x17\x9a\xb5\x56\xb9\x70\x4d\xfb\xc2\xa0\xf8\x77\x4f\xca\xf8\xd5\x27\xdc\xa2\x7b\xf5\x83\xc4\x7b\x0b\x00\x00\xff\xff\x2d\x4d\xb6\x58\x12\x04\x00\x00")

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

	info := bindataFileInfo{name: "assets/templates/index.html", size: 1042, mode: os.FileMode(420), modTime: time.Unix(1488869859, 0)}
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

var _assetsTemplatesTransactionsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8f\xc1\xca\xc2\x30\x10\x84\xef\x79\x8a\x50\x7a\xfc\x69\xe1\x3f\x96\x18\x28\x4a\x9f\x40\x1f\x60\x6b\x17\x2b\x36\xa9\x24\x6b\x2f\x61\xdf\xdd\x8d\xa2\xd6\x83\x39\x04\x66\xe6\x23\x99\x49\x89\xd0\x5d\x27\x20\xd4\xc5\x88\x30\x14\xcc\xca\x8c\xff\x76\x1f\xc0\x47\x38\xd2\x79\xf6\xd1\xd4\x62\x28\x43\xd0\x4f\x68\x55\x4a\x12\x9d\x50\x97\x97\xbf\x72\xd1\xcd\x46\x57\x1d\xb8\x6a\xcd\xeb\xfc\x06\x05\xab\xb4\x1c\x43\x83\x4d\xa9\x5c\xaa\x2e\xcc\xee\x10\x31\x30\x37\x6f\xdd\x6e\x99\x4d\x2d\xc4\x37\xba\xc3\x48\x6b\x34\xeb\x1f\x68\xeb\xe6\x9b\xa7\x57\x24\xb7\x7c\x2b\x15\xd1\x0f\x8f\x12\xcf\xca\xd9\xf9\xcc\xec\x21\x62\x9e\x79\x0f\x00\x00\xff\xff\xf0\xce\x3b\xc5\xfb\x00\x00\x00")

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

	info := bindataFileInfo{name: "assets/templates/transactions.html", size: 251, mode: os.FileMode(420), modTime: time.Unix(1488553875, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesUserhomeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\x4d\x6e\xdb\x3c\x10\xdd\xeb\x14\x03\x22\xab\x20\x9f\x9d\x38\xdf\xca\x90\x04\x04\x01\xb2\x4b\x6a\xa0\xed\x01\x68\x73\x5c\x09\x11\x7f\x42\x51\x71\x05\x41\xa7\xe9\x4d\x7a\xb2\x0e\x49\xc9\x91\xe3\x38\xf5\xa2\xd9\xc4\xe6\xcc\x7b\x7c\x6f\xe6\x49\xee\x3a\x87\xd2\x54\xdc\x21\xb0\x02\xb9\x60\x7d\x9f\xa4\xc5\x22\xff\x5e\xa3\x85\x15\xff\x81\xb0\x84\xae\x9b\x3d\x48\x94\x7d\x9f\xce\xa9\x42\xe5\xdb\xfc\x6e\xb3\xd1\x8d\x72\x35\x9d\xdc\xe6\x49\xd7\x95\x5b\x98\x3d\x62\x1d\xc0\xb7\xb0\xa9\x78\x5d\x67\x4c\x62\x5d\x13\x03\xcb\x89\x20\x14\xc7\x6e\x54\x82\x3a\xbb\x6e\x7e\xf9\xb5\xd0\x3b\x18\xc9\x2e\xe7\x74\x9a\xa4\x4d\xe5\x7b\x2e\xfc\x95\xb0\xcc\x60\xb8\x9b\x8e\x2c\x57\xa4\xe7\xe2\xf9\x0a\x2e\x5e\x63\x85\xcb\xd9\x08\xa6\x0e\xa0\xbf\x20\x05\x5f\xa8\x63\xe6\x2d\x28\x2e\x09\x31\x30\xf8\x7a\x5a\x95\x24\x87\xaa\x4f\x54\xe9\x7b\xf8\xfd\x2b\x7c\xbb\x6f\xac\x45\xe5\xbc\x44\x6a\x18\x98\x46\x99\xf1\x7f\x3a\xf7\xc2\x82\xea\x3b\x21\x46\xd1\x41\x73\xba\xd5\x56\x02\xdf\xb8\x52\xab\x8c\x71\x21\x78\x2c\x32\x90\xe8\x0a\x2d\x32\x66\x74\xed\xd8\x30\xba\x37\x70\x9c\x47\x5a\x2a\xd3\x38\x70\xad\xc1\x8c\x39\xfc\x49\x30\x2f\x3b\x63\xcd\x60\x80\x81\xa5\xcd\x68\x55\xb5\xf0\xca\xab\x86\x2a\xfb\x8d\x30\x28\x4a\x21\x50\x7d\xc2\x32\x68\x89\x44\x86\x3b\x47\xa4\x19\x9b\x75\xff\x5f\x2d\xae\x7b\x4f\xfd\xd2\x94\x16\xc5\x3b\x86\xba\x59\xcb\x92\x38\x3c\x53\xc6\x48\xb3\x57\x3f\xf7\x3e\x69\x06\x61\x08\x2b\xde\xc2\xa3\x56\xd8\x7e\x30\x02\xc3\xdb\x23\xef\x61\xfa\xe4\xf7\x91\x3f\x23\x70\xca\x56\x2b\xf1\x9f\x8e\x00\x20\x79\xb0\x5a\x2e\x93\xb4\xc6\x0a\x37\x6e\xa0\xd8\xd2\x19\x1b\x57\x3a\x46\x08\x3e\xcb\xd0\x39\x39\x0a\x6e\xb4\xf1\x76\xdf\x14\xed\x63\xc5\x0e\x32\xf6\x1f\xbc\xcf\x58\x04\xe6\x93\xdb\x62\xc6\xa6\x9f\xd3\x79\xb4\x91\x27\xdf\xf4\x7b\x4f\x4e\x1f\x39\xba\x3a\xed\xe7\x23\x9d\xa3\xa9\xbe\x5f\x1e\xeb\xfe\xb8\x78\x28\xfb\x48\x66\xba\xb6\x79\x72\x27\xfd\xc5\x4b\xa0\xe7\xea\x60\xa7\xaa\x91\x6b\xb4\x0c\x6a\x87\x26\x63\xd7\xb3\xeb\x9b\x7d\x3c\x03\x82\xd1\xf6\x56\x8d\xa5\xa8\x20\x81\x4f\xc6\xc1\xc4\x16\x16\x2f\x0b\xde\x4e\x87\xd6\x07\xf4\x49\xef\x26\xc1\x3d\x0c\xe9\xa6\x30\xf4\x9a\x3a\x91\xd3\x45\x7e\x5f\x84\xc9\xae\xa8\x67\xa7\xad\x88\xef\x3e\x5f\xfc\x52\x89\xfd\x29\xbd\x1d\x0f\x14\x98\xe1\x7c\x14\xac\x2b\x61\x76\x62\xa2\xf7\x09\x77\x60\xce\x04\x13\xf2\x66\x02\xbd\xd7\x6a\x5b\x92\xfe\xbf\xa3\x16\xe7\x0d\x28\x3a\x9c\x3e\xd8\xd3\x5f\x83\x35\xa7\x41\xd3\x82\xff\x04\x00\x00\xff\xff\x56\xe0\x09\x52\x22\x06\x00\x00")

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

	info := bindataFileInfo{name: "assets/templates/userhome.html", size: 1570, mode: os.FileMode(420), modTime: time.Unix(1488833487, 0)}
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
	"assets/templates/.familypage.html.swp": assetsTemplatesFamilypageHtmlSwp,
	"assets/templates/familypage.html": assetsTemplatesFamilypageHtml,
	"assets/templates/index.html": assetsTemplatesIndexHtml,
	"assets/templates/layout.html": assetsTemplatesLayoutHtml,
	"assets/templates/newaccount.html": assetsTemplatesNewaccountHtml,
	"assets/templates/newuser.html": assetsTemplatesNewuserHtml,
	"assets/templates/transactions.html": assetsTemplatesTransactionsHtml,
	"assets/templates/userhome.html": assetsTemplatesUserhomeHtml,
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
			".familypage.html.swp": &bintree{assetsTemplatesFamilypageHtmlSwp, map[string]*bintree{}},
			"familypage.html": &bintree{assetsTemplatesFamilypageHtml, map[string]*bintree{}},
			"index.html": &bintree{assetsTemplatesIndexHtml, map[string]*bintree{}},
			"layout.html": &bintree{assetsTemplatesLayoutHtml, map[string]*bintree{}},
			"newaccount.html": &bintree{assetsTemplatesNewaccountHtml, map[string]*bintree{}},
			"newuser.html": &bintree{assetsTemplatesNewuserHtml, map[string]*bintree{}},
			"transactions.html": &bintree{assetsTemplatesTransactionsHtml, map[string]*bintree{}},
			"userhome.html": &bintree{assetsTemplatesUserhomeHtml, map[string]*bintree{}},
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

