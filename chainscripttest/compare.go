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

package chainscripttest

import (
	"testing"

	"github.com/stratumn/go-chainscript"
	"github.com/stretchr/testify/require"
)

// LinksEqual compares two links.
// We can't directly compare the structs because protobuf sets some internal
// state data in the XXX_* fields of each underlying struct when serializing.
func LinksEqual(t *testing.T, l1, l2 *chainscript.Link) {
	lh1, err := l1.Hash()
	require.NoError(t, err)

	lh2, err := l2.Hash()
	require.NoError(t, err)

	require.Equal(t, lh1, lh2)
}

// EvidencesEqual compares two evidences.
// We can't directly compare the structs because protobuf sets some internal
// state data in the XXX_* fields of each underlying struct when serializing.
func EvidencesEqual(t *testing.T, e1, e2 *chainscript.Evidence) {
	require.Equal(t, e1.Version, e2.Version)
	require.Equal(t, e1.Backend, e2.Backend)
	require.Equal(t, e1.Provider, e2.Provider)
	require.Equal(t, e1.Proof, e2.Proof)
}

// SegmentsEqual compares two segments.
// We can't directly compare the structs because protobuf sets some internal
// state data in the XXX_* fields of each underlying struct when serializing.
func SegmentsEqual(t *testing.T, s1, s2 *chainscript.Segment) {
	LinksEqual(t, s1.Link, s2.Link)
	require.Equal(t, s1.LinkHash(), s2.LinkHash())

	require.Equal(t, len(s1.Meta.Evidences), len(s2.Meta.Evidences))
	for i := 0; i < len(s1.Meta.Evidences); i++ {
		EvidencesEqual(t, s1.Meta.Evidences[i], s2.Meta.Evidences[i])
	}
}
