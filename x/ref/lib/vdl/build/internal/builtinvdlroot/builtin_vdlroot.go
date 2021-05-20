// Copyright 2020 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by go-bindata.
// sources:
// v23/vdlroot/math/math.vdl
// v23/vdlroot/math/vdl.config
// v23/vdlroot/signature/signature.vdl
// v23/vdlroot/time/time.vdl
// v23/vdlroot/time/vdl.config
// v23/vdlroot/vdltool/config.vdl
// DO NOT EDIT!

package builtinvdlroot

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

var _mathMathVdl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\xc1\x6a\xf3\x30\x10\x04\xe0\x73\xf4\x14\xf3\x02\xb1\xff\x38\xc6\xfc\xd7\x34\xcd\x21\x50\x7a\x68\xda\xde\xd7\xf2\xda\x12\x95\x25\x23\xad\x43\x4d\xc9\xbb\x17\xa5\x50\x72\x28\x94\x1e\x07\x76\x67\xbe\xb2\xc4\x3e\x4c\x4b\xb4\x83\x11\x54\xff\x36\x0d\x9e\x0d\xe3\x95\x3c\x75\x76\x1e\xb1\x9b\xc5\x84\x98\x0a\xec\x9c\xc3\xf5\x28\x21\x72\xe2\x78\xe6\xae\x50\x65\x89\x97\xc4\x08\x3d\xc4\xd8\x84\x14\xe6\xa8\x19\x3a\x74\x0c\x9b\x30\x84\x33\x47\xcf\x1d\xda\x05\x84\xbb\xd3\xfd\x3a\xc9\xe2\x38\x7f\x39\xab\xd9\x27\x86\x18\x12\x68\xf2\x68\x19\x7d\x98\x7d\x07\xeb\x21\x86\xf1\x70\xdc\x1f\x1e\x4f\x07\xf4\xd6\x71\xa1\xd4\x44\xfa\x8d\x06\xc6\x48\x62\x94\xba\x9a\xc7\xc9\xf1\x7b\x53\xe7\x21\x82\xfe\x8a\xf0\xf3\xd8\x72\xbc\xc6\x90\xb8\xcb\xb2\x6d\xb5\x6e\xad\x20\x32\x39\x50\x1e\x18\x69\xb0\x9e\xe2\x82\x89\xa2\xa4\x42\xc9\x32\xf1\x4d\x5f\x92\x38\x6b\xc1\x87\x5a\x3d\xe5\x97\xde\x05\x92\x6d\xa5\x56\xc7\x91\x86\xef\x74\xb9\x45\x6c\xaa\xff\xbf\x2a\x9a\xfa\x0f\x8a\x5c\xf8\x13\xa3\xa9\x6f\x19\x4d\xad\x2e\xea\x33\x00\x00\xff\xff\xce\xca\xd1\x3a\xbf\x01\x00\x00")

func mathMathVdlBytes() ([]byte, error) {
	return bindataRead(
		_mathMathVdl,
		"math/math.vdl",
	)
}

func mathMathVdl() (*asset, error) {
	bytes, err := mathMathVdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "math/math.vdl", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mathVdlConfig = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8d\x41\x4b\xc4\x30\x10\x85\xcf\xc9\xaf\x78\xf4\xbc\xb4\x28\x22\x12\xf0\xb4\x07\x0f\xe2\x9e\x2a\x82\xb7\x98\x4c\xd3\xc1\x34\xa9\xd9\x74\x51\x96\xfe\x77\x49\xa3\x7b\xde\xcb\x3b\x7c\xf3\xe6\x7b\x5d\x87\x7e\x24\x9c\xac\x6f\x4d\x0c\x03\x3b\x0c\xec\x09\x43\x4c\xc8\x23\x61\xd2\x79\xc4\xac\xcd\xa7\x76\x84\x5a\x58\x12\x1d\xc1\x96\xe3\xa4\x33\x1b\x04\x9d\xf9\x44\xc8\x3f\x33\x1d\x65\xd7\x81\x03\x48\x9b\x11\x8e\x02\x25\x9d\xc9\xc2\xeb\xe0\x16\xed\xa8\x95\x7f\x0b\x8f\x65\x2e\xc7\xe8\xdb\xfd\x06\xce\x52\x3c\x45\x85\xb3\x14\xe2\x8d\x13\xf5\xf1\xb0\x39\xfb\xa2\xac\x58\x34\xfb\x38\xcd\x9e\xbe\xef\xef\x1a\x85\x8a\xc4\x33\x07\xab\x70\x58\xa6\x0f\x4a\xbb\x8d\x94\x0f\x85\xc6\x5c\xba\x15\xbf\x53\x2a\xfa\x97\x68\x49\xe1\x35\xf0\xd7\x42\xeb\x76\xa9\xf9\xef\xbe\xb9\x7d\x68\xd4\x95\xee\xd2\xbd\x46\x5e\x62\xdd\xc9\x55\xfe\x06\x00\x00\xff\xff\x03\xf0\x4f\xb5\x68\x01\x00\x00")

func mathVdlConfigBytes() ([]byte, error) {
	return bindataRead(
		_mathVdlConfig,
		"math/vdl.config",
	)
}

func mathVdlConfig() (*asset, error) {
	bytes, err := mathVdlConfigBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "math/vdl.config", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _signatureSignatureVdl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\x41\x8f\xd3\x3c\x14\x3c\x37\xbf\x62\x8e\xdf\x77\x20\x01\x24\xce\xa8\xec\xf6\x50\x09\xda\x4a\x2d\x5c\xaa\x1e\xdc\xf8\xc5\x31\x9b\x3c\x47\x7e\x2f\x2b\x55\x68\xff\x3b\x72\x12\xda\xd2\xa2\x15\xe2\x66\x3f\xcf\xcc\x9b\x99\x28\x45\x81\x87\xd0\x9d\xa2\x77\xb5\xe2\xfd\xdb\x77\x1f\xb0\xab\x09\xdf\x0c\x1b\xeb\xfb\x16\xf3\x5e\xeb\x10\x25\xc7\xbc\x69\x30\x80\x04\x91\x84\xe2\x33\xd9\x3c\x2b\x0a\x7c\x15\x42\xa8\xa0\xb5\x17\x48\xe8\x63\x49\x28\x83\x25\x78\x81\x0b\xcf\x14\x99\x2c\x8e\x27\x18\x7c\xda\x3e\xbe\x11\x3d\x35\x94\x58\x8d\x2f\x89\x85\xa0\xb5\x51\x94\x86\x71\x24\x54\xa1\x67\x0b\xcf\xd0\x9a\xf0\x79\xf9\xb0\x58\x6d\x17\xa8\x7c\x43\x79\x96\x28\x1b\x53\x3e\x19\x47\x10\xef\xd8\x68\x1f\x09\x96\x2a\xcf\x24\xd0\x53\x47\xc9\x55\x97\x8c\xb1\x7a\x76\xf0\xac\x14\x2b\x53\x12\x0c\x5b\xb4\xa4\x75\xb0\x17\xa6\xe4\x59\x77\xab\x36\xec\x58\x9e\x69\x96\xa4\x8c\xfe\x98\xd4\xeb\xeb\x9d\xa1\x82\xe1\x8b\x7c\x9e\xa5\xe5\x57\x3c\xd1\xd8\x97\x8a\x1f\xd9\x6c\x65\x5a\x02\x90\x26\x9e\x5d\x36\xdb\x3c\xb9\x8d\xd1\xfa\x7c\x7f\x0c\x25\xae\xdf\x17\xed\x91\xac\x00\xfb\xc3\x70\x02\x8a\x02\xab\x00\xe9\xa8\xf4\xa6\x41\x88\x96\x12\x30\xcf\x66\x5f\x86\x38\x82\xfd\x61\x3c\x25\xe4\x3a\x3d\x8f\x55\x4f\x69\xd9\xb4\x94\x67\x2f\x43\xae\x51\xf1\xf5\x4c\x94\x30\x96\xec\x5d\xb8\x91\xfc\xcf\xc1\x46\x07\x93\xd3\xbf\xac\x75\xca\x30\x19\x98\xb8\xf7\x0e\x70\xb7\xf3\x32\x59\xf2\x3c\x3a\x49\x93\xfd\x61\x1e\x1d\x86\xaf\xdb\xf5\x0a\x13\x5d\xdf\x12\xab\x64\xb3\x75\xaf\x13\xe8\x8c\x59\xf7\x7a\x03\x5a\xf2\x56\x23\x99\x16\xf8\x98\x30\x17\x21\x19\xc7\xff\x85\x4e\x7d\x60\xd3\xfc\x3f\x08\x4e\xe0\x33\x76\x12\xfc\x03\x78\x67\x46\x7f\x69\xbb\xe1\x13\x2e\x2d\xa9\x71\x32\xf5\x96\x64\x5e\x2b\x0d\xe2\xd9\x35\x74\x36\x3c\x55\x96\x68\x37\x7d\xfd\x56\xd5\xaf\xcb\x2e\x81\x13\x23\x1c\xbf\x53\xa9\xc9\xc3\x30\x1a\xfe\xe8\x6b\xd5\x97\xec\x67\x00\x00\x00\xff\xff\x7a\xf5\x7c\x0e\x29\x04\x00\x00")

func signatureSignatureVdlBytes() ([]byte, error) {
	return bindataRead(
		_signatureSignatureVdl,
		"signature/signature.vdl",
	)
}

func signatureSignatureVdl() (*asset, error) {
	bytes, err := signatureSignatureVdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "signature/signature.vdl", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _timeTimeVdl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x58\xdf\x73\x1b\xb7\x11\x7e\x16\xff\x8a\x1d\xf5\x21\xf6\x98\x3a\x52\x6e\x9a\x56\xf6\x93\x2b\x27\xae\x3b\xa9\x93\xa9\xe4\x76\xa6\x2f\x9c\x25\x6e\x79\xb7\x23\x1c\x70\x05\x70\xa4\x98\x4e\xfe\xf7\xce\x2e\x0e\xc7\x1f\x56\x5d\x25\xa3\x3c\x90\xc4\x2e\xf6\xc7\xf7\x7d\xbb\xf0\x62\x01\xb7\xbe\xdf\x07\x6e\xda\x04\xaf\x97\xd7\x7f\x80\xfb\x96\xe0\x1f\xe8\xb0\xe6\xa1\x83\x77\x43\x6a\x7d\x88\x15\xbc\xb3\x16\xf4\x50\x84\x40\x91\xc2\x96\xea\x6a\xb6\x58\xc0\xe7\x48\xe0\x37\x90\x5a\x8e\x10\xfd\x10\x0c\x81\xf1\x35\x01\x47\x68\xfc\x96\x82\xa3\x1a\xd6\x7b\x40\xf8\xf3\xdd\xfb\xab\x98\xf6\x96\xc4\xca\xb2\x21\x17\x09\x52\x8b\x09\x0c\x3a\x58\x13\x6c\xfc\xe0\x6a\x60\x07\xa9\x25\xf8\xf1\xe3\xed\xf7\x9f\xee\xbe\x87\x0d\x5b\xaa\x66\x62\xf2\x33\x9a\x07\x6c\x08\x12\x77\x04\x35\x6d\xd8\x51\x84\x98\xd0\xd5\x18\x6a\x08\xd4\x4b\x58\x2e\x61\x62\xef\xa2\x84\x84\xeb\xe8\xed\x90\x08\xd0\xc9\xef\x16\x13\x6f\xb3\x79\x94\xc8\xc5\xa7\xa4\x7a\x6e\x59\x53\x34\x81\xd7\x12\x36\x59\xbf\x03\x0c\x72\xe6\xdf\x03\x07\xaa\x21\x79\xe8\x83\xdf\x72\x4d\xb0\xe3\xa0\xa9\x18\xdf\xf5\x98\x78\xcd\x96\xd3\x1e\xd6\x94\x76\x44\x0e\x6a\xde\x6c\x28\x90\x4b\x72\xbe\x09\xd8\x75\xec\x1a\x20\xb7\xe5\xe0\x5d\x47\x2e\xc5\x0a\xe0\x03\x39\x0a\x98\xa8\xce\x25\xdb\xf8\x20\x0e\x0f\xa6\xc7\xc7\x21\xed\x7b\x36\x68\xed\x7e\x8a\x00\x87\xe4\x3b\x4c\x6c\xc0\x78\xb7\xa5\x10\x35\x7e\x76\xc9\x83\xd3\x64\xc5\xdb\x59\x7a\x73\xb9\x05\x22\x77\xbd\xa5\x00\x5c\xf3\xe8\x61\x88\xd8\x50\x35\xeb\x8f\x8a\xac\x55\xff\x7b\x31\x97\xf0\x6b\x4c\x14\xb5\x9c\x5a\x45\xe9\x31\x6a\xfe\x96\x8d\xa6\x91\x7c\xcf\xa6\x02\xb8\xf3\xdd\x88\x0a\x01\x42\x1c\x28\xbe\x11\x67\xd7\x2f\xb5\xe2\x84\x21\xb5\xdf\x44\x90\xd4\xe5\x54\xf0\x39\x38\xf1\xe7\x7c\x82\x40\xcd\x60\x31\x28\xbc\x00\xa0\x4d\xa9\x7f\xb3\x58\x90\xab\x76\xfc\xc0\x3d\xd5\x8c\x95\x0f\xcd\x42\x3e\x2d\x7e\x24\xec\x57\x91\x8c\x77\xf5\x73\x8e\xdf\x7a\x1f\x6a\x76\x12\xec\xea\xb3\x63\x29\x1a\xda\xd5\xbd\x64\xfb\x0c\xeb\x8f\x2e\x51\x70\x1a\x2b\xda\xd5\xbb\xe4\x3b\x36\x93\xf1\x6b\x4d\x2e\x90\x02\xe6\xd0\xc3\x82\xd0\xa8\x85\x57\xdc\x4a\x29\x0d\x5a\x92\xef\xe3\xb3\xb2\x5c\xae\x5e\xec\x09\xc3\xcb\xe7\x9c\xfd\xeb\x60\x19\xdd\xaa\x5c\xf0\x1c\x93\x0f\x81\x1a\x1f\x7e\xab\xd5\xcf\xc1\x5b\xea\x13\x9b\xd5\x93\xf6\x85\x5f\x9d\x8f\x09\xb8\xeb\x7d\x48\xe8\x12\x98\xd6\xb3\x51\x79\xd8\xb5\x94\x5a\x0a\xc2\xa9\x09\xa4\x99\xdd\x98\x71\x35\xb8\xa4\x20\x62\xf3\x10\xc5\x5b\x64\x67\x84\xcd\x40\xbd\x37\x2d\xbc\x90\x1f\xd4\xe0\xe5\x1c\x7c\xc8\x56\x02\xd1\x85\x3a\x91\x72\x23\xf4\x18\x12\x1b\x81\x13\x1c\xe7\xf6\xa2\x7c\xc8\xf6\x15\xc0\x7d\x71\x06\x1d\xf6\x51\x08\x34\x04\x65\x5b\xf2\xd0\x62\xa8\x77\xd2\x56\x63\xbd\x79\x88\x73\x89\xde\xb4\x84\x3d\x24\x7f\xc2\xb1\xb9\x92\x03\xad\xf5\xbb\x08\x43\x94\x46\x67\xaa\x41\x1c\xd6\x29\xa0\x51\x98\x4b\x64\xc2\x9a\x21\x93\x6a\x08\x99\x9a\x8a\x84\xdb\xe3\xb8\xe4\x9e\x03\xed\x3b\x1f\x08\x22\x75\xe8\x52\xf9\x86\x50\xc0\xb4\x19\x6c\x05\x70\x7b\xa4\x01\xa3\x0a\x89\xc3\xa9\x48\x1a\xda\x49\xda\x8a\xd4\x63\xfa\xae\xf7\xca\x58\x5c\xfb\x6d\xe1\x6d\x51\xca\xaf\xf3\xe2\xee\xa7\xd5\x9f\xbe\x5b\x5e\xff\xdf\x83\x77\xfb\x98\xa8\x5b\x4d\xea\x22\xdc\x39\x54\x4f\xa4\xe5\xa0\xda\xbd\x67\x97\x74\x1c\xc8\xa1\x1d\xa7\x16\x06\x29\x38\x38\x74\x3e\x33\x1e\xfa\x40\x86\x25\xe5\x49\xd0\xc7\xaa\x4d\x3e\xa9\x16\x5c\x48\x56\xa5\xce\xb0\xa6\x8d\x54\x52\xe0\xb1\x49\x24\x20\xd9\xf0\x23\xd5\x19\x55\x82\x84\x96\xe0\x17\x0a\xfe\xa9\x00\xc5\x51\x46\xdf\x72\xb9\xbc\xbe\xd2\xbf\xfb\xe5\xf2\x8d\xfe\x55\xcb\xf2\xdf\xbf\xd4\x0d\x0b\x06\x48\x8d\xc4\x57\x5f\xf8\x02\x13\x5f\xa6\x7e\xbc\x55\xcf\x53\x77\xc2\x20\x23\xcc\x29\xd6\x1f\xd1\x24\xf8\x76\xb9\x04\x51\x00\x30\x7b\x23\x03\x71\xb1\x00\x51\x3e\xc8\x75\x88\xda\xc9\xcb\xd8\x11\x06\xaa\x2f\xe7\x40\x2e\x0e\x41\xc0\xa5\xf3\xd5\x79\xb0\x87\xc3\x90\x70\x6d\xb5\x48\x8e\x0c\xc5\x88\x61\x2f\xee\x04\x94\x2c\xfa\xd6\x07\xca\x72\x7c\x18\x92\xac\x52\x1f\xb9\x63\xe1\x51\xf2\xf0\xc1\x6b\x57\x2a\x29\xcf\x1c\xd6\x43\x02\xb4\x3b\xdc\xc7\x32\xbd\x3f\xdf\xdf\x82\xf5\x66\x72\x53\x50\xd1\x78\x8b\xae\x51\x3c\xf4\x0f\x8d\xf2\x74\xf1\xbb\x51\x40\xff\xc7\x55\x3a\xda\x5c\x96\x5c\xb8\x1c\x1c\x3f\xea\xd5\x97\xf9\x5a\x85\xc5\xa1\x29\x79\x2d\xa8\x01\x93\x78\xd3\x82\x5d\xcb\x98\x51\x99\x69\xd1\x8d\x5f\xdd\xfc\x71\x59\x1a\x34\x72\x35\x8d\x3d\xcf\x0d\x4f\x5e\xf6\x91\x21\x66\xf0\xa0\xf8\x1a\xf5\x00\x04\x06\x61\x3f\x1f\xb5\x88\x13\x70\x74\xdf\x24\x40\xd8\xa2\xe5\x3c\x17\xb5\x94\x1d\x3a\x19\xd4\xc2\x78\x83\x16\xb0\xcf\x24\x9b\x78\xfe\x55\x9a\x7c\x76\xfc\x98\x49\x92\xf6\x3d\xe5\x98\xde\x8f\xf0\x55\xde\x94\x0f\x5f\x40\xd3\x62\x2f\x51\x1f\x61\x3d\xaf\x23\x69\xe7\x33\xa1\x62\x61\xd4\x5c\x6b\x27\xce\xbe\xc6\x2a\xbd\x7f\xba\x2d\xa6\x30\x98\x04\xff\x99\x5d\x2c\x16\x70\x37\x42\xef\x2c\x84\x82\xc8\x11\x09\x25\x92\x91\x57\x01\x5d\x93\x09\xea\x87\xa6\xb5\x7b\xf5\xf4\x6a\x71\xf5\xfa\x66\x09\x6b\xb6\x56\x6e\x91\x1e\xc5\x39\x58\x0c\x4d\x69\x9b\xe6\x16\x13\x77\x2a\x53\xb2\xa0\x8c\xcb\xc5\x90\x67\x38\x55\xb3\x8b\xbb\xe9\xe2\xf4\xdd\xb7\xea\xf7\x93\xa4\x74\x1e\xdf\x66\x54\xe1\xbc\x22\x16\x4e\x08\x4b\x0e\x05\x08\xa4\x1a\x94\xc3\xfe\xdb\x10\x93\xba\x5b\x53\x49\x8a\x9d\xb1\x43\x94\x6d\x32\xe7\x53\xaa\xfc\x6a\x71\x75\x73\x73\x33\x1f\xff\xaf\xc4\x4a\x2d\x3f\x3a\x70\x3e\x74\x68\xf9\x17\xaa\x05\x1e\xdd\xfc\x20\xfb\x60\x29\xc6\x9c\xa6\x77\x34\x05\x14\xe8\x44\xc4\x14\xe8\xcb\x93\xc2\x8b\x9e\xbf\x5a\x5c\x69\x96\x15\xc0\x0f\xb2\x5e\x4c\x4e\x8f\x5c\x09\x1a\x7d\xa0\x79\x6e\x0f\x37\x4e\x32\x57\x2b\x75\xd7\x0d\x31\x41\x87\xc9\xb4\xc5\xb3\x4e\xd3\x35\xc1\xb2\x9a\x5d\xe4\x1a\xb2\x4b\xbf\x7f\x3d\xfb\x55\xb1\xf7\x4f\x0e\xf4\x9e\xb0\xb6\xec\xbe\x90\xc6\xba\x7c\xaf\xc3\xd7\x81\xef\x29\x07\x34\x97\x69\x1f\x48\x0f\x4d\x5f\x02\xeb\x58\xa7\xc7\x9e\x4c\xde\x1e\x41\xb6\xa3\xd8\x16\x89\x3e\x76\x39\x02\x48\x24\xca\xd5\x54\xe7\x85\x55\x90\x94\xe7\xbc\xb1\x4c\x4e\x59\x9f\x3c\x44\x12\x46\x4e\xc1\x64\x21\x3d\x8a\x25\xe2\x5e\x0b\xd4\xb1\x93\x49\xb3\x09\xbe\x83\x4b\xe7\x77\x97\x79\x7e\x47\x12\x95\xcc\xa2\x3d\x39\x49\x5e\xd1\x12\xb6\x14\xc6\x50\xf2\x07\x09\xe1\x89\x0c\x4e\x13\xcd\xf9\x68\x74\xc7\x29\x8d\xaa\xf7\x93\x13\xd7\xec\x1a\x2b\xbf\x6d\xd9\xd0\x3c\x2b\xcb\xb8\x3c\xc4\x74\xba\x23\x1d\xa5\x86\x27\x33\x33\xcf\xfb\x8e\xde\x4a\xb5\x33\x54\x55\x90\x9c\xdf\x41\x20\x34\x2d\x9d\xb6\x69\x7e\xda\xb4\x16\x35\x15\x79\xe8\xa8\x46\xfd\xc5\xef\x48\x32\x54\x67\x52\x14\x99\x28\x47\x97\x4f\x6f\x1d\x8d\x39\x66\x8c\xf6\x3e\x89\x5c\xa3\xcd\xcb\x12\xc4\x07\xda\x49\x3a\xe2\x90\x85\x74\x89\x5c\x5e\x65\x82\x5f\x0f\xe7\xa9\x9d\xc4\x93\x17\xba\x22\x3e\x87\x2e\xe5\xfa\xeb\x46\x28\x58\x28\xcb\x14\x9d\xcd\x79\x35\xe0\xa4\x6f\x0a\xf9\xec\x37\xa5\xc9\xbb\x96\x6d\x46\x57\x20\x43\xa2\x20\xe5\x21\x37\x79\x99\xd6\x90\x03\xa4\x8b\x3b\xbf\x73\xe7\x2e\x4f\x46\xe5\xe9\x8b\x0b\x6a\x4f\x79\x44\x98\xbc\xd1\x0a\x5e\xfb\xe0\x7b\x6c\xc6\xdf\xc9\xe2\x3e\x4f\x33\x39\x9a\x67\xf7\x88\x15\x3c\xab\x88\x8e\xa4\x71\x4b\x2b\xa1\xe7\x57\x93\x3c\xaa\x2c\x1f\xeb\x65\x31\xd2\xc2\x17\xd6\x8c\xb6\x99\x2d\x15\x88\x32\xe9\xa4\x32\xbe\xeb\x44\xe0\x0d\x39\x0c\xec\x73\x0d\xbe\x88\x72\x24\x6c\xec\xd0\xda\xfc\xee\x1d\xdf\xc4\x7a\xfa\xc9\xc6\x77\xf8\x90\xf7\x10\x7d\x2b\x8e\xcb\xf0\x7a\x50\x38\xd0\x66\x43\x46\x1f\xe5\xd8\xf7\xc1\xa3\x69\x4b\x1d\x4f\x34\xe6\xb0\x01\xb7\x8a\x88\xfc\xb6\x3d\x2f\xb3\x9c\xa0\x1a\x0e\x56\x52\x3c\xb9\x54\xb7\xe0\xa9\x9d\x79\x63\xd9\xb5\x6c\xda\xc3\xfb\x39\xbf\xaa\x29\x88\x32\x97\x29\xa6\xc0\xca\xff\x72\x70\x0c\x91\xc3\x8e\x7d\x84\xc7\x3c\x26\x4f\x82\x3e\x19\x95\x3f\x04\xdf\x7d\x52\x0e\x3e\xad\x96\x5f\x01\xfa\x3b\xd9\x3f\xc4\x47\xec\xc9\x30\xda\x2b\x83\x71\xe4\xed\xf2\x60\xc2\xae\xd6\x15\x3e\x4e\x98\x09\x34\xe2\xa2\x5c\xf2\x16\xb8\xa2\x4a\x65\xed\xe2\x58\xd7\x38\xc2\x25\x3b\x51\xad\x44\x97\xd5\xec\xa2\xc4\x3a\xed\x1b\xbf\xce\xfe\x1b\x00\x00\xff\xff\xec\x04\xfd\xa5\x23\x12\x00\x00")

func timeTimeVdlBytes() ([]byte, error) {
	return bindataRead(
		_timeTimeVdl,
		"time/time.vdl",
	)
}

func timeTimeVdl() (*asset, error) {
	bytes, err := timeTimeVdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "time/time.vdl", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _timeVdlConfig = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x4d\x8b\xdb\x40\x0c\x3d\xdb\xbf\x42\xf8\xb4\x0b\x61\x0c\xed\xcd\xd0\x43\xd9\xc0\x92\x7e\x78\x0b\x49\x5b\xe8\x4d\xf5\x28\xce\xb4\xf6\xc8\x1d\xcb\x5e\x96\x90\xff\x5e\xe4\x49\x42\xb2\x4d\x0a\x4b\x7b\xb0\xb1\xa5\xa7\xf7\xe6\x49\x62\xf2\x1c\x56\x1b\x82\xd1\x36\xa6\x62\xbf\x76\x35\xac\x5d\x43\xb0\xe6\x00\xb2\x21\xe8\x05\xbd\xc5\x60\x41\x5c\x4b\xd0\x61\xf5\x13\x6b\x82\x88\x1c\x02\xf5\xe0\xac\xe3\x16\xc5\x55\xe0\x51\xdc\x48\x69\x9e\x83\x3c\x75\x9a\xf1\x40\x58\x6d\xa0\x26\x4f\x01\x85\x2c\x34\xe8\xeb\x01\x6b\x32\xe9\x5e\xea\x8d\xea\x0a\x73\x63\xee\xa6\xc0\x36\x4d\xee\xb9\x80\x6d\x9a\x24\x5f\x5d\xa0\x15\x97\x13\xe7\x4a\xf9\x62\x38\xc9\x73\xf8\x32\xff\x30\x1d\xc7\xcc\x87\x80\xe2\xd8\x83\xeb\x21\x50\x17\xa8\x27\xaf\x3a\xd8\xc3\x3d\x9f\x43\x8c\xd6\x66\x87\xbf\x6c\x4f\x96\xbc\x77\xde\x16\x00\x00\xe5\xd0\x7e\xa7\x30\x9b\x82\x2a\x37\x05\xb3\x33\x8a\x2c\x66\x17\x6d\xc7\x41\xf4\x38\xdb\x4f\x28\x9b\x22\xa2\xb2\x19\x94\xd8\xd2\xe1\x6f\xb7\x8b\xe0\x6f\x14\x78\xa2\xda\x7e\x64\x4b\x05\x7c\xf6\xee\xd7\x40\x31\x19\xdf\xa7\x7e\x56\xda\xe3\xeb\x5e\x34\x1d\x7d\xe8\xd7\x9f\x1e\x96\x12\x86\x4a\x2e\x7a\x98\x0a\xfe\xc3\xf9\xef\xd0\xb3\x77\x15\x36\x33\x58\xf4\x31\x99\x99\xf8\x75\x73\x9b\x5d\xf3\xa5\xb3\x9c\x13\xda\xc6\xf9\x4b\xfe\x14\x17\x97\x67\x3f\xb3\x3d\x34\x7a\x3d\x2d\x7e\x91\xe7\x63\xd1\x35\xdf\xa3\x71\x9c\x8f\xaf\x5e\xe7\xa3\x6d\x02\xb3\xe4\xff\xda\x88\x69\x40\x97\xba\xa1\x2f\x7d\xde\xe1\x88\x7f\x5f\xef\xd3\x15\xcd\x38\xd4\xe6\x07\x5b\x34\x17\xf6\xf0\xb0\x03\xcf\x41\x28\x74\x1c\xf6\x41\x76\xf9\xe8\xd6\xf2\x02\xdd\x72\xa9\x14\x0b\x2f\x14\x46\x6c\x9e\xe9\x95\x4b\xd5\x38\xe3\xd7\x3b\xe4\x61\xfe\x70\x23\x6c\xed\xe3\x6d\x01\x6f\xad\x3d\x0e\xf4\xa9\x23\xe8\x87\x4e\x7b\x3f\x5d\x2a\xda\x82\xbe\x0a\xae\x13\x93\xee\xd2\xdf\x01\x00\x00\xff\xff\x46\x95\x86\x66\x7e\x04\x00\x00")

func timeVdlConfigBytes() ([]byte, error) {
	return bindataRead(
		_timeVdlConfig,
		"time/vdl.config",
	)
}

func timeVdlConfig() (*asset, error) {
	bytes, err := timeVdlConfigBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "time/vdl.config", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vdltoolConfigVdl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x5a\x4f\x73\xdc\x36\xb2\x3f\xcf\x7c\x8a\x2e\xd5\xab\x8a\xed\x9a\x90\xef\xe5\xfd\x39\xf8\x55\x0e\x8e\x65\xab\x94\x75\x64\xaf\x25\x3b\x5b\x9b\xca\x01\x43\xf6\xcc\xc0\x22\x01\x06\x00\x47\x9a\xb8\xfc\xdd\xb7\xba\x1b\x00\xc1\x91\x6c\xc7\xd9\x9c\x92\x1c\xe2\x11\x01\xf4\xff\xfe\x75\xa3\xc9\xba\x86\xa7\x76\x38\x38\xbd\xdd\x05\xf8\xe6\x3f\xff\xeb\x7f\xe1\x6a\x87\xf0\x56\x19\xd5\xea\xb1\x87\x27\x63\xd8\x59\xe7\x2b\x78\xd2\x75\xc0\x9b\x3c\x38\xf4\xe8\xf6\xd8\x56\xcb\xba\x86\x37\x1e\xc1\x6e\x20\xec\xb4\x07\x6f\x47\xd7\x20\x34\xb6\x45\xd0\x1e\xb6\x76\x8f\xce\x60\x0b\xeb\x03\x28\xf8\xee\xf2\xf4\x6b\x1f\x0e\x1d\xd2\xa9\x4e\x37\x68\x3c\x42\xd8\xa9\x00\x8d\x32\xb0\x46\xd8\xd8\xd1\xb4\xa0\x0d\x84\x1d\xc2\x8b\xf3\xa7\xcf\x2e\x2e\x9f\xc1\x46\x77\x58\x2d\xe9\xc8\x2b\xd5\x5c\xab\x2d\xc2\xbe\xed\x82\xb5\x1d\xb4\xb8\xd1\x06\x3d\x84\xc3\x80\x1e\x46\x2f\x7c\xe8\xec\xbe\xed\x80\xb7\xe8\xe0\xb1\xdb\xac\x40\x9b\xa6\x1b\x5b\x6d\xb6\xb4\x4c\xb4\x36\xd6\xf5\x2a\x90\xdc\xfb\xb6\xab\x1a\x6b\x36\x7a\xcb\xac\x7c\x05\x17\x16\x06\x61\xe5\xc1\x86\x1d\x3a\x12\xd2\x24\xc2\x4c\xb7\xb1\x7d\xaf\x4c\x2b\x8a\x18\x14\x66\xca\xb4\xb0\xaf\xb4\xad\x6f\x6b\x87\x9b\xba\xd3\xeb\x7a\xdf\x76\x75\x55\x55\x13\x3d\xbf\xb3\x63\xd7\x92\xb0\xa0\x43\xb5\x1c\xe6\x2a\xb1\x9a\x4f\x45\x18\x3f\x60\xa3\x37\x9a\xf4\xdb\x91\x45\xe9\xe1\xe8\x54\xd0\xd6\x90\xf4\x33\x3d\x2b\x80\x2b\x32\xbf\x66\x63\xe8\x46\x75\xdd\x81\x48\x39\x1c\xc8\x55\x26\x20\x9b\xd5\x0e\x74\x5a\x75\x70\x32\x29\x7d\x22\x5a\xd3\x32\xaa\x66\xc7\x24\xa3\x17\xa3\x70\x15\xc0\x33\xd5\xec\x88\xde\x91\xad\x40\xf7\x43\xa7\x1b\x1d\xba\x03\xfd\xb4\x2e\x78\x09\x83\xe2\x64\xb5\xad\xe0\x60\x47\xe8\xd5\x01\x1c\x6e\xc8\x98\x36\x92\x62\xc9\xa3\xb6\xd1\xe9\x77\x04\x83\x1b\x1d\x76\x76\x0c\x80\xb7\x47\xac\xc8\x99\x89\xc8\x92\x42\x20\x1b\x2e\xb8\xb1\x09\xf0\x7e\xb9\xa8\x6b\x38\x43\xf3\x42\x99\xed\xc8\xc6\x77\xe8\x83\xd3\x4d\x10\x93\x7a\x64\xff\x73\xac\x6e\xd1\x60\xb4\x6d\x97\xb6\x57\x00\xe7\x9b\xbc\x53\x7b\xa6\x87\xfd\x10\x0e\x2b\x50\x5d\x07\x7e\x1c\x48\x0e\x6c\xa7\x23\xa0\x1c\xd2\x9a\xbd\xc1\x16\x82\xa5\x98\x8e\x94\x29\x57\x16\x33\x61\x3c\x86\x9f\x8a\x07\x3f\x2f\x99\x7e\xfa\xf3\xeb\xe8\xfe\x66\xee\x79\x4f\x54\x2c\xa4\xff\xce\xac\xe8\xbc\x5c\x7c\xaf\xf6\x2a\x3e\xa4\x9f\xe5\x63\xdf\x38\x3d\x04\x98\x7e\xa6\xc5\xcb\x1b\xbd\x09\x72\x86\x7f\xc6\xe7\x1f\x38\x0a\x0b\xd1\x00\xcd\xd8\xb3\x12\x62\xb8\x6b\x63\x6f\xcc\x27\xec\x26\xde\x38\x26\x40\x0e\x39\xb3\x22\x52\x29\x58\x94\x23\xb1\xb5\x77\xc2\x7f\x6b\xe1\x7e\x63\x24\x46\xf6\x3e\xc7\xff\xa8\x1d\x5e\xd9\x0b\x15\xf4\x1e\xaf\x18\x20\xe6\x09\xd5\xab\x61\xa0\x10\xda\x38\xdb\x83\x82\xb7\xa7\x2f\xe0\x46\x3b\x64\x30\x21\xd7\xe9\xe0\x81\xc4\xad\x6b\x30\x4c\x44\x56\x72\x46\x89\x04\x53\xde\x39\xe5\xb0\x3b\x08\x0a\x11\x10\xa0\xf2\x07\xa2\x93\x24\x46\x26\xa5\x4d\x63\x9d\xc3\x26\x74\x87\xff\x87\xd1\x93\x6d\xb4\x87\x66\x74\x0e\x0d\x85\x76\x8a\x50\x09\x9f\x8c\x1a\x0c\x92\x14\x5b\x53\x16\x30\xb9\x9b\x9d\x0e\xd8\x69\x2f\xe1\x55\xd7\x1f\xd1\x9c\x4e\xf6\xa8\x4c\x60\xe0\xf0\x0d\x1a\xe5\xb4\xf5\x70\xb3\x43\xd2\x98\xff\xaf\x3d\x28\x03\xba\xd5\xb6\x57\x41\x37\x49\x77\x56\x9a\x95\xd2\x86\xf2\xd8\xb1\xdf\x57\xb0\x1e\x03\xa7\xb5\x41\xd2\x17\x7c\x50\xa6\x55\xae\x65\x3b\xce\x6d\xc4\x3c\xc9\xb4\x4c\xaf\xb1\xfd\xa0\x82\x5e\xeb\x4e\x87\x43\xc2\x07\xf2\x07\x1d\x0c\xba\xcf\xa8\x93\xd1\xfd\x34\x81\x1e\x59\xf5\x8a\x76\x6c\xac\x13\xed\xc9\x5f\x33\x82\x22\xd7\x0d\xc2\x0d\x29\x4b\x64\x73\xfa\x49\xbc\x06\xcb\xd0\xcb\x49\x9d\x44\x4e\x9a\x16\xcc\x27\x63\x52\x29\xbc\xc6\x83\x14\x38\x8e\x1a\x06\xd9\x1d\x82\x51\x3d\xa6\xc7\x2c\x3c\x59\xea\x81\xba\x56\xe2\x80\xc3\x80\x0f\x57\xe4\xa0\x66\x07\xfd\xe8\x03\x53\x5b\x27\xb5\x72\x95\x23\xb8\x4d\x1a\x2b\xef\x6d\xa3\x59\x58\x42\xbd\xb4\x5e\x22\xee\x5c\xb0\x32\x05\xad\xa3\xf3\x63\x1f\xc3\x1b\x6f\x29\x2a\x4c\xc3\x22\x2a\x18\x94\x76\x02\x76\x66\x8f\xce\xb3\x5b\x46\xd3\x30\xa4\x44\xbf\xd0\x02\x83\xea\x1a\xc3\x0d\xa2\x88\xc7\x26\x26\xc3\x17\x39\xe0\x57\xfc\xe4\x46\x77\x1d\xa8\x31\x48\xbc\x50\xc5\x01\xfa\x3f\x53\x23\xa9\x5f\xe3\x96\x44\x70\x12\x88\x59\x23\x8f\x99\x33\x9b\xd0\x4f\x2a\x3d\x21\xf9\x63\x8d\x9e\x99\x37\x19\x94\x4c\xff\xdc\xda\xc7\xbc\x1b\x98\x0e\x6c\xac\x4d\xc1\xfe\xe0\x96\x56\x57\x60\xe0\x91\x3c\x78\x08\xe8\x5c\x0c\x96\x69\xfb\x73\x67\xfb\x7c\xe0\x51\x3c\x71\x74\xe0\x6e\x0e\xf5\x6a\xf8\x89\x72\xd3\x6c\x7f\x3e\xb3\xf4\x48\xf0\xfa\x92\x31\xe7\x4a\x6d\x4b\x7c\x51\xe6\xc0\xa8\x25\x78\x14\x68\x51\x8a\x81\xb4\x21\x93\x31\xa6\xf0\x64\x62\x5b\xcb\x2e\xad\xd8\xb9\x05\xe5\x18\x75\xd7\x78\x98\x5a\x9c\x44\xfc\x30\xe0\x57\x9e\xad\x55\x2d\x17\xf3\x33\x49\xde\x9f\x7e\x3e\xb3\x79\x25\xe3\x6c\x7e\x52\x0a\x0e\x5e\x9b\x6d\x87\x73\xe9\xd9\xdd\xc4\x73\xa3\xb1\x63\x60\x92\xb0\xd6\x81\x28\xc5\x8e\x66\x8d\xa0\x86\xa1\xd3\x84\x44\x11\x96\x0b\x0e\x19\x99\x9f\x33\x09\x11\x6c\xb9\x88\x6b\xf4\x3b\x89\xc5\x7e\x6e\x91\x6a\xc3\x3a\xc6\xf2\x99\x95\xe4\xd2\x46\xba\x36\x06\x83\xa3\x5c\x51\x39\x07\xb9\x27\xbd\x44\xbc\xaf\x15\x0b\xe8\x43\xab\x82\xaa\x63\x38\x13\x34\xe1\xad\xea\x87\x6e\x2a\x5b\x22\xc1\xac\x96\xfc\x4d\x53\x5b\x1a\x8b\x1f\xfd\xb6\x1b\xb8\x62\x56\x0b\x5e\x3a\xb3\xf4\x8f\xa4\x65\x0c\xd3\x52\xee\x08\x3a\xda\x1c\xa1\x11\xb5\xa5\x3e\xa0\x6a\x8f\x61\xa4\x5a\x2e\x92\x10\x6c\xa6\xba\x86\x7f\xa2\xb3\x85\xa3\x7e\xa5\x3f\xf7\xaa\x1b\xb9\x39\xe1\xb4\x25\x27\x35\x3b\x6c\xae\x25\x87\x77\x6a\xaf\xad\xab\x96\x0b\x3e\x79\x66\xe9\x1f\x91\x30\x86\x35\x1f\x98\x72\x01\xa8\x55\x77\xba\xc5\x09\xdd\x7c\x92\x2b\x5a\x6b\x02\x0f\x26\x94\x01\x44\x1a\xa5\xd1\x24\xf1\xda\x15\x9f\x6a\x71\xa3\xc6\x2e\x4c\xfb\xb8\x0e\x45\x92\xe4\xc9\x94\xc7\x29\xbd\x73\x1e\x33\xea\xa4\xa7\xab\xc4\xfe\xd1\x94\x8e\xf3\xc4\x4e\x3b\x8b\xc4\x66\x0a\x8f\xee\x90\xb8\x87\x42\x36\x47\x32\x76\x61\x92\xc2\xfe\xe7\xb1\xbd\x55\x52\x30\xc9\xb7\xa9\xe3\xe5\xae\x96\xa0\x36\xe7\xa6\xf0\xfc\x88\xa1\xc5\x76\x94\x06\x6c\xb8\xd8\x76\xf3\x83\xd5\xdd\xe2\xd4\xd9\x46\x4d\x15\x82\xbc\x22\xe8\x2b\xb6\x96\x2e\x80\x0e\x88\x30\x79\x23\x95\xf4\x5c\x5d\x53\xe0\x1c\x4a\x67\x4a\x64\xea\x1e\x2b\xaa\xa9\xc9\x15\x12\xfd\xef\xe3\x5f\xc0\x91\xff\x98\x1b\x44\xce\x87\x55\x5e\xa0\x6d\xbc\x70\x92\x69\x9c\x4c\xab\x14\x6c\xbc\xfa\xfe\x07\xdb\xe2\x63\x78\xaa\x8c\x35\x54\x24\x56\x70\xee\x65\xf1\xa4\x92\x5f\x0f\x1e\x9e\x7c\x98\x0e\x46\x33\x3f\x86\xf7\xef\x5f\xa9\xb0\x7b\x2c\xe4\x4f\x56\x70\xa1\x7a\x4c\x7f\x7d\xc8\x07\x3e\x2c\x17\xc9\x31\x04\x72\xf2\x3b\x43\x09\xe7\xe6\x1c\x4a\x52\xf6\xc6\xd4\xcc\x39\xcf\x5b\x53\x97\x2a\xba\x2e\x17\xdf\xd1\x8d\x6c\x71\x31\xf6\x6b\x74\xfc\x94\x83\xe1\x89\x73\xea\xb0\x5c\x5c\xd2\x05\x76\xb9\xf8\x41\x0d\xcb\xc5\x2b\xab\x4d\xa0\x3d\xe7\x1b\xd5\x60\xe6\x2f\xe2\x14\x12\xe4\xa0\x29\xa1\x2c\xcb\x10\xb7\xcf\x90\x87\x8c\x90\xe0\x24\xf9\x76\x50\x5c\x3c\x54\x80\xd1\xe8\x5f\x46\xea\x3a\x75\x8b\x26\x4c\xbd\xad\x30\xc1\xb6\x68\x67\x98\x50\x11\xcf\x64\xcf\xfb\x3a\x99\xc4\x24\x53\xe4\xa0\xa6\xd3\x15\xc0\xe9\xc8\x60\x76\x66\x05\x0a\x0c\x67\xb5\xf8\x42\x87\xd9\xe5\x13\xde\x8d\x5e\x1a\xb0\xb5\xf2\x98\x18\x10\x19\x69\xd1\xe8\x36\xb8\x96\x8b\x15\x85\xa7\xed\x31\xec\x92\x68\xad\xde\x70\x42\x05\xd0\x9b\x7b\xd5\x81\xd6\xa2\x37\x5f\x51\x37\x4b\x57\xad\x23\x79\xaa\xe5\x82\x95\x3b\xaa\x2b\x0c\x83\x33\x57\x14\x18\x3a\xc1\x65\x5d\xd3\xf6\xd7\xcf\xfe\xfe\xe6\xfc\xf5\xb3\xd3\xc7\xf0\x4c\xf3\x00\x80\xc2\x18\xbe\xfd\x16\xde\xb0\xc5\xc1\xba\x18\xc7\xa4\xb4\xc7\x50\x01\xfc\x88\xd2\x11\x19\x1b\x60\x40\x47\xee\x5d\xe6\xbb\x43\x1d\x5b\xd5\x84\x9f\xdc\x0f\x30\x58\x17\x52\x4c\x15\x48\xc0\x3e\xc7\x01\x33\x97\xa7\xf4\x53\xf0\xc8\x1f\x55\x04\x95\x20\x84\x4c\x69\x5b\x12\x31\xf7\x59\x1c\x2c\x0e\xc3\xe8\x88\xb3\x1b\x11\xf4\x86\x4d\x5b\x5e\x6f\xc4\x12\xb9\x77\xf7\xb9\x22\x4d\x12\x4e\xad\xda\xf9\x26\x4b\x10\x14\xa5\x5f\x2c\xc2\xad\x0d\xf0\xa0\x7a\xb8\x8a\xf1\x20\x1d\x69\xba\x0a\x47\x11\x49\x34\x11\xb2\xca\x9d\x6c\xec\x2c\x48\xfd\xbc\x1a\x25\x96\xa7\xd4\x42\x13\x8d\xb5\x4c\x3b\xce\xe6\x6d\x3d\x59\x9e\x69\x51\x0b\x22\x30\x97\xc5\x73\x3c\x28\x40\xb7\x56\x41\xf7\x24\x09\xde\x92\x8a\xe2\x86\x79\x81\x8b\xb5\xf7\x8e\x8e\x29\xde\x58\xd7\x42\xd5\x4f\x69\x9a\x6c\x7f\xb3\xb3\x5e\xec\xfc\x65\xea\x70\x33\x2d\xa5\x32\x93\x62\xc6\x4c\x2b\xb6\x69\xca\x6d\xc7\x1e\xf9\xaa\xa3\x42\xbc\xc2\x4d\xea\x7d\x52\xbb\xab\x97\xa7\x2f\x1f\x04\xdb\xb6\x37\x0f\x1f\x8b\x0b\x12\x1b\x0a\x5d\x3a\x9a\x23\x9c\x54\xd7\xd4\x23\xf5\x3c\x49\xaa\x96\x8b\x99\x71\xef\x92\x7b\xd2\xd2\x9d\xf0\x12\x03\xef\x8a\x71\x19\x7b\x46\x0f\x9d\xbe\x4e\xee\x11\x30\x48\x91\xa9\xa4\xaf\x28\x14\x60\xb5\x3c\xc6\x60\x2c\x03\xb1\x4c\x6b\xce\x8f\x39\xce\x3b\xec\x64\x5e\xb2\xd3\xc3\xec\x42\x33\xcf\xfb\x99\x81\x96\xf1\xba\xbb\xca\xfd\xee\x71\xf4\x17\xe9\xc9\x3c\x53\xc1\xa8\x6b\x78\x63\x64\x2a\x52\x0e\x19\x54\xb8\x87\x25\xc5\x12\x83\x44\xce\xb4\x19\x2b\xb9\x3e\x31\xbb\x62\xa6\x37\x75\x53\x45\x78\x1a\x99\xb5\xf5\x1c\x31\xd3\x55\x02\x3c\x61\x62\xba\xed\x46\x28\xfa\x28\x0e\xa5\xbb\x74\xea\x22\x4b\x38\x02\xb8\xba\x07\xa5\x88\x97\xc3\x5f\x46\x1e\x96\x24\x1d\xc5\x9d\x53\x6b\x73\x2f\x0c\x45\xac\x5c\x2e\xa2\xb1\x78\x7b\x6e\x0d\x3e\x6f\xba\xbb\xe1\x1c\x8d\x23\x37\xd8\x4c\xe8\x68\xfe\x50\x34\xd7\x13\x31\x89\x3c\x2d\xae\xa0\x65\x6b\xba\xc3\x91\x07\xee\x0e\x7b\xf2\xb4\x62\x2e\x55\x92\x28\x75\x42\xbf\x55\xac\x5f\x73\x0a\x48\x92\xcd\xe4\xa0\x1c\x6f\x84\x93\x07\xd5\x79\x1b\x93\x43\x7b\x19\xe3\x4c\x65\x92\x7a\xc4\x54\xfa\x3e\x17\x1c\x1b\xe5\x83\x04\x07\x8f\x82\x3e\xef\xf5\xc9\x1c\x9f\xf2\xfc\x6f\xf1\x7a\xf6\x74\x4c\x18\xae\xa4\xbf\xdf\xe9\x6c\xa9\x08\xac\x9f\xf7\xf7\xc7\x9d\xc7\x5e\x4b\x43\x26\xa6\xf7\x40\xa0\x6a\xa7\x3c\x77\x2f\xe9\xf2\x19\xfe\xef\x7f\x1e\x96\xbc\x3f\xe2\xd0\x2f\x76\xc2\xda\x52\x6b\x76\xdf\x2d\xee\xc8\xfe\x9c\x39\x64\xb5\x88\x7e\xd3\xa0\xb7\x30\xe3\x3b\xb5\x57\x9f\x9e\x97\x96\xc7\xfe\xe8\x89\xa9\x0c\x77\xff\xaa\x33\xd3\xac\xfd\x9f\x77\x6a\x6a\xdd\xb6\x7a\x67\x5b\x55\x11\x57\xa6\xf7\x27\x19\x99\x46\xf0\xb9\x6f\x40\x2a\xfd\xd8\x9a\x02\x59\x46\x9a\x71\x5c\xa7\x8d\x64\x5b\x21\x50\xf5\x99\xd1\x61\xec\x98\xa6\x98\x3b\x0c\xf8\x1a\x65\xd8\xf2\x05\xd9\xc6\x86\x94\x94\x63\x52\xfc\xf6\xa7\x4c\xb9\x38\x0f\x2c\xa3\xbb\xe0\xf4\xd9\xd8\x9e\xd8\xd1\x81\x34\x20\xde\x74\xfc\xee\x4c\x1b\xbe\xb0\xc1\x8d\x3a\x4c\x46\x2f\x25\x60\x26\x2b\xc0\x6a\x5b\xad\xa2\xe8\xe7\x26\xe0\x16\xdd\x94\x1c\x4d\x1a\x1a\xd2\xe6\x16\xde\xc6\x0d\x2c\x50\xd3\x29\x47\x71\x3e\xc9\xbf\x43\x43\x21\xbb\xe6\x30\x60\x9e\x77\xac\xbc\x12\xac\xdd\xa0\x0a\x63\x7a\x09\xe7\x61\xe4\x9b\x96\x0a\x41\x35\xbb\xf9\xd0\x8a\x9e\xa7\x18\xa4\xa2\xf7\x7c\x1a\x05\xae\x28\x17\xe4\x7d\x25\x1b\xfa\x49\xd3\xa0\xf7\x2f\xa8\xf0\xa6\x13\xd1\xff\xc1\xb2\x20\xc5\x86\xd8\x3a\x1a\x0e\xfb\xa3\xb5\x60\x81\xa0\xe0\xd8\x59\x32\xfb\x9e\xb6\x49\x02\xa4\x44\x9e\x04\xb6\x26\xbb\xa6\x55\x41\xfd\x9e\x84\xcb\x11\x54\xa6\x19\xac\x45\x84\x7f\x3b\xcd\x8e\xc3\xec\x6e\xdc\x4f\x75\xab\x7c\x13\x79\x54\xbd\xe2\xfb\xca\xcf\xd6\xb0\x39\x89\x54\xc9\x84\x45\xf1\x3e\xb3\xa0\xee\xf9\x85\xe7\xfd\x84\xe3\x3b\xf5\xe2\xe5\x35\xd1\xb9\xb0\x21\xf6\x3a\x2d\xfa\x41\x07\xc9\x8e\x92\xba\xbc\x53\xe7\x46\x1a\x14\x6c\xf5\x1e\x4d\x36\x5b\x6c\x24\xf2\x84\x9c\xdf\x7b\xa3\x3b\x70\xe1\x23\x3b\x0f\x4e\x35\x41\x37\x0c\x7b\x08\x3e\xf0\x00\x01\xa5\x90\x51\x8b\x7c\xc7\x89\xa2\x41\x6f\xdb\x51\xbe\xa2\x88\xb8\x35\x49\x2d\x37\x1f\xe4\xf0\xa6\x6a\xe8\xec\x80\x8e\x6a\xac\x69\xd1\x71\xbd\x49\x23\xa8\xb5\x1d\xa9\xf8\x90\x5d\xe2\x05\x89\x08\x0e\xce\xbe\xc3\x26\xf8\x7a\xe3\x54\x8f\x37\xd6\x5d\xfb\x5a\xd8\xf9\x34\x1c\xf9\x91\x2e\x31\xc2\x74\x7d\x90\xa8\xe1\xce\x45\xde\xd7\xd3\xb5\x15\x5b\x38\x61\x49\xe5\xe4\x49\xfc\xc8\xc3\x9a\xa0\xb4\xf1\xf0\xfd\x9b\xcb\xab\xf4\x1d\x46\xa9\xdc\x65\xa1\x1c\x44\x3c\x76\xd6\xf2\x5b\x7a\x2e\xa2\x14\xc2\xa9\xbe\x57\xb3\x8c\x25\x52\x9a\x77\xc1\x4e\xed\x27\x0c\x93\x6f\x1b\xc8\x37\x4c\xe0\x1f\x0c\xff\x51\xc7\x3a\x28\xb7\xc5\x90\x24\x7e\xb3\x46\xf7\xdc\xba\xa7\x2a\x78\xa2\xa6\x02\xd4\x6f\x3c\x3a\x5f\x2b\xe5\xac\xa9\xc7\x35\xdf\xa4\x1a\x15\x7c\xbd\x6f\xbb\x95\x64\x79\xe3\x50\x05\x56\xe4\x53\x9b\xeb\xc6\xf6\x55\xf9\xac\x30\x0e\x23\x06\xcb\xac\x83\x4c\xcd\xa2\x9d\x88\xe8\x49\x21\xd4\x49\x45\x96\x67\xae\x5c\x8f\x02\xb1\xe6\x57\x4d\x93\xa2\xf1\x68\x4a\x62\x0a\x5b\x4d\x9d\x93\x75\x07\x88\x9f\xab\xd0\x01\x3f\xae\xd3\x73\x1e\x1d\x51\xf4\xb8\x90\xdc\x50\xf0\x9c\xb9\x64\x05\x63\x17\x74\xaf\x02\xf5\x6c\x9d\xf4\xab\xcb\x38\x7e\xa6\x7e\x42\x77\xe8\x3e\xf6\x66\x90\x03\x06\xe5\xf3\x21\xe0\x51\x1f\xb9\x8b\x3f\xae\xf1\x29\x20\xd9\x43\x65\xf2\x9d\x8b\x43\x5b\x4b\x77\xa3\x14\x72\x6c\x00\x63\xd7\xb6\x3d\xe4\xba\xac\xd6\x9d\x94\xc3\x82\x50\xc6\x76\x32\x05\xab\xb1\x12\xf5\x5b\xf9\x1c\x85\x25\x41\xd3\xc2\x38\x08\xa6\xe1\x6d\x70\x0a\x3a\x6b\xb6\x54\x80\xbc\xaf\x87\xeb\x6d\x7c\x0d\xf2\xe0\xa9\xed\x0b\xf7\x5d\xa2\xdb\xeb\x06\xfd\x2b\x67\x37\xf2\x0e\xac\x78\x8f\x33\x5f\xe4\xd0\xfb\x8f\xb7\xa7\x2f\x5e\xbf\x7c\x79\x75\x37\x0c\xe2\xde\x7a\xe0\xcd\x0f\x53\x86\x45\xc5\x09\x24\x38\xc0\xc8\x5c\x3d\x19\x7f\xe8\xe6\x59\x32\x87\x9d\xc4\x47\x6c\xc4\xb1\x34\x74\xaa\xe1\xf0\x2c\x23\x4e\x62\x25\x26\x58\x67\xb7\x7c\x3b\x9d\xe0\x20\x66\xd6\x54\x0a\x93\xe9\x89\x17\xbf\x54\xdb\x7f\xf3\xdf\xec\x78\x9e\x75\xe1\x6d\x9c\xce\x06\xf1\x71\xfa\x70\xec\xa9\x75\x08\x19\x45\xa4\xc7\x44\x42\x3f\x71\x3f\x03\x52\x26\x97\x4d\x21\x5f\x54\xed\xc5\x6d\x25\xb9\x64\xd7\x89\x64\x2c\x05\x33\xa8\xff\xa3\xef\x33\xf2\x69\xca\x5f\xf6\x42\x33\xa9\xff\xe7\xbc\xd1\x5c\x5c\x9e\xd2\x7d\xdc\x3a\xb8\xb8\x24\x92\xd4\x85\xba\xbd\x4c\x28\xfe\x72\x77\x19\x01\x96\x2f\xbf\xcc\x7c\x58\xfe\x2b\x00\x00\xff\xff\x88\x99\xf3\xbd\x42\x2a\x00\x00")

func vdltoolConfigVdlBytes() ([]byte, error) {
	return bindataRead(
		_vdltoolConfigVdl,
		"vdltool/config.vdl",
	)
}

func vdltoolConfigVdl() (*asset, error) {
	bytes, err := vdltoolConfigVdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vdltool/config.vdl", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
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
	"math/math.vdl":           mathMathVdl,
	"math/vdl.config":         mathVdlConfig,
	"signature/signature.vdl": signatureSignatureVdl,
	"time/time.vdl":           timeTimeVdl,
	"time/vdl.config":         timeVdlConfig,
	"vdltool/config.vdl":      vdltoolConfigVdl,
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
	"math": {nil, map[string]*bintree{
		"math.vdl":   {mathMathVdl, map[string]*bintree{}},
		"vdl.config": {mathVdlConfig, map[string]*bintree{}},
	}},
	"signature": {nil, map[string]*bintree{
		"signature.vdl": {signatureSignatureVdl, map[string]*bintree{}},
	}},
	"time": {nil, map[string]*bintree{
		"time.vdl":   {timeTimeVdl, map[string]*bintree{}},
		"vdl.config": {timeVdlConfig, map[string]*bintree{}},
	}},
	"vdltool": {nil, map[string]*bintree{
		"config.vdl": {vdltoolConfigVdl, map[string]*bintree{}},
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
