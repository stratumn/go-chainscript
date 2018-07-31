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

func TestLink_Hash(t *testing.T) {
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