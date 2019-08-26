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
	ErrMissingVersion     = errors.New("version is missing")
	ErrOutDegree          = errors.New("cannot create link: parent has too many children")
	ErrUnknownLinkVersion = errors.New("unknown link version")
	ErrUnknownClientID    = errors.New("link was created with a unknown client: can't deserialize it")
)

// compatibleClients contains a list of other libraries that are compatible
// with this one. Only clients in this list are known to produce binary data
// that this library can correctly interpret.
var compatibleClients = map[string]struct{}{
	ClientID: struct{}{},
	"github.com/stratumn/js-chainscript":     struct{}{},
	"github.com/stratumn/java-chainscript":   struct{}{},
	"github.com/stratumn/csharp-chainscript": struct{}{},
}

// compatible returns an error if the link isn't compatible with this package.
func (l *Link) compatible() error {
	if l.Meta == nil {
		return ErrUnknownClientID
	}

	_, ok := compatibleClients[l.Meta.ClientId]
	if !ok {
		return ErrUnknownClientID
	}

	return nil
}

// SetData uses the given object as link's custom data.
func (l *Link) SetData(data interface{}) error {
	if err := l.compatible(); err != nil {
		return err
	}

	switch l.Version {
	case LinkVersion1_0_0:
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return errors.WithStack(err)
		}

		l.Data = dataBytes
	default:
		return ErrUnknownLinkVersion
	}

	return nil
}

// StructurizeData deserializes the link's data into the given object.
// The provided argument should be a pointer to a struct.
func (l *Link) StructurizeData(data interface{}) error {
	if err := l.compatible(); err != nil {
		return err
	}

	switch l.Version {
	case LinkVersion1_0_0:
		return json.Unmarshal(l.Data, data)
	default:
		return ErrUnknownLinkVersion
	}
}

// SetMetadata uses the given object as link's custom metadata.
func (l *Link) SetMetadata(metadata interface{}) error {
	if err := l.compatible(); err != nil {
		return err
	}

	switch l.Version {
	case LinkVersion1_0_0:
		metadataBytes, err := json.Marshal(metadata)
		if err != nil {
			return errors.WithStack(err)
		}

		l.Meta.Data = metadataBytes
	default:
		return ErrUnknownLinkVersion
	}

	return nil
}

// StructurizeMetadata deserializes the link's metadata into the given object.
// The provided argument should be a pointer to a struct.
func (l *Link) StructurizeMetadata(metadata interface{}) error {
	if err := l.compatible(); err != nil {
		return err
	}

	switch l.Version {
	case LinkVersion1_0_0:
		return json.Unmarshal(l.Meta.Data, metadata)
	default:
		return ErrUnknownLinkVersion
	}
}

// Hash serializes the link and computes a hash of the resulting bytes.
// The serialization and hashing algorithm used depend on the link version.
func (l *Link) Hash() (LinkHash, error) {
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

// PrevLinkHash returns the link's parent hash.
// If the link doesn't have a parent, it returns nil.
func (l *Link) PrevLinkHash() LinkHash {
	if l.Meta == nil {
		return nil
	}

	return l.Meta.PrevLinkHash
}

// TagMap returns the tags as a map of string to empty structs (a set).
// It makes it easier to test inclusion.
func (l *Link) TagMap() map[string]struct{} {
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
func (l *Link) Validate(ctx context.Context) error {
	if len(l.Version) == 0 {
		return ErrMissingVersion
	}
	if l.Meta == nil || l.Meta.Process == nil || len(l.Meta.Process.Name) == 0 {
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
	}

	for _, sig := range l.Signatures {
		if err := sig.Validate(l); err != nil {
			return err
		}
	}

	return nil
}
