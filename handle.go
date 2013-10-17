// Copyright 2013 Reed O'Brien. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rescat

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var VERSION string = "2013.10.16"
var MARK string = "??"

// A basic handler.
func HandleFS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", fmt.Sprintf("rescat/%s", VERSION))
	p := &Provide{}
	p.Fetcher = &FetchFile{}
	parseUrl(r.URL, p, w, r)
	fmt.Fprintf(w, "%v", r)
}

func parseUrl(u *url.URL, p *Provide, w http.ResponseWriter, r *http.Request) (err error) {
	v := strings.Split(u.String(), MARK)
	if len(v) != 2 {
		badRequest(w, r)
		return err
	}
	return err
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "400: Bad Request", http.StatusBadRequest)
}
