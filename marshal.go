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
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

// MarshalSegment marshals using protobuf.
func MarshalSegment(s *Segment) ([]byte, error) {
	return proto.Marshal(s)
}

// UnmarshalSegment unmarshals protobuf bytes.
func UnmarshalSegment(b []byte) (*Segment, error) {
	var unmarshalled Segment
	err := proto.Unmarshal(b, &unmarshalled)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &unmarshalled, nil
}

// MarshalLink marshals using protobuf.
func MarshalLink(l *Link) ([]byte, error) {
	return proto.Marshal(l)
}

// UnmarshalLink unmarshals protobuf bytes.
func UnmarshalLink(b []byte) (*Link, error) {
	var unmarshalled Link
	err := proto.Unmarshal(b, &unmarshalled)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &unmarshalled, nil
}

// MarshalEvidence marshals using protobuf.
func MarshalEvidence(e *Evidence) ([]byte, error) {
	return proto.Marshal(e)
}

// UnmarshalEvidence unmarshals protobuf bytes.
func UnmarshalEvidence(b []byte) (*Evidence, error) {
	var unmarshalled Evidence
	err := proto.Unmarshal(b, &unmarshalled)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &unmarshalled, nil
}
