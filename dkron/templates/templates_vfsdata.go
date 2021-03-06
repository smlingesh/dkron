// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package templates

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

// Templates contains project templates.
var Templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 5, 9, 18, 25, 31, 761990784, time.UTC),
		},
		"/busy.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "busy.html.tmpl",
			modTime:          time.Date(2020, 5, 9, 18, 25, 18, 406827122, time.UTC),
			uncompressedSize: 712,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xd1\x6e\xab\x30\x0c\x7d\xe7\x2b\xac\xbc\x73\x23\xf5\x39\x20\xdd\x7b\xdf\xf6\xb0\x49\xeb\x07\xb4\x81\xb8\x25\x55\x96\xa0\xc4\x74\xab\xa2\xfc\xfb\x44\x28\x85\xaa\xd5\xf6\x02\xf6\xf1\xf1\xc1\x1c\x3b\x46\x85\x07\x6d\x11\x58\xeb\x2c\xa1\x25\x96\x52\x21\x94\x3e\x43\x6b\x64\x08\x55\x86\xa5\xb6\xe8\x61\x26\x80\x3d\x96\x63\xec\x9d\x31\xe8\x2b\xf6\x6f\x08\x97\xff\xe4\x0d\xab\x0b\x80\x75\xab\x77\x9f\x19\x03\x10\xdd\xa6\xfe\xdb\x92\x3e\x23\xe0\x17\xb6\x03\x69\x67\x83\xe0\xdd\xe6\x5a\x26\xd9\x18\x9c\xdb\xa6\x24\x3f\xcb\x40\x5e\xf7\xa8\x66\x19\xf2\x53\x30\x25\x5d\xfd\xe2\x1a\xc1\xa9\xbb\x07\xdf\x07\x6b\xb5\x3d\xc2\xab\x53\xf8\x58\xdd\x92\xf4\x84\x0a\x24\x3d\xd6\xde\x06\xea\x87\x27\xf8\x96\x24\x0d\x61\xc1\x05\x9f\x07\x11\xe4\x47\x3b\x3c\xf6\x28\xa9\x62\x08\xda\xae\xfe\x90\xad\x65\x54\x2d\xe4\xc8\xed\x3c\x1e\x2a\x76\x72\x4d\xe0\x31\xee\x63\x04\xfc\x73\x72\xcd\xce\xca\x0f\x84\x94\xf6\x29\xf1\x95\xc0\xd8\xd0\x68\xab\x2a\xb6\xb0\x58\x2d\xb8\xac\x05\x27\x75\x27\xbf\xa6\x5a\xa7\xf0\xc6\xfd\x81\x17\x26\x33\x76\x92\x7e\x21\xba\xec\xcc\x13\xd2\x92\x00\x08\x3d\xaf\x30\x64\xc3\x4a\x7f\xdd\xc4\xd1\x5c\xfa\x4e\xb7\xce\x2e\x51\xd9\x1b\x79\x29\x5b\xed\x5b\x93\x87\xd4\x2b\xd5\xdb\x37\x56\x3e\xf3\x7c\x0f\xf9\xc2\xb8\xd2\xe7\xba\xb8\xbe\x62\x44\xab\x52\x2a\xbe\x03\x00\x00\xff\xff\x67\xc3\x9a\xeb\xc8\x02\x00\x00"),
		},
		"/dashboard.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "dashboard.html.tmpl",
			modTime:          time.Date(2020, 5, 9, 18, 25, 31, 762454955, time.UTC),
			uncompressedSize: 4283,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\xdf\x6f\xdb\x36\x10\x7e\xf7\x5f\x71\x25\x06\x34\x01\x26\xe9\x61\x6f\x9d\x6c\x20\x4b\x0a\xac\xdb\x9a\x06\x5d\x30\x60\x18\x86\xe0\x24\xd1\x12\x63\x8a\xe4\xc8\x93\x17\x23\xf0\xff\x5e\x50\xb2\x64\xc9\x91\x9b\x58\x4d\x9f\xcc\x9f\xdf\x7d\x77\xf7\x91\x47\x39\x7e\x73\xf5\xe9\xf2\xf6\xef\x9b\xf7\x50\x50\x29\x17\xb3\xd8\xff\x80\xca\x03\x34\x66\xce\xb2\x95\xd5\x8a\x2d\x66\xb3\xb8\xe0\x98\x2d\x66\x00\x31\x09\x92\x7c\xf1\xf8\x08\xe1\xa5\x2e\x4b\xad\xc2\x6b\x2c\x39\x6c\xb7\x71\xd4\xcc\xf8\x35\x2e\xb5\xc2\x10\x38\x9b\xce\x59\x6f\xe5\x85\x73\x9c\xdc\x0d\x52\x01\xdb\x6d\x74\xef\x22\x54\x79\x25\xd1\x86\xa5\x50\xe1\xbd\x63\x8b\x38\x6a\x76\x4e\x01\x09\xac\xae\x88\x8f\x40\x9d\x82\x95\xa1\x2b\x12\x8d\x36\x7b\x4a\x47\x0a\xb5\x02\xcb\xe5\x9c\x39\xda\x48\xee\x0a\xce\x89\x41\x61\xf9\xf2\x2b\x90\xa9\x73\x51\x2e\x37\xa6\x10\xa9\x56\x81\x33\x42\x29\x6e\xc3\xd4\x39\x76\x1a\xaf\xfb\xff\x2a\x6e\x37\x47\x9c\x9b\xcc\x2c\xd1\x9a\x1c\x59\x34\x35\x70\xc3\xea\x04\x52\xc3\xed\xd3\x83\xde\xe1\x04\x4a\x93\x58\x8e\xb9\xf9\x2d\x5e\xa2\x12\x25\xee\x94\xb1\x8f\xfc\x74\x38\x63\x26\x24\x30\xfb\xe9\x15\xbd\xb2\x22\x5d\xb9\x02\xff\xef\x1a\x13\xf2\x37\x8e\x31\xed\x10\x76\x27\xf0\x09\xe6\x6b\x89\x55\x67\xbc\x14\xd6\x6a\x1b\x49\x91\xf4\xba\x27\x3a\x7d\x1c\x67\x9a\xe3\x3d\xbc\x52\x67\x3c\xba\xc7\x35\x36\x1b\x7b\xcd\x6f\xc7\xc6\x2c\xd3\x2a\xca\x84\x33\x12\x37\x11\x56\xa4\x2d\x5f\x5a\xee\x8a\xe9\x97\x66\x25\x82\xa3\xde\x1f\xa4\xa9\xd0\x96\xd2\x8a\xc0\x5f\x61\xcf\x66\x6a\x89\x6b\xbf\x2e\x14\xa9\x66\x40\x1b\xc3\xe7\x4c\x94\x98\xf3\xe8\x21\xa8\xf7\x1f\x88\xfe\x95\x30\xf7\x9e\x7b\x7c\x80\x35\x5a\xb8\xfa\xfd\xf3\xa7\xeb\xbb\x8b\x9b\x0f\x77\x37\x17\xb7\xbf\xc2\x1c\x06\x06\x6e\x3e\xec\xd0\xd9\xcf\xf5\x8e\x1f\xce\x96\x95\x4a\x49\x68\x05\x67\xe7\xf0\x58\x8f\xf9\xd1\xb7\xff\x64\x48\x18\x90\xce\x73\xc9\xe7\x8c\xb4\x96\x24\x0c\xfb\xf7\xed\x79\xb8\x6b\x9f\x9d\xd7\x8b\xb7\xfe\x67\x1f\xc3\x38\x6a\x0a\xe6\x2c\x4e\x74\xb6\xa9\xbd\xce\xc4\x1a\x52\x89\xce\xcd\x99\xc2\x75\x82\x16\x9a\x9f\x20\xe3\x4b\xac\x24\xb5\xdd\xa5\x78\xe0\x59\x40\xda\x30\xb0\xda\x1b\x55\xb8\x16\x39\x7a\x6e\xac\x71\xaf\x0f\x95\x6a\x45\x28\x14\xb7\xbb\x39\x80\xf8\x4d\x10\xc0\xa5\x96\x12\x8d\xe3\x19\xec\x77\x43\x10\x74\x6b\x9e\x90\x09\x3c\xdf\x1e\xca\x0e\xe7\xfd\x83\x41\x95\x71\x0b\x49\x45\x34\x80\x00\x88\x77\x63\x4d\x4a\x9a\x0e\x3b\x40\x6d\x02\xc7\x60\x10\xc5\x74\x47\xae\x1d\x46\x9b\x73\x9a\xb3\x70\xb7\xa7\x9b\xde\x9b\xf2\xf9\x35\xa8\x5a\x70\x67\x03\xad\xe4\x86\x2d\x6e\x6b\xc4\x9e\x8f\x71\xe4\xd7\x1d\xdd\x58\x97\xe1\x04\x6d\x2d\xf7\xef\xbc\x30\x8e\x9a\x90\xd4\xf2\xec\x85\xf4\x23\x0a\x05\xf5\x5b\x69\x18\xcd\x5e\x4a\x12\x8b\x2a\x1b\xba\x8f\xed\x9c\xd4\xb9\x1e\x39\x32\xad\x9c\x17\xb1\x28\xf3\x67\xce\xbf\x28\xf3\xa8\x7e\xd9\x05\x1e\x2c\x34\x2a\x67\x10\x2d\xe2\x08\xfb\xe4\x33\xb1\xee\xd4\xd2\x74\x66\x4f\x75\xf1\x72\x79\xb5\x49\x85\x91\xec\xee\xc3\x32\x8a\x06\x10\x57\xb2\x85\x2b\x51\xa8\xc0\x5f\x20\xce\x2f\x6e\xcf\x8c\xc2\xf5\x30\x5c\x52\xf4\xbb\x75\xfc\x8e\x87\x6c\xb0\x12\x20\x16\xad\xb1\xee\xe9\x06\xfb\x47\x5c\xf7\x44\xf4\x19\x17\x87\x7b\xaf\xda\xd9\xa1\xf5\x7e\x68\x7d\x77\x48\xef\x14\xb6\xd1\xbd\x4e\xdc\x89\x94\x09\xdd\xca\x8d\xd2\xfd\x4d\x27\xee\xbb\x31\x4d\x2a\xb7\x39\x91\xa9\x2f\x70\x41\x2a\x6c\x2a\xf9\x28\xdf\x5f\x2a\xb7\x79\x39\xdf\x38\xaa\xe4\x81\xca\xfe\xe4\x68\xd3\xe2\x47\xb8\xde\x0b\x0d\x55\x06\x9f\xb9\xd1\xd0\xa8\xea\x88\xee\x86\x62\x6b\x9b\x56\xe4\x05\xbd\x50\x79\x05\x91\x79\x17\x35\x07\x2f\x14\xbe\x94\xed\xee\xbd\xbb\x44\xa2\x5a\x9d\x18\xa9\x44\xeb\xd5\xb8\x02\x75\x7a\x42\x4a\xfb\x21\xea\x9d\xf9\xae\xd9\x35\x1e\x1f\x89\x97\x46\x22\x71\xa8\x6b\x0e\x57\xc4\x20\xdc\x6e\xbf\x5a\xd9\x9a\x52\x96\x68\x22\x5d\xb6\x63\x42\xad\xb9\xed\x8e\xff\xc8\x4d\x51\x7f\x23\xb1\xb1\xab\xa4\xab\x75\x40\xfc\x81\x82\x94\x2b\x1a\x56\xac\x41\xba\x9e\x49\x8b\x39\xac\x54\xfc\x81\xd8\xc8\x77\x2d\xf4\x86\xfe\xe2\xd6\x79\xc9\xf8\xaf\x5d\x73\xda\x39\x19\x37\xf7\x91\x97\x09\xb7\xef\xfa\x36\x9a\xa1\xee\xa3\xfa\x55\xcc\xfc\x51\xd7\xf6\x81\x99\x66\x68\xa2\x99\xa3\x86\x6e\x45\xc9\x07\x66\xfc\xc0\x73\x06\x5e\xa8\xc1\x38\x6a\x9e\x52\xfe\x6d\x55\xff\x55\xf1\x25\x00\x00\xff\xff\x30\xe9\x9d\x30\xbb\x10\x00\x00"),
		},
		"/executions.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "executions.html.tmpl",
			modTime:          time.Date(2020, 5, 9, 18, 25, 18, 409826767, time.UTC),
			uncompressedSize: 2633,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x51\x6f\xe3\x36\x0c\x7e\xef\xaf\x60\xb5\xa0\x71\xb0\xda\x46\x0f\x7d\xba\xb3\x0d\xdc\x8a\x6d\x40\x71\xbb\x01\xeb\xb0\xd7\x41\xb6\x18\x4b\x57\x47\x32\x24\xfa\xd2\xc0\xf3\x7f\x1f\x24\xc7\x89\xd3\xa4\x38\xb4\x0f\xae\x44\x52\xe4\xc7\x8f\x14\x95\xbe\x17\xb8\x56\x1a\x81\x55\x46\x13\x6a\x62\xc3\x70\x95\x09\xf5\x1d\xaa\x86\x3b\x97\x07\x31\x57\x1a\x2d\x4c\x06\xa0\xeb\xd8\xaf\xad\x69\x1a\xb4\x39\xfb\xf5\x05\xab\x8e\x94\xd1\xee\x81\x6c\xc3\x8a\x2b\x80\xb9\x03\x6b\xb6\x41\xb6\x97\x2a\x91\xb3\x0d\x3a\xc7\x6b\x64\x45\x96\x0a\xf5\x3d\x1c\x38\x2c\x2e\x9f\x94\x77\xc5\x31\x4c\x96\xca\xbb\x49\xfe\xa1\x78\x92\x66\xab\x74\x0d\x5f\xb8\x23\xe8\xfb\xe4\xc1\x74\x9a\x86\x01\xf0\x60\x0f\x6b\x63\xe1\x9b\x29\xbd\xf6\xd1\x94\x5f\xf9\x06\x87\x21\x9c\x07\xc8\x38\x48\x8b\xeb\x9c\xfd\x34\xe6\xd5\xa8\xea\x39\x67\xb6\xd3\x8f\xa6\x8c\x96\xf3\x03\xcb\x55\xb0\x70\xd2\x6c\x73\x76\x6d\x3b\xad\x95\xae\xff\xfd\x66\x4a\x06\xa4\xa8\xc1\x9c\xfd\xd5\x69\x56\x64\x6a\x82\x5f\x37\xbb\x56\xaa\xca\x68\x38\xac\xe2\xb6\xe1\x3b\xb6\x0f\xed\xff\xb8\x55\x3c\x96\x4a\x08\xd4\x39\x23\xdb\x05\x4e\x54\x91\xa5\xbc\x38\x07\xf8\x03\xdf\x95\xa9\x41\x1b\xbb\xe1\x4d\x6c\x55\x2d\x29\x76\xad\xd2\x1a\xed\x0c\xf6\x1c\xf5\x69\xa0\x2c\x95\x1f\x66\x85\xb8\x02\xe8\x7b\xcb\x75\x8d\xb0\xa8\x9f\xe1\x63\x0e\xc9\x2f\xbb\xdf\xad\xe9\xda\xc0\x5c\x26\xef\x8f\xf5\x80\x20\xff\x08\x7d\xbf\xa8\x9f\x87\x01\xa2\xbe\x6f\x50\x43\xa4\xb4\xc0\x17\x58\x24\x41\xed\xbc\x9f\xd5\x30\xac\xb2\x54\xde\x87\x40\xc4\xcb\x06\xa7\x7c\xc6\x4d\xf8\xc6\x8e\xac\x6a\x51\x4c\xa5\x27\x7b\xa0\x82\x64\xf1\x68\xca\x2c\x25\x39\x17\x3d\x11\xb7\x84\x02\x3e\xd3\x6b\xcd\x6f\x4a\x2b\x27\x2f\xaa\xbe\x1a\x81\xaf\x65\x7f\x76\xd4\x76\x67\x96\x4f\x5d\x55\xa1\x73\x47\x71\x96\x4e\x90\x0e\x14\xe1\xf3\x2d\x2c\xf0\xc5\xf3\x74\x39\xed\xf3\x54\x44\xd1\xf7\x0b\x7c\x39\xf6\x57\x96\x92\x38\x57\xef\x93\xfb\x4c\x6f\x19\x4c\x39\xbe\x6d\xe1\x53\xbd\x1c\xe2\xd0\x88\xc7\x26\xeb\x7b\x98\xa1\x82\x61\x88\xbd\xa4\x7e\x9e\x56\x18\x56\x1b\x23\x78\xc3\x40\x70\xe2\x31\x99\xba\xf6\xdd\x7f\x22\xe3\xb6\x46\x7a\xa7\xbb\x62\x76\x2f\x46\xe4\x63\x41\xe0\x3f\x20\xdb\xe9\x8a\x13\x0e\x43\x92\x24\xf3\xcb\x71\x99\xb3\xb1\x62\xf3\x7c\xc7\x9a\x8d\xcb\xeb\x38\x86\x3f\x7c\x44\x88\xe3\xd9\x60\xda\x77\x62\xc0\x02\x6b\x2e\x90\x85\x59\xf5\x2e\x42\x88\x97\xa1\xfc\x39\x8b\xef\x18\x58\xe3\x79\x11\x8a\x37\xa6\x9e\x2e\x7d\xb8\xf0\x0d\x2f\xb1\x69\x50\x94\xbb\x9c\x6d\x76\x01\xcb\x17\x2f\x3a\x50\x70\x06\x28\xde\x7b\x99\x7c\x9a\xaa\xdb\xf8\x59\x3c\x2b\xe1\xd9\x91\x69\x5e\xcf\x79\x3d\xb7\x92\xc8\x05\xda\x13\x23\x80\xac\xec\x88\x8c\x06\xda\xb5\x98\xb3\x71\xc3\x0e\x2f\x42\x63\x1c\xee\x0b\x2d\x94\xdb\xa8\x83\x33\x36\x4b\x2f\x67\x0f\xc1\xae\xc8\x5c\xcb\xf5\xa5\x41\x77\x43\x6a\x83\xee\x53\x96\x7a\x83\x22\x4b\xc7\x30\xaf\x80\xc8\xfb\x53\xb8\x61\xd4\x8e\x95\x39\x65\xee\x1f\x85\x5b\x30\x63\xc3\xf8\x79\x7f\x76\xbb\xc6\xb1\x73\x70\x3c\x3d\x38\x6f\x33\x53\x1a\xb1\x7b\xcd\x4b\x6b\xb1\x98\xf7\xa6\x77\xec\x65\xef\xf3\xbc\x36\x86\xde\xc7\x79\x49\x1a\x4a\xd2\xb1\xc0\x35\xef\x1a\xba\xcc\x7e\x11\x18\xbf\x44\xe4\x2b\x4c\x27\xdb\xd9\x66\xb6\xec\x7b\xd4\x62\x9c\xf5\x69\x18\xcb\xc5\xd5\x51\xb8\xb7\xcb\x5c\x65\x55\x4b\x5e\xb3\x88\xa6\x96\x5c\x25\x16\xb9\xd8\x45\xeb\x4e\x57\xe1\x75\x88\x56\xd0\x07\x97\x6a\x0d\xd1\x56\x69\x61\xb6\x49\x63\x2a\xee\x95\x89\xe4\x4e\xc2\xcd\x0d\x2c\x2e\x6a\x56\x49\x83\xba\x26\x39\x79\x80\x37\xed\x02\x01\xd1\xd2\xbf\x72\xcb\xd5\xa7\x60\x3c\xce\xdc\x45\xb4\x1c\x95\xcb\x55\x62\xf4\x68\x91\x94\x6e\x2f\xbb\x85\x23\x4c\x3c\x46\xf1\x48\x7d\x19\xcc\x1a\x22\x4c\x2c\x36\x9c\x50\xfc\x1d\x66\xda\x0a\xae\x73\x60\x9d\x1e\x7f\x35\x09\x76\x3c\x04\x70\x31\xb9\x1c\x16\x67\x3e\x12\x4e\x64\xa3\xa5\x1f\xb9\x13\xda\x09\xef\xb0\xdf\xcf\x09\xf5\xb8\xc7\xab\xf3\x43\xe4\x52\x39\x32\x76\x97\xb4\x9d\x93\x4f\xc4\x09\x23\xc6\x6e\x61\x72\x95\x84\xdb\x73\x7b\x06\xb4\xe5\x24\xb5\x9f\x6e\x3f\x9f\xa9\x1c\x72\x5b\xc9\x89\xd2\xf0\xdf\x7f\xb3\x74\xaa\xfd\xd4\x13\xff\x07\x00\x00\xff\xff\xe4\xe4\xea\x43\x49\x0a\x00\x00"),
		},
		"/index.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "index.html.tmpl",
			modTime:          time.Date(2019, 7, 4, 21, 24, 27, 840596557, time.UTC),
			uncompressedSize: 1090,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4d\x8f\xda\x30\x10\xbd\xf3\x2b\x46\xbe\xd0\xaa\x85\x48\xbb\xb7\x95\x13\xa9\xed\xa9\x52\xb5\xaa\xd4\xfe\x01\xc3\x0c\xc4\xc2\xd8\xd1\x78\x60\x8b\xd2\xfc\xf7\xca\x21\x21\xa6\xb0\xab\xbd\xc4\xcf\xef\x3d\x8f\x3d\x1f\x69\x5b\xa4\x8d\xf5\x04\x6a\x1d\xbc\x90\x17\xd5\x75\x33\x8d\xf6\x08\x6b\x67\x62\x2c\x7b\xda\x58\x4f\x0c\xa3\x01\xfc\x76\x91\x30\x07\xe7\x88\x4b\xf5\xdd\x23\xfd\xf9\x26\xec\x54\x35\x03\x68\x5b\xa1\x7d\xe3\x8c\x10\xa8\x28\x46\x0e\x51\xc1\xb2\xeb\x66\x00\x79\x54\x0e\x2f\xbd\xfb\x9a\x6d\x8c\x27\x07\xfd\x77\x81\xb4\x31\x07\x27\x83\xeb\x8e\x6f\x51\x93\x41\xeb\xb7\x17\x07\x80\xae\x1f\xaf\x2d\x62\xc5\x91\xaa\x7e\xd8\x23\xe9\xa2\x7e\xbc\xc4\x2a\xd0\x1e\x5f\x0f\xbc\x0a\x78\xca\xa3\xb2\x5d\xef\x62\x6d\x5e\x60\x04\x8b\xd0\x88\x0d\x3e\x96\x6a\x00\x6a\x92\x36\x64\xe4\xc0\x14\x4b\x35\xa2\x4c\x8c\xc4\x36\x49\xe7\x55\x55\xba\x18\xa5\x3b\x4f\xbb\xc0\x01\xbc\x5e\xc2\xfa\xa1\x7a\x0e\x48\x51\x17\xf5\xc3\x40\x89\x59\x39\x1a\xad\xe7\x4d\xff\x5d\x44\x61\xdb\x10\x4e\x75\x15\xce\x52\x95\xba\x7a\x36\x7b\xd2\x85\xd4\xd7\xec\x17\x44\xa6\x18\x6f\x85\x9f\x81\xe5\x96\xfd\xd5\x77\xfe\x96\xff\x6d\xb6\x57\xac\x2e\xa6\xeb\xb5\x70\x1a\x2d\xa6\x86\x8c\x94\x6a\x4f\xfb\x15\x31\x58\x0f\x67\x14\xe1\x2f\x04\x46\xe2\xaf\xa7\xa7\x79\x7a\xe4\x3c\xef\x91\x60\x3a\xbb\xb2\x1e\xc7\x93\xcb\xe4\x49\x25\x16\x7c\xdb\x97\x52\x7b\x8f\x2f\x65\xfa\x1e\xdf\x39\xf7\x3b\xce\x69\x33\x74\x32\xcb\xf6\xc3\x8e\x4e\x9f\xe1\x68\xdc\x81\x3e\x4e\x39\x2f\x53\xbd\x54\x7e\x0e\x40\xc7\xc6\xf8\xe9\xd2\x1d\x9d\xe0\x13\xcc\x9f\x60\x9e\x6e\x4c\x5a\xf5\x9f\xa3\x0f\x7a\x11\xf3\x37\xe4\xbf\x41\xdf\x0b\xbc\xed\x8b\x2e\xfa\xb9\xc9\xe6\x70\x58\xda\x96\x3c\x76\xdd\xec\x5f\x00\x00\x00\xff\xff\x65\x81\xbc\x4b\x42\x04\x00\x00"),
		},
		"/jobs.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "jobs.html.tmpl",
			modTime:          time.Date(2020, 5, 9, 18, 25, 18, 410926858, time.UTC),
			uncompressedSize: 7733,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\xdd\x8e\xdb\xba\x11\xbe\xdf\xa7\x98\x30\x41\x6c\x37\x91\x7d\x02\x1c\xa0\x80\x23\x39\x38\x48\x1a\xa4\x8b\xdd\x93\xa0\xd9\xf6\xa2\x37\x39\x94\x38\xb6\x98\x52\xa4\x40\x52\xbb\x6b\x78\xfd\x44\x7d\x84\xde\x9d\x27\x2b\x48\xfd\x58\x92\xe5\x9f\xa4\x49\xff\x70\x72\xb1\x96\xc8\xe1\x68\xe6\x9b\x6f\x86\x43\x66\xb3\x61\xb8\xe4\x12\x81\x24\x4a\x5a\x94\x96\x6c\xb7\x17\x21\xe3\xb7\x90\x08\x6a\x4c\xe4\x87\x29\x97\xa8\xa1\x16\x00\xb9\x0a\xdc\xb3\x56\x42\xa0\x8e\xc8\xa5\x8a\xaf\xb8\xb1\xaf\xad\x16\x64\x71\x01\xb0\xd9\x58\xcc\x72\x41\x2d\x02\x31\x96\xda\xc2\x10\x98\x6e\xb7\x17\x00\x6d\xbd\x5a\xdd\x79\x69\x80\x70\xa9\x74\x56\x3e\x76\x45\xdc\x78\xb0\xd2\xaa\xc8\x49\x3d\xdd\x15\xe0\x32\x2f\xec\x9e\xc4\x41\x99\x80\x32\xa6\x24\x59\x84\xbc\x9e\x5c\x89\x75\x9e\xf2\x44\x49\x68\x9e\x02\x83\x54\x27\x29\x59\x84\x33\xbe\x08\x67\x8c\xdf\x76\x34\x7b\x75\xc0\x59\x44\x96\x5c\x58\xd4\x04\xec\x3a\xc7\x88\x58\xbc\xb7\xa4\x63\x78\x05\x11\x81\x5c\xd0\x04\x53\x25\x98\x03\xeb\x6d\xb5\x4a\xae\x82\x4c\x31\x14\x11\x29\xbf\x77\xa9\xe2\x2f\x71\xa1\x25\x09\x10\x52\x48\x35\x2e\x23\xf2\x98\x00\xa3\x96\x06\x56\xad\x56\x02\x23\x92\x29\x46\x45\x3d\x46\xf5\x0a\x6d\x44\x1e\x7f\x56\xf1\x4d\x15\x9f\xeb\x72\xde\x72\xeb\xa4\x7f\xc6\x3b\x87\x4d\x47\xb5\xfb\x77\x04\xab\x5c\xb8\xe0\x52\xcd\x69\x90\x72\xc6\x50\x46\xc4\xea\x02\x6b\xf0\x68\xc7\xa3\x2e\x96\x9d\xd7\xd6\x4b\x38\xdb\xd1\x21\xb4\x34\x16\x58\x1b\x50\xbe\xf8\xbf\x81\xb1\x9a\xe7\xc8\x1a\x20\x42\x9b\x22\x65\x2d\xed\x56\x77\xbe\x6d\xd3\xc5\xa5\x8a\xc3\x99\x4d\xfb\xc3\x1f\x93\x14\x59\x21\x70\x70\xae\x48\x12\x34\x66\x68\xea\x0f\x5a\x2b\x3d\x34\x71\x45\x8d\x05\xbc\xc7\xa4\xb0\x5c\xc9\x21\x89\x9f\xf1\xfe\x84\xc4\x47\x9f\x35\x43\x33\x3f\x25\x6e\x4d\x6f\x2a\x9c\xed\xbc\x75\x33\x2d\x24\x42\xbb\x54\xca\xb6\x71\x61\x90\x28\x61\x72\x2a\x23\xf2\xfb\x83\x8c\xcb\x0b\x21\x02\xcd\x57\xa9\xed\x53\xad\x10\x8d\x0c\x5d\x71\x49\x9d\x39\x3d\x19\x80\x50\x70\x5f\x24\x4a\xc1\x0d\xe3\xc6\x05\x8d\xcd\x21\x29\xb4\x46\x69\x3f\xd0\x15\x42\x14\xc1\x0f\xdb\xbd\xa5\x0d\x9b\x4b\x05\x3c\xf9\x9b\xcb\x34\x6d\xfc\x9a\xf1\x84\x2c\xde\xba\x97\x1e\xb5\x4a\xc7\x05\xff\xce\x76\xe4\x1a\x6f\x6b\x33\x7e\xfd\x3b\x7c\xd0\x78\x7b\xd0\x90\x61\x4b\x34\xe6\x48\x6d\x44\x24\x70\x09\x9a\xca\x15\x8e\x73\xba\x42\xf6\x47\x8b\x99\x99\x0a\x94\x2b\x9b\x3e\x6f\x5b\xd7\x79\x81\x67\xb0\xa2\xf9\x04\xc8\x9e\xad\x2d\x27\x69\x62\xf9\x2d\xce\x41\x3a\xc7\x5a\x8b\xb7\xa4\xe5\x89\xc1\x06\xcf\x63\x7e\xc7\x5c\x32\x67\xec\x33\x78\x41\xfa\xe9\x7c\xca\xd5\x7d\xd0\xc7\x2d\x6b\x26\xce\xba\x3d\xd7\x21\x80\x17\xe7\x85\x42\xe2\x7d\xe3\x81\xcf\xa7\x5f\xff\xf1\xaf\x72\xe2\x5b\x9a\x27\xe8\x8e\xb0\xae\x20\x9c\x65\x5b\x38\x2b\xc4\xd1\x82\x69\x59\x2b\xcb\xdb\x79\x1d\xda\x58\xb1\x75\xa7\xfa\xb5\xd9\xf6\x59\xc5\x8e\x6f\x9f\x55\x6c\xe0\x01\xca\x7d\x6b\xde\x6c\x3b\xf0\x00\xc6\x52\x6d\xdf\x6a\x95\xcd\x3b\x20\xfc\x8e\x3b\xe7\x3f\xa0\xf6\xe4\x7b\x00\xc1\x33\x6e\x6f\xd4\xbc\x3d\xdc\x2d\x21\x96\x2d\x42\xea\x3e\x5d\xee\x47\xee\x8b\xb3\xcd\xe6\x97\xcd\xc6\x7d\x7c\x2a\x69\x86\xb0\xdd\xfe\xb2\xdd\xce\x9a\xf2\x67\xc8\x8e\x68\x4e\x86\x71\x93\x0b\xba\xf6\xa2\x8f\x22\x18\x8d\xe0\x15\xf4\xc7\xe7\x8d\xb6\x92\x95\x6d\x64\x2a\x33\xba\x4a\x4d\x55\xe3\xc9\x19\xa2\x65\xc9\xff\x94\xa8\x42\xda\x33\xe4\xd1\xed\x03\x47\xa4\x3b\x11\x6e\x56\x8e\xdd\x52\x47\x92\x4f\x7e\x3d\x44\x51\x04\xb2\x10\x02\x1e\x1e\xa0\x99\xaa\x4c\x81\x05\xf4\xa4\x5f\xc1\x78\x4f\xa8\xd1\xf0\x0a\x46\x12\x6f\x51\x8f\x2a\x98\xda\x52\x93\xf6\x98\xd7\x35\x81\x07\xd7\x1e\xe0\x7c\xb4\x5e\xaf\xd7\xc1\xf5\x75\xc0\x18\xbc\x7b\x37\xcf\xb2\xb9\x31\xf0\xd7\x51\x37\xbc\xa7\xc0\x70\x59\x79\x52\xe1\x80\x96\x5e\x1e\xb4\x13\xf4\x73\x4d\x50\x5b\x98\xd7\x6e\x88\xec\xf5\x23\x65\xb7\x71\xdc\xce\xbd\x86\xa9\xe1\xe8\xe3\x92\x9f\x4f\xb8\x64\x78\x5f\xb2\x33\xe8\xf4\x4d\xc7\x7a\xa9\x83\x6b\x7b\xc9\x5e\x75\x58\x7f\xe1\x78\xb7\xbf\x59\x1e\xed\x46\x71\x8d\x81\xca\x51\x1e\xec\xb2\x7a\x25\xa4\x57\x67\x9e\xca\xd8\xe4\x2f\x0f\x76\x8b\xbb\x8a\xa5\x0b\x79\xa9\xe2\x71\x9d\x58\x13\x3f\x67\x52\x75\x17\x91\x47\xba\x90\x92\xcb\xd5\xa7\x7d\x6f\x9b\xde\xf1\x4f\x85\xfc\xf2\xde\x91\xae\xcf\xed\x1d\x3b\x56\x9f\xe8\xdf\x13\xb5\x02\xa9\x74\x46\xab\x0e\x26\x30\x39\x97\x12\xf5\xe0\xae\x59\x7a\x78\xc4\xc1\x61\x6b\xce\x86\xb5\x64\x4f\x0f\xd9\x0a\xb3\x1b\x3f\x77\xca\x9d\x9c\x16\x06\xf7\x6d\x3f\x17\xb7\x13\x96\x7e\x3d\xc5\x03\x86\x02\xed\x9e\x65\x3b\xd6\xf8\xf9\x13\xb4\x79\x53\xea\x38\x01\x81\xd5\xd4\xa4\x5f\x0f\x41\xe3\xef\x77\x60\xce\x31\x27\x87\x0f\x41\xed\x5a\xd4\x6b\xdc\x77\x9b\x78\x38\xf3\xc7\x9c\xaa\xb7\x0a\x1f\x05\x01\xf8\xb3\x1a\x94\x80\x41\x10\x54\x72\xae\x67\xff\x77\xed\xf4\x35\x74\x3e\xfc\xb0\xa4\xac\x09\xbe\x3b\x0e\x9f\xe2\x89\x3b\xb8\xf9\xc9\x88\x04\x2f\x08\x68\xe5\xe2\xcf\x38\x15\x6a\x55\x55\x01\x41\x63\x14\x02\x59\xbc\x8e\x48\xb6\xf6\xfe\x5e\xb9\x21\x32\x74\x37\x50\xe9\xae\xd6\x57\xda\x54\x52\x64\x28\xed\x81\xdb\x82\x72\x49\x7d\x8b\x71\xe8\xe8\x53\x4a\xb9\x50\xf4\x2a\xf5\x4f\x1a\x61\xad\x0a\x30\x85\xc6\x57\x87\xdb\xb4\x21\x6d\xae\x57\x43\xdd\x3f\x4a\xc5\x85\xb5\x4a\xd6\x92\xb1\x95\x10\x5b\x19\x30\x5c\xd2\x42\xd8\x76\x11\x29\x11\xec\x15\x91\x5b\x2a\x8a\x5d\x06\xf5\x18\xea\x53\x98\x71\x93\xf1\xc6\x08\xb2\x78\xad\xe4\x92\xeb\x2c\x9c\x95\x1f\x1e\xb6\xa6\xbc\xcd\x28\x5f\xc8\x41\xdb\x86\xf5\x0b\x65\x70\x48\xfb\xd9\x27\x7f\xff\xd8\xa7\xfc\x7f\x0b\xd7\x8f\xb2\xbc\x0e\xc0\xff\x14\xc9\x53\xa4\xec\x20\x2d\x07\x89\x90\xb8\x10\x0f\x87\xbf\xe5\x5e\x44\x3c\x15\xc8\x22\x34\x39\x95\xe7\xd4\xed\xa7\x96\x67\x68\x5e\x86\x33\xb7\x60\x71\x80\xa1\xe9\x8f\x5d\xf3\xfd\x1e\x52\x86\xa5\x83\xe4\xae\x27\x1d\xb9\x86\x0b\x46\xf0\xac\x73\x60\x48\x7f\x3c\x42\xce\x83\xc5\xc0\x69\xe5\x92\x97\xc4\xbb\x34\x4a\x46\x56\xb9\x1f\x97\x93\xfd\x63\x74\x68\xf1\xde\x52\x8d\x14\x0a\x1e\x24\x8a\x61\xc6\x7d\xe3\xde\x79\x0b\x54\x6e\x4d\x44\x90\x71\xab\xf4\xfb\x7c\x77\x18\xaa\xee\x05\xab\xcf\xf8\x86\xb9\x52\xf7\x85\x76\x1f\x2d\x3b\xdf\x23\xd1\xbf\xa4\xa8\x15\xb9\x3b\x26\x54\x45\xcd\x39\xba\xab\x69\xe5\xd4\x59\x35\xed\xcf\x5e\xf4\x1b\x16\x1d\x57\x3a\xea\xeb\xd1\xa1\x02\x74\xa0\x2e\xec\x5f\xaa\xfe\x96\xfb\xff\xe1\xdc\x7f\xad\x91\x5a\x1c\x7d\x7d\xc6\x7f\x9f\x9c\xae\x69\xf2\xff\x9c\xdb\x89\x87\xbe\xca\xed\xb6\xc3\xbb\x1c\x2f\xa3\x73\x5e\xdf\xe2\x45\xbf\x49\x8e\x57\x4f\xf5\x8f\x49\x34\xcf\xfd\x2d\xda\x93\x71\x9d\x5b\x93\xa9\x46\xca\xd6\xe3\x65\x21\xfd\x5d\x3b\x8c\x27\xb0\xf1\x6a\x66\x33\xb8\xe2\xc6\xa2\x84\xa5\xd2\x60\x53\x74\xbb\x0a\x68\x75\x67\xc0\x2a\xa0\x79\x8e\x54\x7b\xc1\x27\xe3\xd1\xd4\xf7\xef\xa3\xc9\xd4\x91\x71\x3c\x7a\xf3\xfe\xfa\x63\x11\x5b\x8d\xae\x3a\xf0\x25\x47\x36\x7a\x0e\xbb\x2f\x60\xfd\x09\x00\xbe\x84\x31\x4e\xcb\xd3\xd7\xd4\x9f\x3e\xde\xdd\x5c\x5f\xd5\x17\x90\x0b\xf8\x61\x27\x5a\x0a\xdf\x71\xc9\xd4\xdd\x54\xa8\xc4\xdf\xc5\x4f\x53\x6a\x52\x78\xfa\x14\x9e\x0c\xce\x4c\x2a\x4d\x6d\x2d\x70\x50\xd6\x07\x60\x3c\x72\x47\x9d\xd1\x64\x77\x88\xdc\x5e\xb4\x7f\xb7\x93\x97\x17\x35\x3e\x37\x29\x37\x60\x52\x55\x08\x06\x31\xba\xf6\xcc\xc1\x64\x55\x0e\x02\x6f\x51\x40\x0d\xb2\x03\xcc\x21\xe3\x7e\x9d\xc4\xb2\xb0\x85\x46\x78\xf3\xfe\x1a\x50\x60\x66\x6a\x7d\x36\xa5\x16\xa8\x46\x90\xca\x02\x15\x3e\x32\x90\x6b\x34\x4e\x87\x92\x40\xfc\x08\x99\x56\xb0\xef\x62\xa8\x64\x69\xf6\x34\x36\xa5\x17\xc7\xf0\x76\xe9\xa2\x3c\xee\x1a\x1d\x53\xd9\x8d\x87\x7f\x02\x8f\x22\x20\x85\x2c\xff\x83\x94\x91\x36\x66\x83\xa8\x47\xf0\x64\x4f\xc7\x94\x5a\xab\xc7\x23\x77\x00\xdd\x41\xd8\x03\xae\x6f\x78\x59\x1f\x4f\x9a\x9e\x72\x63\x95\x5e\x4f\xf3\xc2\xa4\x1f\x2d\xb5\x38\x26\xe4\x79\x03\xf1\xd4\x97\xc8\xe7\x7b\x96\xe6\xd4\xa6\xfe\x16\xf5\xd9\xde\x54\xd9\x3c\x57\x56\x6e\xfd\xaf\xfb\x1b\xce\xea\x3c\xd9\x6c\x50\xb2\xed\xf6\xe2\x9f\x01\x00\x00\xff\xff\x81\x9d\x40\xdd\x35\x1e\x00\x00"),
		},
		"/status.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "status.html.tmpl",
			modTime:          time.Date(2019, 10, 23, 20, 55, 44, 840638000, time.UTC),
			uncompressedSize: 1179,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x94\xbd\xee\xd3\x30\x14\xc5\xf7\x3c\xc5\x95\x19\x0a\x83\x89\xc4\xec\x98\x09\x06\x56\xd8\x2b\xc7\xbe\x49\x0d\xb7\x76\xe5\x8f\x52\x14\xe5\xdd\x51\x12\xa7\x1f\xa2\x40\x05\xea\x7f\xa9\xab\x73\x8e\x6f\x7e\xc7\x8e\x32\x0c\x06\x3b\xeb\x10\x58\x4c\x2a\xe5\xc8\xc6\xb1\x12\xc6\x1e\x41\x93\x8a\xb1\x61\xc1\x7f\x67\xb2\x02\x98\x35\x6b\x1a\xb6\xc7\x18\x55\x8f\x4c\x8a\xda\xd8\xa3\xac\xd6\xe5\x77\x7b\x8a\x76\x50\x0e\x09\xe6\x5f\x6e\xb0\x53\x99\xd2\x9c\xb9\x93\xe2\xad\x37\x3f\x60\xc1\xe1\xb3\x52\x92\x00\x22\xd3\x1a\x25\x1b\x13\xb7\x8e\x26\xf6\x82\xbe\xa6\x00\x04\xd9\x35\x97\xf0\x94\xb8\x46\x97\x30\x80\xf6\xc4\xf7\x86\xbf\xbb\x4a\x02\x88\x98\x82\x77\x3d\xb8\x9e\xb7\xd6\x99\x86\x7d\xf5\x6d\x7c\x4b\xe8\xfa\xb4\x9b\x5a\x2e\xf6\xcd\x8e\x83\x14\x71\xaf\x88\xe4\x17\x9f\x14\xc1\x27\xdf\x46\x51\x2f\xca\x05\xa1\x26\xfb\x5f\x40\x52\x28\xd8\x05\xec\x16\xa0\x57\xef\x3b\x4b\x09\x43\x13\xb3\xd6\x18\x23\x9b\x78\x35\x59\xfd\xad\x61\x8b\xf3\x7a\x53\xac\xcd\x1b\x76\x29\x53\xb4\x2e\xd3\x76\x1a\x33\x15\x52\x7f\x2e\xb5\x72\x96\x0b\x58\x9f\x27\x3f\x9f\x27\xbd\x6c\xe1\x4e\x59\x42\x73\xaf\xef\xe2\xdc\xd4\x5d\xa4\x7f\xac\x6a\x94\xeb\x31\x30\xf9\x71\x1e\xf2\xa4\x96\xbf\x5c\xcd\x56\xfb\xec\xd2\xdf\xde\xb4\xab\xd3\xff\x70\x42\x9d\x93\xf5\xee\xa9\x74\x18\x82\x0f\x8f\xb1\x95\xf3\x7a\x90\x4b\xd4\xb9\x38\xe5\xd3\x71\xfe\x53\x96\x61\x40\x67\xc6\xb1\xfa\x19\x00\x00\xff\xff\xe0\x82\x82\x52\x9b\x04\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/busy.html.tmpl"].(os.FileInfo),
		fs["/dashboard.html.tmpl"].(os.FileInfo),
		fs["/executions.html.tmpl"].(os.FileInfo),
		fs["/index.html.tmpl"].(os.FileInfo),
		fs["/jobs.html.tmpl"].(os.FileInfo),
		fs["/status.html.tmpl"].(os.FileInfo),
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
