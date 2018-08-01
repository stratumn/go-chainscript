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

	json "github.com/gibson042/canonicaljson-go"
	"github.com/pkg/errors"
)

// Link errors.
var (
	ErrMissingProcess  = errors.New("link process is missing")
	ErrMissingMapID    = errors.New("link map id is missing")
	ErrMissingLinkHash = errors.New("link hash is missing")
	ErrInvalidPriority = errors.New("priority needs to be positive")
)

// LinkBuilder makes it easy to create links that adhere to the ChainScript
// spec.
// It provides valid default values for required fields and allows the user
// to set fields to valid values.
// Note that link builders are not thread safe. They are meant to build an
// object instance which is generally done in a single go routine.
type LinkBuilder struct {
	link *Link
	refs map[string]struct{}
	err  error
}

// NewLinkBuilder creates a new link builder.
func NewLinkBuilder(process string, mapID string) *LinkBuilder {
	var err error
	if len(process) == 0 {
		err = ErrMissingProcess
	}

	if len(mapID) == 0 {
		err = ErrMissingMapID
	}

	return &LinkBuilder{
		link: &Link{
			Version: LinkVersion,
			Meta: &LinkMeta{
				ClientId: ClientID,
				Process: &Process{
					Name: process,
				},
				MapId: mapID,
			},
		},
		err: err,
	}
}

// WithAction sets the link's action.
// The action is what caused the link to be created.
func (b *LinkBuilder) WithAction(action string) *LinkBuilder {
	b.link.Meta.Action = action
	return b
}

// WithData uses the given object as link's custom data.
func (b *LinkBuilder) WithData(data interface{}) *LinkBuilder {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		b.err = errors.WithStack(err)
		return b
	}

	b.link.Data = dataBytes
	return b
}

// WithMetadata uses the given object as link's custom metadata.
func (b *LinkBuilder) WithMetadata(data interface{}) *LinkBuilder {
	metadataBytes, err := json.Marshal(data)
	if err != nil {
		b.err = errors.WithStack(err)
		return b
	}

	b.link.Meta.Data = metadataBytes
	return b
}

// WithParent sets the link's parent, referenced by its hash.
func (b *LinkBuilder) WithParent(linkHash []byte) *LinkBuilder {
	if len(linkHash) == 0 {
		b.err = ErrMissingLinkHash
		return b
	}

	b.link.Meta.PrevLinkHash = linkHash
	return b
}

// WithPriority sets the link's priority.
func (b *LinkBuilder) WithPriority(priority float64) *LinkBuilder {
	if priority <= 0 {
		b.err = ErrInvalidPriority
		return b
	}

	b.link.Meta.Priority = priority
	return b
}

// WithProcessState sets the state of the process.
// If your process can be represented as a state machine and the current link
// changes the state machine, it allows easy tracking of the process evolution.
func (b *LinkBuilder) WithProcessState(state string) *LinkBuilder {
	b.link.Meta.Process.State = state
	return b
}

// WithRefs references links that are related to the current link.
func (b *LinkBuilder) WithRefs(refs ...*LinkReference) *LinkBuilder {
	if b.refs == nil {
		b.refs = make(map[string]struct{})
	}

	for _, ref := range refs {
		if len(ref.Process) == 0 {
			b.err = ErrMissingProcess
			return b
		}

		if len(ref.LinkHash) == 0 {
			b.err = ErrMissingLinkHash
			return b
		}

		hexLinkHash := hex.EncodeToString(ref.LinkHash)
		if _, ok := b.refs[hexLinkHash]; ok {
			continue
		}

		b.link.Meta.Refs = append(b.link.Meta.Refs, ref)
		b.refs[hexLinkHash] = struct{}{}
	}

	return b
}

// WithStep sets the specific process step represented by the link.
func (b *LinkBuilder) WithStep(step string) *LinkBuilder {
	b.link.Meta.Step = step
	return b
}

// WithTags adds some tags to the link.
func (b *LinkBuilder) WithTags(tags ...string) *LinkBuilder {
	for _, tag := range tags {
		if len(tag) == 0 {
			continue
		}

		b.link.Meta.Tags = append(b.link.Meta.Tags, tag)
	}

	return b
}

// Build returns the corresponding link or an error.
func (b *LinkBuilder) Build() (*Link, error) {
	if b.err != nil {
		return nil, b.err
	}

	return b.link, nil
}
