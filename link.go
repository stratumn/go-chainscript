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

package chainscript

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/golang/protobuf/proto"
)

// Hash serializes the link using protobuf and computes a SHA256 hash of the
// resulting bytes.
func (l *Link) Hash() ([]byte, error) {
	b, err := proto.Marshal(l)
	if err != nil {
		return nil, err
	}

	lh := sha256.Sum256(b)
	return lh[:], nil
}

// HashString returns the hex-encoded link hash.
func (l *Link) HashString() (string, error) {
	lh, err := l.Hash()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(lh), nil
}

// PrevLinkHash returns the link's parent hash.
// If the link doesn't have a parent, it returns nil.
func (l *Link) PrevLinkHash() []byte {
	if l.Meta == nil {
		return nil
	}

	return l.Meta.PrevLinkHash
}
