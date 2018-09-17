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

	"github.com/pkg/errors"
	"github.com/stratumn/go-chainscript"
)

// SimpleSegmentTest tests a segment with custom data and metadata but no
// references, evidences or signatures.
type SimpleSegmentTest struct{}

// NewSimpleSegmentTest creates the test case.
func NewSimpleSegmentTest() TestCase {
	return &SimpleSegmentTest{}
}

// CustomData link data.
type CustomData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// String representation of custom link data.
func (c *CustomData) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", c.Name, c.Age)
}

// Generate encoded segment bytes.
func (t *SimpleSegmentTest) Generate() string {
	link, err := chainscript.NewLinkBuilder("test_process", "test_map").
		WithAction("init").
		WithData(CustomData{Name: "ʙᴀᴛᴍᴀɴ", Age: 42}).
		WithDegree(3).
		WithMetadata("bruce wayne").
		WithParent([]byte{42, 42}).
		WithPriority(42).
		WithProcessState("started").
		WithStep("setup").
		WithTags("tag1", "tag2").
		Build()
	if err != nil {
		panic(err)
	}

	segment, err := link.Segmentify()
	if err != nil {
		panic(err)
	}

	b, err := chainscript.MarshalSegment(segment)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(b)
}

// Validate encoded segment bytes.
func (t *SimpleSegmentTest) Validate(encoded string) error {
	b, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}

	segment, err := chainscript.UnmarshalSegment(b)
	if err != nil {
		return err
	}

	err = segment.Validate(context.Background())
	if err != nil {
		return err
	}

	if segment.Link.Meta.Action != "init" {
		return errors.Errorf("invalid action: %s", segment.Link.Meta.Action)
	}

	if segment.Link.Meta.OutDegree != int32(3) {
		return errors.Errorf("invalid degree: %d", segment.Link.Meta.OutDegree)
	}

	data := CustomData{}
	err = segment.Link.StructurizeData(&data)
	if err != nil {
		return err
	}

	if data.Age != 42 {
		return errors.Errorf("invalid data: %s", data.String())
	}
	if data.Name != "ʙᴀᴛᴍᴀɴ" {
		return errors.Errorf("invalid data: %s", data.String())
	}

	if segment.Link.Meta.MapId != "test_map" {
		return errors.Errorf("invalid map id: %s", segment.Link.Meta.MapId)
	}

	metadata := ""
	err = segment.Link.StructurizeMetadata(&metadata)
	if err != nil {
		return err
	}

	if metadata != "bruce wayne" {
		return errors.Errorf("invalid metadata: %s", metadata)
	}

	if !bytes.Equal(segment.Link.PrevLinkHash(), []byte{42, 42}) {
		return errors.Errorf("invalid parent: %v", segment.Link.PrevLinkHash())
	}

	if segment.Link.Meta.Priority != 42 {
		return errors.Errorf("invalid priority: %f", segment.Link.Meta.Priority)
	}

	if segment.Link.Meta.Process.Name != "test_process" {
		return errors.Errorf("invalid process name: %s", segment.Link.Meta.Process.Name)
	}

	if segment.Link.Meta.Process.State != "started" {
		return errors.Errorf("invalid process state: %s", segment.Link.Meta.Process.State)
	}

	if segment.Link.Meta.Step != "setup" {
		return errors.Errorf("invalid step: %s", segment.Link.Meta.Step)
	}

	if segment.Link.Meta.Tags[0] != "tag1" || segment.Link.Meta.Tags[1] != "tag2" {
		return errors.Errorf("invalid tags: %v", segment.Link.Meta.Tags)
	}

	return nil
}
