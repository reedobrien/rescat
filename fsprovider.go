// Copyright 2013 Reed O'Brien. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rescat

import (
	"errors"
	"io/ioutil"
	"log"
)

var FailedFetch = errors.New("Failed to fetch resource")

// Implement Fetcher interface for the file system.
type FetchFile struct {
}

// open the file n and return its contents.
func (f *FetchFile) Fetch(n string) (b []byte, err error) {
	return getFileContents(n)
}

// TODO: add timeout

// Provide concatenated files from the file system
type Provide struct {
	// The base for fetching files. I.e. /some/path or http://eample.com/some/path
	Base string
	// A concrete implementation of Fetcher. I.e. FetchFile
	Fetcher
	// a list of the files to be concatinated
	Names []string
	// The path provided in the request. In the case of http://exsmple.com/static/css/
	// the calue  would be static/css
	Path string
}

// TODO: add concurrency to fetch file contents
// TODO: and deterministically call concat in cardinal order

// Provide concatenated output
func (p *Provide) Provide() (b []byte, err error) {
	for _, n := range p.Names {
		f, err := p.Fetch(n)
		if err != nil {
			if err == FailedFetch {
				return nil, err
			} else {
				log.Fatalln("Error in p.Provide", err)
			}
		}
		b, err = concatenate(b, f)
		if err != nil {
			log.Fatal(err)
		}
	}
	return b, err
}

// TODO: pass &c and modify in place

// concatenate two byte slices
func concatenate(c []byte, f []byte) (b []byte, err error) {
	b = append(c, f...)
	return b, err
}

func getFileContents(filename string) (b []byte, err error) {
	b, err = ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("error opening file:", filename, err)
	}
	return b, err
}
