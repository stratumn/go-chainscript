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

// Regression tests defined in https://github.com/stratumn/chainscript/tree/master/samples.

package chainscript_test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stratumn/go-chainscript"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegression(t *testing.T) {
	t.Run("v1.0.0", func(t *testing.T) {
		type TestData struct {
			ID   string `json:"id"`
			Data string `json:"data"`
		}

		b, err := ioutil.ReadFile("./proto/samples/1.0.0.json")
		require.NoError(t, err)

		var testData []TestData
		err = json.Unmarshal(b, &testData)
		require.NoError(t, err)

		testSegments := make(map[string]*chainscript.Segment)
		for _, td := range testData {
			data, err := base64.StdEncoding.DecodeString(td.Data)
			require.NoError(t, err)

			var segment chainscript.Segment
			err = proto.Unmarshal(data, &segment)
			require.NoError(t, err)

			err = segment.Validate(context.Background(), nil)
			require.NoError(t, err)

			testSegments[td.ID] = &segment
		}

		t.Run("simple-segment", func(t *testing.T) {
			s, ok := testSegments["simple-segment"]
			require.True(t, ok)

			assert.Equal(t, "1.0.0", s.Link.Version)
			assert.Equal(t, "github.com/stratumn/go-chainscript", s.Link.Meta.ClientId)
			assert.Equal(t, []byte{42, 42}, s.Link.Meta.PrevLinkHash)
			assert.Equal(t, 42.0, s.Link.Meta.Priority)
			assert.Equal(t, "test_process", s.Link.Meta.Process.Name)
			assert.Equal(t, "started", s.Link.Meta.Process.State)
			assert.Equal(t, "test_map", s.Link.Meta.MapId)
			assert.Equal(t, "init", s.Link.Meta.Action)
			assert.Equal(t, "setup", s.Link.Meta.Step)
			assert.ElementsMatch(t, []string{"tag1", "tag2"}, s.Link.Meta.Tags)

			type CustomData struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}

			linkData := CustomData{}
			err = s.Link.StructurizeData(&linkData)
			require.NoError(t, err)
			assert.Equal(t, "batman", linkData.Name)
			assert.Equal(t, 42, linkData.Age)

			linkMetadata := ""
			err = s.Link.StructurizeMetadata(&linkMetadata)
			require.NoError(t, err)
			assert.Equal(t, "bruce wayne", linkMetadata)
		})

		t.Run("segment-references", func(t *testing.T) {
			s, ok := testSegments["segment-references"]
			require.True(t, ok)

			assert.Equal(t, "1.0.0", s.Link.Version)
			assert.Equal(t, "test_process", s.Link.Meta.Process.Name)
			assert.Equal(t, "test_map", s.Link.Meta.MapId)

			refs := s.Link.Meta.Refs
			require.Len(t, refs, 2)

			assert.Equal(t, "p1", refs[0].Process)
			assert.Equal(t, []byte{42}, refs[0].LinkHash)

			assert.Equal(t, "p2", refs[1].Process)
			assert.Equal(t, []byte{24}, refs[1].LinkHash)
		})

		t.Run("segment-evidences", func(t *testing.T) {
			s, ok := testSegments["segment-evidences"]
			require.True(t, ok)

			assert.Equal(t, "1.0.0", s.Link.Version)
			assert.Equal(t, "test_process", s.Link.Meta.Process.Name)
			assert.Equal(t, "test_map", s.Link.Meta.MapId)

			require.Len(t, s.Meta.Evidences, 2)

			btc := s.GetEvidence("bitcoin", "testnet")
			require.NotNil(t, btc)
			assert.Equal(t, "0.1.0", btc.Version)
			assert.Equal(t, []byte{42}, btc.Proof)

			eth := s.GetEvidence("ethereum", "mainnet")
			require.NotNil(t, eth)
			assert.Equal(t, "1.0.3", eth.Version)
			assert.Equal(t, []byte{24}, eth.Proof)
		})

		t.Run("segment-signatures", func(t *testing.T) {
			s, ok := testSegments["segment-signatures"]
			require.True(t, ok)

			assert.Equal(t, "1.0.0", s.Link.Version)
			assert.Equal(t, "test_process", s.Link.Meta.Process.Name)
			assert.Equal(t, "test_map", s.Link.Meta.MapId)

			require.Len(t, s.Link.Signatures, 2)
			assert.NoError(t, s.Link.Signatures[0].Validate(s.Link))
			assert.NoError(t, s.Link.Signatures[1].Validate(s.Link))

			assert.Equal(t, "1.0.0", s.Link.Signatures[0].Version)
			assert.Equal(t, "[version,data,meta]", s.Link.Signatures[0].PayloadPath)

			assert.Equal(t, "1.0.0", s.Link.Signatures[1].Version)
			assert.Equal(t, "[version,meta.mapId]", s.Link.Signatures[1].PayloadPath)
		})
	})
}
