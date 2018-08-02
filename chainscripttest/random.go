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
	"math/rand"
	"testing"

	"github.com/stratumn/go-crypto/keys"
	"github.com/stretchr/testify/require"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomBytes returns a random byte array of the specified length.
func RandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

// RandomHash creates a random hash.
func RandomHash() []byte {
	return RandomBytes(32)
}

// RandomString generates a random string.
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandomPrivateKey generates a private key that can be used to sign links.
func RandomPrivateKey(t *testing.T) []byte {
	_, privKey, err := keys.NewEd25519KeyPair()
	require.NoError(t, err)

	keyBytes, err := keys.EncodeED25519SecretKey(privKey)
	require.NoError(t, err)

	return keyBytes
}
