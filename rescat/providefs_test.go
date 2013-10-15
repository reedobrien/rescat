// Copyright 2013 Reed O'Brien. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rescat

import (
	"strconv"
	"strings"
	"testing"
)

func Test_ProvidesInterface(t *testing.T) {
	names := []string{"one", "two", "three"}
	pi := &Provide{names, &DummyFetch{}}
	for i, n := range pi.Names {
		if n != names[i] {
			t.Errorf("pi.Names[%i] %s != %s", i, n, names[i])
		}
	}

}

type DummyFetch struct {
}

func (d *DummyFetch) Fetch(n string) (b []byte, err error) {
	return getDummyFile(n)
}

func getDummyFile(n string) (b []byte, err error) {
	if n == "fail" {
		err = FailedFetch
		return b, err
	}
	return []byte(strings.ToUpper(n)), err
}

func TestProvideFetch(t *testing.T) {
	pi := &Provide{[]string{"one", "two", "three"}, &DummyFetch{}}
	for _, n := range pi.Names {
		expected := strings.ToUpper(n)
		actual, err := pi.Fetch(n)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if string(actual) != expected {
			t.Errorf("pi.Fetch(n) failed. Expected '%s', got '%s'", expected, actual)
		}
	}
}

func TestProvideFetchFails(t *testing.T) {
	pi := &Provide{[]string{"fail", "fail", "fail"}, &DummyFetch{}}
	for _, n := range pi.Names {
		actual, err := pi.Fetch(n)
		if err != FailedFetch {
			t.Errorf("Error: %s", err)
		}
		if actual != nil {
			t.Error("Should have failed but didn't")
		}
	}
}

func TestProvidesFS(t *testing.T) {
	pi := &Provide{[]string{"../testfiles/zero.txt", "../testfiles/one.txt", "../testfiles/two.txt"}, &FetchFile{}}
	for i, n := range pi.Names {
		expected := strconv.Itoa(i)
		actual, err := pi.Fetch(n)
		if err != nil {
			t.Errorf("Error in TestProvidesFS: %s", err)
		}
		if string(actual) != expected {
			t.Errorf("Errof in TestProvidesFS: i %d, expected %s, got %s", i, expected, actual)
		}
	}
}

func TestProvidesFSConcat(t *testing.T) {
	pi := &Provide{[]string{"../testfiles/zero.txt", "../testfiles/one.txt", "../testfiles/two.txt"}, &FetchFile{}}
	concat, err := pi.Provide()
	if err != nil {
		t.Errorf("unexpected Error:%s", err)
	}
	expected := "012"
	if string(concat) != expected {
		t.Errorf("Expected %s, got %s", expected, concat)
	}
}