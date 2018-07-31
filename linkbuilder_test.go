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
	t.Run("missing process", func(t *testing.T) {
		l, err := chainscript.NewLinkBuilder("", "mapID_123").Build()
		assert.EqualError(t, err, chainscript.ErrMissingProcess.Error())
		assert.Nil(t, l)
	})

	t.Run("missing map ID", func(t *testing.T) {
		l, err := chainscript.NewLinkBuilder("process1", "").Build()
		assert.EqualError(t, err, chainscript.ErrMissingMapID.Error())
		assert.Nil(t, l)
	})

	t.Run("version", func(t *testing.T) {
		l, err := chainscript.NewLinkBuilder("p1", "map1").Build()
		require.NoError(t, err)
		require.NotNil(t, l)
		assert.Equal(t, chainscript.LinkVersion, l.Version)
	})
}
