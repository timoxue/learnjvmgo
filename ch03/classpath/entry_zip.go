package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

//ZipEntry for zip file path
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (zipEntry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(zipEntry.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, zipEntry, err
		}

	}
	return nil, nil, errors.New("class not found: " + className)
}

func (zipEntry *ZipEntry) String() string {
	return zipEntry.absPath
}
