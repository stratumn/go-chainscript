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
	"github.com/stratumn/go-chainscript/chainscripttest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSegment_LinkHash(t *testing.T) {
	l := chainscripttest.NewLinkBuilder(t).Build()
	s := &chainscript.Segment{Link: l}

	err := s.SetLinkHash()
	require.NoError(t, err)

	lh := s.LinkHashString()
	assert.NotEmpty(t, lh)

	err = s.SetLinkHash()
	require.NoError(t, err)
	lhh := s.LinkHashString()
	assert.Equal(t, lh, lhh)
}

func TestSegment_Validate(t *testing.T) {
	t.Run("missing link hash", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).Build()
		s := &chainscript.Segment{Link: l}

		err := s.Validate(context.Background(), nil)
		assert.EqualError(t, err, chainscript.ErrMissingLinkHash.Error())
	})

	t.Run("link hash mismatch", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).Build()
		s := &chainscript.Segment{
			Link: l,
			Meta: &chainscript.SegmentMeta{
				LinkHash: []byte{42, 42, 42},
			},
		}

		err := s.Validate(context.Background(), nil)
		assert.EqualError(t, err, chainscript.ErrLinkHashMismatch.Error())
	})

	t.Run("invalid link", func(t *testing.T) {
		s, err := chainscripttest.NewLinkBuilder(t).WithInvalidFields().Build().Segmentify()
		require.NoError(t, err)

		err = s.Validate(context.Background(), nil)
		assert.Error(t, err)
	})

	t.Run("valid segment", func(t *testing.T) {
		s, err := chainscripttest.NewLinkBuilder(t).Build().Segmentify()
		require.NoError(t, err)

		err = s.Validate(context.Background(), nil)
		require.NoError(t, err)
	})
}
