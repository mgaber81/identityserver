// Code generated by go-bindata.
// sources:
// index.html
// registration.html
// login.html
// home.html
// error.html
// apidocumentation.html
// DO NOT EDIT!

package html

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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x57\xef\x6e\xdb\x36\x10\xff\x5e\xa0\xef\x70\x55\x3f\xf8\x4b\x24\x75\x49\xb6\x0e\xad\xed\x21\x48\x8a\x22\xc0\x80\x16\x75\x80\xa2\x9f\x02\x5a\x3a\x89\x44\x25\x52\x25\x29\xbb\xc6\x36\xa0\x0f\xb2\xbd\x5c\x9f\x64\x77\xa4\x54\xdb\x71\xb3\x7a\xff\x10\x18\x22\xc5\xfb\xfb\xfb\xdd\x9d\x98\xe9\xa3\xab\x57\x97\x37\xef\x5e\xbf\x00\xe9\xdb\x66\xfe\xf0\xc1\x94\x9f\xd0\x08\x5d\xcf\x12\xd4\x49\x78\x83\xa2\xa4\x27\xc0\xb4\x45\x2f\x48\xd0\x77\x29\x7e\xe8\xd5\x6a\x96\x5c\x1a\xed\x51\xfb\xf4\x66\xd3\x61\x02\x45\xdc\xcd\x12\x8f\x1f\x7d\xce\x86\x9e\x43\x21\x85\x75\xe8\x67\xbd\xaf\xd2\x1f\x13\xc8\xa3\x21\xaf\x7c\x83\xf3\x6b\x3f\x71\xf0\xce\xf4\xf0\x4a\x37\x4a\xe3\x34\x8f\xaf\x83\x04\xbd\x78\x0f\x16\x9b\x59\xe2\xfc\xa6\x41\x27\x11\x7d\x02\xd2\x62\x35\x4b\x84\x23\x8b\x2e\x2f\x9c\xcb\xc3\x61\x46\xab\x24\xff\x5b\x7a\x15\x85\x9a\x8a\x35\x3a\xd3\x62\xd6\x2a\x7d\x60\x22\x68\x4c\x38\x57\xf7\x2c\x0f\xe2\x2e\xab\x8d\xa9\x1b\x14\x9d\x72\x59\x61\x5a\xb6\xf3\x53\x25\x5a\xd5\x6c\x66\x6f\xb1\x69\xaa\x86\x5c\xfd\x7a\x61\x57\xe6\xd9\xf9\x93\x27\x27\x4f\xe9\x47\x4f\xe5\x45\xa3\x0a\xde\xc5\xd5\x24\x04\x37\xd9\x06\x37\x01\x4f\xe0\xcd\x26\x01\x33\x32\x39\x89\x41\x04\x81\xb0\x94\xa7\xf0\x0b\xf0\x69\xea\xa4\x28\xcd\xfa\x19\x9c\x76\x1f\xc3\xef\xf1\xd9\x25\xff\x3d\x87\xdf\x1e\x3e\x08\x4a\xf9\xa8\x35\xcd\x07\xd2\xa6\x4b\x53\x6e\x46\x12\xd1\x82\x2a\x67\x49\x5c\x26\xd1\x4f\xa9\x56\x50\x34\x84\xcc\x2c\x19\xe8\x8b\x07\xc3\x11\xcb\x37\xa6\x36\xc9\x7c\x2a\x06\x14\x1f\x27\x73\xb8\xf6\x8e\x78\xcb\x22\x6f\xe4\x57\xcc\xa7\x39\x89\xef\x6a\x0e\x46\x9d\xaa\x75\xba\xb6\xa2\x1b\xcd\xd2\xf1\x68\xc9\x62\xad\x9c\xe7\x50\xa6\x6a\x94\xaf\x04\x54\x22\xed\x1d\xda\xb4\x6b\x7a\x47\x47\xb9\x9a\xbf\x19\x04\xd9\xd1\x81\x15\x0a\x4f\xe9\x43\x13\xc1\x71\x38\x20\x03\x0b\xda\x80\xd2\x5b\xfd\x6d\xb8\xe3\x2a\x42\x86\x76\xce\x58\x4e\x1f\xa5\x29\x2c\x7e\xbe\xbe\x7a\xb1\x80\xc5\xcd\xc5\x9b\x1b\x48\x53\x96\x19\x21\x71\x8d\x2a\xf1\xbb\x6f\x43\x28\x4f\xe7\x84\x93\x05\x13\x81\x22\x25\x4d\x35\xbe\x21\x5f\xa7\x3b\x22\x4a\xc3\x26\x48\xad\x35\x48\xa1\x4b\xf7\xe5\x7c\x1b\xdc\xf0\xdc\x0b\xe0\xf4\xfe\x00\xe0\x2b\x54\x7c\xe8\x8d\x47\x77\xcb\x22\x82\xa2\xb1\x3b\x8c\xc8\xb3\x7d\xa1\x64\x3e\x50\xdb\x59\x74\xa8\x0b\x84\x16\x85\x76\x04\x21\xb1\x20\x0a\xaf\x8c\xce\xe0\xd2\xb4\x9d\xd0\x0a\x1d\x68\xc4\x32\xa6\x30\x66\x08\x6b\xa9\x0a\xc9\xef\x06\x13\x9e\x74\xa1\x54\x55\x85\x96\x37\x6b\xb1\x71\x27\x50\x62\x87\xba\x54\xba\x26\x80\xc0\x4b\xa4\x09\xc2\x26\x37\x27\xd0\xa1\x75\xf4\xce\xd8\x38\x54\x3e\xfa\x8c\x30\x39\xfb\x0b\xf2\xbe\x86\xcf\xf7\xc7\x12\xc4\x7e\xdb\x5e\xab\x42\x78\x04\xb3\xa2\x3e\x59\x09\xab\x4c\xef\x78\x7c\x69\x8d\x8d\x03\xa2\xe5\x90\xb5\x6d\x3e\x43\x94\x0e\x96\x3d\x65\xca\x73\x4d\x34\x9c\x23\x23\x70\x24\x9b\x3f\xfc\x8f\x6c\x7e\xfe\xf4\xfb\x8d\xc4\x88\xf1\x4b\xab\x34\x51\xe3\xa5\xe9\x6b\xe9\xc1\x54\xc0\x43\xd0\x4b\xa6\x81\x8e\x69\xc6\xe8\x89\x87\x25\x56\xc6\xe2\x23\x78\x2b\x05\x25\x54\xc1\xa5\xb4\xd4\x82\xad\x20\xd6\x48\x68\x50\x26\x06\x0d\x3a\x16\x27\x00\x11\x2a\x6b\x5a\x10\xe0\x3c\x69\x66\x87\x9a\x59\x96\x11\xad\x52\x74\xbc\x8a\xf5\x24\xa0\x51\x9e\x66\x3e\x2c\x95\x87\x96\x1d\x7e\xfe\xf4\xc7\x0e\xd3\x94\x8c\x6a\x6b\x70\xb6\xf8\x32\xbd\x69\x9f\x97\x36\x75\xd8\x3b\x97\x75\xba\x4e\x08\x69\xea\x6e\x9a\x03\x58\xf9\x61\x84\x47\x14\xce\x47\x14\x44\x4f\xf1\xda\x5b\x2d\x5a\xbc\xad\xad\xd8\x24\xf3\x2b\x9b\xc1\x82\x2d\x10\x27\xf2\x7c\xab\xb3\xb4\xf9\x3d\x9b\x43\x4c\x2f\xf4\x6e\xb5\x13\xb8\xdb\x22\x62\x28\x15\xb5\x85\xf1\x20\x96\x86\x2a\x22\xd4\x76\x2c\x25\xae\x69\xde\x76\xd6\xac\x48\xdf\x9e\x0c\xe5\x12\xe4\xd6\xd2\x84\x9e\x11\x96\x30\x16\xcd\x7b\x36\xe4\x0d\xb7\x4e\xec\xa7\x7f\xd1\x0c\x67\x47\x35\xc3\x42\xb2\xeb\xbe\x63\xaf\x25\x77\x83\xd2\x54\x07\xad\xe0\x8e\xbf\x53\xff\x5c\xe9\x34\xd7\x36\x20\x1c\xb4\x3d\x05\x27\x42\xb5\x53\x6b\x6b\x7f\x64\xc9\x9f\x1f\xf7\x11\xfa\x27\x15\x7f\x65\x62\x15\x83\xa8\xac\x50\x25\xe7\xe3\xa4\x59\xc7\x29\x55\x09\x9a\x68\x34\x67\xda\x50\xa8\x31\x67\xe6\x64\x27\xd9\x98\x8a\x72\x32\x6a\x06\x09\xc3\xc5\x8d\x76\xa5\x8a\x2d\x7d\x3c\x18\xd8\x8d\xeb\x49\x22\x4e\x71\xb2\x64\xf7\x4c\x51\x29\x38\x51\xe1\x71\xac\xed\x7c\x78\x03\x46\x4f\xbf\x06\xc5\x1d\x94\x8e\xc5\x69\x5f\x8e\x08\xf4\x46\xa7\x05\xf2\x50\x8f\x37\x8c\xb8\xde\x55\xd8\x57\x29\x4d\xb1\x7f\x18\xbe\xc3\x5e\xd8\x9a\xae\x78\xc9\xed\x92\xee\x8d\xef\x93\xed\x3d\x61\x5f\x12\xe0\xca\x14\x7d\x4b\x2e\x02\x2c\x77\xcc\xec\x7c\xda\xf7\x01\xba\x27\x0e\x1a\x3d\xc5\xe9\xd1\xc1\xd0\xa5\xad\xdc\x75\x7e\x18\xdb\xc5\xeb\xeb\xff\x3c\xbe\xb3\xa3\xe3\x1b\x2f\x9a\xb5\xf2\xb2\x5f\x86\xeb\xa5\xf2\x8e\xea\x29\xde\x1c\x0e\xc3\x9d\x3a\xfa\x50\x8e\xee\xa2\x56\xca\x17\xd7\x43\x49\x80\x05\x95\x65\xc1\x1f\xd7\x12\xf9\x43\xfb\x32\x48\x1f\x18\xcc\xd9\xe2\xdd\x78\xbf\x91\xf5\xfe\x7e\xef\x0e\x78\xff\x58\xca\x5d\x61\x55\xe7\xc3\x72\xbc\x9e\xe6\xc3\xbf\x1f\x7f\x06\x00\x00\xff\xff\xd1\x2c\x15\x9f\x90\x0c\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 3216, mode: os.FileMode(493), modTime: time.Unix(1457982012, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _registrationHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x58\x7b\x6f\xdb\xba\x15\xff\x7b\x05\xfa\x1d\x38\xae\x98\x12\x20\x96\x1c\xe7\xe5\xa6\x96\x8b\xa0\x0f\xa0\x7f\x14\x0d\xb0\x14\x43\xb1\x15\x05\x2d\x1d\x49\x4c\x28\x51\x25\x29\x27\x5e\xd7\xef\x7e\x0f\x29\xc9\xf1\x43\x76\xe2\xdb\xe6\x1a\x08\x42\x52\xbf\xf3\x7e\xe8\x50\xa3\xbf\xbf\xfd\xf4\xe6\xea\xcb\xe5\x3b\x92\x99\x5c\x8c\x9f\x3f\x1b\xd9\xff\x44\xb0\x22\x0d\x29\x14\x14\x4f\x08\x19\x65\xc0\x62\x5c\xfd\x8d\xe0\x3a\x07\xc3\x10\x6c\xca\x1e\x7c\xaf\xf8\x34\xa4\x6f\x64\x61\xa0\x30\xbd\xab\x59\x09\x94\x44\xf5\x2e\xa4\x06\xee\x4c\x60\x99\xbd\x22\x51\xc6\x94\x06\x13\x56\x26\xe9\x0d\x29\x09\x1c\x53\xf7\x1b\x19\x6e\x04\x8c\x3f\x18\x4f\x93\x2f\xb2\x22\x9f\x0a\xc1\x0b\x18\x05\xf5\xf1\xf3\x67\xf7\x40\x3c\xbf\x21\x0a\x44\xe8\x69\x33\x13\xa0\x33\x00\x43\x4a\x05\x09\x98\x28\xf3\x48\x86\xab\xd0\xb3\x6a\xe9\xf3\x20\x60\xd7\xec\xce\x4f\xa5\x4c\x05\xb0\x92\x6b\x3f\x92\xb9\x3b\x0b\x04\x9f\xe8\x00\x6d\xab\x04\x53\xdf\x72\x66\x40\x71\x26\x82\x43\xbf\xef\xf7\xdb\xe3\x5e\x7b\xec\xe7\xbc\xf0\x23\xad\xbd\x71\x87\x16\xf4\x5e\x0b\x5a\x0b\xa7\x4c\xa3\x8d\x3a\x40\x8a\xc0\x3d\xb4\xb4\x34\xd8\x9d\x38\x41\x0f\xf6\xd8\x2d\x68\x99\x43\xab\xc3\x3a\x9f\x65\x83\x2d\x8d\x5e\xb5\x18\xe9\x5e\x27\x2c\xe7\x62\x16\xfe\x1b\x84\x48\x04\xca\xfb\xff\x85\x9a\xca\xf3\xe3\x7e\xff\xe0\x0c\xff\xf0\x3f\x37\x4c\xf0\xc8\xee\xea\x95\xb7\xea\x64\x8f\x18\x0c\x6c\xe8\xb9\x78\x2e\xbb\x83\x8c\x1c\x6c\x31\x4e\xee\xf7\x9f\x22\xfd\xef\x79\x24\x24\xbb\xf9\x7a\x60\x37\xbd\x76\xed\xb7\x6b\xf2\x63\x99\xc2\xfe\x62\xae\x4b\xc1\x66\xe7\xa4\x90\x05\xbc\x5a\x7e\xfe\x73\x41\x66\xd0\x0a\xc5\x65\x93\x98\x98\xb6\x13\x19\xcf\x08\xb2\x67\x65\x19\x52\x05\x29\xd7\x46\x31\xc3\x65\x71\x51\x96\xb4\x86\x58\x30\x28\xc2\xe3\x90\xd6\xcb\x26\xbb\x63\x3e\x25\x91\xc0\x08\x84\xb4\xc9\x5e\xda\xd8\xe8\x1e\x59\xbc\x90\xa9\xa4\xe3\x11\x6b\xa2\xe5\x07\x74\x4c\x3e\x18\x8d\x29\xeb\xd7\x29\x8b\xba\xb0\xf1\x28\x40\xfc\x22\x69\xc3\x55\xf3\xb4\xe8\xdd\x2a\x56\xd2\xb9\xef\xe6\xac\x90\x33\xc7\x32\x1b\xf1\x16\x9c\x30\x92\xb0\x9e\x23\x71\x0f\x02\x3e\xfe\x17\x6e\x08\x2f\xac\x88\x86\xf9\x5c\x50\xbb\xaa\x5d\x01\xca\x59\x8a\xdb\x05\xe9\x49\x25\x84\x8e\x14\x40\x91\x48\x95\xcf\x4d\xb3\x1b\x32\x8f\x07\x56\x75\x26\xd1\xd0\x52\x6a\xcc\x49\x16\x59\xcf\xb5\x7e\x44\x47\x61\x43\x98\xc9\xca\x58\x07\x89\x2a\x2f\x28\x29\x58\x0e\xb5\x61\x55\xe9\xd8\x92\x45\xbb\x91\x2d\x2f\xb8\xc1\x34\x9a\x03\x7c\x0d\x6a\x0a\x4a\xf3\x18\xa6\x98\x67\xb1\x8b\x0d\x28\x25\x95\x26\x21\xf9\x41\x79\xe1\x8e\x6d\xd2\x9b\x8c\x17\x29\x3d\x27\x46\x55\xf0\xd3\x5b\xf7\x2a\x26\xf2\xdd\x5c\x21\x25\x6f\x29\x69\x76\xbd\x3b\x3d\xd7\x70\x31\x4b\x5b\xa2\x90\x9e\xf4\xd7\x4c\x19\x2f\x67\xda\x28\x8f\x51\xf7\x12\x99\xd9\x5c\x60\x18\x5a\x35\x5e\xcf\xd5\x91\x60\x13\x10\xe3\xcf\x68\x94\xf5\xc4\x28\xa8\xf7\x1d\x40\xc7\xcb\x3a\x24\x97\xb1\x2d\xfd\x3a\xde\x04\xa5\xe4\xec\x4e\x40\x91\x9a\x2c\xa4\x47\xa8\x16\x56\x7a\xbb\x1d\x50\x2c\x42\x6c\xae\x0a\xe2\xc6\xd1\x0d\x95\xab\x44\xd7\x59\x31\x46\x95\x91\x89\x8c\x2a\xdd\x25\xb5\x09\x42\x0e\x5a\xb3\x14\xf4\x62\xa4\xb6\x04\xc2\xa9\x65\xf9\xf6\x32\x7c\x6a\x53\x51\x68\x58\x75\x50\xa7\x84\x90\xc6\x55\x89\xdd\x03\x5b\x67\xd5\xf8\x84\x8e\xaf\x32\xae\x49\xbb\x25\xb8\x66\x42\x61\x92\xce\x88\x61\x37\x50\x2c\x86\x75\x89\x73\xc7\xf9\x28\xe8\x8c\xca\x2f\x44\xee\x5d\xce\xb8\xd8\x21\x6c\x60\xf1\x6b\x61\x69\x4e\xeb\xb0\xd4\x9b\x27\xd7\xfc\x12\x8b\xfa\x56\xaa\x78\x07\xe5\xcb\x86\x64\x4d\xff\xfb\x07\xb5\x09\xf3\xfd\x5f\x66\x85\x9d\x17\x12\xae\x72\x97\x85\x7f\xc2\xa4\xfb\x14\xde\x68\xdc\x22\x04\x19\x94\xcc\x60\x47\x2b\x1e\x6f\xbc\x93\xbf\xad\xa0\xd6\x05\xf9\x2f\x5c\x41\x75\x97\x4e\x47\xf1\x34\x3a\x35\x25\x13\x4b\xd0\xf8\x12\x34\x04\xbd\x12\x65\xc4\x64\x40\xca\x79\xd0\x7f\xb9\x68\x36\x91\xac\xb7\xd6\x0d\x0d\xb2\xb3\x05\x77\xe9\xa4\x4b\x56\x38\x24\xb6\x70\xbb\xde\xe4\xd7\xd5\x77\x4b\xdd\xaa\x4f\xfb\x1b\xdd\xf7\xa8\x44\x9b\xa3\xeb\x9c\x1a\xf4\xde\xe3\x5b\x4d\x2a\xdb\x37\x33\x7c\xc5\xdb\x5e\x85\x91\xc2\xfc\x8b\xb7\xb4\xef\x39\x93\xd5\xfc\x33\xd2\x94\x96\x74\xb5\x93\x9f\x2e\x35\xf2\xd3\xb5\xa4\x9c\xd3\x6d\x13\xf6\xe4\xdd\x7b\x83\xa4\xb0\x7d\x0b\xdf\x6b\x79\x61\x07\x0f\x77\xe6\x3c\x45\x6e\x99\x26\x29\x9f\x6e\x6e\xdf\x73\xd6\x5b\x9e\x3f\x90\x99\x4b\xd0\xef\xca\xc9\x75\x36\xdb\x91\x04\x5d\xea\xcc\x45\x4a\xa5\xc0\xcd\x29\x3d\x01\x53\x1b\x93\x8f\x94\x68\xfe\x3f\xb4\x62\xd0\xc7\xf7\x29\xba\x86\x85\x14\x0d\xb1\x01\xc7\x11\xd9\xda\x14\x34\x33\x9b\x74\x33\xdb\x6b\x0d\x38\x16\x99\xf0\x70\x70\x74\xbc\xf8\xf7\x4f\xae\x75\x05\x2a\x5c\x1a\xf0\xda\x41\x7d\x1b\xc3\xf3\xdd\x38\x62\x59\xd4\xd6\x3d\xb2\xa0\xdd\xf9\x03\x45\xd5\xd2\xad\xd7\xf9\xf2\xe4\xb8\x52\xf7\x5b\x4a\xf9\x41\x89\x18\xcb\x49\x65\x0c\xd6\x52\xdd\x47\x75\x35\xc9\x39\x0e\x28\xcd\xf8\x89\x8f\x15\xe3\x1a\x2b\x00\x57\xa5\xe2\x39\x53\x33\x5a\x8f\xb4\x55\xe9\x32\xa1\xa6\xde\x41\xec\xd2\x44\x18\xd8\xb2\x70\x13\x70\x6b\x19\xae\x71\xde\xe5\xa5\x21\x5a\x45\x3b\x5e\x0a\xaf\x35\xde\x06\x8f\xfd\x61\xbb\xf7\xaf\xb5\x9d\x40\x6b\x7e\xe3\xdf\xc9\xb9\xc7\x0a\x6e\xaf\x9a\xee\x8a\xf7\x74\x52\xda\x26\xf2\xc4\x62\x18\xde\x98\x7f\xaf\x88\x47\x5d\xcf\xbb\xa5\x35\xa9\xd1\x86\x10\x5b\x76\x25\x60\xcf\x5b\xb9\x15\x7a\x78\x37\xf5\x8a\xf4\x63\xc3\xcf\x3b\xb0\x9b\xc6\x5b\xb8\xc9\x65\x21\x31\xef\x22\x88\xfd\xba\x4a\xbd\xaf\xfb\xf7\x29\xe8\xbb\xb9\x25\xdd\x4b\xaa\xc2\xf5\xa0\xbd\x17\x79\x7c\x95\x01\xea\x94\x5e\x2a\x39\xc5\x06\xac\xf6\x57\x2f\xb9\xeb\x10\x3f\x86\x04\x3b\xc1\x25\x13\x80\x03\xc0\x9e\x37\x11\x15\x70\x9d\xa1\x66\x1d\xf7\x63\xef\xa4\xef\x9d\x13\xef\x1f\xc9\x59\x32\x49\x62\xef\xa0\x03\x71\xd8\xaf\x21\x13\xbc\x0a\x42\x37\x64\xd0\x40\x86\x10\x0d\x61\xd0\x09\x39\x6a\x20\x27\x27\x2c\x8e\x8f\x3a\x21\xc7\x0d\xe4\x28\x62\x87\x51\xb7\xa0\x93\x16\xd2\x7f\x79\x38\x99\x74\x42\x4e\x1b\xc8\x80\x9d\x01\xeb\x16\x74\xd6\x42\x8e\x4e\x27\x43\xd6\x09\x19\x36\x90\xc3\xf8\x64\x78\xd6\x6d\xd1\xcb\x16\x72\x76\x8c\x46\x75\x42\x2e\x5a\xdf\x6d\x71\xef\xc5\xe0\x61\xff\x5e\x3c\xc2\x35\x17\x8f\xb0\xca\xbe\x18\x15\xd3\xe6\x2d\x24\xac\x12\xe6\x8d\x14\x52\x59\x1a\xc1\xd3\xcc\x6c\xa7\x60\xea\xc6\xc1\xb5\xc5\x9f\xf4\x09\x5a\x46\x50\x73\x82\x71\x25\xa8\x1d\xb1\xa6\x12\x6b\x0b\x71\xca\xae\x7c\x69\xd9\x7f\xf5\x60\xd6\xe2\x00\x95\x63\xb6\xc6\xb5\x6a\xde\xfe\xba\x32\x7e\xd3\xe8\xd7\x32\x7b\x01\xeb\x24\x6d\x68\x15\xd4\x64\x5c\xc5\x25\x53\x66\xd6\x7c\x19\xab\x6b\xb0\x97\x02\xce\x0a\x0c\xe7\xb8\xe0\xba\x3d\xc3\x26\x40\x37\xb5\x9c\x5d\xf8\x7c\xfb\x7c\xf5\x7e\xb8\x1b\xb3\xb6\x27\xd5\x0c\x56\xb6\xab\xac\xac\xb1\xf6\x2b\x55\xfd\xb5\xa6\xfe\xda\xfa\x47\x00\x00\x00\xff\xff\xe2\x85\x8d\xbd\x7f\x15\x00\x00")

func registrationHtmlBytes() ([]byte, error) {
	return bindataRead(
		_registrationHtml,
		"registration.html",
	)
}

func registrationHtml() (*asset, error) {
	bytes, err := registrationHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "registration.html", size: 5503, mode: os.FileMode(420), modTime: time.Unix(1457983475, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _loginHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x57\x69\x6f\xdc\x36\x13\xfe\xfc\xee\xaf\xe0\xcb\x16\x50\x0c\x58\xd2\xfa\xde\x38\x92\x0a\x23\x69\x81\x00\x2d\x62\x14\x29\x8a\xa0\x0d\x0a\xae\x34\x92\x98\x50\xa4\x4a\x52\x6b\x2f\xd2\xfc\xf7\x0e\x75\xec\x25\x39\x4e\x8a\x44\x80\x61\x92\xf3\xcc\xf1\xcc\x0c\x8f\x8d\xfe\xff\xe2\xd5\xf3\xd7\x6f\x6e\x7f\x24\xa5\xad\x44\x32\x8b\xdc\x3f\x22\x98\x2c\x62\x0a\x92\x26\x33\x42\xa2\x12\x58\x96\xcc\xfe\x47\x70\x58\x81\x65\x88\xb4\xb5\x0f\x7f\x37\x7c\x15\xd3\xe7\x4a\x5a\x90\xd6\x7f\xbd\xae\x81\x92\xb4\x9b\xc5\xd4\xc2\xbd\x0d\x9d\xa9\x67\x24\x2d\x99\x36\x60\xe3\xc6\xe6\xfe\x82\x92\xd0\x99\x6c\xbf\xc8\x72\x2b\x20\xf9\x59\x15\x84\x4b\xe2\x93\x97\xd6\x33\xe4\x8d\x6a\xc8\x2b\x29\xb8\x84\x28\xec\xe4\xb3\x0d\x1e\x57\xdf\x13\x0d\x22\xf6\x8c\x5d\x0b\x30\x25\x80\x25\xb5\x86\x1c\x6c\x5a\x7a\xa4\xc4\x51\xec\xb9\xe0\xcc\x75\x18\xb2\x77\xec\x3e\x28\x94\x2a\x04\xb0\x9a\x9b\x20\x55\x55\xbb\x16\x0a\xbe\x34\x21\xf2\x6b\x04\xd3\x7f\x55\xcc\x82\xe6\x4c\x84\x27\xc1\x3c\x98\x0f\xcb\xfe\xb0\x1c\x54\x5c\x06\xa9\x31\x5e\x32\x0e\x82\x6e\x83\xa0\x9d\x6f\xca\x0c\x12\x35\x21\x2a\x84\xad\xd0\xa9\xd2\xf0\x8b\x75\x73\xcc\xa2\xcf\xee\xc0\xa8\x0a\x86\x08\x46\x66\xf6\xd9\x3a\x15\x73\x48\x17\xd5\x7e\xc8\x59\xc5\xc5\x3a\xfe\x1d\x84\xc8\x05\xba\xfb\xe7\x46\xaf\xd4\xf5\xf9\x7c\x7e\x7c\x85\x7f\xf8\x9f\x5b\x26\x78\xea\x66\xdd\xc8\x3b\xcc\xb0\x47\x2c\xd6\x36\xf6\xda\x92\xee\xe5\x82\x44\x2d\x6a\x5b\xa1\xf6\xfb\x43\x16\x7f\x5e\xa7\x42\xb1\xf7\x6f\x8f\xdd\xc4\x1f\xc6\xc1\x30\x26\x1f\xf6\x14\xdc\x97\x71\x53\x0b\xb6\xbe\x26\x52\x49\x78\xb6\x27\xfe\xb8\x75\x17\xf6\xfe\x70\xd4\xf5\xe4\x2c\x5a\xaa\x6c\x4d\xd0\x30\xab\xeb\x98\x0a\x55\x70\x79\x53\xd7\xd4\x49\x1c\x02\x34\xe1\x59\x4c\xbb\x61\xd7\xcc\x19\x5f\x91\x54\x60\xb6\x63\xda\x77\x2b\xed\x08\xb5\x12\x87\x46\x2b\x8a\x26\x11\xeb\x0b\x13\x84\x34\xc1\xde\x34\xd8\x99\x41\xd7\x99\xe8\x9e\x25\x51\x88\xf8\x1d\xcd\xde\xa6\xe1\x85\xf4\xef\x34\xab\xe9\x90\xa6\x8d\x21\x0d\x05\x37\xd6\xc5\x11\xf1\x01\x9e\x33\x92\x33\xbf\x31\xa0\xfd\x5a\x34\x06\x45\x21\x4f\x7e\xed\x81\xce\x4d\xe7\x60\xf0\xd5\x0f\x3a\xfa\xa0\x91\xe6\x6c\xd7\x79\xde\x08\x61\x52\x0d\x20\x73\xa5\xab\x81\x97\x1b\x93\x4d\xee\x71\x07\x97\x0a\x59\xd6\xca\x60\xef\xb1\xd4\x72\x25\xfb\xcc\x51\xdc\xf7\x6b\xd5\x58\x97\x19\xd1\x54\x38\x97\xac\x82\x5e\xd8\x5a\x24\x3b\x84\xd1\x22\x97\xdc\xc6\xde\x46\x1c\x20\x8d\x15\x68\xc3\x33\x58\x61\x2b\x65\xcc\xd9\x06\xad\x95\x36\x24\x26\x1f\x28\x97\xed\xb2\x6b\x6b\x5b\x72\x59\xd0\x6b\x62\x75\x03\x1f\xbd\x51\x32\xb1\x55\xef\x37\xc1\x68\x75\x47\x77\x7a\x6e\x90\xef\x2a\x6d\x04\x87\x04\x1c\xd0\x2f\xac\x7f\x8f\xd9\x39\x9f\xf7\x73\x37\x39\x99\xcf\x77\x8c\xb6\xfa\x55\x86\x84\xea\xc6\xfa\xae\x2f\x18\xd6\x59\x27\xa3\x3e\x8d\x04\x5b\x82\x48\x7e\x43\xa6\x2e\x37\x51\xd8\xcd\xc7\xb8\xd6\x92\xcb\x51\xa5\x32\xb7\xe1\xfb\x04\xe3\x6e\x16\x20\x0b\x5b\xc6\xf4\x94\xe2\x4e\xc3\x43\x54\x43\xb6\x9b\x67\xda\x6d\xb7\xf6\x04\xc5\xfa\x34\x56\xe5\x2a\x6d\xcc\x41\xac\xe1\x54\xb0\xff\x99\xce\x2d\x76\xcf\x9d\xd2\xd9\xe7\xd3\xa9\x7b\x8d\x11\x87\xad\xa0\xa3\xb1\x99\x7f\x53\x02\xa7\xfe\x4f\xd8\xc9\x4a\xbb\x7c\x95\xb8\xa5\x79\xda\xf6\x1e\xde\x47\xd9\x17\xd4\xc8\x2a\x5b\x3b\x0d\x2c\x53\x86\x57\xc0\xfd\x50\xa9\xcb\xbd\xba\x5d\x8e\x38\x6f\xf4\x3e\x87\xe3\x56\x3a\xee\xde\xc3\xb6\xee\x87\xd3\x2d\xbe\xbf\x2d\x5a\xb1\xa9\x99\x1c\x8c\xb8\xf1\x81\x78\x62\x83\x4c\x64\x05\x83\x5e\x36\xd6\x62\xf6\xba\x0a\x9a\x66\x59\x71\x6c\xc5\xfe\x8c\x41\xb1\x66\xdc\x20\x79\x1c\xd5\x9a\x57\x4c\xaf\x69\x7f\x89\xb7\x8c\x3b\xe5\x09\xc3\xfd\xa1\x51\x81\x31\xac\x00\xb3\x73\xae\x7c\xe2\xe0\x68\x6b\xe1\x76\x81\x5f\xa2\xd4\x1d\x98\xc2\xc0\x26\x98\x16\x33\x41\x62\xc2\x5f\x3c\x1c\x3f\x78\x3e\x66\xae\x45\xd0\x10\x4d\x5e\x76\x6b\x64\x67\xf1\xa0\x30\x0f\xd4\xeb\xa1\xa5\x87\x4b\xb0\x5b\xcc\x28\x74\xb4\xdd\x41\xde\xad\xcd\x22\x3c\xb6\x79\x6d\x89\xd1\xe9\x17\x3e\x61\xde\x19\x7c\xbb\x9c\x07\x8b\x61\x1e\xbc\x33\xee\x3c\xed\xec\x25\x5f\xd1\xb0\xcf\x24\x77\xef\xa2\xf6\x45\xf2\xcd\x9c\x0c\xdd\xf1\x6d\xbd\x30\x7c\xdc\x7d\x55\x0f\x9f\xf5\x90\x9c\x74\xd6\x35\xc8\x50\x3c\x3c\x89\x1a\x01\x4f\xbc\xe1\x25\xe3\xe1\x13\xca\x93\xc5\x2f\xbd\x19\xef\xd8\x4d\xfa\x14\x79\x6f\x8f\x36\xcd\x85\x51\xc9\x9c\x17\x4f\xf2\x46\xb6\x37\xfa\x93\xef\xab\xec\x75\x09\xe8\xb5\xb8\xd5\x6a\x85\x5b\x47\x1f\x1d\x3c\xbb\xc6\x88\x20\x83\x1c\x0f\xa9\x5b\x26\xc0\x5a\x0c\x62\x29\x1a\xe0\xa6\xc4\x18\xc6\x0f\x36\xef\x62\xee\x5d\x13\xef\xbb\xfc\x2a\x5f\xe6\x99\x77\x3c\x06\xe0\x0d\xdb\x22\x96\xf8\x52\x81\x49\xc4\x69\x8f\x58\x40\xba\x80\xd3\x29\xc4\x59\x8f\xb8\xb8\x60\x59\x76\x36\x85\x38\xef\x11\x67\x29\x3b\x49\x27\xbd\x5c\x0c\x88\xf9\xd3\x93\xe5\x72\x0a\x71\xd9\x23\x4e\xd9\x15\xb0\x49\x2f\x57\x03\xe2\xec\x72\xb9\x60\x53\x88\x45\x8f\x38\xc9\x2e\x16\x57\x93\x5c\x9e\x0e\x88\xab\x73\xa4\x33\x85\xb8\x19\x52\xf6\x70\x52\x6f\x4e\x1f\xcd\xea\xcd\xe3\x29\xb9\x79\x9c\x8f\xbb\xb1\x34\x33\xf6\x05\xe4\xac\x11\xf6\xb9\x12\x4a\x3b\x15\xc1\x8b\xd2\x7e\x52\x81\xe9\xf7\x2d\xda\x38\xf8\xc5\x9c\x20\x29\x82\x51\x13\xac\x25\xc1\xd0\x88\x63\x49\x1c\x0f\xd2\x46\xba\xff\xd6\x3f\x7a\xf6\x58\x8f\xe2\xfd\x5e\x61\x6f\x66\x5d\x5c\xde\xd1\x28\x92\xa0\xbf\x94\x46\x6d\xbc\x85\x3a\x37\x3b\x1b\x31\x74\x3f\x22\xda\x57\x75\xfb\x13\xf8\xdf\x00\x00\x00\xff\xff\x7e\x6e\x89\xef\x13\x0f\x00\x00")

func loginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_loginHtml,
		"login.html",
	)
}

func loginHtml() (*asset, error) {
	bytes, err := loginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "login.html", size: 3859, mode: os.FileMode(420), modTime: time.Unix(1457983445, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _homeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x58\x6b\x6f\xdb\x36\x17\xfe\xfc\xfa\x57\xf0\x55\x0b\x28\x01\x22\xc9\xb9\x35\x6e\x62\x6b\x30\xd2\x7d\xe8\x87\x2e\xc5\xd6\x61\x28\xb6\x62\xa0\xa5\x23\x99\x0d\x45\xaa\x24\x65\xc7\xdb\xfa\xdf\x77\xa8\x8b\x2d\x5b\x72\x93\x34\x2d\x30\x01\x81\x49\xf1\x39\x97\xe7\x5c\x48\x2a\xe3\xff\xbf\xba\xb9\x7e\xf7\xfe\xed\x8f\x64\x6e\x32\x1e\x0e\xc6\xf6\x87\x70\x2a\xd2\x89\x03\xc2\x09\x07\x84\x8c\xe7\x40\xe3\x70\xf0\x3f\x82\xc3\x0c\x0c\x45\xa4\xc9\x3d\xf8\x54\xb0\xc5\xc4\xb9\x96\xc2\x80\x30\xde\xbb\x55\x0e\x0e\x89\xaa\xd9\xc4\x31\x70\x67\x02\xab\xea\x8a\x44\x73\xaa\x34\x98\x49\x61\x12\x6f\xe4\x90\xc0\xaa\x2c\x9f\xb1\x61\x86\x43\xf8\xda\xb8\x9a\xbc\x97\x05\xb9\x11\x9c\x09\x18\x07\xd5\xeb\xc1\x1a\x86\x6f\x6f\x89\x02\x3e\x71\xb5\x59\x71\xd0\x73\x00\x43\x72\x05\x09\x98\x68\xee\x92\x39\x8e\x26\xae\xf5\x49\x5f\x06\x01\xfd\x48\xef\xfc\x54\xca\x94\x03\xcd\x99\xf6\x23\x99\x95\xef\x02\xce\x66\x3a\x40\x5a\x05\xa7\xea\xcf\x8c\x1a\x50\x8c\xf2\xe0\xd8\x1f\xfa\xc3\xe6\xb5\xd7\xbc\xf6\x33\x26\xfc\x48\x6b\x37\xec\x3a\xe1\x6c\x9c\x70\x2a\xdb\x0e\xd5\xc8\x4f\x07\x28\x10\x94\x8b\x56\xd4\x09\x1e\x2d\x9b\x60\xf0\x3c\xba\x04\x2d\x33\x68\x3c\xe8\xa8\xd9\x66\x6b\x45\xf4\x2e\x5d\x14\xfb\x21\xa1\x19\xe3\xab\xc9\x6f\xc0\x79\xc2\xd1\xdc\x3f\x53\xb5\x90\x97\x67\xc3\xe1\xd1\x05\xfe\xe1\x2f\x33\x94\xb3\xc8\xce\xaa\x91\xbb\x1b\x61\x97\x18\x4c\xe9\xc4\x2d\x33\xb9\x15\x0b\x32\x2e\x51\x9b\x0c\x95\xcf\xef\x22\xfd\xe3\x32\xe2\x92\xde\x7e\x38\xb2\x13\xaf\x19\xfb\xcd\x98\xfc\xbd\x25\x60\x9f\x98\xe9\x9c\xd3\xd5\x25\x11\x52\xc0\xd5\xd6\xf2\xe7\x8d\xb9\xa0\xb6\x87\xa3\xaa\x14\x07\xe3\x99\x8c\x57\x04\x15\xd3\x3c\x9f\x38\xcc\xe8\x95\x2c\x64\x59\x3e\xd3\x3c\x77\x2c\xc0\x02\x41\x11\x16\x4f\x9c\x6a\x58\x95\x72\xcc\x16\x24\xe2\x18\xf4\x89\x53\xd7\xaa\x53\xf1\x2a\x57\x2c\x9a\xcb\x54\x3a\xe1\x98\xd6\xf9\x79\xe6\x84\xe4\xb5\xd1\x58\x9f\x7e\x55\x9f\xe8\x04\x0d\xc7\x01\xc2\xeb\x00\xb4\x75\x6a\x96\x0a\x6f\xa9\x68\xee\xb4\xa2\xd5\xa8\x12\xd2\xb0\x84\x45\xd4\x30\x29\x34\x9a\x60\x8d\x54\x42\x49\x42\xbd\x19\xe6\x0a\xdf\x06\x2c\xdc\x8a\xc3\x38\x8b\x3d\x23\x25\x37\x2c\x27\x4d\x24\xb7\x11\xb6\x77\xe6\x74\x01\x18\x44\x22\x73\x10\x98\xc9\x4f\x05\x68\xa3\xb7\xf5\x04\x1b\x45\x2d\xe7\x90\xcc\xa0\x6d\x2a\x03\x51\x10\xfc\x95\x49\x62\x9b\xd6\xf1\xce\xc9\xe9\xd0\xd9\x71\xa9\x61\xe4\x54\x1e\xb1\xe8\x76\xe2\x3c\xcf\xe2\x1b\x34\xfe\x06\xe5\x0f\x9e\xc3\x02\x23\x7b\xd8\xe5\x98\xf3\x42\x57\x1c\x77\x57\x22\xaa\xc0\x78\xb1\x5c\x8a\x9e\x18\x3c\x24\x0a\xf6\xb9\x56\x80\x1d\x4c\x04\x2c\x7d\xdf\xdf\xd5\xd0\xcb\xbf\x13\x83\x76\x1c\xbc\xba\x44\xf6\x5b\x5c\x23\x99\x81\x2c\xb4\xb3\x59\x61\x8c\x14\x4d\xf1\xa0\x27\x52\xa5\x54\xb0\xbf\xca\xac\x3b\xe1\x4f\xb0\x24\xed\x37\xa5\x57\x95\x4c\x58\x8e\x37\xda\xbe\xc2\x16\x76\x7f\x4e\xc5\xaa\x32\x53\x4f\x1e\x6c\x61\xb3\x58\xd3\x6e\x57\x49\xbd\xb4\xee\xf9\x75\x09\xd8\x76\x29\x4c\x37\xd3\x65\x27\x54\x2b\x98\xcd\x5f\x70\x46\x70\xb6\x0e\x75\xdd\x40\xeb\x41\xd5\xd9\xa0\xd0\xc0\xa0\xdd\x50\x49\xc1\xb9\x8e\x14\x80\x48\xa4\xca\xda\xbd\x8a\xdb\xda\x1d\x9e\x53\xd8\xf8\x58\xa5\x4a\x2e\xdb\x1d\xd7\xac\x87\x6b\x33\xad\x85\x46\x26\x92\xbc\xc8\x84\x53\x02\xbd\xd4\x78\x77\x68\x6e\x34\xac\xe7\x76\x72\x3c\xec\x14\xbe\x95\xc7\x62\x58\x30\x58\x76\x74\xf7\x98\xda\xf5\x61\x4d\xb6\xda\x3d\xd6\x83\x31\x52\x64\xb9\x21\x5a\x45\x8f\x3c\xc9\x3e\x6a\x3c\xc2\xce\xfc\x51\x33\xf7\x3f\xe2\x46\x8d\x1b\x66\xa9\x2f\xdc\x52\xec\x3c\x49\xb1\xa7\x30\x68\x80\xea\x9d\x3d\xea\x9f\xe6\x37\x56\x97\xd6\x34\x05\x5d\x1e\x7c\xfb\x49\x3c\xd1\x0a\xf6\x9d\x3d\xe3\xbf\xb3\x11\xbc\x42\x7c\x53\x0b\x0f\xba\xae\xf4\x1a\xab\xca\xae\xa9\x8d\x4c\xc6\x05\x87\x03\x77\xe7\xbc\x74\xf1\xbc\x76\x45\xfa\xa6\xd6\xe6\x1e\xe1\xe4\x67\x9b\x6d\x5c\xb0\xef\xeb\xcc\xb8\x1f\x0e\xd7\xd5\x8d\x7e\x8a\x84\xa5\x07\x49\x21\x22\xbb\x8d\x1d\xe0\xee\xff\x6e\x0e\xe8\x47\xfa\x56\xc9\x05\xc3\x4e\x3e\xdc\x39\xee\xbb\x08\x3f\x86\x04\x3d\x78\x4b\x39\x18\x83\x6e\xcd\x78\x01\x4c\xcf\xd1\x6a\xf7\xa2\xe0\x9e\x0f\xdd\x4b\xe2\x3e\x4b\x2e\x92\x59\x12\xbb\x47\x5d\x00\x76\x6b\x89\x98\xe1\x36\x02\xbd\x88\x93\x1a\x31\x82\x68\x04\x27\x7d\x88\xd3\x1a\x71\x7e\x4e\xe3\xf8\xb4\x0f\x71\x56\x23\x4e\x23\x7a\x1c\xf5\x5a\x39\x6f\x10\xc3\x97\xc7\xb3\x59\x1f\xe2\x45\x8d\x38\xa1\x17\x40\x7b\xad\x5c\x34\x88\xd3\x17\xb3\x11\xed\x43\x8c\x6a\xc4\x71\x7c\x3e\xba\xe8\xe5\xf2\xb2\x41\x5c\x9c\x21\x9d\x3e\xc4\xb4\x09\xd9\xfe\xa0\x4e\x4f\xee\x8d\xea\xf4\xfe\x90\x4c\xef\xe7\x63\x8f\x1c\x45\xb5\x79\x05\x09\x2d\xb8\xb9\x96\x5c\x2a\x2b\xc2\x59\x3a\x37\x5f\x14\xa0\xea\xb6\x44\x6b\x0b\x3f\x1f\x12\x24\x45\xd0\x6b\xbc\xb5\x0c\x09\xba\x46\x2c\x4b\x62\x79\x90\xd2\xd3\xed\x3b\xe6\xe1\xd5\x7d\x35\x6a\x70\x8e\xb5\x19\x57\x7e\xb9\x87\x1d\x4f\xfc\x5c\xe1\x9e\xa2\x56\x9d\x32\xde\x40\x3f\x7f\xa9\x71\xca\x7d\x75\x6f\xd3\x6c\xad\x76\x6d\x2f\xe7\x20\x0e\xdc\x60\xe7\x9a\xd1\xdb\x41\xf6\xc1\x43\x1f\xaf\xdb\x06\x7e\x55\xfc\xd2\x86\x30\xcb\xf1\xda\x8d\x5f\x10\x41\x5b\x3a\xb0\xa7\x9b\xb6\x3a\x7d\xfb\x01\xd7\x13\x7c\xfb\x94\xf1\x97\x9c\x83\x42\x4d\x37\x2d\xf1\xeb\xf5\xc2\xbd\x92\x53\x8d\xb2\x8b\xcc\xed\xc0\x3e\xf7\x84\xb9\xa6\xda\xf6\x54\x07\x97\x29\x97\x33\xfc\x7c\x89\xbf\x09\xe5\x18\x3f\x6e\x19\xff\x2f\xb2\xde\xdc\xed\x1e\x4f\xb4\x16\xfc\x8a\xb4\x5e\x57\x92\xdf\x99\x5b\xe5\x1f\x83\x27\x65\x73\x9b\xe4\x23\x13\xf9\xfd\x78\x5e\xb5\xf6\x80\xab\xc1\x9e\xbb\xc0\xbe\xa2\xdc\x98\xfb\xd2\x4d\x62\x9f\xb4\x06\xb5\x60\x11\x3c\x50\xb4\x89\xdf\xa3\x6d\x36\x82\x7b\xcc\x05\xf6\x43\xbd\xbc\xde\x97\xff\x5d\xfa\x37\x00\x00\xff\xff\x60\xfe\xdf\xb8\x6e\x12\x00\x00")

func homeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_homeHtml,
		"home.html",
	)
}

func homeHtml() (*asset, error) {
	bytes, err := homeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "home.html", size: 4718, mode: os.FileMode(420), modTime: time.Unix(1457983036, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _errorHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x54\xdf\x6b\xdb\x30\x10\x7e\x1f\xec\x7f\xb8\x6a\x0f\x7e\xa9\xa3\xac\xdb\xe8\xe8\xec\x8c\xd2\x96\x11\x18\xb4\x2c\x81\xd1\xa7\xa2\xd8\x17\x5b\x54\x96\x5c\xeb\xdc\x2e\xb0\x3f\x7e\xe7\x5f\x6d\x9c\x74\xb4\x83\x3d\x18\x9d\x74\xdf\x7d\xf7\xf9\x93\xcf\xd1\xc1\xf9\xe5\xd9\xf2\xfa\xea\x02\x72\x2a\xcc\xec\xed\x9b\xa8\x59\xc1\x28\x9b\xc5\x02\xad\x68\x4f\x50\xa5\xbc\x02\x44\x05\x92\x62\x20\x95\x21\xde\xd5\xfa\x3e\x16\x67\xce\x12\x5a\x0a\x97\x9b\x12\x05\x24\xdd\x2e\x16\x84\xbf\x48\x36\x44\x5f\x20\xc9\x55\xe5\x91\xe2\x9a\xd6\xe1\x67\x01\xb2\x23\x22\x4d\x06\x67\x73\x0a\x3c\x5c\xbb\x1a\x2e\xad\xd1\x16\x23\xd9\x1d\xb7\x08\x3e\xb8\x85\x0a\x4d\x2c\x3c\x6d\x0c\xfa\x1c\x91\x04\xe4\x15\xae\x63\xa1\x3c\x33\x7a\x99\x78\x2f\xdb\xe4\x84\x23\x21\xff\xa9\x6e\xcd\x52\x43\xf5\x80\xde\x15\x38\x29\xb4\xdd\xa3\x68\x2b\x82\xe6\x5d\xfd\x89\x6c\xe1\x7e\x92\x39\x97\x19\x54\xa5\xf6\x93\xc4\x15\x0d\xcf\xd7\xb5\x2a\xb4\xd9\xc4\x3f\xd1\x98\xb5\xe1\x56\xbf\x4f\xab\x7b\x77\xf2\x71\x3a\x3d\x3c\xe6\x87\x57\x4d\xca\xe8\xa4\xd9\x75\x51\xd0\x8a\x0b\x9e\xc4\x05\x40\x6c\x5e\x1c\xb4\x9e\x31\x65\xd0\x58\x2e\x7b\xcf\xa3\x95\x4b\x37\xc3\x1d\x60\x05\x3a\x8d\x45\x17\x8a\x4e\x6b\xaa\xef\x21\x31\xfc\x62\xb1\xe8\xdd\xef\x12\x7d\xaa\xc1\x1b\x97\x39\x31\x8b\x54\x6f\xc2\x44\x8a\x19\xcc\xc9\xb3\xef\x93\xce\x77\x88\xa4\x9a\x45\x92\xf1\x1d\x67\x1f\x75\x22\xb0\xe2\x90\x37\x07\x61\x08\x8b\xef\xf3\xf3\x8b\x05\x2c\x96\xa7\x3f\x96\x10\x86\x0d\x66\x68\xe2\x8d\x4e\xf1\xfd\xcb\xa2\xf2\xa3\xd9\xa7\xe9\x94\xa9\x8f\xb6\x4e\x2e\x5d\xe9\x0f\x81\x72\xed\xe1\x41\x79\xb0\x8e\xc0\xd7\x65\xe9\x3c\xa6\x40\x0e\x72\x55\x96\x68\x1f\x6b\x9e\xf4\x3d\x2a\x1e\xa9\x38\x1e\x39\xf0\xbc\x8e\x71\xf2\xae\x76\x84\xfe\xa6\xc1\x28\xf6\xa3\x7a\x42\x8d\x71\xab\x9a\xc8\xd9\x30\x61\x22\xbe\x8b\xe6\xc2\xfa\x78\xbb\x60\x5c\x92\xba\x64\x9c\xe4\xb4\x02\x52\x55\xc6\x33\x21\x6e\x56\x3c\x68\xb7\xc3\xe7\xf9\x6e\x17\x09\x70\xee\x92\xba\xe0\x16\x8a\xb4\xb3\x3b\x34\x7c\x69\xa3\xae\x8f\x76\xfc\x45\x07\xf0\x73\xf4\x6a\x31\xfc\x95\xa7\xdb\xcd\xf7\xb5\x9d\x5e\xcd\xff\xbb\xbe\x0f\xaf\xd6\x37\x4c\x66\xa6\x29\xaf\x57\xed\x3c\x6a\xf2\x1b\x57\xbb\xf6\x9b\xde\x97\x1b\xf9\x52\xd9\xa1\x5d\x57\x15\x36\x93\xbe\x8f\x04\x58\xb8\xba\x4a\x90\xff\x67\x29\x82\xb3\xf0\xad\x45\xef\x11\xca\x86\x71\x57\xef\x0b\x6f\x3d\xde\x6f\xef\x9e\x1b\xc0\x61\x1d\xfe\x02\xb2\xff\x49\xff\x09\x00\x00\xff\xff\x1c\x94\x6c\xce\xb6\x05\x00\x00")

func errorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_errorHtml,
		"error.html",
	)
}

func errorHtml() (*asset, error) {
	bytes, err := errorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "error.html", size: 1462, mode: os.FileMode(420), modTime: time.Unix(1457983046, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _apidocumentationHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x56\x51\x6f\xdb\x36\x10\x7e\x9e\x7f\xc5\x95\xdb\xe0\x16\x88\x24\xb7\x48\xd1\xcc\x96\x34\x14\xd9\x1e\xf2\xb2\xee\xa1\xc0\x50\x6c\xc3\x40\x4b\x27\x89\x0b\x45\x6a\x24\x15\x47\xdd\xf6\xdf\x77\xa4\x62\x47\x8a\x9d\xa2\xe8\x4b\x57\x01\x89\xa8\xe3\xdd\x77\xdf\x1d\xef\x8e\x4e\x9f\xfc\xf0\xe6\xf2\xed\xbb\x9f\x7f\x84\xc6\xb5\x32\x5f\xa4\xfe\x05\x92\xab\x3a\x63\xa8\x58\xbe\x00\x48\x1b\xe4\x65\xbe\xf8\x0a\x68\xd9\xa2\xe3\xa4\xe9\xba\x08\xff\xea\xc5\x4d\xc6\x2e\xb5\x72\xa8\x5c\xf4\x76\xe8\x90\x41\x31\x7e\x65\xcc\xe1\xad\x4b\x3c\xd4\x06\x8a\x86\x1b\x8b\x2e\xeb\x5d\x15\x5d\x30\x48\x3c\x64\x78\x52\x27\x9c\xc4\xfc\xca\x2d\x2d\xbc\xd3\x3d\xbc\x51\x52\x28\x4c\x93\x51\xbc\x38\xa8\x91\xf4\x1a\x0c\xca\x8c\x59\x37\x48\xb4\x0d\xa2\x63\xd0\x18\xac\x32\xc6\x2d\x41\xdb\xa4\xb0\x36\x09\x9b\x31\xad\xd8\xc4\xc5\x47\xda\x56\xc4\x3b\xe2\x3b\xb4\xba\xc5\xb8\x15\xea\x24\x4c\xb0\x5a\xfa\xe0\xed\x3a\x09\x26\x36\xae\xb5\xae\x25\xf2\x4e\xd8\xb8\xd0\xad\xc7\xfa\xbe\xe2\xad\x90\x43\xf6\x0b\x4a\x59\x49\x72\xf7\xcf\x6b\x73\xa3\xd7\xe7\xab\xd5\xd9\x2b\xfa\xa3\xb7\x70\x5c\x8a\xc2\x7f\x8d\xab\x65\x20\xb8\xbc\x27\xb8\x04\x47\xd9\xcc\x96\x21\x89\x04\xb9\x3c\x10\x81\x34\x68\xdd\x27\x27\x3c\xbf\xaa\xfa\xb7\x75\x21\x35\xbf\xfe\xfd\xcc\x7f\x44\xfb\x75\xbc\x5f\xc3\xdf\x33\x03\xff\x94\xc2\x76\x92\x0f\x6b\x50\x5a\xe1\x66\xb6\xfd\xef\x1c\x3e\xae\x7a\x29\x6d\x61\x10\x55\xa9\x8b\xbe\xa5\x03\xe6\x4e\x68\x75\x04\x2a\xb1\x72\x6b\x58\x6d\x1e\x88\x8d\xa8\x9b\x53\xf2\xf7\x91\x50\x25\xde\x1e\xed\xb4\xdc\xd4\x42\x45\x4e\x77\x6b\x78\xbe\xea\x6e\x1f\x92\x9b\x71\x13\xce\x0e\xba\xd7\xa1\x74\x7c\x9d\xa2\x39\x11\xeb\x4e\x94\xae\xf1\x60\xab\x6f\x37\xb0\xe5\xc5\x75\x6d\x74\xaf\xca\x35\x98\x7a\xcb\x9f\x9e\x5f\x9c\xc1\xf3\xf3\x97\xf4\xef\xe2\xd5\x19\xac\xe2\xef\x5e\x3c\xdb\x40\x83\x23\xeb\x0b\x4f\x00\x36\x07\xae\x2f\x36\x33\x7a\x2f\x3f\x4c\xef\x6b\xc3\x5b\xf9\xc7\x87\xb3\xf6\xe0\x28\xe0\x31\xbc\x34\x79\x78\xf6\x93\xb2\x64\x54\x82\x33\x37\xa3\xb2\x4d\x48\x1e\x51\x57\x5a\x2d\x31\x92\x3e\xa4\xc8\x35\xd8\x8e\x7d\x72\xdc\x18\x85\xa4\xa6\xa0\xee\xf5\x2a\xbe\xf7\xd3\x64\xec\xfd\x45\xba\xd5\xe5\x00\x54\x4e\xbc\xeb\x32\xe6\xa3\xba\x1c\x41\x5f\x77\xdd\xc1\xcc\x8b\x0f\xce\xbc\x01\x03\x6f\x5a\x8a\x9b\xbd\xc6\xf1\x69\x8d\x13\x66\xa2\x72\x37\x42\xd8\x58\xf3\x61\x47\x94\x19\x93\xba\xd6\x2c\x4f\xf9\x5d\xb8\x71\xc2\x72\xb8\x72\x96\xe6\x46\x3c\xce\x0d\xe2\xca\xf3\x34\x21\xfd\x80\x38\x2e\xee\x5e\x73\x12\x8f\x94\xf3\xc4\xe3\xa4\xdf\xa6\x31\x81\x35\x45\xc6\x92\xa3\x54\x7b\x9d\xa4\xb7\x68\x6c\xec\x97\xcc\x1f\x28\xdf\x52\x0e\x42\x1e\x23\xbb\x13\xae\x68\xa8\x30\xf7\xe2\x11\x53\x0a\x3f\x37\x6b\x54\x68\xb8\xd3\xf7\xbb\x61\xfc\xd1\xd1\x58\xdd\x9b\x02\x2d\x39\x97\x92\x77\x16\x4b\x8a\x6e\xca\x66\x42\xf2\x49\x14\xc1\x4f\x9a\xe2\xeb\x11\x76\xcd\x00\xdb\xde\x01\xf9\x06\xaf\x0e\x7b\xf2\x94\xf5\x81\x60\xa9\x8c\x8d\x85\xce\xe8\x0e\x0d\x09\x44\x05\x57\x40\x23\xa2\x0c\x06\x21\x08\xa0\xe8\xc0\xed\x44\x81\x10\x45\x13\x2f\x3e\x85\xa1\x58\x32\x36\x1f\x1e\x2c\x9f\x95\xec\x97\x91\xb3\xe4\x13\xcf\x59\x9b\x9a\x2b\xf1\x3e\x48\x3e\x17\xf7\x8f\x26\x4b\x77\x52\x47\x6c\xf1\x0b\x20\xaa\x9c\xe1\x85\xfb\xdf\x13\xf5\x25\xfb\x19\x4b\xe0\x30\xd1\xee\x27\x1b\x8d\x32\xd1\xb9\xc7\x48\x8f\xbb\xf3\x7b\xe0\xc6\x4f\x3e\x13\xff\x69\x69\xa0\xde\x29\xe4\x9f\x80\x73\x1a\xc0\x27\xf5\x9b\x58\x69\xba\x1f\x2a\xfa\x71\xe3\x9e\x3e\xdb\x2c\x26\x4a\x89\xbf\x17\xfc\x7b\xfc\xb9\xf9\x5f\x00\x00\x00\xff\xff\x8e\x41\x17\xc0\x7f\x0a\x00\x00")

func apidocumentationHtmlBytes() ([]byte, error) {
	return bindataRead(
		_apidocumentationHtml,
		"apidocumentation.html",
	)
}

func apidocumentationHtml() (*asset, error) {
	bytes, err := apidocumentationHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "apidocumentation.html", size: 2687, mode: os.FileMode(420), modTime: time.Unix(1457983065, 0)}
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
	"index.html": indexHtml,
	"registration.html": registrationHtml,
	"login.html": loginHtml,
	"home.html": homeHtml,
	"error.html": errorHtml,
	"apidocumentation.html": apidocumentationHtml,
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
	"apidocumentation.html": &bintree{apidocumentationHtml, map[string]*bintree{}},
	"error.html": &bintree{errorHtml, map[string]*bintree{}},
	"home.html": &bintree{homeHtml, map[string]*bintree{}},
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
	"login.html": &bintree{loginHtml, map[string]*bintree{}},
	"registration.html": &bintree{registrationHtml, map[string]*bintree{}},
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

