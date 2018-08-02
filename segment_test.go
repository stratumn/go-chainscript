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

func TestSegment_Evidence(t *testing.T) {
	createBtcEvidence := func() *chainscript.Evidence {
		return &chainscript.Evidence{
			Version:  "1.0.0",
			Backend:  "bitcoin",
			Provider: "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f",
			Proof:    []byte{42},
		}
	}

	t.Run("add evidence", func(t *testing.T) {
		t.Run("missing version", func(t *testing.T) {
			s, _ := chainscripttest.NewLinkBuilder(t).Build().Segmentify()
			e := createBtcEvidence()
			e.Version = ""
			err := s.AddEvidence(e)
			assert.EqualError(t, err, chainscript.ErrMissingVersion.Error())
		})

		t.Run("missing backend", func(t *testing.T) {
			s, _ := chainscripttest.NewLinkBuilder(t).Build().Segmentify()
			e := createBtcEvidence()
			e.Backend = ""
			err := s.AddEvidence(e)
			assert.EqualError(t, err, chainscript.ErrMissingBackend.Error())
		})

		t.Run("missing provider", func(t *testing.T) {
			s, _ := chainscripttest.NewLinkBuilder(t).Build().Segmentify()
			e := createBtcEvidence()
			e.Provider = ""
			err := s.AddEvidence(e)
			assert.EqualError(t, err, chainscript.ErrMissingProvider.Error())
		})

		t.Run("missing proof", func(t *testing.T) {
			s, _ := chainscripttest.NewLinkBuilder(t).Build().Segmentify()
			e := createBtcEvidence()
			e.Proof = nil
			err := s.AddEvidence(e)
			assert.EqualError(t, err, chainscript.ErrMissingProof.Error())
		})

		t.Run("valid evidence", func(t *testing.T) {
			s, _ := chainscripttest.NewLinkBuilder(t).Build().Segmentify()
			err := s.AddEvidence(createBtcEvidence())
			require.NoError(t, err)
		})

		t.Run("add duplicate evidence", func(t *testing.T) {
			s, _ := chainscripttest.NewLinkBuilder(t).Build().Segmentify()
			e := createBtcEvidence()
			s.AddEvidence(e)

			e2 := createBtcEvidence()
			e2.Proof = []byte{63}
			err := s.AddEvidence(e2)
			assert.EqualError(t, err, chainscript.ErrDuplicateEvidence.Error())

			storedEvidence := s.GetEvidence(e.Backend, e.Provider)
			require.NotNil(t, storedEvidence)
			assert.Equal(t, e.Proof, storedEvidence.Proof)
		})
	})

	t.Run("get evidence", func(t *testing.T) {
		t.Run("valid evidence", func(t *testing.T) {
			s, _ := chainscripttest.NewLinkBuilder(t).Build().Segmentify()
			e := createBtcEvidence()
			s.AddEvidence(e)

			storedEvidence := s.GetEvidence(e.Backend, e.Provider)
			assert.Equal(t, e, storedEvidence)
		})

		t.Run("no result", func(t *testing.T) {
			s, _ := chainscripttest.NewLinkBuilder(t).Build().Segmentify()
			e := s.GetEvidence("santa", "doesn't exist")
			assert.Nil(t, e)
		})
	})

	t.Run("find evidence", func(t *testing.T) {
		t.Run("valid evidences", func(t *testing.T) {
			mainnet := "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"
			e1 := createBtcEvidence()
			e1.Provider = mainnet

			testnet := "000000000933ea01ad0ee984209779baaec3ced90fa3f408719526f8d77f4943"
			e2 := createBtcEvidence()
			e2.Provider = testnet

			s, _ := chainscripttest.NewLinkBuilder(t).Build().Segmentify()
			s.AddEvidence(e1)
			s.AddEvidence(e2)

			e := s.FindEvidences(e1.Backend)
			assert.ElementsMatch(t, []*chainscript.Evidence{e1, e2}, e)
		})

		t.Run("no result", func(t *testing.T) {
			s, _ := chainscripttest.NewLinkBuilder(t).Build().Segmentify()
			s.AddEvidence(createBtcEvidence())

			e := s.FindEvidences("santa")
			assert.Empty(t, e)
		})
	})
}
