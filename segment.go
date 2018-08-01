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
	"bytes"
	"context"
	"encoding/hex"

	"github.com/pkg/errors"
)

// Segment errors.
var (
	ErrLinkHashMismatch = errors.New("link hash from meta doesn't equal hashed link")
)

// GetLinkHash returns the link hash.
func (s *Segment) GetLinkHash() []byte {
	return s.Meta.LinkHash
}

// GetLinkHashString returns the hex-encoded link hash.
func (s *Segment) GetLinkHashString() string {
	return hex.EncodeToString(s.Meta.LinkHash)
}

// SetLinkHash computes and sets the link hash.
func (s *Segment) SetLinkHash() error {
	linkHash, err := s.Link.Hash()
	if err != nil {
		return err
	}

	if s.Meta == nil {
		s.Meta = &SegmentMeta{}
	}

	s.Meta.LinkHash = linkHash
	return nil
}

// Validate checks for errors in a segment
func (s *Segment) Validate(ctx context.Context, getSegment GetSegmentFunc) error {
	if s.Meta == nil || len(s.Meta.LinkHash) == 0 {
		return ErrMissingLinkHash
	}

	linkHash, err := s.Link.Hash()
	if err != nil {
		return err
	}

	if !bytes.Equal(linkHash, s.Meta.LinkHash) {
		return ErrLinkHashMismatch
	}

	return s.Link.Validate(ctx, getSegment)
}
