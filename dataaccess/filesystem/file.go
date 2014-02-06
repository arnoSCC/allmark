// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filesystem

import (
	"fmt"
	"github.com/andreaskoch/allmark2/common/route"
	"github.com/andreaskoch/allmark2/common/util/fsutil"
	"github.com/andreaskoch/allmark2/dataaccess"
	"io/ioutil"
	"path/filepath"
)

func getFiles(repository *Repository, itemDirectory, filesDirectory string) []*dataaccess.File {

	childs := make([]*dataaccess.File, 0)

	filesDirectoryEntries, err := ioutil.ReadDir(filesDirectory)
	if err != nil {
		return childs
	}

	for _, directoryEntry := range filesDirectoryEntries {

		filePath := filepath.Join(filesDirectory, directoryEntry.Name())

		// recurse if the path is a directory
		if isDir, _ := fsutil.IsDirectory(filePath); isDir {
			childs = append(childs, getFiles(repository, itemDirectory, filePath)...)
			continue
		}

		// append new file
		file, err := newFile(repository, itemDirectory, filePath)
		if err != nil {
			fmt.Printf("Unable to add file %q to index.\nError: %s\n", filePath, err)
		}

		childs = append(childs, file)
	}

	return childs
}

func newFile(repository *Repository, itemDirectory, filePath string) (*dataaccess.File, error) {

	// check if the file path is a file
	if isFile, _ := fsutil.IsFile(filePath); !isFile {
		return nil, fmt.Errorf("%q is not a file.", filePath)
	}

	// parent route
	parentRoute, err := route.NewFromFilePath(repository.Path(), itemDirectory)
	if err != nil {
		return nil, fmt.Errorf("Cannot create a parent route for the File with the file path %q. Error: %s", filePath, err)
	}

	// route
	route, err := route.NewFromFilePath(itemDirectory, filePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot create a File for the file path %q. Error: %s", filePath, err)
	}

	// content provider
	contentProvider := newContentProvider(filePath, route)

	// create the file
	file, err := dataaccess.NewFile(route, parentRoute, contentProvider)

	if err != nil {
		return nil, err
	}

	return file, nil
}