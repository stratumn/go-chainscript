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

package chainscript_test

import (
	"testing"

	"github.com/stratumn/go-chainscript"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLinkBuilder(t *testing.T) {
	process := "test_process"
	mapID := "test_map_1"

	testLink, err := chainscript.NewLinkBuilder(process, mapID).Build()
	require.NoError(t, err)

	testLinkHash, err := testLink.Hash()
	require.NoError(t, err)

	testCases := []struct {
		name     string
		builder  *chainscript.LinkBuilder
		validate func(*testing.T, *chainscript.Link)
		err      error
	}{{
		"missing process",
		chainscript.NewLinkBuilder("", mapID),
		nil,
		chainscript.ErrMissingProcess,
	}, {
		"missing map id",
		chainscript.NewLinkBuilder(process, ""),
		nil,
		chainscript.ErrMissingMapID,
	}, {
		"default values",
		chainscript.NewLinkBuilder(process, mapID),
		func(t *testing.T, l *chainscript.Link) {
			assert.Equal(t, chainscript.LinkVersion, l.Version)
			assert.Equal(t, 0.0, l.Meta.Priority)
			assert.Nil(t, l.Data)
			assert.Nil(t, l.Meta.Data)
		},
		nil,
	}, {
		"client ID",
		chainscript.NewLinkBuilder(process, mapID),
		func(t *testing.T, l *chainscript.Link) {
			assert.Equal(t, chainscript.ClientID, l.Meta.ClientId)
		},
		nil,
	}, {
		"action",
		chainscript.NewLinkBuilder(process, mapID).WithAction("receive-document"),
		func(t *testing.T, l *chainscript.Link) {
			assert.Equal(t, "receive-document", l.Meta.Action)
		},
		nil,
	}, {
		"priority",
		chainscript.NewLinkBuilder(process, mapID).WithPriority(0.42),
		func(t *testing.T, l *chainscript.Link) {
			assert.Equal(t, 0.42, l.Meta.Priority)
		},
		nil,
	}, {
		"negative priority",
		chainscript.NewLinkBuilder(process, mapID).WithPriority(-0.42),
		nil,
		chainscript.ErrInvalidPriority,
	}, {
		"process state",
		chainscript.NewLinkBuilder(process, mapID).WithProcessState("all-documents-gathered"),
		func(t *testing.T, l *chainscript.Link) {
			assert.Equal(t, "all-documents-gathered", l.Meta.Process.State)
		},
		nil,
	}, {
		"step",
		chainscript.NewLinkBuilder(process, mapID).WithStep("document-handoff"),
		func(t *testing.T, l *chainscript.Link) {
			assert.Equal(t, "document-handoff", l.Meta.Step)
		},
		nil,
	}, {
		"tags",
		chainscript.NewLinkBuilder(process, mapID).WithTags("tag1", "tag2").WithTags("", "tag3"),
		func(t *testing.T, l *chainscript.Link) {
			assert.Len(t, l.Meta.Tags, 3)
			assert.ElementsMatch(t, []string{"tag1", "tag2", "tag3"}, l.Meta.Tags)
		},
		nil,
	}, {
		"parent",
		chainscript.NewLinkBuilder(process, mapID).WithParent(testLinkHash),
		func(t *testing.T, l *chainscript.Link) {
			assert.Equal(t, testLinkHash, l.PrevLinkHash())
		},
		nil,
	}, {
		"parent missing",
		chainscript.NewLinkBuilder(process, mapID).WithParent(nil),
		nil,
		chainscript.ErrMissingLinkHash,
	}, {
		"refs",
		chainscript.NewLinkBuilder(process, mapID).WithRefs(
			&chainscript.LinkReference{Process: process, LinkHash: testLinkHash},
			&chainscript.LinkReference{Process: process, LinkHash: testLinkHash},
			&chainscript.LinkReference{Process: process, LinkHash: make([]byte, 32)},
		),
		func(t *testing.T, l *chainscript.Link) {
			expected := []*chainscript.LinkReference{
				&chainscript.LinkReference{Process: process, LinkHash: testLinkHash},
				&chainscript.LinkReference{Process: process, LinkHash: make([]byte, 32)},
			}
			assert.ElementsMatch(t, expected, l.Meta.Refs)
		},
		nil,
	}, {
		"invalid refs",
		chainscript.NewLinkBuilder(process, mapID).WithRefs(
			&chainscript.LinkReference{Process: "", LinkHash: testLinkHash},
			&chainscript.LinkReference{Process: process, LinkHash: testLinkHash},
		),
		nil,
		chainscript.ErrMissingProcess,
	}, {
		"custom data",
		chainscript.NewLinkBuilder(process, mapID).WithData(map[string]string{
			"orderID":     "123",
			"description": "stuff",
		}),
		func(t *testing.T, l *chainscript.Link) {
			assert.NotEmpty(t, l.Data)
		},
		nil,
	}, {
		"custom metadata",
		chainscript.NewLinkBuilder(process, mapID).WithMetadata(map[string]string{
			"orderDate": "Monday",
		}),
		func(t *testing.T, l *chainscript.Link) {
			assert.NotEmpty(t, l.Meta.Data)
		},
		nil,
	}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			l, err := tt.builder.Build()
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
				assert.Nil(t, l)
			} else {
				require.NoError(t, err)
				require.NotNil(t, l)
				tt.validate(t, l)
			}
		})
	}
}
