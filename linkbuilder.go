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

import "github.com/pkg/errors"

const (
	// LinkVersion is the current version of the link encoding.
	LinkVersion = "1.0.0"

	// LinkHashSize is the size of a link hash. In the current version we use
	// SHA-256 so this should be 32.
	LinkHashSize = 32
)

// Link errors.
var (
	ErrMissingProcess  = errors.New("link process is missing")
	ErrMissingMapID    = errors.New("link map id is missing")
	ErrInvalidPriority = errors.New("priority needs to be positive")
	ErrInvalidLinkHash = errors.New("invalid link hash")
)

// LinkBuilder makes it easy to create links that adhere to the ChainScript
// spec.
// It provides valid default values for required fields and allows the user
// to set fields to valid values.
type LinkBuilder struct {
	link *Link
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

// WithParent sets the link's parent, referenced by its hash.
func (b *LinkBuilder) WithParent(linkHash []byte) *LinkBuilder {
	if len(linkHash) != LinkHashSize {
		b.err = ErrInvalidLinkHash
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

// WithStep sets the specific process step represented by the link.
func (b *LinkBuilder) WithStep(step string) *LinkBuilder {
	b.link.Meta.Step = step
	return b
}

// WithTags adds some tags to the link.
func (b *LinkBuilder) WithTags(tags ...string) *LinkBuilder {
	b.link.Meta.Tags = append(b.link.Meta.Tags, tags...)
	return b
}

// Build returns the corresponding link or an error.
func (b *LinkBuilder) Build() (*Link, error) {
	if b.err != nil {
		return nil, b.err
	}

	return b.link, nil
}
