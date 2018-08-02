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

func TestEvidence_New(t *testing.T) {
	testCases := []struct {
		name     string
		version  string
		backend  string
		provider string
		proof    []byte
		err      error
	}{{
		"missing version",
		"",
		"bitcoin",
		"0x42",
		[]byte{42},
		chainscript.ErrMissingVersion,
	}, {
		"missing backend",
		"0.1.0",
		"",
		"0x42",
		[]byte{42},
		chainscript.ErrMissingBackend,
	}, {
		"missing provider",
		"0.1.1",
		"bitcoin",
		"",
		[]byte{42},
		chainscript.ErrMissingProvider,
	}, {
		"missing proof",
		"0.1.1",
		"bitcoin",
		"0x42",
		nil,
		chainscript.ErrMissingProof,
	}, {
		"valid evidence",
		"0.1.1",
		"bitcoin",
		"0x42",
		[]byte{42},
		nil,
	}}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			e, err := chainscript.NewEvidence(tt.version, tt.backend, tt.provider, tt.proof)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
				assert.NotNil(t, e)
				assert.Equal(t, tt.version, e.Version)
				assert.Equal(t, tt.backend, e.Backend)
				assert.Equal(t, tt.provider, e.Provider)
				assert.Equal(t, tt.proof, e.Proof)
				assert.NoError(t, e.Validate())
			}
		})
	}
}
