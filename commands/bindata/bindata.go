// Code generated by go-bindata.
// sources:
// ../templates/client_helper_resource.tmpl
// ../templates/client_resource.tmpl
// ../templates/generic_main.tmpl
// ../templates/python_client.tmpl
// ../templates/python_client_utils.tmpl
// ../templates/server_main.tmpl
// ../templates/server_resources_api.tmpl
// ../templates/server_resources_interface.tmpl
// ../templates/struct.tmpl
// DO NOT EDIT!

package bindata

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

var _TemplatesClient_helper_resourceTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x55\x51\x6f\xd3\x30\x10\x7e\xae\x7f\xc5\x11\x09\x64\xb7\x21\x05\xed\x69\x93\x2a\x24\x36\x4d\x08\x89\x01\xed\x10\x0f\xd3\xb4\xb9\xcd\x75\x0d\x6b\xec\xd4\xb9\x14\xaa\xaa\xff\x1d\x9f\x93\x66\x6d\x19\xd3\x5e\x90\x78\xc9\xec\xf3\x7d\xbe\xef\xfb\x7c\xb7\xae\xd7\x29\x4e\x33\x83\x10\x4d\xe6\x19\x1a\xba\x99\xe1\xbc\x40\x77\xe3\xb0\xb4\x95\x9b\x60\x19\x6d\x36\xa2\xd0\x93\x7b\x7d\x87\x50\xa7\x08\x91\xe5\x85\x75\x04\x52\x74\xa2\xf1\x8a\x7c\x8e\x5f\xa0\x99\xd8\x34\x33\x77\xfd\x1f\xa5\x35\x1c\xc8\x2c\x7f\x0d\x52\x7f\x46\x54\xf0\x7a\x9a\x13\xff\xa1\x2c\xc7\x48\x28\x21\xa6\x95\x99\x40\xc0\xe1\x7b\x9b\xae\x64\xaa\x49\x43\x66\x08\xdd\x54\x4f\x70\xbd\x51\x20\x33\x9b\x0c\x51\xa7\xe8\x62\x40\xe7\xac\x53\xb0\x16\x9d\x71\xd8\xc0\xc9\x00\xb8\x56\xf2\x49\xbb\x72\xa6\xe7\x01\xae\x44\x27\x9b\x86\xd3\x17\x03\x30\xd9\x9c\xd3\x3b\x0e\xa9\x72\x86\xb7\x01\x28\x3a\x1b\xb1\x8d\x05\xfa\xc9\x05\xfe\xac\xab\xc8\xb1\x8a\x39\x4f\x6c\x1a\x76\xa9\x1d\xe2\xe2\x7b\x46\xb3\x40\x30\x47\x9a\xd9\x34\x86\xca\xcd\x47\xe4\xa0\x24\xe7\x05\xc7\x70\xc8\x3b\x6e\x8c\x02\x16\x9e\x9c\x86\x75\x0c\xb3\x50\xa1\x84\x5c\x17\x57\x35\xf2\x7a\x0f\xb3\x28\xbf\x68\xa7\xf3\xe6\x56\xaf\xbd\x1b\xe0\x43\x2c\x0b\x6b\x4a\xdc\x33\xc0\x93\x69\x3d\x38\x30\xf0\xb9\x0e\x88\x4e\xbf\x0f\x13\x87\x9a\x10\x68\x86\xe0\x70\x51\x61\x49\xec\xcc\xa2\xbd\x3b\x30\x08\xee\x84\xc3\x43\x03\x7a\x5b\xd2\x31\x30\xa5\xe7\x97\x9e\x5a\x07\xf7\x31\x2c\xb9\x86\xd3\xc6\xb7\xd6\xd6\x1d\x86\x74\x98\x4c\xf2\x21\x44\x92\x11\x92\xf4\xa9\xbe\x77\x92\x51\xe1\x9d\xa1\xa9\x8c\x5e\x2e\xa3\x78\xa9\x54\x7d\x57\x53\xa0\xb6\x3c\x39\xb3\xd2\x83\x55\xfb\x80\xe3\x2a\x9b\xa7\x5f\x2b\x74\xab\x51\xf0\xb5\x6e\xb2\xc7\xdf\x40\x35\xde\xaf\x83\x8e\x39\x9a\xda\x50\x18\x0c\xe0\xcd\xae\x96\x28\xaa\x2b\x8f\x75\x89\xe1\x6a\x96\x11\xbd\x8b\x1e\xd3\x15\xca\x31\xf6\x21\xb9\x37\x80\x7b\x6f\x5d\x34\x88\xfc\xf7\x41\x97\x5c\x2a\x8e\xbe\x8a\xf6\x64\xb5\xb0\xab\x13\x26\xd4\x6e\xd5\xeb\xb7\xd7\x2c\xb2\xdf\x3f\xe3\x17\x74\x58\xf8\x89\xe5\x9e\x1b\x9e\x9f\x1e\x1d\x1d\x1f\x73\x61\x14\xb4\x2a\x10\x42\x02\x0f\x5d\x72\xe9\x3f\x0c\x69\x26\xe6\xe3\xe8\xf3\x05\xd8\xa5\x7f\x95\x2c\x45\xef\x49\x1b\xac\xad\x93\x04\x5d\xc6\x2a\xd8\xc9\x97\xbe\x31\xaf\xae\x79\x6c\x76\x1b\xb2\x21\x5b\x1f\xc8\xb6\x96\xec\x92\x4a\xce\xad\xcb\x35\xc9\xdb\xe8\xd6\xcb\x0b\x47\x81\xe2\xd1\xb1\xdf\xfa\xa0\x7a\x98\xb8\x96\xd8\x25\xfe\xa2\x3f\x88\x71\xf0\x2f\xc4\xf8\xe8\xdf\x12\xfb\x66\xf2\xc7\x3c\xab\xcc\x13\xae\xed\x61\xe4\xb8\x21\xa1\x6a\x76\x4c\x8e\xca\x76\xd0\x42\x79\x3f\x4a\x25\x32\x9f\xde\x2e\x9b\x9e\x0f\xc4\x4d\x67\xfa\xff\x4f\x4f\x0d\x59\x3b\x5f\x5d\x82\x41\x78\x77\x49\xa5\x12\x3b\x33\x78\xa0\x66\xdf\xe8\xca\x3c\x61\xf5\x1e\xe6\x3f\x52\x73\x40\xb3\x99\xf2\xed\x28\xef\xf4\xc0\xfe\xe3\x6f\xf3\xf8\x8a\xf5\x1a\x4d\xea\x7f\xe6\x7e\x07\x00\x00\xff\xff\x51\x79\xd6\x10\x0c\x07\x00\x00")

func TemplatesClient_helper_resourceTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesClient_helper_resourceTmpl,
		"../templates/client_helper_resource.tmpl",
	)
}

func TemplatesClient_helper_resourceTmpl() (*asset, error) {
	bytes, err := TemplatesClient_helper_resourceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/client_helper_resource.tmpl", size: 1804, mode: os.FileMode(420), modTime: time.Unix(1455625500, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesClient_resourceTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x55\x4d\x4f\xdb\x40\x10\x3d\x7b\x7f\xc5\xd4\x4a\x91\x4d\x2d\xe7\x5e\x29\x17\x20\xfd\x90\x28\xa2\x01\xda\x63\x65\xec\x31\x31\x24\x5e\x67\x77\x1d\x84\x5c\xff\xf7\xce\xec\xda\x8e\x09\x2d\x52\x91\xca\xa9\x39\x6d\xe6\xf3\xcd\x9b\xb7\xeb\xa6\xc9\x30\x2f\x4a\x04\x3f\x5d\x15\x58\x9a\x1f\x0a\xb5\xac\x55\x8a\x7e\xdb\x8a\x2a\x49\xef\x92\x1b\x04\xe7\x12\xa2\x58\x57\x52\x19\x08\x84\xe7\x97\x68\xa6\x4b\x63\x2a\x5f\x78\x9e\x8f\x65\x2a\xb3\xa2\xbc\x99\xde\x6a\x59\x5a\x4b\xbe\x36\xbe\x08\x85\x48\x65\xa9\x6d\x82\x92\xd2\x5c\x2d\x4e\x61\x06\x7e\xd3\xc4\x47\x89\xc6\xab\xc5\xe7\xb6\xb5\x41\x4d\x33\x49\xaa\xe2\x2c\x59\x23\xbc\x9f\x41\xcc\x07\x6a\x6e\x1e\x2a\x04\x8a\x75\x7f\x41\x1b\x55\xa7\x06\x1a\xe1\x39\x34\xc0\xdd\xe3\x63\x87\xac\x15\x22\xaf\xcb\x14\xce\xf0\x7e\xc8\x08\x42\x38\xdc\xa5\x73\x1e\x57\x2f\xf1\x3e\x18\xac\x21\x19\xe3\xae\xdc\x6c\x5c\xb0\x69\x09\x31\x9a\x5a\x95\x90\x72\xf1\xa6\x01\x95\x94\xc4\xc4\xe4\x2e\x82\xc9\xd6\xc2\xfc\x82\x66\x29\x33\x0d\x04\x75\xe4\xce\xd9\x9f\x73\xc0\x64\x1b\x7f\x20\x4c\xc7\x72\xbd\xa6\x8a\x1c\x37\x9d\xd2\x3c\xe4\x65\x38\x0d\x96\x19\x65\x32\xea\x20\x65\xa0\x3d\x05\x84\xca\x46\x75\xf5\xbb\x59\x46\x96\xf3\x44\x25\x6b\x8a\x72\xb6\x05\xea\xea\x48\x66\x0f\xb6\x66\x91\xd3\x7c\x30\xb2\x82\xef\x53\xdf\xa8\x6f\x77\x68\x47\x64\x27\xad\x05\x23\x54\x4a\xaa\x90\xa9\xd9\x68\x5b\x95\x61\x5f\xd7\xc5\x2a\xfb\x5a\xa3\x7a\xb8\x30\x8a\x76\x1a\x6c\xf8\x6c\xdd\x9a\xf8\xb2\x4d\x70\xc3\x4d\xbe\xa1\xba\x06\xff\xe3\xfc\x92\xa5\xe2\x79\x7f\x6c\x0f\xdb\x44\x41\x0d\x4f\xe1\x3a\x0a\x28\x95\x88\x49\x15\x26\x06\x41\x21\xb5\x23\xc5\xc8\xeb\x5b\x4c\x0d\xb9\xc8\x10\x01\x01\x65\x68\x16\x3d\xad\x78\xe1\x82\x02\xdb\x3b\x82\x5e\x5a\x8f\x01\x58\x11\x9f\x27\x66\xd9\x81\x78\xd7\x37\x1c\x80\x0c\x11\xd6\xdb\x51\x10\x41\x59\xac\x68\x4e\x8f\xc7\xa4\xb6\x6f\x66\x6c\x60\x8e\x9e\x9b\xb0\xd3\x4a\x6d\xb3\x2d\x5e\x17\x8f\x2b\x8d\x3b\xf7\x9e\xcf\x4e\xef\x79\x96\x81\x5c\x2a\x20\x69\x59\x65\x39\x29\x2d\x31\xc9\x50\x69\xea\x0c\xfd\x8f\xb8\x88\x3f\x59\x73\x7c\x81\x26\xa0\x78\xba\x66\xf1\x45\x45\x6b\x32\x79\xe0\xbf\xdd\x12\x19\xdb\x30\xb4\x09\x1d\xaf\x99\x04\xb3\x1c\x68\xb5\x7c\xea\x6a\x20\xb4\x17\x7f\x7c\x22\x03\x0a\x79\xe5\xb1\x3d\x8f\x9e\x1e\x54\xc0\x90\x62\xae\x4a\x97\x4f\x6a\x0c\x42\xf1\x9c\x9a\xb8\xca\xae\xaf\x1b\x87\x9f\x1d\x16\xc6\x09\xd2\x3b\x84\x2a\x18\x2a\x86\xb1\x33\x05\x07\x75\x28\x76\xc8\x46\x35\x5c\x01\x82\x28\x76\xe0\x44\x17\x08\x7b\x52\x3f\x99\x9f\xce\x2f\xe7\x4e\xed\x2f\x95\x6c\x57\x63\x50\xed\x3f\x17\xed\xd3\x25\xbc\xb6\xe0\xdc\x33\xba\x2f\x35\x31\xde\xc6\xcb\xdf\x0e\x07\x73\xa4\xe9\x4c\x12\xd7\xdf\x0b\xb3\xe4\xd0\xc0\xb7\x99\xbc\x3e\xfa\xd6\xfc\x2d\xe9\xbf\xe1\x3c\x7a\xfc\xca\x6c\x3a\xa4\x6d\x7b\xf0\x38\xc7\x39\x7e\xc2\xa5\x3c\x95\xf7\xa8\x38\xb1\x67\x20\xea\x99\x8e\xfa\xf5\xfd\xbf\x77\x1d\xb8\xfe\x60\xbf\xa8\x74\xe4\xb6\xa3\xe3\xaf\x00\x00\x00\xff\xff\x97\x0e\x7e\xf8\xac\x08\x00\x00")

func TemplatesClient_resourceTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesClient_resourceTmpl,
		"../templates/client_resource.tmpl",
	)
}

func TemplatesClient_resourceTmpl() (*asset, error) {
	bytes, err := TemplatesClient_resourceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/client_resource.tmpl", size: 2220, mode: os.FileMode(420), modTime: time.Unix(1455716159, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesGeneric_mainTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\x4f\xcd\x4b\x2d\xca\x4c\x8e\xcf\x4d\xcc\xcc\x8b\x2f\x49\xcd\x2d\xc8\x49\x2c\x49\x55\xaa\xad\xe5\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x00\x49\x70\x71\xa5\x95\xe6\x25\x83\x99\x1a\x9a\xd5\x5c\x5c\xb5\x5c\xd5\xd5\xa9\x79\x29\x40\x55\x80\x00\x00\x00\xff\xff\xdc\x57\x73\x81\x49\x00\x00\x00")

func TemplatesGeneric_mainTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesGeneric_mainTmpl,
		"../templates/generic_main.tmpl",
	)
}

func TemplatesGeneric_mainTmpl() (*asset, error) {
	bytes, err := TemplatesGeneric_mainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/generic_main.tmpl", size: 73, mode: os.FileMode(420), modTime: time.Unix(1455530577, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesPython_clientTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x51\x41\xcf\xda\x30\x0c\xbd\xe7\x57\x58\x15\x07\x10\xa8\xda\x19\x89\x03\xa0\x4d\x9a\xb4\x4d\x0c\xb6\x5d\xab\xd0\xba\x6d\xb4\x34\x05\x27\x61\x42\x5d\xfe\xfb\x9c\x96\xf5\x2b\x7c\x3d\x34\x8e\xfd\xfc\xfc\xf2\xdc\x75\x50\x60\xa9\x0c\x42\x72\xb9\xbb\xba\x35\x59\xae\x15\x1a\x97\x39\x6c\x2e\x5a\x3a\x4c\x20\x04\xa1\x9a\x4b\x4b\x0e\x08\xaf\x1e\xad\xb3\xa2\xa4\xb6\x81\x07\xd0\x3b\xa5\x2d\x3c\x10\x67\xaf\x74\xf1\xdd\x23\xdd\x4f\x8e\x94\xa9\x84\xd8\x6d\x4f\x1f\xb3\x9f\xc7\xcf\xb0\x81\xa4\xeb\xd2\x9d\xb4\xc8\xb7\x10\x12\x21\x72\x2d\xad\x85\x7d\x4f\xb3\x16\xc0\x1f\x4b\x81\x2c\x53\x46\xb9\x2c\x9b\x5b\xd4\xe5\x62\xc8\xc7\x2f\x5e\x53\x4f\x9a\x89\xfe\x73\x0a\xc1\xf2\x49\x9a\x0a\x61\xf6\x7b\x05\xb3\x1b\xac\x37\x90\x1e\xfa\x77\x7c\x45\xfe\x17\x36\xaa\x1f\xa9\xbb\x6e\x76\x4b\x87\xc2\x37\xd9\x60\x08\xf3\x3e\x73\x90\x24\x1b\x1b\xc2\x64\x18\xe7\xf3\xb6\x69\x58\x98\x46\x53\xb9\x3a\x12\x73\xc4\x23\xd2\x4f\xde\xe4\xfb\xa1\xc6\x3d\x93\x0e\x55\x02\xdb\xf8\xd2\xf7\x81\x05\x24\x49\x32\x81\x8d\x82\xcb\xa8\xb8\x8c\xcc\x2f\xac\xdc\x12\x75\x95\xf1\x40\x53\x4c\x86\x3c\x33\x3d\xd7\x3c\x29\xb6\xa6\x7f\xd0\x11\x6d\xeb\x29\xc7\x83\x74\xf5\x04\x41\xe8\x3c\x99\x71\x8b\x69\x8f\xfd\x85\x74\x86\xbf\xf0\xa3\xfd\xd2\xfe\x41\x62\x4b\x46\x9f\x97\x3d\xe5\xf2\xdd\x4e\xe7\xd7\x18\x0f\xa6\x2d\x06\x03\x8f\x5b\xaa\xd8\x0c\x58\xd5\x28\x0b\x24\xbb\x79\x9c\x8b\xb8\x20\x96\xd9\x6f\xe1\x2d\xfc\x17\x00\x00\xff\xff\x76\x5e\x4d\xac\x75\x02\x00\x00")

func TemplatesPython_clientTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesPython_clientTmpl,
		"../templates/python_client.tmpl",
	)
}

func TemplatesPython_clientTmpl() (*asset, error) {
	bytes, err := TemplatesPython_clientTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/python_client.tmpl", size: 629, mode: os.FileMode(420), modTime: time.Unix(1455716159, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesPython_client_utilsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8f\xcd\x6a\xc3\x30\x10\x84\xef\x7a\x8a\x65\x0f\xc5\xc6\xed\xa1\x57\x83\xe8\x1b\x94\x96\x1c\x43\x30\x4e\xbc\x4e\x96\xc8\xeb\x58\x3f\x07\x23\xfc\xee\x91\x1c\x43\x3c\x07\x21\xcd\xa7\x99\x65\x63\xec\xa8\x67\x21\xc0\xc7\xec\x6f\xa3\x34\x17\xc3\x24\xbe\x09\x9e\x8d\xc3\x65\x51\x09\xc3\x39\xb0\xe9\xfe\x03\xd9\xf9\xe0\x2d\xcb\xb5\x98\xf2\xfd\xaf\xb5\xed\xe0\xf4\xef\x28\x54\xd6\x0a\x92\xb8\x87\x1d\x01\x76\x90\xe1\x8b\x65\x59\xf2\xc1\x0a\x20\xaa\xd5\x9a\x1c\x68\xc0\x1f\x5c\x1f\xfd\x68\xe1\x4e\xf3\x27\x90\xa1\x01\x58\xf6\x4d\xef\x86\x14\xa9\x74\xfe\x07\x15\xa0\xc6\x74\x3a\x6f\x8b\x1c\x29\xb3\xf3\xb1\x35\x6f\x83\x26\x77\xac\xbf\xbe\x4f\x2a\x46\x92\x2e\xed\xf2\x0c\x00\x00\xff\xff\xe8\x33\x27\xf9\xed\x00\x00\x00")

func TemplatesPython_client_utilsTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesPython_client_utilsTmpl,
		"../templates/python_client_utils.tmpl",
	)
}

func TemplatesPython_client_utilsTmpl() (*asset, error) {
	bytes, err := TemplatesPython_client_utilsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/python_client_utils.tmpl", size: 237, mode: os.FileMode(420), modTime: time.Unix(1455530577, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesServer_mainTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x94\x51\x6f\xda\x30\x10\xc7\x9f\xe3\x4f\x71\xb3\xfa\x60\x43\x14\x2a\xf1\x52\x2a\xf1\x50\xad\xaa\xd4\x69\x63\x08\xba\xa7\x69\x6a\x0d\x1c\x60\x35\x71\x22\xfb\xd2\xb5\x8a\xf8\xee\x3b\x27\x2d\x85\x75\xeb\xdb\xa4\xbd\x80\x73\xfe\xff\xcf\xbf\xbb\xe4\xdc\x34\x2b\x5c\x5b\x87\x20\x03\xfa\x07\xf4\xb7\x85\xb1\xee\x96\xb0\xa8\x72\x43\x28\x77\x3b\x51\x99\xe5\xbd\xd9\x20\x34\x4d\x36\xed\x96\x13\x53\x20\x6f\x08\x5b\x54\xa5\x27\x50\x22\x91\x79\xb9\x91\xfc\xe7\x90\x06\x5b\xa2\x2a\xae\xc9\x16\x28\x05\x2f\x36\x96\xb6\xf5\x22\x5b\x96\xc5\x60\x53\x7a\x9b\xe7\x66\x50\xd4\x8f\x52\x68\x21\x06\x83\x4b\x3e\x05\x3c\x56\x1e\x03\x3a\x82\xd9\xd5\xc7\xe1\x70\x34\x82\x15\x87\x05\x3d\x55\x08\xad\x20\xe6\xca\x6e\xf8\x27\x5a\xbe\x18\x1f\xb6\x26\xff\x34\xff\x3a\x81\x92\x91\xbd\x5d\x21\x14\xaf\x41\xb1\xae\xdd\x12\x14\x41\x2f\x7a\x35\x1c\xe8\x95\x06\xf5\xfd\xc7\xe2\x89\x30\x05\x36\x96\x5e\x43\x23\x12\x8f\x54\x7b\x07\xdd\x86\xda\x9f\xa5\x7a\xa4\xb3\xab\xd2\x17\x86\xd4\x9d\xbc\x83\x7e\x87\xd1\x22\x0e\x47\xfc\xc8\x41\xad\x53\x70\x36\x17\xbb\x03\xb0\x1b\x7c\xa4\x37\x60\x31\xf8\x17\xb0\xb8\xf5\x6f\xc1\xbe\xb9\xe2\x4f\x3d\xab\xdd\x3b\x5d\x3b\xf2\xa8\xc5\x33\x84\xee\xe8\x22\x1c\x85\x16\x15\xce\xc7\xdd\xf1\x53\x96\x63\xe4\xe9\x1f\xd2\xf4\x39\x90\x42\x20\x6f\xdd\x46\x2d\xb4\x16\x89\x5d\xb7\xb6\x0f\xe3\xc8\x17\x13\xbd\x94\xc9\x51\x91\x30\x6f\xd2\x23\x18\xb7\xef\x5d\x51\xd0\xfb\x2e\xbc\xad\xe6\xb8\xd1\xb5\x7b\xa7\xd5\x47\x9e\xff\xa8\x9a\xdf\x30\xe7\x5d\x66\xfd\x7c\xc4\xc1\x37\x70\xfc\xf2\x5f\x74\xfb\x14\x71\x6a\x55\xf7\xcd\xc4\x12\x78\xbe\xb2\x09\xfe\x9c\x95\x35\xa1\x67\x99\x48\x9a\x06\xbc\x71\x3c\xc6\x27\xf7\x29\x9c\x3c\x44\x51\x36\xc3\x50\xd6\x7e\x89\xe1\x12\xd7\xc0\x03\xcd\xa2\xac\x9b\xed\x6b\xc7\xbe\xb5\x59\x62\x9b\x21\x28\x9f\xc2\x7e\xef\x62\x7a\xdd\xec\x74\x9b\x11\xdd\x2a\xfa\x44\xc2\xe3\x9f\x4d\x19\x89\x72\xa7\x64\x20\xe3\x29\xd2\x77\x37\x8a\x64\x6d\xbc\x13\xb2\xcf\x36\x10\xba\x0b\xb7\x9a\xc7\xb8\x92\xe7\x67\xa7\x67\xa7\x32\x05\xdf\x96\xf1\x9a\xee\x57\x00\x00\x00\xff\xff\x29\x06\x41\xfe\x93\x04\x00\x00")

func TemplatesServer_mainTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesServer_mainTmpl,
		"../templates/server_main.tmpl",
	)
}

func TemplatesServer_mainTmpl() (*asset, error) {
	bytes, err := TemplatesServer_mainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/server_main.tmpl", size: 1171, mode: os.FileMode(420), modTime: time.Unix(1455716585, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesServer_resources_apiTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x52\xcd\x8e\xd3\x30\x10\x3e\xc7\x4f\x31\x58\x15\x4a\x50\xe5\xee\x81\x13\xab\x3d\x80\x00\x01\x12\xa5\xb0\x12\x1c\x57\x26\x99\xb6\xa1\x89\x9d\x38\x4e\xa3\xca\xca\xbb\x33\x63\xa7\x94\xcd\x69\x3c\x33\xdf\x8f\x3f\x27\x84\x0a\xf7\xb5\x41\x90\x0e\x07\x3b\xba\x12\x9f\x74\x57\x3f\x79\x6c\xbb\x46\x7b\x94\xf3\x2c\x3a\x5d\x9e\xf4\x01\x21\x04\xb5\x4b\xe5\x56\xb7\x48\x03\x51\xb7\x9d\x75\x1e\x72\x91\x85\x00\xf5\x1e\xd4\x16\xb1\xfa\xf2\xf8\x6d\x0b\xf3\x2c\xd1\x94\xb6\xaa\xcd\x61\xf3\x67\xb0\x46\xd2\x02\x9a\x8a\xfa\x22\x93\x06\xfd\xe6\xe8\x7d\x27\x45\x21\x42\x58\x91\x1e\x13\xc2\x9b\x07\x22\x48\xcc\xfe\xd2\x45\xbd\x74\x7c\xbb\xfb\x0c\x83\x77\x63\xe9\x21\x08\x92\x25\x10\x38\x6d\xc8\xd2\xea\xb4\x86\xd5\x39\x22\xbf\xa2\x3f\xda\x6a\x60\x85\xfd\x68\xca\x9c\x58\xe1\x46\x1e\x49\x0a\x6e\x9c\x97\xcd\xd4\xcd\x27\x60\x27\xea\x07\x0e\x9d\x35\x03\xfe\x72\xb5\x47\xb7\x06\x07\xaf\x96\x7e\x3f\xe2\xe0\x09\xc9\x77\xbc\x8a\xf6\xac\xda\xb3\x2c\xd1\x7d\x1f\xd1\x5d\x76\xda\x11\x1f\x41\x07\xd2\x07\xfa\x36\x1b\x16\x3b\xf5\xf3\xcc\x6b\x0e\x7b\xf5\xd1\xba\xf6\xa7\x6e\x46\xcc\xe5\x32\x91\x05\x93\x52\x2c\x9c\x4a\x08\x1c\x20\xe9\xbd\xb3\xd5\x85\x1b\x67\xed\x18\xc7\x47\x8e\xe2\x36\x11\x19\x6d\xa2\x73\x4c\xcc\xd9\x52\xea\xd3\x7b\xa4\xb0\xd1\xe5\x4e\xf1\x52\xa1\xd2\x39\x7f\xb9\x10\x14\xf7\x11\xf0\xe2\x01\x4c\xdd\xf0\x5d\xb2\x49\xc5\xab\x7e\x42\xcd\xb0\xd7\x77\x77\xe4\x25\x73\xe8\x47\x67\x44\x16\xed\x5c\xdf\xeb\xdf\xe3\x72\x46\xd1\xcd\xcd\xdd\xd2\x88\xf6\x52\xcd\xb3\xab\xa9\x0f\x26\x99\x9a\x0a\x95\x4a\xf6\x93\xd6\x8a\xfb\xe7\x12\x14\x17\x3d\x9a\x6d\x5b\x34\x1e\x7e\x63\x63\x27\x68\xf8\xaf\xf4\x16\x74\x55\xc1\x31\xda\x8c\x6b\x93\x4a\x9e\xd5\x23\xfa\x5c\x9e\xf0\x22\xd7\xf2\xcc\xb1\x52\x9a\xb3\xf8\x8f\xf3\x56\xfe\x0d\x00\x00\xff\xff\x12\x29\xc2\x9d\xe4\x02\x00\x00")

func TemplatesServer_resources_apiTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesServer_resources_apiTmpl,
		"../templates/server_resources_api.tmpl",
	)
}

func TemplatesServer_resources_apiTmpl() (*asset, error) {
	bytes, err := TemplatesServer_resources_apiTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/server_resources_api.tmpl", size: 740, mode: os.FileMode(420), modTime: time.Unix(1455716585, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesServer_resources_interfaceTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\x4f\x6f\xd4\x30\x10\xc5\xcf\xf1\xa7\x18\x45\x7b\xd8\xad\xda\xe4\x8e\xc4\x89\x3f\x82\x03\x08\x55\x08\x8e\x95\x37\x9e\x24\x56\xe3\x71\xb0\xc7\xdb\x56\x96\xbf\x3b\xe3\xdd\x14\x10\x54\xe2\x94\x89\xe7\xcd\x9b\xf7\xb3\x73\x36\x38\x5a\x42\x68\x03\x46\x9f\xc2\x80\x77\x76\xbc\x63\x74\xeb\xa2\x19\xdb\x52\xd4\xaa\x87\x7b\x3d\x21\xe4\xdc\x7d\xb9\x94\x9f\xb5\x43\x69\xa8\xbe\xff\x3a\xdb\x08\xa3\x5d\x10\xe4\xab\x13\xfb\x9b\x09\x09\x83\x4c\x1a\x38\x3e\xc1\xe4\x6f\x82\x76\x8b\x08\xdf\x7a\x20\xcf\x80\xc6\x32\xf0\xaf\x21\x91\xcc\x9a\x0c\x44\x4b\x83\x58\x30\x3c\xd8\x65\x81\x23\x82\x3f\x61\x78\x08\x96\x19\x09\x4c\x0a\x96\x26\x99\x42\x20\x7c\x64\xd8\x36\x58\x4f\x4a\x59\xb7\xfa\xc0\xb0\x57\x4d\x3b\x59\x9e\xd3\xb1\x1b\xbc\xeb\x27\x1f\xc4\x47\xf7\x2e\x3d\xb6\xd2\x21\xe4\x7e\x66\x5e\x5b\x75\x50\x8a\x9f\xd6\x33\xca\x85\xe1\x23\x31\x86\x51\xd7\xe5\xcf\x55\x56\x4d\xce\x10\x34\x09\xf2\xee\xfe\x1a\x76\x27\x78\xf5\x1a\xba\x4f\xc8\xb3\x37\x11\x84\xbb\x01\xf8\x43\x31\x56\xc9\x58\x35\xbb\x53\xf7\x3e\xd1\xf0\xc6\x3b\x87\xc4\xb1\x94\xbe\x17\xa1\x34\x4b\xc9\x19\xc9\xd4\xd1\xa6\x1e\x6c\x66\x97\x08\xfb\x1a\xad\xbb\xc5\xb8\x7a\x8a\xf8\x5d\xa0\x31\x5c\xc3\xd5\x76\xfa\x23\x61\xe4\xc3\x39\x92\x38\xd4\xed\x72\xef\xa3\x6c\x79\x81\xe1\xd6\x27\xc6\xb8\x0f\x70\x25\xe0\xdd\xf9\x4f\x9c\xec\x0b\xca\x83\x10\xfc\x1f\xb3\x09\xdd\x07\x79\x9d\x05\x2b\xd5\xbe\x3d\x27\x7f\x47\x66\xf5\x72\x57\xa5\xb4\x62\xdd\xfd\x43\x73\x78\x76\xd8\xf4\xdf\x30\x1c\x45\xfb\x17\xc1\xef\xfa\x67\x00\x00\x00\xff\xff\x18\x3f\x3b\x71\x7f\x02\x00\x00")

func TemplatesServer_resources_interfaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesServer_resources_interfaceTmpl,
		"../templates/server_resources_interface.tmpl",
	)
}

func TemplatesServer_resources_interfaceTmpl() (*asset, error) {
	bytes, err := TemplatesServer_resources_interfaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/server_resources_interface.tmpl", size: 639, mode: os.FileMode(420), modTime: time.Unix(1455716585, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesStructTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8e\x41\x4b\xc4\x30\x10\x85\xef\xf9\x15\x43\xe8\x51\xb2\xf7\x05\x4f\x8a\xe0\x45\x3c\x78\x77\x43\x3b\x95\xd8\x34\x89\x4d\x2a\x94\x61\xfe\xbb\x93\x26\xc2\xe6\xf4\x32\x6f\xde\xbc\x8f\x68\xc2\xd9\x05\x04\x9d\xcb\xb6\x8f\xe5\xb3\xe0\x9a\xbc\x2d\xa8\x99\x55\xb2\xe3\x62\xbf\x10\x88\xcc\x7b\x93\x6f\x76\x45\x31\x14\x91\x9b\x41\x52\x83\x79\xc6\x3c\x6e\x2e\x15\x17\x03\x68\x0d\xcc\x97\x0b\x11\x86\x89\xb9\xc6\xee\x5c\x89\x95\x23\xd5\x63\x60\xea\x19\x59\x85\xd6\x09\xa4\x40\x9e\x18\x9b\x0d\x52\x37\x2c\x78\x3c\xc0\xf0\x6b\xfd\x8e\x70\x7d\x04\xf3\xe2\xd0\x4f\x59\x02\x44\x6d\x6a\x1a\x47\xcd\x08\x07\xfe\xf4\x65\xf3\x9a\x9f\xe2\x9a\x62\x76\x27\xce\x6c\x7d\xc6\x93\xa3\xdb\x1f\xd2\x2f\xff\xdb\x77\x8e\xe1\xaa\x65\x2c\x45\xcc\xfa\x06\x9d\xb8\x63\x34\xcd\xea\x5f\xfd\x05\x00\x00\xff\xff\x65\x3c\x23\x77\x24\x01\x00\x00")

func TemplatesStructTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesStructTmpl,
		"../templates/struct.tmpl",
	)
}

func TemplatesStructTmpl() (*asset, error) {
	bytes, err := TemplatesStructTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/struct.tmpl", size: 292, mode: os.FileMode(420), modTime: time.Unix(1455716585, 0)}
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
	"../templates/client_helper_resource.tmpl": TemplatesClient_helper_resourceTmpl,
	"../templates/client_resource.tmpl": TemplatesClient_resourceTmpl,
	"../templates/generic_main.tmpl": TemplatesGeneric_mainTmpl,
	"../templates/python_client.tmpl": TemplatesPython_clientTmpl,
	"../templates/python_client_utils.tmpl": TemplatesPython_client_utilsTmpl,
	"../templates/server_main.tmpl": TemplatesServer_mainTmpl,
	"../templates/server_resources_api.tmpl": TemplatesServer_resources_apiTmpl,
	"../templates/server_resources_interface.tmpl": TemplatesServer_resources_interfaceTmpl,
	"../templates/struct.tmpl": TemplatesStructTmpl,
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
	"..": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"client_helper_resource.tmpl": &bintree{TemplatesClient_helper_resourceTmpl, map[string]*bintree{}},
			"client_resource.tmpl": &bintree{TemplatesClient_resourceTmpl, map[string]*bintree{}},
			"generic_main.tmpl": &bintree{TemplatesGeneric_mainTmpl, map[string]*bintree{}},
			"python_client.tmpl": &bintree{TemplatesPython_clientTmpl, map[string]*bintree{}},
			"python_client_utils.tmpl": &bintree{TemplatesPython_client_utilsTmpl, map[string]*bintree{}},
			"server_main.tmpl": &bintree{TemplatesServer_mainTmpl, map[string]*bintree{}},
			"server_resources_api.tmpl": &bintree{TemplatesServer_resources_apiTmpl, map[string]*bintree{}},
			"server_resources_interface.tmpl": &bintree{TemplatesServer_resources_interfaceTmpl, map[string]*bintree{}},
			"struct.tmpl": &bintree{TemplatesStructTmpl, map[string]*bintree{}},
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

