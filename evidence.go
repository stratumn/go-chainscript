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

// NewEvidence creates a new evidence that can be added to a segment.
func NewEvidence(version, backend, provider string, proofData []byte) (*Evidence, error) {
	e := &Evidence{
		Version:  version,
		Backend:  backend,
		Provider: provider,
		Proof:    proofData,
	}

	err := e.Validate()
	if err != nil {
		return nil, err
	}

	return e, nil
}

// Validate that the evidence is well-formed.
// The proof is opaque bytes so it isn't validated here.
func (e *Evidence) Validate() error {
	if len(e.Version) == 0 {
		return ErrMissingVersion
	}

	if len(e.Backend) == 0 {
		return ErrMissingBackend
	}

	if len(e.Provider) == 0 {
		return ErrMissingProvider
	}

	if len(e.Proof) == 0 {
		return ErrMissingProof
	}

	return nil
}
