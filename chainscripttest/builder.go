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

	"github.com/stratumn/go-chainscript"
	"github.com/stretchr/testify/require"
)

// LinkBuilder allows building links easily in tests.
type LinkBuilder struct {
	Link *chainscript.Link
}

// NewLinkBuilder creates a new LinkBuilder.
func NewLinkBuilder(t *testing.T) *LinkBuilder {
	l, err := chainscript.NewLinkBuilder("process", "mapID").Build()
	require.NoError(t, err)

	return &LinkBuilder{Link: l}
}

// Branch uses the provided link as its parent and copies its mapID and process.
func (lb *LinkBuilder) Branch(t *testing.T, parent *chainscript.Link) *LinkBuilder {
	lh, err := parent.Hash()
	require.NoError(t, err)

	lb.Link.Meta.PrevLinkHash = lh
	lb.Link.Meta.MapId = parent.Meta.MapId
	lb.Link.Meta.Process = &chainscript.Process{
		Name:  parent.Meta.Process.Name,
		State: parent.Meta.Process.State,
	}

	return lb
}

// From clones the given link.
func (lb *LinkBuilder) From(t *testing.T, l *chainscript.Link) *LinkBuilder {
	var err error
	lb.Link, err = l.Clone()
	require.NoError(t, err)

	return lb
}

// WithAction fills the link's action.
func (lb *LinkBuilder) WithAction(action string) *LinkBuilder {
	lb.Link.Meta.Action = action
	return lb
}

// WithClientID sets the link's clientID.
func (lb *LinkBuilder) WithClientID(clientID string) *LinkBuilder {
	lb.Link.Meta.ClientId = clientID
	return lb
}

// WithData fills the link's data.
func (lb *LinkBuilder) WithData(t *testing.T, data interface{}) *LinkBuilder {
	err := lb.Link.SetData(data)
	require.NoError(t, err)

	return lb
}

// WithInvalidFields makes the link invalid by setting some fields to invalid
// values.
func (lb *LinkBuilder) WithInvalidFields() *LinkBuilder {
	lb.Link.Meta.Process = nil
	lb.Link.Meta.MapId = ""
	return lb
}

// WithInvalidSignature adds an invalid signature to the link.
func (lb *LinkBuilder) WithInvalidSignature(t *testing.T) *LinkBuilder {
	err := lb.Link.Sign(RandomPrivateKey(t), "")
	require.NoError(t, err)

	lb.Link.Signatures[len(lb.Link.Signatures)-1].Signature = []byte("this is not the signature you're looking for")
	return lb
}

// WithMapID fills the link's mapID.
func (lb *LinkBuilder) WithMapID(mapID string) *LinkBuilder {
	lb.Link.Meta.MapId = mapID
	return lb
}

// WithMetadata sets the link meta.Data field.
func (lb *LinkBuilder) WithMetadata(t *testing.T, metadata interface{}) *LinkBuilder {
	err := lb.Link.SetMetadata(metadata)
	require.NoError(t, err)

	return lb
}

// WithoutParent removes the link's parent (prevLinkHash).
func (lb *LinkBuilder) WithoutParent() *LinkBuilder {
	lb.Link.Meta.PrevLinkHash = nil
	return lb
}

// WithParent fills the link's prevLinkHash with the given parent's hash.
func (lb *LinkBuilder) WithParent(t *testing.T, link *chainscript.Link) *LinkBuilder {
	linkHash, err := link.Hash()
	require.NoError(t, err)

	lb.Link.Meta.PrevLinkHash = linkHash
	return lb
}

// WithParentHash fills the link's prevLinkHash.
func (lb *LinkBuilder) WithParentHash(linkHash []byte) *LinkBuilder {
	lb.Link.Meta.PrevLinkHash = linkHash
	return lb
}

// WithPriority fills the link's priority.
func (lb *LinkBuilder) WithPriority(priority float64) *LinkBuilder {
	lb.Link.Meta.Priority = priority
	return lb
}

// WithProcess fills the link's process.
func (lb *LinkBuilder) WithProcess(process string) *LinkBuilder {
	if lb.Link.Meta.Process == nil {
		lb.Link.Meta.Process = &chainscript.Process{}
	}

	lb.Link.Meta.Process.Name = process
	return lb
}

// WithProcessState fills the link's process state.
func (lb *LinkBuilder) WithProcessState(state string) *LinkBuilder {
	if lb.Link.Meta.Process == nil {
		lb.Link.Meta.Process = &chainscript.Process{}
	}

	lb.Link.Meta.Process.State = state
	return lb
}

// WithRandomData sets random data in most fields of the link.
func (lb *LinkBuilder) WithRandomData() *LinkBuilder {
	lb.Link.Data = RandomBytes(42)
	lb.Link.Meta.Action = RandomString(12)
	lb.Link.Meta.Data = RandomBytes(24)
	lb.Link.Meta.MapId = RandomString(24)
	lb.Link.Meta.PrevLinkHash = RandomHash()
	lb.Link.Meta.Priority = rand.Float64()
	lb.Link.Meta.Process = &chainscript.Process{
		Name:  RandomString(24),
		State: RandomString(24),
	}
	lb.Link.Meta.Step = RandomString(12)
	lb.Link.Meta.Tags = []string{RandomString(12), RandomString(12)}

	return lb
}

// WithRef adds a reference to the link.
func (lb *LinkBuilder) WithRef(t *testing.T, link *chainscript.Link) *LinkBuilder {
	require.NotNil(t, link.Meta.Process)

	refHash, err := link.Hash()
	require.NoError(t, err)

	lb.Link.Meta.Refs = append(lb.Link.Meta.Refs, &chainscript.LinkReference{
		LinkHash: refHash,
		Process:  link.Meta.Process.Name,
	})

	return lb
}

// WithSignature signs the link with a random key.
// Provide an empty payload to sign the whole link.
func (lb *LinkBuilder) WithSignature(t *testing.T, payloadPath string) *LinkBuilder {
	err := lb.Link.Sign(RandomPrivateKey(t), payloadPath)
	require.NoError(t, err)

	return lb
}

// WithSignatureFromKey signs the link with the given key.
// Provide an empty payload to sign the whole link.
func (lb *LinkBuilder) WithSignatureFromKey(t *testing.T, key []byte, payloadPath string) *LinkBuilder {
	err := lb.Link.Sign(key, payloadPath)
	require.NoError(t, err)

	return lb
}

// WithStep fills the link's step.
func (lb *LinkBuilder) WithStep(step string) *LinkBuilder {
	lb.Link.Meta.Step = step
	return lb
}

// WithTag adds a tag to the link.
func (lb *LinkBuilder) WithTag(tag string) *LinkBuilder {
	lb.Link.Meta.Tags = append(lb.Link.Meta.Tags, tag)
	return lb
}

// WithTags replaces the link's tags.
func (lb *LinkBuilder) WithTags(tags ...string) *LinkBuilder {
	lb.Link.Meta.Tags = tags
	return lb
}

// WithVersion sets the link's version.
func (lb *LinkBuilder) WithVersion(version string) *LinkBuilder {
	lb.Link.Version = version
	return lb
}

// Build returns the underlying link.
func (lb *LinkBuilder) Build() *chainscript.Link {
	return lb.Link
}
