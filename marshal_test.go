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

	json "github.com/gibson042/canonicaljson-go"
	"github.com/golang/protobuf/proto"
	"github.com/stratumn/go-chainscript"
	"github.com/stratumn/go-chainscript/chainscripttest"
	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	data := map[string]interface{}{
		"user_id":   42,
		"user_name": "spongebob",
	}

	metadata := map[string]interface{}{
		"location": "france",
		"age":      42,
	}

	link, err := chainscript.NewLinkBuilder("proc", "my_map").
		WithAction("init").
		WithPriority(42.0).
		WithRefs(&chainscript.LinkReference{
			LinkHash: []byte{42, 42, 42},
			Process:  "some other process",
		}).
		WithTags("t1", "t2").
		WithData(data).
		WithMetadata(metadata).
		Build()
	require.NoError(t, err)

	key1 := chainscripttest.RandomPrivateKey(t)
	key2 := chainscripttest.RandomPrivateKey(t)

	require.NoError(t, link.Sign(key1, ""))
	require.NoError(t, link.Sign(key2, "[version,state]"))

	segment, err := link.Segmentify()
	require.NoError(t, err)

	err = segment.AddEvidence(&chainscript.Evidence{
		Version:  "0.1.0",
		Backend:  "dummy",
		Provider: "dummy1",
		Proof:    []byte{42, 42},
	})
	require.NoError(t, err)

	require.NoError(t, segment.Validate(context.Background(), nil))

	t.Run("proto", func(t *testing.T) {
		protoBytes, err := proto.Marshal(segment)
		require.NoError(t, err)

		var unmarshalled chainscript.Segment
		err = proto.Unmarshal(protoBytes, &unmarshalled)
		require.NoError(t, err)

		err = unmarshalled.Validate(context.Background(), nil)
		require.NoError(t, err)

		chainscripttest.SegmentsEqual(t, segment, &unmarshalled)
	})

	t.Run("json", func(t *testing.T) {
		jsonBytes, err := json.Marshal(segment)
		require.NoError(t, err)

		var unmarshalled chainscript.Segment
		err = json.Unmarshal(jsonBytes, &unmarshalled)
		require.NoError(t, err)

		err = unmarshalled.Validate(context.Background(), nil)
		require.NoError(t, err)

		chainscripttest.SegmentsEqual(t, segment, &unmarshalled)
	})
}
