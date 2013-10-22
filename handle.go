// Copyright 2013 Reed O'Brien. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rescat

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var VERSION string = "2013.10.16"
var MARK string = "??"

// TODO: This should 'implement' provider and fetcher
// TODO: so we don't have to pass the w and r around
type HandleFS struct {
	p *Provide
}

// A basic handler.
func (h *HandleFS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", fmt.Sprintf("rescat/%s", VERSION))
	p := &Provide{Fetcher: &FetchFile{}}
	// p.Fetcher = &FetchFile{}
	err := h.parseUrl(w, r, p)
	if err != nil {
		log.Println("error rescat.HandleFS.ServeHTTP: ", err)
		return
	}
	x, err := p.Provide()
	// TODO: refactor with HandleFS as Provider and Fetcher
	if err != nil {
		if err == errors.New("not found") {
			http.NotFound(w, r)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
	}
	fmt.Fprintf(w, "%s", x)
}

func (h *HandleFS) parseUrl(w http.ResponseWriter, r *http.Request, p *Provide) error {
	v := strings.Split(r.URL.RequestURI(), MARK)
	if len(v) != 2 {
		badRequest(w, r)
		return errors.New("bad request")
	} else {
		// strip the leading and training / we'll join on them else
		p.Path = strings.Trim(v[0], "/")
		p.Names = strings.Split(v[1], ",")
	}
	return nil
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "400: Bad Request", http.StatusBadRequest)
}
