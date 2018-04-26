// Code generated by go-bindata.
// sources:
// crdb/1517299952_init.down.sql
// crdb/1517299952_init.up.sql
// DO NOT EDIT!

// Copyright (c) 2018 Joyent, Inc.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.


package migrations

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

var __1517299952_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xc1\x4a\xc6\x30\x10\x84\xef\xff\x53\xe4\x3d\xfe\x93\xe2\x2f\x14\x04\xc5\xf6\xe0\x6d\x49\x63\x5a\xa2\xe9\x26\x6c\x37\x91\xfa\xf4\xd2\x5a\x6f\x4e\xe9\x31\xcc\xf0\xed\xcc\x90\x87\xd7\xe7\x17\xd3\xdd\xdd\x3f\xdd\x4c\xf3\x68\x6e\x6f\x4d\xdb\xb5\xa6\x4e\xd7\xcb\xbf\x82\x63\x2a\xfc\xee\x25\xda\x85\x42\x86\x26\x20\xcc\xa5\x67\xaf\x54\x39\x50\x8d\x16\xb9\x24\x15\xf5\x42\xbb\x79\x7b\x21\x9e\x77\x45\x82\x2e\x34\x4a\x2a\x79\xe5\xba\x73\x4e\x29\xf1\x1c\xf3\x54\xc4\xc0\xea\x65\xb0\x0e\x21\xd7\x5c\x78\xad\x83\xd4\xa9\xff\x20\x5d\x32\xe2\xfe\xa6\x38\xde\x1a\x9e\xb5\xce\xa5\xc2\x4a\x93\x85\x9b\x6d\x04\x9c\x1a\x81\xbf\x81\x30\x58\x17\xe2\xba\x2c\x7b\xfd\x4a\xf2\x49\x2a\x96\xe7\x80\x2e\xfc\xd9\x51\x79\x3f\x86\x84\xbe\x50\xcd\xa8\xd4\x5e\x1b\x0d\x2e\xe3\xf5\xf2\x13\x00\x00\xff\xff\xd0\x6d\x25\x7b\x12\x03\x00\x00")

func _1517299952_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1517299952_initDownSql,
		"1517299952_init.down.sql",
	)
}

func _1517299952_initDownSql() (*asset, error) {
	bytes, err := _1517299952_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1517299952_init.down.sql", size: 786, mode: os.FileMode(420), modTime: time.Unix(1519862399, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1517299952_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x7b\x5d\x73\x1b\x37\xb2\xf6\xbd\x7e\x45\x97\x6f\x48\xbe\x45\xd1\x49\xde\x64\x6b\xe3\x24\x5b\x87\x4b\xd3\x0e\x6b\x65\x4a\xa1\x28\x25\xeb\x8b\xc3\x02\x67\x40\x12\xd1\x0c\x30\x01\x30\xa4\xe8\x5f\x7f\xaa\x1b\xc0\x7c\x71\x38\xb2\x65\xe5\xca\xd6\xa0\xd1\x00\x1a\xfd\xf1\x74\xa3\x79\x79\x09\xf7\x37\x13\xd8\x31\x03\x0c\x36\xfc\x00\xb1\xd8\x6c\xb8\xe6\xd2\x42\x2c\x52\x2e\x8d\x50\xd2\xbc\xb9\xb8\xbc\xbc\xb8\xbc\x84\xff\x07\xe3\x38\x15\x52\x18\xab\x99\x15\x7b\x0e\x37\x09\x93\x1c\xfa\x51\x6e\xac\x4a\xb9\x06\xb5\x01\xbb\xe3\xc8\x71\xe0\xe8\x27\x4a\x5a\xad\x92\xa7\x09\xdf\x32\xcb\x02\x55\x6e\xb8\x36\x48\xc2\x20\xe6\x59\xa2\x8e\x3c\x06\xc3\xf5\x5e\x44\x7c\x70\x81\xe4\xd7\x7a\xcb\xa4\xf8\xc4\xac\x50\x12\xfa\xaf\xae\xf5\xf6\xd5\x00\x84\x01\x26\x81\xd5\x37\x18\x29\x69\xac\xce\x23\x3b\x02\x18\x4b\x50\xd5\x89\xc2\x80\x54\x16\xf9\x69\x4e\x47\x8e\x38\x5b\x27\x1c\xd6\x47\x60\xf2\x88\x53\x69\xeb\x4a\x43\x8c\x9b\xcb\x68\x73\x5c\x5a\x61\x8f\xa3\x8b\xc9\x62\x3a\x5e\x4e\x61\x39\xfe\xf7\xd5\x14\x66\xef\x60\x7e\xbd\x84\xe9\x1f\xb3\xdb\xe5\x2d\x2e\x02\xfd\x0b\x00\x11\xc3\xdd\xdd\xec\x2d\xdc\x2c\x66\x1f\xc6\x8b\xff\xc2\x7f\xa6\xff\x85\xb7\xd3\x77\xe3\xbb\xab\x25\x6c\xb9\x5c\x69\x26\x63\x95\xae\xf2\x5c\xc4\xfd\xc1\xf0\x02\x40\xb2\x94\xc3\x72\xfa\xc7\xf2\x62\xf0\x13\x9d\x73\x1c\x45\x2a\x97\xd6\x00\xd3\x9c\xe4\x95\x69\x91\x32\x7d\x84\x5c\x0a\x8b\x02\xda\x6a\x26\xf3\x84\x69\x61\x8f\xb0\x51\x9a\xae\x52\xad\xff\xe4\x91\x35\xee\xc0\x9e\x03\x08\x83\xfc\x22\xcd\x99\xe5\xb1\x3b\x61\x4d\x8a\x4d\xea\xda\x6a\x8e\x23\xa8\x83\xe4\x1a\x97\x21\x56\x5e\x3a\x15\xa9\x08\x6e\x3a\xe5\xc2\x3c\xf7\xaa\x6c\x70\x78\x7e\x77\x75\xd5\x2d\x18\xa5\xb7\xab\xe6\x8c\x9a\xc0\xf0\x8f\xd9\xfc\xed\xf4\x8f\xbe\x23\xa5\x59\x15\xb9\xfb\xcf\x43\xf0\x43\x77\xf3\xd9\x6f\x77\xd3\xbe\xff\x6b\x72\x3d\xbf\x5d\x2e\xc6\xb3\xf9\xd2\xaf\xb4\xda\x3c\xc0\xbb\xeb\xc5\x74\xf6\x7e\x5e\x99\x3d\x80\xc5\xf4\xdd\x74\x31\x9d\x4f\xa6\x74\xc9\x38\xfd\x62\x00\xb3\xf9\x72\xba\xb8\x9a\x8e\xef\xa7\x30\x9b\xc3\xcd\x78\x31\x75\x7c\xc2\x2c\x77\x95\x78\x33\xa8\xa1\x10\xa9\x24\xe1\x11\x29\xa0\xda\x80\xc9\xd7\x92\xe3\x05\xcb\x42\xc5\xf1\xe6\x96\x3b\x0e\x29\xcb\x32\x21\xb7\x48\x75\x7f\x33\x31\x60\x55\xa0\x46\x7e\xc2\xc0\x8e\xc9\x38\x71\xb7\x79\xff\xc7\xd5\x78\x0e\xb3\xb7\x8e\xd1\xfd\xd5\x78\x6e\x86\xa0\xb9\xc9\x70\xa5\x3d\x4f\xba\x15\x76\x9f\x45\xcf\x53\x58\x7f\xa3\x4f\xdf\x4d\x45\xc2\xe5\x9c\xa6\x94\x1b\x12\xf6\x84\xe1\x92\xdc\xf5\x96\xb3\x07\xc1\x48\x16\x7c\xeb\xad\x79\xcb\xd5\x56\xb3\x6c\x27\x22\xd8\xb0\x3c\xb1\x10\xab\x94\x09\xd4\xed\xdf\x72\x65\x51\x96\xef\x95\xda\x26\x1c\x26\x89\xca\x63\xf4\x36\x76\xa3\x74\x1a\xbc\xdb\xce\xda\xcc\xbc\x79\xfd\x3a\xc2\xd1\xd1\x96\x48\x47\x91\x4a\x5f\xc7\x2a\x32\xaf\x03\xf3\xe3\x25\x93\xf1\xa5\xa6\x45\x8d\x9f\xe9\xb6\xe0\xcc\x54\xc8\x98\x67\x5c\xc6\xe8\x41\x2b\x1b\x62\x9a\x33\xb4\x2a\x66\xc9\x27\x09\x43\xf6\xfb\x49\x49\xbc\xef\x2b\x15\x91\x15\xd2\xcd\x1e\x84\xdd\x09\x09\x7e\x05\xb0\x5c\xc6\x78\xf7\x3b\xb6\xe7\xa0\x55\x2e\xe3\x4b\xab\x45\x06\x92\xdb\x83\xd2\x0f\x90\x30\xcb\x65\x24\x38\x39\xcc\x5c\xc6\x5c\xc3\x0f\xa9\x01\x25\xd1\x84\x91\xdf\x8f\x3f\xd8\x1d\x64\x5c\x47\x68\xa4\x09\xef\xd4\x04\xb7\x68\x50\x06\xbc\xbd\xda\x9d\x56\x2d\xaa\x72\x03\xef\x58\x24\x12\x74\x41\x9a\x67\x9a\x1b\x4e\x1e\xcb\x39\x4d\x5c\x14\xdd\xbd\xf7\xa1\xa9\x8a\xf3\x84\x8f\x00\x26\x3e\x10\x18\x88\x15\x7a\x61\x10\x48\xc7\x22\x72\xc7\xb1\xd0\x3c\xb2\xc9\x91\x44\x01\x1b\xc7\x5d\x70\x33\x04\x25\x93\x23\x8c\x3f\xa2\x89\x5c\x67\x5c\x33\xab\xb4\x09\x76\x03\xfc\xaf\x5c\x64\x29\xca\x1d\x8d\xc0\xcb\xa7\x2a\x52\x16\x78\x1d\xff\x7e\x95\x18\xd3\xdd\x3a\x93\x77\x41\xcc\xed\x4c\x73\x46\x9e\xba\xbe\x1e\xda\xaa\xca\x75\xc4\x4d\xb1\x57\x17\x97\xb6\xe4\x9b\x3f\xa2\x9a\x80\xd9\xa9\x3c\x89\x61\xed\xa2\x9a\x88\xb9\xe6\x31\x30\x30\x42\xe2\x01\x36\x4c\x24\xb9\xe6\x5e\xe7\x5b\xd9\xcc\x30\x00\xa2\x82\x58\xe5\xf7\xe4\xec\xe4\xd2\xaa\x84\x6b\x86\xdb\xcb\xb2\x44\x78\x5d\x74\xd2\xdf\x89\xed\x0e\x79\xb0\x3d\x13\x09\x5b\x93\xf4\x86\x70\x54\x79\xd8\x8e\x67\x74\x54\xb9\xae\x4f\x67\x91\x56\xc6\x40\x9a\x27\x56\x64\x09\x77\xba\x0e\xf5\x3d\x91\x62\xf3\x24\x83\x4c\x2b\x8b\x31\x86\x6d\x99\x90\xc6\x42\x2e\xf9\x23\x3a\x2f\x1e\x87\x83\x75\xc7\x96\x70\xb1\xcf\x08\x2e\x85\xa3\xaa\xa9\xba\xdb\xdf\xea\x29\x2b\x28\xc8\x86\xc4\xa7\x1a\x5c\x9a\x7f\x9f\x06\x9b\x62\x72\xd3\x13\x16\x03\x35\x87\xe8\xbe\x76\x45\x1d\x4f\x51\x4e\x6f\xd8\xe7\xdc\xfb\x8c\xa5\xce\x0d\x45\xf9\xdc\x70\x72\x2e\x29\xd3\x0f\x14\xf2\xad\x66\xd2\x64\x4a\x5b\x37\xb4\xe6\xf6\xc0\xb9\x04\x7b\x50\xce\x30\x03\x36\x2c\x8d\x72\x04\x70\x9b\x67\x38\x05\x59\x15\xd3\x53\x6e\x77\x2a\xc6\xfb\x8e\x92\x3c\xe6\x25\x82\xec\x65\x09\x13\xb2\xf7\x06\x75\x5f\xc8\x4b\xcb\x1f\xed\x10\xaf\x5b\x46\xfa\x98\xd5\x78\x78\xfa\xd9\x8d\xe1\x51\xef\x0d\xd0\xbf\x50\xa3\xdb\x6c\x44\x54\xec\x51\xc8\x58\xec\x45\x9c\xb3\x04\x26\x73\x52\x34\xce\xa2\xc2\x7d\x1c\x3f\x47\x7b\x56\xde\x69\xac\x68\x0f\xc2\x41\x15\xa3\xa3\x55\x41\xd0\x16\xe9\x62\x63\xbb\x09\x4a\xa9\xd4\x74\x09\x26\xbf\x4e\x27\xff\x81\x7e\x39\x3c\x9b\xf7\xbd\x78\x86\xfe\xd8\x83\x13\x24\xd3\xd8\xce\xb0\xb9\x7c\x55\xe5\x1a\x43\xc3\xe6\x59\x9a\x0a\xd9\x18\x6e\xaa\x65\x73\x76\x55\x39\xc3\xf7\x16\x35\x6f\xec\xa2\xc9\xb5\xb9\xff\x73\x5c\x0b\x50\x5c\xf1\x44\xe4\x19\xa1\xff\x6a\xfc\xd1\x25\x00\xa8\xc2\x45\x8e\x11\x73\x23\xb6\x92\x05\xa8\xc5\x80\xbc\x38\xba\x9a\x3d\x3a\x4f\x52\x68\x4f\xc2\x63\x78\x15\xd6\x7a\x85\xd1\xaa\xee\x4d\xbb\x61\xed\xa7\x67\x38\x9d\x0e\xf7\x42\xfe\xe8\x76\xb9\x98\xcd\xdf\xf7\xbf\x1d\x9c\x68\x0b\x0d\xa3\xa2\xb0\xde\xb0\xb7\xee\x0d\x7b\x51\x6f\xd8\x8b\x7b\xc3\x1e\xef\x0d\x7b\x9b\xde\xb0\xb7\xed\x0d\x7b\xbb\xde\xb0\x27\x5a\x94\xa7\xc3\x5d\xbd\xa0\x7b\xaa\x70\x6d\x2e\xf8\x85\x9e\x6b\x2f\x05\xf4\x0b\x70\xeb\x72\x3c\x48\x99\x41\x4c\x61\x29\x59\xc3\x8b\x4d\x12\xb8\x9f\xcf\xc8\xe4\x73\xc3\xa1\x19\xe9\x71\x8c\x42\x98\xe6\x90\xa9\x2c\x4f\xe8\xbe\x95\x84\x98\xa7\x88\x14\x30\x1c\x33\xd8\x8a\x3d\x97\xd5\x69\xb3\x98\xb3\x24\x39\x82\xb0\x70\x08\x81\x37\x53\xc6\x88\x75\x42\xb8\xca\x2a\x9f\x4f\x41\xa6\x79\x8c\x61\xcf\x81\xbf\x47\x5e\xc9\x97\xbd\xde\x10\x78\xc5\xd3\xac\x7c\x1c\x4d\xf8\x4a\xc4\x8f\x70\x3d\xc7\x8f\x06\x08\xfb\xfe\xfe\xeb\x74\x31\x05\x21\x57\x78\x88\x5f\x60\xb9\xb8\x9b\xfe\x74\x86\x0b\x3a\xe7\x27\x19\xbc\x1b\x5f\xdd\x12\x07\x8a\x01\x4a\x83\x54\x87\x21\x49\x08\x81\x98\xe6\x60\x44\x8a\x81\x39\x20\x33\x9e\xb8\x34\x24\x8f\x1e\x20\xcf\xc8\x96\x30\x08\x4b\x43\x52\xce\xb8\x46\xa4\x82\x90\x89\x02\xb8\xe1\x7f\xe5\x88\x28\x59\x02\x26\x62\x14\xc9\x13\xa5\x1e\x9c\x34\x25\x14\xe7\x44\xf1\x37\xae\xa4\x33\x17\xc1\x1b\xbf\x00\xe8\x74\xa8\x48\x84\x1a\xda\xb4\x0d\xfc\xfe\x2f\xf8\x06\xc6\xf3\xb7\x44\xf3\x33\x7c\x07\xff\x0b\xdf\x7d\x3f\x18\x5e\x5c\x00\x5c\x5e\xc2\xef\x3b\x2e\x31\xdd\x41\xae\xb3\xdb\x92\x27\x1d\xd6\x7f\x2f\xdc\x88\xa6\x48\xe7\xb2\x5e\xf2\x1f\x4e\x45\xdc\x79\x1c\x3f\x26\x61\xfc\x71\x04\x27\x7c\x3d\x4f\x61\x1c\xb5\xa9\x88\x03\xe5\xa3\x79\x6e\xf8\x08\x0f\xe2\xa6\xe0\x01\xc3\x16\xf9\x63\x26\x34\x8f\x57\xac\x0c\xd0\x54\x9b\xb1\x2a\x5d\x1b\x8b\xee\xce\xe9\x2b\x32\x66\x71\xac\xb9\x31\x0e\x56\x15\xd0\x2e\xd3\x7c\xcf\xa5\xf5\x3b\x74\xa6\x41\x2b\xc2\x9a\x6f\x94\xaf\x24\x18\x8b\xfa\xea\xcb\x30\xe6\x68\x2c\x4f\x11\x56\xee\xb9\xde\x22\xd4\x82\xea\x3e\x96\xb3\x0f\xd3\xdb\xe5\xf8\xc3\x0d\xfc\x3e\x5b\xfe\x4a\x7f\xc2\xc7\xeb\xf9\x74\x58\x90\xad\xd8\x06\x4d\x92\x4c\xfb\x7e\x7c\x75\xea\x04\x7b\x3f\x7e\x03\x31\x3b\x9a\x1e\x1d\xb3\xea\x92\x6a\xf1\x69\x2f\x45\xd5\x7b\xec\xa5\x18\x42\x33\x5e\xb9\x44\xd0\x49\xae\xe9\xb2\x3a\xc2\xcc\x33\x02\x97\x5b\xa3\xc9\xc8\xaf\x5c\xe5\xb1\xcf\xa2\x5a\x84\xf2\x89\xba\x73\x57\x21\x51\x8b\x94\xb4\x4c\x48\x1e\x97\xe9\xc8\xfd\xcd\xa4\x96\x0b\x85\x24\x88\x28\xc8\x65\x05\x4e\xd5\x29\x43\x4a\x98\x98\x4f\xfa\x95\xa6\x7c\x7f\xe4\xed\x7c\x79\xfd\xf6\xba\x6f\x38\x93\xd1\xff\x0c\xde\xf8\x1b\xb5\xc5\x1e\x84\xb4\x0a\x18\x4c\x66\x6f\x17\xce\xec\xed\x31\xe3\xc8\x62\x36\x5f\x5e\xae\x59\xf4\xc0\xe3\xf2\x3b\x39\x39\x4c\x58\x59\x92\xa8\x83\x21\xbd\xc3\x1d\x6e\xb9\x76\x1c\x52\x66\x77\x0d\xcd\xa3\xd0\x0a\xfc\x31\x4a\x72\x83\x71\xd7\xd5\xdf\x98\x90\xae\x6a\x81\xfe\x45\x33\xb9\xe5\xa0\x28\x7d\xc3\x44\x61\x04\xb0\xbc\xfa\xe9\xed\xe2\x0d\xdc\x28\x63\xb7\x9a\xdf\xfe\x76\xd5\x33\xcd\x3d\x16\x2e\x58\x62\xaa\x67\x09\x8e\xba\x94\x58\x6c\x77\x16\xd8\x5a\xe5\x16\x1d\x5c\xa7\x73\xf1\xc2\xfc\xf2\x70\x5d\xb1\xd3\x9a\x23\xf2\x06\xb8\xa2\x2d\xb6\x82\xbb\x1a\x05\x86\xed\xd9\xcd\xfe\x7b\x82\x77\xfb\x7f\xf8\x00\x1d\x2e\xe7\x04\x09\x64\x9a\x6f\xc4\xe3\x2a\xe1\xb2\xd5\xdb\x55\x86\x83\xd3\xeb\xd7\xd7\xfb\x05\xdc\x72\x34\x56\x21\xff\xf9\x17\xf8\xff\xdf\x0d\xe0\x7a\x01\x6d\xf4\xff\x68\xa3\xff\xf6\xbb\x7f\x0e\x06\x5f\x6b\x20\x4d\x38\xe2\x28\xcf\x54\xe5\x82\x0f\xf0\x34\x5e\x4a\x67\x11\x04\x2e\xe1\x57\xae\x15\x4f\xe1\xc3\xd8\x57\xde\x72\x29\xfe\xca\x6b\x75\x35\x1c\x2a\x7d\x68\xb0\xb1\xa2\x0e\x3a\x02\xa4\x20\xbd\x2d\x88\x10\x44\x78\x46\x68\x4a\x05\x6d\x61\x94\x68\x9f\x6b\x1e\x31\x17\x67\xef\xe7\xb3\x09\xa4\xec\x88\x9a\x9b\xaa\x3d\x8f\x91\x5b\xc8\x58\xa8\xac\x57\x58\xb6\x8b\x2e\xe5\xd2\x14\x53\x98\x84\xd9\x0d\xcd\xc4\x18\x3f\xe2\xa3\xc0\x53\x98\x92\x1f\xd9\xb4\xe4\x07\x88\x39\xd5\xc8\xc9\xa7\xab\x24\xae\x9e\x0f\x1d\xfe\x31\x94\x85\x68\x70\xa7\x5c\x0e\xe8\x7d\x04\x66\xda\x5a\xa5\xce\x4a\x9d\x27\xda\xb0\xc8\x9d\x12\x4d\xff\xc0\x74\x8c\x62\x2b\x47\x84\x0c\x35\x7c\x58\x6b\x11\x6f\xb9\x2f\x5c\x96\xc4\x54\xa8\x2c\xe8\xfd\x77\x43\x80\x2d\x43\x67\x63\x0d\x22\x70\x4b\x4e\xd1\xaa\xd6\x6d\xfb\xcf\x78\xba\xfa\x5d\x8c\x00\x6e\x3e\x93\x07\x4b\x8c\x02\xcd\x23\x2e\xf6\x78\x27\x5b\x74\x3c\xb9\x15\x2a\x27\x9f\x34\x5e\xdc\x40\xff\xd5\xfb\xf1\xe2\xe6\xd5\x00\xd0\x07\x92\x1c\x88\xd1\x99\x63\x07\x1c\x44\xa1\x93\x53\xc9\xcd\x05\x52\x64\xe7\x77\x5a\xd9\xe6\xbb\x16\x2e\x06\x0c\x4f\x36\x97\x2e\x7a\x82\x8b\x9e\xdf\xef\xdc\xab\x87\x90\x2c\xb2\x62\x8f\x39\xab\x77\x99\xe4\xe8\x18\xec\x98\x8e\x5d\xc0\x2d\x92\x9a\x7f\xee\x74\xa8\x17\xd3\xf4\x3d\x4b\x72\x0e\x07\x66\x20\xda\x29\xc3\x25\xac\x8f\xee\x82\x25\xce\x20\x7e\xbe\x0a\xe9\x8b\xd1\x69\xaa\x24\x89\x20\x62\xd1\x8e\xc3\x72\x79\x35\xa4\x25\xf1\x6c\xc8\x1c\x19\x6c\x04\x06\x01\x03\xc2\xf6\x0c\x7c\xf7\x18\x4e\x5a\xf2\x65\x09\xae\x8e\x93\x47\x00\xd7\xd2\x15\xa8\x08\x1f\x46\x1e\x63\x08\x83\x0b\x1e\x08\x51\x26\x22\x12\x36\x39\xfa\xca\x1f\x61\x49\x57\x40\x20\x50\x22\x22\xda\x0e\x97\x56\xfb\xea\x42\xf9\xf1\x20\x92\x04\xc8\x57\xe6\xd2\x8a\x24\xd8\x02\xc6\x2f\x6f\x24\x43\x3a\xa2\x85\xc3\x4e\x44\x3b\xc8\x94\x90\x05\x57\xa6\x33\xc7\x60\xad\x39\x7b\x80\xef\x2f\xe9\x78\x5e\x89\x37\xb9\xcd\x75\x77\x6d\x34\x94\x9e\x53\x16\x3d\x23\x82\x74\x95\xca\x91\xe3\x89\xf3\x3f\x17\x72\x5c\x0c\x6b\x19\xfa\x12\xf8\x58\x77\x7c\xd5\x10\x4e\xb1\x9e\x1e\x79\x88\x5d\x55\x93\xff\x4e\x50\x58\x8f\x0b\xa5\xb0\x86\x28\x9c\x6a\x38\x48\x59\x34\xac\x08\xb3\x1e\x37\x86\xa5\x74\x9a\xa1\xea\xfc\xb3\x43\x85\x59\xc7\xf3\x43\x23\xea\x15\xeb\xb4\x07\xbf\xea\x46\xaa\x4c\xdd\xd7\x6a\xdc\x0b\x98\xf1\xd6\x41\x93\xd9\x4d\xc8\x3f\x0e\x3b\x61\x79\xe2\xdf\x09\x66\x37\xa6\x0d\x3f\xba\x90\xe1\x98\x7e\x06\xf0\x59\x89\xec\x05\xb1\xcf\x59\x45\x04\x10\x59\x77\xd5\xf4\x44\x48\x43\x10\x59\xf5\x2a\xcf\x0e\x14\x92\xcb\xba\x6b\x17\x2f\x7c\x47\xed\x90\xa3\x41\x59\xb2\xf3\x0f\x53\x2a\xb7\xae\x0c\x15\x69\xb1\x46\xec\x00\xda\x7d\x12\xd2\x58\x26\x23\x0c\x98\xe3\xe6\x37\xc0\x64\xda\xd5\xf8\xa5\x2b\x80\x3b\xb0\x8b\xa6\xab\xb4\x4f\xdc\xfd\xab\x61\xdf\x7b\x2f\xc3\x52\x1a\x2c\xeb\xb3\x08\x2d\x06\xb0\x17\x45\x02\x51\x46\x9e\xee\x07\x20\xb7\x99\x97\xd2\x92\xd3\x37\xa3\xb6\xbb\x7c\xb1\xc4\xeb\xda\x3d\x54\x2f\x11\xcb\x52\x7c\x0b\xf6\x83\x88\x23\xe4\xba\x45\xb1\xda\x3f\x6b\x23\xf2\x2d\xe2\x80\x89\x76\x3c\x65\x78\x2f\x26\x04\xf3\x0d\x67\x18\x1b\x1c\xf4\x63\x71\xcc\x09\x8c\x1e\xea\xd3\x25\x77\xf0\x63\xed\x49\xba\x1b\x04\xd6\x7f\x3a\xbc\x7d\xee\xd1\xf5\x6b\x1e\x2e\xaa\xaf\x0f\x28\x97\xd9\xfc\x76\xba\x58\xa2\x02\x5f\x57\x16\xa6\x61\xb8\x1f\x5f\xdd\x4d\x6f\xa1\xdf\xdb\xa7\xbd\xcf\x25\x75\x1a\xd2\x0b\xcf\xda\x1e\x8b\x32\xd8\x0b\x6d\x73\x96\xc0\x7c\x36\x21\xb5\xa6\x11\xaf\xcb\xcc\x18\xb1\xf5\x00\x2d\x18\xc1\x10\xee\x3f\x0c\x51\x69\xfd\xf3\x5d\xed\x9d\xce\x25\xc6\x88\x6a\xaa\x38\xae\x28\x7c\xb8\x25\x2b\x4c\xf1\xea\x4a\x84\x56\x82\x31\x7f\x47\x0e\xfb\x05\xc8\x76\x5f\xdb\x22\x82\x73\x02\x88\x6b\x0e\xb1\x30\xa5\xd9\x11\x0a\x64\xd2\xf3\x08\xf9\xf5\xfc\x7a\x39\x7d\x43\x11\xfb\xc3\xdd\xed\x12\x27\x55\xb2\x81\x06\x86\x7f\xa2\xce\x75\xfe\xd1\xfd\x6b\x21\x85\x8b\xdd\x78\x85\x65\x3d\x8b\x48\xbc\x0d\x78\xb1\xf8\x12\x57\x72\x84\xdc\x50\x5d\xcf\x49\x04\xe1\x5f\x71\xfd\x6e\xb6\x63\x88\x1f\x46\xf4\x5f\xf7\x37\xca\xa2\xac\x35\x8c\x2d\xf8\xd7\x5f\xc4\x79\xfe\x59\xd7\x21\x57\x79\x74\x85\x03\xea\x53\x89\x76\x3c\x7a\x30\x4e\xbc\x24\x05\xab\x3c\xfb\x1d\x87\x3d\xd3\x88\xcb\x6b\xc6\x35\x02\x98\x55\x51\x1a\xb2\x8f\xa8\x1c\x40\x39\x56\x1c\x0b\x8f\x41\x65\x9e\x24\x8e\x55\xa4\x92\x3c\x95\x94\x3d\xc4\x8a\xf0\xee\x10\x76\xea\xc0\xf7\x5c\xe3\x6c\x57\x4c\x08\x36\x9b\x72\x66\x72\x5f\x16\xab\x14\x3c\x4d\x1e\xed\x80\x39\x7e\xce\x2d\xc0\x46\x68\x63\x7d\x1e\xb6\x4f\x8b\xca\x9d\xd7\xe8\xf2\x6f\x4a\xd0\xfb\x8e\xa2\x52\x66\xc4\xb4\xda\xf1\x2b\x27\xf8\x6a\xe1\x68\x34\x1a\xa0\x68\xfd\x9d\xb9\x82\x20\x94\xd7\x50\x54\x08\x2b\x58\xc7\xd1\xba\xda\xe6\xdf\x83\x74\xc2\xf2\x27\xfd\x30\xfe\x7b\xbd\x23\xc6\x7f\xac\x3a\x64\xd2\x27\xc4\xe6\x2c\x33\xb5\x08\x86\x80\x46\xb9\x62\x14\x99\xe3\x53\xa6\x12\xa0\x8b\xfb\x7f\x2b\xe4\x38\xfb\x59\xc6\xfc\xb1\x5e\x40\x09\x66\xf5\xcd\x09\x30\x71\xec\x87\x8e\x5d\xa3\x14\x59\x8e\x20\xc7\x93\x98\xe5\x08\x4e\x82\x96\xfb\x5c\x8f\x5a\x52\x44\x2d\xd2\x2e\x40\x5a\x5b\x6f\x4c\xd6\x0e\x53\x56\x22\xeb\x7a\x10\xa6\x95\xc2\x16\x3c\xd0\xe4\x51\x4e\x76\xf8\x5e\xab\x3c\x33\x65\xcf\x06\x30\x30\x9c\xa2\x65\x28\x48\x6d\x44\x62\x39\xe5\x97\xcd\x49\x4c\x93\x7b\xa5\xf7\x7e\x67\x44\x78\x8b\xd4\x73\xd6\xa0\x45\x07\xc2\x31\x2b\xa5\x47\x17\x41\x15\x41\xcb\x1f\x6d\x48\x32\x25\xb5\x1e\x0a\x57\x1e\xd4\x79\xc2\x0d\x64\xb9\xd9\x61\x78\xce\xad\x7b\x9a\xb6\xd1\x0e\xbd\x53\xb1\x82\x75\x96\x1a\x6f\x79\xf5\xcd\xd7\xd5\x3e\x0c\x86\x1c\xf7\xb2\x17\x89\x18\x51\x34\xb3\x9e\x2d\x31\xaa\x56\xf1\xbb\xb5\xce\xf8\x63\xac\xb6\x74\x8c\xaf\xe8\x2b\xa0\x37\x62\x02\x81\x19\xa5\xed\xe1\x5b\x57\x4a\xf8\x82\xf6\x7c\xa6\x69\xa7\x71\x4d\x0b\x94\x51\xfb\x4d\x87\x72\xee\x9e\x27\x47\xc8\xb8\x4e\x85\x31\x54\x49\x91\x54\x7d\xa2\x97\x84\x4d\x9e\x7c\x81\x34\x57\x78\x23\xcf\x10\x69\x83\x4b\x6b\xe0\x73\xcd\x43\x41\xce\x2d\x49\xa7\x90\x3d\xef\xa0\x4b\xd2\xd9\xbc\xdf\xa3\x57\x77\x95\x5b\xaa\xca\x3a\x37\x7d\xa3\x95\x55\x91\x4a\x60\x9e\xa7\x6b\x84\xc5\xb7\x9c\xc3\x6b\x6e\xa3\xd7\x99\x1f\x31\x03\x4f\xa9\xb4\x75\x95\x6d\xf7\xf7\x6c\xf2\xe1\xc6\x15\xae\x51\xcf\x23\x15\xfb\xef\xb7\xd4\xe5\x43\x58\x9d\x8a\x56\xae\x90\x53\xa6\xe2\x7d\x2a\x7b\x4b\x65\x69\xc0\x33\xaf\xdf\x89\xfb\x76\x7f\x33\xf1\x83\xe4\x07\xdc\xff\x43\xa3\x07\xd5\x8e\xfd\xd6\x51\x85\xdc\x71\xcb\x4f\x2e\xe6\xc0\xf5\xa2\x24\xfb\xd7\x2f\xf0\x8d\x3b\xb7\xd1\xd1\x0a\x11\xf2\xca\x58\x46\x3d\x09\xcb\x21\x50\x8f\x5b\x28\x26\xa0\xf5\x95\xe7\x3b\xa0\x0d\x17\x6c\x84\x81\x6f\xa1\x8f\xa3\x83\x2a\x2b\x2e\x63\xc7\x08\xce\xb0\x42\x09\x75\xb2\x8a\x8d\x6d\xee\xaa\xfa\x35\x2c\x10\x0e\x10\x89\x58\x97\xc6\x67\x6c\xf5\x83\x93\xd5\xaf\xca\x58\x03\x87\xd7\xce\xaf\xb8\x17\xca\x86\x4d\xb8\xc7\x0f\x44\xfd\x49\xa2\x0e\xa1\x53\x06\xbd\x08\x15\xb4\x50\x89\x47\x7e\xbd\x76\xcd\x0c\x8b\x9f\x1b\x2d\xae\xd2\x94\x6e\xae\x60\xed\x71\x86\x83\x0a\xbe\x68\x67\x32\x26\x7d\x6e\x87\x7a\x95\x70\xf6\x50\xf4\xc1\xe0\x44\xc7\xf0\xc0\x8e\x20\x08\xe9\x1d\x1d\x00\x73\x15\xd4\x88\x1b\xc3\xf4\xd1\x71\x34\x61\xe7\xb5\x27\x48\xb7\xdd\x96\x57\xc9\x5b\x9f\x71\x9e\x6e\x33\x08\xa0\x56\x0d\x28\x0e\xde\xf8\xea\x98\x8d\x3f\x76\x30\x62\x9f\x9a\x4c\xaa\x5f\x9a\xcd\x30\x4d\xc1\xb6\x65\x93\x27\x44\x27\x7d\x2d\x4d\x82\x7a\xa0\xad\x8d\xb6\xc5\xed\xb6\xfb\x6f\xeb\x9d\xf9\xca\x75\x5a\x35\xa9\xad\x9b\xe6\x05\xce\xd3\x9e\x81\x97\x23\xe7\x5e\x77\x1a\xdb\x6d\x67\x53\x8e\x7c\x06\x9b\x9a\x72\xb5\x4a\xb5\xab\x8a\x73\x46\x8a\xe7\xf8\xd5\x06\x3f\x8f\x5f\xa1\xb3\x6d\x7b\xa3\x81\x7a\x70\xfe\x74\x66\x4f\xad\x3c\x8a\x81\x16\x1e\x67\xab\x52\xf5\xfb\x3d\x55\x86\xd6\xf8\xef\x73\xe1\xcc\xf9\xe4\xb2\x5d\x21\x4c\x87\xad\xc7\x03\xb5\x54\x9e\xb0\x54\x00\xfa\x5f\x02\x00\x8a\xf4\xf7\x3c\xa0\xff\x22\x64\xd0\x0e\xe2\x4f\xca\xc3\x01\x0a\x7f\x01\x5e\xf6\x55\xbc\x50\x9f\x2d\x8a\x0c\x24\x2b\x17\x3b\x2a\x55\x8d\xa2\x5a\xf0\x64\x8d\xad\xd0\xb3\x82\xe3\x33\xba\xc8\x6a\x79\xe7\xe7\x56\x67\x53\xd6\x2e\xf1\xb3\x57\x51\xeb\x20\x0b\x4b\x36\x6b\xed\x65\x7d\xfe\xfc\x60\x4b\x8f\x59\x60\x77\xd2\x63\x16\x06\xea\x3d\x66\xf4\xb5\x85\x8f\x5b\xb6\xc9\xe4\x74\x33\x2d\x38\x79\x95\xb2\xa8\xfb\xed\xe0\x19\xa9\xdd\xd9\x3e\x37\x77\x80\xf2\x74\x35\x25\xf3\xb5\x5a\x1a\xf5\xcd\x65\xee\x47\x4d\xc6\xa8\x48\x38\xac\x58\x69\xc9\x2d\xea\xc0\xa1\x66\x16\x20\xb9\x6f\xd8\x08\x6f\xac\x56\x81\x39\xa6\x29\xb7\x5a\x44\xd4\xd2\xb6\x49\xd4\xa1\x8d\xd1\x17\xa8\xad\xdb\xe3\x8b\xaa\xec\x19\x35\x6b\x55\xa1\xaa\xe3\x97\x76\x73\xb6\x2f\xf7\x29\x9a\xf0\xde\x70\xc2\x6e\xd8\x32\xbd\xd9\x5a\xdb\x9c\x71\xca\xa5\x2b\x98\x39\x8a\xae\x88\xe6\x79\x9c\x1a\xc0\x89\xef\x78\x22\xce\xb5\x2f\xd5\x72\xc0\xcf\x5c\xea\x0b\xb5\xdb\xfb\xce\xfb\xf9\x8c\x5a\x95\x9c\xeb\x2c\x5e\x26\xac\xc2\xf4\xdf\xf5\xb3\xf9\x1f\x2f\xb5\xb7\x64\x7e\xc6\x03\xd7\x5e\x8a\xd5\x3e\x61\xf2\xe9\x16\xc2\x0e\xff\xd8\xec\x2e\xa4\x6f\x09\xa3\x6e\xdd\xd6\xae\x43\x3f\x86\x29\x94\x6b\x3d\xf4\x1f\x7e\xfe\x05\xbe\xff\xe6\xc7\x1f\x4e\x02\x54\xb3\xe5\x6d\x18\x66\x9c\x79\x02\xeb\xe8\xd8\xc6\x03\xb7\x7a\xa6\xfa\xac\x86\x8b\x6a\x6f\xac\x6b\x29\x45\x9d\xaa\xe7\x13\xe0\xa8\x23\xb0\x9e\x76\xfa\x39\xed\x98\xa8\x34\x43\x5f\x32\xc7\x14\xb0\xff\x6a\x32\x2f\x5b\xb9\xb3\xdd\xd1\xa0\xcb\xa2\x37\x02\xae\xe9\x17\x6d\x4a\x52\x07\x2e\xa9\xc8\x4e\x19\xfa\x0d\xcf\xfd\x87\x6e\xd7\x15\xc9\xe7\xfd\xbe\xad\x53\x83\x5e\xae\xed\x30\x94\x64\x26\x73\xb8\x93\x31\xd7\x09\x3b\x16\x45\x53\x14\x43\x5e\x7e\xec\x63\xe6\x57\x83\x60\x93\xf9\x13\x47\x5f\x85\xe9\xa1\x86\x1a\xc9\xe2\x3c\xb5\x03\x55\xe9\x3a\x9f\x6f\x23\xd7\xd0\x5d\xa1\xaf\x2a\x6e\xf3\x73\x45\x4e\x51\x5b\x37\x79\x74\xd2\x49\x1e\x75\xfe\xc8\x25\x92\x7e\x8a\x2f\x31\x7f\xf0\xbf\xfc\x0d\x0f\xa8\x9d\xa5\xe4\xf4\x19\x21\xab\x22\xaf\x5a\xef\x5f\x47\xf9\xce\x72\x9d\xfa\xe2\xce\xca\xff\xb4\x09\x63\xf7\xbf\xaf\xaf\x5b\xda\x20\xa8\x4f\x9b\x9c\x4c\xda\xd1\x48\x18\x06\x67\xf3\x7e\x6f\xbd\x3b\xee\x79\x6f\xd8\x7b\xd8\xa7\xbd\x61\xef\x4f\x26\x92\xde\xb0\xf7\x49\x49\xde\xd2\xf2\xef\x2f\xeb\x29\x20\xf6\x82\xaf\x06\x2f\x7d\xcb\xff\x17\x00\x00\xff\xff\xa0\xa9\xbd\x7e\xa3\x3e\x00\x00")

func _1517299952_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1517299952_initUpSql,
		"1517299952_init.up.sql",
	)
}

func _1517299952_initUpSql() (*asset, error) {
	bytes, err := _1517299952_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1517299952_init.up.sql", size: 16035, mode: os.FileMode(420), modTime: time.Unix(1519862399, 0)}
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
	"1517299952_init.down.sql": _1517299952_initDownSql,
	"1517299952_init.up.sql": _1517299952_initUpSql,
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
	"1517299952_init.down.sql": &bintree{_1517299952_initDownSql, map[string]*bintree{}},
	"1517299952_init.up.sql": &bintree{_1517299952_initUpSql, map[string]*bintree{}},
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

