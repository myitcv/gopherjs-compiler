// Code generated by vfsgen; DO NOT EDIT.

// +build !gopherjsdev

package gopherjspkg

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// FS is a virtual filesystem that contains core GopherJS packages.
var FS = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Time{},
		},
		"/js": &vfsgen۰DirInfo{
			name:    "js",
			modTime: time.Time{},
		},
		"/js/js.go": &vfsgen۰CompressedFileInfo{
			name:             "js.go",
			modTime:          time.Time{},
			uncompressedSize: 10926,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\x5f\x73\xdb\x36\x90\x7f\x16\x3f\xc5\x96\xd3\x19\x8b\x89\x22\xf5\x4f\xc6\xd3\x71\xce\x0f\x69\x73\xcd\xa5\xd7\xb8\x99\xba\xb9\x3e\x64\x32\x1e\x88\x5c\x4a\x88\x29\x80\x05\x20\x29\xaa\xed\xef\x7e\xb3\x58\x80\x22\x45\x2a\x76\x9a\xde\xcc\xf9\xc5\x12\xb1\xf8\xed\x0f\xbb\x8b\xdd\x05\xa8\xd9\x0c\xde\x88\xfc\x5a\x2c\x10\x3e\x58\xa8\x8d\xde\xc8\x02\x2d\x94\x6b\x95\x3b\xa9\x95\x85\x52\x1b\x90\xca\xa1\x11\xb9\x93\x6a\x01\x5b\xe9\x96\xa0\x84\x93\x1b\x84\x5f\xc4\x46\x5c\xe6\x46\xd6\x0e\x9e\xbf\x79\x65\xa7\xf0\x93\xa8\x2a\x0b\x4e\x83\x5b\xa2\xc5\x16\x8a\x30\x08\xce\xa0\x70\x58\x80\xad\x31\x97\xa2\xaa\x76\x30\xdf\xc1\x4b\x5d\x2f\xd1\xfc\x72\x09\x42\x15\xe0\x8c\x50\xb6\xf2\x42\x85\x34\x98\xbb\x6a\x17\xc0\xa4\x81\x5c\x1b\x83\xb6\xd6\xaa\x20\x1a\x2d\xd5\x76\xa7\x9c\xf8\x38\x4d\x66\xb3\x64\x36\x83\xb7\x16\xe1\xb5\xb8\xc6\x3f\x8d\xa8\x6b\x34\x34\x1f\x3f\xd6\xda\x22\xac\xd0\x2d\x75\xe1\xe9\xed\x67\x4f\x9b\x09\x3f\xaf\xab\xea\xf8\xa4\xe7\x17\x2f\xa0\x94\x58\xf5\xe7\xff\xb9\x44\x05\xb5\xb0\x96\x68\x6d\x44\xb5\x46\xdb\xb0\x9f\x10\x77\x28\x75\x55\xe9\x2d\x0d\xbb\x5d\x8d\x90\x6b\xb5\x41\x63\x1b\xbb\xd4\x68\x4a\x6d\x56\x58\x9c\x85\x25\xc0\x2d\xbc\xd4\x2c\xdb\xfd\xbb\x6d\x2f\xbb\x35\x7e\x0b\x3f\xb5\x30\xe7\x22\xbf\x26\x92\xde\x6b\xa5\xc8\xf1\xe6\x0e\x6e\x03\xee\x93\xa1\xbf\xcf\x7d\xde\x96\x08\xb8\x73\xad\x2b\xe8\xfd\xdd\xc2\x8f\x5a\x57\x28\x54\xef\xf9\xb0\x7c\x4b\x22\xe0\xd2\x1a\x16\x68\xac\x0f\x8f\xb2\xd2\xc2\x59\x3f\xff\x62\xbd\x9a\xa3\xe9\xeb\xf3\x22\xa7\x4f\xef\xc5\xb5\xce\x90\x3f\x7a\xf3\x2f\x8f\x3c\x1f\x96\xef\xe3\xbe\x7b\x2f\x95\xfb\xa1\x3f\xff\x95\x72\x3f\x3c\x37\x46\xec\x0e\x9e\x0f\xcb\x1f\xc1\xfd\xf6\x74\x08\xf7\xdb\xd3\x1e\xf0\x31\xf9\x23\xb8\xdf\x7f\x37\xe1\x0f\x1d\xdc\xef\xbf\x3b\x86\x7b\x9c\x6e\x0b\x77\x3d\xb0\xb0\x5b\x78\x2b\x87\x0c\x71\x4c\xfe\x18\xee\xe1\xc2\x18\xb7\x6f\x88\x63\xf2\xc7\x70\xd9\x10\xeb\x66\x89\x8c\xdb\x37\xc4\x6d\x47\xea\xd3\xb8\x3e\x22\xbf\xff\xee\x80\xef\xcf\xfc\xf4\x00\xf8\x98\xfc\x51\xdc\x83\x48\x0f\xb8\xa7\x4f\x8f\xe1\x1e\xdd\x19\x11\x57\x54\x15\x68\xb7\x44\x03\xb6\x92\x39\xda\x38\xbf\x1f\xbb\xb0\x8f\x87\x26\xcb\x7c\x02\x97\xe6\xdb\xfe\x7c\x8b\xc8\x9a\x3a\xe9\xee\xd8\xf3\x3e\xee\xbe\xc2\x1c\xd8\x21\x3c\x3f\xd4\x47\xf2\xe3\xe9\x74\xda\x62\x9d\xc1\xa3\x0f\x76\xfa\xdb\xfc\x03\xe6\xae\xc1\x75\x72\x85\xd3\x3f\xe4\x0a\x0f\xe6\xbf\x10\x6e\x88\xcd\x11\xf9\x3e\xdf\x27\xc3\xa3\x20\x95\x75\x42\xe5\xa8\x4b\xb8\xd0\xc5\x3e\xaf\xb7\xa8\x7d\x12\x77\x25\x6a\x3b\xa1\x2c\xb5\xce\x9d\x1d\xc6\x6d\xc1\x78\xf9\x77\x9c\xd3\x86\x1d\x78\x1b\x4a\xd1\xf3\xa2\x90\x64\x47\x2a\xd7\x13\xdf\x0b\x88\xa0\x85\xca\x98\x13\x52\x51\x5a\x14\x6d\x9e\xbe\x4a\x4e\x40\x2b\x2a\xde\x4b\x5f\xee\x1c\x2a\x07\xba\xe4\x62\x48\xc3\xb0\x95\x55\x05\x73\xf4\x75\x13\x8b\x6e\x49\xf5\xb9\x7e\x43\xbe\xa7\x92\x26\xa6\x49\xdd\x34\x28\x09\x71\x0a\x7a\xa4\x05\x11\x49\xa0\x09\xdc\xfa\x8d\x89\xf6\xd2\xad\xd6\x44\x3a\xdb\x54\xf5\x7f\xa1\x2d\xe9\x37\x22\xf0\x1c\x94\xac\xa0\xd6\xde\xb2\x24\xb9\x67\x8c\x7f\xad\x45\xd5\x5d\xee\x89\x85\x54\xad\xab\x2a\x9d\x46\xb9\x5c\x28\x50\xda\x91\x7d\xd6\x64\x1d\x41\x2b\x5d\x89\x1a\xae\x71\x37\x4d\xfc\x86\x08\x92\xec\x8a\x9b\xb0\x48\x78\x14\x1e\xdf\x79\x3b\xbd\x44\x07\x06\xdd\xda\x28\xeb\x2d\xcf\x42\x27\xbe\xcb\xab\xd1\xb8\x1d\xf7\x72\x34\xb4\x90\x1b\x54\x0c\x4f\x3b\x04\xc6\x3a\x62\x65\x04\x33\xbe\xc6\x5d\x28\x81\x59\xa3\xe4\x26\x80\x83\x9e\x06\x1b\x07\xc9\x2c\xe8\xbf\x44\x07\xd4\x16\x2d\x82\x7e\xdf\x1b\x05\xc3\xfd\x53\x32\x97\x1d\x32\x93\x80\xd9\xd9\xcd\x37\x7b\x42\x41\x3a\x88\x45\x5e\x2f\xb0\x42\x87\x60\x70\xa5\x37\xf8\x45\xa6\x61\xa4\x8e\x75\x5a\xda\xf7\xa3\x51\xf3\xaf\xa8\x16\x6e\x39\xec\x94\xb4\xf2\x83\x69\x43\x61\x12\x1a\x45\xc7\xfb\x43\x2a\x37\xc0\x80\x11\xc7\x19\x0d\x0f\x78\xa4\x19\x66\xfd\xaf\x54\x81\x1f\x3b\xea\xe5\x89\x5b\x02\x56\xb8\x0a\x3b\x54\x28\x4e\xd5\x03\xaa\xfc\xe4\xb1\x24\x4d\x9f\x0a\x82\x20\xd6\x0a\x02\xd6\x6a\xd1\x7d\xb6\xca\x38\x99\xb5\x3e\xc0\xdb\x41\xfa\xc0\xe1\xb4\xf5\x21\xe7\xfd\xdf\x36\x39\x67\x81\x43\x57\x2b\xb1\xc2\x01\x2e\x04\x32\xa6\xb1\x26\xf6\x84\x59\x58\xe8\xd5\x92\xa3\x86\x69\x00\x78\xe6\x74\x3a\xdd\xbb\x65\xa3\xaf\xb1\xc7\x90\x32\x15\x56\xe5\x14\xfe\x58\x4a\xcb\x19\xb3\x14\xb2\x02\x59\x82\xf4\xc9\x84\x72\x84\x68\x4a\xe0\xa0\xcb\x08\x78\xfc\x99\x44\x5b\xb3\x5a\x24\x2f\x70\x0b\xb9\x4f\x95\x94\x8d\x14\x6e\x9b\xda\xc2\x99\x5d\x5a\x2e\xd5\x31\xdf\x0e\x92\xee\x32\x86\x71\xae\x15\xa7\x30\x6d\xb2\x01\xfe\x17\xb8\xfd\x5c\xf2\x71\x4a\x8b\x39\x9d\x41\x06\xf6\x5c\x77\x7b\xf9\x03\x89\xc8\x73\x6d\xfc\xf1\xb2\x5b\x90\x0e\x8f\x6d\x03\x54\x49\xc9\x38\x63\x98\x3e\xab\x30\x1a\xb6\x04\x9f\x25\xee\x63\x14\x8e\x1c\x5f\xc0\x89\x15\x8d\xb3\x08\xd5\xe7\xd5\x48\xc4\x40\x1c\xaa\x18\xbd\x3c\xf4\x60\x4e\x30\xae\x85\xb1\xf8\x4a\xb9\x21\xef\xbe\x52\xee\x68\xe2\xe2\xb1\x86\xd5\xe9\xd3\x87\xf0\x3a\x7d\xfa\xef\x31\x3b\x7d\xca\xdc\x4e\x9f\x0e\xb3\xf3\xe3\xcc\xef\xad\x7c\x10\xc1\xf5\xbf\xc9\x90\x75\x8e\xb3\x88\xda\xe7\xd8\x48\x30\x49\x7f\x30\xb8\x97\x63\x3c\x24\x7c\x26\x49\x0f\x3e\x44\xd3\x0f\x8c\xb3\x06\xb7\x4f\x33\x4a\x34\xae\xe6\x4d\xfe\x10\x77\xc7\x74\x30\x85\x4b\x44\x70\x62\x5e\x51\x6d\x80\xd8\x2d\xe6\x7a\xe5\x4b\x0c\x35\x86\x05\x3a\x21\xab\xa1\x3d\xd2\x68\x64\x77\x37\x9d\xf0\xa0\xd3\x1b\xc9\xe0\x78\x65\x45\x39\x48\x95\x3a\x36\xe5\x7d\x53\x3b\x33\x81\xed\x52\xe6\x4b\xdf\xd6\xcd\xb1\xb5\x8c\x8d\x14\xb0\xf6\x18\xd3\x37\xdc\x2c\x4e\xe1\x42\x3b\xcf\x43\x15\x58\x78\xea\xf5\x7a\x5e\xc9\x9c\x1a\xc1\xa1\x30\xf0\xb3\x43\x18\xd4\xce\x0c\xc5\x41\x14\x61\xce\xff\x69\x8c\x36\x80\x2a\x17\xb5\x5d\x57\x3e\x9b\xb7\xfc\x8b\x34\x6a\x29\x79\x6b\x8b\xdc\x1d\xaf\x8d\xc2\x82\x28\x69\x10\xf0\x52\x43\x2d\x94\xcc\x7d\x5b\xbc\x12\x3b\x5a\x8f\xc1\x5c\x6f\xd0\x60\x31\xa1\x02\xea\x53\x96\x82\x47\xac\xc7\x2d\x85\x83\xa5\xf6\xb7\x66\x4b\xec\x69\x8a\xc5\x82\x7b\x5a\x9e\x12\x4e\x17\x37\xc9\x28\xac\x32\x69\x13\x6f\xdb\x7a\x85\xd6\x92\xa3\xc3\xc1\xa2\xb5\xa6\xe2\xb8\x26\x36\x21\x1a\x13\x28\x66\x0c\xdc\x4a\x92\xc9\x28\x98\x30\x3d\x04\x39\x83\x14\x1e\xd3\x47\xdf\xe9\xa6\x41\x7f\x9a\x35\x69\x34\x89\x09\x5e\xe4\xd7\x1d\xaa\xd6\x3f\x69\x9a\xcb\x2f\x64\xec\xf1\x87\x18\x37\xd4\xbc\xbe\x3e\xb1\x97\x95\x9e\x8b\xca\xf7\x39\xb6\x7b\x02\x59\xf0\x48\x08\xdf\x71\xba\x95\xaa\xd0\xdb\xd4\x47\xe0\xdc\xe8\xad\x8d\x77\x70\xe9\xcb\x5f\x7f\xfb\xf1\xf9\xaf\x3c\x42\x47\xd5\xe9\x07\x9b\x4d\x93\x8d\x30\x11\x3d\xba\x8d\x14\xbe\xd6\xc5\xba\xc2\xa0\x70\x7f\x06\x08\xeb\x4f\x57\x7e\x38\x85\x8d\x30\xd2\x6f\x5f\x8b\x8e\x4e\x5f\x01\x77\x0a\xff\x25\x95\x3b\xe3\x83\x04\xb0\xb0\xbf\x97\x35\x8e\x9b\xb6\x93\x0f\x76\xca\x2a\x78\xd9\x3c\x66\x69\xe1\xfb\xaf\x17\x62\x85\xe9\x84\x5a\x88\xec\x84\x89\x06\x56\x6d\xa2\x6f\x55\x81\xa5\xa4\x48\xdf\x73\x6d\x79\x84\x69\xa7\xeb\x28\x95\x32\xd0\x7e\x56\x1b\xeb\x05\xce\xd7\x8b\x05\x1a\x58\x50\xcb\x9b\xeb\x55\x2d\xab\xc3\x33\x2e\x35\xfc\x45\x90\x7b\x96\x52\x7c\x38\xdf\x10\x07\x77\x47\x88\x71\x06\x37\xad\xcc\xa8\x44\x15\x1a\x9f\x4e\x0f\x1f\x86\xfa\xa7\x5e\xde\x7f\x06\x6b\x83\x16\x95\xb3\x20\x1f\x92\x60\xba\xaa\xb8\xf7\x1e\x68\xbd\x9a\xa8\x53\xb2\x0a\xf1\xc5\xd7\xe8\x2a\x87\xad\x11\xb5\x6d\x77\x7a\x14\x3a\x6c\x59\x91\xe7\x68\xe3\x3b\x82\x78\x5f\xae\xcb\x03\xdb\x50\x3f\x99\x72\xc0\x09\xb3\x58\x93\x69\x6c\x4a\xa7\xb0\xad\x36\x45\xcc\xe3\x51\xdd\xb8\x54\x7c\xb1\xe3\xbb\xd0\x40\xd0\x77\xd9\x3c\x11\xde\xbd\x6f\x32\xe6\x3d\x6b\xe1\x18\xe6\x5e\x3d\xfd\x7a\x15\x14\xa4\x93\x43\xa3\x94\x2a\x8b\x9b\xea\xbf\x71\x67\x3b\xfe\xb8\xa6\x07\x21\xc4\xf9\x48\xd1\xbf\x8e\xe0\x05\xd0\xd4\x76\x3a\x7f\xf7\x7e\xbf\xa5\x65\x09\x1a\xce\xcf\xfd\x55\xc2\xed\x2d\x7f\xde\xc7\xdb\x4d\x32\x6a\x9b\x7f\x74\x97\x8c\x04\x9c\x9d\x47\xfe\x7e\x37\x30\x6a\x9a\x85\xd5\x10\xad\x74\x02\x3a\x4b\x46\x96\x44\x69\x71\xe3\xa8\x71\x02\xa2\x39\x2c\x66\xc9\xc8\xbf\xf4\x21\xa1\x6f\x9e\x81\x84\xff\x68\x0d\x3e\x03\xf9\xf8\xb1\x57\x6f\xdf\xc9\xf7\x70\x0e\xa2\x39\xf1\xed\xb3\x0d\xd1\x09\xec\x6c\x2b\x34\xe2\xdb\x95\xfd\x31\xa2\x1f\xb1\x5c\x2a\x97\xc2\xfa\x18\xaa\x29\xed\x94\xbe\x90\xc4\x9d\x8f\x45\x73\x7b\xa3\x4b\x0a\xe8\xb7\xd6\x0f\x55\x32\x97\x8e\xb6\x9c\x43\xe3\x03\xc7\xf2\xc7\xd6\x5b\x9f\xf0\x4a\x27\x54\x98\xc1\xb7\x39\xfb\xc0\x0a\x64\x3f\x11\xfe\x1b\x32\xd0\xe1\x66\xc9\x92\x91\x3e\xea\x08\x3a\x9c\x90\x00\xa7\xa7\xab\xab\xb8\x73\xaf\x78\xf1\x57\x57\xe9\x04\x36\x59\x32\x8a\x9c\xcf\xce\x61\xc3\x10\xad\x83\x52\x9a\xc5\xf2\xe3\x85\xd2\x01\x77\x85\xa1\x01\xa7\xad\xbc\xe7\xc3\x70\x74\x5c\x32\xa2\x68\x5b\x31\x6c\x7d\xbd\x68\x15\x0e\xf8\xea\x1c\xd2\x14\x6e\x60\x36\xf3\x87\xb7\xe8\x83\x64\x34\x1a\xe5\x5a\x39\xa9\xd6\x98\x8c\xc8\xdf\x61\x55\x01\x85\xce\xb9\x2d\x98\x09\xef\xcf\x78\x96\x6b\x02\xbe\x65\xcd\xd1\xf0\x16\xc4\x8f\x6c\x22\xf9\x37\xc6\x3b\x5d\x32\x92\xd7\x12\x19\x1b\x5d\xb7\x74\x65\x93\xb8\x14\xb7\xab\xd3\x6c\x02\xce\xac\x31\x6e\x02\x51\xd7\xd5\x8e\x00\xf8\x10\x4e\x4b\xbf\xeb\xc4\xab\xee\xa4\xb2\xfd\x1b\xc1\x2f\x8c\x59\x82\x6c\x87\xed\x84\x42\x94\x1a\x43\x34\x08\x92\xef\x32\xc7\xad\x1b\x43\x91\xc5\x30\xf5\x19\x72\x12\x90\x8b\x10\xe0\x96\xf0\xf6\x41\xee\xbf\x8e\x97\xce\xd5\xf6\x6c\x36\x2b\x70\x83\x15\x75\x1a\xd3\x95\xfe\x5b\x56\x95\x98\x6a\xb3\x98\xa1\x7a\xf2\xf6\x72\x56\xe8\xdc\xce\xfe\xc4\xf9\x6c\xbf\x8a\xd9\xef\x58\xa2\x41\x95\xe3\x8c\x4d\x7f\xc5\x4e\xb1\x33\xfe\x3f\xe3\x9c\xf3\x26\x34\x2f\x19\xe9\x8a\xcb\x53\x5a\x3d\xc1\xd5\x1c\x0b\x2a\x26\xcd\xfe\x0c\x3b\x8b\xb7\xe7\xff\x70\x86\xe7\xb4\x1f\x9a\x5e\x7e\x3b\x1c\xec\x11\x97\x12\x56\xc6\x5d\xe7\x12\x57\x16\x37\x15\xda\x66\xe1\xdb\x25\xaa\x06\x65\x02\xf3\xb5\xa3\x3e\x53\xae\x48\xa7\x50\x8e\xaf\x5b\xc1\xe9\x84\x23\x15\xc9\xa6\xbe\xfc\xf1\x65\x45\x84\x09\x57\x48\x36\x38\xb4\x00\xad\x00\x45\xbe\x0c\xd0\x9d\xca\xd2\x78\xff\x1f\x24\x81\xfc\xd8\xe6\x4d\xee\x4f\x10\xc9\xa8\x6b\x72\x12\xf7\xfb\xe7\xba\xb9\x7b\x2a\x95\xdf\x48\xfe\x69\xdc\x4d\x99\xdf\x46\xba\xbe\x2f\xfd\xf8\x74\x71\x35\x81\x92\x04\x8d\x50\x0b\xf4\x70\x7e\x0f\x96\x63\x5d\x67\x61\x43\x7f\xa2\x96\x74\xf9\x51\x55\x99\xc0\xf5\x04\xfc\xdc\xbb\x1e\xff\x63\xa9\xce\xb3\xd7\xf5\xbe\xfe\x31\x7f\xce\x8e\xde\x51\x21\x1f\xde\x91\x4d\x68\x70\x36\x83\x5c\xe4\xbe\x63\x06\x01\x16\x95\x95\xd4\x36\xfa\xe3\x03\x9b\x26\x61\xa9\x2d\x42\xa1\xd5\x89\x83\xad\xf0\x51\x11\x02\x05\x84\xda\xc5\x03\xa2\xa5\xc3\xa2\xef\x18\xc2\x03\x9e\x69\x35\x4f\x06\xab\x9b\xbb\x3e\xa0\x75\x08\x7e\xcf\x34\xdf\xc1\x52\xa8\x82\xf4\xb8\x9d\xb7\x75\x1e\x7b\x6e\x12\x6d\x37\xdd\xa3\x51\x7d\xbd\x68\x49\x74\x53\x2b\x21\xd0\x91\xed\x8c\x32\x2c\xa7\x60\xb7\xab\xdf\x7d\xf3\x9e\x2a\xfd\xc9\xa3\x13\x76\x08\x49\x9c\x43\xfa\x28\xf5\x4e\x09\xce\x6b\xe7\xfa\x0a\xd5\xd8\xed\xea\x56\x8e\x8f\x48\x92\x91\xa6\x01\xc9\xd3\x3d\xe7\x91\xc7\xdf\x9e\xbd\xf7\xcf\xe6\x06\xc5\x35\x7d\xba\x8b\xf8\xf5\xf5\xe2\x0f\x5e\x17\x91\x7f\x0c\xe9\x94\x0e\x3d\x44\xe3\x31\xcd\x4d\x46\x3d\xe7\x7e\x4d\xe6\x3f\xe6\xce\x9e\x3f\x19\x7f\x9f\x78\x93\x11\x75\xd2\x21\x65\xc4\x36\xba\x5d\x01\xdb\x61\xe8\x5f\x43\xee\x23\x99\xaa\x96\x6d\x19\xb8\x29\x89\xcf\xe8\xf9\x57\x87\x0d\x53\x04\xdd\x57\x40\x8e\xe6\x5c\xab\x5c\xb8\x74\x02\x2b\xcb\xb5\x60\x36\x03\x59\x52\x18\x50\x2e\x12\xcd\x9b\x9c\xf0\x02\xc3\x27\xa2\xa2\x49\x73\xa5\xd1\xab\x78\x9f\xed\xa7\x62\x65\x31\xbe\xf9\x8a\xfb\x3e\xdc\xe5\xf2\x85\xe8\x52\x6c\x38\xc5\xf9\x15\x60\x6b\x01\x04\x43\xec\xb1\x4f\x3e\x68\x3b\x87\x70\xd6\xe1\xef\x54\xfd\xef\x5f\x17\x1e\x98\x87\x56\xc9\x34\x3b\xc8\xf9\x21\xf2\xdd\xff\x93\xd6\xe2\x20\xe2\x86\xeb\xfe\x43\x03\xf0\x01\x9d\xc8\xff\x69\x2b\xd2\x3b\x4b\x1c\x14\x9b\xec\x13\xbd\x8a\xdf\x33\xb1\x63\x91\x65\x0c\xc1\x7e\xb0\x1c\x78\x8d\xe5\x06\x9c\x36\xf2\x45\x20\x0c\xb7\x9c\x46\xe0\x5f\x95\xed\x53\x35\x16\x69\x16\xef\xb1\xd9\x48\x2d\x27\x79\x2f\x1d\xba\xa9\xfc\x2c\x37\x35\x7e\x5a\xa0\x8b\x5e\x3a\x74\xcb\x68\x93\xb7\x32\x42\x70\x4b\xae\xeb\xdd\xab\xf2\x77\xfc\x6b\x2d\x0d\x16\x8d\x47\xd2\xaf\x37\xa2\x0a\xbd\xf2\x30\x13\xa2\xd2\xf2\x4e\xc6\x2a\xee\x73\x3d\x29\xc8\xbb\x33\xef\xf7\xa8\x87\xbe\xcb\x3a\xeb\xb4\xfb\x75\x7e\xd8\x1c\x1a\x63\xb4\xd8\xf4\x57\x1a\x6b\x28\xb3\xf8\xb0\xf9\x47\x2c\x46\x5d\xf3\x5c\x1e\x35\xcf\x04\x16\x9b\x36\xf1\xbb\xd0\x17\xf4\x7b\xe5\x0b\xdc\xfa\xc4\xfc\xe3\xba\x2c\x8f\xb5\xca\x6d\x01\x9f\x31\x05\xcc\x77\x2e\xfc\xc8\x23\x74\x5d\x5d\x9c\xf1\x1c\xde\xbd\x27\x99\x4e\x14\xf0\x8f\x42\xfa\x3d\xd7\x9c\xce\x55\x65\x69\xd1\xd1\x20\xa3\xf2\x3a\xf9\x69\x9a\xf1\x3b\x85\x64\xc4\xef\x59\x0f\xa5\xc2\xdb\xd7\x46\x2a\x1e\x5f\x5b\x22\x22\x14\x1f\xff\x6d\xee\x39\x36\x0d\x91\x97\xa3\x3e\xc8\x2b\x8b\xff\x1f\x33\x6a\xbc\x29\x78\xcd\x7d\xbe\x95\xab\xba\x42\xff\x42\x9f\x2a\xe7\x14\x5e\xf9\xb7\x79\xcd\xa5\x8c\x7f\xdd\x6f\x97\xda\xb8\xa5\xff\xd5\x9b\x36\xfd\x33\x87\x85\xf1\x1c\x4b\x6d\xda\xb7\xf1\x59\xb8\x47\x7d\x7d\xe4\xd7\x1d\x7c\x37\xd9\xe1\xb0\xff\x89\xcd\x67\xb2\x08\xbf\xe7\x39\x4e\xe2\xb2\xfb\xd3\xa0\x84\x3d\x2c\x95\x74\x9c\x3f\xa8\xe9\xdf\x68\x59\x40\x81\xa2\x80\x5c\x17\x08\x58\xc9\x95\x54\xbe\xcd\x4a\x46\xde\xc7\xfe\xbe\xf3\xe6\x2e\x19\x5d\x51\xdd\x4b\xee\x92\xff\x0d\x00\x00\xff\xff\x14\x03\x9e\xa4\xae\x2a\x00\x00"),
		},
		"/nosync": &vfsgen۰DirInfo{
			name:    "nosync",
			modTime: time.Time{},
		},
		"/nosync/map.go": &vfsgen۰CompressedFileInfo{
			name:             "map.go",
			modTime:          time.Time{},
			uncompressedSize: 1958,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x4d\x8f\xdb\x46\x0c\x3d\x5b\xbf\x82\x3d\x55\x2e\x14\xe7\x9e\x62\x0f\x05\x7a\x29\xd0\x34\x40\xdb\x5b\x90\x03\x2d\x71\xac\x81\xe7\x43\x1d\x52\xeb\x2a\x8b\xfd\xef\x05\x39\xb2\x57\xde\x24\x45\x0f\xbd\xd9\x23\x0e\xf9\xf8\xde\x23\x67\xc2\xfe\x8c\x27\x82\x94\x79\x49\x7d\xd3\xbc\x7d\x0b\xef\x71\x02\xcf\x80\xd0\xe7\xd4\xcf\xa5\x50\x12\x88\x38\xc1\xc5\xcb\x08\x18\x73\x11\xff\x99\x86\x37\x7d\x4e\x2c\x98\xe4\x8d\xf8\x48\x10\x32\x0e\xdc\x01\x4b\x2e\xc4\x1d\x60\x1a\x60\xa0\x40\x42\x7c\xd0\x9c\xbf\x88\xa6\x64\x74\x04\x2e\x17\x88\x73\x10\x3f\x05\x82\x53\x2e\x79\x16\x9f\x88\x41\x32\xf4\x18\x02\xa0\x02\xf8\x9e\x21\x92\x8c\x79\xe0\x0d\x8a\xb0\x68\x2e\x4d\xf7\xe7\x48\xf0\x99\x4a\xbe\x62\x7d\xc4\xe0\x07\x2b\x4a\x71\x92\x5b\xd8\x4f\xf6\x3d\xce\x2c\x90\xb2\xc0\x91\xa0\xcf\x93\xa7\x01\xd0\x09\x15\x70\xbe\xb0\xc0\xcc\x74\x68\x64\x99\xc8\x82\x59\xca\xdc\x0b\x3c\x35\xbb\xa8\x4d\x7f\xf4\x49\xa8\x38\xec\xe9\xe9\xf9\xd3\xe6\x77\xf3\x6c\x54\xfd\x9a\x71\x80\x42\x32\x97\xc4\x20\x23\x29\x90\x99\x2a\x0b\x03\xf8\x64\x67\xca\x9d\x36\x8d\x70\xa6\xa5\x83\x5c\x20\xf9\x00\xde\x41\xca\x9a\xa3\x5e\xf1\x0c\x53\x21\xa6\x24\x87\x6b\x83\xf9\x0c\x85\x78\x0e\x02\x3e\x0d\xbe\x47\x21\x86\xcb\x48\x32\x52\x59\x2f\x5d\x90\xc1\xe5\x39\x6d\x4b\x1d\x1a\x37\xa7\x1e\xda\x08\x3f\xbc\xc7\x69\x6f\x10\xdb\x33\x2d\xb0\x41\xbf\x87\x76\xad\xfa\x72\xd6\x69\xbd\x63\xce\x61\xaf\xcd\xdb\x67\x3b\x7a\x80\x78\x88\x1f\xcf\xb4\x7c\x6a\x76\xb5\x53\xb8\x7d\x5c\x59\xf8\x43\xdb\x05\x26\xd9\x72\x70\xeb\xf8\x35\x20\x8b\x6e\x8d\x8a\x2f\x40\x58\x6d\xef\xb4\x24\x3c\x3c\x18\x4f\x4f\xcd\x6e\x67\x7f\x21\xe2\x99\xda\x7f\xd1\x64\xdf\xec\x9e\x9b\xdd\x15\x2d\x3c\xd4\xf4\x1b\xa5\x3e\x94\x8a\x74\x2b\x18\xfd\xed\x59\x7c\x3a\x6d\x50\xeb\xb1\x11\xe6\xee\x24\xf9\xa0\xc4\x5f\x3c\x53\x07\x5e\x56\xa3\x9b\xe5\xb6\xe9\x4e\xfe\x91\x56\x82\x6e\x3a\xea\x68\xd0\x70\xd3\x92\x41\x8a\x76\xed\x36\x64\xa9\x90\x35\xac\x03\x87\x81\xed\x73\x75\xd1\xd7\xf4\x5c\x1b\xf9\x26\x89\x2d\xf6\x32\x63\xb8\x97\x77\x85\x71\x93\xd8\xbb\x17\x21\xe1\xdd\x8b\xcc\x3f\xea\x7f\x65\xfd\x5e\x6d\x05\x6d\x04\xff\xcf\xf2\xbc\x2a\x63\xdd\xaf\x9a\xfd\x6c\x0b\xe4\xba\x47\xfe\x8b\xb7\xea\x8d\x2f\xed\xfe\x55\x57\xd5\xc2\x86\xaa\x96\x68\xe3\x21\x76\x9a\x76\xbf\x02\xf8\x1d\xd3\x89\x6c\x2b\x31\x38\x60\xfa\x6b\xa6\x24\x1e\x43\x58\x0c\x02\x61\x3f\x9a\x53\xd4\x05\x15\xd9\x6a\x98\xbb\x79\xd4\xf5\xe7\xc0\xdd\x7c\x62\x2d\x76\x50\x2c\x39\x4b\x9e\x6a\x6b\x5e\xa8\xa0\xf8\x9c\xae\xdb\xab\x56\x1f\x32\xb1\x6d\xaf\x44\x3d\x31\x63\xf1\x61\x81\x3e\x97\x42\x3c\xe5\x34\xe8\xda\xc4\xa4\x27\x89\x3d\x8b\xd6\xe6\x84\x13\x8f\x59\x20\x57\x8b\xd9\x3a\xd5\x84\x7d\x4e\x1a\xc0\xef\x20\x65\xc3\x7d\xf1\x21\xe8\x56\x7c\xf4\xec\x85\x06\x88\x3a\x1d\x32\x62\x82\x9c\x7a\xea\xe0\x38\xcb\xbd\x4f\x8d\xf8\xb4\xe8\x65\x4d\xa8\x2b\xbd\xae\xba\x5c\x56\x99\x86\xbb\x7d\xdd\xad\x4d\x44\x5c\xa0\x90\x0b\xd4\x8b\xdd\x8f\x38\x4d\x3a\x74\x75\xdc\x50\xae\x09\x5d\xc9\xd1\x02\xa6\xec\x93\xc0\x30\x17\x8d\xd2\xfa\x2f\x52\xdc\xd3\xa3\x99\x8f\x04\x1f\xda\xdf\xf6\xf5\x81\xd2\xe0\x34\xc7\x23\x15\xed\x9f\x02\x45\x6d\x79\xbb\x8b\x49\x47\xd4\x6f\x14\xb1\xca\x36\x75\xf5\x5d\xb0\x97\xcf\xde\xb6\x4d\x26\x73\xc1\x6b\xbf\x19\x86\xd6\x81\x9e\x7e\x73\x1a\x6f\x13\xa7\xdd\x9e\x3b\x78\xd4\x69\xab\xea\xab\x23\xd5\x8a\xde\xc1\x77\xae\xd5\x6f\x16\xb8\xdb\x1d\x0b\xe1\xb9\xd9\xa9\x37\xf5\xad\xf9\x27\x00\x00\xff\xff\xe8\x19\x65\x16\xa6\x07\x00\x00"),
		},
		"/nosync/mutex.go": &vfsgen۰CompressedFileInfo{
			name:             "mutex.go",
			modTime:          time.Time{},
			uncompressedSize: 2073,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\xcb\x6e\xdb\x30\x10\x3c\x4b\x5f\xb1\xc9\xc9\x4e\x62\xa5\xbd\xb6\xf5\xa1\x68\x81\x22\x40\x7a\x09\x50\xe4\x4c\x53\x2b\x99\xb0\x44\x1a\x24\x55\xd5\x4d\xf2\xef\xc5\xf2\x21\xcb\x92\xec\xc4\x2d\xaa\x93\xb0\xe4\xce\xce\xec\x0c\xb8\x65\x7c\xc3\x4a\x04\xa9\xcc\x4e\xf2\x34\xbd\xbd\x85\xef\x8d\xc5\x5f\x20\x0c\x30\xc8\x9b\xba\xde\x41\xbb\x16\x7c\x4d\x05\xa9\xe4\x62\x55\x29\xbe\x11\xb2\xcc\x52\xbb\xdb\x62\xb8\x6c\xac\x6e\xb8\x85\xa7\x34\xa1\x53\xcc\x61\xa5\x54\x95\xbe\x38\xb8\x7b\xc5\x37\x40\x65\x03\x75\x06\x77\xd6\x23\xeb\x46\x2e\xac\xa8\x11\x50\x6b\xa5\x41\x14\x50\xbb\x83\x4a\x23\xcb\x77\xe0\x61\xb2\xb4\x68\x24\x87\x59\x0d\x57\x6e\xce\xdc\x81\xcd\xe6\x34\x88\x3a\xb2\x30\xed\x29\x4d\x92\x2d\x93\x82\xcf\x2e\xbd\x8e\x0f\x50\x77\x22\x0e\x10\x2f\xe7\x69\xf2\x92\x26\x5d\xe7\x12\xac\x6e\x30\x30\xfd\x21\xa9\x0a\x8d\x7c\x2b\x5b\xa9\xec\x51\xa6\x1e\xac\xe3\x7a\x71\x8a\xac\x9f\x08\xaa\x08\x7f\x98\x7b\xfe\x63\xb6\x05\xab\x4c\xa4\xfb\xf0\x78\x96\x53\xf1\xfa\xde\xab\x56\x0b\x8b\xf7\x1e\x9a\x3e\x67\x5a\x42\xeb\xa2\xe2\x17\xd5\x48\x8b\x1a\x84\xb4\x13\x4e\x42\xa1\x34\x10\x00\x0d\x38\xb1\x27\xdd\x8e\x4d\x70\xbd\x54\x10\xb2\x84\x1e\x4c\xd8\xa1\x6e\xe1\x2a\x90\x1d\x18\xae\xdb\x6c\xc8\xee\x62\x09\xef\xe0\xf9\x99\x8e\xfa\x72\xce\x4e\xc4\xa0\xff\x54\x2e\x74\x7b\x9e\xf8\x7d\x4a\x0e\xfa\xa6\xd4\x0e\x43\xf3\xba\xaa\x57\xa2\x33\x92\x75\x10\xa0\x91\xa1\xc1\x94\xff\x69\xe8\xc3\xd0\xd1\x7f\xb5\x6d\x90\x88\xeb\xeb\xa8\xae\xb3\x2d\x57\x48\x5a\x8c\x90\x65\x85\x41\x35\x67\x55\xf5\x11\x84\x05\x77\x48\x16\xb1\xa2\x40\x6e\x41\xd9\x35\x6a\x30\xa2\x6e\x2a\xcb\x24\xaa\xc6\x38\x65\xa8\xcd\xd9\x4e\xc7\x6d\x4e\xae\x61\x60\xf5\x44\xb4\x97\x14\xed\xbf\xb2\x7c\x80\xb4\x58\x84\x95\x3c\x32\x61\xbf\x69\xd5\x6c\xdf\xfa\x66\xec\x1b\xf6\xaf\x06\x1f\xbd\x0b\x9f\xf3\x1c\x58\x9e\x1b\xc8\xb1\xb2\xec\x26\x20\xd6\x6c\x07\x2b\x04\x89\x25\xb3\xe2\x27\xde\x80\x55\x60\xd7\x7d\xcc\xbb\xc2\x15\x22\x60\xe9\x9c\xe8\xae\x13\xaa\x53\x6e\xe2\x02\xdb\x12\xae\xba\xee\x39\x5d\x98\xb9\x89\x44\xc5\xed\xb1\x2d\xb3\x08\x76\xbd\xf4\x6c\xdc\x72\x7b\xf5\x4f\x87\x3b\xf5\x1b\x8d\x43\x7b\xdc\xc2\x7d\xbf\x53\x2f\xf3\xab\x92\x08\x39\x72\x8d\x35\x4a\x6b\x06\x62\x42\xc3\x11\xae\xd4\x3b\x8b\x1c\x89\xf8\xe2\xfd\xbc\x67\x4a\x10\x4a\x49\x9a\x44\x8d\xe1\xfa\x8d\x5a\x1d\x99\x40\xbf\x5d\x9a\x7a\x82\x2f\x96\x53\x8a\xc7\x13\x22\x7c\x54\xfc\x27\x00\x00\xff\xff\xec\x95\x29\x83\x19\x08\x00\x00"),
		},
		"/nosync/once.go": &vfsgen۰CompressedFileInfo{
			name:             "once.go",
			modTime:          time.Time{},
			uncompressedSize: 1072,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x53\xcb\x92\xda\x40\x0c\x3c\xdb\x5f\xd1\xb5\x27\x9c\xa2\xe0\xbe\xa9\x1c\x52\xc5\x65\x4f\x39\xe4\x0b\xc4\x58\x03\xca\x0e\x1a\x32\x0f\x58\x67\x8b\x7f\x4f\x69\x6c\x08\xb9\xd9\x92\xba\xd5\x6a\x69\xce\xe4\xde\xe9\xc0\xd0\x98\x27\x75\x7d\xbf\xdd\xe2\x87\x3a\x86\x64\x90\x22\xee\x7f\xb1\x2b\x28\x47\x2a\xb8\x4a\x08\x38\x73\xf2\x31\x9d\xc0\x1f\xe4\x4a\x98\x10\x95\x41\xae\x48\xd4\x4d\x5f\xa6\x33\xcf\xe0\x5c\x52\x75\x05\x9f\x7d\x37\x46\xd1\x03\xf6\x31\x06\xfb\x56\xc6\xfc\x7d\x6b\x8d\x76\x11\x8e\x42\xc8\x28\x47\x86\xaf\xda\x78\xe0\x21\x1e\xa4\x23\xa2\x86\xc9\xbe\x77\xd1\xd4\xec\xd9\x98\xac\x9e\x47\xf8\x98\x0c\x64\x24\x5e\x52\x2e\x28\x72\xe2\x25\x2a\x19\xa2\xb9\x90\x09\x89\xbe\x09\xda\xe0\x4d\x11\xcb\x91\x13\xae\x31\x8d\x79\x8d\x83\x5c\x58\x0d\xde\x5d\x28\x21\x5a\xad\x15\x5a\x44\x7c\xfb\xdf\xec\xe2\xca\x0f\xd6\x79\xe9\x79\xaa\xa1\xc8\x39\x70\xeb\x95\xd7\xb3\xbc\xa6\xbc\x29\xb0\xaa\xd9\x23\xd1\x4b\x7c\x67\xf8\xb5\xb1\xf1\x85\xd5\x28\x3d\x8e\x94\x41\x18\xc5\x7b\x4e\xac\x05\x17\x0a\x95\x21\x0a\x26\x77\x6c\x20\x47\xcd\x48\xe0\x3b\x94\xaf\xcf\x53\x3c\xaf\x25\xf1\xef\x2a\x69\x31\xa1\x61\x1f\xd6\x95\x08\xfe\x60\x57\x0b\x6f\xfa\xed\x76\xb1\xb8\xf9\x51\x58\xc7\x05\x22\x2a\x45\x28\xc8\x1f\x9a\x31\xb6\xdb\x53\xcd\x05\x7b\x46\xaa\xfa\xb4\x5a\x33\x0e\x3f\xc5\xfa\x36\x05\x92\xa1\x12\x68\x14\xb7\x86\x14\x9c\x68\x32\x8c\xb2\xe3\x9c\x29\x4d\xd6\xbe\x66\x06\xfd\x13\x14\xa4\x70\xa2\x60\x19\x47\xe7\x52\x13\xdf\xd7\x46\xe9\x50\x4f\xac\x25\x5b\x8e\xfe\x1b\x61\xcf\x8b\x85\x23\xf6\x13\x76\xf1\xb5\xed\xc9\x45\xf5\x72\xd8\x3c\x56\x53\xd5\xad\x06\x7c\x62\x89\xdb\x54\x2b\x2f\x81\x95\x4e\x3c\xe0\x36\x2c\x06\xbc\x99\xf5\x8e\x6a\xe6\x6c\x66\xcc\xf4\xf3\x46\xdb\x10\xf3\x55\x93\x8a\xdb\x3c\x23\x5a\x24\xaf\xdb\x89\x46\xcd\x32\x72\xca\x56\x5e\x22\x8e\x74\x61\x24\x2e\x35\x29\x8f\x5f\xe1\x6b\x1b\x6b\x3e\xe4\xd8\xae\x75\x4e\x1a\xd7\x55\xca\x31\xd6\xf9\x38\xec\x7c\x7d\x6b\x62\xda\xb1\x8a\xf8\x62\x2b\x1d\x60\xd3\x60\x9e\x67\xb0\x37\x63\x07\xb8\x69\x8f\xe5\xb3\xef\xba\x85\xac\xbb\x3d\x12\x46\x64\x99\xa6\x71\xf5\x32\xbf\xdc\xd7\xfb\x6b\xe2\xb1\x75\x15\x85\x7f\x19\x1a\xec\x8e\xf9\x86\x92\x2a\xf7\xdd\xc8\x9e\x13\xee\x06\xf6\xdd\x53\x81\xa7\x90\x79\x89\x28\x3f\x10\xb7\xd5\xd0\x77\x7e\x35\xf4\xb7\xfe\x6f\x00\x00\x00\xff\xff\xf9\x72\xbe\xa9\x30\x04\x00\x00"),
		},
		"/nosync/pool.go": &vfsgen۰CompressedFileInfo{
			name:             "pool.go",
			modTime:          time.Time{},
			uncompressedSize: 2130,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x55\x3f\x93\xdb\xc6\x0f\xad\x4f\x9f\x02\xbf\xea\x77\xca\xe8\x74\x49\xeb\x99\x2b\x32\x29\x1c\x37\x89\x8b\x74\x1e\x17\x10\x09\x8a\x88\x97\x0b\x06\xc0\x4a\xa2\x3d\xf7\xdd\x33\x58\xfe\x39\x39\xee\x44\xee\xf2\xe1\xe1\xbd\x07\x68\xc4\xe6\x0b\x9e\x09\xb2\xd8\x94\x9b\xdd\xee\xf9\x19\x7e\x85\x8f\x22\x09\xd8\x00\xc1\xc8\x41\x3a\x70\x1a\x46\x51\xd4\x09\xe4\xf4\x37\x35\x6e\xe0\x3d\x3a\x0c\x38\xc1\x89\x80\x73\xcb\x17\x6e\x0b\xa6\x34\x81\xe1\x85\x5a\xc0\xdc\x06\x94\x92\x2b\xd3\x85\xda\xe3\xee\xf9\xb9\x62\xe7\x09\xd8\x69\x00\x73\x51\x6a\x81\x33\x78\x4f\x73\xc1\x05\x4d\x69\x90\x0a\x51\x5c\x06\x74\x6e\x2a\x2c\x3a\x60\x9e\xc0\x79\x20\xb8\xb2\xf7\x52\x3c\xf0\xb2\x38\x77\xdc\xa0\xb3\xe4\x23\x7c\xe8\xde\xd0\x7a\x49\xad\xd5\x47\xc9\x69\x02\xa5\x8e\x94\x72\x43\x70\xed\x29\x8a\xb2\x41\x8f\xe3\x48\xd9\x0e\x71\x2b\xc0\x2a\xb1\x81\xcf\xbd\x07\x8f\x96\x30\x25\x69\xd0\xef\xd8\x6f\xca\x18\x76\x04\x9d\x28\x14\x23\x38\x4d\x30\x94\xe4\x3c\x26\x82\xb3\xa8\x14\xe7\x4c\x06\xc6\xf1\x16\x33\x49\xb1\x34\xad\x18\x81\xf0\x7f\x83\xb1\xe8\x28\x46\x81\xe5\x02\x0d\x36\x3d\xc1\x56\x0f\x4e\xc5\xa1\xe4\x62\xa1\x90\xd3\x60\xb5\x54\x42\x27\x05\xa5\x62\x74\x98\xc5\x4d\x4c\x17\xce\x67\x18\x95\xcc\x8a\x46\xab\xb5\xe3\x33\xea\x29\x4c\x6d\x24\x25\x6a\x5c\xf4\x08\x7f\x85\x5f\x6c\x07\xe0\xb0\xed\x0b\x59\xfc\x20\xb4\x09\x5c\x02\xec\x54\x38\xb5\x40\x5d\xc7\x0d\x53\xf6\xd0\x44\x09\xdb\xa7\xb9\x51\x25\x82\xc4\xe6\x76\x84\xdf\xe5\x4a\x17\xd2\x0a\xc4\x16\x06\x80\x15\x76\x3c\xa5\x59\x10\x4c\x29\xf0\xee\x3e\xd9\xac\x07\x1c\x47\x95\x51\x19\x9d\xaa\x70\xd2\x01\x6e\x92\xba\xc0\x80\x39\x68\x23\x9c\x55\xca\xf8\x7d\xf0\xaa\x0e\x81\x63\x9c\x28\x7b\x24\xad\xc7\x88\x10\x0e\x92\xcf\x11\x38\x18\xc5\x29\x3b\xd7\xbc\x54\x99\xda\xb0\xa6\x91\xdc\x14\x55\xca\x1e\x41\xa5\x91\x72\x4b\xb9\x86\xa7\x49\xd1\xaa\xcd\x34\x96\x41\x38\xce\x7c\x46\x95\x0b\xb7\x14\x23\x70\xc5\xd0\x28\xca\xa8\xf3\xd7\xcd\x25\x96\x0c\x72\x21\xed\x09\x6b\xd4\xb1\x51\x31\x8b\x16\xa6\x15\xf8\xae\x73\xba\xe1\x10\xf1\x90\x0e\xce\x22\xed\x8f\xdd\x2f\x83\xd0\x0d\xbe\x32\x39\xc0\xb5\xe7\xa6\x87\x01\x39\x3b\x72\x36\xc0\x00\x6b\xa7\x8c\xc3\x3c\x14\x4f\xc6\x5f\xa9\x9d\x47\xe9\x3f\x53\x5a\x7c\x2c\x0e\xa7\xd2\x75\xa4\x16\xee\xd3\x72\xcd\x1a\x4c\x64\x50\x72\x4b\x1a\x70\x49\xb0\x85\xc7\x3a\x13\x95\xfa\x5d\x7e\x51\x09\xb0\x71\xbe\x50\x9a\x60\x54\xce\xce\xf9\xbc\xaf\x4a\x5b\xaf\x9c\xbf\x58\x9d\xa5\x40\xf9\xa7\x30\x59\x43\xd9\xd7\x96\xff\x9c\xdb\x11\xef\x49\xa1\xc7\xdc\x1e\x00\xdf\x32\xb1\xf5\x14\xf6\x19\x8c\xa8\x3e\xab\x61\xbd\xa8\x3f\x25\x8e\xf9\x9f\x37\x0d\xb0\x2d\x73\x1e\xc7\x6b\xd0\x42\xbe\x1a\xb6\xaa\xdf\x01\x8c\x63\xb2\x6b\xc5\xc5\x12\x68\x85\xe6\x74\x6e\xc6\x5d\x29\x25\xe0\xca\xb7\x6e\xaf\x20\x8c\xca\x72\x84\x0f\x35\xca\x43\xe8\xb3\x4d\x40\x78\xde\xe3\x85\xc0\x4a\xd3\x6f\x6b\x8f\xc3\xc5\xa1\x1e\xf7\xc4\x0a\x72\xcd\xdf\xa5\xbd\xf6\xef\xd3\xb8\x2c\x21\x73\x2d\x8d\xc3\xb7\xdd\xc3\xac\xfe\xa7\xcf\x9c\x9d\xb4\xc3\x86\xbe\xbd\xee\x1e\xfe\xa0\x2b\x00\x74\x25\x37\x8f\x7b\xb8\x3f\x79\xad\x8b\xf8\x3d\x39\x18\xa5\x5a\x18\x33\xa0\x9e\xd8\xb7\x59\x80\x4e\x65\xd8\xd6\xdd\x61\x59\x9b\x75\xac\xd7\x93\x75\xdd\x1c\xaa\x67\x4a\x5e\x34\xd7\x0b\x2e\xf5\xc3\x08\x11\xe9\x71\x2d\x15\xfb\xb7\xe9\x25\xb6\x92\x0b\xf0\x39\x07\xe3\xb8\x37\x46\x2b\x01\xe1\x4a\xb1\x45\x3c\x4c\xa3\x61\xf4\xba\xd4\xe0\xb7\x0a\x63\x61\x5e\x49\xed\xac\xb9\x59\x19\xa8\x6e\x6c\xa5\x34\x0f\xcb\x89\xfc\x4a\x94\xe1\x82\xa9\x50\x98\x6e\x31\xa0\x2e\xf0\xb1\xf8\xfa\x7f\x11\xd5\x96\xf3\x99\xee\x3c\xc2\xef\x69\x0b\xd6\x87\xae\x72\xbd\xd6\x52\x35\x5e\x57\x36\x5a\x6e\x43\xe6\x99\xe8\x78\x0c\x69\xeb\x7a\xca\x4f\x99\xd3\xa1\x7e\xb4\x28\xb0\x16\x52\xb2\x92\x6a\xf0\x42\x88\xba\x47\xe3\xb3\xe3\x2e\x0c\x81\xc7\x11\x7e\x0a\xf1\xf6\xf1\xe9\xf7\xf6\x84\x9f\xdc\x41\xa2\xfc\x38\x1e\xab\xb1\x7b\x78\x79\x81\x9f\xe3\x7d\x1c\xcc\xd5\xff\xf7\x52\xe9\xc4\xbb\x87\x85\x5e\x3d\x78\xdc\xef\x1e\x1e\x5e\x77\xdb\xcb\xcc\x69\x17\xcf\x37\x78\xf7\x02\x0b\xde\xa7\x7b\xec\xa7\x5f\x3e\xef\x1e\x96\x07\x78\xbb\xf2\xee\x87\x3b\x0b\xe0\x6d\x89\x4f\xd5\xb5\x6d\x0d\x6e\xab\xe1\x61\xe4\x0f\xed\x7d\x2c\xfe\x78\xbb\x6f\x6f\xbf\xf4\x77\x8b\xa6\xd6\x16\x66\xec\x4a\xf4\x8d\x4a\xfd\xff\x6c\x57\x12\x07\xb8\xed\x77\xaf\xbb\x7f\x03\x00\x00\xff\xff\x07\xba\x3e\x57\x52\x08\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/js"].(os.FileInfo),
		fs["/nosync"].(os.FileInfo),
	}
	fs["/js"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/js/js.go"].(os.FileInfo),
	}
	fs["/nosync"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/nosync/map.go"].(os.FileInfo),
		fs["/nosync/mutex.go"].(os.FileInfo),
		fs["/nosync/once.go"].(os.FileInfo),
		fs["/nosync/pool.go"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
