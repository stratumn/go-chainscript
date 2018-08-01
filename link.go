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
	"context"
	"crypto/sha256"
	"encoding/hex"

	json "github.com/gibson042/canonicaljson-go"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

const (
	// LinkVersion1_0_0 is the first version of the link encoding.
	// In that version we encode interfaces (link.data and link.meta.data) with
	// canonical JSON and hash the protobuf-encoded link bytes with SHA-256.
	LinkVersion1_0_0 = "1.0.0"

	// LinkVersion is the version used for new links.
	LinkVersion = LinkVersion1_0_0
)

// Link errors.
var (
	ErrUnknownLinkVersion = errors.New("unknown link version")
	ErrRefNotFound        = errors.New("referenced link could not be retrieved")
)

// GetSegmentFunc is the function signature to retrieve a Segment.
type GetSegmentFunc func(ctx context.Context, linkHash []byte) (*Segment, error)

// Hash serializes the link and computes a hash of the resulting bytes.
// The serialization and hashing algorithm used depend on the link version.
func (l *Link) Hash() ([]byte, error) {
	switch l.Version {
	case LinkVersion1_0_0:
		b, err := proto.Marshal(l)
		if err != nil {
			return nil, err
		}

		lh := sha256.Sum256(b)
		return lh[:], nil
	default:
		return nil, ErrUnknownLinkVersion
	}
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

// GetTagMap returns the tags as a map of string to empty structs (a set).
// It makes it easier to test inclusion.
func (l *Link) GetTagMap() map[string]struct{} {
	tags := make(map[string]struct{})
	for _, v := range l.Meta.Tags {
		tags[v] = struct{}{}
	}
	return tags
}

// Segmentify returns a segment from a link, filling the link hash.
func (l *Link) Segmentify() (*Segment, error) {
	lh, err := l.Hash()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Segment{
		Link: l,
		Meta: &SegmentMeta{
			LinkHash: lh,
		},
	}, nil
}

// Clone returns a copy of the link.
func (l *Link) Clone() (*Link, error) {
	var clone Link
	js, err := json.Marshal(l)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := json.Unmarshal(js, &clone); err != nil {
		return nil, errors.WithStack(err)
	}

	return &clone, nil
}

// Validate checks for errors in a link.
func (l *Link) Validate(ctx context.Context, getSegment GetSegmentFunc) error {
	if l.Meta.Process == nil || len(l.Meta.Process.Name) == 0 {
		return ErrMissingProcess
	}
	if len(l.Meta.MapId) == 0 {
		return ErrMissingMapID
	}

	if _, err := l.Hash(); err != nil {
		return err
	}

	for _, ref := range l.Meta.Refs {
		if len(ref.Process) == 0 {
			return ErrMissingProcess
		}

		if len(ref.LinkHash) == 0 {
			return ErrMissingLinkHash
		}

		// We only check the referenced segment if it's in the same process.
		if l.Meta.Process.Name == ref.Process && getSegment != nil {
			seg, err := getSegment(ctx, ref.LinkHash)
			if err != nil || seg == nil {
				return ErrRefNotFound
			}
		}
	}

	return nil
}
