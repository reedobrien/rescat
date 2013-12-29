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
	pi := &Provide{Fetcher: &DummyFetch{}, Base: "", Names: names, Path: ""}
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
	pi := &Provide{&DummyFetch{}, "", []string{"one", "two", "three"}, ""}
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
	pi := &Provide{&DummyFetch{}, "", []string{"fail", "fail", "fail"}, ""}
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
	names := []string{"testfiles/zero.txt", "testfiles/one.txt", "testfiles/two.txt"}
	pi := &Provide{&FetchFile{}, "", names, ""}
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
	names := []string{"testfiles/zero.txt", "testfiles/one.txt", "testfiles/two.txt"}
	pi := &Provide{&FetchFile{}, "", names, ""}
	concat, err := pi.Provide()
	if err != nil {
		t.Errorf("unexpected Error:%s", err)
	}
	expected := "012"
	if string(concat) != expected {
		t.Errorf("Expected %s, got %s", expected, concat)
	}
}

func TestProvidesFSFailedFetch(t *testing.T) {
	names := []string{"nothere/fails", "testfiles/one.txt", "testfiles/two.txt"}
	pi := &Provide{&FetchFile{}, "", names, ""}
	_, err := pi.Provide()
	if err == nil {
		t.Error("Should have failed but didn't")
	}
	if err != FailedFetch {
		t.Errorf("unexpected Error:%s", err)
	}
}

func BenchmarkProvideFS(b *testing.B) {
	names := []string{"testfiles/zero.txt", "testfiles/one.txt", "testfiles/two.txt"}
	pi := &Provide{&FetchFile{}, "", names, ""}
	for i := 0; i < b.N; i++ {
		_, _ = pi.Provide()
	}
}
