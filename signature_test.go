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
	"github.com/stratumn/go-chainscript/chainscripttest"
	"github.com/stratumn/go-crypto/keys"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// GeneratePrivateKey generates a private key that can be used to sign links.
func GeneratePrivateKey(t *testing.T) []byte {
	_, privKey, err := keys.NewEd25519KeyPair()
	require.NoError(t, err)

	keyBytes, err := keys.EncodeED25519SecretKey(privKey)
	require.NoError(t, err)

	return keyBytes
}

func TestLink_Sign(t *testing.T) {
	t.Run("invalid private key", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).Build()
		err := l.Sign([]byte("this is not the droid you're looking for"), "")
		assert.Error(t, err)
		assert.Empty(t, l.Signatures)
	})

	t.Run("invalid payload path", func(t *testing.T) {
		sk := GeneratePrivateKey(t)
		l := chainscripttest.NewLinkBuilder(t).Build()

		err := l.Sign(sk, "[version")
		assert.Error(t, err)
		assert.Empty(t, l.Signatures)
	})

	t.Run("valid signatures", func(t *testing.T) {
		sk1 := GeneratePrivateKey(t)
		sk2 := GeneratePrivateKey(t)
		l := chainscripttest.NewLinkBuilder(t).Build()

		payloadPaths := []string{"", "[data]"}
		err := l.Sign(sk1, payloadPaths[0])
		require.NoError(t, err)

		err = l.Sign(sk2, payloadPaths[1])
		require.NoError(t, err)

		assert.Len(t, l.Signatures, 2)
		for i, s := range l.Signatures {
			assert.Equal(t, chainscript.SignatureVersion, s.Version)
			assert.Equal(t, payloadPaths[i], s.PayloadPath)
			assert.Equal(t, "ED25519", s.Type)
			assert.Len(t, s.PublicKey, 129)
			assert.Len(t, s.Signature, 136)
		}
	})
}

func TestLink_SignedBytes(t *testing.T) {
	t.Run("unknown version", func(t *testing.T) {
		l := chainscripttest.NewLinkBuilder(t).Build()
		_, err := l.SignedBytes("0.1.0", "")
		assert.EqualError(t, err, chainscript.ErrUnknownSignatureVersion.Error())
	})

	t.Run("version 1.0.0", func(t *testing.T) {
		v1 := chainscript.SignatureVersion1_0_0

		t.Run("include data and meta if no path provided", func(t *testing.T) {
			l := chainscripttest.NewLinkBuilder(t).WithData("b4tm4n").Build()

			b1, err := l.SignedBytes(v1, "[version,data,meta]")
			require.NoError(t, err)

			b2, err := l.SignedBytes(v1, "")
			require.NoError(t, err)

			assert.Equal(t, b1, b2)
			assert.Len(t, b1, 32)

			b3, err := l.SignedBytes(v1, "[version,data]")
			require.NoError(t, err)

			assert.NotEqual(t, b1, b3)
		})

		t.Run("doesn't include signatures by default", func(t *testing.T) {
			l := chainscripttest.NewLinkBuilder(t).Build()

			b1, err := l.SignedBytes(v1, "")
			require.NoError(t, err)

			err = l.Sign(GeneratePrivateKey(t), "")
			require.NoError(t, err)
			require.Len(t, l.Signatures, 1)

			b2, err := l.SignedBytes(v1, "")
			require.NoError(t, err)
			assert.Equal(t, b1, b2)

			b3, err := l.SignedBytes(v1, "[version,data,meta,signatures]")
			require.NoError(t, err)
			assert.NotEqual(t, b1, b3)
		})

		t.Run("partial meta", func(t *testing.T) {
			l := chainscripttest.NewLinkBuilder(t).Build()

			b1, err := l.SignedBytes(v1, "[meta.action,meta.process.name,meta.map_id]")
			require.NoError(t, err)

			b2, err := l.SignedBytes(v1, "[meta.action,meta.process.name]")
			require.NoError(t, err)

			assert.NotEqual(t, b1, b2)
		})

		t.Run("partial meta and link data", func(t *testing.T) {
			l1 := chainscripttest.NewLinkBuilder(t).WithData(map[string]int{
				"user":   42,
				"random": 63,
			}).Build()

			path := "[data,meta.process.name,meta.map_id]"
			b1, err := l1.SignedBytes(v1, path)
			require.NoError(t, err)

			l2 := chainscripttest.NewLinkBuilder(t).Build()
			b2, err := l2.SignedBytes(v1, path)
			require.NoError(t, err)

			assert.NotEqual(t, b1, b2)
		})
	})
}
