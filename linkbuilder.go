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
)

// Link errors.
var (
	ErrMissingProcess = errors.New("link process is missing")
	ErrMissingMapID   = errors.New("link map id is missing")
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

// Build returns the corresponding link or an error.
func (b *LinkBuilder) Build() (*Link, error) {
	if b.err != nil {
		return nil, b.err
	}

	return b.link, nil
}
