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
	"context"
	"testing"

	"github.com/stratumn/go-chainscript"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLink_Hash(t *testing.T) {
	t.Run("unknown version", func(t *testing.T) {
		l, _ := chainscript.NewLinkBuilder("p1", "m1").Build()
		l.Version = "0.42.0"

		lh, err := l.Hash()
		assert.EqualError(t, err, chainscript.ErrUnknownLinkVersion.Error())
		assert.Nil(t, lh)
	})

	t.Run("version 1.0.0", func(t *testing.T) {
		l1, _ := chainscript.NewLinkBuilder("p1", "m1").Build()
		l2, _ := chainscript.NewLinkBuilder("p1", "m1").Build()
		l3, _ := chainscript.NewLinkBuilder("p2", "m42").Build()

		h1, err := l1.Hash()
		require.NoError(t, err)
		assert.Len(t, h1, 32)

		h2, err := l2.Hash()
		require.NoError(t, err)
		assert.Equal(t, h1, h2)

		h3, err := l3.Hash()
		require.NoError(t, err)
		assert.NotEqual(t, h1, h3)
	})
}

func TestLink_HashString(t *testing.T) {
	l, _ := chainscript.NewLinkBuilder("p1", "m1").Build()

	h, err := l.HashString()
	require.NoError(t, err)
	assert.NotEmpty(t, h)
}

func TestLink_PrevLinkHash(t *testing.T) {
	l1, _ := chainscript.NewLinkBuilder("p1", "m1").Build()
	assert.Nil(t, l1.PrevLinkHash())
}

func TestLink_GetTagMap(t *testing.T) {
	t.Run("empty tags", func(t *testing.T) {
		l, _ := chainscript.NewLinkBuilder("p", "m").Build()
		tags := l.GetTagMap()
		assert.Empty(t, tags)
	})

	t.Run("with tags", func(t *testing.T) {
		l, _ := chainscript.NewLinkBuilder("p", "m").WithTags("t1", "t2").Build()
		tags := l.GetTagMap()
		assert.Contains(t, tags, "t1")
		assert.Contains(t, tags, "t2")
		assert.NotContains(t, tags, "t3")
	})
}

func TestLink_Clone(t *testing.T) {
	l, _ := chainscript.NewLinkBuilder("p", "m").Build()

	ll, err := l.Clone()
	require.NoError(t, err)

	assert.Equal(t, l, ll)
	assert.False(t, l == ll)
}

func TestLink_Segmentify(t *testing.T) {
	l, _ := chainscript.NewLinkBuilder("p", "m").Build()
	lh, _ := l.Hash()

	s, err := l.Segmentify()
	require.NoError(t, err)
	assert.Equal(t, l, s.Link)
	assert.Equal(t, lh, s.Meta.LinkHash)
}

func TestLink_Validate(t *testing.T) {
	testCases := []struct {
		name       string
		link       func(*testing.T) *chainscript.Link
		getSegment chainscript.GetSegmentFunc
		err        error
	}{{
		"missing process",
		func(*testing.T) *chainscript.Link {
			l, _ := chainscript.NewLinkBuilder("p", "m").Build()
			l.Meta.Process = nil
			return l
		},
		nil,
		chainscript.ErrMissingProcess,
	}, {
		"missing map ID",
		func(*testing.T) *chainscript.Link {
			l, _ := chainscript.NewLinkBuilder("p", "m").Build()
			l.Meta.MapId = ""
			return l
		},
		nil,
		chainscript.ErrMissingMapID,
	}, {
		"invalid ref",
		func(*testing.T) *chainscript.Link {
			l, _ := chainscript.NewLinkBuilder("p", "m").Build()
			l.Meta.Refs = []*chainscript.LinkReference{
				&chainscript.LinkReference{Process: "p"},
			}
			return l
		},
		nil,
		chainscript.ErrMissingLinkHash,
	}, {
		"missing ref",
		func(*testing.T) *chainscript.Link {
			l, _ := chainscript.NewLinkBuilder("p", "m").Build()
			l.Meta.Refs = []*chainscript.LinkReference{
				&chainscript.LinkReference{Process: "p", LinkHash: make([]byte, 32)},
			}
			return l
		},
		func(context.Context, []byte) (*chainscript.Segment, error) {
			return nil, nil
		},
		chainscript.ErrRefNotFound,
	}, {
		"valid refs",
		func(*testing.T) *chainscript.Link {
			l, _ := chainscript.NewLinkBuilder("p", "m").Build()
			l.Meta.Refs = []*chainscript.LinkReference{
				&chainscript.LinkReference{Process: "p", LinkHash: make([]byte, 32)},
				&chainscript.LinkReference{Process: "p2", LinkHash: make([]byte, 32)},
			}
			return l
		},
		func(context.Context, []byte) (*chainscript.Segment, error) {
			l, _ := chainscript.NewLinkBuilder("p", "m").Build()
			s, _ := l.Segmentify()
			return s, nil
		},
		nil,
	}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.link(t)
			err := l.Validate(context.Background(), tt.getSegment)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
