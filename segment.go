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
	ErrLinkHashMismatch  = errors.New("link hash from meta doesn't equal hashed link")
	ErrMissingBackend    = errors.New("evidence backend is missing")
	ErrMissingProvider   = errors.New("evidence provider is missing")
	ErrMissingProof      = errors.New("evidence proof is missing")
	ErrDuplicateEvidence = errors.New("evidence already exists for the given backend and provider")
)

// LinkHash returns the link hash.
func (s *Segment) LinkHash() []byte {
	return s.Meta.LinkHash
}

// LinkHashString returns the hex-encoded link hash.
func (s *Segment) LinkHashString() string {
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

// AddEvidence adds an evidence to the segment.
func (s *Segment) AddEvidence(evidence *Evidence) error {
	err := evidence.Validate()
	if err != nil {
		return err
	}

	if e := s.GetEvidence(evidence.Backend, evidence.Provider); e != nil {
		return ErrDuplicateEvidence
	}

	if s.Meta == nil {
		s.Meta = &SegmentMeta{}
	}

	s.Meta.Evidences = append(s.Meta.Evidences, evidence)
	return nil
}

// GetEvidence gets an evidence from a provider in a given backend.
func (s *Segment) GetEvidence(backend, provider string) *Evidence {
	if s.Meta == nil {
		return nil
	}

	for _, e := range s.Meta.Evidences {
		if e.Backend == backend && e.Provider == provider {
			return e
		}
	}

	return nil
}

// FindEvidences finds all evidences from a specific backend.
func (s *Segment) FindEvidences(backend string) []*Evidence {
	if s.Meta == nil {
		return nil
	}

	var results []*Evidence
	for _, e := range s.Meta.Evidences {
		if e.Backend == backend {
			results = append(results, e)
		}
	}

	return results
}
