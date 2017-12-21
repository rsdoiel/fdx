//
// fdx is a package encoding/decoding fdx formatted XML files.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// BSD 2-Clause License
//
// Copyright (c) 2017, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package fdx

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

var (
	expectedDocs map[string][]byte
)

func testFdxFile(t *testing.T, fname string) {
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "skipping %s, %s", fname, err)
		return
	}
	fdx := new(FinalDraft)
	if err := xml.Unmarshal(src, &fdx); err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	} else {
		os.RemoveAll(path.Join("testout", path.Base(fname)))
		if src2, err := xml.MarshalIndent(fdx, " ", "    "); err != nil {
			t.Errorf("%s", err)
		} else {
			if err := ioutil.WriteFile(path.Join("testout", path.Base(fname)), src2, 0666); err != nil {
				t.Errorf("%s", err)
			}
		}
	}
}

func TestConversion(t *testing.T) {
	fileList := []string{
		"Big%20Fish.fdx",
		"Brick%20&%20Steel.fdx",
		"The%20Last%20Birthday%20Card.fdx",
		"sample-01.fdx",
		"sample-02.fdx",
		"sample-03.fdx",
	}
	for _, fname := range fileList {
		testFdxFile(t, path.Join("testdata", fname))
	}
}

func TestTitlePageToMap(t *testing.T) {
	// The following shouldn't return populated maps
	noTitlePages := []string{
		"sample-01.fdx",
		"sample-02.fdx",
	}
	screenplay := new(FinalDraft)
	for _, fname := range noTitlePages {
		src, err := ioutil.ReadFile(path.Join("testdata", fname))
		if err != nil {
			t.Errorf("%s", err)
		} else {
			if err := xml.Unmarshal(src, &screenplay); err != nil {
				t.Errorf("Can't Unmarshal %s, %s", fname, err)
			} else {
				m := screenplay.TitlePageAsMap()
				if len(m) > 0 {
					t.Errorf("was expecting an empty map, got %+v\n", m)
				}
			}
		}

	}

	haveTitlePages := map[string][]string{
		"Big%20Fish.fdx":                   []string{"Title", "Credit", "Author", "Source", "Contact"},
		"Brick%20&%20Steel.fdx":            []string{"Title", "Credit", "Author", "Source", "Draft date", "Contact"},
		"The%20Last%20Birthday%20Card.fdx": []string{"Title", "Credit"},
		"sample-03.fdx":                    []string{"Title", "Credit", "Author", "Draft date", "Contact"},
	}
	for fname, fieldnames := range haveTitlePages {
		src, err := ioutil.ReadFile(path.Join("testdata", fname))
		if err != nil {
			if strings.HasPrefix(fname, "sample") {
				t.Errorf("%s", err)
				t.FailNow()
			} else {
				fmt.Printf("Skipping %s", fname)
			}
		} else {
			if err := xml.Unmarshal(src, &screenplay); err != nil {
				t.Errorf("Can't Unmarshal %s, %s", fname, err)
			} else {
				m := screenplay.TitlePageAsMap()
				for _, field := range fieldnames {
					if _, ok := m[field]; ok != true {
						t.Errorf("%s is missing map value for %s", fname, field)
					}
				}
			}
		}
	}
}

func TestMain(m *testing.M) {
	// Setup everything, process flags, etc.
	os.Exit(m.Run())
}
