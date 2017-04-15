// Code generated by go-bindata.
// sources:
// data/config_schema_v3.0.json
// data/config_schema_v3.1.json
// data/config_schema_v3.2.json
// DO NOT EDIT!

package schema

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

var _dataConfig_schema_v30Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4b\x8f\xdb\x38\x12\xbe\xfb\x57\x08\x4a\x6e\x71\x77\x07\xd8\x60\x81\xcd\x6d\x8f\x7b\xda\x39\x4f\x43\x11\x68\xa9\x6c\x33\x4d\x91\x4c\x91\x72\xda\x09\xfc\xdf\x07\xd4\xcb\x14\x4d\x8a\xb2\xad\x3c\x30\x98\x53\xb7\xc5\xaa\x62\xbd\xf8\x55\xb1\xa4\xef\xab\x24\x49\xdf\xaa\x62\x0f\x15\x49\x3f\x26\xe9\x5e\x6b\xf9\xf1\xe9\xe9\xb3\x12\xfc\xa1\x7d\xfa\x28\x70\xf7\x54\x22\xd9\xea\x87\xf7\x1f\x9e\xda\x67\x6f\xd2\xb5\xe1\xa3\xa5\x61\x29\x04\xdf\xd2\x5d\xde\xae\xe4\x87\x7f\x3d\xbe\x7f\x34\xec\x2d\x89\x3e\x4a\x30\x44\x62\xf3\x19\x0a\xdd\x3e\x43\xf8\x52\x53\x04\xc3\xfc\x9c\x1e\x00\x15\x15\x3c\xcd\xd6\x2b\xb3\x26\x51\x48\x40\x4d\x41\xa5\x1f\x13\xa3\x5c\x92\x0c\x24\xfd\x03\x4b\xac\xd2\x48\xf9\x2e\x6d\x1e\x9f\x1a\x09\x49\x92\x2a\xc0\x03\x2d\x2c\x09\x83\xaa\x6f\x9e\xce\xf2\x9f\x06\xb2\xb5\x2b\xd5\x52\xb6\x79\x2e\x89\xd6\x80\xfc\x8f\x4b\xdd\x9a\xe5\x4f\xcf\xe4\xe1\xdb\x7f\x1f\xfe\x7c\xff\xf0\x9f\xc7\xfc\x21\x7b\xf7\x76\xb4\x6c\xfc\x8b\xb0\x6d\xb7\x2f\x61\x4b\x39\xd5\x54\xf0\x61\xff\x74\xa0\x3c\x75\xff\x9d\x86\x8d\x49\x59\x36\xc4\x84\x8d\xf6\xde\x12\xa6\x60\x6c\x33\x07\xfd\x55\xe0\x4b\xcc\xe6\x81\xec\x17\xd9\xdc\xed\xef\xb1\x79\x6c\xce\x41\xb0\xba\x8a\x46\xb0\xa7\xfa\x45\xc6\xb4\xdb\xdf\x17\xbf\x55\x6f\xf4\x24\x6d\x4b\x61\xed\xdd\x28\x38\xca\x76\x9f\xab\x7c\xd9\x16\xf6\xd5\xe0\xac\x80\x97\x4a\x90\x4c\x1c\xcd\xb3\x80\x3f\x5a\x82\x0a\xb8\x4e\x07\x17\x24\x49\xba\xa9\x29\x2b\x5d\x8f\x0a\x0e\xff\x37\x22\x9e\xad\x87\x49\xf2\xdd\x3d\xd8\x96\x9c\x66\x7d\xf4\x2b\x1c\xf0\x61\x3d\x60\xcb\xb0\x5e\x08\xae\xe1\x55\x37\x46\x4d\x6f\xdd\xba\x40\x14\x2f\x80\x5b\xca\x60\x2e\x07\xc1\x9d\x9a\x70\x19\xa3\x4a\xe7\x02\xf3\x92\x16\x3a\x3d\x39\xec\x17\xf2\xe2\xf9\x34\xb0\x5a\xbf\xb2\x95\x47\x60\x5a\x10\x99\x93\xb2\x1c\xd9\x41\x10\xc9\x31\x5d\x27\x29\xd5\x50\x29\xbf\x89\x49\x5a\x73\xfa\xa5\x86\xff\x75\x24\x1a\x6b\x70\xe5\x96\x28\xe4\xf2\x82\x77\x28\x6a\x99\x4b\x82\x26\xc1\xa6\xdd\x9f\x16\xa2\xaa\x08\x5f\x2a\xeb\xae\xb1\x63\x86\xe7\x05\xd7\x84\x72\xc0\x9c\x93\x2a\x96\x48\xe6\xd4\x01\x2f\x55\xde\xd6\xbf\xc9\x34\xda\xe6\x2d\xbf\x72\x04\x0c\xc5\x70\xd1\x78\x94\x7c\x2a\xb1\x5b\x31\x26\xb5\x8d\x6e\xa9\xc3\x98\x2b\x20\x58\xec\x6f\xe4\x17\x15\xa1\x7c\x8e\xef\x80\x6b\x3c\x4a\x41\xdb\x7c\xf9\xed\x12\x01\xf8\x21\x1f\xb0\xe4\x6a\x37\x00\x3f\x50\x14\xbc\xea\x4f\xc3\x1c\x80\x19\x40\xde\xf0\xbf\x4a\xa1\xc0\x75\x8c\x63\xa0\xbd\x34\x98\x3a\xf2\x49\xcf\xf1\xdc\x1b\xbe\x4e\x52\x5e\x57\x1b\x40\xd3\xd2\x8d\x28\xb7\x02\x2b\x62\x94\xed\xf7\xb6\x96\x47\x9e\xf6\x64\x9e\xed\x40\xdb\x06\x53\xd6\x09\xcb\x19\xe5\x2f\xcb\xa7\x38\xbc\x6a\x24\xf9\x5e\x28\x3d\x1f\xc3\x2d\xf6\x3d\x10\xa6\xf7\xc5\x1e\x8a\x97\x09\x76\x9b\x6a\xc4\x2d\x94\x9e\x93\xe4\xb4\x22\xbb\x38\x91\x2c\x62\x24\x8c\x6c\x80\xdd\x64\xe7\xa2\xce\xb7\xc4\x8a\xdd\xce\x90\x86\x32\xee\xa2\x73\xe9\x96\x63\x35\xbf\x44\x7a\x00\x9c\x5b\xc0\x85\x3c\x37\x5c\xee\x62\xbc\x01\x49\xe2\xdd\xe7\x88\xf4\xd3\x63\xdb\x7c\x4e\x9c\xaa\xe6\x3f\xc6\xd2\xcc\x6d\x17\x12\xa7\xee\xfb\x9e\x38\x16\xce\x6b\x28\x46\x51\xa9\x48\x61\xfa\x06\x04\x15\x88\xeb\x99\xb4\x6b\xf6\xf3\x4a\x94\xa1\x04\xbd\x20\x76\x7d\x13\x44\xea\xab\x0b\x61\x72\x53\xff\x38\x2b\x74\xd1\x0b\x44\xc4\x9a\x90\x7a\x73\xd5\x3c\xab\x1b\x4f\xb1\x86\x8e\x30\x4a\x14\xc4\x0f\x7b\xd0\x91\x23\x69\x54\x1e\x3e\xcc\xcc\x09\x1f\xef\xbf\x27\x79\x03\xac\x41\x99\xf3\x7b\xe4\x88\xa8\xb3\x2a\xcd\x71\xf3\x29\x92\x45\x4e\xdb\x0f\x6e\xe1\x25\x2d\xc3\x58\xd1\x20\x84\x7d\xc0\xa4\x40\x7d\x71\xba\x7e\x4e\xb9\x6f\xb7\xbe\xbb\xda\x4b\xa4\x07\xca\x60\x07\xe3\x5b\xcb\x46\x08\x06\x84\x8f\xa0\x07\x81\x94\xb9\xe0\xec\x38\x83\x52\x69\x82\xd1\x0b\x85\x82\xa2\x46\xaa\x8f\xb9\x90\x7a\xf1\x3e\x43\xed\xab\x5c\xd1\x6f\x30\x8e\xe6\x19\xef\x3b\x41\xd9\x88\xe7\xa8\x0a\x7d\x5b\xbd\x56\xba\xa4\x3c\x17\x12\x78\xd4\x3b\x4a\x0b\x99\xef\x90\x14\x90\x4b\x40\x2a\x4a\x9f\x81\x6b\x3b\xd6\x65\x8d\xc4\xec\x7f\x29\x46\xd1\x1d\x27\x2c\xe6\x68\x5d\xc9\xed\x8d\x17\x0b\xad\xe3\xe1\xae\x19\xad\x68\xf8\x1c\x78\x00\x76\x46\x0d\x68\xf1\xdf\x0f\xfb\x13\x90\x7f\xd6\x94\x72\x0d\x3b\x40\x1f\x52\x4e\x74\x1d\xd3\x4d\xc7\x8c\x6e\x63\x4f\x70\x1c\xd0\x09\x3d\x1a\x06\x25\xb6\xda\xcf\xe0\xeb\x45\xbc\x7a\x8d\x86\xbf\x8d\xbc\x75\xa7\x48\xe6\xa5\xbf\x0a\xce\x5d\x35\xb2\x20\xa2\x9e\xbc\x88\x5a\xab\x68\x63\xd8\xd0\x70\x35\xd5\xd4\x0c\xa4\xd6\x14\x73\x51\xbc\x30\x8d\x92\x39\x04\x25\xf5\x6b\xbb\x72\x2c\xbb\x62\x8e\xec\xdc\x59\x7a\x01\xbe\x89\xa2\x4d\x1a\x9d\xc0\x4e\x4f\x37\x3b\xa2\xe0\xe4\x91\x2a\xb2\x71\x66\x6e\xbe\xc3\x6d\xb2\x11\x0f\x71\x8c\x41\xd0\x48\x9d\xb8\x74\x68\x3b\xc2\x13\x50\xbf\xe7\xe0\x40\xd3\x0a\x44\xed\xaf\x59\x2b\x3b\xbf\x3b\xa6\xd4\x9a\xcc\x46\x82\x6a\x51\xba\x31\x7d\x1e\x82\xda\xf7\x17\xd1\xc0\xcd\x39\x24\x08\x92\xd1\x82\xa8\x18\x10\xdd\x71\x41\xad\x65\x49\x34\xe4\xed\x8b\xaa\xab\xa0\x7f\x02\xf3\x25\x41\xc2\x18\x30\xaa\xaa\x39\x18\x9a\x96\xc0\xc8\xf1\xa6\xf2\xd9\xb0\x6f\x09\x65\x35\x42\x4e\x0a\xdd\xbd\x0b\x8b\xe4\x5c\x5a\x09\x4e\xb5\xf0\x22\xc4\xbc\x2d\x2b\xf2\x9a\xf7\xdb\x36\x24\xde\x03\x13\x6c\xeb\xe6\xde\x2d\xad\x4c\x50\xa2\xc6\xe2\xc2\xd9\x37\x87\xe8\x5c\xeb\x03\x19\xd3\xef\x78\x61\x3a\x82\x32\x48\x32\x5c\xfd\xa3\xfc\xd1\xd2\xd2\xf5\x99\xb9\x14\x8c\x16\xc7\xa5\x2c\x2c\x04\x6f\x9d\x3c\x27\x21\xee\xcc\x40\x93\x0e\xa6\x15\xaa\xa4\x8e\x1e\xd6\x86\xe1\x2b\xe5\xa5\xf8\x7a\xc5\x86\xcb\xa5\x92\x64\xa4\x00\x07\xef\xee\x75\xb4\xd2\x48\x28\xd7\x57\x97\xf3\x7b\xcd\xba\xa3\x9a\x0f\xf9\x19\x41\xfd\x81\x2e\xfe\x26\x35\x80\xf4\x85\xac\xa3\xf3\xa0\x0a\x2a\x81\xde\x04\x5c\xe0\xcd\x77\xcc\xc4\x9e\x6c\x81\xaa\x36\x6b\x80\xd8\x51\x99\xfb\xe2\xe2\xb7\x8d\xf8\x90\x30\x8b\x03\x12\x95\xa4\x5a\xea\x74\xcc\x1e\xa9\xa6\xde\x1a\x9c\x4c\x8f\x22\x92\xf0\x38\x22\xa6\x75\x5c\xf7\x8e\x42\xd5\x1b\x0e\x93\x1d\x95\xe5\x4f\xdf\x7b\xde\xf9\xd7\x94\x53\xf8\x52\x72\x1f\xe8\xf5\x6f\x43\x02\x51\x7d\x1e\x7a\xe6\xf5\xe0\xab\x6c\x76\x88\x83\xaf\x22\x96\xd3\xbf\x69\xdf\xdd\x11\x81\xaf\xcf\xbf\xb2\x13\xbc\x03\x5c\xba\x4f\x3c\x22\xd8\xd2\x51\xfd\x03\x2d\x7f\x93\x44\xfc\x79\xf9\xe5\x4c\xb3\xac\x3c\xbb\xbc\x68\x4e\xa5\xc4\xec\x31\x7e\xc7\x91\x8d\xd5\x70\xc9\x3c\x5f\xda\x8d\x61\x79\x6a\x7a\xd1\x93\x04\xc6\xba\xce\xa6\x9d\x13\xa7\x2d\x5f\x30\xc3\x1f\xdf\x4d\x14\x9f\xa9\xd7\x6d\x3f\x08\xb5\x17\x98\x0c\xf9\x63\xea\x74\xac\xbd\x77\x2f\x3f\x17\x0b\x80\x9a\xc5\x7f\xf1\xf1\x98\xb1\x93\x1f\x2f\x06\x21\xdf\xc7\xd3\xbd\xf6\xc3\xaf\x6c\xe4\x1f\x87\xa4\x7d\x79\x6d\x41\x4a\x66\x37\xf1\xa1\x30\x7a\x3f\x29\x73\x67\x8b\xfd\xa7\x5d\x99\x1f\xae\x56\xf6\xdf\xe6\x33\xbc\xd5\x69\xf5\x57\x00\x00\x00\xff\xff\x78\x30\xec\x51\x0e\x2b\x00\x00")

func dataConfig_schema_v30JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v30Json,
		"data/config_schema_v3.0.json",
	)
}

func dataConfig_schema_v30Json() (*asset, error) {
	bytes, err := dataConfig_schema_v30JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/config_schema_v3.0.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataConfig_schema_v31Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1a\xcb\x8e\xdb\x36\xf0\xee\xaf\x10\x94\xdc\xe2\xdd\x4d\xd1\xa0\x40\x73\xeb\xb1\xa7\xf6\xdc\x85\x23\xd0\xd2\x58\x66\x96\x22\x19\x92\x72\xd6\x09\xfc\xef\x05\xf5\x32\x45\x91\x22\x6d\x2b\xd9\x45\xd1\xd3\xae\xc5\x99\xe1\xbc\x67\x38\xe4\xf7\x55\x92\xa4\x6f\x65\xbe\x87\x0a\xa5\x1f\x93\x74\xaf\x14\xff\xf8\xf0\xf0\x59\x32\x7a\xd7\x7e\xbd\x67\xa2\x7c\x28\x04\xda\xa9\xbb\xf7\x1f\x1e\xda\x6f\x6f\xd2\xb5\xc6\xc3\x85\x46\xc9\x19\xdd\xe1\x32\x6b\x57\xb2\xc3\xaf\xf7\xbf\xdc\x6b\xf4\x16\x44\x1d\x39\x68\x20\xb6\xfd\x0c\xb9\x6a\xbf\x09\xf8\x52\x63\x01\x1a\xf9\x31\x3d\x80\x90\x98\xd1\x74\xb3\x5e\xe9\x35\x2e\x18\x07\xa1\x30\xc8\xf4\x63\xa2\x99\x4b\x92\x01\xa4\xff\x60\x90\x95\x4a\x60\x5a\xa6\xcd\xe7\x53\x43\x21\x49\x52\x09\xe2\x80\x73\x83\xc2\xc0\xea\x9b\x87\x33\xfd\x87\x01\x6c\x6d\x53\x35\x98\x6d\xbe\x73\xa4\x14\x08\xfa\xf7\x94\xb7\x66\xf9\xd3\x23\xba\xfb\xf6\xc7\xdd\x3f\xef\xef\x7e\xbf\xcf\xee\x36\xef\xde\x8e\x96\xb5\x7e\x05\xec\xda\xed\x0b\xd8\x61\x8a\x15\x66\x74\xd8\x3f\x1d\x20\x4f\xdd\x7f\xa7\x61\x63\x54\x14\x0d\x30\x22\xa3\xbd\x77\x88\x48\x18\xcb\x4c\x41\x7d\x65\xe2\x29\x24\xf3\x00\xf6\x42\x32\x77\xfb\x3b\x64\x1e\x8b\x73\x60\xa4\xae\x82\x16\xec\xa1\x5e\x48\x98\x76\xfb\x65\xec\x27\x21\x17\xa0\xc2\x2e\xdb\x42\xbd\x98\xc7\xea\xed\x6f\x13\x78\xd5\x0b\x3d\x0b\xdb\x42\x18\x7b\x37\x0c\x8e\xc2\xdb\xa5\x2a\x57\x78\xf9\x75\x35\x28\xcb\xa3\xa5\x02\x38\x61\x47\xfd\xcd\xa3\x8f\x16\xa0\x02\xaa\xd2\x41\x05\x49\x92\x6e\x6b\x4c\x0a\x5b\xa3\x8c\xc2\x5f\x9a\xc4\xa3\xf1\x31\x49\xbe\xdb\x99\xcc\xa0\xd3\xac\x8f\x7e\xf9\x0d\x3e\xac\x7b\x64\x19\xd6\x73\x46\x15\x3c\xab\x46\xa8\xf9\xad\x5b\x15\xb0\xfc\x09\xc4\x0e\x13\x88\xc5\x40\xa2\x94\x33\x2a\x23\x58\xaa\x8c\x89\xac\xc0\xb9\x4a\x4f\x16\xfa\x84\x5e\xd8\x9f\x06\x54\xe3\xd7\x66\xe5\x20\x98\xe6\x88\x67\xa8\x28\x46\x72\x20\x21\xd0\x31\x5d\x27\x29\x56\x50\x49\xb7\x88\x49\x5a\x53\xfc\xa5\x86\x3f\x3b\x10\x25\x6a\xb0\xe9\x16\x82\xf1\xe5\x09\x97\x82\xd5\x3c\xe3\x48\x68\x07\x9b\x57\x7f\x9a\xb3\xaa\x42\x74\x29\xaf\xbb\x44\x8e\x08\xcd\x33\xaa\x10\xa6\x20\x32\x8a\xaa\x90\x23\xe9\xa8\x03\x5a\xc8\xac\x2d\xf8\xb3\x6e\xb4\xcb\x5a\x7c\x69\x11\x18\xaa\xff\xa2\xf6\x28\xe8\x9c\x63\xb7\x64\xb4\x6b\x6b\xde\x52\x0b\x31\x93\x80\x44\xbe\xbf\x12\x9f\x55\x08\xd3\x18\xdd\x01\x55\xe2\xc8\x19\x6e\xfd\xe5\xd5\x39\x02\xd0\x43\x36\xe4\x92\x8b\xd5\x00\xf4\x80\x05\xa3\x55\x1f\x0d\x31\x09\x66\x48\xf2\x1a\xff\x99\x33\x09\xb6\x62\x2c\x01\xcd\xa5\x41\xd4\x91\x4e\x7a\x8c\xc7\x5e\xf0\x75\x92\xd2\xba\xda\x82\xd0\x3d\xec\x08\x72\xc7\x44\x85\x34\xb3\xfd\xde\xc6\xf2\x48\xd3\x0e\xcf\x33\x15\x68\xca\xa0\xcb\x3a\x22\x19\xc1\xf4\x69\x79\x17\x87\x67\x25\x50\xb6\x67\x52\xc5\xe7\x70\x03\x7d\x0f\x88\xa8\x7d\xbe\x87\xfc\x69\x06\xdd\x84\x1a\x61\x33\xa9\x62\x9c\x1c\x57\xa8\x0c\x03\xf1\x3c\x04\x42\xd0\x16\xc8\x55\x72\x2e\xaa\x7c\x83\x2c\x2b\x4b\x0d\xea\xf3\xb8\x49\xe7\xd2\x2d\x87\x6a\x7e\x21\xf0\x01\x44\x6c\x01\x67\xfc\xdc\x70\xd9\x8b\xe1\x06\x24\x09\x77\x9f\x23\xd0\x4f\xf7\x6d\xf3\x39\x13\x55\xcd\x7f\x84\xa4\x1b\xbb\x5d\x48\xac\xba\xef\xfa\x62\x49\x18\xd7\x50\x8c\xac\x52\xa1\x5c\xf7\x0d\x02\xa4\xc7\xae\x67\xd0\xee\x74\x93\x55\xac\xf0\x39\xe8\x04\xd8\xd6\x8d\x37\x53\x5f\x5c\x08\x93\xab\xfa\xc7\x28\xd3\x05\x0f\x10\x01\x69\x7c\xec\xc5\xb2\x79\x66\x37\xec\x62\x0d\x1c\x22\x18\x49\x08\x07\xbb\x57\x91\x23\x6a\x98\x1f\x3e\x44\xfa\x84\x0b\xf7\xb7\x59\x5c\x0f\xaa\x97\x66\x7c\x8f\x1c\x20\x75\x66\xa5\x09\x37\x17\x23\x9b\x40\xb4\xfd\xe0\x16\x9e\xe3\xc2\x9f\x2b\x9a\x0c\x61\x06\x18\x67\x42\x4d\xa2\xeb\xe7\x94\xfb\x76\xeb\x9b\xab\x3d\x17\xf8\x80\x09\x94\x30\x3e\xb5\x6c\x19\x23\x80\xe8\x28\xf5\x08\x40\x45\xc6\x28\x39\x46\x40\x4a\x85\x44\xf0\x40\x21\x21\xaf\x05\x56\xc7\x8c\x71\xb5\x78\x9f\x21\xf7\x55\x26\xf1\x37\x18\x5b\xf3\x9c\xef\x3b\x42\x1b\x8b\x21\x6b\x42\x72\xa5\x41\x7d\x29\x29\x1c\xc6\x8e\x44\x18\x4c\x54\xe1\x14\x95\x4a\x56\x8b\x3c\xf6\x80\xad\xf7\x44\xa2\x84\xd8\x23\xbc\x76\xb7\x71\xd8\xcc\x03\x97\x97\x00\x4f\x0a\x5d\x67\xc2\x50\x55\xb6\x7f\x9b\x79\xe5\xe4\x0c\x7d\x79\x94\xb9\xba\xae\x5b\x93\xaa\xc0\x34\x63\x1c\x68\x30\x36\xa4\x62\x3c\x2b\x05\xca\x21\xe3\x20\x30\x73\xaa\x62\x6d\x46\x7a\x51\x0b\xa4\xf7\x9f\x92\x91\xb8\xa4\x88\x84\xc2\x4c\x55\x7c\x77\xe5\xb1\x52\xa9\x70\xb0\xd7\x04\x57\xd8\x1f\x34\x0e\xaf\x8d\xe8\x00\xda\xea\xef\x2e\xfa\x33\x05\xff\xcc\x29\xa6\x0a\x4a\xed\x26\x53\xa7\x9a\xe9\x39\xe7\x5b\xce\x88\x5e\x73\x8f\xc4\xd8\xa0\x33\x7c\x24\x6d\x60\xee\x94\x1b\xc1\xd5\x89\x3a\xf9\x1a\xdd\x75\x34\xf4\xd6\x1d\x23\x1b\x27\xfc\x45\xc5\xdc\x66\x63\xe3\xad\xa7\xee\xa0\xaa\x65\xf0\x58\xd0\xc0\x50\x39\xd7\xd2\x0e\xa0\xc6\xd0\x7e\xd1\x6a\xa1\xdb\x64\x1d\x04\x05\x76\x73\xbb\xb2\x24\xbb\x60\xec\x6e\x9d\x58\x7b\x02\xae\x79\xb2\x09\x1a\x9c\xbf\xcf\xcf\xb6\x3b\x20\xef\xdc\x19\x4b\xb4\xb5\x26\xae\xae\xe0\xd6\xde\x28\x0e\xe1\x1c\x23\x40\x09\x6c\xd9\xa5\x4f\xd4\x66\x3e\x01\xf9\x3a\xc7\x46\x0a\x57\xc0\x6a\x77\xc1\x5b\x99\xfe\xdd\x21\xa5\xc6\x5c\x3e\x60\x54\x03\xd2\xb6\xe9\xe3\x60\xd4\xbe\xbb\x0c\x1a\x2e\x26\x48\x04\x70\x82\x73\x24\x43\x89\xe8\x86\xf1\x44\xcd\x0b\xa4\x20\x6b\xef\x65\x2f\x4a\xfd\x33\x39\x9f\x23\x81\x08\x01\x82\x65\x15\x93\x43\xd3\x02\x08\x3a\x5e\x55\x3e\x1b\xf4\x1d\xc2\xa4\x16\x90\xa1\x5c\x75\x57\xbf\x01\x9f\x4b\x2b\x46\xb1\x62\xce\x0c\x11\xb7\x65\x85\x9e\xb3\x7e\xdb\x06\x24\xd4\xd9\x8c\x9b\xfa\xd8\xc9\x82\xe1\x09\x6d\xe3\x77\x59\x75\x9e\x31\xd1\xb9\xd6\x7b\x3c\xa6\xdf\x71\x22\xba\x00\xa9\x33\xc9\x30\xf8\x09\xe2\x07\x4b\x4b\x77\xca\xc8\x38\x23\x38\x3f\x2e\x25\x61\xce\x68\xab\xe4\x18\x87\xb8\xd1\x03\xb5\x3b\xe8\x56\xa8\xe2\x2a\x18\xac\x0d\xc2\x57\x4c\x0b\xf6\xf5\x82\x0d\x97\x73\x25\x4e\x50\x0e\x56\xbe\xbb\x55\xd1\x52\x09\x84\xa9\xba\xb8\x9c\xdf\x2a\xd6\x0d\xd5\x7c\xf0\xcf\x40\xd6\x1f\xe0\xc2\xf7\xe8\x9e\x4c\x9f\xf3\x3a\x38\x0d\xac\xa0\x62\xc2\xe9\x80\x0b\x3c\xf4\x08\x89\xd8\x83\x2d\x50\xd5\xa2\xc6\xc7\x1d\x54\xc6\xf8\xf2\xa7\x8d\xf0\x88\x78\x13\x4e\x48\x98\xa3\x6a\xa9\xe8\x88\x1e\xa8\xa7\xce\x1a\x9c\xcc\xcf\x2d\x12\xff\xec\x22\xc4\x75\x98\xf7\x0e\x42\xd6\x5b\xea\x19\x21\x4c\x4f\x19\xae\x5b\xfe\xf8\x63\xca\xc9\x7f\x28\xb9\x2d\xe9\xf5\x77\x61\x1e\xab\x3e\x0e\x3d\xf3\x7a\xd0\xd5\x26\xda\xc4\xde\x8b\xa8\xe5\xf8\x6f\xda\x77\x7b\x44\xe0\xea\xf3\x2f\xec\x04\x6f\x48\x2e\xdd\x8b\xa6\x40\x6e\xe9\xa0\xfe\x4f\x2d\xff\x11\x47\xfc\x79\xfe\xd5\x3d\x20\x0b\xbe\xdc\x6a\xa0\xae\x2e\xce\x11\xcf\x95\x5e\x81\xcd\x5e\xda\x14\xe3\xc1\xa2\x61\x92\xe9\x99\x7f\x4e\x93\xd1\xf7\x69\x1d\xc6\x66\xcc\x86\x0d\xe6\x78\xe3\x3b\xae\x90\x73\x83\xa4\x1e\xc4\x73\xbf\x62\x6d\xda\x29\x71\x5e\xf2\x05\x93\xcd\xfd\xbb\x99\x3e\x60\xee\xde\xfb\x07\x15\xd0\x05\x86\x74\x6e\x9b\x5a\x87\x87\x5e\xbb\xd3\x77\x9b\x9e\xf8\x37\xf0\x27\xaf\x38\xb5\x9c\xf4\x38\x99\x49\x7d\x1f\x0f\x5a\xdb\x17\x98\x9b\x91\x7e\x2c\x90\xf6\x15\x89\x91\xdd\x37\xe6\x79\xca\x67\x46\xe7\xdb\x4e\x7b\xcc\xdb\xbf\xb1\xf4\xdc\x6a\xac\xcc\xbf\xcd\x7b\xd8\xd5\x69\xf5\x6f\x00\x00\x00\xff\xff\xfc\xf3\x11\x6a\x88\x2f\x00\x00")

func dataConfig_schema_v31JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v31Json,
		"data/config_schema_v3.1.json",
	)
}

func dataConfig_schema_v31Json() (*asset, error) {
	bytes, err := dataConfig_schema_v31JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/config_schema_v3.1.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataConfig_schema_v32Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1b\x4d\x73\xdc\xa8\xf2\x3e\xbf\x42\xa5\xe4\x16\x7f\xa4\xde\x4b\xbd\xaa\x97\xdb\x3b\xbe\xd3\xee\x79\x5d\x13\x15\x83\x7a\x34\xc4\x12\x10\x40\x63\x4f\x52\xfe\xef\x5b\xfa\x1c\x40\x20\xd0\x8c\x1c\x67\xb7\xf6\x64\x5b\x74\x37\xf4\x77\x37\x8d\x7f\x6c\x92\x24\x7d\x2f\xf1\x01\x2a\x94\x7e\x4e\xd2\x83\x52\xfc\xf3\xfd\xfd\x57\xc9\xe8\x6d\xf7\xf5\x8e\x89\xe2\x3e\x17\x68\xaf\x6e\x3f\x7e\xba\xef\xbe\xbd\x4b\x6f\x1a\x3c\x92\x37\x28\x98\xd1\x3d\x29\xb2\x6e\x25\x3b\xfe\xfb\xee\x5f\x77\x0d\x7a\x07\xa2\x4e\x1c\x1a\x20\xb6\xfb\x0a\x58\x75\xdf\x04\x7c\xab\x89\x80\x06\xf9\x21\x3d\x82\x90\x84\xd1\x74\x7b\xb3\x69\xd6\xb8\x60\x1c\x84\x22\x20\xd3\xcf\x49\x73\xb8\x24\x19\x41\x86\x0f\x1a\x59\xa9\x04\xa1\x45\xda\x7e\x7e\x69\x29\x24\x49\x2a\x41\x1c\x09\xd6\x28\x8c\x47\x7d\x77\x7f\xa6\x7f\x3f\x82\xdd\xd8\x54\xb5\xc3\xb6\xdf\x39\x52\x0a\x04\xfd\x7d\x7a\xb6\x76\xf9\xcb\x03\xba\xfd\xfe\xbf\xdb\x3f\x3e\xde\xfe\xf7\x2e\xbb\xdd\x7e\x78\x6f\x2c\x37\xf2\x15\xb0\xef\xb6\xcf\x61\x4f\x28\x51\x84\xd1\x71\xff\x74\x84\x7c\xe9\x7f\x7b\x19\x37\x46\x79\xde\x02\xa3\xd2\xd8\x7b\x8f\x4a\x09\x26\xcf\x14\xd4\x13\x13\x8f\x21\x9e\x47\xb0\x37\xe2\xb9\xdf\xdf\xc1\xb3\xc9\xce\x91\x95\x75\x15\xd4\xe0\x00\xf5\x46\xcc\x74\xdb\xaf\xa3\x3f\x09\x58\x80\x0a\x9b\x6c\x07\xf5\x66\x16\xdb\x6c\x7f\x1d\xc3\x9b\x81\xe9\x59\xd8\x0e\x42\xdb\xbb\x3d\xa0\xe1\xde\x2e\x51\xb9\xdc\xcb\x2f\xab\x51\x58\x1e\x29\xe5\xc0\x4b\x76\x6a\xbe\x79\xe4\xd1\x01\x54\x40\x55\x3a\x8a\x20\x49\xd2\x5d\x4d\xca\xdc\x96\x28\xa3\xf0\x5b\x43\xe2\x41\xfb\x98\x24\x3f\xec\x48\xa6\xd1\x69\xd7\x8d\xbf\xfc\x0a\x1f\xd7\x3d\xbc\x8c\xeb\x98\x51\x05\xcf\xaa\x65\x6a\x7e\xeb\x4e\x04\x0c\x3f\x82\xd8\x93\x12\x62\x31\x90\x28\xe4\x8c\xc8\x4a\x22\x55\xc6\x44\x96\x13\xac\x9c\xf8\x25\xda\x41\x79\x15\x05\x8c\xf0\x01\xb2\xbd\x60\x55\x90\xca\x3e\xeb\x38\x91\xe9\x8b\x45\x67\x42\x38\x6c\xda\x23\xaa\xf6\xd7\x76\xe3\x20\x98\x62\xc4\x33\x94\xe7\x86\x48\x91\x10\xe8\x94\xde\x24\x29\x51\x50\x49\xb7\xb4\x93\xb4\xa6\xe4\x5b\x0d\xff\xef\x41\x94\xa8\xc1\xa6\x9b\x0b\xc6\xd7\x27\x5c\x08\x56\xf3\x8c\x23\xd1\xd8\xfa\xbc\x25\xa4\x98\x55\x15\xa2\x6b\x39\xc0\x12\x3e\x22\x24\xcf\xa8\x42\x84\x82\xc8\x28\xaa\x42\x36\xdd\x04\x00\xa0\xb9\xcc\xba\xda\x23\xd6\x92\x0c\x02\x63\x21\xb2\xaa\x3e\x72\x3a\xe7\x21\x1d\x99\xc6\x47\x9a\xb3\xa5\x16\x62\x26\x01\x09\x7c\xb8\x10\x9f\x55\x88\xd0\x18\xd9\x01\x55\xe2\xc4\x19\xe9\xec\xe5\x97\x33\x04\xa0\xc7\x6c\x0c\x6b\x8b\xc5\x00\xf4\x48\x04\xa3\xd5\xe0\x0d\x71\x91\x4a\xc3\x7f\xe6\x4c\x82\x2d\x18\x8b\x41\x7d\x69\x64\xd5\x90\xc9\x80\xf1\x30\x30\x7e\x93\xa4\xb4\xae\x76\x20\x9a\x72\xda\x80\xdc\x33\x51\xa1\xe6\xb0\xc3\xde\xda\xb2\x21\x69\x87\xe5\xe9\x02\xd4\x79\x68\x2a\x0c\x54\x66\x25\xa1\x8f\xeb\x9b\x38\x3c\x2b\x81\xb2\x03\x93\xea\x92\x64\x90\x1e\x00\x95\xea\x80\x0f\x80\x1f\x67\xd0\x75\x28\x03\x9b\x49\x15\x63\xe4\xa4\x42\x45\x18\x88\xe3\x10\xc8\xc5\x49\x2f\x5d\x55\xf8\x1a\x59\x56\x14\x0d\xa8\xcf\xe2\x26\x45\x54\xbf\x1c\x2a\x3f\x72\x41\x8e\x20\x62\x6b\x09\xc6\xcf\xb5\x9f\xbd\x18\xae\x85\x92\x70\x21\x6c\x80\x7e\xb9\xeb\xea\xe0\x19\xaf\x6a\x7f\x2b\xcb\x74\x6b\x97\x0b\x89\x95\xf7\x5d\x5f\x2c\x0e\xe3\x0a\x0a\x43\x2b\x15\xc2\x4d\xdd\x20\x40\x7a\xf4\x7a\x06\xed\x1b\xad\xac\x62\xb9\xcf\x40\x27\xc0\xb6\x6c\xbc\x91\x7a\x71\x22\x4c\x2e\x2a\x65\xa3\x54\x17\xec\x65\x02\xdc\xf8\x8e\x17\x7b\xcc\xf3\x71\xc3\x26\xd6\xc2\xa1\x92\x20\x09\x61\x67\xf7\x0a\xd2\xa0\x46\xf8\xf1\x53\xa4\x4d\xb8\x70\xff\x33\x8b\xeb\x41\xf5\xd2\x8c\xaf\x91\x03\xa4\xce\x47\x69\xdd\xcd\x75\x90\x6d\xc0\xdb\x5e\xb9\x84\xe7\x24\xf7\xc7\x8a\x36\x42\xe8\x0e\xc6\x99\x50\x13\xef\x5a\x9e\xee\x7d\x16\xac\x8b\x6b\x88\x53\xe7\x84\xdf\x6d\x3e\x91\xc6\x44\xdd\x51\x48\x53\xff\x0b\xfa\x47\xd8\x33\xd2\x99\x28\xe5\x80\x56\x48\x14\x60\xb6\x21\x84\x2a\x28\x40\x78\x10\x78\xbd\x2b\x89\x3c\x40\xbe\x04\x47\x30\xc5\x30\x2b\xe3\x1c\xc3\xd9\x09\xc7\x3b\x83\x49\x70\x7b\x75\x6d\xc6\x05\x39\x92\x12\x0a\x8b\xe3\x1d\x63\x25\x20\x6a\x24\x0a\x01\x28\xcf\x18\x2d\x4f\x11\x90\x52\x21\x11\x6c\xff\x24\xe0\x5a\x10\x75\xca\x18\x57\xab\x57\x85\xf2\x50\x65\x92\x7c\x07\xd3\xf7\xce\x56\xdf\x13\xda\x5a\x07\xb2\xae\xd6\x92\xd7\x72\x3f\x9f\xd9\xbe\x92\xdb\x48\x56\x0b\x7c\x9d\xe3\xcc\xc2\xd7\x66\x90\x9b\x07\x2e\x96\x00\x4f\x1c\xbe\x57\x61\xa8\x86\x9a\x75\x15\x67\xa0\x96\x27\x89\xd5\x65\xb5\xb5\x54\x39\xa1\x19\xe3\x40\x83\xbe\x21\x15\xe3\x59\x21\x10\x86\x8c\x83\x20\xcc\x29\x0a\x23\xc0\xe6\xb5\x40\xcd\xfe\x53\x32\x92\x14\x14\xb9\xe3\x8e\x06\xaa\x2a\xbe\xbf\xf0\x12\x40\xa9\xb0\xb3\xd7\x25\xa9\x88\xdf\x69\x1c\x56\x1b\x51\xaf\x75\xb5\x9a\xbb\x44\x9b\x29\xcf\xa2\x42\xf6\x4c\x87\x30\xdf\x20\x44\x74\x06\x07\x24\x16\xa4\x8e\xd6\x31\xf7\x9e\xfc\xe4\xea\x1b\x9c\xe7\x32\x86\x64\x2d\xbd\x9b\xfe\x20\x5b\x27\xfc\xa2\xd2\xcb\x3e\xc6\xd6\x5b\xfd\xb8\x9d\xaa\x96\xc1\x26\xae\x85\xa1\x72\xae\x01\x19\x41\xa7\xd3\x9e\xe4\x2f\x11\xa1\x0d\x1d\xb5\xe0\x0e\xdd\x44\xc4\xf1\x7e\xa7\xc8\xd8\xf9\xda\x51\x3f\xba\x22\xd0\x70\x30\xa3\x92\x48\x05\x14\x9f\xe2\x37\xda\x91\xc9\x2d\xf1\x54\x28\xf3\x7d\x57\x5c\xd7\xd5\x42\xa1\xa2\x8b\xb7\xd1\x8d\x4e\xbc\xaf\xf6\x83\xc0\x9f\xc2\x0a\x65\x98\x71\x8f\x6a\xe2\xd9\x58\x9a\x66\xad\xab\x8b\x99\x3a\xd4\x17\x32\x9e\x98\x78\x6c\x12\x52\x4e\xdc\x91\x63\x63\xa1\x2c\x98\x9d\x5a\x77\x7d\x03\x01\xd7\x50\x50\x07\x0d\x0e\x51\xe7\x07\x94\x3d\x90\x77\x78\x48\x24\xda\x59\x63\x33\x57\xa2\x6d\x32\x83\x38\x86\xf3\xbd\x00\x25\x88\x35\x4a\x18\x8a\x26\x3d\xb7\x83\xfc\x35\x2f\xdc\x15\xa9\x80\xd5\xee\x30\xb4\xd1\x0d\xa7\x47\x4a\xb5\xe1\x6a\x40\xa9\x1a\xa4\xad\xd3\x87\x51\xa9\x43\x5f\x1e\x54\x5c\x4c\xc2\x02\x9a\xb7\xa3\x8d\xa8\xec\x26\x80\x97\x04\x23\x19\xaa\x20\xae\xb8\x05\xae\x79\x8e\x14\x64\xdd\x4b\x9c\x45\x35\xdb\x4c\xb1\xc6\x91\x40\x65\x09\x25\x91\x55\x4c\xf1\x93\xe6\x50\x22\x67\xf4\x0f\xd6\xbd\x2d\xfa\x1e\x91\xb2\x16\x90\x21\xec\x0d\xd3\x16\x46\xc5\x28\x51\xcc\x19\x4e\xe2\xb6\xac\xd0\x73\x36\x6c\xdb\x82\x84\x5a\x12\xb3\x1b\x8f\xbd\xc0\xd5\x2c\xa1\xcb\xdd\xcb\xca\xea\x19\x15\x9d\x8b\x74\x8f\xc5\x0c\x3b\x4e\x58\x17\x20\x9b\xb0\x33\xde\xaf\x07\xf1\x83\x01\xbe\xbf\x1e\xc8\x38\x2b\x49\x57\x05\xac\xc1\x21\x66\xb4\x13\x72\x8c\x41\x5c\x69\x81\x8d\x39\x34\x3d\x4c\xc5\x55\xd0\x59\x5b\x84\x27\x42\x73\xf6\xb4\x60\xc3\xf5\x4c\x89\x97\x08\x83\x15\x1c\xaf\x15\xb4\x54\x02\x11\xaa\x16\x8f\x93\xae\x65\xeb\x8a\xd4\x3f\xda\x67\x20\x45\x8c\x70\xc1\xa4\xef\x4b\x0b\x98\xd7\xc1\xa1\x4b\x05\x15\x13\xee\x02\xf8\x0a\x1e\x87\x37\x73\x01\x16\x07\xb0\x15\x52\x60\xd4\x94\xae\x87\xca\x18\x5f\xff\x9a\x20\x3c\x89\xdb\x86\x03\x12\xe1\xa8\x5a\xcb\x3b\xa2\xe7\x96\xa9\x33\x07\x27\xf3\xed\x6c\xe2\x6f\x69\x43\xa7\x0e\x9f\xbd\x87\x90\xf5\x8e\x7a\xba\xc0\x69\x33\xb0\xe6\x6d\xf6\x8a\x41\x6f\x78\x72\xe0\xd1\xea\xc3\x58\x60\xdf\x8c\xb2\xda\x46\xab\xd8\x3b\xef\x5f\xef\xfc\x6d\xad\x6f\xdf\xed\xb9\x9a\x02\xa4\x14\xc2\x87\xa8\xfe\x61\x61\xd1\x78\x45\x1c\x9a\x74\xb9\xce\x30\xd4\x43\xfd\x13\x85\xfe\x26\x36\xfb\xf3\xec\xab\x7f\x5d\x1c\x7c\xd6\xdb\x42\x5d\x9c\xc7\x23\xde\xb2\xfe\x02\x3a\x7b\x6b\x55\x98\xc3\x03\x4d\x25\xd3\xbb\x84\x39\x49\x2e\x7d\x7d\xbb\x35\x8f\x61\x83\x39\xfe\x01\xc4\x4c\xa6\x73\xa3\xc5\x01\xc4\x73\x77\x65\x6d\xda\x0b\x71\x9e\xf3\x15\x83\xcd\xdd\x87\x99\x92\x61\xee\x25\xd2\x2b\xe5\xda\x15\xc6\xb6\x6e\x9d\x5a\x7d\xc6\x20\xdd\xe9\xa3\x7e\x8f\xff\x6b\xf8\x93\x27\xfe\x0d\x9f\xf4\x34\xb9\xeb\xfa\x61\x5e\xd4\x77\xcf\xf3\xb7\x86\x7c\x2c\x90\xee\x5d\x9f\x16\xdd\xb7\x7a\xeb\xe5\x53\xa3\xf3\xe1\xbf\x3d\x26\x18\x1e\xe0\x7b\x26\x97\x1b\xfd\x67\xfb\xcf\x12\x9b\x97\xcd\x9f\x01\x00\x00\xff\xff\x0c\x18\x50\x44\xa5\x35\x00\x00")

func dataConfig_schema_v32JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v32Json,
		"data/config_schema_v3.2.json",
	)
}

func dataConfig_schema_v32Json() (*asset, error) {
	bytes, err := dataConfig_schema_v32JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/config_schema_v3.2.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"data/config_schema_v3.0.json": dataConfig_schema_v30Json,
	"data/config_schema_v3.1.json": dataConfig_schema_v31Json,
	"data/config_schema_v3.2.json": dataConfig_schema_v32Json,
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
		"config_schema_v3.0.json": &bintree{dataConfig_schema_v30Json, map[string]*bintree{}},
		"config_schema_v3.1.json": &bintree{dataConfig_schema_v31Json, map[string]*bintree{}},
		"config_schema_v3.2.json": &bintree{dataConfig_schema_v32Json, map[string]*bintree{}},
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
