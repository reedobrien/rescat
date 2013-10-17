// Copyright 2013 Reed O'Brien. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rescat

import (
	"io"
)

// Fetcher interface to be provided for FS, HTTP etc...
type Fetcher interface {
	// Fetch the resource n and return it as a byte array.
	Fetch(n string) (b []byte, err error)
}

// Maybe doesn't need to be an interface?
type Provider interface {
	// return the concatinated string as a ??? ready to write to
	// and HTTP Response
	Provide(names []string) (r io.Reader, err error)
	Fetcher
}
