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

package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/stratumn/go-chainscript"
)

// ReferencesTest tests a segment with references.
type ReferencesTest struct{}

// NewReferencesTest creates the test case.
func NewReferencesTest() TestCase {
	return &ReferencesTest{}
}

// Generate encoded segment bytes.
func (t *ReferencesTest) Generate() string {
	link, err := chainscript.NewLinkBuilder("test_process", "test_map").
		WithRefs(
			&chainscript.LinkReference{LinkHash: []byte{42}, Process: "p1"},
			&chainscript.LinkReference{LinkHash: []byte{24}, Process: "p2"},
		).
		Build()
	if err != nil {
		panic(err)
	}

	segment, err := link.Segmentify()
	if err != nil {
		panic(err)
	}

	b, err := proto.Marshal(segment)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(b)
}

// Validate encoded segment bytes.
func (t *ReferencesTest) Validate(encoded string) error {
	b, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}

	var segment chainscript.Segment
	err = proto.Unmarshal(b, &segment)
	if err != nil {
		return err
	}

	err = segment.Validate(context.Background(), nil)
	if err != nil {
		return err
	}

	refs := segment.Link.Meta.Refs

	if len(refs) != 2 {
		return fmt.Errorf("invalid references count: %d", len(refs))
	}

	if refs[0].Process != "p1" {
		return fmt.Errorf("invalid first reference process: %s", refs[0].Process)
	}
	if !bytes.Equal(refs[0].LinkHash, []byte{42}) {
		return fmt.Errorf("invalid first reference link hash: %v", refs[0].LinkHash)
	}

	if refs[1].Process != "p2" {
		return fmt.Errorf("invalid second reference process: %s", refs[1].Process)
	}
	if !bytes.Equal(refs[1].LinkHash, []byte{24}) {
		return fmt.Errorf("invalid second reference link hash: %v", refs[1].LinkHash)
	}

	return nil
}
