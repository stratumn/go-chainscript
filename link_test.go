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

	"github.com/pkg/errors"
	"github.com/stratumn/go-chainscript"
	"github.com/stratumn/go-chainscript/chainscripttest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLink_Data(t *testing.T) {
	t.Run("unknown version", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).WithVersion("0.42.63").Build()

		err := l.SetData("yolo")
		assert.EqualError(t, err, chainscript.ErrUnknownLinkVersion.Error())

		var data string
		err = l.StructurizeData(&data)
		assert.EqualError(t, err, chainscript.ErrUnknownLinkVersion.Error())
	})

	t.Run("unknown client ID", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).WithClientID("github.com/someguy/someapp").Build()

		err := l.SetData("yolo")
		assert.EqualError(t, err, chainscript.ErrUnknownClientID.Error())

		var data string
		err = l.StructurizeData(&data)
		assert.EqualError(t, err, chainscript.ErrUnknownClientID.Error())
	})

	t.Run("js-chainscript interoperability", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).WithClientID("github.com/stratumn/js-chainscript").Build()

		err := l.SetData("yolo")
		require.NoError(t, err)

		var data string
		err = l.StructurizeData(&data)
		require.NoError(t, err)
		assert.Equal(t, "yolo", data)
	})

	t.Run("version 1.0.0", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).Build()
		require.Nil(t, l.Data)

		data := map[string]interface{}{
			"user": "ʙᴀᴛᴍᴀɴ",
			"id":   123,
		}
		err := l.SetData(data)
		require.NoError(t, err)
		require.NotNil(t, l.Data)

		var data2 map[string]interface{}
		err = l.StructurizeData(&data2)
		require.NoError(t, err)
		assert.Len(t, data2, 2)
		assert.Equal(t, "ʙᴀᴛᴍᴀɴ", data2["user"])
		// In version 1.0.0 we use JSON marshalling which only uses float64 for
		// numeric values. So we need an explicit convertion to get an int.
		assert.Equal(t, 123, int(data2["id"].(float64)))
	})
}

func TestLink_Metadata(t *testing.T) {
	t.Run("unknown version", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).WithVersion("0.42.63").Build()

		err := l.SetMetadata("yolo")
		assert.EqualError(t, err, chainscript.ErrUnknownLinkVersion.Error())

		var metadata string
		err = l.StructurizeMetadata(&metadata)
		assert.EqualError(t, err, chainscript.ErrUnknownLinkVersion.Error())
	})

	t.Run("unknown client ID", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).WithClientID("github.com/someguy/someapp").Build()

		err := l.SetMetadata("yolo")
		assert.EqualError(t, err, chainscript.ErrUnknownClientID.Error())

		var metadata string
		err = l.StructurizeMetadata(&metadata)
		assert.EqualError(t, err, chainscript.ErrUnknownClientID.Error())
	})

	t.Run("js-chainscript interoperability", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).WithClientID("github.com/stratumn/js-chainscript").Build()

		err := l.SetMetadata("yolo")
		require.NoError(t, err)

		var metadata string
		err = l.StructurizeMetadata(&metadata)
		require.NoError(t, err)
		assert.Equal(t, "yolo", metadata)
	})

	t.Run("version 1.0.0", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).Build()
		require.Nil(t, l.Meta.Data)

		err := l.SetMetadata("spongebob rocks")
		require.NoError(t, err)
		require.NotNil(t, l.Meta.Data)

		var metadata string
		err = l.StructurizeMetadata(&metadata)
		require.NoError(t, err)
		assert.Equal(t, "spongebob rocks", metadata)
	})
}

func TestLink_Hash(t *testing.T) {
	t.Run("unknown version", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).WithVersion("0.42.63").Build()
		lh, err := l.Hash()
		assert.EqualError(t, err, chainscript.ErrUnknownLinkVersion.Error())
		assert.Nil(t, lh)
	})

	t.Run("version 1.0.0", func(t *testing.T) {
		l1 := chainscripttest.NewLinkBuilder(t).Build()
		l2 := chainscripttest.NewLinkBuilder(t).Build()
		l3 := chainscripttest.NewLinkBuilder(t).WithRandomData().Build()

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
	l := chainscripttest.NewLinkBuilder(t).Build()
	h, err := l.Hash()
	require.NoError(t, err)
	assert.NotEmpty(t, h)

	hs := h.String()
	assert.NotEmpty(t, hs)
}

func TestLink_PrevLinkHash(t *testing.T) {
	l := chainscripttest.NewLinkBuilder(t).Build()
	assert.Nil(t, l.PrevLinkHash())
}

func TestLink_TagMap(t *testing.T) {
	t.Run("empty tags", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).Build()
		tags := l.TagMap()
		assert.Empty(t, tags)
	})

	t.Run("with tags", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).WithTags("t1", "t2").Build()
		tags := l.TagMap()
		assert.Contains(t, tags, "t1")
		assert.Contains(t, tags, "t2")
		assert.NotContains(t, tags, "t3")
	})
}

func TestLink_Clone(t *testing.T) {
	l := chainscripttest.NewLinkBuilder(t).Build()

	ll, err := l.Clone()
	require.NoError(t, err)

	assert.Equal(t, l, ll)
	assert.False(t, l == ll)
}

func TestLink_Segmentify(t *testing.T) {
	l := chainscripttest.NewLinkBuilder(t).Build()
	lh, err := l.Hash()
	require.NoError(t, err)

	s, err := l.Segmentify()
	require.NoError(t, err)
	assert.Equal(t, l, s.Link)
	assert.Equal(t, []byte(lh), s.Meta.LinkHash)
}

func TestLink_Validate(t *testing.T) {
	testCases := []struct {
		name string
		link func(*testing.T) *chainscript.Link
		err  error
	}{{
		"missing version",
		func(*testing.T) *chainscript.Link {
			l := chainscripttest.NewLinkBuilder(t).WithVersion("").Build()
			return l
		},
		chainscript.ErrMissingVersion,
	}, {
		"missing meta",
		func(*testing.T) *chainscript.Link {
			l := chainscripttest.NewLinkBuilder(t).Build()
			l.Meta = nil
			return l
		},
		chainscript.ErrMissingProcess,
	}, {
		"missing process",
		func(*testing.T) *chainscript.Link {
			l := chainscripttest.NewLinkBuilder(t).WithProcess("").Build()
			return l
		},
		chainscript.ErrMissingProcess,
	}, {
		"missing map ID",
		func(*testing.T) *chainscript.Link {
			l := chainscripttest.NewLinkBuilder(t).WithMapID("").Build()
			return l
		},
		chainscript.ErrMissingMapID,
	}, {
		"invalid ref",
		func(*testing.T) *chainscript.Link {
			l := chainscripttest.NewLinkBuilder(t).Build()
			l.Meta.Refs = []*chainscript.LinkReference{
				&chainscript.LinkReference{Process: "p"},
			}
			return l
		},
		chainscript.ErrMissingLinkHash,
	}, {
		"invalid signature",
		func(*testing.T) *chainscript.Link {
			l := chainscripttest.NewLinkBuilder(t).WithSignature(t, "").WithInvalidSignature(t).Build()
			return l
		},
		chainscript.ErrInvalidSignature,
	}, {
		"valid signatures",
		func(*testing.T) *chainscript.Link {
			l := chainscripttest.NewLinkBuilder(t).WithSignature(t, "").WithSignature(t, "").Build()
			return l
		},
		nil,
	}, {
		"valid refs",
		func(*testing.T) *chainscript.Link {
			l := chainscripttest.NewLinkBuilder(t).Build()
			l.Meta.Refs = []*chainscript.LinkReference{
				&chainscript.LinkReference{Process: "p", LinkHash: make([]byte, 32)},
				&chainscript.LinkReference{Process: "p2", LinkHash: make([]byte, 32)},
			}
			return l
		},
		nil,
	}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.link(t)
			err := l.Validate(context.Background())
			if tt.err != nil {
				assert.EqualError(t, errors.Cause(err), tt.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
