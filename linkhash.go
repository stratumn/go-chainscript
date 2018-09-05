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
	"encoding/hex"
)

// LinkHash is a byte array for which we provide utility methods.
type LinkHash []byte

// NewLinkHashFromString decodes a string and returns a link hash.
func NewLinkHashFromString(lh string) (LinkHash, error) {
	b, err := hex.DecodeString(lh)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// String encodes the link hash to a string.
func (lh LinkHash) String() string {
	return hex.EncodeToString(lh)
}
