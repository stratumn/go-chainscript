// Copyright 2017-2018 Stratumn SAS. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main defines the end-to-end compatibility tests.
//
// Every implementation of ChainScript needs to generate the same test suite
// to test that encoding/decoding works across all implementations.
// When a new version of ChainScript is released:
//  * this test suite should be updated to cover the new features
//  * snapshot encoded bytes of the previous version should be added to the
//  tests in https://github.com/stratumn/chainscript/samples.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	// TestCases included in the compatibility test suite.
	TestCases = map[string]TestCase{
		"simple-segment":     NewSimpleSegmentTest(),
		"segment-references": NewReferencesTest(),
		"segment-evidences":  NewEvidencesTest(),
	}
)

func main() {
	action := os.Args[1]
	path := os.Args[2]

	switch action {
	case "generate":
		generate(path)
	case "validate":
		validate(path)
	default:
		panic(fmt.Sprintf("Unknown action %s", action))
	}
}

// generate encoded test segments and save them at the specified path.
func generate(path string) {
	var results []TestData
	for id, t := range TestCases {
		results = append(results, TestData{
			ID:   id,
			Data: t.Generate(),
		})
	}

	b, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, b, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

// validate encoded test segments from the specified path.
func validate(path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var testData []TestData
	err = json.Unmarshal(b, &testData)
	if err != nil {
		panic(err)
	}

	for _, t := range testData {
		tt, ok := TestCases[t.ID]
		if !ok {
			fmt.Printf("Unkown test case: %s\n", t.ID)
			continue
		}

		err = tt.Validate(t.Data)
		if err != nil {
			fmt.Printf("[%s] FAILED: %s\n", t.ID, err.Error())
		} else {
			fmt.Printf("[%s] SUCCESS\n", t.ID)
		}
	}
}
